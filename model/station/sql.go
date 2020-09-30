package station

import ()

var sqlGetDam = ` SELECT a.id
					, a.dam_oldcode
					, a.dam_name
					, agt.agency_name
					, agt.agency_shortname
					, a.dam_lat
					, a.dam_long
					, g.geocode
					, sb.subbasin_name
				FROM m_dam a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				LEFT JOIN subbasin sb ON sb.id = a.subbasin_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.dam_oldcode `

var sqlGetMediumDam = ` SELECT a.id
					, a.mediumdam_oldcode
					, a.mediumdam_name
					, agt.agency_name
					, agt.agency_shortname
					, a.mediumdam_lat
					, a.mediumdam_long
					, g.geocode
					, sb.subbasin_name
				FROM m_medium_dam a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				LEFT JOIN subbasin sb ON sb.id = a.subbasin_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.mediumdam_oldcode `

var sqlGetGround = ` SELECT a.id
					, a.ground_oldcode
					, agt.agency_name
					, agt.agency_shortname
					, a.ground_lat
					, a.ground_long
					, g.geocode
				FROM m_ground a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.ground_oldcode `

var sqlGetFloodForecastStation = ` SELECT a.id
					, a.floodforecast_station_oldcode
					, a.floodforecast_station_name
					, agt.agency_name
					, agt.agency_shortname
					, a.floodforecast_station_lat
					, a.floodforecast_station_long
					, g.geocode
					, sb.subbasin_name
				FROM m_floodforecast_station a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				LEFT JOIN subbasin sb ON sb.id = a.subbasin_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.floodforecast_station_oldcode `

var sqlGetFordStation = ` SELECT a.id
					, a.ford_station_oldcode
					, a.ford_station_name
					, agt.agency_name
					, agt.agency_shortname
					, a.ford_station_lat
					, a.ford_station_long
					, g.geocode
					, '' AS subbasin_name
				FROM m_ford_station a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.ford_station_oldcode `

var sqlGetSwanStation = ` SELECT a.id
					, a.swan_oldcode
					, a.swan_name
					, agt.agency_name
					, agt.agency_shortname
					, a.swan_lat
					, a.swan_long
					, g.geocode
					, '' AS subbasin_name
				FROM m_swan_station a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.swan_oldcode `

var sqlGetTeleStation = ` SELECT a.id
					, a.tele_station_oldcode
					, a.tele_station_name
					, agt.agency_name
					, agt.agency_shortname
					, a.tele_station_lat
					, a.tele_station_long
					, g.geocode
					, sb.subbasin_name
				FROM m_tele_station a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				LEFT JOIN subbasin sb ON sb.id = a.subbasin_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.tele_station_oldcode `

var sqlGetWaterQualityStation = ` SELECT a.id
					, a.waterquality_station_oldcode
					, a.waterquality_station_name
					, agt.agency_name
					, agt.agency_shortname
					, a.waterquality_station_lat
					, a.waterquality_station_long
					, g.geocode
					, '' AS subbasin_name
				FROM m_waterquality_station a
				LEFT JOIN agency agt ON agt.id = a.agency_id
				LEFT JOIN lt_geocode g ON g.id = a.geocode_id
				WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
				ORDER BY a.waterquality_station_oldcode `

var sqlGetAirStation = ` SELECT a.id
							, a.air_station_oldcode
							, a.air_station_name
							, agt.agency_name
							, agt.agency_shortname
							, a.air_station_lat
							, a.air_station_long
							, g.geocode
							, '' AS subbasin_name
						FROM m_air_station a
						LEFT JOIN agency agt ON agt.id = a.agency_id
						LEFT JOIN lt_geocode g ON g.id = a.geocode_id
						WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
						ORDER BY a.air_station_oldcode `

var sqlGetCanalStation = ` SELECT a.id
							, a.canal_station_oldcode
							, a.canal_station_name
							, agt.agency_name
							, agt.agency_shortname
							, a.canal_station_lat
							, a.canal_station_long
							, g.geocode
							, sb.subbasin_name
						FROM m_canal_station a
						LEFT JOIN agency agt ON agt.id = a.agency_id
						LEFT JOIN lt_geocode g ON g.id = a.geocode_id
						LEFT JOIN subbasin sb ON sb.id = a.subbasin_id
						WHERE a.#{ColumnName} IS NULL AND a.agency_id = $1
						ORDER BY a.canal_station_oldcode `
