package event_log_category

import (
	"encoding/json"
	model_setting "haii.or.th/api/thaiwater30/model/setting"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// update event category
//  Parameters:
//		userId
//					user id update event category
//		params
//					Struct_EventLogCategory_InputParam
//  Return:
//		result:ok , data:Struct_EventCode
func PutEventLogCategory(userId int64, param *Struct_EventLogCategory_InputParam) (*Struct_ELC, error) {
	//Check id is not null
	if param.ID == "" {
		return nil, errors.New("'id' is not null.")
	}

	//Convert id type from string to int64
	intID, err := strconv.ParseInt(param.ID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
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

	//Update lt_event_log_category table
	err = updateEventLogCategory(tx, intID, param.Code, param.Description, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Update system_setting table
	objSettingParam := &model_setting.Struct_SystemSetting{}
	objSettingParam.User_id = 1
	objSettingParam.Name = "bof.EventMgt.EventLogCategory.Color_" + strconv.FormatInt(intID, 10)
	objSettingParam.Is_public = false
	objSettingParam.Value = param.Color
	//objSettingParam.Description = "Color Setting for EventLogCategory ID '" + strconv.FormatInt(intID, 10) + "'"

	_, err = model_setting.PutSystemSetting(userId, objSettingParam)
	if err != nil {
		defer tx.Rollback()
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	data := &Struct_ELC{ID: intID, Code: param.Code, Description: param.Description, Color: param.Color}

	//Return result
	return data, nil
}

//Update lt_event_log_category table
//  Parameters:
//		tx
//			transaction database
//		id
//			event category id
//		code
//			event code
//		description
//			description
//		userId
//			user id for field created by 
//  Return:
//		event code id
func updateEventLogCategory(tx *pqx.Tx, id int64, code string, description json.RawMessage, userId int64) error {
	var (
		jsonDescription interface{} = nil
		err             error
	)

	//Convert description to db-json type
	if description != nil {
		jsonDescription, err = description.MarshalJSON()
		if err != nil {
			return err
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateEventLogCategory)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(id, code, jsonDescription, userId)
	if err != nil {
		return err
	}

	return nil
}
