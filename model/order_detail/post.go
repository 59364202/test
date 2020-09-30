// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"strconv"

	model "haii.or.th/api/server/model"

	"haii.or.th/api/util/pqx"
)

//	inser order_detail
func InsertOrderDetail(OrderDetails []*Param_OrderDetail, tx *pqx.Tx) error {

	// cd dvd , download
	smtDetail, err := tx.Prepare(SQL_InserOrderDetail)
	if err != nil {
		return err
	}
	defer smtDetail.Close()
	// webservice 7 day
	smtDetailService1, err := tx.Prepare(SQL_InserOrderDetail_service1)
	if err != nil {
		return err
	}
	defer smtDetailService1.Close()
	// webservice latest

	smtDetailService4, err := tx.Prepare(SQL_InserOrderDetail_service4)
	if err != nil {
		return err
	}
	defer smtDetailService4.Close()

	for _, detail := range OrderDetails {
		_encrpt, _ := model.GetCipher().EncryptText(strconv.FormatInt(detail.Metadata_Id, 10))
		if detail.Service_Id == 1 { // history 7 day
			_, err = smtDetailService1.Exec(detail.Metadata_Id, detail.Order_Header_Id, detail.Status_Id, detail.Detail_Frequency, detail.Detail_Remark, detail.Detail_Province,
				detail.Detail_Basin, _encrpt, detail.Detail_Source_Result_Itf, detail.Detail_Source_Result_Date_Itf)
		} else if detail.Service_Id == 4 { // latest
			_, err = smtDetailService4.Exec(detail.Metadata_Id, detail.Order_Header_Id, detail.Status_Id, detail.Detail_Frequency, detail.Detail_Remark, detail.Detail_Province,
				detail.Detail_Basin, _encrpt, detail.Detail_Source_Result_Itf, detail.Detail_Source_Result_Date_Itf)
		} else { // download, cd/dvd
			_, err = smtDetail.Exec(detail.Metadata_Id, detail.Order_Header_Id, detail.Status_Id, detail.Service_Id, detail.Detail_Frequency, detail.Detail_Fromdate, detail.Detail_Todate, detail.Detail_Remark,
				detail.Detail_Province, detail.Detail_Basin, _encrpt, detail.Detail_Source_Result_Itf, detail.Detail_Source_Result_Date_Itf)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
