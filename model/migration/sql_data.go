package migrate

import ()

var (
	NhcCanal  = "SELECT canal_station_oldid	,wl_canal_date+wl_canal_time,wl_canal,comm_status FROM wl_canal wt LEFT JOIN canal_station te ON wt.canal_station_id = te.canal_station_id WHERE (anhc_id = '15009') and (wl_canal_date=>$1 and wl_canal_date<=$2) order by canal_station_oldid,wl_canal_date+wl_canal_time"
	TW30Canal = "SELECT canal_station_oldcode,cw.id,cw.canal_station_id,canal_waterlevel_datetime,canal_waterlevel_value,comm_status from canal_waterlevel cw LEFT JOIN m_canal_station mts ON mts.id=cw.canal_station_id where cw.deleted_at=to_timestamp(0) and (canal_waterlevel_datetime>=$1 and canal_waterlevel_datetime<=$2) order by canal_station_oldcode,canal_waterlevel_datetime"

	NHCDamDaily  = "SELECT dam_tname,dam_date,dam_level,dam_storage,dam_inflow,dam_released,dam_spilled,dam_losses,dam_evalp,dam_uses_water,dam_storage_percent,dam_uses_water_percent,dam_inflow_acc_percent FROM dam_daily LEFT JOIN dam ON dam.dam_id = dam_daily.dam_id WHERE (anhc_id = '50504' or (anhc_id = '07003' AND dam.dam_id LIKE 'ldam%')) and (dam_date>=$1 and dam_date<=$2) order by dam_tname"
	TW30DamDaily = "SELECT md.dam_name->>'th',dd.dam_date,dam_level,dam_storage,dam_inflow,dam_released,dam_spilled,dam_losses,dam_evap,dam_uses_water,dam_storage_percent,dam_uses_water_percent,dam_inflow_avg,dam_released_acc,dam_inflow_acc,dam_inflow_acc_percent from dam_daily dd LEFT JOIN public.m_dam md ON md.id=dd.dam_id where dd.deleted_at=to_timestamp(0) and (dam_date>=$1 and dam_date<=$2) order by md.dam_name->>'th'"

	NHCDamHourly  = "SELECT dam_tname,dam_date+dam_time,dam_level,dam_storage,dam_inflow,dam_released,dam_spilled,dam_losses,dam_evalp FROM dam_hourly LEFT JOIN dam ON dam.dam_id = dam_hourly.dam_id WHERE anhc_id = '50504' and (dam_date+dam_time>=$1 and dam_date+dam_time<=$2)  order by dam_tname"
	TW30DamHourly = "SELECT md.dam_name->>'th',dam_datetime,dam_level,dam_storage,dam_inflow,dam_released,dam_spilled,dam_losses,dam_evap from dam_hourly dd LEFT JOIN public.m_dam md ON md.id=dd.dam_id where dd.deleted_at=to_timestamp(0) and (dam_datetime>=$1 and dam_datetime<=$2) order by md.dam_name->>'th'"

	NHCFord  = "SELECT te.ford_station_oldid,wl_ford_date+wl_ford_time,wl_ford,comm_status FROM wl_ford wt LEFT JOIN ford_station te ON wt.ford_station_id = te.ford_station_id WHERE anhc_id = '15009' and (wl_ford_date+wl_ford_time>=$1 and wl_ford_date+wl_ford_time<=$2 order by te.ford_station_oldid"
	TW30Ford = "SELECT ford_station_oldcode,ford_waterlevel_datetime,ford_waterlevel_value,comm_status from ford_waterlevel fw LEFT JOIN m_ford_station mts ON mts.id=fw.ford_station_id where fw.deleted_at=to_timestamp(0) and (ford_waterlevel_datetime>=$1 and ford_waterlevel_datetime<=$2)"

	NHCHumid  = "SELECT tele_station_oldid,humid_date + humid_time,humid_value FROM humid wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012' or  anhc_id = '02005' ) and (humid_date + humid_time>=$1 and humid_date + humid_time<=$2) order by tele_station_oldid,humid_date + humid_time"
	TW30Humid = "SELECT mts.tele_station_oldcode, humid_datetime, humid_value FROM public.humid h LEFT JOIN public.m_tele_station mts ON mts.id=h.tele_station_id WHERE h.deleted_at=to_timestamp(0) and (humid_datetime>=$1 and humid_datetime<=$2) order by tele_station_oldcode,humid_datetime"

	NHCMediumDam  = "SELECT dam_tname,dam_date,dam_level,dam_storage,dam_inflow,dam_released,dam_spilled,dam_losses,dam_evalp,dam_uses_water,dam_storage_percent,dam_uses_water_percent,dam_inflow_acc_percent FROM dam_daily LEFT JOIN dam ON dam.dam_id = dam_daily.dam_id WHERE (anhc_id = '07003' AND dam.dam_id like 'mdamrid%') and (dam_date>=$1 and dam_date<=$2) order by dam_tname"
	TW30MediumDam = "SELECT mediumdam_name->>'th', mediumdam_date, mediumdam_storage, mediumdam_inflow, mediumdam_released, mediumdam_uses_water, mediumdam_storage_percent FROM public.medium_dam md LEFT JOIN m_medium_dam mmd ON mmd.id=md.mediumdam_id WHERE md.deleted_at=to_timestamp(0) and mediumdam_date>=$1 and mediumdam_date<=$2 order by mediumdam_name->>'th', mediumdam_date"

	NHCPressure  = "SELECT tele_station_oldid,pressure_date + pressure_time,pressure_value FROM pressure wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE anhc_id = '19012' and  (pressure_date + pressure_time>=$1 pressure_date + pressure_time<=$2) ORDER BY tele_station_oldid,pressure_date + pressure_time"
	TW30Pressure = "SELECT tele_station_oldcode, pressure_datetime, pressure_value FROM public.pressure p LEFT JOIN public.m_tele_station mts ON mts.id=p.tele_station_id WHERE p.deleted_at=to_timestamp(0) and (pressure_datetime>=$1 and pressure_datetime<=$2) order by tele_station_oldcode, pressure_datetime"

	NHCRainfall  = "SELECT tele_station_oldid,rainfall_date + rainfall_time as rainfall_date, rainfall_date_calc :: timestamp with time zone ,rainfall5m, rainfall10m, rainfall15m, rainfall30m, rainfall1h, rainfall3h, rainfall6h, rainfall12h, rainfall24h, rainfall_acc FROM rainfall wt left join tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '09006' or anhc_id = '02005' or (anhc_id = '07003' AND tele_station_oldid like 'rid%') or anhc_id = '11004' or anhc_id = '50504') and (rainfall_date + rainfall_time>=$1 and rainfall_date + rainfall_time<=$2) order by tele_station_oldid ,rainfall_date + rainfall_time"
	TW30Rainfall = "SELECT tele_station_oldcode, rainfall_datetime, rainfall5m, rainfall_date_calc, rainfall10m, rainfall15m, rainfall30m, rainfall1h, rainfall3h, rainfall6h, rainfall12h, rainfall24h, rainfall_acc, rainfall_today FROM public.rainfall r LEFT JOIN public.m_tele_station mts ON mts.id=r.tele_station_id WHERE r.deleted_at=to_timestamp(0) and (rainfall_datetime>=$1 and rainfall_datetime<=$2) order by tele_station_oldcode,rainfall_datetime"

	NHCRainfall1H  = "SELECT tele_station_oldid, (rainfall_date + rainfall_time) as rainfall_date,rainfall_date_calc,rainfall1h FROM rainfall1h_test wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id  where (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '50504' ) and (rainfall_date=$1 and rainfall_date<=$2) order by tele_station_oldid, rainfall_date + rainfall_time"
	TW30Rainfall1H = "SELECT tele_station_oldcode, rainfall_datetime, rainfall_datetime_calc,rainfall1h FROM public.rainfall_1h r LEFT JOIN public.m_tele_station mts ON mts.id=r.tele_station_id WHERE r.deleted_at=to_timestamp(0) and (rainfall_datetime>=$1 and rainfall_datetime<=$2) order by tele_station_oldcode, rainfall_datetime;"

	NHCRainfall24H  = "SELECT tele_station_oldid, rainfall_date + rainfall_time as rainfall_date,rainfall_date_calc ,rainfall24h FROM rainfall24h wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id where  (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '09006' or anhc_id = '11004' or (anhc_id = '07003' AND tele_station_oldid like 'rid%')  or anhc_id = '50504' ) and (rainfall_date + rainfall_time>=$1 and rainfall_date + rainfall_time<=$2)"
	TW30Rainfall24H = "SELECT tele_station_oldcode, rainfall_datetime, rainfall_datetime_calc,rainfall24h FROM public.rainfall_24h r LEFT JOIN public.m_tele_station mts ON mts.id=r.tele_station_id WHERE r.deleted_at=to_timestamp(0) and (rainfall_datetime>=$1 and rainfall_datetime<=$2) order by tele_station_oldcode, rainfall_datetime"

	NHCRainfallDaily  = "SELECT tele_station_oldid, rainfall_date :: timestamp with time zone as rainfall_date,rainfall_daily FROM rainfall_daily wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id where (anhc_id ='11004' or anhc_id = '09006' or anhc_id = '50504' or anhc_id = '19012' or anhc_id = '15009') and (rainfall_date>=$1 and rainfall_date<=$2) order by  tele_station_oldid, rainfall_date"
	TW30RainfallDaily = "SELECT tele_station_oldcode, rainfall_datetime, rainfall_value  FROM public.rainfall_daily r LEFT JOIN public.m_tele_station mts ON mts.id=r.tele_station_id WHERE r.deleted_at=to_timestamp(0) and (rainfall_datetime>=$1 and rainfall_datetime<=$2) order by tele_station_oldcode,rainfall_datetime"

	NHCSoilMoisutre  = "SELECT tele_station_oldid,(soil_date + soil_time) :: timestamp with time zone as soil_date ,soil_value FROM soil_moisture wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE anhc_id = '09006' and (soil_date + soil_time>=$1 and soil_date + soil_time<=$2) order by tele_station_oldid,soil_date + soil_time"
	TW30SoilMoisutre = "SELECT tele_station_oldcode, soil_datetime, soil_value FROM public.soilmoisture s LEFT JOIN public.m_tele_station mts ON mts.id=s.tele_station_id WHERE s.deleted_at=to_timestamp(0) and (soil_datetime>=$1 and soil_datetime<=$2) order by tele_station_oldcode, soil_datetime"

	NHCSolar  = "SELECT tele_station_oldid,(solar_date + solar_time) :: timestamp with time zone as solar_date ,solar_value FROM solar wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id  WHERE anhc_id = '19012' and (solar_date + solar_time>=$1 and solar_date + solar_time<=$2) order by tele_station_oldid,(solar_date + solar_time)"
	TW30Solar = "SELECT tele_station_oldcode, solar_datetime,solar_value FROM public.solar s LEFT JOIN public.m_tele_station mts ON mts.id=s.tele_station_id WHERE s.deleted_at=to_timestamp(0) and (solar_datetime>=$1 and solar_datetime<=$2)order by tele_station_oldcode, solar_datetime"

	NHCTeleWaterGate  = "SELECT tele_station_oldid,(wl_tele_date + wl_tele_time) :: timestamp with time zone as wl_tele_date,wl_in,wl_out,wl_out2 FROM wl_tele wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '15009')  and (wl_tele_date + wl_tele_time>=$1 and wl_tele_date + wl_tele_time<=$2) order by tele_station_oldid,(wl_tele_date + wl_tele_time)"
	TW30TeleWaterGate = "SELECT tele_station_oldcode, watergate_datetime, watergate_in, watergate_out, watergate_out2 FROM public.tele_watergate r LEFT JOIN public.m_tele_station mts ON mts.id=r.tele_station_id WHERE r.deleted_at=to_timestamp(0) and (watergate_datetime>=$1 and watergate_datetime<=$2) order by tele_station_oldcode, watergate_datetime"

	NHCTeleWaterlevel  = "SELECT tele_station_oldid,(wl_tele_date + wl_tele_time) :: timestamp with time zone as wl_tele_date ,wl_m,wl_msl,flow_rate,discharge FROM wl_tele wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id  WHERE (anhc_id = '19012' or  anhc_id = '08003' or anhc_id = '09006' or anhc_id = '02005' or (anhc_id = '07003' AND tele_station_oldid like 'rdw%') or (anhc_id = '07003' AND tele_station_oldid like 'rid%') or  anhc_id = '50504') and (wl_tele_date + wl_tele_time>=$1 and wl_tele_date + wl_tele_time<=$) order by  tele_station_oldid,(wl_tele_date + wl_tele_time)"
	TW30TeleWaterlevel = "SELECT tele_station_oldcode, waterlevel_datetime, waterlevel_m, waterlevel_msl, flow_rate, discharge FROM public.tele_waterlevel r LEFT JOIN public.m_tele_station mts ON mts.id=r.tele_station_id WHERE r.deleted_at=to_timestamp(0) and waterlevel_datetime>=$1 and waterlevel_datetime=$2) order by tele_station_oldcode, waterlevel_datetime"

	NHCTemp  = "SELECT tele_station_oldid,temp_date+temp_time,temp_value FROM temp wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012' or anhc_id = '02005' or anhc_id = '50504') and (temp_date+temp_time>=$1 and temp_date+temp_time<=$2) and tele_station_oldid='BKK001'  order by tele_station_oldid "
	TW30Temp = "SELECT mts.tele_station_oldcode,t.temp_datetime,temp_value FROM public.temperature t LEFT JOIN public.m_tele_station mts ON mts.id=t.tele_station_id WHERE t.deleted_at=to_timestamp(0) and (temp_datetime>=$1 and temp_datetime<=$2) and tele_station_oldcode='BKK001' order by tele_station_oldcode,temp_datetime"

	NHCWind  = "SELECT tele_station_oldid,(wind_date + wind_time) :: timestamp with time zone as wind_date ,wind_speed,wind_dir,wind_dir_value FROM wind wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012' or anhc_id = '02005') and (wind_date + wind_time>=$1 and wind_date + wind_time<=$2)"
	TW30Wind = "SELECT mts.tele_station_oldcode,t.wind_datetime,wind_speed,wind_dir,wind_dir_value FROM public.wind t LEFT JOIN public.m_tele_station mts ON mts.id=t.tele_station_id WHERE t.deleted_at=to_timestamp(0) and (wind_datetime>=$1 and wind_datetime<=$2) order by tele_station_oldcode,wind_datetime"

	NHCMDam = "SELECT dam_tname::TEXT,dam_ename::TEXT,dam_lat::double precision,dam_long::double precision,max_water_level,normal_water_level,min_water_level,normal_watergate_level,emer_watergate_level,service_watergate_level,max_old_storage,maxos_date,min_old_storage,minos_date,top_spillway_level,ridge_spillway_level,max_storage,normal_storage,min_storage,uses_water,avg_inflow,avg_inflow_intyear::TEXT,avg_inflow_endyear::TEXT,max_inflow,max_inflow_date,downstream_storage,water_shed,rainfall_yearly,power_install,power_intake_storage,power_intake_level,tailrace_level,used_genpower FROM dam WHERE anhc_id = '50504'  or (anhc_id = '07003' AND dam.dam_id LIKE 'ldam%') order by dam_tname"
	TW30Dam = "SELECT dam_name->>'th' as name_th,dam_name->>'en' as name_en,dam_lat::double precision,dam_long::double precision,max_water_level,normal_water_level,min_water_level,normal_watergate_level,emer_watergate_level,service_watergate_level,max_old_storage,maxos_date,min_old_storage,minos_date,top_spillway_level,ridge_spillway_level,max_storage,normal_storage,min_storage,uses_water,avg_inflow,avg_inflow_intyear::TEXT,avg_inflow_endyear::TEXT,max_inflow,max_inflow_date,downstream_storage,water_shed,rainfall_yearly,power_install,power_intake_storage,power_intake_level,tailrace_level,used_genpower FROM public.m_dam WHERE deleted_at=to_timestamp(0) order by dam_name->>'th'"

	NHCCanal   = "SELECT canal_station_oldid::TEXT,canal_station_name::TEXT,canal_station_long::double precision, canal_station_lat::double precision,district_id::TEXT,amphoe_id::TEXT,province_id::TEXT,mbasin_id::TEXT,sbasin_id::TEXT FROM canal_station order by canal_station_oldid"
	TW30MCanal = "SELECT canal_station_oldcode,canal_station_name->>'th' as canal_station_name_th,canal_station_long::double precision, canal_station_lat::double precision,subbasin_id,geocode_id,agency_id from m_canal_station where deleted_at=to_timestamp(0) order by canal_station_oldcode "

	NHCMFord  = "SELECT ford_station_name::TEXT,ford_station_lat::double precision,ford_station_long::double precision,ford_station_oldid::TEXT FROM ford_station WHERE anhc_id = '15009' order by ford_station_name"
	TW30MFord = "SELECT ford_station_name->>'th' as for_station_name_th,ford_station_lat::double precision,ford_station_long::double precision,ford_station_oldcode FROM public.m_ford_station WHERE deleted_at=to_timestamp(0) order by ford_station_name->>'th'"

	NHCTele  = "SELECT tele_station_name::TEXT,tele_station_oldid::TEXT,tele_station_lat::double precision,tele_station_long::double precision,anhc_id::TEXT,left_bank,right_bank,ground_level FROM tele_station order by tele_station_oldid"
	TW30Tele = "SELECT tele_station_name->>'th' as tele_station_name_th,tele_station_oldcode,tele_station_lat::double precision,tele_station_long::double precision,agency_id,left_bank,right_bank,ground_level FROM public.m_tele_station WHERE deleted_at=to_timestamp(0) order by tele_station_oldcode"

	NHCMedium  = "SELECT dam_tname::TEXT,dam_ename::TEXT,dam_lat::double precision,dam_long::double precision,normal_storage,min_storage,max_storage FROM dam WHERE anhc_id = '07003' AND dam.dam_id like 'mdamrid%' order by dam_tname"
	TW30Medium = "SELECT mediumdam_name->>'th' as name_th,mediumdam_name->>'en' as name_en,mediumdam_oldcode,mediumdam_lat::double precision, mediumdam_long::double precision,normal_storage,min_storage,max_storage FROM m_medium_dam WHERE deleted_at=to_timestamp(0) order by mediumdam_name->>'th'"

	NHCAgency  = "SELECT id,agency_id::TEXT,agency_name from agency_document order by agency_id"
	TW30Agency = "SELECT id,agency_shortname->>'en' as agency_shortname_en ,agency_name->>'th' as agency_name FROM public.agency order by agency_shortname->>'en'"

	NHCBasin  = "SELECT mbasin_id::TEXT,basin_tname,basin_ename FROM basin order by basin_tname"
	TW30Basin = "SELECT id,basin_code,basin_name->>'th' as basin_name_th,basin_name->>'en' as basin_name_en FROM public.basin order by basin_name->>'th'"

	NHCSubbasin  = "SELECT sbasin_id::TEXT,mbasin_id::TEXT,sbasin_tname,sbasin_ename,sbasin_area,sbasin_area_km,sbasin_perimeter,sbasin_acres FROM subbasin order by sbasin_tname"
	TW30Subbasin = "SELECT id, basin_id,subbasin_code, subbasin_name->>'th' as subbasin_name, subbasin_area, subbasin_areakm, subbasin_perimeter, subbasin_acres FROM public.subbasin order by subbasin_name->>'th'"

	NHCGeocode  = "SELECT geocode,area_code::TEXT,province_code::TEXT,amphoe_code::TEXT,tumbon_code::TEXT,area_name,province_name,amphoe_name,tumbon_name FROM public.ref_geocode order by geocode"
	TW30Geocode = "SELECT id, geocode, area_code, province_code, amphoe_code, tumbon_code, area_name->>'th' as area_name, province_name->>'th' as province_name, amphoe_name->>'th' as amphoe_name, tumbon_name->>'th' as tumbon_name, warning_zone, zone_detail as zone_detail, tmd_area_code, tmd_area_name as tmd_area_name FROM public.lt_geocode order by geocode"

	NHCMediaType  = "SELECT media_type_id, media_type_name, media_subtype_name FROM public.media_type order by media_type_id"
	TW30MediaType = "SELECT id, media_type_name, media_subtype_name, media_category FROM public.lt_media_type order by id"

	NHCPowerPlant  = "SELECT power_plant_id::TEXT,power_plant_oldid,power_plant_name,power_plant_loc,power_plant_lat::double precision,power_plant_long::double precision,power_plant_type,power_producer_status,capacity_mw,sold_mw,fuel,secon_fuel FROM public.power_plant order by power_plant_oldid"
	TW30PowerPlant = "SELECT id, power_plant_oldcode, power_plant_name, power_plant_location, power_plant_lat, power_plant_long, power_plant_type, power_producer_status, capacity_mw, sold_mw, fuel, secon_fuel FROM public.power_plant order by power_plant_oldcode"

	NHCEgatMetadata  = `SELECT "Dam_ID","Station",name_th,name_en,"Tambon","Amphoe","Province","Lat","Long","North_UTM","East_UTM","BedLevel","LeftBank","RightBank",tele_station_oldid FROM public.egat_new_metada order by tele_station_oldid`
	TW30EgatMetadata = "SELECT tele_station_name->>'th' as tele_station_name,egat_namenv, tele_station_lat::double precision,tele_station_long::double precision,tele_station_oldcode,left_bank,right_bank,ground_level FROM public.egat_metadata em LEFT JOIN public.m_tele_station mts ON em.tele_station_id=mts.id order by tele_station_oldcode"

	NHCMSzone = "SELECT sfz_id::TEXT,district_id::TEXT,amphoe_id::TEXT,province_id::TEXT,sfz_name,sfz_oldid,mooban_name,mooban_id,sfz_lat::double precision,sfz_long::double precision FROM safety_zone order by sfz_name"
	TW30Szone = "SELECT id, geocode_id, safety_zone_oldcode, mooban_name, mooban_id, safety_zone_lat, safety_zone_long, safety_zone_name->>'th' as safety_zone_name FROM public.safety_zone order by safety_zone_name->>'th'"

	NHCXsection  = "SELECT xs_no::TEXT,survey_year,xs_location,xs_oldid,xs_filepath FROM xsection_station order by xs_oldid"
	TW30Xsection = "SELECT id, agency_id, survey_year, section_location, section_oldcode, section_filepath FROM public.m_crosssection_station order by section_oldcode"

	NHCRuleCurve  = "SELECT rt.dam_name,jan_value,feb_value,mar_value,apr_value,may_value,jun_value,jul_value,aug_value,	sep_value,oct_value,nov_value,dec_value FROM public.rule_curve r LEFT JOIN rule_curve_type rt ON rt.rc_type_id=r.rc_type_id order by rt.dam_name"
	TW30RuleCurve = "SELECT dam_name,rc_datetime,rc_unit,urc_old,lrc_old,urc_new,lrc_new FROM public.rulecurve order by dam_name,rc_datetime,rc_unit"

	NHCMGround  = "SELECT district_id::TEXT,amphoe_id::TEXT,province_id::TEXT,gwr_oldid,well_ownner,mooban_id::TEXT,mooban,gwr_lat::double precision,gwr_long::double precision,gwr_location,map_sheet FROM ground_wr order by gwr_oldid asc limit 2000 "
	TW30MGround = "SELECT geocode_id,ground_oldcode,well_ownner,mooban_id,mooban,ground_lat,ground_long,ground_location,map_sheet FROM public.m_ground order by ground_oldcode asc limit 2000"

	NHCGroundWater  = "SELECT gwr_oldid,gwr_size,gwr_depth,gwr_wl,aquifer FROM wl_groundwater wl LEFT JOIN ground_wr g ON g.gwr_id=wl.gwr_id order by gwr_oldid asc limit 2000"
	TW30GroundWater = "SELECT mg.ground_oldcode,ground_size,ground_depth,ground_waterlevel,ground_aquifer FROM public.ground_waterlevel wl LEFT JOIN m_ground mg ON mg.id=wl.ground_id order by ground_oldcode asc limit 2000"

	NHCXSectionData  = "SELECT xs_oldid,distance,water_level_msl,water_level_m,marker FROM xsection_data d LEFT JOIN xsection_station s ON s.xs_no=d.xs_no order by xs_oldid asc"
	TW30XSectionData = "SELECT s.section_oldcode,point_id,distance,water_level_msl,water_level_m,remark FROM public.crosssection d LEFT JOIN m_crosssection_station s ON s.id=d.section_station_id where d.deleted_at=to_timestamp(0) order by section_oldcode asc"

	NHCWaterResource  = "SELECT replace(wrs_oldid , 'lddswr','') as wrs_oldid ,projectname,projecttype,fiscal_year::TEXT,mooban,coordination,benefit_household,benefit_area,capacity::TEXT,standard_cost::TEXT,budget::TEXT,contract_signdate,contract_enddate,rec_date FROM water_resource order by wrs_oldid asc"
	TW30WaterResource = "SELECT water_resource_oldcode, projectname, projecttype, fiscal_year, mooban, coordination, benefit_household, benefit_area,capacity, standard_cost, budget, contract_signdate, contract_enddate, rec_date FROM public.water_resource where deleted_at=to_timestamp(0)"

	NHCTMDRegion  = "SELECT region_name::TEXT FROM lt_region_tmd"
	TW30TMDRegion = "select tmd_area_name::json->>'th' as tmd_area_name from public.lt_geocode group by tmd_area_name"

	//	SumNHC = `SELECT * FROM (SELECT 'wl_canal' as table, count(*) FROM wl_canal wt LEFT JOIN canal_station te ON wt.canal_station_id = te.canal_station_id WHERE (anhc_id = '15009') and wl_canal_date<='2017-07-10'
	//union
	//SELECT 'dam_daily' as table, count(*) FROM dam_daily LEFT JOIN dam ON dam.dam_id = dam_daily.dam_id WHERE (anhc_id = '50504' or (anhc_id = '07003' AND dam.dam_id LIKE 'ldam%')) and dam_date<='2017-07-10'
	//union
	//SELECT 'dam_hourly' as table, count(*) FROM dam_hourly LEFT JOIN dam ON dam.dam_id = dam_hourly.dam_id WHERE (anhc_id = '50504') and dam_date<='2017-07-10'
	//union
	//SELECT 'medium_dam' as table,count(*) FROM dam_daily LEFT JOIN dam ON dam.dam_id = dam_daily.dam_id WHERE (anhc_id = '07003' AND dam.dam_id like 'mdamrid%') and dam_date<='2017-07-10'
	//union
	//SELECT 'wl_ford' as table, count(*) FROM wl_ford wt LEFT JOIN ford_station te ON wt.ford_station_id = te.ford_station_id WHERE (anhc_id = '15009' ) and wl_ford_date<='2017-07-10'
	//union
	//SELECT 'pressure' as table, count(*) FROM pressure wt LEFT JOIN tele_station te
	//on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012' or anhc_id = '02005') and pressure_date<='2017-07-10'
	//union
	//SELECT 'rainfall' as table, count(*) FROM rainfall wt left join tele_station te
	//on wt.tele_station_id = te.tele_station_id
	//WHERE (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '09006' or anhc_id = '02005' or (anhc_id = '07003' AND tele_station_oldid like 'rid%')
	//or anhc_id = '11004' or anhc_id = '50504') and rainfall_date<='2017-07-10'
	//union
	//SELECT 'rainfall1h' as table, count(*) FROM rainfall1h_test wt
	//LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id
	//where (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '50504' ) and rainfall_date<='2017-07-10'
	//union
	//SELECT 'rainfall24h' as table, count(*) FROM rainfall24h wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id
	//where  (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '09006' or
	//anhc_id = '11004' or (anhc_id = '07003' AND tele_station_oldid like 'rid%')
	//or anhc_id = '50504' ) and rainfall_date<='2017-07-10'
	//union
	//SELECT 'rainfall_daily' as table, count(*) FROM rainfall_daily wt LEFT JOIN tele_station te
	//ON wt.tele_station_id = te.tele_station_id where (anhc_id ='11004' or
	//anhc_id = '09006' or anhc_id = '50504' or anhc_id = '19012' or anhc_id = '15009') and rainfall_date<='2017-07-10'
	//union
	//SELECT 'soil_moisture' as table, count(*) FROM soil_moisture wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '09006') and soil_date<='2017-07-10'
	//union
	//SELECT 'solar' as table, count(*) FROM solar wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id  WHERE anhc_id = '19012' and solar_date<='2017-07-10'
	//union
	//SELECT 'tele_watergate' as table, count(*) FROM wl_tele wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '15009')  and wl_tele_date<='2017-07-10'
	//union
	//SELECT 'tele_waterlevel' as table, count(*) FROM wl_tele wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id
	//WHERE (anhc_id = '19012' or  anhc_id = '08003' or anhc_id = '09006' or anhc_id = '02005' or
	//(anhc_id = '07003' AND tele_station_oldid like 'rdw%') or (anhc_id = '07003' AND tele_station_oldid like 'rid%') or  anhc_id = '50504') and wl_tele_date<='2017-07-10'
	//union
	//SELECT 'temp' as table, count(*) FROM temp wt
	//LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id
	//WHERE (anhc_id = '19012' or anhc_id = '02005' or anhc_id = '50504') and temp_date<='2017-07-10'
	//union
	//SELECT 'wind' as table, count(*) FROM wind wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id
	//WHERE (anhc_id = '19012' or anhc_id = '02005') and wind_date<='2017-07-10'
	//union
	//SELECT 'humid' as table,count(*) FROM humid wt
	//LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012'
	//or  anhc_id = '02005' ) and humid_date<='2017-07-10'
	//) as sumtable ORDER BY 1`

	SumNHC = []string{
		"SELECT 'wl_canal' as table, count(*) FROM wl_canal wt LEFT JOIN canal_station te ON wt.canal_station_id = te.canal_station_id WHERE (anhc_id = '15009') ",
		"SELECT 'dam_daily' as table, count(*) FROM dam_daily LEFT JOIN dam ON dam.dam_id = dam_daily.dam_id WHERE (anhc_id = '50504' or (anhc_id = '07003' AND dam.dam_id LIKE 'ldam%')) ",
		"SELECT 'dam_hourly' as table, count(*) FROM dam_hourly LEFT JOIN dam ON dam.dam_id = dam_hourly.dam_id WHERE (anhc_id = '50504') ",
		"SELECT 'medium_dam' as table,count(*) FROM dam_daily LEFT JOIN dam ON dam.dam_id = dam_daily.dam_id WHERE (anhc_id = '07003' AND dam.dam_id like 'mdamrid%') ",
		"SELECT 'wl_ford' as table, count(*) FROM wl_ford wt LEFT JOIN ford_station te ON wt.ford_station_id = te.ford_station_id WHERE (anhc_id = '15009' ) ",
		"SELECT 'pressure' as table, count(*) FROM pressure wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '19012' or anhc_id = '02005') ",
		`SELECT 'rainfall' as table, count(*) FROM rainfall wt left join tele_station te on wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '09006' or anhc_id = '02005' or (anhc_id = '07003' AND tele_station_oldid like 'rid%') 
		 or anhc_id = '11004' or anhc_id = '50504') `,
		`SELECT 'rainfall1h' as table, count(*) FROM rainfall1h_test wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '50504' ) `,
		`SELECT 'rainfall24h' as table, count(*) FROM rainfall24h wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id
		 WHERE  (anhc_id = '19012' or anhc_id = '15009' or anhc_id = '09006' or 
		 anhc_id = '11004' or (anhc_id = '07003' AND tele_station_oldid like 'rid%')  or anhc_id = '50504' ) `,
		`SELECT 'rainfall_daily' as table, count(*) FROM rainfall_daily wt LEFT JOIN tele_station te ON wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id ='11004' or anhc_id = '09006' or anhc_id = '50504' or anhc_id = '19012' or anhc_id = '15009') `,
		"SELECT 'soil_moisture' as table, count(*) FROM soil_moisture wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '09006') ",
		"SELECT 'solar' as table, count(*) FROM solar wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id  WHERE anhc_id = '19012' ",
		"SELECT 'tele_watergate' as table, count(*) FROM wl_tele wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id WHERE (anhc_id = '15009') ",
		`SELECT 'tele_waterlevel' as table, count(*) FROM wl_tele wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id = '19012' or  anhc_id = '08003' or anhc_id = '09006' or anhc_id = '02005' or
		 (anhc_id = '07003' AND tele_station_oldid like 'rdw%') or (anhc_id = '07003' AND tele_station_oldid like 'rid%') or  anhc_id = '50504') `,
		`SELECT 'temp' as table, count(*) FROM temp wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id = '19012' or anhc_id = '02005' or anhc_id = '50504') `,
		`SELECT 'wind' as table, count(*) FROM wind wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id = '19012' or anhc_id = '02005') `,
		`SELECT 'humid' as table,count(*) FROM humid wt LEFT JOIN tele_station te on wt.tele_station_id = te.tele_station_id 
		 WHERE (anhc_id = '19012' or  anhc_id = '02005' ) `,
	}

	//	SumTW30 = `SELECT * FROM (select 'wl_canal' as table, count(id) from canal_waterlevel where deleted_at=to_timestamp(0) and
	//canal_waterlevel_datetime::date<='2017-07-10'
	//union
	//select 'dam_daily' as table, count(id) from dam_daily where deleted_at=to_timestamp(0) and
	//dam_date<='2017-07-10'
	//union
	//select 'dam_hourly' as table, count(id) from dam_hourly where deleted_at=to_timestamp(0) and
	//dam_datetime::date<='2017-07-10'
	//union
	//select 'medium_dam' as table, count(id) from medium_dam where deleted_at=to_timestamp(0) and
	//mediumdam_date<='2017-07-10'
	//union
	//select 'wl_ford' as table, count(id) from ford_waterlevel where deleted_at=to_timestamp(0) and
	//ford_waterlevel_datetime::date<='2017-07-10'
	//union
	//select 'pressure' as table, count(id) from pressure where deleted_at=to_timestamp(0) and
	//pressure_datetime::date<='2017-07-10'
	//union
	//select 'rainfall' as table, count(id) from rainfall where deleted_at=to_timestamp(0) and
	//rainfall_datetime::date<='2017-07-10'
	//union
	//select 'rainfall1h' as table, count(id) from rainfall_1h where deleted_at=to_timestamp(0) and
	//rainfall_datetime::date<='2017-07-10'
	//union
	//select 'rainfall24h' as table, count(id) from rainfall_24h where deleted_at=to_timestamp(0) and
	//rainfall_datetime::date<='2017-07-10'
	//union
	//select 'rainfall_daily' as table, count(id) from rainfall_daily where deleted_at=to_timestamp(0) and
	//rainfall_datetime::date<='2017-07-10'
	//union
	//select 'soil_moisture' as table, count(id) from soilmoisture where deleted_at=to_timestamp(0) and
	//soil_datetime::date<='2017-07-10'
	//union
	//select 'solar' as table, count(id) from solar where deleted_at=to_timestamp(0) and
	//solar_datetime::date<='2017-07-10'
	//union
	//select 'tele_watergate' as table, count(id) from tele_watergate where deleted_at=to_timestamp(0) and
	//watergate_datetime::date<='2017-07-10'
	//union
	//select 'tele_waterlevel' as table, count(id) from tele_waterlevel where deleted_at=to_timestamp(0) and
	//waterlevel_datetime::date<='2017-07-10'
	//union
	//select 'temp' as table, count(id) from temperature where deleted_at=to_timestamp(0) and
	//temp_datetime::date<='2017-07-10'
	//union
	//select 'wind' as table, count(id) from wind where deleted_at=to_timestamp(0) and
	//wind_datetime::date<='2017-07-10'
	//union
	//select 'humid' as table, count(id) from humid where deleted_at=to_timestamp(0) and
	//humid_datetime::date <= '2017-07-10'
	//) as alltable ORDER BY 1`

	SumTW30 = []string{
		"SELECT 'wl_canal' as table, count(id) FROM canal_waterlevel WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'dam_daily' as table, count(id) FROM dam_daily WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'dam_hourly' as table, count(id) FROM dam_hourly WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'medium_dam' as table, count(id) FROM medium_dam WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'wl_ford' as table, count(id) FROM ford_waterlevel WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'pressure' as table, count(id) FROM pressure WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'rainfall' as table, count(id) FROM rainfall WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'rainfall1h' as table, count(id) FROM rainfall_1h WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'rainfall24h' as table, count(id) FROM rainfall_24h WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'rainfall_daily' as table, count(id) FROM rainfall_daily WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'soil_moisture' as table, count(id) FROM soilmoisture WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'solar' as table, count(id) FROM solar WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'tele_watergate' as table, count(id) FROM tele_watergate WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'tele_waterlevel' as table, count(id) FROM tele_waterlevel WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'temp' as table, count(id) FROM temperature WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'wind' as table, count(id) FROM wind WHERE deleted_at=to_timestamp(0) ",
		"SELECT 'humid' as table, count(id) FROM humid WHERE deleted_at=to_timestamp(0) ",
	}

	SumMasterNHC = `SELECT * FROM (
SELECT 'canal_station' as table,count(*) FROM canal_station 	
union
SELECT 'dam' as table,count(*) FROM dam WHERE anhc_id = '50504'  or (anhc_id = '07003' AND dam.dam_id LIKE 'ldam%')
union
SELECT 'ford_station' as table,count(*) FROM ford_station WHERE anhc_id = '15009' 
union
SELECT 'tele_station' as table,count(*) FROM tele_station 
union
SELECT 'm_medium_dam' as table,count(*) FROM dam WHERE anhc_id = '07003' AND dam.dam_id like 'mdamrid%'
union
SELECT 'agency_document' as table,count(*) FROM agency_document
union
SELECT 'basin' as table,count(*) FROM basin
union
SELECT 'subbasin' as table,count(*) FROM subbasin
union
SELECT 'egat_new_metada' as table,count(*) FROM egat_new_metada
union
SELECT 'media_type' as table,count(*) FROM media_type
union
SELECT 'power_plant' as table,count(*) FROM power_plant
union
SELECT 'safety_zone' as table,count(*) FROM safety_zone
union
SELECT 'xsection_station' as table,count(*) FROM xsection_station
union
SELECT 'water_resource' as table, count(*) FROM water_resource
union
SELECT 'lt_region_tmd' as table, count(*) FROM lt_region_tmd
union
SELECT 'wl_groundwater' as table, count(*) FROM wl_groundwater
union
SELECT 'xsection_data' as table, count(*) FROM xsection_data
union
SELECT 'rule_curve' as table, count(*) FROM rule_curve
union
SELECT 'ground_wr' as table, count(*) FROM ground_wr
	) as sumtable ORDER BY 1`

	SumMasterTW30 = `SELECT * FROM (
SELECT 'canal_station' as table, count(id) from m_canal_station where deleted_at=to_timestamp(0)
union
SELECT 'dam' as table, count(id) FROM public.m_dam WHERE deleted_at=to_timestamp(0)
union
SELECT 'ford_station' as table, count(id) FROM public.m_ford_station WHERE deleted_at=to_timestamp(0)
union
SELECT 'tele_station' as table, count(id) FROM public.m_tele_station WHERE deleted_at=to_timestamp(0)
union
SELECT 'agency_document' as table, count(id) FROM agency WHERE deleted_at=to_timestamp(0)
union
SELECT 'basin' as table, count(id) FROM basin WHERE deleted_at=to_timestamp(0)
union
SELECT 'subbasin' as table, count(id) FROM subbasin WHERE deleted_at=to_timestamp(0)
union
SELECT 'egat_new_metada' as table, count(id) FROM egat_metadata WHERE deleted_at=to_timestamp(0)
union
SELECT 'media_type' as table, count(id) FROM lt_media_type WHERE deleted_at=to_timestamp(0)
union
SELECT 'power_plant' as table, count(id) FROM power_plant WHERE deleted_at=to_timestamp(0)
union
SELECT 'ref_geocode' as table, count(id) FROM lt_geocode WHERE deleted_at=to_timestamp(0)
union
SELECT 'safety_zone' as table, count(id) FROM safety_zone WHERE deleted_at=to_timestamp(0)
union
SELECT 'water_resource' as table, count(id) FROM water_resource WHERE deleted_at=to_timestamp(0)
union
SELECT 'xsection_station' as table, count(id) FROM m_crosssection_station WHERE deleted_at=to_timestamp(0)
union
SELECT 'm_medium_dam' as table, count(id) FROM m_medium_dam WHERE deleted_at=to_timestamp(0)
union
SELECT 'wl_groundwater' as table, count(id) FROM ground_waterlevel WHERE deleted_at=to_timestamp(0)
union
SELECT 'xsection_data' as table, count(id) FROM crosssection WHERE deleted_at=to_timestamp(0)
union
SELECT 'ground_wr' as table, count(id) FROM m_ground WHERE deleted_at=to_timestamp(0)
union
SELECT 'rule_curve' as table, count(id) FROM rulecurve WHERE deleted_at=to_timestamp(0)
union
SELECT 'lt_region_tmd' as table ,count(t.tmd_area_name) FROM (SELECT tmd_area_name FROM public.lt_geocode group by tmd_area_name) t 
	) as alltable ORDER BY 1`
)
