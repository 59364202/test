package dam_daily

import (
	"time"

	"haii.or.th/api/util/datatype"
)

var arrDamDailyLastestColumn = []string{
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

var sqlGetDamDaily = ` SELECT id, dam_date
							  , dam_storage, dam_storage_percent
							  , dam_inflow, dam_inflow_acc_percent
							  , dam_uses_water, dam_uses_water_percent
							  , dam_level, dam_released, dam_spilled, dam_losses, dam_evap
					   FROM dam_daily
					   WHERE dam_id = $1
					     AND dam_date BETWEEN $2 AND $3
					     AND deleted_at = '1970-01-01 07:00:00+07' `

var sqlGetDamDailyOrderBy = ` ORDER BY dam_date DESC, dam_id `

var SQL_GetDamThailand = "SELECT dam_daily.dam_id, dam_daily.dam_date, dam_daily.dam_storage, dam_daily.dam_storage_percent, dam_daily.dam_inflow, dam_daily.dam_released, dam_daily.dam_uses_water, " +
	" dam_daily.dam_uses_water_percent, m_dam.geocode_id, m_dam.agency_id, m_dam.dam_lat, m_dam.dam_long, m.max_storage, m_dam.dam_name " +
	" FROM dam_daily " +
	" INNER JOIN m_dam ON dam_daily.dam_id = m_dam.id " +
	" WHERE dam_daily.dam_date = (select max(dam_daily.dam_date)from dam_daily) AND m_dam.agency_id = $1 " +
	" ORDER BY m_dam.id "

func SQL_GetDamGraph(column, year string) string {
	y := datatype.MakeInt(year)
	if y == 0 {
		y = 1
	}
	t := time.Now().AddDate(-1*int(y), 0, 0) // ย้อนหลัง x ปี

	// temporary by pass QC Rule
	// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
	return ` SELECT dam_date , ` + column + ` FROM dam_daily WHERE dam_date >= '` + t.Format("2006-01-02 15:04") + `'
	AND dam_id = $1  ORDER BY dam_date ASC `
}

var SQL_GetDamDailyLastest = `
	SELECT dd.id AS dam_daily_id, dd.dam_date
		, st.id AS dam_id, st.dam_name, st.dam_lat, st.dam_long, st.max_storage
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
		, geo.id AS geocode_id, geo.geocode, geo.rid_area_code, geo.rid_area_name, geo.province_name, geo.amphoe_name, geo.tumbon_name, geo.province_code, st.dam_oldcode
	FROM (SELECT dam_id, max(dam_date) AS dam_date
		  FROM dam_daily
		  WHERE deleted_at = '1970-01-01 07:00:00+07' `

// temporary by pass QC Rule
// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
var SQL_GetDamDailyLastest2 = ` GROUP BY dam_id) md
	INNER JOIN dam_daily dd ON dd.dam_id = md.dam_id AND dd.dam_date = md.dam_date
	INNER JOIN m_dam st ON st.id = dd.dam_id
	INNER JOIN agency agt ON agt.id = st.agency_id
	LEFT  JOIN subbasin sb ON sb.id = st.subbasin_id
	LEFT  JOIN basin b ON b.id = sb.basin_id
	LEFT  JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE dd.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.is_ignore = 'false' `

var SQL_GetDamDaily_ = `
	SELECT --replace AS dam_daily_id, dd.dam_date
		, st.id AS dam_id, st.dam_name, st.dam_lat, st.dam_long, st.max_storage
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
		, geo.id AS geocode_id, geo.geocode, geo.rid_area_code, geo.rid_area_name, geo.province_name, geo.amphoe_name, geo.tumbon_name, geo.province_code, st.dam_oldcode, st.subbasin_id
		, cctv.id, cctv.cctv_url, cctv.cctv_filename `

// temporary by pass QC Rule
// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
var SQL_GetDamDaily_2 = `
		INNER JOIN m_dam st ON st.id = dd.dam_id
		INNER JOIN agency agt ON agt.id = st.agency_id
		LEFT  JOIN subbasin sb ON sb.id = st.subbasin_id
		LEFT  JOIN basin b ON b.id = sb.basin_id
		LEFT  JOIN lt_geocode geo ON geo.id = st.geocode_id
		LEFT  JOIN cctv ON (cctv.dam_id_rid = dd.dam_id AND cctv.is_active = 'true')
		WHERE dd.deleted_at = '1970-01-01 07:00:00+07'
			AND st.deleted_at = '1970-01-01 07:00:00+07'
			AND st.is_ignore = 'false' `

var SQL_GetDamDailyLastestOrderBy = ` ORDER BY geo.rid_area_code, st.id `

var SQL_GetDamDailyByDate = `
	SELECT dd.dam_id, st.dam_name, dd.dam_date
		, dd.dam_inflow, dd.dam_storage, dd.dam_storage_percent
		, dd.dam_uses_water, dd.dam_uses_water_percent, dd.dam_released
		, st.agency_id
		, st.dam_lat, st.dam_long, st.max_storage
		, b.id AS basin_id, b.basin_name
		, st.geocode_id, geo.rid_area_code, geo.rid_area_name
		, geo.province_name, geo.amphoe_name, geo.tumbon_name
	FROM dam_daily dd
	INNER JOIN m_dam st ON st.id = dd.dam_id
	INNER JOIN subbasin sb ON sb.id = st.subbasin_id
	INNER JOIN basin b ON b.id = sb.basin_id
	INNER JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE dd.dam_date = '2016-10-23'
	  AND dd.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.agency_id = 12
	  AND b.id = 26
`

var SQL_GetDamDailyByDateOrderBy = ` ORDER BY geo.rid_area_code, dd.dam_id  `

var sqlUpdateToDeleteDamDaily = ` UPDATE dam_daily
								SET deleted_by = $1
								  , deleted_at = NOW()
								  , updated_by = $1
								  , updated_at = NOW() `

var sqlGetErrorDamDaily = ` SELECT dd.id
								  , dam_oldcode
								  , dam_date
								  , dam_name
								  , province_name
								  , agency_name
								  , agency_shortname
								  , dam_storage, dam_storage_percent
								  , dam_inflow, dam_inflow_acc_percent
								  , dam_uses_water, dam_uses_water_percent
								  , dam_level, dam_released, dam_spilled, dam_losses, dam_evap, d.id AS station_id
							FROM dam_daily dd
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
								, 'dam_daily' AS data_category
								, d.id AS station_id
								, d.dam_oldcode
								, d.dam_name
								, g.province_name
								, agt.agency_shortname
								--, agt.agency_name
								, dd.id
								, dd.dam_date
								, '#{Remarks}' AS remark
								, #{UserID} AS user_created
								, #{UserID} AS user_updated
								, NOW()
								, NOW()
								, dam_storage AS data_value
							FROM dam_daily dd
							LEFT JOIN m_dam d ON dd.dam_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE true `

// WHERE d.is_ignore <> $1 `

// temporary by pass QC Rule
// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
var sqlSelectDamDailyGraphAnalyst = `
	SELECT dd.dam_date,sum(dd.dam_storage),sum(dd.dam_inflow),sum(dd.dam_released)
		,sum(dd.dam_spilled),sum(dd.dam_losses),sum(dd.dam_evap),sum(dd.dam_uses_water)
		,sum(dam_inflow_avg),sum(dam_released_acc),sum(dam_inflow_acc)
	FROM public.dam_daily dd
	WHERE dd.deleted_at=to_timestamp(0)
	`

// temporary by pass QC Rule
// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
var sqlSelectFourDamLatest = `
	SELECT DISTINCT ON (md.dam_name::jsonb)dd.dam_date,dd.dam_id,md.dam_name::jsonb,dd.dam_storage,dd.dam_inflow,dd.dam_released, dd.dam_storage_percent, dd.dam_uses_water
	FROM public.dam_daily dd LEFT JOIN public.m_dam md ON md.id=dd.dam_id
	WHERE (dd.dam_id=$1 OR dd.dam_id=$2 OR dd.dam_id=$3 OR dd.dam_id=$4) AND dd.deleted_at=to_timestamp(0)
	GROUP BY dd.dam_date,dd.dam_id,md.dam_name::jsonb,dd.dam_storage,dd.dam_inflow,dd.dam_released,dd.dam_storage_percent,dd.dam_uses_water ORDER BY md.dam_name::jsonb,dd.dam_date DESC
	`

var sqlSelectMappingNearDam = `SELECT dam_id
		FROM mapping_near_dam
		WHERE mapping_near_dam.province_id = $1`

var SQL_GetDamDailyLastestProvince = `
	SELECT dd.id AS dam_daily_id, dd.dam_date
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
	FROM (SELECT dam_id, max(dam_date) AS dam_date
		  FROM dam_daily
		  WHERE deleted_at = '1970-01-01 07:00:00+07' `

var SQL_GetDamDailyLastestProvince_2 = ` GROUP BY dam_id) md
	INNER JOIN dam_daily dd ON dd.dam_id = md.dam_id AND dd.dam_date = md.dam_date
	INNER JOIN m_dam st ON st.id = dd.dam_id
	INNER JOIN agency agt ON agt.id = st.agency_id
	LEFT  JOIN subbasin sb ON sb.id = st.subbasin_id
	LEFT  JOIN basin b ON b.id = sb.basin_id
	LEFT  JOIN lt_geocode geo ON geo.id = st.geocode_id
	WHERE dd.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.deleted_at = '1970-01-01 07:00:00+07'
	  AND st.is_ignore = 'false' `

var SQL_GetDamDailyLastestProvinceOrderBy = ` ORDER BY dam_storage_percent DESC `

// สำหรับหน้า thaiwater main ข้อมูลเขื่อนรวม 6 ภาค
var SQL_GetDamDailySummary = `WITH main AS (
	SELECT
		dam_date,
		rid_area_name->>'th' as region,
		sum( dam_storage ) OVER ( PARTITION BY rid_area_code ) AS dam_storage,
		sum( dam_uses_water ) OVER ( PARTITION BY rid_area_code ) AS dam_uses_water,		
		sum( dam_storage_percent ) OVER ( PARTITION BY rid_area_code ) / count( 1 ) OVER ( PARTITION BY rid_area_code ) AS dam_storage_percent,
		sum( dam_uses_water_percent ) OVER ( PARTITION BY rid_area_code ) / count( 1 ) OVER ( PARTITION BY rid_area_code ) AS dam_uses_water_percent 		,
		g.geocode,
		g.area_code,
		g.province_code,
		g.amphoe_code,
		g.tumbon_code,
		g.area_name,
		g.province_name,
		g.amphoe_name,
		g.tumbon_name,
		rid_area_code,
		row_number () OVER ( PARTITION BY rid_area_code ) AS ROW
	FROM
		latest.dam_daily dd
		LEFT JOIN m_dam d ON dd.dam_id = d.id
		LEFT JOIN lt_geocode g ON d.geocode_id = g.id 
	WHERE d.agency_id = 12 
	)
	select * from main 	WHERE main.ROW = 1
	ORDER BY rid_area_code `

// สำหรับหน้า thaiwater main ข้อมูลน้ำใช้การของ 4 เขื่อนหลัก รายวัน
var SQL_GetDamDailySummary4Dam = `SELECT
		dam_date,
		sum( total_dam_storage ) OVER ( PARTITION BY dam_date ) AS dam_storage,
		sum( total_dam_uses_water ) OVER ( PARTITION BY dam_date ) AS dam_uses_water,		
		sum( total_dam_inflow ) OVER ( PARTITION BY dam_date ) / count( 1 ) OVER ( PARTITION BY dam_date ) AS dam_inflow,
		sum( total_dam_released ) OVER ( PARTITION BY dam_date ) / count( 1 ) OVER ( PARTITION BY dam_date ) AS dam_released,
		row_number () OVER ( PARTITION BY dam_date ) AS ROW
	FROM
		public.dam_daily_sum_by_date dd 
	ORDER BY dam_date DESC`
