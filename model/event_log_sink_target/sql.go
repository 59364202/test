package event_log_sink_target

import ()

var (
	getTargetList = "SELECT elst.id, elsc.description, elsm.description, elst.language_code, perg.description " +
		"FROM api.event_log_sink_target elst LEFT JOIN api.event_log_sink_condition elsc ON elst.event_log_sink_condition_id=elsc.id " +
		"LEFT JOIN api.event_log_sink_method elsm ON elst.event_log_sink_method_id=elsm.id LEFT JOIN api.permission_group perg ON elst.target_permission_group_id=perg.id " +
		"WHERE elst.deleted_at IS NULL ORDER BY elsc.description,perg.description"

	getTarget = "SELECT elst.id, elsc.id, elsm.id, elst.language_code, perg.id " +
		"FROM api.event_log_sink_target elst LEFT JOIN api.event_log_sink_condition elsc ON elst.event_log_sink_condition_id=elsc.id " +
		"LEFT JOIN api.event_log_sink_method elsm ON elst.event_log_sink_method_id=elsm.id LEFT JOIN api.permission_group perg ON elst.target_permission_group_id=perg.id " +
		"WHERE elst.id=$1 AND elst.deleted_at IS NULL ORDER BY elst.id"

	insTarget = "INSERT INTO api.event_log_sink_target(event_log_sink_condition_id, event_log_sink_method_id, language_code, target_permission_group_id, created_by, created_at, updated_by,updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), $5, NOW()) RETURNING id"

	updateTarget = "UPDATE api.event_log_sink_target SET event_log_sink_condition_id=$1, event_log_sink_method_id=$2, language_code=$3, target_permission_group_id=$4,updated_by=$5, updated_at=NOW() WHERE id=$6"

	deleteTarget = "UPDATE api.event_log_sink_target SET deleted_by=$2, deleted_at=NOW() WHERE id=$1"

	getSinkCondition = "SELECT id, description FROM api.event_log_sink_condition WHERE deleted_at IS NULL ORDER BY description"

	getSinkMethod = "SELECT id, description FROM api.event_log_sink_method WHERE deleted_at IS NULL ORDER BY description"

	getPerGroup = "SELECT id, description FROM api.permission_group WHERE deleted_at IS NULL AND category='user' ORDER BY description"

	upsertColor = "INSERT INTO api.system_setting (user_id, name, is_public, value, description, created_by, created_at, updated_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, NOW(), $6, NOW()) ON CONFLICT (user_id, name) DO UPDATE SET value=$4, description=$5, updated_by=$6, updated_at=NOW() RETURNING id"
)
