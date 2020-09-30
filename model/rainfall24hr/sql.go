package rainfall24hr

import (
	"strconv"
	"time"
)

var SQLSelectIgnoreData = ` SELECT NOW()
								, 'rainfall_24h' AS data_category
								, d.id AS station_id
								, d.tele_station_oldcode
								, d.tele_station_name
								, g.province_name
								, agt.agency_shortname
								--, agt.agency_name
								, dd.id
								, dd.rainfall_datetime
								, '#{Remarks}' AS remark
								, #{UserID} AS user_created
								, #{UserID} AS user_updated
								, NOW()
								, NOW()
								, rainfall24h AS data_value
							FROM rainfall_24h dd 	   
							LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id
							LEFT JOIN agency agt ON d.agency_id = agt.id
							LEFT JOIN lt_geocode g ON d.geocode_id = g.id
							WHERE true `

// WHERE d.is_ignore <> $1 `

func SQL_AdvRainDiagram(agency_id []int64, date string) (string, []interface{}) {
	var text = ""
	var (
		d   time.Time
		err error
	)
	d, err = time.Parse("2006-01-02", date)
	if err != nil {
		d, _ = time.Parse("2006-01-02 15:04", date)
	}
	d = d.Add(32*time.Hour + 40*time.Minute) //   + interval ' 1 day 8 hour 40 minute'
	var p = []interface{}{d.Format("2006-01-02 15:04:05")}
	for i, v := range agency_id {
		if i != 0 {
			text += ","
		}
		text += "$" + strconv.Itoa(i+2)
		p = append(p, v)
	}
	var q = `
SELECT	   m.tele_station_name, 
           m.tele_station_lat, 
           m.tele_station_long, 
           r24.rainfall24h 
FROM       rainfall_24h r24 
INNER JOIN m_tele_station m 
ON         r24.tele_station_id = m.id  AND m.agency_id IN (` + text + `)
--WHERE      r24.rainfall_datetime = $1::timestamp + interval ' 1 day 8 hour 40 minute'
WHERE      r24.rainfall_datetime = $1
AND        r24.rainfall24h IS NOT NULL
AND 	   r24.deleted_at = to_timestamp(0)
AND 	   (qc_status IS NULL OR qc_status->>'is_pass' = 'true')
	`
	return q, p
}

//var SQL_AdvRainDiagram = `
//SELECT     m.tele_station_name,
//           m.tele_station_lat,
//           m.tele_station_long,
//           r24.rainfall24h
//FROM       rainfall_24h r24
//INNER JOIN m_tele_station m
//ON         r24.tele_station_id = m.id  AND m.agency_id IN ($1::integer[])
//WHERE      r24.rainfall_datetime = $2::timestamp + interval ' 1 day 8 hour 40 minute'
//AND        r24.rainfall24h IS NOT NULL
//AND r24.deleted_at = to_timestamp(0)
//`
var SQL_AdvOnload_Agency = `
SELECT m.agency_id, 
        a.agency_name :: text, 
        Max(Date_part('year', r24.rainfall_datetime)), 
        Min(Date_part('year', r24.rainfall_datetime)) 
FROM   rainfall_24h r24 
        INNER JOIN m_tele_station m 
                ON r24.tele_station_id = m.id 
                AND r24.deleted_at = To_timestamp(0) 
        INNER JOIN agency a 
			    ON m.agency_id = a.id 
WHERE (qc_status IS NULL OR qc_status->>'is_pass' = 'true')
GROUP  BY m.agency_id, 
    		a.agency_name :: text 
ORDER  BY m.agency_id 
`

func Gen_SQLAdvRainSum(p *Param_AdvRainSum) (string, []interface{}) {
	var q string = ""
	var itf []interface{} = make([]interface{}, 0)

	q = `
SELECT   St_x(St_transform((St_setsrid(St_point(s.tele_station_long,s.tele_station_lat), 4326)), 32647)) AS x,
         St_y(St_transform((St_setsrid(St_point(s.tele_station_long,s.tele_station_lat), 4326)), 32647)) AS y,
         Sum(r.rainfall24h)                                                                              AS sum,
         Count(r.rainfall24h)                                                                            AS count
FROM     m_tele_station s, 
         rainfall_24h r 
WHERE    r.tele_station_id = s.id 
AND      r.rainfall24h < 300 
AND      r.rainfall24h >= 0 
AND      r.rainfall_datetime >= $2
AND      r.rainfall_datetime <= $3
AND      (r.qc_status IS NULL OR r.qc_status->>'is_pass' = 'true')
AND      date_part('hour', r.rainfall_datetime) = 7
AND      s.tele_station_oldcode IN 
         ( 
                SELECT s.tele_station_oldcode 
                FROM   m_tele_station s 
                WHERE  St_contains(St_polygonfromtext($4 ,32647),St_transform(St_setsrid(St_point(s.tele_station_long,s.tele_station_lat), 4326), 32647))
                AND    agency_id = $1 ) 
GROUP BY s.id, 
         s.tele_station_long, 
         s.tele_station_lat 
HAVING   Count(r.rainfall24h)>(( $3::date - $2::date)* 0.5)
	`
	itf = append(itf, p.AgencyId)
	itf = append(itf, p.DateStart)
	itf = append(itf, p.DateEnd)
	itf = append(itf, p.Boundary)

	return q, itf
}
