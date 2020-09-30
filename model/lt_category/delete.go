// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package category is a model for public.lt_category table. This table store lt_category information.
package lt_category

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	delete data
//	Parameters:
//		categoryId
//			รหัสประเภทข้อมูล
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		Delete Successful
func DeleteCategory(categoryId string, userId int64) (string, error) {
	hasChild := false
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	row, err := db.Query(sqlCheckChild, categoryId)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	for row.Next() {
		hasChild = true
	}

	if hasChild {
		return "", rest.NewError(422, "หมวดหมู่หลักได้ถูกใช้โดย บัญชีข้อมูล", nil)
	}

	_, err = db.Exec(sqlUpdateToDeleteCategory, categoryId, userId)
	if err != nil {
		return "", err
	}

	return "Delete Successful", nil
}
