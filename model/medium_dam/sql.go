package medium_dam

var sqlGetDamGroupbyAgency = ` SELECT agt.id AS agency_id
									, agt.agency_shortname
									, agt.agency_name
									, string_agg(CAST(d.id AS TEXT) || '##' || CAST(mediumdam_name AS TEXT), '|') AS dam
							   FROM m_medium_dam d
							   LEFT JOIN agency agt ON d.agency_id = agt.id
							   WHERE EXISTS(SELECT id FROM medium_dam dd WHERE dd.deleted_by IS NULL AND d.id = dd.mediumdam_id)
							     AND d.deleted_by IS NULL AND agt.deleted_by IS NULL
							   GROUP BY agt.id `

var SQL_GetDamLastest = `  
	SELECT dd.id, dd.mediumdam_date
		, st.id AS mediumdam_id, st.mediumdam_name, st.mediumdam_lat, st.mediumdam_long
		, dd.mediumdam_inflow
		, dd.mediumdam_storage, dd.mediumdam_storage_percent
		, dd.mediumdam_uses_water
		, dd.mediumdam_released
		, agt.id AS agency_id, agt.agency_shortname, agt.agency_name
		, b.id AS basin_id, b.basin_code, b.basin_name
		, geo.id AS geocode_id, geo.geocode, geo.rid_area_code AS area_code, geo.tmd_area_name AS area_name, geo.province_name
		, geo.amphoe_name, geo.tumbon_name, geo.province_code, st.mediumdam_oldcode, st.subbasin_id	
	FROM medium_dam dd`

var SQL_GetDamLastest2 = ` 
	INNER JOIN m_medium_dam st ON st.id = dd.mediumdam_id
	INNER JOIN agency agt ON agt.id = st.agency_id
	LEFT  JOIN subbasin sb ON sb.id = st.subbasin_id
	LEFT  JOIN basin b ON b.id = sb.basin_id
	INNER JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE dd.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.deleted_at = '1970-01-01 07:00:00+07' 
	  AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' ) 
	  AND st.is_ignore = 'false'
	  AND mediumdam_date = (SELECT max(mediumdam_date) FROM medium_dam)`

var SQL_GetDamLastest_OrderBy = ` ORDER BY geo.rid_area_code, st.id `
