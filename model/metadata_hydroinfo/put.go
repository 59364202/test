// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_hydroinfo is a model for public.metadata_hydroinfo table. This table store metadata_hydroinfo information.
package metadata_hydroinfo

import (
//	result "haii.or.th/api/thaiwater30/util/result"
//	"haii.or.th/api/util/errors"
//	"haii.or.th/api/util/pqx"
//	"haii.or.th/api/util/rest"
)

//	Function for insert data in 'metadta_hydroinfo' table with metadata_id and hydroinfo_id
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		hydroinfoId
//			รหัสกลุ่มข้อมูล9ด้าน
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
//func PutMetadataHydroinfo(param *Struct_MetadataHydroinfo_InputParam, userID int64) (*result.Result, error) {
//	//Check 'metadata_id' is not null.
//	if param.MetadataID == 0 {
//		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
//	}
//
//	//Check 'hydroinfo_id' is not null.
//	if len(param.ListHydroInfoID) == 0 {
//		return nil, rest.NewError(422, "hydroinfo_id is not null.", errors.New("parameter 'hydroinfo_id' is not null."))
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
//	//Delete data in 'metadata_hydroinfo' table with metadata_id
//	err = FncDeleteDataByMetadata(tx, param.MetadataID)
//	if err != nil {
//		return nil, err
//	}
//
//	//Loop for insert to 'metadata_hydroinfo' table with hydroinfo_id's array.
//	for _, intHydroInfo := range param.ListHydroInfoID {
//		err := FncInsertData(tx, param.MetadataID, intHydroInfo, userID)
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
