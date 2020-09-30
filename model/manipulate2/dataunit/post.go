package dataunit

import (
	"encoding/json"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlInsertDataunit = ` INSERT INTO lt_dataunit (dataunit_name, created_by, updated_by) VALUES ($1, $2, $2) RETURNING id `

func PostDataunit(userId int64, dataunitName json.RawMessage) (*result.Result, error) {
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
	return result.Result1(data), nil
}

//Insert lt_dataunit table
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
