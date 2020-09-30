// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for public.metadata table. This table store metadata information.
package metadata

import (
	model_server "haii.or.th/api/server/model"
	model_metadata_frequency "haii.or.th/api/thaiwater30/model/metadata_frequency"
	model_metadata_hydroinfo "haii.or.th/api/thaiwater30/model/metadata_hydroinfo"
	model_metadata_servicemethod "haii.or.th/api/thaiwater30/model/metadata_servicemethod"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
	"strconv"
)

//	delete metadata (soft delete)
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//		userID
//			รหัสผู้ใช้
//	Return:
//		Delete Successful
func DeleteMetadata(metadataID string, userID int64) (string, error) {
	//Check 'metadata_id' is not null.
	if metadataID == "" {
		return "", rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
	}

	//Decrypt metadata_id
	strMetadataID, err := model_server.GetCipher().DecryptText(metadataID)
	if err != nil {
		return "", rest.NewError(422, "Invalid metadata_id", errors.New("Invalid metadata_id."))
	}
	intMetadataID, err := strconv.ParseInt(strMetadataID, 10, 64)
	if err != nil {
		return "", rest.NewError(422, "metadata_id is not a number.", err)
	}

	//	log.Println(metadataID)
	//	log.Println(intMetadataID)

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Delete data
	err = fncUpdateToDeleteByMetadata(tx, intMetadataID, userID)
	if err != nil {
		return "", err
	}

	//Update metadata_hydroinfo table
	err = model_metadata_hydroinfo.FncUpdateToDeleteDataByMetadata(tx, intMetadataID, userID)
	if err != nil {
		return "", err
	}

	//Delete metadata_servicemethod table
	err = model_metadata_servicemethod.FncUpdateToDeleteDataByMetadata(tx, intMetadataID, userID)
	if err != nil {
		return "", err
	}

	//Update metadata_frequency table
	err = model_metadata_frequency.FncUpdateToDeleteDataByMetadata(tx, intMetadataID, userID)
	if err != nil {
		return "", err
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful", nil
}

//	delete data in table metadata (soft delete)
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func fncUpdateToDeleteByMetadata(tx *pqx.Tx, metadataID int64, userID int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateToDeleteMetadata)
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
