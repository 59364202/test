// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	delete order_detail ตาม order_header_id
func DeleteOrderDetailByOrderHeaderId(order_header_id int64) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	_, err = db.Exec(SQL_DeleteOrderDetailByOrderHeaderId, order_header_id)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

//	ปรับ is_enabled ของ order_detail เป็น false (ปิดการใช้งาน)
func DisableOrderDetail(param *Param) error {
	if param.Id <= 0 {
		return rest.NewError(422, "invalid id", nil)
	}
	db, err := pqx.Open()
	if err != nil {
		return err
	}
	_, err = db.Exec(SQL_UpdateIsEnable, false, param.Id)
	if err != nil {
		return err
	}

	return nil
}
