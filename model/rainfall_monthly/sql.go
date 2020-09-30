package rainfall_monthly

import "time"

//	"strconv"

func Gen_SQL_SelectRain(p *Param_RainfallMonthly) (string, []interface{}) {
	var param []interface{}

	// slow query require optimize (2020-02-21 inspect by manorot)
	// strSql := `
	// SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code,
	// lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, a.agency_shortname, r.id as data_id, mts.tele_station_oldcode,
	// mts.basin_id,
	// bs.basin_code,
	// bs.basin_name,
	// bss.subbasin_code AS subbasin_id
	// FROM rainfall_monthly r
	// INNER JOIN (
	// SELECT tele_station_id , max(rainfall_datetime) as rainfall_datetime
	// FROM rainfall_monthly
	// WHERE deleted_at = '1970-01-01 07:00:00+07'
	// AND rainfall_value IS NOT NULL AND rainfall_value <> 0
	// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') `
	// strSql += ` GROUP BY tele_station_id
	// ) tr ON r.tele_station_id = tr.tele_station_id AND r.rainfall_datetime = tr.rainfall_datetime
	// INNER JOIN m_tele_station mts ON r.tele_station_id = mts.id
	// INNER JOIN lt_geocode lg ON mts.geocode_id = lg.id
	// INNER JOIN agency a ON mts.agency_id = a.id
	// INNER JOIN basin bs ON mts.basin_id = bs.id
	// INNER JOIN subbasin bss ON mts.subbasin_id = bss.id
	// WHERE r.deleted_at = '1970-01-01 07:00:00+07'
	// AND mts.deleted_at = '1970-01-01 07:00:00+07'
	// AND mts.tele_station_type NOT IN ('W','G')
	// AND mts.is_ignore <> 't'
	// ORDER BY r.rainfall_value DESC NULLS LAST
	// `

	strSql := `
	SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code, 
	lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, a.agency_shortname, r.id as data_id, mts.tele_station_oldcode, 
	mts.basin_id,
	bs.basin_code,
	bs.basin_name,
	bss.subbasin_code AS subbasin_id 
	FROM rainfall_monthly r 
	INNER JOIN ( 
	SELECT tele_station_id , rainfall_datetime 
		FROM latest.rainfall_monthly
		WHERE deleted_at = '1970-01-01 07:00:00+07' 
		AND rainfall_value IS NOT NULL AND rainfall_value <> 0
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true')
	) tr
	ON r.tele_station_id = tr.tele_station_id AND r.rainfall_datetime = tr.rainfall_datetime 
	INNER JOIN m_tele_station mts ON r.tele_station_id = mts.id 
	INNER JOIN lt_geocode lg ON mts.geocode_id = lg.id 
	INNER JOIN agency a ON mts.agency_id = a.id
	INNER JOIN basin bs ON mts.basin_id = bs.id
	INNER JOIN subbasin bss ON mts.subbasin_id = bss.id 
	WHERE r.deleted_at = '1970-01-01 07:00:00+07' 
	AND mts.deleted_at = '1970-01-01 07:00:00+07'
	AND mts.tele_station_type NOT IN ('W','G')
	AND mts.is_ignore <> 't'
	ORDER BY r.rainfall_value DESC NULLS LAST
	`
	//	if p.Order == "asc" {
	//		strSql += " ORDER BY r.rainfall_value ASC NULLS LAST "
	//	} else if p.Order == "desc" {
	//		strSql += " ORDER BY r.rainfall_value DESC NULLS LAST "
	//	}
	//	if p.DataLimit != 0 {
	//		param = append(param, p.DataLimit)
	//		strSql += " LIMIT $" + strconv.Itoa(len(param))
	//	}
	return strSql, param
}

func Gen_SQL_SelectRainGraph(p *Param_RainfallMonthly) (string, []interface{}) {
	var param []interface{}
	param = append(param, p.StationId)
	strSql := `
	SELECT rainfall_datetime, rainfall_datetime::timestamp + INTERVAL '1 month' - INTERVAL '1 day' AS rainfall_datetime2, rainfall_value, day_count FROM rainfall_monthly
	WHERE deleted_at = '1970-01-01 07:00:00+07' AND tele_station_id = $1 
    AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') `
	var (
		gsFrom string
		gsTo   string
	)
	if p.Year > 0 {
		// param = append(param, p.Year)
		// strSql += " AND date_part('year', rainfall_datetime) = $2 "
		// gsFrom = " make_date($2::int, 1, 1) "
		// gsTo = " make_date($2::int, 12, 1) "
		ds := time.Date(p.Year, time.January, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		de := time.Date(p.Year, time.December, 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		strSql += " AND rainfall_datetime BETWEEN '" + ds + "' AND '" + de + " 23:59:59'"
		gsFrom = "'" + ds + "'::timestamp"
		gsTo = "'" + de + "'"
	}
	//	AND rainfall_datetime BETWEEN $2 AND $3
	strSql += ` ORDER BY rainfall_datetime `

	strSql = `
	SELECT gs.datetime, d.rainfall_value, d.day_count
	FROM ( ` + strSql + ` )d
	RIGHT JOIN (
  		SELECT generate_series(` + gsFrom + `, ` + gsTo + `, '1 month' ) + interval '1 month' - interval '1 day' as datetime
	) gs ON  gs.datetime = d.rainfall_datetime2 
	`
	return strSql, param
}
