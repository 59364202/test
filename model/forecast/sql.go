package forecast

var (
	getMaxDateFloodForecast = `SELECT max(floodforecast_datetime) --replace
								as floodforecast_datetime 
								FROM public.floodforecast `

	getMaxDateFloodForecastOrderBy = " ORDER  BY floodforecast_datetime DESC NULLS LAST limit 1"

	getCpyInterval = "SELECT mffc.floodforecast_station_oldcode, ffc.id, ffc.floodforecast_datetime, ffc.floodforecast_value, mffc.id, mffc.floodforecast_station_name, mffc.floodforecast_station_lat, mffc.floodforecast_station_long, mffc.floodforecast_station_alarm, mffc.floodforecast_station_warning, mffc.floodforecast_station_unit, a.id, a.agency_name, a.agency_shortname FROM public.floodforecast ffc " +
		"LEFT JOIN public.m_floodforecast_station mffc ON ffc.floodforecast_station_id=mffc.id " +
		"LEFT JOIN public.agency a ON mffc.agency_id=a.id"

	getCpyIntervalOrderBy = " ORDER  BY floodforecast_station_id, floodforecast_datetime DESC NULLS LAST"

	getObserveWaterlevel = ""

	getSwanStation = "SELECT swan_name, swan_lat, swan_long FROM public.m_swan_station WHERE deleted_at=to_timestamp(0)"

	getSwanForecast = "SELECT DISTINCT ON (swan_station_id,swan_datetime) mss.swan_name, s.id, s.swan_datetime, s.swan_depth, s.swan_highsig, s.swan_direction, s.swan_period_top, s.swan_period_average, s.swan_windx, s.swan_windy, mss.id, mss.swan_lat, mss.swan_long " +
		" FROM public.swan s LEFT JOIN public.m_swan_station mss ON s.swan_station_id=mss.id WHERE swan_datetime > (SELECT DISTINCT ON (swan_datetime)swan_datetime FROM public.swan WHERE deleted_at=to_timestamp(0) ORDER  BY swan_datetime DESC NULLS LAST limit 1) - interval '7 day' "
	getSwanForecastOrderBy = "  ORDER  BY swan_station_id, swan_datetime NULLS LAST"
	
	getFloodForecast = "SELECT DISTINCT ON (sb.subbasin_name::jsonb)sb.subbasin_name::jsonb,ffs.floodforecast_station_name::json,ff.floodforecast_datetime,ff.floodforecast_value " +
		"FROM public.floodforecast ff " +
		"LEFT JOIN public.m_floodforecast_station ffs ON ffs.id=ff.floodforecast_station_id " +
		"LEFT JOIN public.subbasin sb ON sb.id=ffs.subbasin_id " +
		"WHERE ffs.subbasin_id=$1 and ff.deleted_at=to_timestamp(0) ORDER BY sb.subbasin_name::jsonb,sb.id"
)
