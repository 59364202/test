// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_frequency is a model for public.metadata_frequency table. This table store metadata_frequency information.
package metadata_frequency

import (
	//	result "haii.or.th/api/thaiwater30/util/result"
	//	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//	"haii.or.th/api/util/rest"
)

//	Function for delete data in 'metadata_frequency' table with metadata_id
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//		userID
//			รหัสผู้ใช้
//
//func DeleteMetadataFrequencyByMetadataID(metadataID int64, userID int64) (*result.Result, error) {
//	//Check 'metadata_id' is not null.
//	if metadataID == 0 {
//		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
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
//	//Delete data
//	err = FncUpdateToDeleteDataByMetadata(tx, metadataID, userID)
//	if err != nil {
//		return nil, err
//	}
//
//	//Commit transaction
//	tx.Commit()
//
//	//Return result
//	return result.Result1("Delete Successful."), nil
//}

//	Function for delete data in 'metadata_frequency' table with metadata_id
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//	Return:
//		nil, error
func FncDeleteDataByMetadata(tx *pqx.Tx, metadataID int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteByMetadataID)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(metadataID)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

//	Function for soft delete data in 'metadata_frequency' table with metadata_id
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func FncUpdateToDeleteDataByMetadata(tx *pqx.Tx, metadataID int64, userID int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateToDeletedByMetadataID)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(metadataID, userID)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
