// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_history is a model for public.metadata_history table. This table store metadata_history information.
package metadata_history

import (
	"database/sql"
	model_setting "haii.or.th/api/server/model/setting"
	model_user "haii.or.th/api/server/model/user"
	"haii.or.th/api/thaiwater30/util/sqltime"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
)

//	get netadada hustiry
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_MetadataHistory
func GetMetadataHistory(metadataID int64) ([]*Struct_MetadataHistory, error) {
	//Check 'metadata_id' is not null.
	if metadataID == 0 {
		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	/*//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer tx.Rollback()*/

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Variables
	var (
		data               []*Struct_MetadataHistory
		objMetadataHistory *Struct_MetadataHistory

		_id                  sql.NullInt64
		_history_description sql.NullString
		_history_datetime    sqltime.NullTime
		_user_id             sql.NullInt64
		_user_fullname       sql.NullString

		_result *sql.Rows
	)

	//Query
	//	log.Printf(sqlSelect, metadataID)
	_result, err = db.Query(sqlSelect, metadataID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_MetadataHistory, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_history_datetime, &_history_description, &_user_id, &_user_fullname)
		if err != nil {
			return nil, err
		}

		objMetadataHistory = &Struct_MetadataHistory{}
		objMetadataHistory.ID = _id.Int64
		objMetadataHistory.Description = _history_description.String
		objMetadataHistory.UserFullname = _user_fullname.String
		objMetadataHistory.Datetime = _history_datetime.Time.Format(strDatetimeFormat)

		objMetadataHistory.User = &model_user.User{}
		objMetadataHistory.User.ID = _user_id.Int64
		objMetadataHistory.User.FullName = _user_fullname.String

		data = append(data, objMetadataHistory)
	}

	//Return Data
	return data, nil
}
