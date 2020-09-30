// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_subcategory is a model for public.lt_subcategory table. This table store lt_subcategory information.
package lt_subcategory

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	update data
//	Parameters:
//		subCategoryId
//			รหัสหมวดหมู่ย่อย
//		userId
//			รหัสผู้ใช้งาน
//		categoryId
//			รหัสหมวดข้อมูลหลัก
//		mapTxt
//			ชื่อหมวดหมู่ย่อย
//	Return:
//		Update Successful
func PutSubCategory(subCategoryId string, userId int64, categoryId string, mapTxt json.RawMessage) (string, error) {
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}

	_, err = db.Exec(sqlUpdateSubCategory, categoryId, userId, jsonText, subCategoryId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Update Successful", nil
}
