package water_map

var (
	// getDamDaily = "SELECT d.dam_name::jsonb,a.agency_name::jsonb,b.basin_name::jsonb,sb.subbasin_name::jsonb,lg.area_name::jsonb,lg.province_name::jsonb,lg.amphoe_name::jsonb,lg.tumbon_name::jsonb, " +
	// 	"d.dam_lat,d.dam_long,d.dam_oldcode,d.max_water_level,d.normal_water_level,d.min_water_level,dd.dam_date,dd.dam_inflow,dd.dam_released,dd.dam_storage FROM public.dam_daily dd " +
	// 	"RIGHT JOIN (SELECT dam_id,MAX(dam_date) AS max_date FROM public.dam_daily GROUP BY dam_id ORDER BY dam_id) dm ON dd.dam_id=dm.dam_id AND dd.dam_date=dm.max_date " +
	// 	"LEFT JOIN public.m_dam d ON dd.dam_id=d.id LEFT JOIN public.agency a ON d.agency_id=a.id LEFT JOIN public.lt_geocode lg ON d.geocode_id=lg.id LEFT JOIN public.subbasin sb ON d.subbasin_id=sb.id " +
	// 	"LEFT JOIN public.basin b ON sb.basin_id=b.id " +
	// 	"GROUP BY d.dam_name::jsonb,a.agency_name::jsonb,b.basin_name::jsonb,sb.subbasin_name::jsonb,lg.area_name::jsonb,lg.province_name::jsonb,lg.amphoe_name::jsonb, " +
	// 	"lg.tumbon_name::jsonb,d.dam_lat,d.dam_long,d.dam_oldcode,d.max_water_level,d.normal_water_level,d.min_water_level,dd.dam_date,dd.dam_inflow,dd.dam_released,dd.dam_storage " +
	// 	"ORDER BY b.basin_name::jsonb,d.dam_name::jsonb"
	getDamDaily = `
	SELECT
		d.dam_name :: jsonb,
		a.agency_name :: jsonb,
		b.basin_name :: jsonb,
		sb.subbasin_name :: jsonb,
		lg.area_name :: jsonb,
		lg.province_name :: jsonb,
		lg.amphoe_name :: jsonb,
		lg.tumbon_name :: jsonb,
		d.dam_lat,
		d.dam_long,
		d.dam_oldcode,
		d.max_water_level,
		d.normal_water_level,
		d.min_water_level,
		dd.dam_date,
		dd.dam_inflow,
		dd.dam_released,
		dd.dam_storage
	FROM
		latest.dam_daily dd
		LEFT JOIN public.m_dam d ON dd.dam_id = d.id
		LEFT JOIN public.agency a ON d.agency_id = a.id
		LEFT JOIN public.lt_geocode lg ON d.geocode_id = lg.id
		LEFT JOIN public.subbasin sb ON d.subbasin_id = sb.id
		LEFT JOIN public.basin b ON sb.basin_id = b.id
	ORDER BY
		b.basin_name :: jsonb,
		d.dam_name :: jsonb
	`

	// getTeleWaterLevel = "SELECT mts.tele_station_name::jsonb, a.agency_name::jsonb,b.basin_name::jsonb,sb.subbasin_name::jsonb,lg.area_name::jsonb,lg.province_name::jsonb,lg.amphoe_name::jsonb, " +
	// 	"lg.tumbon_name::jsonb, mts.tele_station_lat, mts.tele_station_long, mts.tele_station_oldcode, mts.right_bank, mts.left_bank, tw.waterlevel_datetime , tw.waterlevel_m, tw.waterlevel_msl, tw.flow_rate, tw.discharge " +
	// 	"FROM tele_waterlevel tw RIGHT JOIN (SELECT tele_station_id,MAX(waterlevel_datetime) AS max_date FROM public.tele_waterlevel GROUP BY tele_station_id ORDER BY tele_station_id) twm " +
	// 	"ON tw.tele_station_id=twm.tele_station_id AND tw.waterlevel_datetime=twm.max_date LEFT JOIN public.m_tele_station mts ON tw.tele_station_id=mts.id " +
	// 	"LEFT JOIN public.agency a ON mts.agency_id=a.id LEFT JOIN public.lt_geocode lg ON mts.geocode_id=lg.id LEFT JOIN public.subbasin sb ON mts.subbasin_id=sb.id " +
	// 	"LEFT JOIN public.basin b ON sb.basin_id=b.id " +
	// 	"GROUP BY mts.tele_station_name::jsonb, a.agency_name::jsonb,b.basin_name::jsonb,sb.subbasin_name::jsonb,lg.area_name::jsonb,lg.province_name::jsonb,lg.amphoe_name::jsonb,lg.tumbon_name::jsonb, mts.tele_station_lat, mts.tele_station_long, mts.tele_station_oldcode, mts.right_bank, mts.left_bank, tw.waterlevel_datetime , tw.waterlevel_m, tw.waterlevel_msl, tw.flow_rate, tw.discharge  " +
	// 	"ORDER BY b.basin_name::jsonb,mts.tele_station_name::jsonb"
	getTeleWaterLevel = `
	SELECT
		mts.tele_station_name :: jsonb,
		a.agency_name :: jsonb,
		b.basin_name :: jsonb,
		sb.subbasin_name :: jsonb,
		lg.area_name :: jsonb,
		lg.province_name :: jsonb,
		lg.amphoe_name :: jsonb,
		lg.tumbon_name :: jsonb,
		mts.tele_station_lat,
		mts.tele_station_long,
		mts.tele_station_oldcode,
		mts.right_bank,
		mts.left_bank,
		tw.waterlevel_datetime,
		tw.waterlevel_m,
		tw.waterlevel_msl,
		tw.flow_rate,
		tw.discharge
	FROM
		latest.tele_waterlevel tw
		INNER JOIN public.m_tele_station mts ON tw.tele_station_id = mts.id
		LEFT JOIN public.agency a ON mts.agency_id = a.id
		LEFT JOIN public.lt_geocode lg ON mts.geocode_id = lg.id
		LEFT JOIN public.subbasin sb ON mts.subbasin_id = sb.id
		LEFT JOIN public.basin b ON sb.basin_id = b.id
	ORDER BY
		b.basin_name :: jsonb,
		mts.tele_station_name :: jsonb
	`
)
