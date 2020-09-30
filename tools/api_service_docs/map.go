package main

import ()

var MapFieldExample map[string]map[string]string = map[string]map[string]string{
	// air
	"air": map[string]string{
		"agency_id":      "14",
		"air_aqi":        "1",
		"air_co":         "0.12",
		"air_datetime":   "2006-01-02T15:04:05Z07:00",
		"air_no2":        "1",
		"air_o3":         "3",
		"air_pm10":       "1",
		"air_pm25":       "",
		"air_so2":        "1",
		"air_station_id": "9",
		"id":             "9",
		"qc_status":      "null",
	},

	// canal waterlevel
	"canal_waterlevel": map[string]string{
		"canal_station_id":          "9",
		"canal_waterlevel_datetime": "2006-01-02T15:04:05Z07:00",
		"canal_waterlevel_value":    "2.99",
		"comm_status":               "Success, Fail",
		"id":                        "3194172",
		"qc_status":                 "null",
	},

	// crosssection
	"crosssection": map[string]string{
		"distance":           "23",
		"id":                 "13",
		"point_id":           "",
		"qc_status":          "null",
		"remark":             "LB",
		"section_lat":        "",
		"section_long":       "",
		"section_station_id": "9",
		"water_level_m":      "2",
		"water_level_msl":    "12",
	},

	// dam_daily
	"dam_daily": map[string]string{
		"dam_date":               "2006-01-02",
		"dam_evap":               "0.0187",
		"dam_id":                 "9",
		"dam_inflow":             "2.24",
		"dam_inflow_acc":         "130",
		"dam_inflow_acc_percent": "48.529999",
		"dam_inflow_avg":         "24.167",
		"dam_level":              "57.57",
		"dam_losses":             "0",
		"dam_released":           "0.09",
		"dam_released_acc":       "60.988896",
		"dam_spilled":            "2.7479",
		"dam_storage":            "10.0249",
		"dam_storage_percent":    "26.959999",
		"dam_uses_water":         "40.560001",
		"dam_uses_water_percent": "17.059999",
		"id":        "3628",
		"qc_status": "null",
	},

	// dam_hourly
	"dam_hourly": map[string]string{
		"dam_datetime":           "2006-01-02T15:04:05Z07:00",
		"dam_evap":               "0.0187",
		"dam_id":                 "9",
		"dam_inflow":             "2.24",
		"dam_inflow_acc":         "130",
		"dam_inflow_acc_percent": "48.529999",
		"dam_inflow_avg":         "24.167",
		"dam_level":              "57.57",
		"dam_losses":             "0",
		"dam_released":           "0.09",
		"dam_released_acc":       "60.988896",
		"dam_spilled":            "2.7479",
		"dam_storage":            "10.0249",
		"dam_storage_percent":    "26.959999",
		"dam_uses_water":         "40.560001",
		"dam_uses_water_percent": "17.059999",
		"id":        "3628",
		"qc_status": "null",
	},

	// flood_situation
	"flood_situation": map[string]string{
		"agency_id":         "2",
		"flood_author":      "http://www.dmr.go.th/",
		"flood_colorlevel":  " ",
		"flood_datetime":    "2006-01-02T15:04:05Z07:00",
		"flood_description": " ",
		"flood_link":        "http://www.dmr.go.th/ewt_news.php?nid=102862",
		"flood_name":        "รายงานสถานการณ์ธรณีพิบัติภัยประจำวัน วันอังคารที่ ๒๒ สิงหาคม พ.ศ. ๒๕๖๐",
		"flood_remark":      " ",
		"geocode_id":        " ",
		"id":                "1",
	},

	// floodforecast
	"floodforecast": map[string]string{
		"floodforecast_datetime":   "2006-01-02T15:04:05Z07:00",
		"floodforecast_station_id": "9",
		"floodforecast_value":      "31",
		"id":        "21",
		"qc_status": "null",
	},

	// ford_waterlevel
	"ford_waterlevel": map[string]string{
		"comm_status":              "Connect, Fail",
		"ford_station_id":          "39",
		"ford_waterlevel_datetime": "2006-01-02T15:04:05Z07:00",
		"ford_waterlevel_value":    "39.3",
		"id":        "12301159",
		"qc_status": "null",
	},

	// geohazard_situation
	"geohazard_situation": map[string]string{
		"agency_id":             "2",
		"geocode_id":            "1601",
		"geohazard_author":      "http://www.dmr.go.th/",
		"geohazard_colorlevel":  "null",
		"geohazard_datetime":    "2006-01-02T15:04:05Z07:00",
		"geohazard_description": "null",
		"geohazard_link":        "http://www.dmr.go.th/ewt_news.php?nid=102885",
		"geohazard_name":        "บ.ใหม่ ม.6 ต.สาริกา อ.เมือง จ.นครนายก_เช้านี้ท้องฟ้าโปร่ง อากาศแจ่มใส เมื่อวานนี้มีฝนตกเล็กน้อย",
		"geohazard_remark":      "null",
		"id":                    "61",
	},

	// humid
	"humid": map[string]string{
		"humid_datetime":  "2006-01-02T15:04:05Z07:00",
		"humid_value":     "95.33",
		"id":              "73346851",
		"qc_status":       "null",
		"tele_station_id": "899",
	},

	// m_air_station
	"m_air_station": map[string]string{
		"agency_id":           "14",
		"air_staiton_type":    "GROUND",
		"air_station_lat":     "13.783143",
		"air_station_long":    "100.540529",
		"air_station_name":    `{"en":"The Government Public Relations Department","th":"กรมประชาสัมพันธ์"}`,
		"air_station_oldcode": "59t",
		"id": "9",
	},

	// m_dam
	"m_dam": map[string]string{
		"agency_id":               "8",
		"avg_inflow":              "2578.85",
		"avg_inflow_endyear":      "",
		"avg_inflow_intyear":      "1986",
		"dam_lat":                 "8.966667",
		"dam_long":                "98.783333",
		"dam_name":                `{"th":"รัชชประภา","en":"RAJJAPRABHA DAM","jp":" "}`,
		"dam_oldcode":             "15",
		"downstream_storage":      "1100",
		"emer_watergate_level":    "",
		"geocode_id":              "7580",
		"id":                      "49",
		"max_inflow":              "161.62",
		"max_inflow_date":         "1997-08-24",
		"max_old_storage":         "90.7",
		"max_storage":             "6144.38",
		"max_water_level":         "97.65",
		"maxos_date":              "1997-09-28",
		"min_old_storage":         "64.39",
		"min_storage":             "1351.54",
		"min_water_level":         "62",
		"minos_date":              "1992-07-25",
		"normal_storage":          "5638.84",
		"normal_water_level":      "95",
		"normal_watergate_level":  "95",
		"power_install":           "240",
		"power_intake_level":      "53",
		"power_intake_storage":    "721.17",
		"rainfall_yearly":         "1967.46",
		"ridge_spillway_level":    "87.5",
		"service_watergate_level": "",
		"subbasin_id":             "",
		"tailrace_level":          "12",
		"top_spillway_level":      "95.5",
		"used_genpower":           "7.98",
		"uses_water":              "4200",
		"water_shed":              "1435",
	},

	// m_floodforecast_station
	"m_floodforecast_station": map[string]string{
		"agency_id":                      "9",
		"floodforecast_station_alarm":    "33.15",
		"floodforecast_station_critical": "34.34",
		"floodforecast_station_lat":      "16.270330",
		"floodforecast_station_long":     "100.413960",
		"floodforecast_station_name":     "",
		"floodforecast_station_oldcode":  "NAN007",
		"floodforecast_station_type":     "ระดับน้ำ",
		"floodforecast_station_unit":     "ม.รทก.",
		"floodforecast_station_warning":  "33.75",
		"geocode_id":                     "6269",
		"id":                             "9",
		"subbasin_id":                    "1",
	},

	// m_swan_station
	"m_swan_station": map[string]string{
		"agency_id":    "9",
		"geocode_id":   "256",
		"id":           "9",
		"swan_lat":     "12.871000",
		"swan_long":    "100.844000",
		"swan_name":    `{"en":"Pattaya"}`,
		"swan_oldcode": "",
	},

	// m_tele_station
	"m_tele_station": map[string]string{
		"agency_id":             "9",
		"distance":              "",
		"floodgate":             "",
		"geocode_id":            "41",
		"ground_level":          "-2.678",
		"id":                    "19",
		"left_bank":             "1.782",
		"max_waterlevel_20y":    "",
		"pump":                  "",
		"right_bank":            "0.646",
		"riverbank":             "",
		"sort_order":            "",
		"subbasin_id":           "198",
		"tele_station_lat":      "13.589267",
		"tele_station_long":     "100.802235",
		"tele_station_name":     `{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}`,
		"tele_station_oldcode":  "BKK012",
		"tele_station_type":     "W",
		"water_storage_station": "",
		"waterflow_time":        "",
	},

	// media
	"media": map[string]string{
		"agency_id":      "13",
		"media_type_id":  "30",
		"media_datetime": "2006-01-02T15:04:05Z07:00",
		"media_path":     "http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/image?image=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHNIfCEaLG3zGgZ2LEuTaHsMjGebDH8UN",
		"media_desc":     "เรดาร์ตรวจอากาศ",
		"filename":       "cri240_201612090100.gif",
		"refer_source":   "",
	},

	// media_animation
	"media_animation": map[string]string{
		"agency_id":      "50",
		"media_type_id":  "141",
		"media_datetime": "2006-01-02T15:04:05Z07:00",
		"media_path":     "http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/image?image=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJBMDrGmg4sd2EqTleGoMGzbfRwOOn9GdVm9blDTBH42TRFzCh4Sws-QyEtcntJfW62mIK0Q==",
		"media_desc":     "00Latest.jpg",
		"filename":       "ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi",
		"refer_source":   "http://weather.is.kochi-u.ac.jp/SE/00Latest.jpg",
	},

	// media_other
	"media_other": map[string]string{
		"agency_id":      "50",
		"media_type_id":  "141",
		"media_datetime": "2006-01-02T15:04:05Z07:00",
		"media_path":     "http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/image?image=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJBMDrGmg4sd2EqTleGoMGzbfRwOOn9GdVm9blDTBH42TRFzCh4Sws-QyEtcntJfW62mIK0Q==",
		"media_desc":     "00Latest.jpg",
		"filename":       "ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi",
		"refer_source":   "http://weather.is.kochi-u.ac.jp/SE/00Latest.jpg",
	},

	// medium_dam
	"medium_dam": map[string]string{
		"id":                        "274200",
		"mediumdam_date":            "2006-01-02",
		"mediumdam_id":              "9",
		"mediumdam_inflow":          "0.878",
		"mediumdam_released":        "1.76",
		"mediumdam_storage":         "15.801",
		"mediumdam_storage_percent": "57.043",
		"mediumdam_uses_water":      "14.651",
		"qc_status":                 "null",
	},

	// pressure
	"pressure": map[string]string{
		"id":                "74105189",
		"pressure_datetime": "2006-01-02T15:04:05Z07:00",
		"pressure_value":    "1012.67",
		"qc_status":         "null",
		"tele_station_id":   "899",
	},

	// rainfall
	"rainfall": map[string]string{
		"id":                 "109503496",
		"qc_status":          "null",
		"rainfall10m":        "0",
		"rainfall12h":        "4.5",
		"rainfall15m":        "0",
		"rainfall1h":         "1.5",
		"rainfall24h":        "12.5",
		"rainfall30m":        "0",
		"rainfall3h":         "3",
		"rainfall5m":         "0",
		"rainfall6h":         "3.5",
		"rainfall_acc":       "241",
		"rainfall_date_calc": "2006-01-02",
		"rainfall_datetime":  "2006-01-02T15:04:05Z07:00",
		"rainfall_today":     "",
		"tele_station_id":    "2073",
	},

	// rainfall_daily
	"rainfall_daily": map[string]string{
		"id":                 "612726",
		"qc_status":          "null",
		"rainfall_date_calc": "2006-01-02",
		"rainfall_datetime":  "2006-01-02T15:04:05Z07:00",
		"rainfall_value":     "3.2",
		"tele_station_id":    "135",
	},

	// rainforecast
	"rainforecast": map[string]string{
		"agency_id":              "",
		"geocode_id":             "325",
		"id":                     "9",
		"qc_status":              "null",
		"rainforecast_datetime":  "2006-01-02T15:04:05Z07:00",
		"rainforecast_level":     "2",
		"rainforecast_leveltext": "ฝนตกเล็กน้อย",
		"rainforecast_value":     "2.3596",
	},

	// soilmoisture
	"soilmoisture": map[string]string{
		"id":              "37672770",
		"qc_status":       "null",
		"soil_datetime":   "2006-01-02T15:04:05Z07:00",
		"soil_value":      "61.6",
		"tele_station_id": "1969",
	},

	// solar
	"solar": map[string]string{
		"id":              "73447761",
		"qc_status":       "null",
		"solar_datetime":  "2006-01-02T15:04:05Z07:00",
		"solar_value":     "120.37",
		"tele_station_id": "151",
	},

	// swan
	"swan": map[string]string{
		"id":                  "1",
		"qc_status":           "null",
		"swan_datetime":       "2006-01-02T15:04:05Z07:00",
		"swan_depth":          "2.3737",
		"swan_direction":      "88.637",
		"swan_highsig":        "0.10686",
		"swan_period_average": "2.6946",
		"swan_period_top":     "3.3597",
		"swan_station_id":     "1",
		"swan_windx":          "1.0075",
		"swan_windy":          "0.3487",
	},

	// tele_watergate
	"tele_watergate": map[string]string{
		"floodgate_height":   "",
		"floodgate_open":     "",
		"id":                 "785132",
		"pump_on":            "",
		"qc_status":          "null",
		"tele_station_id":    "956",
		"watergate_datetime": "2006-01-02T15:04:05Z07:00",
		"watergate_in":       "0.04",
		"watergate_out":      "0.44",
		"watergate_out2":     "0",
	},

	// tele_waterlevel
	"tele_waterlevel": map[string]string{
		"discharge":           "",
		"flow_rate":           "",
		"id":                  "11697144",
		"qc_status":           "null",
		"tele_station_id":     "3458",
		"waterlevel_datetime": "2006-01-02T15:04:05Z07:00",
		"waterlevel_m":        "",
		"waterlevel_msl":      "22.549",
	},

	// temperature
	"temperature": map[string]string{
		"id":              "127726769",
		"qc_status":       "null",
		"tele_station_id": "899",
		"temp_datetime":   "2006-01-02T15:04:05Z07:00",
		"temp_value":      "26.46",
	},

	// temperature_daily
	"temperature_daily": map[string]string{
		"diffmaxtemperature": "-1.3",
		"diffmintemperature": "-0.7",
		"id":                 "855",
		"maxtemperature":     "34.5",
		"mintemperature":     "25.2",
		"qc_status":          "null",
		"tele_station_id":    "3634",
		"temperature_date":   "2006-01-02",
		"temperature_value":  "26.2",
	},

	// water_resource
	"water_resource": map[string]string{
		"agency_id":              "5",
		"benefit_area":           "1100",
		"benefit_household":      "250",
		"budget":                 "0",
		"capacity":               "3700",
		"contract_enddate":       "2004-07-29",
		"contract_signdate":      "2004-03-01",
		"coordination":           "MB338244",
		"fiscal_year":            "2547",
		"geocode_id":             "4966",
		"id":                     "4894",
		"mooban":                 "เมืองแปง ม.1",
		"projectname":            "งานปรับปรุงพื้นที่และจัดทำระบบส่งน้ำในไร่นา ( 70 แห่ง ) ปี 2547",
		"projecttype":            "คลองส่งน้ำ",
		"qc_status":              "null",
		"rec_date":               " ",
		"standard_cost":          "8326000",
		"water_resource_oldcode": "5309",
	},

	// waterquality
	"waterquality": map[string]string{
		"id":                        "15",
		"qc_status":                 "null",
		"waterquality_ammonium":     "0",
		"waterquality_bod":          "129.1",
		"waterquality_chlorophyll":  "6",
		"waterquality_colorstatus":  "",
		"waterquality_conductivity": "232",
		"waterquality_datetime":     "2006-01-02T15:04:05Z07:00",
		"waterquality_do":           "0",
		"waterquality_fcb":          "",
		"waterquality_id":           "117",
		"waterquality_nh3n":         "",
		"waterquality_nitrate":      "0",
		"waterquality_ph":           "4",
		"waterquality_salinity":     "0.09",
		"waterquality_status":       "ปกติ",
		"waterquality_tcb":          "",
		"waterquality_tds":          "130",
		"waterquality_temp":         "30.22",
		"waterquality_turbid":       "0",
		"waterquality_wqi":          "",
	},

	// weather_forecast
	"weather_forecast": map[string]string{
		"agency_id":       "",
		"geocode_id":      "",
		"id":              "",
		"overall_forcast": "",
		"region_forcast":  "",
		"weather_date":    "",
	},

	// wind
	"wind": map[string]string{
		"id":              "12168431",
		"qc_status":       "null",
		"tele_station_id": "60",
		"wind_datetime":   "2006-01-02T15:04:05Z07:00",
		"wind_dir":        "",
		"wind_dir_value":  "228",
		"wind_speed":      "8",
	},

	"all": map[string]string{
		"province_name": `{"th": "กรุงเทพมหานคร"}`,
		"amphoe_name":   `{"th": "พระบรมมหาราชวัง"}`,
		"tumbon_name":   `{"th": "พระนคร"}`,
	},
}
