package event_code

import (
	"encoding/json"
	model_event_log_category "haii.or.th/api/thaiwater30/model/event_log_category"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// update eventcode
//  Parameters:
//		userId
//					user id update event code
//		params
//					Struct_EventCode_InputParam
//  Return:
//		result:ok , data:Struct_EventCode
func PutEventCode(userId int64, param *Struct_EventCode_InputParam) (*result.Result, error) {
	//Check id is not null
	if param.ID == "" {
		return nil, errors.New("'id' is not null.")
	}

	//Convert id type from string to int64
	intID, err := strconv.ParseInt(param.ID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Check event_log_category_id is not null
	if param.EventCategoryID == "" {
		return nil, errors.New("'event_log_category_id' is not null.")
	}

	//Convert event_log_category_id type from string to int64
	intEventLogCatID, err := strconv.ParseInt(param.EventCategoryID, 10, 64)
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

	//Update lt_event_code table
	err = updateEventCode(tx, intID, intEventLogCatID, param.Code, param.Description, param.IsAutoclose, param.Troubleshoot, param.SubtypeCategory, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	objEventLogCat := &model_event_log_category.Struct_EventLogCategory{}
	objEventLogCat.ID = intEventLogCatID
	data := &Struct_EventCode{ID: intID, EventCategory: objEventLogCat, Code: param.Code, Description: param.Description, IsAutoclose: param.IsAutoclose, Troubleshoot: param.Troubleshoot, SubtypeCategory: param.SubtypeCategory}

	//Return result
	return result.Result1(data), nil
}

//Update lt_event_code table
//  Parameters:
//		tx
//			transaction database
//		id
//			eventcode id
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
func updateEventCode(tx *pqx.Tx, id int64, eventLogCategoryID int64, code string, description json.RawMessage, isAutoClose bool, troubleshoot string, subtypeCategory string, userId int64) error {
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
	statement, err := tx.Prepare(sqlUpdateEventCode)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(id, eventLogCategoryID, code, jsonDescription, isAutoClose, troubleshoot, subtypeCategory, userId)
	if err != nil {
		return err
	}

	return nil
}
