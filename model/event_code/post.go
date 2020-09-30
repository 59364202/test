package event_code

import (
	"encoding/json"
	model_event_log_category "haii.or.th/api/thaiwater30/model/event_log_category"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

//  Add new eventcode
//  Parameters:
//		userId
//					user id update event code
//		params
//					Struct_EventCode_InputParam
//  Return:
//		result:ok , data:Struct_EventCode
func PostEventCode(userId int64, param *Struct_EventCode_InputParam) (*result.Result, error) {
	//Check event_log_category_id is not null
	if param.EventCategoryID == "" {
		return nil, errors.New("'event_log_category_id' is not null.")
	}

	//Convert event_log_category_id type from string to int64
	intEventLogCatID, err := strconv.ParseInt(param.EventCategoryID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

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

	//Insert lt_event_code table
	newId, err := insertEventCode(tx, intEventLogCatID, param.Code, param.Description, param.IsAutoclose, param.Troubleshoot, param.SubtypeCategory, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Set object fot return result
	objEventLogCat := &model_event_log_category.Struct_EventLogCategory{}
	objEventLogCat.ID = intEventLogCatID
	data := &Struct_EventCode{ID: newId, EventCategory: objEventLogCat, Code: param.Code, Description: param.Description, IsAutoclose: param.IsAutoclose, Troubleshoot: param.Troubleshoot, SubtypeCategory: param.SubtypeCategory}

	//Return data
	return result.Result1(data), nil
}

//Insert to lt_event_code table
//  Parameters:
//		tx
//			transaction database
//		eventLogCategoryID
//			event log category id
//		code
//			event code
//		description
//			description
//		isAutoClose
//			event auto close
//		troubleshoot
//			how to fix this problem
//		subtypeCategory
//			sub category
//		userId
//			user id for field created by 
//  Return:
//		event code id
func insertEventCode(tx *pqx.Tx, eventLogCategoryID int64, code string, description json.RawMessage, isAutoClose bool, troubleshoot string, subtypeCategory string, userId int64) (int64, error) {

	var (
		_id int64

		jsonDescription interface{} = nil
		err             error
	)

	//Convert description to db-json type
	if description != nil {
		jsonDescription, err = description.MarshalJSON()
		if err != nil {
			return 0, err
		}
	}

	//log.Printf(sqlInsertEventCode, eventLogCategoryID, code, jsonDescription, isAutoClose, troubleshoot, subtypeCategory, userId)

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertEventCode)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(eventLogCategoryID, code, jsonDescription, isAutoClose, troubleshoot, subtypeCategory, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	//Return Value
	return _id, nil
}
