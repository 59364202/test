package sea_waterlevel

import ()

var (
	sqlSelectSeaWaterlevelLatest = "SELECT DISTINCT ON (sea_station_id) " +
		"sea_station_id, sea_station_name, waterlevel_datetime, waterlevel_value,sea_station_lat,sea_station_long,sea_station_oldcode, a.id, a.agency_name,a.agency_shortname, gc.id, area_code,area_name,province_code,province_name,amphoe_code,amphoe_name,tumbon_code,tumbon_name " +
		"FROM public.sea_waterlevel sw LEFT JOIN public.m_sea_station msw ON sw.sea_station_id=msw.id LEFT JOIN public.agency a ON msw.agency_id=a.id LEFT JOIN public.lt_geocode gc ON msw.geocode_id=gc.id " +
		"WHERE sw.deleted_at=to_timestamp(0) " +
		"ORDER BY sea_station_id, waterlevel_datetime DESC"

	sqlSelectSeaForecast = "SELECT m.id,m.sea_station_name,a.datetime,a.seaforecast_value,sea_station_lat,sea_station_long,sea_station_oldcode, ag.id, ag.agency_name,ag.agency_shortname, " +
		"gc.id,area_code,area_name,province_code,province_name,amphoe_code,amphoe_name,tumbon_code,tumbon_name  FROM public.m_sea_station m LEFT JOIN " +
		"(SELECT gs.datetime,seaforecast_value  " +
		"FROM (SELECT generate_series($2::timestamp,$3::timestamp,'4 hour') AS datetime)  " +
		"gs LEFT JOIN public.sea_water_forecast sw ON gs.datetime=sw.seaforecast_datetime  " +
		"LEFT JOIN public.m_sea_station m ON m.id=sw.sea_station_id AND sw.sea_station_id=$1 " +
		"ORDER BY gs.datetime ASC) a ON m.id=$1" + 
		"LEFT JOIN public.agency ag ON m.agency_id=ag.id " + 
		"LEFT JOIN public.lt_geocode gc ON m.geocode_id=gc.id"

	sqlSelectSeaWaterlevelRealByStation = "SELECT m.sea_station_name, a.datetime, a.waterlevel_value FROM (SELECT gs.datetime, sw.waterlevel_value " +
		"FROM public.sea_waterlevel sw  " +
		"INNER JOIN public.m_sea_station m ON m.id = sw.sea_station_id AND sw.sea_station_id=$1 " +
		"RIGHT JOIN ( select generate_series($2::date,$3, '5 min') as datetime  " +
		")gs ON sw.waterlevel_datetime=gs.datetime AND sw.deleted_at=to_timestamp(0)  " +
		"ORDER BY gs.datetime DESC) a  " +
		"LEFT JOIN public.m_sea_station m ON m.id=$1 " +
		"WHERE CASE WHEN m.agency_id = 6 THEN  " +
		"( date_part('minute', a.datetime )::integer % 60) = 0  " +
		"WHEN m.agency_id = 3 THEN	 " +
		"( date_part('minute', a.datetime )::integer % 15) = 0  " +
		"WHEN m.agency_id = 10 THEN	  " +
		"( date_part('minute', a.datetime )::integer % 15) = 0  " +
		"WHEN m.agency_id = 9 THEN	 " +
		"( date_part('minute', a.datetime )::integer % 10) = 0  " +
		"END AND m.sea_station_name IS NOT NULL"
		
	sqlSelectSeaWaterlevelForecastByStation = "SELECT m.sea_station_name, a.datetime, a.seaforecast_value FROM (SELECT gs.datetime, sw.seaforecast_value " +
		"FROM public.sea_water_forecast sw  " +
		"INNER JOIN public.m_sea_station m ON m.id = sw.sea_station_id AND sw.sea_station_id=$1 " +
		"RIGHT JOIN ( select generate_series($2::date,$3, '5 min') as datetime  " +
		")gs ON sw.seaforecast_datetime=gs.datetime AND sw.deleted_at=to_timestamp(0)  " +
		"ORDER BY gs.datetime DESC) a  " +
		"LEFT JOIN public.m_sea_station m ON m.id=$1 " +
		"WHERE CASE WHEN m.agency_id = 6 THEN  " +
		"( date_part('minute', a.datetime )::integer % 60) = 0  " +
		"WHEN m.agency_id = 3 THEN	 " +
		"( date_part('minute', a.datetime )::integer % 15) = 0  " +
		"WHEN m.agency_id = 10 THEN	  " +
		"( date_part('minute', a.datetime )::integer % 15) = 0  " +
		"WHEN m.agency_id = 9 THEN	 " +
		"( date_part('minute', a.datetime )::integer % 10) = 0  " +
		"END"	
)
