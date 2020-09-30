// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package order_status is a model for dataservice.order_status table. This table store order_status information.
package order_status

import (
	"haii.or.th/api/util/pqx"
)

//	get all order status
//	Return:
//		Array Struct_OrderStatus
func GetOrderStatusAll() ([]*Struct_OrderStatus, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		data        []*Struct_OrderStatus
		orderStatus *Struct_OrderStatus

		_id     int64
		_status string
	)
	row, err := db.Query(SQL_selectAllStatus)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for row.Next() {
		err = row.Scan(&_id, &_status)
		if err != nil {
			return nil, err
		}

		orderStatus = &Struct_OrderStatus{Id: _id, Order_Status: _status}

		data = append(data, orderStatus)
	}

	return data, nil
}
