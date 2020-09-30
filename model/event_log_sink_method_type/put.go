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

// update lt event log sink method tpye
//  Parameters:
//		id
//			sink method type id
//		uid
//			user id update data
//		description 
//			description sink method type
//  Return:
//		sink method typd id
func UpdateSinkMethodType(id int64, description string, uid int64) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := putLtSinkMethodType
	p := []interface{}{id, description, uid}

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
