// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_template is a model for api.event_log_sink_template table. This table store html template email.
package event_log_sink_template

import (
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// update email template
//  Parameters:
//		uid
//			user id 
//		id
//			sink template id
//		name
//			template name
//		subjectJson
//			subject email
//		bodyJson
//			body email
//  Return:
//		sink template id
func UpdateTemplate(uid, id int64, name string, subjectJson, bodyJson interface{}) (int64, error) {
	b, err := json.Marshal(subjectJson)
	if err != nil {
		return 0, err
	}
	subject := string(b[:])

	b, err = json.Marshal(bodyJson)
	if err != nil {
		return 0, err
	}
	body := string(b[:])

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := updateTemplate
	p := []interface{}{name, subject, body, uid, id}

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

	return id, nil
}
