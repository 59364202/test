// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	insert data
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		param
//			parameter ที่หน้าจอส่งมา
//	Return:
//		Inserted Successful
func PostMetadataShow(userId int64, param *Struct_MetadataShow_Param) (string, error) {

	//Try to open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	err = FncInsertData(tx, param.MetadataID, param.MetadataShowSystemID, param.SubCategoryID, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	return "Inserted Successful", nil
}

//	Function for insert data in 'metadata_data_show' table with metadata_id and hydroinfo_id
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
//		userID
//			รหัสผู้ใช้
//	Return:
//		nil, error
func FncInsertData(tx *pqx.Tx, metadataID int64, MetadataShowSystemID int64, SubCategoryID int64, userID int64) error {
	var _id int64

	//Prepare statement
	statement, err := tx.Prepare(sqlInsert)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters
	err = statement.QueryRow(metadataID, MetadataShowSystemID, SubCategoryID, userID).Scan(&_id)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
