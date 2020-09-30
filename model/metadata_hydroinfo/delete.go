// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_hydroinfo is a model for public.metadata_hydroinfo table. This table store metadata_hydroinfo information.
package metadata_hydroinfo

import (
	//result "haii.or.th/api/thaiwater30/util/result"
	//"haii.or.th/api/util/errors"
	//"haii.or.th/api/util/rest"
	"haii.or.th/api/util/pqx"
)

/*
//Function for delete data in 'metadta_hydroinfo' table with metadata_id
func DeleteMetadataHydroInfoByMetadataID(metadataID int64, userID int64) (*result.Result, error) {
	//Check 'metadata_id' is not null.
	if (metadataID == 0){
		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
	}

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Delete data
	err = FncUpdateToDeleteDataByMetadata(tx, metadataID, userID)
	if(err != nil){
		return nil, err
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return result.Result1("Delete Successful."), nil
}
*/

//	Function for delete data in 'metadta_hydroinfo' table with metadata_id
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

//	Function for soft delete data in 'metadta_hydroinfo' table with metadata_id
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
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
