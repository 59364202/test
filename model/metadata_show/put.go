// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	Function for insert data in 'metadata_data_show' table with metadata_id, metadata_show_system_id, subcategory_id
//	Parameters:
//		tx
//			transaction
//		metadataID
//			รหัสบัญชีข้อมูล
//		MetadataShowSystemID
//			รหัสระบบที่นำข้อมูลไปแสดง
//		SubCategoryID
//			รหัสส่วนที่นำข้อมูลไปแสดง เช่น แสดงผลที่ web thaiwater30 ส่วนฝน
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func PutMetadataShow(param *Struct_MetadataShow_Param, userID int64) (*result.Result, error) {
	//	Check 'metadata_id' is not null.
	if param.MetadataID == 0 {
		return nil, rest.NewError(422, "metadata_id is not null.", errors.New("parameter 'metadata_id' is not null."))
	}

	if param.MetadataShowSystemID == 0 {
		return nil, rest.NewError(422, "metadata_show_system_id is not null.", errors.New("parameter 'metadata_show_system_id' is not null."))
	}

	if param.SubCategoryID == 0 {
		return nil, rest.NewError(422, "subcategory_id is not null.", errors.New("parameter 'subcategory_id' is not null."))
	}

	//	Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//	Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//	Delete data in 'metadata_data_show' table with metadata_id
	//	err = FncDeleteDataByMetadata(tx, param.MetadataID)
	//	if err != nil {
	//		return nil, err
	//	}

	err = FncInsertData(tx, param.MetadataID, param.MetadataShowSystemID, param.SubCategoryID, userID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//	Commit transaction
	tx.Commit()

	//	Return result
	return result.Result1("Updated Successful"), nil
}
