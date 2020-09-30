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

// update eventlog sink method
//  Parameters:
//		id
//			sink method id
//		uid
//			user id for updated by
//		sink_method_type
//			sink method type
//		name
//			sink method name
//		description
//			sink method description
//  Return:
//		sink method id
func PutEventLogSinkMethodSystemSetting(id string, uid int64, sink_method_type, name, description string) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	iid, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
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

	q := putEventLogSinkMethod
	p := []interface{}{description, dst, uid, sink_method_type, id}

	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return int64(iid), nil

}

// update eventlog sink method
//  Parameters:
//		userId
//			sink method id
//		param
//			Struct_EventLogSinkMethod_InputParam
//  Return:
//		result ok and data Struct_EventLogSinkMethod
func PutEventLogSinkMethod(userId int64, param *Struct_EventLogSinkMethod_InputParam) (*result.Result, error) {
	//Check id is not null
	if param.Id == "" {
		return nil, errors.New("'id' is not null.")
	}

	//Check event_log_sink_method_type_id is not null
	if param.Event_log_sink_method_type_id == "" {
		return nil, errors.New("'event_log_sink_method_type_id' is not null.")
	}

	//Convert id type from string to int64
	intID, err := strconv.ParseInt(param.Id, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Convert event_log_sink_method_type_id type from string to int64
	intEventLogSinkMethodTypeID, err := strconv.ParseInt(param.Event_log_sink_method_type_id, 10, 64)
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

	//Update lt_event_log_sink_method table
	err = updateEventLogSinkMethod(tx, intID, intEventLogSinkMethodTypeID, param.Description, param.Sink_params, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	objEventLogSinkMethodType := &model_method_type.Struct_EventLogSinkMethodType{}
	objEventLogSinkMethodType.Id = intEventLogSinkMethodTypeID
	data := &Struct_EventLogSinkMethod{Id: intID, Event_log_sink_method_type: objEventLogSinkMethodType, Description: param.Description, Sink_params: param.Sink_params}

	//Return result
	return result.Result1(data), nil
}

//Update lt_event_log_sink_method table
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
func updateEventLogSinkMethod(tx *pqx.Tx, id int64, eventLogSinkMethodTypeID int64, description string, sinkParams json.RawMessage, userId int64) error {
	var (
		jsonSinkParams interface{} = nil
		err            error
	)

	//Convert sinkParams to db-json type
	if sinkParams != nil {
		jsonSinkParams, err = sinkParams.MarshalJSON()
		if err != nil {
			return err
		}
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateEventLogSinkMethod)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(id, eventLogSinkMethodTypeID, description, jsonSinkParams, userId)
	if err != nil {
		return err
	}

	return nil
}
