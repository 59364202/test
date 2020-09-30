package event_log_category

import (
	"encoding/json"
	model_setting "haii.or.th/api/thaiwater30/model/setting"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"strconv"
)

//  Add new category
//  Parameters:
//		userId
//					user id update event category
//		params
//					Struct_EventLogCategory_InputParam
//  Return:
//		result:ok , data:Struct_EventCode
func PostEventLogCategory(userId int64, param *Struct_EventLogCategory_InputParam) (*Struct_ELC, error) {
	if param.Code == "" {
		return nil, rest.NewError(422, "code is not null.", errors.New("parameter 'code' is not null."))
	}

	//Try to open database
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

	//Insert lt_event_log_category table
	newId, err := insertEventLogCategory(tx, param.Code, param.Description, userId)
	if err != nil {
		return nil, err
	}

	//Insert system_setting table

	objSettingParam := &model_setting.Struct_SystemSetting{}
	objSettingParam.User_id = 1
	objSettingParam.Name = "bof.EventMgt.EventLogCategory.Color_" + strconv.FormatInt(newId, 10)
	objSettingParam.Is_public = false
	objSettingParam.Value = param.Color
	objSettingParam.Description = "Color Setting for EventLogCategory ID '" + strconv.FormatInt(newId, 10) + "'"

	_, err = model_setting.PostSystemSetting(userId, objSettingParam)
	if err != nil {
		defer tx.Rollback()
		return nil, pqx.GetRESTError(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	data := &Struct_ELC{ID: newId, Code: param.Code, Description: param.Description, Color: param.Color}
	return data, nil
}

//Insert to lt_event_log_category table
//  Parameters:
//		tx
//			transaction database
//		code
//			event category
//		description
//			description
//		userId
//			user id for field created by 
//  Return:
//		event category id
func insertEventLogCategory(tx *pqx.Tx, code string, description json.RawMessage, userId int64) (int64, error) {

	var (
		_id int64

		jsonDescription interface{} = nil
		err             error
	)

	//Convert description to db-json type
	if description != nil {
		jsonDescription, err = description.MarshalJSON()
		if err != nil {
			return 0, errors.Repack(err)
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertEventLogCategory)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(code, jsonDescription, userId).Scan(&_id)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	//Return Value
	return _id, nil
}
