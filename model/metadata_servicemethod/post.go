// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_servicemethod is a model for public.metadata_servicemethod table. This table store metadata_servicemethod information.
package metadata_servicemethod

import (
	//	result "haii.or.th/api/thaiwater30/util/result"
	//	"haii.or.th/api/util/errors"
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/pqx"
)

//	Function for insert data in 'metadata_servicemethod' table
//func PostMetadataServiceMethod(metadataID int64, arrServiceMethodID []int64, userID int64) (*result.Result, error) {
//	//Check 'metadata_id' is not null.
//	if metadataID == 0 {
//		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
//	}
//
//	//Check 'servicemethod_id' is not null.
//	if len(arrServiceMethodID) == 0 {
//		return nil, rest.NewError(422, "servicemethod_id is not null.", errors.New("parameter 'servicemethod_id' is not null."))
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
//	//Loop for insert to 'metadata_servicemethod' table with servicemethod's array.
//	for _, intServiceMethodID := range arrServiceMethodID {
//		err := FncInsertData(tx, metadataID, intServiceMethodID, userID)
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

// 	Function for insert data in 'metadata_servicemethod' table
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		serviceMethodID
//			รหัสวิธีการให้บริการ
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func FncInsertData(tx *pqx.Tx, metadataID int64, serviceMethodID int64, userID int64) error {
	var _id int64

	//Prepare statement
	statement, err := tx.Prepare(sqlInsert)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters (metadata_id, servicemethod_id, user_id)
	err = statement.QueryRow(metadataID, serviceMethodID, userID).Scan(&_id)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
