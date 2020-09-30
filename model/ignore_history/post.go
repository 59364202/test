package ignore_history

import (
	"fmt"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
)

//	insert ignore_history
//	Parameters:
//		sqlCmd
//			sql query ที่ต้องใช้
//		arrParam
//			parameter ที่ใช้ร่วมกับ query
//	Return:
//		Insert History Successful
func PostIgnoreHistoryWithSqlCmd(sqlCmd string, arrParam []interface{}) (int64, error) {

	//Check input params
	if sqlCmd == "" {
		return 0, rest.NewError(422, "'sqlCmd' is not null.", errors.New("parameter 'sqlCmd' is not null."))
	}

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Insert ignore_history table
	fmt.Println("1.", tx)
	fmt.Println("2.", sqlCmd)
	fmt.Println(arrParam...)
	newID, err := insertHistoryWithSqlCmd(tx, sqlCmd, arrParam)
	if err != nil {
		return 0, err
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	return newID, nil
}

//	Insert to ignore_history table
//	Parameters:
//		tx
//			Transaction
//		sqlCmd
//			sql query ที่ต้องใช้
//		arrParam
//			parameter ที่ใช้ร่วมกับ query
//	Return:
//		รหัส ignore_history
func insertHistoryWithSqlCmd(tx *pqx.Tx, sqlCmd string, arrParam []interface{}) (int64, error) {
	var _id int64

	// delete statement
	// _, err := tx.Prepare(sqlDeleteHistory + " WHERE id = 901 ")
	// fmt.Println(err)
	// if err != nil {
	// 	return 0, pqx.GetRESTError(err)
	// }

	//Prepare Statement
	//	log.Printf(sqlInsertHistory+sqlCmd+" RETURNING id  ", arrParam...)
	statement, err := tx.Prepare(sqlInsertHistory + sqlCmd + " RETURNING id  ")
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(arrParam...).Scan(&_id)
	fmt.Println(_id)
	fmt.Println(arrParam...)
	fmt.Println(err)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return _id, nil
}
