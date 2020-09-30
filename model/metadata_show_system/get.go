// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show_system

// Package metadata_show_system is a model for public.lt_metadata_show_system table ระบบแสดงข้อมูล  เช่น คลังข้อมูลน้ำฯ,เว็บจังหวัด

import (
	"database/sql"
	"encoding/json"

	"haii.or.th/api/util/pqx"
)

//	get metadata show system
//	Return:
//		Array Struct_MetadataShowSystem
func GetMetadataShowSystem() ([]*Struct_MetadataShowSystem, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Variables
	var (
		data []*Struct_MetadataShowSystem
		obj  *Struct_MetadataShowSystem

		_id                   sql.NullInt64
		_metadata_show_system sql.NullString
	)

	//Query
	_result, err := db.Query(sqlSelect)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_MetadataShowSystem, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_metadata_show_system)
		if err != nil {
			return nil, err
		}

		if !_metadata_show_system.Valid || _metadata_show_system.String == "" {
			_metadata_show_system.String = "{}"
		}

		obj = &Struct_MetadataShowSystem{}
		obj.ID = _id.Int64
		obj.MetadataShowSystem = json.RawMessage(_metadata_show_system.String)

		data = append(data, obj)
	}
	return data, nil
}
