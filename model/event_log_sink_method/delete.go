// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_method is a model for api.event_log_sink_method table. This table store method config e-mail infomation.
package event_log_sink_method

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// Soft deleted event log sink method 
//  Parameters:
//		userId
//			user id stamp updated at
//		param
//			Struct_EventLogSinkMethod_InputParam
//  Return:
//		result ok and message "Delete Successful."
func DeleteEventLogSinkMethod(userId int64, param *Struct_EventLogSinkMethod_InputParam) (*result.Result, error) {
	//Check id is not null
	if param.Id == "" {
		return nil, errors.New("'id' is not null.")
	}

	//Convert id type from string to int64
	intID, err := strconv.ParseInt(param.Id, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

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

	//Check child table of event_log_sink_method
	isHasChild, err := checkEventLogSinkMethodChild(db, intID)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return result.Result0(""), nil
	}

	//Delete event_log_sink_method table
	err = deleteEventLogSinkMethodById(tx, intID)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return result.Result1("Delete Successful."), nil
}

//Check child table of event_log_sink_method
//  Parameters:
//		db
//			connection database
//		id
//			sink method id
//  Return:
//		result ok and message "Delete Successful."
func checkEventLogSinkMethodChild(db *pqx.DB, id int64) (bool, error) {

	//Set default of return value
	var isHasChild bool = false

	//Query statement with parameters
	row, err := db.Query(sqlCheckEventLogSinkMethodChild, id)
	if err != nil {
		return isHasChild, err
	}

	//Check child
	for row.Next() {
		isHasChild = true
	}

	//Return result
	return isHasChild, nil
}

//Delete data at event_log_sink_method table
//  Parameters:
//		tx
//			transection database
//		id
//			sink method id
//  Return:
//		result ok and message "Delete Successful."
func deleteEventLogSinkMethodById(tx *pqx.Tx, id int64) error {

	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteEventLogSinkMethod)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
