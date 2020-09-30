// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_hydroinfo is a model for public.metadata_hydroinfo table. This table store metadata_hydroinfo information.
package metadata_hydroinfo

import (
	"database/sql"
	"encoding/json"
	model_hydroinfo "haii.or.th/api/thaiwater30/model/hydroinfo"
	"haii.or.th/api/util/pqx"
	//	"log"
)

//	get metadata hydroinfo
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_Hydroinfo
func GetMetadataHydroinfo(metadataID int64) ([]*model_hydroinfo.Struct_Hydroinfo, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	/*//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer tx.Rollback()
	*/

	//Variables
	var (
		data                 []*model_hydroinfo.Struct_Hydroinfo
		objMetadataHydroinfo *model_hydroinfo.Struct_Hydroinfo

		_hydroinfo_id   sql.NullInt64
		_hydroinfo_name sql.NullString

		_result *sql.Rows
	)

	//Query
	//	log.Printf(sqlSelect, metadataID)
	_result, err = db.Query(sqlSelect, metadataID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*model_hydroinfo.Struct_Hydroinfo, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_hydroinfo_id, &_hydroinfo_name)
		if err != nil {
			return nil, err
		}

		if !_hydroinfo_name.Valid || _hydroinfo_name.String == "" {
			_hydroinfo_name.String = "{}"
		}

		objMetadataHydroinfo = &model_hydroinfo.Struct_Hydroinfo{}
		objMetadataHydroinfo.ID = _hydroinfo_id.Int64
		objMetadataHydroinfo.HydroinfoName = json.RawMessage(_hydroinfo_name.String)

		data = append(data, objMetadataHydroinfo)
	}

	//Return Data
	return data, nil
}
