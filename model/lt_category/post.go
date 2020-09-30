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

//	insert data
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		mapTxt
//			ชื่อประเภทข้อมูล
//	Return:
//		Struct_category
func PostCategory(userId int64, mapTxt json.RawMessage) (*Struct_category, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var newId int64
	err = db.QueryRow(sqlInsertCategory, jsonText, userId).Scan(&newId)
	if err != nil {
		return nil, err
	}

	data := &Struct_category{Id: newId, Category_name: mapTxt}
	return data, nil
}
