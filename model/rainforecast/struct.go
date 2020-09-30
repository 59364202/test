package rainforecast

import (
	"encoding/json"
)

type Rainforecast struct {
	Id                     int64   `json:"id"`                              // รหัสข้อมูลคาดการณ์ฝน
	Rainforecast_Datetime  string  `json:"rainforecast_datetime,omitempty"` // วันเวลาที่เก็บข้อมูลคาดการณ์ฝน
	Rainforecast_Value     float64 `json:"rainforecast_value,omitempty"`    // ข้อมูลคาดการณ์ฝน
	Rainforecast_Level     string  `json:"rainforecast_level"`              // เกณฑ์ของการคาดการณ์
	Rainforecast_Leveltext string  `json:"rainforecast_leveltext"`          // รายละเอียดเกณฑ์ของการคาดการณ์
}

type Struct_RainforecastCurrentDate struct {
	Province_Code string          `json:"province_code"` // example:`10` รหัสจังหวัด
	Province_Name json.RawMessage `json:"province_name"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัด
	Value         float64         `json:"value"`         // example:`70` ข้อมูลคาดการณ์ฝน
	Area_Code     string          `json:"area_code"`     // example:`1` รหัสภาค
	Zone_Code     string          `json:"zone_code"`     // example:`06` รหัสภาคของประเทศไทยตามกรมอุตุนิยมวิทยา
}

type Struct_Rainforecast struct {
	Province_Code      string           `json:"province_code"`      // example:`10` รหัสจังหวัด
	Province_Name      *json.RawMessage `json:"province_name"`      // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัด
	Rainforecast_Level int64            `json:"rainforecast_level"` // example:`4` เกณฑ์ของการคาดการณ์
}

type Struct_RainForecast3Day struct {
	Date      string `json:"date"`      // example:`วันนี้` วันที่
	Status    string `json:"status"`    // example:`ไม่มีฝน` สถานะการคาดการณ์
	Status_id int64  `json:"status_id"` // example:`1` รหัสสถานะการคาดการณ์
}

type Struct_RainForcastRegion struct {
	Date   string                             `json:"date"`   // example:`พุธ 25 พ.ค.` วันที่
	Region []*Struct_RainForcastRegion_Region `json:"region"` // array ของรหัสภาคและสถานะฝน
}
type Struct_RainForcastRegion_Region struct {
	Region_id int64 `json:"region_id"` // exmaple:`1` รหัสภาค
	Status    int64 `json:"status"`    // example:`2` สถานะฝน
}

//	create new Struct_RainForcastRegion
//	Return:
//		new Struct_RainForcastRegion
func (s *Struct_RainForcastRegion) New(date string) *Struct_RainForcastRegion {
	st := new(Struct_RainForcastRegion)

	st_r := make([]*Struct_RainForcastRegion_Region, 0)

	st.Date = date
	st.Region = st_r

	return st
}

//	Add new Region into  Struct_RainForcastRegion.Region
//	Parameters:
//		region_id
//			รหัสภาค
//		status
//			สถานะฝน
func (s *Struct_RainForcastRegion) AddRegion(region_id, status int64) {
	r := &Struct_RainForcastRegion_Region{
		Region_id: region_id,
		Status:    status}
	s.Region = append(s.Region, r)
}

type Struct_RainForecastData struct {
	Forecast_date  string `json:"forecast_date"`  // example:`2017-01-14` วันที่คาดการณ์
	Province_id    string `json:"province_id"`    // example:`19` รหัสจังหวัด
	Province_name  string `json:"province_name"`  // example:`สระบุรี` ชื่อจังหวัด
	Rainfall       string `json:"rainfall"`       // example:`0.0253` ปริมาณฝน
	Rainfall_text  string `json:"rainfall_text"`  // example:`ไม่มีฝน` สถานะ
	Rainfall_level int64  `json:"rainfall_level"` // example:`1` ระดับสถานะ
}

//"forecast_date": "2018-01-17",
//"province_id": "10",
//"rainfall": "0.0000",
//"rainfall_text": "ไม่มีฝน",
//"rainfall_level": 1,
//"date_thai": "17 ม.ค. 2561",
//"province_name": "กรุงเทพมหานคร",
//"image": "http://live1.haii.or.th/product/latest/wrfroms/NHC_Province/Lampang_day1.png"
