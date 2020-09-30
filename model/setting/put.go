// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package setting is a model for api.system_setting table. This table store system_setting data.
package setting

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	Update system setting
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		param
//			Struct_SystemSetting
//	Return:
//		Struct_SystemSetting
func PutSystemSetting(userId int64, param *Struct_SystemSetting) (*Struct_SystemSetting, error) {
	//Check name is not null
	if param.Name == "" {
		return nil, errors.New("'name' must not be null.")
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

	//Update system_setting table
	err = updateSystemSetting(tx, param.User_id, param.Name, param.Is_public, param.Value, param.Description, userId)
	//string(param.Data[:])
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	data := &Struct_SystemSetting{User_id: param.User_id, Name: param.Name, Is_public: param.Is_public, Value: param.Value, Description: param.Description}

	//Return result
	return data, nil
}

//	Update system_setting table
//	Parameters:
//		tx
//			transaction
//		settingUserID
//			รหัสผู้ใช้
//		name
//			ชื่อ setting
//		isPublic
//			เป็นสาธารณะ
//		value
//			ค่า setting
//		description
//			คำอธิบาย
//		userId
//			รหัสผู้ใช้
//	Return:
//		nil, error
func updateSystemSetting(tx *pqx.Tx, settingUserID int64, name string, ispublic bool, value string, description string, userId int64) error {

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateSystemSetting)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(settingUserID, name, ispublic, value, userId)
	if err != nil {
		return err
	}

	return nil
}
