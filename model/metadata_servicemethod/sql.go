package metadata_servicemethod

import ()

var SQL_GetServicemethodFromMetadataId = `SELECT s.id, s.servicemethod_name 
 FROM metadata_servicemethod ms
 INNER JOIN lt_servicemethod s ON ms.servicemethod_id = s.id AND s.deleted_at = '1970-01-01 07:00:00+07'
 WHERE ms.metadata_id = $1 AND ms.deleted_at = '1970-01-01 07:00:00+07'
`

var sqlInsert = ` INSERT INTO metadata_servicemethod (metadata_id, servicemethod_id, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $3, NOW(), NOW()) RETURNING id `

var sqlUpdateToDeletedByMetadataID = ` UPDATE metadata_servicemethod 
										  SET deleted_by = $2
										    , updated_by = $2
										    , deleted_at = NOW()
										    , updated_at = NOW()
										WHERE metadata_id = $1 AND deleted_by IS NULL `

var sqlDeleteByMetadataID = ` DELETE FROM metadata_servicemethod WHERE metadata_id = $1 AND deleted_by IS NULL `
