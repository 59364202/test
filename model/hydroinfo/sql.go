package hydroinfo

import ()

var sqlGetHydroinfo = ` SELECT h.id, h.hydroinfo_name, h.hydroinfo_number
							  , ha.agency_id, ag.agency_shortname, ag.agency_name
						 FROM lt_hydroinfo h 
						 LEFT JOIN lt_hydroinfo_agency ha ON h.id = ha.hydroinfo_id
						 LEFT JOIN agency ag ON ha.agency_id = ag.id
						 WHERE h.deleted_at = to_timestamp(0) AND ag.deleted_at = to_timestamp(0) `

var sqlGetHydroinfoByMetadata = ` 	SELECT h.id, h.hydroinfo_name
										  , ha.agency_id, ag.agency_shortname, ag.agency_name
									FROM lt_hydroinfo h
									LEFT JOIN lt_hydroinfo_agency ha ON h.id = ha.hydroinfo_id
									LEFT JOIN agency ag ON ha.agency_id = ag.id
									WHERE h.deleted_by IS NULL 
									  AND ag.deleted_by IS NULL
									  AND EXISTS (SELECT mh.hydroinfo_id FROM metadata_hydroinfo mh WHERE mh.metadata_id = $1 AND mh.hydroinfo_id = h.id) `

var sqlInsertHydroinfo = "INSERT INTO lt_hydroinfo (hydroinfo_name, hydroinfo_number, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $3, NOW(), NOW()) RETURNING id "

var sqlInsertHydroinfoAgency = "INSERT INTO lt_hydroinfo_agency (hydroinfo_id, agency_id, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $3, NOW(), NOW()) "

var sqlUpdateHydroinfo = ` UPDATE lt_hydroinfo
						   SET hydroinfo_name = $2
						     , hydroinfo_number = $3
						     , updated_by = $4
						     , updated_at = NOW()
						   WHERE id = $1 `

var sqlUpdateToDeleteHydroinfo = ` UPDATE lt_hydroinfo
								   SET deleted_by = $2
									 , deleted_at = NOW()
									 , updated_by = $2
									 , updated_at = NOW()
								   WHERE id = $1 `

var sqlUpdateToDeleteHydroinfoAgency = ` UPDATE lt_hydroinfo_agency
										 SET deleted_by = $2
										   , deleted_at = NOW()
										   , updated_by = $2
										   , updated_at = NOW()
										   WHERE hydroinfo_id = $1 `

var sqlDeleteHydroinfoAgency = ` DELETE FROM lt_hydroinfo_agency WHERE hydroinfo_id = $1 `