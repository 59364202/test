// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show

// Package metadata_data_show is a model for public.metadata_data_show table. เก็บข้อมูลว่า เอาข้อมูลไปแสดงที่ระบบไหนบ้าง

import (
	"database/sql"
	//	"encoding/json"

	"haii.or.th/api/util/pqx"
)

//	get metadata show system
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_MetadataShow_Param
func GetMetadataShow() ([]*Struct_MetadataShow_Param, error) {
	//	func GetMetadataDataShow(metadataID int64) ([]*Struct_MetadataShow_Param, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Variables
	var (
		data            []*Struct_MetadataShow_Param
		objMetadataShow *Struct_MetadataShow_Param

		_metadataShow_id         sql.NullInt64
		_metadata_id             sql.NullInt64
		_metadata_name           sql.NullString
		_agency_id               sql.NullInt64
		_agency_name             sql.NullString
		_metadataShowSystem_id   sql.NullInt64
		_metadataShowSystem_name sql.NullString
		_subCategory_id          sql.NullInt64
		_subCategory_name        sql.NullString
		_connection_format       sql.NullString
		_metadata_method         sql.NullString

		_result *sql.Rows
	)

	//Query
	//	log.Printf(sqlSelect, metadataID)
	//	_result, err = db.Query(sqlSelect, metadataID)
	_result, err = db.Query(sqlSelect)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_MetadataShow_Param, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_metadataShow_id, &_metadata_id, &_metadata_name, &_agency_id, &_agency_name, &_metadataShowSystem_id, &_metadataShowSystem_name, &_subCategory_id, &_subCategory_name, &_connection_format, &_metadata_method)
		if err != nil {
			return nil, err
		}

		objMetadataShow = &Struct_MetadataShow_Param{}
		objMetadataShow.ID = _metadataShow_id.Int64
		objMetadataShow.MetadataID = _metadata_id.Int64
		objMetadataShow.MetadataName = _metadata_name.String
		objMetadataShow.AgencyID = _agency_id.Int64
		objMetadataShow.AgencyName = _agency_name.String
		objMetadataShow.MetadataShowSystemID = _metadataShowSystem_id.Int64
		objMetadataShow.MetadataShowSystemName = _metadataShowSystem_name.String
		objMetadataShow.SubCategoryID = _subCategory_id.Int64
		objMetadataShow.SubCategoryName = _subCategory_name.String
		objMetadataShow.ConnectionFormat = _connection_format.String
		objMetadataShow.MetadataMethod = _metadata_method.String

		data = append(data, objMetadataShow)
	}

	//Return Data
	return data, nil
}
