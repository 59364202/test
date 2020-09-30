package event_log_sink_method

import ()

var sqlGetEventLogSinkMethod = ` SELECT sm.id
									, sm.description
									, sm.sink_params
									, sm.event_log_sink_method_type_id AS method_type_id
									, smt.description AS method_type_description
								FROM api.event_log_sink_method sm
								LEFT JOIN api.lt_event_log_sink_method_type smt ON smt.id = sm.event_log_sink_method_type_id
								WHERE sm.deleted_at IS NULL AND sm.deleted_by IS NULL
								  AND smt.deleted_at IS NULL AND smt.deleted_by IS NULL `

var sqlGetEventLogSinkMethodOrderby = ` ORDER BY sm.description `

var sqlUpdateEventLogSinkMethod = ` UPDATE api.event_log_sink_method
								  SET event_log_sink_method_type_id = $2
								    , description = $3
								    , sink_params = $4
								    , updated_by = $5
								    , updated_at = NOW()
								  WHERE id = $1 
								    AND deleted_at IS NULL AND deleted_by IS NULL `

var sqlInsertEventLogSinkMethod = ` INSERT INTO api.event_log_sink_method(
										event_log_sink_method_type_id
									  , description
									  , sink_params
									  , created_by, updated_by
									  , created_at, updated_at)
							      VALUES ($1, $2, $3, $4, $4, NOW(), NOW()) RETURNING id `

var sqlDeleteEventLogSinkMethod = ` DELETE FROM api.event_log_sink_method WHERE id = $1 `

var sqlCheckEventLogSinkMethodChild = ` SELECT event_log_sink_method_id FROM api.event_log_sink_target WHERE event_log_sink_method_id = $1 AND deleted_at IS NULL AND deleted_by IS NULL `

var postEventLogSinkMethod = "INSERT INTO api.event_log_sink_method(description, sink_params, created_by, created_at, updated_by, updated_at, event_log_sink_method_type_id) VALUES($1, $2, $3, NOW(), $3, NOW(), $4) RETURNING id"

var putEventLogSinkMethod = "UPDATE api.event_log_sink_method SET description=$1, sink_params=$2, updated_by=$3, updated_at=NOW(), event_log_sink_method_type_id=$4 WHERE id=$5"