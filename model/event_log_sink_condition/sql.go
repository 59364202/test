package event_log_sink_condition

import ()

var (
	getSinkConditionList = "SELECT elsc.id,elsc.description, lelc.description, lec.description, elst.name AS event_log_template, post_start_interval " +
		"FROM api.event_log_sink_condition elsc LEFT JOIN api.lt_event_code lec ON elsc.event_code_id=lec.id LEFT JOIN api.lt_event_log_category lelc ON elsc.event_log_category_id=lelc.id " +
		"LEFT JOIN api.event_log_sink_template elst ON elsc.event_log_sink_template_id=elst.id "

	getSinkCondition = "SELECT elsc.id,elsc.description, lelc.id, lelc.description, lec.id, lec.description, elst.id, elst.name AS event_log_template, post_start_interval " +
		"FROM api.event_log_sink_condition elsc LEFT JOIN api.lt_event_code lec ON elsc.event_code_id=lec.id LEFT JOIN api.lt_event_log_category lelc ON elsc.event_log_category_id=lelc.id " +
		"LEFT JOIN api.event_log_sink_template elst ON elsc.event_log_sink_template_id=elst.id "

	insSinkCondition = "INSERT INTO api.event_log_sink_condition(event_log_channel_id, event_log_category_id, event_code_id, service_id, agent_user_id, user_id, " +
		"event_log_sink_template_id, created_by, created_at, updated_by, updated_at, description, post_start_interval) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), $8, NOW(), $9, $10) RETURNING id"

	updateSinkCondition = "UPDATE api.event_log_sink_condition SET event_log_channel_id=$1, event_log_category_id=$2, event_code_id=$3, service_id=$4, agent_user_id=$5, user_id=$6, event_log_sink_template_id=$7, updated_by=$8, updated_at=NOW(), description=$9, post_start_interval=$11 WHERE id=$10"

	deleteSinkCondition = "UPDATE api.event_log_sink_condition SET deleted_by=$2, deleted_at=NOW() WHERE id=$1"

	getEventLogChannel = "SELECT id,description FROM api.lt_event_log_channel WHERE deleted_at IS NULL ORDER BY id"

	getEventLogCategory = "SELECT elc.id,elc.description::jsonb,lec.id,lec.description::jsonb FROM api.lt_event_log_category elc LEFT JOIN api.lt_event_code lec ON elc.id=lec.event_log_category_id WHERE elc.deleted_at IS NULL ORDER BY elc.description::jsonb,lec.description::jsonb"

	getTemplate = "SELECT id,name FROM api.event_log_sink_template WHERE deleted_at IS NULL ORDER BY name"
)
