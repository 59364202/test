// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_target is a model for api.event_log_sink_target table. This table store target send email information.
package event_log_sink_target

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// update eventlog sink target
//  Parameters:
//		uid
//			user id for updated_by
//		id
//			sink target id
//		condition
//			sink condition id
//		method
//			sink method id 
//		group
//			array permission group id
//		lang
//			lang for send email template
//		color
//			color of group
//  Return:
//		sink target id
func UpdateTarget(uid, id, condition, method, group int64, lang, color string) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := updateTarget
	p := []interface{}{condition, method, lang, group, uid, id}

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
	makeDetailsSystemSetting(color, uid, group)
	return id, nil
}
