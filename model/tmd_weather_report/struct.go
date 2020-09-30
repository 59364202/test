package tmd_weather_report

type Struct_Weather_today struct {
	Id                     int64  `json:"id"`
	Wmo_code               int64  `json:"wmo_code"`
	Import_date            string `json:"import_date"`
	Mean_sea_level_presure string `json:"mean_sea_level_pressure"`
	Temperature            string `json:"temperature"`
	Max_temperature        string `json:"max_temperature"`
	Diff_max_temperature   string `json:"diff_max_temperature"`
	Min_temperature        string `json:"min_temperature"`
	Diff_min_temperature   string `json:"diff_min_temperature"`
	Relative_humidity      string `json:"relative_humidity"`
	Wind_direction         string `json:"wind_direction"`
	Wind_speed             string `json:"wind_speed"`
	Rainfall               string `json:"rainfall"`
}
