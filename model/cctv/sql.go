package cctv

var (
	getCCTV = "SELECT pcctv.id,m.tele_station_name,basin_name,pcctv.geocode_id,lt.geocode,lt.tumbon_code,lt.tumbon_name,lt.amphoe_code,lt.amphoe_name" +
		",lt.province_code,lt.province_name,lt.area_code,lt.area_name,dam_id_rid,tele_station_id, pcctv.basin_id, cctv_lat, cctv_long, cctv_title " +
		",cctv_description, cctv_mediatype, cctv_url,cctv_filename,cctv_flash,cctv_quicktime,b.basin_code, b.basin_name,pcctv.subbasin_id, is_active " +
		"FROM public.cctv pcctv " +
		"LEFT JOIN m_tele_station m ON pcctv.tele_station_id::integer = m.id " +
		"LEFT JOIN lt_geocode lt  ON pcctv.geocode_id = lt.id " +
		"LEFT JOIN basin b ON  pcctv.basin_id = b.id " +
		"WHERE pcctv.deleted_at=to_timestamp(0) "
)
