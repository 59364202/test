// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_frequency is a model for public.metadata_frequency table. This table store metadata_frequency information.
package metadata_frequency

import (
	//	result "haii.or.th/api/thaiwater30/util/result"
	//	"haii.or.th/api/util/errors"
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/pqx"
)

//  Function for insert data in 'metadata_frequency' table
//func PostMetadataFrequencyMethod(metadataID int64, arrDataFrequency []string, userID int64) (*result.Result, error) {
//	//Check 'metadata_id' is not null.
//	if metadataID == 0 {
//		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
//	}
//
//	//Check 'dataFrequency' is not null.
//	if len(arrDataFrequency) == 0 {
//		return nil, rest.NewError(422, "dataFrequency is not null.", errors.New("parameter 'dataFrequency' is not null."))
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
//	//Loop for insert to 'metadata_frequency' table with servicemethod's array.
//	for _, strDataFrequency := range arrDataFrequency {
//		err := FncInsertData(tx, metadataID, strDataFrequency, userID)
//		if err != nil {
//			return nil, pqx.GetRESTError(err)
//		}
//	}
//
//	//Commit transaction
//	tx.Commit()
//
//	//Return data
//	return result.Result1("Created Successful."), nil
//}

//	Function for insert data in 'metadata_frequency' table
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		dataFrequency
//			ความถี่
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func FncInsertData(tx *pqx.Tx, metadataID int64, dataFrequency string, userID int64) error {
	var _id int64

	//Prepare statement
	statement, err := tx.Prepare(sqlInsert)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters (metadata_id, datafrequency, user_id)
	err = statement.QueryRow(metadataID, dataFrequency, userID).Scan(&_id)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
