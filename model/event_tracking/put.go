// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_tracking is a model for api.event_log table. This table store eventlog information.
package event_tracking

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// update event tracking
//  Parameters:
//		eventLogId
//			event log id
//		msg
//			message
//		uid
//			user id
//  Return:
//		"Update EventLog Solve Successful"
func UpdateEventTracking(eventLogId []int64, msg string, uid int64) (string, error) {

	if len(eventLogId) < 0 {
		return "", errors.New("not enought parameter for tracking function")
	}

	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := updateEventTrackingMessage
	p := []interface{}{uid, msg}

	count := 3
	for i, v := range eventLogId {
		if i > 0 {
			q += " OR id=$" + strconv.Itoa(count)
		} else {
			q += " WHERE id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	return "Update EventLog Solve Successful", nil
}

// update event tracking when send invalid data
//  Parameters:
//		eventLogId
//			array event log id
//		sendErrDate
//			stamp date when send error to agency
//		uid
//			user id
//  Return:
//		"Update EventLog Send Error At Successful"
func UpdateSendEventInvalidData(eventLogId []int64, sendErrDate string, uid int64) (string, error) {

	if len(eventLogId) < 0 {
		return "", errors.New("not enought parameter for send invalid data function")
	}

	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := updateEventSendInvalidData
	p := []interface{}{uid, sendErrDate}

	count := 3
	for i, v := range eventLogId {
		if i > 0 {
			q += " OR id=$" + strconv.Itoa(count)
		} else {
			q += " WHERE id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	return "Update EventLog Send Error At Successful", nil
}

// update event tracking invalid data
//  Parameters:
//		eventLogId
//			array event log id
//		msg
//			message invalid data
//		uid
//			user id
//  Return:
//		"Update EventLog Solve Successful"
func UpdateEventTrackingInvalidData(eventLogId []int64, msg string, uid int64) (string, error) {

	if len(eventLogId) < 0 {
		return "", errors.New("not enought parameter for tracking invalid data function")
	}

	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := updateEventTrackingMessage
	p := []interface{}{uid, msg}

	count := 3
	for i, v := range eventLogId {
		if i > 0 {
			q += " OR id=$" + strconv.Itoa(count)
		} else {
			q += " WHERE id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	return "Update EventLog Solve Successful", nil
}
