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

// soft deleted event log sink target
//  Parameters:
//		templateID
//			template id
// 		uid
//			user id soft deleted data
//  Return:
//		Array information cctv
func DeleteTarget(templateID int64, uid int64) (interface{}, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := deleteTarget
	p := []interface{}{templateID, uid}

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	return templateID, nil
}
