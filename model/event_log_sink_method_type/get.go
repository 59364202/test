// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_method_type is a model for api.lt_event_log_sink_method_type table. This table store method type email information.
package event_log_sink_method_type
import (
	//result "haii.or.th/api/thaiwater30/util/result"
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

// get sink method type description
//  Parameters:
//		None
//  Return:
//		Array Struct_EventLogSinkMethodType
func GetEventLogSinkMethodType() ([]*Struct_EventLogSinkMethodType, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data                      []*Struct_EventLogSinkMethodType
		objEventLogSinkMethodType *Struct_EventLogSinkMethodType

		_id          sql.NullInt64
		_description sql.NullString

		_result *sql.Rows
	)

	//Query
	_result, err = db.Query(sqlGetEventLogSinkMethodType + sqlGetEventLogSinkMethodTypeOrderby)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_EventLogSinkMethodType, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_description)

		if err != nil {
			return nil, err
		}

		//Generate EventLogSinkMethodType object
		objEventLogSinkMethodType = &Struct_EventLogSinkMethodType{}
		objEventLogSinkMethodType.Id = _id.Int64
		objEventLogSinkMethodType.Description = _description.String

		data = append(data, objEventLogSinkMethodType)
	}

	//Return result
	return data, nil
}

// get lt eventlog sink method type
//  Parameters:
//		None
//  Return:
//		Array Struct_EventLogSinkMethodType
func getLtEventLogMethodType() ([]*Struct_EventLogSinkMethodType, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	q := getLtSinkMethodType
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	
	data := make([]*Struct_EventLogSinkMethodType,0)
	for rows.Next() {
		var (
			id          sql.NullInt64
			description sql.NullString
		)
		dataRow := &Struct_EventLogSinkMethodType{}
		rows.Scan(&id, &description)
		dataRow.Id = id.Int64
		dataRow.Description = description.String 
		data = append(data,dataRow)
	}
	
	return data,nil
}
