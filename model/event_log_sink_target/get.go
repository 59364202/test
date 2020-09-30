// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_target is a model for api.event_log_sink_target table. This table store target send email information.
package event_log_sink_target

import (
	"database/sql"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// get sink target select option
//  Parameters:
//		None
//  Return:
//		TargetSelectOption
func GetSinkTargetSelectOption() (*TargetSelectOption, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getSinkCondition
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	conditionList := make([]*selectoption.Option, 0)
	for rows.Next() {
		dd := &selectoption.Option{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		dd.Value = id.Int64
		dd.Text = name.String
		conditionList = append(conditionList, dd)
	}

	q = getPerGroup
	p = []interface{}{}
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	groupList := make([]*GroupTargetColor, 0)
	for rows.Next() {
		dd := &GroupTargetColor{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		dd.Value = id.Int64
		dd.Text = name.String
		dd.Color = setting.GetSystemSetting("bof.EventMgt.PermissionGroup.Color_" + strconv.FormatInt(id.Int64, 10))
		groupList = append(groupList, dd)
	}

	q = getSinkMethod
	p = []interface{}{}
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	methodList := make([]*selectoption.Option, 0)
	for rows.Next() {
		dd := &selectoption.Option{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		dd.Value = id.Int64
		dd.Text = name.String
		methodList = append(methodList, dd)
	}
	lang := setting.GetSystemSettingJSON("bof.EventMgt.EventLogSinkTarget.Lang")
	data := &TargetSelectOption{}
	data.Condition = conditionList
	data.Group = groupList
	data.Lang = lang
	data.Method = methodList

	return data, nil
}

// get list of all target
//  Parameters:
//		None
//  Return:
//		Array TargetDetails
func GetTargets() ([]*TargetDetails, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getTargetList
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	list := make([]*TargetDetails, 0)
	for rows.Next() {
		data := &TargetDetails{}
		var (
			id            sql.NullInt64
			conditionName sql.NullString
			methodName    sql.NullString
			lang          sql.NullString
			groupName     sql.NullString
		)
		rows.Scan(&id, &conditionName, &methodName, &lang, &groupName)
		data.ID = id.Int64
		data.Condition = conditionName.String
		data.Method = methodName.String
		data.Lang = lang.String
		data.Group = groupName.String
		list = append(list, data)
	}
	return list, nil
}

// get one sink target
//  Parameters:
//		templateID
//			sink template id
//  Return:
//		TargetEdit
func GetTarget(templateID string) (*TargetEdit, error) {
	templateIDint, _ := strconv.Atoi(templateID)
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getTarget
	p := []interface{}{templateIDint}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := &TargetEdit{}
	for rows.Next() {
		var (
			id          sql.NullInt64
			conditionID sql.NullInt64
			methodID    sql.NullInt64
			lang        sql.NullString
			groupID     sql.NullInt64
		)
		rows.Scan(&id, &conditionID, &methodID, &lang, &groupID)
		data.ID = id.Int64
		data.Condition = conditionID.Int64
		data.Method = methodID.Int64
		data.Lang = lang.String
		data.Group = groupID.Int64
	}
	return data, nil
}
