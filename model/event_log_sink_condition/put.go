// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_condition is a model for api.event_log_sink_condition table. This table store condition send e-mail.
package event_log_sink_condition

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// update condition send e-mail by id
//  Parameters:
//		uid
//			user id stamp updated at
//		id
//			sink condition id
//		desc
//			description 
//		eventChannel
//			event channel id 
//		eventCategory
//			event category id
//		eventCode
//			event code id 
//		service
//			service 
//		agent
//			agent user id 
//		user
//			user id 
//		template
//			template id
//		postStartInterval
//			time cron for send email after stamp event log
//  Return:
//		sink condition id
func UpdateSinkCondition(uid, id int64, desc string, eventChannel, eventCategory, eventCode, service, agent, user interface{}, template int64, postStartInterval string) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := updateSinkCondition
	p := []interface{}{eventChannel, eventCategory, eventCode, service, agent, user, template, uid, desc, id, postStartInterval}

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
