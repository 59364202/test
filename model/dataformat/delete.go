// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataformat is a model for public.lt_dataformat This table store dataformat information.
package dataformat

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	delete data (soft delete)
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		dataformatId
//			รหัสรูปแบบข้อมูล
//	Return:
//		Delete Successful
func DeleteDataformat(userId int64, dataformatId string) (string, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	//Check child table of lt_dataformat
	isHasChild, err := checkDataformatChild(db, dataformatId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return "", nil
	}

	//Update to Delete agency table
	err = updateToDeleteDataformat(tx, dataformatId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful", nil
}

//  Check child table of lt_dataformat
//	Parameters:
//		db
//			sql.DB
//		dataformatId
//			รหัสรูปแบบข้อมูล
//	Return:
//		true
func checkDataformatChild(db *pqx.DB, dataformatId string) (bool, error) {
	//Set SQL for check child
	var sqlCheckChild string = ` SELECT dataformat_id FROM metadata WHERE dataformat_id = $1 LIMIT 1 `

	//Set default of return value
	var isHasChild bool = false

	//Query statement with parameters
	row, err := db.Query(sqlCheckChild, dataformatId)
	if err != nil {
		return isHasChild, err
	}

	//Check child
	for row.Next() {
		isHasChild = true
	}

	//Return result
	return isHasChild, nil
}
