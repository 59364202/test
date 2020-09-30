package rainfall

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"strings"
)

//	Update rainfall table to set 'Delete'
//	Parameters:
//		param
//			Rainfall_InputParam
//		userId
//			รหัสผู้ใช้
//	Return:
//		"Delete Successful"
func UpdateToDeleteRainfall(param *Rainfall_InputParam, userId int64) (string, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	if param.Id == "" {
		return "", rest.NewError(404, "param 'id' is null.", nil)
	}

	sqlCmd := ""
	if len(strings.Split(param.Id, ",")) == 1 {
		sqlCmd = "WHERE id = " + param.Id
	} else {
		sqlCmd = "WHERE id IN (" + param.Id + ")"
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateToDeleteRainfall + sqlCmd)
	if err != nil {
		return "", errors.Repack(err)
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful.", nil
}
