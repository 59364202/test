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
)

//	update agency
//	Parameters:
//		param
// 	Return:
//		Update Successful ถ้าอัพเดทสำเร็จ
func UpdateAgency(param *Param_Agency) (string, error) {
	//Convert departmentId type from string to int64
	intDepartmentId, err := strconv.ParseInt(param.DepartmentId, 10, 64)
	if err != nil {
		return "", errors.Repack(err)
	}
	//Convert agencyId type from string to int64
	intAgencyId, err := strconv.ParseInt(param.Id, 10, 64)
	if err != nil {
		return "", errors.Repack(err)
	}
	//Convert agencyName to db-json type
	jsonAgencyName, err := param.AgencyName.MarshalJSON()
	if err != nil {
		return "", err
	}
	//Convert agencyShortName to db-json type
	jsonAgencyShortName, err := param.AgencyShortName.MarshalJSON()
	if err != nil {
		return "", err
	}

	db, err := pqx.Open()
	if err != nil {
		return "", err
	}
	_, err = db.Exec(SQL_UpdateAgency, intAgencyId, jsonAgencyName, jsonAgencyShortName, intDepartmentId, param.UserId)
	if err != nil {
		return "", err
	}

	return "Update Successful", nil
}
