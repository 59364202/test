// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for public.metadata table. This table store metadata.
package metadata

const (
	MetadataStatus_Connect     = "1"
	MetadataStatus_WaitUpdate  = "3"
	MetadataStatus_WaitConnect = "2"
	MetadataStatus_Cancel      = "4"
)

func GetColumnDateTime(table string) string {
	return GetTablePartitionField(table)
}

func GetColumnMasterId(table string) string {
	return GetTableMasterId(table)
}

var filterValue = "999999"

//Fields = กรณีมี dataset มี mapping ต้องใส่ ว่าอนุญาติให้เค้า mapping column ไหนบ้าง
//เช่น  mapping table : m_tele_station โดยใช้ tele_station_oldcode ในการเช็ค แล้วเอา id ไปใช้
//ต้องใส่ Field : id, tele_station_oldcode
//HasProvince = ถ้ามี geocode_id ในตาราง
//HasBasin = ถ้ามี subbasin_id ในตาราง

var mapStrctTable = map[string]*Struct_Table{
	"agency": &Struct_Table{
		Table:    "agency",
		Fields:   "id, agency_name, agency_name->>'th', agency_name->>'en'",
		IsMaster: true},

	"air": &Struct_Table{
		Table:          "air",
		PartitionField: "air_datetime",
		MasterId:       "air_station_id",
		MasterTable:    "m_air_station",
		HasProvince:    true},

	"basin": &Struct_Table{
		Table:    "basin",
		Fields:   "id, agency_id, basin_code, basin_name, basin_name->>'th', basin_name->>'en'",
		IsMaster: true},

	"canal_waterlevel": &Struct_Table{
		Table:          "canal_waterlevel",
		PartitionField: "canal_waterlevel_datetime",
		MasterId:       "canal_station_id",
		MasterTable:    "m_canal_station",
		HasProvince:    true,
		HasBasin:       true},

	"cctv": &Struct_Table{
		Table:       "cctv",
		HasBasin:    true,
		HasProvince: true,
		IsMaster:    true},

	"crosssection": &Struct_Table{
		Table:       "crosssection",
		MasterId:    "section_station_id",
		MasterTable: "m_crosssection_station",
		HasProvince: false},

	"dam_daily": &Struct_Table{
		Table:          "dam_daily",
		PartitionField: "dam_date",
		Fields:         "id, dam_uses_water, dam_id, dam_inflow_avg,dam_storage_percent",
		MasterId:       "dam_id",
		MasterTable:    "m_dam",
		HasProvince:    true,
		HasBasin:       true},

	"dam_hourly": &Struct_Table{
		Table:          "dam_hourly",
		PartitionField: "dam_datetime",
		Fields:         "id, dam_id, dam_storage",
		MasterId:       "dam_id",
		MasterTable:    "m_dam",
		HasProvince:    true,
		HasBasin:       true},

	"mapping_near_dam": &Struct_Table{
		Table:    "mapping_near_dam",
		MasterId: "province_id"},

	"drought_area": &Struct_Table{
		Table:          "drought_area",
		PartitionField: "drought_datetime",
		HasProvince:    true},

	"egat_metadata": &Struct_Table{
		Table:       "egat_metadata",
		Fields:      "id, tele_station_id, egat_deviceid, egat_namenv",
		HasProvince: false},

	"flood_situation": &Struct_Table{
		Table:          "flood_situation",
		PartitionField: "flood_datetime",
		HasProvince:    false},

	"floodforecast": &Struct_Table{
		Table:          "floodforecast",
		PartitionField: "floodforecast_datetime",
		MasterId:       "floodforecast_station_id",
		MasterTable:    "m_floodforecast_station",
		HasProvince:    true,
		HasBasin:       true},

	"ford_waterlevel": &Struct_Table{
		Table:          "ford_waterlevel",
		PartitionField: "ford_waterlevel_datetime",
		MasterId:       "ford_station_id",
		MasterTable:    "m_ford_station",
		HasProvince:    true},

	"geohazard_situation": &Struct_Table{
		Table:          "geohazard_situation",
		Fields:         "id, geocode_id, agency_id, geohazard_name, geohazard_name ->>'th', geohazard_name ->>'en'",
		PartitionField: "geohazard_datetime",
		HasProvince:    true},

	"ground_waterlevel": &Struct_Table{
		Table:       "ground_waterlevel",
		MasterTable: "m_ground",
		MasterId:    "ground_id"},

	"ground_waterquality": &Struct_Table{
		Table:       "ground_waterquality",
		MasterTable: "m_ground",
		MasterId:    "ground_id"},

	"humid": &Struct_Table{
		Table:          "humid",
		PartitionField: "humid_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		Where:          "humid_value <> " + filterValue,
		HasProvince:    true,
		HasBasin:       true},

	"lt_geocode": &Struct_Table{
		Table:    "lt_geocode",
		Fields:   "id,geocode,province_code,amphoe_code,province_name, province_name->>'th', province_name->>'en', amphoe_name, amphoe_name->>'th', amphoe_name->>'en',tumbon_name, tumbon_name->>'th', tumbon_name->>'en'",
		IsMaster: true},

	"lt_subcategory": &Struct_Table{
		Table:    "lt_subcategory",
		Fields:   "id, subcategory_name, subcategory_name->>'th', subcategory_name->>'en'",
		IsMaster: true},

	"landslide_area": &Struct_Table{
		Table:       "landslide_area",
		IsMaster:    true,
		HasProvince: true,
		HasBasin:    true},

	"latest_media": &Struct_Table{
		Table:          "latest_media",
		PartitionField: "media_datetime"},

	"m_air_station": &Struct_Table{
		Table:        "m_air_station",
		SelectColumn: " air_station_name, air_station_lat, air_station_long ",
		Fields:       "id, air_station_oldcode, air_station_name, air_station_lat, air_station_long, air_staiton_type, agency_id",
		IsMaster:     true},

	"m_canal_station": &Struct_Table{
		Table:        "m_canal_station",
		SelectColumn: " canal_station_name, canal_station_lat, canal_station_long ",
		Fields:       "id, canal_station_oldcode",
		IsMaster:     true,
		HasProvince:  true,
		HasBasin:     true},

	"m_crosssection_station": &Struct_Table{
		Table:        "m_crosssection_station",
		Fields:       "id, section_oldcode",
		SelectColumn: " section_location ",
		IsMaster:     true,
	},

	"m_dam": &Struct_Table{
		Table:        "m_dam",
		SelectColumn: " dam_lat, dam_long ",
		Fields:       "id,dam_oldcode,normal_storage,dam_name, dam_name->>'th', dam_name->>'en',agency_id,geocode_id",
		IsMaster:     true,
		HasProvince:  true,
		HasBasin:     true},

	"m_floodforecast_station": &Struct_Table{
		Table:        "m_floodforecast_station",
		SelectColumn: " floodforecast_name, floodforecast_lat, floodforecast_long ",
		Fields:       "id, floodforecast_station_oldcode",
		IsMaster:     true,
		HasProvince:  true,
		HasBasin:     true},

	"m_ford_station": &Struct_Table{
		Table:        "m_ford_station",
		SelectColumn: " ford_station_name, ford_station_lat, ford_station_long ",
		Fields:       "id, geocode_id, agency_id, ford_station_name, ford_station_name ->>'th', ford_station_name ->>'en', ford_station_name ->>'jp'",
		IsMaster:     true,
		HasProvince:  true},

	"m_ground": &Struct_Table{
		Table:        "m_ground",
		SelectColumn: " ground_location, ground_lat, ground_long ",
		IsMaster:     true,
	},

	"m_medium_dam": &Struct_Table{
		Table:        "m_medium_dam",
		SelectColumn: " mediumdam_name, mediumdam_lat, mediumdam_long ",
		IsMaster:     true,
		HasProvince:  true},

	"m_small_dam": &Struct_Table{
		Table:       "m_small_dam",
		IsMaster:    true,
		HasProvince: true,
		HasBasin:    true},

	"m_sea_station": &Struct_Table{
		Table:        "m_sea_station",
		SelectColumn: " sea_station_name, sea_station_lat, sea_station_long",
		IsMaster:     true,
		HasProvince:  true,
		HasBasin:     true},

	"m_swan_station": &Struct_Table{
		Table:        "m_swan_station",
		SelectColumn: " swan_name, swan_lat, swan_long ",
		IsMaster:     true,
		HasProvince:  true},

	"m_tele_station": &Struct_Table{
		Table:        "m_tele_station",
		SelectColumn: " tele_station_name, tele_station_name->>'th', tele_station_name->>'en', tele_station_lat, tele_station_long ",
		Fields:       "id, tele_station_oldcode, agency_id, tele_station_type, tele_station_name, tele_station_lat, tele_station_long",
		WhereHydro:   "hydro_id IS NOT NULL ",
		IsMaster:     true,
		HasProvince:  true,
		HasBasin:     true},

	"m_waterquality_station": &Struct_Table{
		Table:        "m_waterquality_station",
		SelectColumn: " waterquality_station_name, waterquality_station_lat, waterquality_station_long ",
		Fields:       "id,waterquality_station_oldcode,waterquality_station_name, waterquality_station_name->>'th', waterquality_station_name->>'en'",
		IsMaster:     true,
		HasProvince:  true,
		HasBasin:     true},

	"media": &Struct_Table{
		Table:          "media",
		PartitionField: "media_datetime"},

	"media_animation": &Struct_Table{
		Table:          "media_animation",
		PartitionField: "media_datetime"},

	"media_other": &Struct_Table{
		Table:          "media_other",
		PartitionField: "media_datetime"},

	"medium_dam": &Struct_Table{
		Table:          "medium_dam",
		PartitionField: "mediumdam_date",
		MasterId:       "mediumdam_id",
		MasterTable:    "m_medium_dam",
		HasProvince:    true,
	},

	"metadata": &Struct_Table{
		Table:    "metadata",
		IsMaster: true,
	},

	// เพิ่มเงื่อนไข query ข้อมูล value <> 999999
	//Where:          "pressure_value <> " + filterValue,
	"pressure": &Struct_Table{
		Table:          "pressure",
		PartitionField: "pressure_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		Where:          "pressure_value <> " + filterValue,
		HasProvince:    true,
		HasBasin:       true,
	},

	//HasProvince = join table province
	//	HasBasin = join table basiin
	"rainfall": &Struct_Table{
		Table:          "rainfall",
		PartitionField: "rainfall_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		Where:          "(rainfall10m is not null OR rainfall1h is not null OR rainfall24h is not null OR rainfall_today is not null) AND EXTRACT(MINUTE FROM rainfall_datetime) = 0",
		// WhereHAII:      " m.tele_station_type NOT IN ( 'W' )",
		HasProvince: true,
		HasBasin:    true},

	"rainfall_1h": &Struct_Table{
		Table:          "rainfall_1h",
		PartitionField: "rainfall_datetime",
		Fields:         "id, tele_station_id",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"rainfall_24h": &Struct_Table{
		Table:          "rainfall_24h",
		PartitionField: "rainfall_datetime",
		Fields:         "id, tele_station_id",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"rainfall_daily": &Struct_Table{
		Table:          "rainfall_daily",
		PartitionField: "rainfall_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		WhereHAII:      " m.tele_station_type = 'R' ",
		Where:          "rainfall_value <> " + filterValue,
		HasProvince:    true,
		HasBasin:       true},

	"rainfall_today": &Struct_Table{
		Table:          "rainfall_today",
		PartitionField: "rainfall_datetime",
		Fields:         "id, tele_station_id",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"rainfall_yearly": &Struct_Table{
		Table:          "rainfall_yearly",
		PartitionField: "rainfall_datetime",
		Fields:         "id, tele_station_id",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"rainfall_monthly": &Struct_Table{
		Table:          "rainfall_monthly",
		PartitionField: "rainfall_datetime",
		Fields:         "id, tele_station_id",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"rainforecast": &Struct_Table{
		Table:          "rainforecast",
		PartitionField: "rainforecast_datetime",
		HasProvince:    true},

	"rainforecast_7day": &Struct_Table{
		Table:          "rainforecast_7day",
		PartitionField: "rainfall_datetime",
		Fields:         "id, geocode_id",
		MasterId:       "geocode_id",
		MasterTable:    "lt_geocode"},

	"rulecurve": &Struct_Table{
		Table:          "rulecurve",
		PartitionField: "rc_datetime",
		MasterTable:    "m_dam",
		MasterId:       "dam_id"},

	"safety_zone": &Struct_Table{
		Table:       "safety_zone",
		IsMaster:    true,
		HasProvince: true},

	"sea_water_forecast": &Struct_Table{
		Table:          "sea_water_forecast",
		PartitionField: "seaforecast_datetime",
		MasterId:       "sea_station_id",
		MasterTable:    "m_sea_station"},

	"sea_waterlevel": &Struct_Table{
		Table:          "sea_waterlevel",
		PartitionField: "waterlevel_datetime",
		MasterId:       "sea_station_id",
		MasterTable:    "m_sea_station"},

	"soilmoisture": &Struct_Table{
		Table:          "soilmoisture",
		PartitionField: "soil_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		Where:          "soil_value <> " + filterValue,
		HasProvince:    true,
		HasBasin:       true},

	"solar": &Struct_Table{
		Table:          "solar",
		PartitionField: "solar_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		Where:          "solar_value <> " + filterValue,
		HasProvince:    true,
		HasBasin:       true},

	"storm": &Struct_Table{
		Table:          "storm",
		PartitionField: "storm_datetime"},

	"subbasin": &Struct_Table{
		Table:    "subbasin",
		Fields:   "id, subbasin_code",
		MasterId: "basin_id"},

	"swan": &Struct_Table{
		Table:          "swan",
		PartitionField: "swan_datetime",
		MasterId:       "swan_station_id",
		MasterTable:    "m_swan_station",
		HasProvince:    true},

	"tele_watergate": &Struct_Table{
		Table:          "tele_watergate",
		PartitionField: "watergate_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		WhereHAII:      "m.tele_station_type = 'G'",
		HasProvince:    true,
		HasBasin:       true},

	"tele_waterlevel": &Struct_Table{
		Table:          "tele_waterlevel",
		PartitionField: "waterlevel_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		WhereHAII:      "m.tele_station_type IN ('W','A')",
		WhereHydro:     "hydro_id IS NOT NULL ",
		HasProvince:    true,
		HasBasin:       true},

	"temperature": &Struct_Table{
		Table:          "temperature",
		PartitionField: "temp_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		Where:          "temp_value <> " + filterValue,
		HasProvince:    true,
		HasBasin:       true},

	"temperature_daily": &Struct_Table{
		Table:          "temperature_daily",
		PartitionField: "temperature_date",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"water_resource": &Struct_Table{Table: "water_resource",
		PartitionField: "created_at"},

	"waterquality": &Struct_Table{
		Table:          "waterquality",
		PartitionField: "waterquality_datetime",
		MasterId:       "waterquality_id",
		MasterTable:    "m_waterquality_station",
		HasProvince:    true,
		HasBasin:       true},

	"weather_forecast": &Struct_Table{
		Table:          "weather_forecast",
		PartitionField: "weather_date",
		HasProvince:    true},

	//wind_dir_value
	//wind_speed
	"wind": &Struct_Table{
		Table:          "wind",
		PartitionField: "wind_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station",
		HasProvince:    true,
		HasBasin:       true},

	"rainfall_monthly_tmd": &Struct_Table{
		Table:          "rainfall_monthly_tmd",
		PartitionField: "rainfall_datetime",
		MasterId:       "tele_station_id",
		MasterTable:    "m_tele_station"},

	"m_floodroad_station": &Struct_Table{
		Table:        "m_floodroad_station",
		SelectColumn: " floodroad_station_name, floodroad_station_lat, floodroad_station_long ",
		Fields:       "id, floodroad_station_oldcode",
		IsMaster:     true,
		HasProvince:  true},

	"floodroad": &Struct_Table{
		Table:          "floodroad",
		PartitionField: "floodroad_datetime",
		MasterId:       "floodroad_station_id",
		MasterTable:    "m_floodroad_station",
		HasProvince:    true},

	"m_flow_station": &Struct_Table{
		Table:        "m_flow_station",
		SelectColumn: " flow_station_name, flow_station_lat, flow_station_long ",
		Fields:       "id, flow_station_oldcode",
		IsMaster:     true,
		HasProvince:  true},

	"flow": &Struct_Table{
		Table:          "flow",
		PartitionField: "flow_datetime",
		MasterId:       "flow_station_id",
		MasterTable:    "m_flow_station",
		HasProvince:    true},
}

//	get map table
//	Parameters:
//		table
//			ชื่อตาราง
//	Return:
//		map table ที่เตรียมไว้ล่วงหน้า
func GetTable(table string) *Struct_Table {
	if s, ok := mapStrctTable[table]; ok {
		return s
	}
	return nil
}

//	get parition field from map table
//	Parameters:
//		table
//			ชื่อตาราง
//	Return:
//		parition field จาก map table ที่เตรียมไว้ล่วงหน้า
func GetTablePartitionField(table string) string {
	s := GetTable(table)
	if s != nil {
		return s.PartitionField
	}
	return ""
}

//	get master id field from map table
//	Parameters:
//		table
//			ชื่อตาราง
//	Return:
//		master id field จาก map table ที่เตรียมไว้ล่วงหน้า
func GetTableMasterId(table string) string {
	s := GetTable(table)
	if s != nil {
		return s.MasterId
	}
	return ""
}

//	get master table from map table
//	Parameters:
//		table
//			ชื่อตาราง
//	Return:
//		master table จาก map table ที่เตรียมไว้ล่วงหน้า
func GetTableMasterTable(table string) string {
	s := GetTable(table)
	if s != nil {
		return s.MasterTable
	}
	return ""
}

//	get select columm from map table
//	Parameters:
//		table
//			ชื่อตาราง
//	Return:
//		select column จาก map table ที่เตรียมไว้ล่วงหน้า
func GetTableSelectColum(table string) string {
	s := GetTable(table)
	if s != nil {
		return s.SelectColumn
	}
	return ""
}

//	get select columm from map table
//	Parameters:
//		table
//			ชื่อตาราง
//	Return:
//		select column จาก map table ที่เตรียมไว้ล่วงหน้า
func GetSelectColumn(table string) string {
	return GetTableSelectColum(table)
}

// 	เช็คว่าเป็น table media ?
// 	Parameters:
//		s
//			ชื่อตาราง
//	Return:
//		true ถ้าเป็นตาราง media
func IsMedia(s string) bool {
	if s == "media" || s == "latest_media" || s == "media_animation" || s == "media_other" {
		return true
	}
	return false
}

//	เช็คว่าเป็นรหัสหน่วยงานของ สสนก ?
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//	Return:
//		true ถ้าเป็นหน่วยงาน สสนก
func IsHAII(agency_id string) bool {
	if agency_id == "9" {
		return true
	}
	return false
}
