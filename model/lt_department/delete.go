// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_department is a model for public.lt_department table. This table store lt_department information.
package lt_department

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	delete data (soft delete)
//	Parameters:
//		departmentId
//			รหัสกรม
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		Delete Successful
func DeleteDepartment(departmentId string, userId int64) (string, error) {
	hasChild := false
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	row, err := db.Query(sqlCheckChild, departmentId)
	if err != nil {
		return "", errors.Repack(err)
	}
	for row.Next() {
		hasChild = true
	}
	if hasChild {
		return "", rest.NewError(422, "กรมได้ถูกใช้โดย บัญชีข้อมูล", nil)
	}
	_, err = db.Exec(sqlUpdateToDeleteDepartment, departmentId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Delete Successful", nil
}
