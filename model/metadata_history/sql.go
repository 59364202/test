package metadata_history

import ()

var sqlSelect = " SELECT h.id AS history_id, h.metadata_datetime AS history_datetime , h.history_description , u.id, u.full_name " +
	" FROM metadata_history h " +
	" LEFT JOIN api.user u ON h.created_by = u.id " +
	" WHERE h.metadata_id = $1 AND h.deleted_by IS NULL "

var sqlInsert = ` INSERT INTO metadata_history (metadata_id, metadata_datetime, history_description, created_by, updated_by, created_at, updated_at)
				  VALUES ($1, NOW(), $2, $3, $3, NOW(), NOW()) RETURNING id `
