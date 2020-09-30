package metadata_hydroinfo

import ()

var sqlSelect = ` 	SELECT h.id AS hydroinfo_id, h.hydroinfo_name
					FROM metadata_hydroinfo mh
					LEFT JOIN metadata m ON mh.metadata_id = m.id
					LEFT JOIN lt_hydroinfo h ON mh.hydroinfo_id = h.id 
					WHERE mh.metadata_id = $1 AND m.deleted_by IS NULL AND h.deleted_by IS NULL `

var sqlInsert = ` INSERT INTO metadata_hydroinfo (metadata_id, hydroinfo_id, created_by, updated_by, created_at, updated_at)
				  VALUES ($1, $2, $3, $3, NOW(), NOW()) RETURNING id `

var sqlDeleteByMetadataID = ` DELETE FROM metadata_hydroinfo WHERE metadata_id = $1 AND deleted_by IS NULL `

var sqlUpdateToDeletedByMetadataID = ` UPDATE metadata_hydroinfo
									   SET updated_by = $2
									     , deleted_by = $2
									     , updated_at = NOW()
									     , deleted_at = NOW()
									   WHERE metadata_id = $1 AND deleted_by IS NULL `
