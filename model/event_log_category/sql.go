package event_log_category

import ()

var sqlGetEventLogCategory = ` SELECT elc.id, elc.code, elc.description, sst.value AS color
							   FROM api.lt_event_log_category elc
							   LEFT JOIN api.system_setting sst ON CONCAT('bof.EventMgt.EventLogCategory.Color_',elc.id) = sst.name
							   WHERE elc.deleted_by IS NULL AND (elc.deleted_at IS NULL OR elc.deleted_at = '1970-01-01 07:00:00+07') `

var sqlGetEventLogCategoryOrderby = ` ORDER BY elc.description->>'en' `

var sqlUpdateEventLogCategory = ` UPDATE api.lt_event_log_category
								  SET code = $2
								    , description = $3
								    , updated_by = $4
								    , updated_at = NOW()
								  WHERE id = $1 
								    AND deleted_at IS NULL AND deleted_by IS NULL `

var sqlInsertEventLogCategory = ` INSERT INTO api.lt_event_log_category(
										code, description
									  , created_by, updated_by
									  , created_at, updated_at)
							      VALUES ($1, $2, $3, $3, NOW(), NOW()) RETURNING id `

var sqlDeleteEventLogCategory = ` DELETE FROM api.lt_event_log_category WHERE id = $1 `

var sqlCheckEventLogCategoryChild = ` SELECT event_log_category_id FROM api.lt_event_code WHERE event_log_category_id = $1 AND deleted_at IS NULL AND deleted_by IS NULL `
