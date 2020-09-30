// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package humid

var sqlGraph = `SELECT data.id,gs.datetime, humid_value
	FROM public.humid data
	INNER JOIN m_tele_station m ON m.id = data.tele_station_id  AND data.tele_station_id = $1
	RIGHT JOIN ( SELECT generate_series ($2::date, $3, '1 hour' ) AS datetime ) gs 
		ON data.humid_datetime between $2 AND $3
		AND data.humid_datetime = gs.datetime  AND data.deleted_at = to_timestamp (0)  AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
	ORDER BY gs.datetime ASC `

var sqlHumid = `SELECT 
        dd.humid_datetime,
        humid_value,
        d.id AS station_id,
        d.tele_station_oldcode,
        d.tele_station_name,
        d.tele_station_lat,
        d.tele_station_long,
        g.tmd_area_code AS area_code,
        g.tmd_area_name AS area_name,
        g.province_code,
        g.amphoe_code,
        g.tumbon_code,
        g.province_name,
        g.amphoe_name,
        g.tumbon_name,
        d.agency_id,
        agency_name
   FROM ((latest.humid dd
     LEFT JOIN m_tele_station d ON ((dd.tele_station_id = d.id)))
     LEFT JOIN agency agt ON ((d.agency_id = agt.id))
     LEFT JOIN lt_geocode g ON ((d.geocode_id = g.id)))
  WHERE 
  humid_datetime::date = CURRENT_DATE AND ((d.is_ignore = false) AND ((dd.qc_status IS NULL) OR ((dd.qc_status ->> 'is_pass'::text) = 'true'::text)))
`
