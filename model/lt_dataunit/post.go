// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_dataunit is a model for public.lt_dataunit table. This table store lt_dataunit information.
package lt_dataunit

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlInsertDataunit = ` INSERT INTO lt_dataunit (dataunit_name, created_by, updated_by) VALUES ($1, $2, $2) RETURNING id `

//	insert data unit
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		dataunitName
//			ชื่อหน่วยข้อมูล
//	Return:
//		Dataunit_struct
func PostDataunit(userId int64, dataunitName json.RawMessage) (*Dataunit_struct, error) {
	//Open DB
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Insert lt_dataunit table
	newId, err := insertDataunit(tx, dataunitName, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx.Commit()

	data := &Dataunit_struct{Id: newId, DataunitName: dataunitName}
	return data, nil
}

//	Insert data to lt_dataunit table
//	Parameters:
//		tx
//			transaction
//		dataunitName
//			ชื่อหน่วยข้อมูล
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		new id
func insertDataunit(tx *pqx.Tx, dataunitName json.RawMessage, userId int64) (int64, error) {
	var (
		_id int64

		jsonDataunitName interface{} = nil
		err              error
	)

	//Convert dataunitName to db-json type
	if dataunitName != nil {
		jsonDataunitName, err = dataunitName.MarshalJSON()
		if err != nil {
			return 0, err
		}
	}

	//Prepare statement
	statement, err := tx.Prepare(sqlInsertDataunit)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute statement
	err = statement.QueryRow(jsonDataunitName, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}
