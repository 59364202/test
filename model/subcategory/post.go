// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package subcategory is a model for public.subcategory table. This table store subcategory information.
package subcategory

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	insert subcategory
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		categoryId
//			รหัสหมวดหมู่หลัก
//		mapTxt
//			ชื่อหมวดหมู่ย่อย
//	Return:
//		SubCategory_struct
func PostSubCategory(userId int64, categoryId string, mapTxt json.RawMessage) (*SubCategory_struct, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return nil, errors.Repack(err)
	}
	var (
		_id int64
	)
	err = db.QueryRow(sqlInsertSubCategory, categoryId, userId, jsonText).Scan(&_id)
	if err != nil {
		return nil, err
	}

	data := &SubCategory_struct{Id: _id, SubCategoryName: mapTxt}
	return data, nil
}
