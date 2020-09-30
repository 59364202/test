// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package subcategory is a model for public.subcategory table. This table store subcategory information.
package subcategory

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	soft delete sbucategory
//	Parameters:
//		subcategoryId
//			รหัสหมวดหมู่ย่อย
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		"Delete Successful"
func DeleteSubCategory(subcategoryId string, userId int64) (string, error) {
	hasChild := false
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	row, err := db.Query(sqlCheckChild, subcategoryId)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	for row.Next() {
		hasChild = true
	}

	if hasChild {
		return "", rest.NewError(422, "หมวดหมู่ย่อยได้ถูกใช้โดย บัญชีข้อมูล", nil)
	}

	_, err = db.Exec(sqlUpdateToDeleteSubCategory, subcategoryId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Delete Successful", nil
}
