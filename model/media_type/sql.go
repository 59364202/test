package media_type

import ()

var sqlGetMediaType = ` SELECT mt.id
						, mt.media_type_name
						, mt.media_subtype_name
						, mt.media_type_name || (CASE WHEN mt.media_subtype_name IS NOT NULL THEN ' - ' || mt.media_subtype_name ELSE '' END) AS media_type_subtype_name
					FROM lt_media_type mt
					WHERE mt.media_category = 'image'
					  AND mt.deleted_by IS NULL
					  AND mt.deleted_at = '1970-01-01 07:00:00+07' `

var sqlGetMediaTypeOrderBy = ` ORDER BY mt.media_type_name, mt.media_subtype_name `
