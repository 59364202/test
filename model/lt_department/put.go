// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_department is a model for public.lt_department table. This table store lt_department information.
package lt_department

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	update department
//	Parameters:
//		departmentId
//			รหัสกรม
//		departmentCode
//			โค้ดกรม
//		userId
//			ชื่อผู้ใช้งาน
//		ministryId
//			รหัสกระทรวง
//		mapTxt
//			ชื่อกรม
//		mapShortTxt
//			ชื่อย่อกรม
//	Return:
//		Update Successful
func PutDepartment(departmentId, departmentCode string, userId int64, ministryId string, mapTxt, mapShortTxt json.RawMessage) (string, error) {
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}
	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}
	jsonShortText, err := mapShortTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}
	_, err = db.Exec(sqlUpdateDepartment, departmentCode, jsonShortText, jsonText, ministryId, userId, departmentId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Update Successful", nil
}
