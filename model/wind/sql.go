//     Author: Thitiporn  Meeprasert <thitiporn@hii.or.th>
package wind

var sqlWind = `SELECT 
        dd.wind_datetime,
        wind_speed,wind_dir_value,wind_dir,
        d.id AS station_id,
        d.tele_station_oldcode,
        d.tele_station_name,
        d.tele_station_lat,
        d.tele_station_long,
        g.tmd_area_code AS area_code,
        g.province_code,
        g.amphoe_code,
        g.tumbon_code,
        g.tmd_area_name AS area_name,
        g.province_name,
        g.amphoe_name,
        g.tumbon_name,
        d.agency_id,
				agency_name
   FROM ((latest.wind dd
     LEFT JOIN m_tele_station d ON ((dd.tele_station_id = d.id)))
		 LEFT JOIN agency agt ON ((d.agency_id = agt.id))
     LEFT JOIN lt_geocode g ON ((d.geocode_id = g.id)))
  WHERE 
  wind_datetime::date = CURRENT_DATE AND ((d.is_ignore = false) AND ((dd.qc_status IS NULL) OR ((dd.qc_status ->> 'is_pass'::text) = 'true'::text)))
`
