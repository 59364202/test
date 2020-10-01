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

//	insert agency
//	Parameters:
//		param
//			โครงสร้างหน่วยงาน ที่ผู้ช้กรอกเข้ามา
// 	Return:
//		Struct_Agency
func InsertAgency(param *Param_Agency) (*Struct_Agency, error) {
	//Convert departmentId type from string to int64
	intDepartmentId, err := strconv.ParseInt(param.DepartmentId, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var _id int64
	//Convert agencyName to db-json type
	jsonAgencyName, err := param.AgencyName.MarshalJSON()
	if err != nil {
		return nil, err
	}
	//Convert agencyshortName to db-json type
	jsonAgencyShortName, err := param.AgencyShortName.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(SQL_InsertAgency, jsonAgencyName, jsonAgencyShortName, intDepartmentId, param.UserId, param.Logo).Scan(&_id)
	if err != nil {
		return nil, err
	}

	data := &Struct_Agency{Agency_shortname: param.AgencyShortName, Department_id: intDepartmentId}
	data.Id = _id
	data.Agency_name = param.AgencyName
	return data, nil
}
