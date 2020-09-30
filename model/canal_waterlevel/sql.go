// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package canal_waterlevel is a model for public.canal_waterlevel table. This table store canal_waterlevel.
package canal_waterlevel

var SQL_GetCanalWaterLevelByStationAndDate = ` SELECT id
											  , canal_waterlevel_datetime
											  , canal_waterlevel_value
									 FROM canal_waterlevel
									 WHERE canal_station_id = $1
									   AND canal_waterlevel_datetime BETWEEN $2 AND $3
									   AND deleted_at = '1970-01-01 07:00:00+07' `

var sqlSelectCanal = "SELECT m.id,a.* FROM (SELECT gs.datetime,cw.canal_waterlevel_value " +
	"FROM  (select generate_series($2::date,$3, '15 min') as datetime " +
	") gs LEFT JOIN public.canal_waterlevel cw ON cw.canal_waterlevel_datetime BETWEEN $2 AND $3 AND gs.datetime=cw.canal_waterlevel_datetime AND cw.deleted_at=to_timestamp(0) " +
	" AND canal_station_id = $1 AND (qc_status IS NULL OR qc_status->>'is_pass' = 'true') " +
	"ORDER BY gs.datetime ASC )a  " +
	"LEFT JOIN m_canal_station m ON m.id=$1"
