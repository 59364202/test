package rainfall_today

import "time"

//	"strconv"
//	"time"

func Gen_SQL_SelectRain(p *Param_RainfallToday) (string, []interface{}) {
	var param []interface{}
	strSQL := `
	SELECT r.tele_station_id 
		, mts.tele_station_name 
		, r.rainfall_value 
		, r.rainfall_datetime 
		, lg.province_code AS province_code 
		, lg.province_name 
		, mts.tele_station_lat 
		, mts.tele_station_long 
		, lg.amphoe_name 
		, lg.tumbon_name 
		, a.agency_name 
		, a.agency_shortname 
		, r.id             AS data_id 
		, mts.tele_station_oldcode
		, mts.basin_id
		, bs.basin_code
		, bs.basin_name
		, bss.subbasin_code AS subbasin_id
	FROM   latest.rainfall_today r 
		INNER JOIN m_tele_station mts 
				ON r.tele_station_id = mts.id 
		INNER JOIN lt_geocode lg 
				ON mts.geocode_id = lg.id 
		INNER JOIN agency a 
				ON mts.agency_id = a.id 
		INNER JOIN basin bs 
				ON mts.basin_id = bs.id
		INNER JOIN subbasin bss
				ON mts.subbasin_id = bss.id
	WHERE  r.deleted_at = '1970-01-01 07:00:00+07' 
		AND mts.deleted_at = '1970-01-01 07:00:00+07' 
		AND rainfall_value IS NOT NULL 
		AND rainfall_value <> 999999 
		AND rainfall_value <> 9999 
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true')
		AND mts.tele_station_type NOT IN ('W','G')
		AND mts.is_ignore <> 't'
`
	if p.Order == "asc" {
		strSQL += " ORDER BY r.rainfall_value ASC NULLS LAST "
	} else if p.Order == "desc" {
		strSQL += " ORDER BY r.rainfall_value DESC NULLS LAST "
	}
	//	if p.DataLimit != 0 {
	//		param = append(param, p.DataLimit)
	//		strSQL += " LIMIT $" + strconv.Itoa(len(param))
	//	}
	return strSQL, param
}

func Gen_SQL_SelectRainGraph(p *Param_RainfallToday) (string, []interface{}) {
	var param []interface{}
	param = append(param, p.StationId)
	now := time.Now()
	DateStart := now.Format("2006-01-02") + " 07:00" // CURRENT_DATE + interval '7 hours' "
	DateEnd := now.Format(time.RFC3339Nano)        // NOW()
		
	param = append(param, DateStart)
	param = append(param, DateEnd)
	
	strSQL := `
	SELECT gs.datetime, d.rainfall_value 
	FROM (
		SELECT date_trunc('hour',rainfall_datetime)as rainfall_datetime, sum(rainfall_value) as rainfall_value FROM rainfall_today
		WHERE deleted_at = '1970-01-01 07:00:00+07' AND tele_station_id = $1
		AND rainfall_datetime BETWEEN $2 AND $3
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true')
		GROUP BY date_trunc('hour',rainfall_datetime)
		ORDER BY date_trunc('hour',rainfall_datetime)
	) d
	RIGHT JOIN (
  		SELECT generate_series($2::date, $3, '1 hour' ) as datetime
	) gs ON  gs.datetime = d.rainfall_datetime 
	`
	/* edit 2019-05-16 by permporn
	strSQL := `
	SELECT gs.datetime, d.rainfall_value 
	FROM (
		SELECT date_trunc('hour',rainfall_datetime)as rainfall_datetime, sum(rainfall_value) as rainfall_value FROM rainfall_today
		WHERE deleted_at = '1970-01-01 07:00:00+07' AND tele_station_id = $1
		AND rainfall_datetime BETWEEN '` + now.Format("2006-01-02") + ` 07:00' AND '` + now.Format("2006-01-02 15:04") + `'
		AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true')
		GROUP BY date_trunc('hour',rainfall_datetime)
		ORDER BY date_trunc('hour',rainfall_datetime)
	) d
	RIGHT JOIN (
  		SELECT generate_series('` + now.Format("2006-01-02") + ` 07:00::timestamp', ` + now.Format("2006-01-02 15:04") + `,  '1 hour' ) as datetime
	) gs ON  gs.datetime = d.rainfall_datetime 
	`*/
	return strSQL, param
}
