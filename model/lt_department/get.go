// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_department is a model for public.lt_department table. This table store lt_department information.
package lt_department

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	get data
//	Parameters:
//		ministryId
//			[]รหัสกระทรวง
//	Return:
//		[]Department_struct
func GetDepartment(ministryId []int) ([]*Department_struct, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		data       []*Department_struct
		department *Department_struct

		_id            sql.NullInt64
		_deptcode      sql.NullString
		_shorttext     sql.NullString
		_text          sql.NullString
		_ministry_id   sql.NullInt64
		_ministry_text sql.NullString
	)

	var _sql = sqlGetDepartment + sqlGetDepartmentWhereDeletedIsNull
	for i, v := range ministryId {
		if i == 0 {
			_sql += " AND d.ministry_id IN ("
		} else {
			_sql += ","
		}
		_sql += strconv.Itoa(v)
	}
	if len(ministryId) != 0 {
		_sql += ") "
	}

	_result, err := db.Query(_sql + sqlGetDepartmentOrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Department_struct, 0)
	// loop scan data
	for _result.Next() {
		err := _result.Scan(&_id, &_deptcode, &_shorttext, &_text, &_ministry_id, &_ministry_text)
		if err != nil {
			return nil, err
		}
		if !_shorttext.Valid || _shorttext.String == "" {
			_shorttext.String = "{}"
		}
		if !_text.Valid || _text.String == "" {
			_text.String = "{}"
		}
		if !_ministry_text.Valid || _ministry_text.String == "" {
			_ministry_text.String = "{}"
		}

		department = &Department_struct{Id: _id.Int64, DepartmentCode: _deptcode.String}
		department.DepartmentShortName = json.RawMessage(_shorttext.String)
		department.DepartmentName = json.RawMessage(_text.String)
		department.MinistryId = _ministry_id.Int64
		department.MinistryName = json.RawMessage(_ministry_text.String)

		data = append(data, department)
	}

	return data, nil
}

//	get data
//	Parameters:
//		ministry_id
//			รหัสกระทรวง
//	Return:
//		[]Struct_Department
func getDepartment(ministry_id int64) ([]*Struct_Department, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		strSql     string
		data       []*Struct_Department
		department *Struct_Department

		_id              int64
		_department_name sql.NullString
	)
	strSql = SQL_selectDepartment
	if ministry_id != 0 {
		strSql += " WHERE ministry_id = " + strconv.FormatInt(ministry_id, 10)
	}

	row, err := db.Query(strSql)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_department_name)
		if err != nil {
			return nil, err
		}

		if !_department_name.Valid || _department_name.String == "" {
			_department_name.String = "{}"
		}

		department = &Struct_Department{Id: _id, Department_Name: json.RawMessage(_department_name.String)}

		data = append(data, department)
	}

	return data, nil
}

//	get all department
//	Return:
//		[]Struct_Department
func GetAllDepartment() ([]*Struct_Department, error) {
	return getDepartment(0)
}

//	get department from ministy id
//	Parameters:
//		ministry_id
//			รหัสกระทรวง
//	Return:
//		[]Struct_Department
func GetDepartmentFromMinistryId(ministry_id int64) ([]*Struct_Department, error) {
	return getDepartment(ministry_id)
}
