// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_frequencyunit is a model for public.lt_frequencyunit table. This table store lt_frequencyunit information.
package lt_frequencyunit

import (
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlUpdateFrequencyUnit = ` UPDATE lt_frequencyunit
						       SET frequencyunit_name = $2
						         , convert_minute = $3
						         , updated_by = $4
						         , updated_at = NOW()
						       WHERE id = $1 `

var sqlUpdateFrequencyUnitToDelete = ` UPDATE lt_frequencyunit
									   SET deleted_by = $2
									     , deleted_at = NOW()
									     , updated_by = $2
									     , updated_at = NOW()
									   WHERE id = $1 `

//	update frequencyunit
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		frequencyUnitId
//			รหัสหน่วยความถี่
//		frequencyUnitName
//			ชื่อหน่วยความถี่
//		convertMinute
//			แปลงค่าเป็นนาที
//	Return:
//		FrequencyUnit_struct
func PutFrequencyUnit(userId int64, frequencyUnitId string, frequencyUnitName json.RawMessage, convertMinute string) (*FrequencyUnit_struct, error) {

	//Convert frequencyUnitId type from string to int64
	intFrequencyUnitId, err := strconv.ParseInt(frequencyUnitId, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Convert convertMinute type from string to int64
	intConvertMinute, err := strconv.ParseInt(convertMinute, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Update lt_frequencyunit table
	err = updateFrequencyUnit(tx, frequencyUnitId, frequencyUnitName, convertMinute, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	data := &FrequencyUnit_struct{Id: intFrequencyUnitId, FrequencyUnitName: frequencyUnitName, ConvertMinute: intConvertMinute}

	//Return result
	return data, nil
}

//	Update lt_frequencyunit table
//	Parameters:
//		tx
//			transaction
//		frequencyUnitId
//			รหัสหน่วยความถี่
//		frequencyUnitName
//			ชื่อหน่วยความถี่
//		convertMinute
//			แปลงค่าเป็นนาที
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		nil, error
func updateFrequencyUnit(tx *pqx.Tx, frequencyUnitId string, frequencyUnitName json.RawMessage, convertMinute string, userId int64) error {
	var (
		jsonFrequencyUnitName interface{} = nil
		err                   error
	)

	//Convert agencyName to db-json type
	if frequencyUnitName != nil {
		jsonFrequencyUnitName, err = frequencyUnitName.MarshalJSON()
		if err != nil {
			return err
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateFrequencyUnit)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(frequencyUnitId, jsonFrequencyUnitName, convertMinute, userId)
	if err != nil {
		return err
	}

	return nil
}

//	Update lt_frequencyunit table to set 'Delete'
//	Parameters:
//		tx
//			transaction
//		frequencyUnitId
//			รหัสหน่วยความถี่
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		nil, error
func updateFrequencyUnitToDelete(tx *pqx.Tx, frequencyUnitId string, userId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateFrequencyUnitToDelete)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(frequencyUnitId, userId)
	if err != nil {
		return err
	}

	//Return result
	return nil
}
