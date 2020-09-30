package rainforecast_7day_province

type Struct_RainForecastData struct {
	Forecast_day  		string `json:"forecast_day"`  		// example:`1` วัน
	Forecast_datetime  	string `json:"forecast_datetime"`  // example:`2017-01-14 00:00:00` วันที่คาดการณ์
	Province_id    	string `json:"province_id"`    		// example:`19` รหัสจังหวัด
	Province_name  	string `json:"province_name"`  		// example:`สระบุรี` ชื่อจังหวัด
	Rainfall       	string `json:"rainfall"`       		// example:`0.0253` ปริมาณฝน
	Rainfall_text  	string `json:"rainfall_text"`  		// example:`ไม่มีฝน` สถานะ
}
