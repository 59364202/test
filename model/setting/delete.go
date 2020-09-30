// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package setting is a model for api.system_setting table. This table store system_setting data.
package setting

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	delete system setting
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		param
//			Struct_SystemSetting
//	Return:
//		result.Result
func DeleteSystemSetting(userId int64, param *Struct_SystemSetting) (*result.Result, error) {

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

	//Delete system_setting table
	err = deleteSystemSetting(tx, param.Name)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return result.Result1("Delete Successful."), nil
}

//	Delete data at system_setting table
//	Parameters:
//		tx
//			transaction
//		name
//			ชื่อ setting
//	Return:
//		nil, error
func deleteSystemSetting(tx *pqx.Tx, name string) error {

	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteSystemSetting)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(name)
	if err != nil {
		return err
	}

	return nil
}
