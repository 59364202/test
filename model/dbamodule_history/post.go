package dbamodule_history

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//"encoding/json"
)

//	insert create partition table history
//	Parameters:
//		userId
//			user id
//		param
//			ใช้ในส่วน TableName บันทึกว่าของตารางไหน
//			ใช้ในส่วน Year บันทึกว่าเพิ่มของปีไหน
//			ใช้ในส่วน Month บันทึกว่าเพิ่มของเดือนไหน
//	Return:
//		"Insert History Successful"
func PostDBAModuleHistory(userId int64, param *Struct_DBAModuleHistory_InputParam) (*result.Result, error) {

	//Check input params
	if param.TableName == "" {
		return nil, rest.NewError(422, "'table_name' is not null.", errors.New("parameter 'table_name' is not null."))
	}
	if param.Year == "" {
		return nil, rest.NewError(422, "'year' is not null.", errors.New("parameter 'year' is not null."))
	}
	if param.Month == "" {
		return nil, rest.NewError(422, "'month' is not null.", errors.New("parameter 'month' is not null."))
	}

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Insert lt_dataformat table
	/*var strRemarks string
	if (intTypeOfHistory == 0){
		strRemarks = "Delete " + param.TableName + "_y" + param.Year + "m" + param.Month + " table"
	}else{
		strRemarks = "Create " + param.TableName + "_y" + param.Year + "m" + param.Month + " table"
	}*/
	_, err = insertHistory(tx, param.TableName, param.Year, param.Month, param.Remarks, userId)
	if err != nil {
		return nil, err
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	return result.Result1("Insert History Successful."), nil
}

//Insert to dbamodule_history table
func insertHistory(tx *pqx.Tx, tableName string, year string, month string, remarks string, userId int64) (int64, error) {
	var _id int64

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertHistory)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(tableName, year, month, remarks, userId).Scan(&_id)

	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return _id, nil
}
