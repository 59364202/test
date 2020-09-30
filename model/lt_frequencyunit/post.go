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

var sqlInsertFrequencyUnit = "INSERT INTO lt_frequencyunit (frequencyunit_name, convert_minute, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $3, NOW(), NOW()) RETURNING id "

//	insert frequencyunit
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		frequencyUnitName
//			ชื่อหน่วยความถี่
//		convertMinute
//			แปลงค่าเป็นนาที
//	Return:
//		FrequencyUnit_struct
func PostFrequencyUnit(userId int64, frequencyUnitName json.RawMessage, convertMinute string) (*FrequencyUnit_struct, error) {
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

	//Insert lt_frequencyunit table
	newId, err := insertFrequencyUnit(tx, frequencyUnitName, convertMinute, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	data := &FrequencyUnit_struct{Id: newId, FrequencyUnitName: frequencyUnitName, ConvertMinute: intConvertMinute}
	return data, nil
}

//	Insert to lt_frequencyunit table
//	Parameters:
//		tx
//			transaction
//		frequencyUnitName
//			ชื่อหน่วยความถี่
//		convertMinute
//			แปลงค่าเป็นนาที
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		new id
func insertFrequencyUnit(tx *pqx.Tx, frequencyUnitName json.RawMessage, convertMinute string, userId int64) (int64, error) {
	var (
		_id int64

		jsonFrequencyUnitName interface{} = nil
		err                   error
	)

	//Convert frequencyUnitName to db-json type
	if frequencyUnitName != nil {
		jsonFrequencyUnitName, err = frequencyUnitName.MarshalJSON()
		if err != nil {
			return 0, err
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertFrequencyUnit)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(jsonFrequencyUnitName, convertMinute, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}
