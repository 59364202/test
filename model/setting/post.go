// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package setting is a model for api.system_setting table. This table store system_setting data.
package setting

import (
//	model_setting "haii.or.th/api/server/model/setting"
	//result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	insery system setting
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		param
//			Struct_SystemSetting
//	Return:
//		Struct_SystemSetting
func PostSystemSetting(userId int64, param *Struct_SystemSetting) (*Struct_SystemSetting, error) {

	//Try to open database
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

	//Insert system_setting table
	newId, err := insertSystemSetting(tx, param.User_id, param.Name, param.Is_public, param.Value, param.Description, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Set object fot return result
	data := &Struct_SystemSetting{Id: newId, User_id: param.User_id, Name: param.Name, Is_public: param.Is_public, Value: param.Value, Description: param.Description}

	//Return data
	return data, nil
}

//	Insert system_setting table
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
//		new id
func insertSystemSetting(tx *pqx.Tx, settingUserID int64, name string, isPublic bool, value string, description string, userId int64) (int64, error) {

	var _id int64

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertSystemSetting)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(settingUserID, name, isPublic, value, description, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	//Return Value
	return _id, nil
}
