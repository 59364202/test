// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

// model for public.temperature
package temperature

var sqlTemperature = `SELECT * FROM latest.v_main_temperature WHERE tele_station_lat is not null`

var sqlTemperatureMaxMinByRegion = ` SELECT * FROM latest.v_main_temperature_maxmin_by_tmd_region `

var sqlTemperatureGraph = `SELECT data.id,gs.datetime, temp_value
	FROM public.temperature data
	INNER JOIN m_tele_station m ON m.id = data.tele_station_id  AND data.tele_station_id = $1
	RIGHT JOIN ( SELECT generate_series ($2::date, $3, '1 hour' ) AS datetime ) gs 
		ON data.temp_datetime between $2 AND $3
		AND data.temp_datetime = gs.datetime  AND data.deleted_at = to_timestamp (0)  AND ( qc_status IS NULL OR qc_status ->> 'is_pass' = 'true' ) 
	ORDER BY gs.datetime ASC `
