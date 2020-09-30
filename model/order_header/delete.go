// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package order_header is a model for dataservice.order_header table. This table store order_header.
package order_header

import (
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	"haii.or.th/api/util/pqx"
)

//	ยกเลิก order by order header id ( update status เป็น  ยกเลิก)
//	Parameters:
//		order_header_id
//			รหัสคำขอ
//		user_id
//			รหัสผู้ใช้
//	Return:
//		nil, error
func DeleteOrderHeaderById(order_header_id, user_id int64) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//	order header ที่ถูกยกเลิก บังคับให้ order detail result เป็น ไม่อนุมัติ
	stm, err := tx.Prepare(model_order_detail.SQL_UpdateDetailToCancelByOrderHeaderId)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(order_header_id)
	if err != nil {
		return err
	}
	// เปลี่ยนสถานะ order header เป็น ยกเลิก( status_id = 3 )
	stmt, err := tx.Prepare(SQL_UpdateOrderHeaderStatusToCancelByOrderHeaderId)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(order_header_id, user_id)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
