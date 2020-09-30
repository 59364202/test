package dam_hourly

var arrDamHourlyByStationAndDateColumn = []string{
	"id",
	//"dam_id",
	"dam_date",
	"dam_level",
	"dam_storage",
	"dam_storage_percent",
	"dam_inflow",
	"dam_inflow_acc_percent",
	"dam_uses_water",
	"dam_uses_water_percent",
	"dam_released",
	"dam_spilled",
	"dam_losses",
	"dam_evap",
}

var sqlGetHourlyByStationAndDate = ` SELECT id, dam_datetime
								  , dam_storage, dam_storage_percent
								  , dam_inflow, dam_inflow_acc_percent
								  , dam_uses_water, dam_uses_water_percent
								  , dam_level, dam_released, dam_spilled, dam_losses, dam_evap
							 FROM dam_hourly
							 WHERE dam_id = $1
							   AND dam_datetime BETWEEN $2 AND $3
							   AND deleted_at = '1970-01-01 07:00:00+07'`

var sqlGetHourlyByStationAndDateOrderBy = ` ORDER BY dam_datetime DESC `

var sqlUpdateToDeleteDamHourly = ` UPDATE dam_hourly
									SET deleted_by = $1
									  , deleted_at = NOW()
									  , updated_by = $1
									  , updated_at = NOW() `

var SQL_GetDamHourlyLastest = `
	SELECT dd.id AS dam_hourly_id, dd.dam_datetime
		, st.id AS dam_id, st.dam_name, st.dam_lat, st.dam_long
		, dd.dam_inflow, dd.dam_inflow_acc_percent
		, dd.dam_storage, dd.dam_storage_percent
		, dd.dam_uses_water, dd.dam_uses_water_percent
		, dd.dam_released
		, dd.dam_level
		, dd.dam_spilled
		, dd.dam_losses
		, dd.dam_evap
		, agt.id AS agency_id, agt.agency_shortname, agt.agency_name
		, b.id AS basin_id, b.basin_code, b.basin_name
		, geo.id AS geocode_id, geo.geocode, geo.rid_area_code, geo.rid_area_name, geo.province_name, geo.amphoe_name, geo.tumbon_name, geo.province_code, st.dam_oldcode,subbasin_id
	FROM (SELECT dam_id, max(dam_datetime) AS dam_datetime
		  FROM dam_hourly
		  WHERE deleted_at = '1970-01-01 07:00:00+07' `

// temporary by pass QC Rule
// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
var SQL_GetDamHourlyLastest2 = ` GROUP BY dam_id) md
	INNER JOIN dam_hourly dd ON dd.dam_id = md.dam_id AND dd.dam_datetime = md.dam_datetime
	INNER JOIN m_dam st ON st.id = dd.dam_id
	INNER JOIN agency agt ON agt.id = st.agency_id
	LEFT  JOIN subbasin sb ON sb.id = st.subbasin_id
	LEFT  JOIN basin b ON b.id = sb.basin_id
	INNER JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE dd.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.is_ignore = 'false'`

var SQL_GetDamHourlyLastestOrderBy = ` ORDER BY geo.rid_area_code, st.id `

var sqlGetErrorDamHourly = ` SELECT dd.id
								  , dam_oldcode
								  , dam_datetime
								  , dam_name
								  , province_name
								  , agency_name
								  , agency_shortname
								  , dam_storage, dam_storage_percent
								  , dam_inflow, dam_inflow_acc_percent
								  , dam_uses_water, dam_uses_water_percent
								  , dam_level, dam_released, dam_spilled, dam_losses, dam_evap, d.id AS station_id
							FROM dam_hourly dd
							LEFT JOIN m_dam d ON dd.dam_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE ((dd.dam_storage = -9999 OR dd.dam_storage = 999999)
								  OR (dd.dam_storage_percent = -9999 OR dd.dam_storage_percent = 999999)
								  OR (dd.dam_inflow = -9999 OR dd.dam_inflow = 999999)
								  OR (dd.dam_inflow_acc_percent = -9999 OR dd.dam_inflow_acc_percent = 999999)
								  OR (dd.dam_uses_water = -9999 OR dd.dam_uses_water = 999999)
								  OR (dd.dam_uses_water_percent = -9999 OR dd.dam_uses_water_percent = 999999)
								  OR (dd.dam_level = -9999 OR dd.dam_level = 999999)
								  OR (dd.dam_released = -9999 OR dd.dam_released = 999999)
								  OR (dd.dam_spilled = -9999 OR dd.dam_spilled = 999999)
								  OR (dd.dam_losses = -9999 OR dd.dam_losses = 999999)
								  OR (dd.dam_evap = -9999 OR dd.dam_evap = 999999)) AND (dd.deleted_by IS NULL) `

var SQLSelectIgnoreData = ` SELECT NOW()
								, 'dam_hourly' AS data_category
								, d.id AS station_id
								, d.dam_oldcode
								, d.dam_name
								, g.province_name
								, agt.agency_shortname
								--, agt.agency_name
								, dd.id
								, dd.dam_datetime
								, '#{Remarks}' AS remark
								, #{UserID} AS user_created
								, #{UserID} AS user_updated
								, NOW()
								, NOW()
								, dam_storage AS data_value
							FROM dam_hourly dd
							LEFT JOIN m_dam d ON dd.dam_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE true `

// WHERE d.is_ignore <> $1 `
