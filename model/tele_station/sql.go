package tele_station

import ()

var SQL_GetTeleStaion = " SELECT ts.id, ts.tele_station_name, ts.tele_station_oldcode, ts.tele_station_lat, ts.tele_station_long, lg.geocode, lg.province_name, lg.province_code, ts.agency_id, ts.tele_station_type " +
	" FROM m_tele_station ts " +
	" INNER JOIN lt_geocode lg ON ts.geocode_id = lg.id "

var SQL_GetTeleStaion_OrderBy = " ORDER BY tele_station_name->>'th' "
var SQL_GetTeleStation_OrderByProvince = "ORDER BY province_code"

var SQL_GetTeleStationById = `SELECT id, subbasin_id, agency_id, geocode_id, tele_station_name, tele_station_lat, 
       tele_station_long, tele_station_oldcode, tele_station_type, left_bank, 
       right_bank, ground_level FROM public.m_tele_station WHERE id = $1`

var sqlGetTeleStationByDataType = ` SELECT ts.id
								   , tele_station_oldcode
								   , tele_station_name
								   , agency_id
								   , geocode_id
								   , tele_station_lat
								   , tele_station_long
								   , tele_station_type
								   , subbasin_id
								   , agency_shortname
								   , agency_name
							  FROM m_tele_station ts
							  LEFT JOIN agency agt ON ts.agency_id =  agt.id `
var sqlGetTeleStationByDataTypeOrderBy = ` ORDER BY tele_station_name->>'th', agt.agency_shortname->>'en', tele_station_oldcode `

// https://wiki.postgresql.org/wiki/Loose_indexscan
//var sqlConditionTeleStationByRainfall = ` ts.id IN (
//	WITH RECURSIVE t AS (
//	   (SELECT tele_station_id FROM rainfall ORDER BY tele_station_id LIMIT 1)
//	   UNION ALL
//	   SELECT (SELECT tele_station_id FROM rainfall WHERE tele_station_id > t.tele_station_id ORDER BY tele_station_id LIMIT 1)
//	   FROM t
//	   WHERE t.tele_station_id IS NOT NULL
//	)
//	SELECT tele_station_id FROM t WHERE tele_station_id IS NOT NULL
//)`
//	เปลี่ยนไปใช้ tele_station_id จาก lastest.rainfall แทน
var sqlConditionTeleStationByRainfall = ` ts.id IN (
	SELECT tele_station_id FROM latest.rainfall WHERE tele_station_id IS NOT NULL
)`

//var sqlConditionTeleStationByWaterLevel = ` ts.id IN (
//	WITH RECURSIVE t AS (
//	   (SELECT tele_station_id FROM tele_waterlevel  ORDER BY tele_station_id LIMIT 1)
//	   UNION ALL
//	   SELECT (SELECT tele_station_id FROM tele_waterlevel WHERE tele_station_id > t.tele_station_id ORDER BY tele_station_id LIMIT 1)
//	   FROM t
//	   WHERE t.tele_station_id IS NOT NULL
//	)
//	SELECT tele_station_id FROM t WHERE tele_station_id IS NOT NULL
//)`
//	เปลี่ยนไปใช้ tele_station_id จาก lastest.tele_waterlevel แทน
var sqlConditionTeleStationByWaterLevel = ` ts.id IN (
	SELECT tele_station_id FROM latest.tele_waterlevel WHERE tele_station_id IS NOT NULL
)`

var sqlGetWaterlevelCanalStation = ` SELECT geo.province_code
											 , geo.province_name
											 , ss.station_id
											 , ss.station_oldcode
											 , ss.station_name
											 --, ss.agency_id
											 --, ss.geocode_id
											 --, ss.subbasin_id
											 , ss.station_lat
											 , ss.station_long
									 FROM (
											SELECT 't_' || CAST(ts.id AS TEXT) AS station_id, ts.tele_station_name::jsonb AS station_name, agency_id, geocode_id, subbasin_id, tele_station_oldcode AS station_oldcode, tele_station_lat AS station_lat, tele_station_long AS station_long
											FROM m_tele_station ts
											WHERE ts.is_ignore = false AND ts.deleted_at = '1970-01-01 07:00:00+07' AND (upper(ts.tele_station_type) = 'A' OR upper(ts.tele_station_type) = 'W' OR ts.tele_station_type IS NULL)
											UNION
											SELECT 'c_' || CAST(cs.id AS TEXT) AS station_id, cs.canal_station_name::jsonb AS station_name, agency_id, geocode_id, subbasin_id, canal_station_oldcode, canal_station_lat, canal_station_long
											FROM m_canal_station cs
											WHERE cs.is_ignore = false AND cs.deleted_at = '1970-01-01 07:00:00+07') ss
									 LEFT JOIN lt_geocode geo ON ss.geocode_id = geo.id `
var sqlGetWaterlevelCanalStation_OrderBy = ` ORDER BY province_code, station_name->'th' `

// watergate สถานี ปตร./ฝาย
var sqlGetWeirStation = `SELECT lt.province_code, lt.province_name, m_tele_station.id, tele_station_oldcode, tele_station_name, 
	tele_station_lat, tele_station_long, geocode_id
FROM latest.tele_watergate wg LEFT JOIN  m_tele_station ON wg.tele_station_id = m_tele_station.id 
LEFT JOIN lt_geocode lt  ON m_tele_station.geocode_id = lt.id  
WHERE  m_tele_station.deleted_at=to_timestamp(0) AND province_name IS NOT NULL
ORDER  BY lt.province_code, tele_station_oldcode`

// watergate สถานี ปตร./ฝาย
var SQL_GetWeirStation = ` SELECT ts.id, ts.tele_station_name, ts.tele_station_oldcode, ts.tele_station_lat, ts.tele_station_long, lg.geocode, lg.province_name, lg.province_code, ts.agency_id 
FROM latest.tele_watergate wg LEFT JOIN  m_tele_station ts ON wg.tele_station_id = ts.id
INNER JOIN lt_geocode lg ON ts.geocode_id = lg.id `
