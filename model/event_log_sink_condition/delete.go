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

// soft deleted sink condition by sink condition id
//  Parameters:
//		sinkConditionID
//			sink condition id
//		uid
//			user id for stamp updated by
//  Return:
//		sinkConditionID
func DeleteSinkCondition(sinkConditionID int64, uid int64) (interface{}, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := deleteSinkCondition
	p := []interface{}{sinkConditionID, uid}

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

	return sinkConditionID, nil
}
