// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package category is a model for public.lt_category table. This table store lt_category information.
package lt_category

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	update data
//	Parameters:
//		categoryId
//			รหัสประเภทข้อมูล
//		userId
//			รหัสผู้ใช้
//		mapTxt
//			ชื่อประเภทข้อมูล
//	Return:
//		Update Successful
func PutCategory(categoryId string, userId int64, mapTxt json.RawMessage) (string, error) {
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}

	_, err = db.Exec(sqlUpdateCategory, jsonText, userId, categoryId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Update Successful", nil
}
