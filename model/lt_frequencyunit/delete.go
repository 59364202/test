// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_frequencyunit is a model for public.lt_frequencyunit table. This table store lt_frequencyunit information.
package lt_frequencyunit

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlDeleteFrequencyUnit = ` DELETE FROM lt_frequencyunit WHERE id = $1 `

//	delete data
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		frequencyUnitId
//			รหัสหน่วยความถี่
//	Return:
//		Delete Successful
func DeleteFrequencyUnit(userId int64, frequencyUnitId string) (string, error) {
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

	//Check child table of lt_frequencyunit
	isHasChild, err := checkFrequencyUnitChild(db, frequencyUnitId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return "", errors.New("Can't delete this data. It's has been used.")
	}

	//Delete lt_frequencyunit table
	err = deleteFrequencyUnit(tx, frequencyUnitId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful", nil
}

//	Check child table of lt_frequencyunit
//	Parameters:
//		db
//			database connection
//		frequencyUnitId
//			รหัสหน่วยความถี่
//	Return:
//		true, false
func checkFrequencyUnitChild(db *pqx.DB, frequencyUnitId string) (bool, error) {
	//
	//	//Set SQL for check child
	//	var sqlCheckChild string = ` SELECT dataformat_id FROM metadata WHERE dataformat_id = $1 LIMIT 1 `
	//
	//	//Set default of return value
	//	var isHasChild bool = false
	//
	//	//Query statement with parameters
	//	row, err := db.Query(sqlCheckChild, dataformatId)
	//	if err != nil {
	//		return isHasChild, err
	//	}
	//
	//	//Check child
	//	for row.Next() {
	//		isHasChild = true
	//	}
	//
	//	//Return result
	//	return isHasChild, nil

	return false, nil
}

//	Delete lt_frequencyunit table
//	Parameters:
//		tx
//			transaction
//		frequencyUnitId
//			รหัสหน่วยความถี่
//	Return:
//		nil, error
func deleteFrequencyUnit(tx *pqx.Tx, frequencyUnitId string) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteFrequencyUnit)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(frequencyUnitId)
	if err != nil {
		return err
	}

	return nil
}
