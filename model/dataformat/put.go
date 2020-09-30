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
	"strconv"
)

var sqlUpdateDataformat = ` UPDATE lt_dataformat
						    SET dataformat_name = $2
						      , updated_by = $3
						      , updated_at = NOW()
						      , metadata_method_id = $4
						    WHERE id = $1 `

var sqlUpdateToDeleteDataformat = ` UPDATE lt_dataformat
									SET deleted_by = $2
									  , deleted_at = NOW()
									  , updated_by = $2
									  , updated_at = NOW()
									WHERE id = $1 `

//  update data
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		dataformatId
//			รหัสรูปแบบข้อมูล
//		dataformatName
//			ชื่อรูปแบบข้อมูล
//		metdata_method_id
//			รหัสวิธีการได้มาซึ่งข้อมูล
//	Return:
//		Dataformat_struct
func PutDataformat(userId int64, dataformatId string, dataformatName json.RawMessage, metdata_method_id int64) (*Dataformat_struct, error) {
	//Convert dataformatId type from string to int64
	intDataformatId, err := strconv.ParseInt(dataformatId, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

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

	//Update lt_dataformat table
	err = updateDataformat(tx, userId, dataformatId, dataformatName, inf_metdata_method_id)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	data := &Dataformat_struct{Id: intDataformatId, DataformatName: dataformatName}

	//Return result
	return data, nil
}

//	Update lt_dataformat table
//	Parameters:
//		tx
//			transaction
//		userId
//			รหัสผู้ใช้งาน
//		dataformatId
//			รหัสรูปแบบข้อมูล
//		dataformatName
//			ชื่อรูปแบบข้อมูล
//		metdata_method_id
//			รหัสวิธีการได้มาซึ่งข้อมูล
//	Return:
//		nil, error
func updateDataformat(tx *pqx.Tx, userId int64, dataformatId string, dataformatName json.RawMessage, metdata_method_id interface{}) error {
	//Convert dataformatName to db-json type
	jsonDataformatName, err := dataformatName.MarshalJSON()
	if err != nil {
		return err
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateDataformat)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(dataformatId, jsonDataformatName, userId, metdata_method_id)
	if err != nil {
		return err
	}

	return nil
}

//	Update lt_dataformat table to set 'Delete'
//	Parameters:
//		tx
//			transaction
//		dataformatId
//			รหัสรูปแบบข้อมูล
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		nil, error
func updateToDeleteDataformat(tx *pqx.Tx, dataformatId string, userId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateToDeleteDataformat)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(dataformatId, userId)
	if err != nil {
		return err
	}

	//Return result
	return nil
}
