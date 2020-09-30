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

//	insert department
//	Parameters:
//		userId
//			รหัสผู้้ใช้งาน
//		deptCode
//			โค้ดกรม
//		ministryId
//			รหัสกระทรวง
//		mapTxt
//			ชื่อกรม
//		mapShortTxt
//			ชื่อย่อกรม
//	Return:
//		Department_struct
func PostDepartment(userId int64, deptCode string, ministryId string, mapTxt, mapShortTxt json.RawMessage) (*Department_struct, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return nil, err
	}
	jsonShortText, err := mapShortTxt.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var newId int64
	err = db.QueryRow(sqlInsertDepartment, deptCode, jsonShortText, jsonText, ministryId, userId).Scan(&newId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	data := &Department_struct{Id: newId, DepartmentCode: deptCode, DepartmentShortName: mapShortTxt, DepartmentName: mapTxt}
	return data, nil
}
