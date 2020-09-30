// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_method is a model for api.event_log_sink_method table. This table store method config e-mail infomation.
package event_log_sink_method

import (
	"encoding/json"
	model_method_type "haii.or.th/api/thaiwater30/model/event_log_sink_method_type"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

const (
	SettingSMTPServerPrefix = "thaiwater30.service.event_management.smtpserver."
)

// new event log sink method
//  Parameters:
//		uid
//			user id for stamp updated_by
//		sink_method_type
//			sink method type
//		name
//			sink method name
//		description
//			sink method description
//  Return:
//		sinkMethodID
func PostEventLogSinkMethodSystemSetting(uid int64, sink_method_type, name, description string) (int64, error) {
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	type SinkParams struct {
		SystemSettingName string `json:"system_setting_name"`
	}

	sp := &SinkParams{}
	sp.SystemSettingName = SettingSMTPServerPrefix + name

	b, err := json.Marshal(sp)
	if err != nil {
		return 0, nil
	}
	dst := string(b[:])

	q := postEventLogSinkMethod
	p := []interface{}{description, dst, uid, sink_method_type}

	var sinkMethodID int64
	err = db.QueryRow(q, p...).Scan(&sinkMethodID)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	return sinkMethodID, nil
}

// new event log sink method
//  Parameters:
//		userId
//			user id for stamp updated_by
//		param
//			Struct_EventLogSinkMethod_InputParam
//  Return:
//		result ok and data Struct_EventLogSinkMethod
func PostEventLogSinkMethod(userId int64, param *Struct_EventLogSinkMethod_InputParam) (*result.Result, error) {
	//Check event_log_sink_method_type_id is not null
	if param.Event_log_sink_method_type_id == "" {
		return nil, errors.New("'event_log_sink_method_type_id' is not null.")
	}

	//Convert event_log_sink_method_type_id type from string to int64
	intEventLogSinkMethodTypeID, err := strconv.ParseInt(param.Event_log_sink_method_type_id, 10, 64)
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

	//Insert event_log_sink_method table
	newId, err := insertEventLogSinkMethod(tx, intEventLogSinkMethodTypeID, param.Description, param.Sink_params, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Set object fot return result
	objEventLogSinkMethodType := &model_method_type.Struct_EventLogSinkMethodType{}
	objEventLogSinkMethodType.Id = intEventLogSinkMethodTypeID
	data := &Struct_EventLogSinkMethod{Id: newId, Event_log_sink_method_type: objEventLogSinkMethodType, Description: param.Description, Sink_params: param.Sink_params}

	//Return data
	return result.Result1(data), nil
}

//Insert to event_log_sink_method table
//  Parameters:
//		tx
//			transection connect database 
//		eventLogSinkMethodID
//			event log sink method id
//		description
//			sink method description
//		sinkParams
//			sink method params 
//		userId
//			user id for created by
//  Return:
//		result ok and data Struct_EventLogSinkMethod
func insertEventLogSinkMethod(tx *pqx.Tx, eventLogSinkMethodID int64, description string, sinkParams json.RawMessage, userId int64) (int64, error) {

	var (
		_id int64

		jsonSinkParams interface{} = nil
		err            error
	)

	//Convert sinkParams to db-json type
	if sinkParams != nil {
		jsonSinkParams, err = sinkParams.MarshalJSON()
		if err != nil {
			return 0, err
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertEventLogSinkMethod)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(eventLogSinkMethodID, description, jsonSinkParams, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	//Return Value
	return _id, nil
}
