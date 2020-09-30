// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataformat is a model for public.lt_dataformat This table store dataformat information.
package dataformat

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlInsertDataformat = "INSERT INTO lt_dataformat (dataformat_name, created_by, updated_by, created_at, updated_at, metadata_method_id) VALUES ($1, $2, $2, NOW(), NOW(), $3) RETURNING id "

//  insert data
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		dataformatName
//			ชื่อรูปแบบข้อมูล
//		metdata_method_id
//			รหัสวิธีการได้มาซึ่งข้อมูล
//	Return:
//		Dataformat_struct
func PostDataformat(userId int64, dataformatName json.RawMessage, metdata_method_id int64) (*Dataformat_struct, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()
	var inf_metdata_method_id interface{}
	if metdata_method_id != 0 {
		inf_metdata_method_id = metdata_method_id
	}

	//Insert lt_dataformat table
	newId, err := insertDataformat(tx, userId, dataformatName, inf_metdata_method_id)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	data := &Dataformat_struct{Id: newId, DataformatName: dataformatName}
	return data, nil
}

//  Insert to lt_dataformat table
//	Parameters:
//		tx
//			transaction
//		userId
//			รหัสผู้ใช้งาน
//		dataformatName
//			ชื่อรูปแบบข้อมูล
//		metdata_method_id
//			รหัสวิธีการได้มาซึ่งข้อมูล
//	Return:
//		new id
func insertDataformat(tx *pqx.Tx, userId int64, dataformatName json.RawMessage, metadata_method interface{}) (int64, error) {
	var _id int64

	//Convert dataformatName to db-json type
	jsonDataformatName, err := dataformatName.MarshalJSON()
	if err != nil {
		return 0, err
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertDataformat)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	//err = statement.QueryRow(string(jsonDataformatName[:]), userId).Scan(&_id)
	err = statement.QueryRow(jsonDataformatName, userId, metadata_method).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}
