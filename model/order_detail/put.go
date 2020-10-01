// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"fmt"

	"haii.or.th/api/util/pqx"
)

//	update letterno
func UpdateLetterno(param *Param_OrderLetter) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stm, err := tx.Prepare(SQL_UpdateOrderLetterno)
	if err != nil {
		return err
	}
	defer stm.Close()

	for i, _ := range param.Agency {
		_, err = stm.Exec(param.Order_Header_Id, param.Agency[i], param.LetterNo[i], param.Date[i], param.UserId)
		if err != nil {
			return pqx.GetRESTError(err)
		}
	}

	_, err = db.Exec(SQL_UpdateOrderHeaderStatus, param.UserId, param.Order_Header_Id, 2)
	if err != nil {
		return nil
	}

	tx.Commit()

	return nil
}

//	update letter path
func UpdaeLetterPath(param *Param_OrderLetterPath, user_id int64) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	_, err = db.Exec(SQL_UpdateOrderLetterPath, param.Order_Header_Id, param.Agency_Id, param.Detail_Letterpath, user_id)
	if err != nil {
		return err
	}

	return nil
}

//	update ผลอนุมัติคำขอ จาก หน่วยงานเจ้าของข้อมูล
func UpdateSourceResult(p []*Pram_OrderApprove_Put, user_id int64) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	smt, err := tx.Prepare(SQL_UpdateOrderSourceResult)
	if err != nil {
		return err
	}
	defer smt.Close()

	mapOH := map[int]bool{}
	for _, param := range p {
		_statusId := 4
		if param.Detail_Source_Result == "AP" {
			_statusId = 3
		}
		var _ohID int

		err = smt.QueryRow(param.Id, param.Detail_Source_Result, user_id, _statusId).Scan(&_ohID)
		if err != nil {
			return err
		}
		mapOH[_ohID] = true

	}

	smtOh, err := tx.Prepare(SQL_UpdateOrderHeaderStatus)
	if err != nil {
		return err
	}
	defer smtOh.Close()
	for ohId := range mapOH {
		_, err = smtOh.Exec(user_id, ohId, 4)
		if err != nil {
			return err
		}
	}
	tx.Commit()

	return nil
}

//	update วันหมดอายุตาม order_detail id
func UpdateExpireDate(p *Param_OrderExpireDate_Put) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	_, err = db.Exec(SQL_UpdateOrderExpireDate, p.Detail_Expire_Date, p.Id)
	fmt.Println("SQL_UpdateOrderExpireDate : ", p.Detail_Expire_Date, p.Id)
	if err != nil {
		return err
	}
	return nil
}
