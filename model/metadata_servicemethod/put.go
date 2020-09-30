// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_servicemethod is a model for public.metadata_servicemethod table. This table store metadata_servicemethod information.
package metadata_servicemethod

import (
//	result "haii.or.th/api/thaiwater30/util/result"
//	"haii.or.th/api/util/errors"
//	"haii.or.th/api/util/pqx"
//	"haii.or.th/api/util/rest"
)

//func PutMetadataServiceMethod(metadataID int64, arrServiceMethodID []int64, userID int64) (*result.Result, error) {
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
//		return nil, errors.Repack(err)
//	}
//
//	//Begin Transaction
//	tx, err := db.Begin()
//	if err != nil {
//		return nil, errors.Repack(err)
//	}
//	defer tx.Rollback()
//
//	//Delete data in 'metadata_servicemethod' table with metadata_id
//	err = FncDeleteDataByMetadata(tx, metadataID)
//	if err != nil {
//		return nil, err
//	}
//
//	//Loop for insert to 'metadata_servicemethod' table with servicemethod_id's array.
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
//	//Return result
//	return result.Result1("Updated Successful"), nil
//}
