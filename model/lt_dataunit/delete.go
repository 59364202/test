// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_dataunit is a model for public.lt_dataunit table. This table store lt_dataunit information.
package lt_dataunit

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlDeleteDataunit = ` DELETE FROM lt_dataunit WHERE id = $1 `

//	delete data
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		dataunitId
//			รหัสหน่วยข้อมูล
//	Return:
//		Delete Successful
func DeleteDataunit(userId int64, dataunitId int64) (string, error) {
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

	//Check child table of lt_dataunit
	isHasChild, err := checkDataunitChild(db, dataunitId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return "", errors.New("Can't delete this data. It's has been used.")
	}

	/*//Update to Delete lt_dataunit
	err = updateToDeleteDataUnit(tx, dataunitId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}*/

	//Delete lt_dataunit table
	err = deleteDataunit(tx, dataunitId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful", nil
}

//	Delete data in lt_dataunit table
//	Parameters:
//		tx
//			transaction
//		dataunitId
//			รหัสหน่วยข้อมูล
//	Return:
//		nil, error
func deleteDataunit(tx *pqx.Tx, dataunitId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteDataunit)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(dataunitId)
	if err != nil {
		return err
	}

	return nil
}

//	Check child table of lt_dataunit
//	Parameters:
//		db
//			database connection
//		dataunitId
//			รหัสหน่วยข้อมูล
//	Return:
//		true, false
func checkDataunitChild(db *pqx.DB, dataunitId int64) (bool, error) {
	//Set SQL for check child
	var sqlCheckChild string = ` SELECT dataunit_id FROM metadata WHERE dataunit_id = $1 LIMIT 1 `

	//Set default of return value
	var isHasChild bool = false

	//Query statement with parameters
	row, err := db.Query(sqlCheckChild, dataunitId)
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
