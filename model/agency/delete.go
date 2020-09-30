// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agency is a model for public.agency table. This table store agency.
package agency

import (
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	delete agency (soft delete)
//	Parameters:
//		param
//			ใช้แค่ส่วนของ id
// 	Return:
//		Delete Successful ถ้าลบสำเร็จ(soft delete)
func DeleteAgency(param *Param_Agency) (string, error) {
	//Convert agencyId type from string to int64
	intAgencyId, err := strconv.ParseInt(param.Id, 10, 64)
	if err != nil {
		return "", errors.Repack(err)
	}
	//Check child table of agency
	isHasChild, err := checkAgencyChild(intAgencyId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return "", rest.NewError(422, "ไม่สามารถลบหน่วยงานนี้ได้ เนื่องจากถูกใช้งานอยู่", nil)
	}

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	_, err = db.Exec(SQL_UpdateAgencyToDelete, intAgencyId, param.UserId)
	if err != nil {
		return "", err
	}

	//Return result
	return "Delete Successful.", nil
}

//	Check child table of agency
//	เช็คว่าหน่วยงานนี้มีบัญชีข้อมูลใช้อยู่
//	Parameters:
//		agencyId
//			agency id
// 	Return:
//		true ถ้ามีบัญชีข้อมูลใช้อยู่
func checkAgencyChild(agencyId int64) (bool, error) {

	//Set default of return value
	var isHasChild bool = false

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return isHasChild, errors.Repack(err)
	}
	//Query statement with parameters
	row, err := db.Query(SQL_CheckChild, agencyId)
	if err != nil {
		return isHasChild, err
	}

	//Check child
	for row.Next() {
		isHasChild = true
	}

	//Return result
	return isHasChild, nil
}
