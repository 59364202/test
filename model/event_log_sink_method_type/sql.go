// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_method_type is a model for api.lt_event_log_sink_method_type table. This table store method type email information.
package event_log_sink_method_type

import ()

var (
	sqlGetEventLogSinkMethodType = ` SELECT id, description
							   FROM api.lt_event_log_sink_method_type
							   WHERE deleted_at IS NULL AND deleted_by IS NULL `

	sqlGetEventLogSinkMethodTypeOrderby = ` ORDER BY description `
	
	
	getLtSinkMethodType = "SELECT id,description FROM api.lt_event_log_sink_method_type WHERE deleted_at IS NULL AND deleted_by IS NULL ORDER BY description"
	
	postLtSinkMethodType = "INSERT INTO api.lt_event_log_sink_method_type(description, created_by, created_at, updated_by, updated_at)VALUES ($1, $2, NOW(), $2, NOW()) RETURNING id"
	
	putLtSinkMethodType = "UPDATE api.lt_event_log_sink_method_type SET description=$2, updated_by=$3, updated_at=NOW() WHERE id=$1 "
)
