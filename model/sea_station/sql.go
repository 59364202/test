package sea_station

import (

)

var (
	sqlSeaStationByAgency = "SELECT a.id,a.agency_name,a.agency_shortname,mss.id,mss.sea_station_lat,mss.sea_station_long,mss.sea_station_oldcode,sea_station_name FROM m_sea_station mss LEFT JOIN (SELECT sea_station_id FROM public.sea_waterlevel GROUP BY sea_station_id) ss ON mss.id=ss.sea_station_id LEFT JOIN public.agency a ON mss.agency_id=a.id"

	sqlSeaForecastStationByAgency = "SELECT a.id,a.agency_name,a.agency_shortname,mss.id,mss.sea_station_lat,mss.sea_station_long,mss.sea_station_oldcode,sea_station_name FROM m_sea_station mss LEFT JOIN (SELECT sea_station_id FROM public.sea_water_forecast GROUP BY sea_station_id) ss ON mss.id=ss.sea_station_id LEFT JOIN public.agency a ON mss.agency_id=a.id"

)