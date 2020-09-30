// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_template is a model for api.event_log_sink_template table. This table store html template email.
package event_log_sink_template

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// get email template list
//  Parameters:
//		None
//  Return:
//		Array TemplateList
func GetTemplates() ([]*TemplateList, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getTemplateList
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	list := make([]*TemplateList, 0)
	for rows.Next() {
		data := &TemplateList{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		data.ID = id.Int64
		data.Name = name.String
		list = append(list, data)
	}
	return list, nil
}

// get email template by id
//  Parameters:
//		templateID
//			sink template id 
//  Return:
//		TemplateDetails
func GetTemplate(templateID string) (*TemplateDetails, error) {
	templateIDint, _ := strconv.Atoi(templateID)
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getTemplate
	p := []interface{}{templateIDint}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := &TemplateDetails{}
	for rows.Next() {
		var (
			id              sql.NullInt64
			name            sql.NullString
			message_subject sql.NullString
			message_body    sql.NullString
		)
		rows.Scan(&id, &name, &message_subject, &message_body)
		data.ID = id.Int64
		data.Name = name.String
		data.MessageSubject = json.RawMessage(message_subject.String)
		data.MessageBody = json.RawMessage(message_body.String)
	}
	return data, nil
}
