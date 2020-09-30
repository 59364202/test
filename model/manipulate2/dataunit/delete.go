package dataunit

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlDeleteDataunit = ` DELETE FROM lt_dataunit WHERE id = $1 `

func DeleteDataunit(userId int64, dataunitId int64) (*result.Result, error) {
	//Open database
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

	//Check child table of lt_dataunit
	isHasChild, err := checkDataunitChild(db, dataunitId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return result.Result0(""), nil
	}

	/*//Update to Delete lt_dataunit
	err = updateToDeleteDataUnit(tx, dataunitId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}*/

	//Delete lt_dataunit table
	err = deleteDataunit(tx, dataunitId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return result.Result1("Delete Successful."), nil
}

//Delete lt_dataunit table
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

//Check child table of lt_dataunit
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
