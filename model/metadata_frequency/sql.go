package metadata_frequency

import ()

var SQL_GetFrequencyFromMetadataId = `SELECT id, datafrequency
 FROM metadata_frequency
 WHERE metadata_id = $1 AND deleted_at = '1970-01-01 07:00:00+07'
`
var sqlInsert = ` INSERT INTO metadata_frequency (metadata_id, datafrequency, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $3, NOW(), NOW()) RETURNING id `

var sqlUpdateToDeletedByMetadataID = ` UPDATE metadata_frequency 
										  SET deleted_by = $2
										    , updated_by = $2
										    , deleted_at = NOW()
										    , updated_at = NOW()
										WHERE metadata_id = $1 AND deleted_by IS NULL `

var sqlDeleteByMetadataID = ` DELETE FROM metadata_frequency WHERE metadata_id = $1 AND deleted_by IS NULL `
