package rainfall_daily

import (
	"time"
)

func Gen_SQL_SelectRain(p *Param_RainfallDaily) (string, []interface{}) {
	var param []interface{}
	//	strSQL := `
	//	SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code,
	//	lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, lg.warning_zone, mts.basin_id, bs.basin_code, bs.basin_name, bss.subbasin_code AS subbasin_id
	//	FROM rainfall_daily r
	//	INNER JOIN (
	//	SELECT tele_station_id , max(rainfall_datetime) as rainfall_datetime
	//	FROM rainfall_daily
	//	WHERE deleted_at = '1970-01-01 07:00:00+07'
	//	AND rainfall_value IS NOT NULL
	//	AND rainfall_value <> 0
	//	AND rainfall_value <> 999999
	//	AND rainfall_value <> 9999
	//	AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') `

	// slow query need optimize (2020-02-20 inspect by Manorot)
	// strSQL := `
	// SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code,
	// lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, lg.warning_zone, mts.basin_id, bs.basin_code, bs.basin_name, bss.subbasin_code AS subbasin_id
	// FROM rainfall_daily r
	// INNER JOIN (
	// SELECT tele_station_id , max(rainfall_datetime) as rainfall_datetime
	// FROM rainfall_daily
	// WHERE deleted_at = '1970-01-01 07:00:00+07'
	// AND rainfall_value IS NOT NULL
	// AND rainfall_value <> 999999
	// AND rainfall_value <> 9999
	// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') `
	// if p.IsDaily {
	// 	ds := time.Now().AddDate(0, 0, -2).Format("2006-01-02 ") + "07:00"
	// 	de := time.Now().AddDate(0, 0, -1).Format("2006-01-02 ") + "07:00"
	// 	param = append(param, ds, de)
	// 	strSQL += " AND rainfall_datetime BETWEEN ($1) AND ($2) "
	// }
	// strSQL += ` GROUP BY tele_station_id
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

	strSQL := `
	SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code, 
	lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, lg.warning_zone, mts.basin_id, bs.basin_code, bs.basin_name, bss.subbasin_code AS subbasin_id
	FROM rainfall_daily r 
	INNER JOIN m_tele_station mts ON r.tele_station_id = mts.id 
	INNER JOIN lt_geocode lg ON mts.geocode_id = lg.id 
	INNER JOIN agency a ON mts.agency_id = a.id
	INNER JOIN basin bs ON mts.basin_id = bs.id
	INNER JOIN subbasin bss ON mts.subbasin_id = bss.id
	WHERE r.deleted_at = '1970-01-01 07:00:00+07'  `
	if p.IsDaily {
		ds := time.Now().AddDate(0, 0, -2).Format("2006-01-02 ") + "07:00"
		de := time.Now().AddDate(0, 0, -1).Format("2006-01-02 ") + "07:00"
		param = append(param, ds, de)
		strSQL += " AND r.rainfall_datetime BETWEEN ($1) AND ($2) "
	}
	strSQL += `AND r.rainfall_value IS NOT NULL
	AND r.rainfall_value <> 999999 
	AND r.rainfall_value <> 9999 
	AND ( r.qc_status IS NULL OR r.qc_status->>'is_pass' = 'true')
	AND mts.deleted_at = '1970-01-01 07:00:00+07'
	AND mts.tele_station_type NOT IN ('W','G')
	AND mts.is_ignore <> 't'
	ORDER BY r.rainfall_value DESC NULLS LAST`

	//	if p.Order == "asc" {
	//		strSQL += " ORDER BY r.rainfall_value ASC NULLS LAST "
	//	} else if p.Order == "desc" {
	//		strSQL += " ORDER BY r.rainfall_value DESC NULLS LAST "
	//	}
	//	if p.DataLimit != 0 {
	//		param = append(param, p.DataLimit)
	//		strSQL += " LIMIT $" + strconv.Itoa(len(param))
	//	}
	return strSQL, param
}

func Gen_SQL_SelectRainGraph(p *Param_RainfallDaily) (string, []interface{}) {
	var param []interface{}
	param = append(param, p.StationId)

	strSQL := `
	SELECT rainfall_datetime::date, sum(rainfall_value) as rainfall_value FROM rainfall_daily
	WHERE deleted_at = '1970-01-01 07:00:00+07' 
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') 
		AND tele_station_id = $1`
	var (
		gsFrom string
		gsTo   string
	)
	if p.IsDaily {
		param = append(param, p.StratDate)
		param = append(param, p.EndDate)
		gsFrom = "$2::timestamp"
		gsTo = "$3"
		strSQL += " AND rainfall_datetime BETWEEN $2 AND $3 "
	} else if p.IsMonthly {
		// param = append(param, p.Year)
		// param = append(param, p.Month)
		// strSQL += " AND date_part('year', rainfall_datetime) = $2 AND date_part('month', rainfall_datetime) = $3 "
		ds := time.Date(p.Year, time.Month(p.Month), 1, 0, 0, 0, 0, time.UTC)
		de := ds.AddDate(0, 1, -1)
		// param = append(param, p.Year)
		// param = append(param, p.Month)
		// strSQL += " AND date_part('year', rainfall_datetime) = $2 AND date_part('month', rainfall_datetime) = $3 "
		param = append(param, ds.Format("2006-01-02"))
		param = append(param, de.Format("2006-01-02")+" 23:59:59")
		strSQL += " AND rainfall_datetime BETWEEN $2 AND $3 "
		gsFrom = "$2::timestamp"
		gsTo = "$3"
	}
	//	AND rainfall_datetime BETWEEN $2 AND $3
	strSQL += `GROUP BY rainfall_datetime::date
	ORDER BY rainfall_datetime::date
	`
	strSQL = `
	SELECT gs.datetime, d.rainfall_value 
	FROM ( ` + strSQL + ` )d
	RIGHT JOIN (
  		SELECT generate_series(` + gsFrom + `, ` + gsTo + `, '1 day' ) as datetime
	) gs ON  gs.datetime = d.rainfall_datetime 
	`
	return strSQL, param
}
