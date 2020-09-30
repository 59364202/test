package event_log_sink_template

import ()

var (
	getTemplateList = "SELECT id,name FROM api.event_log_sink_template WHERE deleted_at IS NULL ORDER BY name"

	getTemplate = "SELECT id,name,message_subject,message_body FROM api.event_log_sink_template WHERE deleted_at IS NULL AND id = $1"

	insTemplate = "INSERT INTO api.event_log_sink_template( name, created_by, created_at, updated_by, updated_at,message_subject, message_body) VALUES ($1, $4, NOW(), $4, NOW(), $2, $3) RETURNING id"

	updateTemplate = "UPDATE api.event_log_sink_template SET name=$1,updated_by=$4, updated_at=NOW(), message_subject=$2, message_body=$3 WHERE id=$5"

	deleteTemplate = "UPDATE api.event_log_sink_template SET deleted_by=$2, deleted_at=NOW() WHERE id=$1"
)
