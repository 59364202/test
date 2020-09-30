package lt_geocode

var SQL_selectGeocodeFromGeocode = " SELECT id, geocode, area_code, province_code, amphoe_code, tumbon_code, area_name, province_name, amphoe_name, tumbon_name " +
	" FROM public.lt_geocode "

var SQL_selectGeocodeFromGeocode_AllProvince = " WHERE (amphoe_code = '  ' AND tumbon_code = '  ' OR amphoe_name->>'th' = '' AND tumbon_name->>'th' = '' OR geocode = '999999')"

var SQL_selectGeocodeFromGeocode_OrderBy = " ORDER BY province_name->>'th' "

var sqlGetAllArea = "SELECT tmd_area_code,tmd_area_name_json::jsonb FROM public.lt_geocode WHERE deleted_at=to_timestamp(0) AND tmd_area_name_json->>'th' != '' AND tmd_area_code != ' ' AND tmd_area_code IS NOT NULL GROUP BY tmd_area_code,tmd_area_name_json::jsonb ORDER BY tmd_area_name_json"

var sqlGetAllProvince = `
	SELECT province_code,province_name::jsonb,mt.tele_station_oldcode,tele_station_name,mt.id, rmax.max, rmin.min FROM public.lt_geocode ge LEFT JOIN m_tele_station mt ON mt.geocode_id=ge.id 
	LEFT JOIN (SELECT tele_station_id,MAX(rainfall_datetime) FROM public.rainfall_monthly WHERE deleted_at=to_timestamp(0) AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') GROUP BY tele_station_id ) rmax ON rmax.tele_station_id=mt.id 
	LEFT JOIN (SELECT tele_station_id,MIN(rainfall_datetime) FROM public.rainfall_monthly WHERE deleted_at=to_timestamp(0) AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') GROUP BY tele_station_id ) rmin ON rmin.tele_station_id=mt.id 
	WHERE ge.deleted_at=to_timestamp(0) AND province_code != '99' GROUP BY province_code,province_name::jsonb,mt.tele_station_oldcode,tele_station_name::jsonb,mt.id, rmax.max,rmin.min ORDER BY province_name::jsonb
	`
