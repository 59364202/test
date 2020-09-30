// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_condition is a model for api.event_log_sink_condition table. This table store condition send e-mail.
package event_log_sink_condition

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// get sink select option eventlog condition
//  Parameters:
//		None
//  Return:
//		SinkConditionSelectOption
func GetSinkSelectOption() (*SinkConditionSelectOption, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getEventLogCategory
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	categoryList := make([]*CategoryOptionList, 0)
	dataCategory := &CategoryOptionList{}
	codeList := make([]*SelectOptionJson, 0)
	for rows.Next() {
		dd := &SelectOptionJson{}
		cc := &SelectOptionJson{}
		var (
			id       sql.NullInt64
			name     pqx.JSONRaw
			codeID   sql.NullInt64
			codeName pqx.JSONRaw
		)
		rows.Scan(&id, &name, &codeID, &codeName)
		if dataCategory.Category == nil {
			dd.Value = id.Int64
			dd.Text = name.JSON()
			dataCategory.Category = dd
			cc = &SelectOptionJson{}
			cc.Text = codeName.JSON()
			cc.Value = codeID.Int64
			codeList = append(codeList, cc)
		} else if dataCategory.Category.Value.(int64) != id.Int64 {
			dataCategory.Code = codeList
			categoryList = append(categoryList, dataCategory)
			dataCategory = &CategoryOptionList{}
			codeList = make([]*SelectOptionJson, 0)
			dd = &SelectOptionJson{}
			dd.Value = id.Int64
			dd.Text = name.JSON()
			dataCategory.Category = dd
			cc = &SelectOptionJson{}
			cc.Text = codeName.JSON()
			cc.Value = codeID.Int64
			codeList = append(codeList, cc)
		} else {
			cc = &SelectOptionJson{}
			cc.Text = codeName.JSON()
			cc.Value = codeID.Int64
			codeList = append(codeList, cc)
		}
	}

	if len(codeList) > 0 {
		dataCategory.Code = codeList
		categoryList = append(categoryList, dataCategory)
	}
	q = getTemplate
	p = []interface{}{}
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	templateList := make([]*selectoption.Option, 0)
	for rows.Next() {
		dd := &selectoption.Option{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		dd.Value = id.Int64
		dd.Text = name.String
		templateList = append(templateList, dd)
	}

	data := &SinkConditionSelectOption{}
	data.Category = categoryList
	data.Template = templateList

	return data, nil
}

// get sink condition by array id return list of sink condition
//  Parameters:
//		sinkConditionID
//			array sink condition
//  Return:
//		SinkConditionList
func GetSinkConditions(sinkConditionID []int64) ([]*SinkConditionList, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getSinkConditionList
	p := []interface{}{}

	if len(sinkConditionID) > 0 {
		var condition string
		for i, v := range sinkConditionID {
			if i > 0 {
				condition += " OR elsc.id=$" + strconv.Itoa(i+1)
				p = append(p, v)
			} else {
				condition += " elsc.id=$" + strconv.Itoa(i+1)
				p = append(p, v)
			}
		}
		condition = " WHERE (" + condition + ")"
		q += condition + " AND elsc.deleted_at IS NULL"
	} else {
		q += " WHERE elsc.deleted_at IS NULL"
	}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	list := make([]*SinkConditionList, 0)
	for rows.Next() {
		data := &SinkConditionList{}
		var (
			id                sql.NullInt64
			desc              sql.NullString
			category          sql.NullString
			code              sql.NullString
			template          sql.NullString
			postStartInterval sql.NullString
		)
		rows.Scan(&id, &desc, &category, &code, &template, &postStartInterval)
		data.ID = id.Int64
		data.Condition = desc.String
		data.Category = json.RawMessage(category.String)
		data.Code = json.RawMessage(code.String)
		data.Template = template.String
		data.PostStartInterval = postStartInterval.String
		list = append(list, data)
	}
	return list, nil
}

// get sink condition by id return one sink condition information
//  Parameters:
//		sinkConditionID
//			event log sink condition id 
//  Return:
//		SinkCondition
func GetSinkCondition(sinkConditionID string) (*SinkCondition, error) {
	sinkConditionIDint, _ := strconv.Atoi(sinkConditionID)
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getSinkCondition
	p := []interface{}{sinkConditionIDint}

	q += " WHERE elsc.id=$1 AND elsc.deleted_at IS NULL"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := &SinkCondition{}
	for rows.Next() {
		var (
			id                sql.NullInt64
			desc              sql.NullString
			categortID        sql.NullInt64
			category          sql.NullString
			codeID            sql.NullInt64
			code              sql.NullString
			templateID        sql.NullInt64
			template          sql.NullString
			postStartInterval sql.NullString
		)
		rows.Scan(&id, &desc, &categortID, &category, &codeID, &code, &templateID, &template, &postStartInterval)
		data.ID = id.Int64
		data.Condition = desc.String
		data.CategoryID = categortID.Int64
		if categortID.Valid {
			data.CategoryID = categortID.Int64
		} else {
			data.CategoryID = nil
		}
		data.Category = json.RawMessage(category.String)
		if codeID.Valid {
			data.CodeID = codeID.Int64
		} else {
			data.CodeID = nil
		}
		data.Code = json.RawMessage(code.String)
		data.TemplateID = templateID.Int64
		data.Template = template.String
		data.PostStartInterval = postStartInterval.String
	}
	return data, nil
}
