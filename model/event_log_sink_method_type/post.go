// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_method_type is a model for api.lt_event_log_sink_method_type table. This table store method type email information.
package event_log_sink_method_type
import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// new lt eventlog sink method type
//  Parameters:
//		uid
//			user id insert data
//		description 
//			description sink method type
//  Return:
//		sink method typd id
func AddSinkMethodType(description string, uid int64) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := postLtSinkMethodType
	p := []interface{}{description,uid}

	var methodTypeID int64
	err = db.QueryRow(q, p...).Scan(&methodTypeID)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	return methodTypeID, nil
}
