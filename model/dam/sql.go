package dam

import ()

var sqlGetDamFromAgency = "SELECT id, geocode_id, agency_id, dam_lat, dam_long, dam_name, max_water_level, min_water_level " +
	" FROM m_dam " +
	" WHERE m_dam.agency_id = $1 " +
	" ORDER BY m_dam.id "

var sqlGetDam = "SELECT m_dam.id, geocode_id, agency_id, dam_lat, dam_long, dam_name, max_storage, min_storage, agt.agency_name " +
	" FROM m_dam LEFT JOIN agency agt ON agt.id = m_dam.agency_id " +
	" WHERE "
var sqlGetDam_Order = " ORDER BY m_dam.id "

var sqlGetDamStationByDataType = ` SELECT d.id
							  , geocode_id
							  , agency_id
							  , dam_lat
							  , dam_long
							  , dam_oldcode
							  , dam_name
							  , max_water_level
							  , min_water_level
							  , max_storage
							  , min_storage
							  , agency_shortname
						      , agency_name
						  FROM m_dam d
						  LEFT JOIN agency agt ON d.agency_id =  agt.id `

var sqlGetDamStationByDataTypeOrderBy = ` ORDER BY d.dam_name->>'th', agt.agency_shortname->>'en', d.dam_oldcode `

var sqlConditionDamStationByDaily = ` EXISTS (SELECT * FROM dam_daily dd WHERE d.id = dd.dam_id) `
var sqlConditionDamStationByHourly = ` EXISTS (SELECT * FROM dam_hourly dh WHERE d.id = dh.dam_id) `
var sqlGetDamGroupbyAgency = ` SELECT agt.id AS agency_id
									, agt.agency_shortname
									, agt.agency_name
									, string_agg(CAST(d.id AS TEXT) || '##' || CAST(dam_name AS TEXT), '|') AS dam
							   FROM m_dam d
							   LEFT JOIN agency agt ON d.agency_id = agt.id
							   WHERE EXISTS(SELECT id FROM dam_daily dd WHERE dd.deleted_by IS NULL AND d.id = dd.dam_id)
							     AND d.deleted_by IS NULL AND agt.deleted_by IS NULL
							   GROUP BY agt.id `