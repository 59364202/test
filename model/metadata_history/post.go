// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_history is a model for public.metadata_history table. This table store metadata_history information.
package metadata_history

import (
	//	result "haii.or.th/api/thaiwater30/util/result"
	//	"haii.or.th/api/util/errors"
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/pqx"
)

//	Function for insert data in 'metadta_history' table
//	Parameters:
//		param
//			Struct_MetadataHistory_InputParam
//		userID
//			รหัสผู้ใช้
//func PostMetadataHistory(param *Struct_MetadataHistory_InputParam, userID int64) (*result.Result, error) {
//	//Check 'metadata_id' is not null.
//	if param.MetadataID == 0 {
//		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
//	}
//
//	//Check 'history_description' is not null.
//	if param.Description == "" {
//		return nil, rest.NewError(422, "history_description is not null.", errors.New("parameter 'history_description' is not null."))
//	}
//
//	//Open database
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//
//	//Begin transaction
//	tx, err := db.Begin()
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//	defer tx.Rollback()
//
//	//Loop for insert to 'metadata_history' table
//	err = FncInsertData(tx, param.MetadataID, param.Description, userID)
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//
//	//Commit transaction
//	tx.Commit()
//
//	//Return data
//	return result.Result1("Created Successful."), nil
//}

//	Function for insert data in 'metadta_history' table
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		historyDesc
//			คำอธิบายการแก้ไข
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func FncInsertData(tx *pqx.Tx, metadataID int64, historyDesc string, userID int64) error {
	var _id int64

	//Prepare statement
	statement, err := tx.Prepare(sqlInsert)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters (metadata_id, description, user_id)
	err = statement.QueryRow(metadataID, historyDesc, userID).Scan(&_id)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
