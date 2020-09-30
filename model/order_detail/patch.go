// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"strconv"

	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	model "haii.or.th/api/server/model"
)

//	Regenerate e_id
func RegenerateKey(pId int64) (string, error) {
	if pId <= 0 {
		return "", rest.NewError(422, "invalid id", nil)
	}
	strId := strconv.FormatInt(pId, 10)
	newEId, err := model.GetCipher().EncryptText(strId)
	if err != nil {
		return "", rest.NewError(422, "invalid id", err)
	}

	db, err := pqx.Open()
	if err != nil {
		return "", err
	}

	_, err = db.Exec(SQL_UpdateEId, newEId, strId)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	return newEId, nil
}

//	ปรับ is_enabled ของ order_detail เป็น true (เปิดการใช้งาน)
func EnableOrderDetail(pId int64) error {
	if pId <= 0 {
		return rest.NewError(422, "invalid id", nil)
	}
	db, err := pqx.Open()
	if err != nil {
		return err
	}

	_, err = db.Exec(SQL_UpdateIsEnable, true, pId)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	return nil
}
