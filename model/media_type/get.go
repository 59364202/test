// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package media_type is a model for public.media_type table. This table store media_type information.
package media_type

import (
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

// 	get media type ตามหน่วยงาน
//	Parameters:
//		agencyID
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_MediaType
func getMediaType(agencyID int64) ([]*Struct_MediaType, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data         []*Struct_MediaType
		objMediaType *Struct_MediaType

		_id                sql.NullInt64
		_type_name         sql.NullString
		_subtype_name      sql.NullString
		_type_subtype_name sql.NullString

		_result *sql.Rows
	)

	//Where Condition
	var sqlCmdWhere string = ""
	var arrParam = make([]interface{}, 0)
	if agencyID != 0 {
		arrParam = append(arrParam, agencyID)
		sqlCmdWhere = ` AND EXISTS (SELECT id FROM media m WHERE m.agency_id = $1 AND m.deleted_at = '1970-01-01 07:00:00+07' AND mt.id = m.media_type_id) `
	}

	//Query
	//log.Println(sqlGetMedia + sqlGetMediaOrderBy)
	_result, err = db.Query(sqlGetMediaType+sqlCmdWhere+sqlGetMediaTypeOrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_MediaType, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_type_name, &_subtype_name, &_type_subtype_name)
		if err != nil {
			return nil, err
		}

		//Generate MediaType object
		objMediaType = &Struct_MediaType{}
		objMediaType.Id = _id.Int64
		objMediaType.Type_name = _type_name.String
		objMediaType.Subtype_name = _subtype_name.String
		objMediaType.Type_subtype_name = _type_subtype_name.String

		data = append(data, objMediaType)
	}

	//Return Data
	return data, nil
}

// 	get media type ตามหน่วยงาน
//	Parameters:
//		agencyID
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_MediaType
func GetMediaTypeByAgency(agencyID int64) ([]*Struct_MediaType, error) {
	return getMediaType(agencyID)
}

// 	get media type ทั้งหมด
//	Parameters:
//		agencyID
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_MediaType
func GetAllMediaType() ([]*Struct_MediaType, error) {
	return getMediaType(0)
}
