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

var sqlUpdateDataunit = ` UPDATE lt_dataunit
						  SET dataunit_name = $2
						    , updated_by = $3
						    , updated_at = NOW()
						  WHERE id = $1 `

var sqlUpdateDataunitToDelete = ` UPDATE lt_dataunit
								  SET deleted_by = $2
								    , deleted_at = NOW()
									, updated_by = $2
									, updated_at = NOW()
								  WHERE id = $1 `

//	update data unit
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		dataunitId
//			รหัสหน่วยข้อมูล
//		dataunitName
//			ชื่อหน่วยข้อมูล
//	Return:
//		Dataunit_struct
func PutDataunit(userId int64, dataunitId int64, dataunitName json.RawMessage) (*Dataunit_struct, error) {
	//Open DB
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

	//Update lt_dataunit table
	err = updateDataunit(tx, dataunitId, dataunitName, userId)
	if err != nil {
		return nil, err
	}

	//Commit transaction
	tx.Commit()

	//Return Result
	data := &Dataunit_struct{Id: dataunitId, DataunitName: dataunitName}
	return data, nil
}

//	Update data to lt_dataunit table
//	Parameters:
//		tx
//			transaction
//		dataunitId
//			รหัสหน่วยข้อมูล
//		dataunitName
//			ชื่อหน่วยข้อมูล
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		nil, error
func updateDataunit(tx *pqx.Tx, dataunitId int64, dataunitName json.RawMessage, userId int64) error {
	var (
		jsonDataunitName interface{} = nil
		err              error
	)

	//Convert dataunitName to db-json type
	if dataunitName != nil {
		jsonDataunitName, err = dataunitName.MarshalJSON()
		if err != nil {
			return err
		}
	}

	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateDataunit)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(dataunitId, jsonDataunitName, userId)
	if err != nil {
		return err
	}

	return nil
}

//Update lt_dataunit to set 'Delete'
func updateDataunitToDelete(tx *pqx.Tx, dataunitId int64, userId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateDataunitToDelete)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(dataunitId, userId)
	if err != nil {
		return err
	}

	return nil
}
