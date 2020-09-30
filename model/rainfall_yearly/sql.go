package rainfall_yearly

//	"strconv"

func Gen_SQL_SelectRain(p *Param_RainfallYearly) (string, []interface{}) {
	var param []interface{}
	// strSql := `
	// SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code,
	// lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, a.agency_shortname, r.id as data_id, mts.tele_station_oldcode,
	// mts.basin_id,
	// bs.basin_code,
	// bs.basin_name,
	// bss.subbasin_code AS subbasin_id
	// FROM rainfall_yearly r
	// INNER JOIN (
	// SELECT tele_station_id , max(rainfall_datetime) as rainfall_datetime
	// FROM rainfall_yearly
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
	// `

	strSql := `
	SELECT r.tele_station_id, mts.tele_station_name, r.rainfall_value, r.rainfall_datetime, lg.province_code as province_code, 
	lg.province_name, mts.tele_station_lat, mts.tele_station_long, lg.amphoe_name, lg.tumbon_name, a.agency_name, a.agency_shortname, r.id as data_id, mts.tele_station_oldcode, 
	mts.basin_id,
	bs.basin_code,
	bs.basin_name,
	bss.subbasin_code AS subbasin_id
	FROM latest.rainfall_yearly r 
	INNER JOIN ( 
	SELECT tele_station_id , rainfall_datetime
	FROM rainfall_yearly
	WHERE deleted_at = '1970-01-01 07:00:00+07' 
	AND rainfall_value IS NOT NULL AND rainfall_value <> 0
	AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true')
	) tr ON r.tele_station_id = tr.tele_station_id AND r.rainfall_datetime = tr.rainfall_datetime 
	INNER JOIN m_tele_station mts ON r.tele_station_id = mts.id 
	INNER JOIN lt_geocode lg ON mts.geocode_id = lg.id 
	INNER JOIN agency a ON mts.agency_id = a.id 
	INNER JOIN basin bs ON mts.basin_id = bs.id 
	INNER JOIN subbasin bss ON mts.subbasin_id = bss.id
	WHERE r.deleted_at = '1970-01-01 07:00:00+07' 
	AND mts.deleted_at = '1970-01-01 07:00:00+07'
	AND mts.tele_station_type NOT IN ('W','G')
	AND mts.is_ignore <> 't'
	`
	if p.Order == "asc" {
		strSql += " ORDER BY r.rainfall_value ASC NULLS LAST "
	} else if p.Order == "desc" {
		strSql += " ORDER BY r.rainfall_value DESC NULLS LAST "
	}
	//	if p.DataLimit != 0 {
	//		param = append(param, p.DataLimit)
	//		strSql += " LIMIT $" + strconv.Itoa(len(param))
	//	}
	return strSql, param
}
