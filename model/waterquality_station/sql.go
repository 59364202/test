package waterquality_station

import ()

var SQL_GetAllWaterQualityStation = "SELECT ws.id, ws.agency_id, ws.geocode_id, ws.waterquality_station_name, ws.waterquality_station_lat, " +
	" ws.waterquality_station_long, ws.waterquality_station_oldcode, ws.is_active " +
	" FROM m_waterquality_station ws " +
	" INNER JOIN lt_geocode lg ON ws.geocode_id = lg.id " +
	" WHERE ws.deleted_at = '1970-01-01 07:00:00+07' "
