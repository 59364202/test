// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package order_header is a model for dataservice.order_header table. This table store order_header.
package order_header

import (
	//	"haii.or.th/api/thaiwater3nil/util/result"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"

	"haii.or.th/api/util/pqx"
)

//	insert order
//	Parameters:
//		Order_Header
//			Param_OrderHeader
//	Return:
//		Param_OrderHeader
func InsertOrder(Order_Header *Param_OrderHeader) (*Param_OrderHeader, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	statment, err := tx.Prepare(SQL_InsertOrderHeader)
	if err != nil {
		return nil, err
	}
	defer statment.Close()

	var _statusId = 4
	if Order_Header.Order_Forexternal {
		// สำหรับบุคคลภายนอก status เป็น รับคำขอ
		_statusId = 1
	}
	Order_Header.Order_Quality = int64(len(Order_Header.OrderDetail))
	err = statment.QueryRow(Order_Header.User_Id, Order_Header.Order_Quality, Order_Header.Order_Purpose, Order_Header.Order_Forexternal, _statusId).Scan(&Order_Header.Id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for _, detail := range Order_Header.OrderDetail {
		detail.Order_Header_Id = Order_Header.Id
		if !Order_Header.Order_Forexternal {
			// สำหรับหน่วยงาน ตั้ง result เป็น approve
			detail.Detail_Source_Result_Itf = "AP"
			detail.Detail_Source_Result_Date_Itf = "NOW()"
			detail.Status_Id = 3
		} else {
			// สำหรับบุคคลภายนอก result เป็น null
			detail.Detail_Source_Result_Itf = nil
			detail.Detail_Source_Result_Date_Itf = nil
			detail.Status_Id = 5
		}
	}

	err = model_order_detail.InsertOrderDetail(Order_Header.OrderDetail, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return Order_Header, nil
}
