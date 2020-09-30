// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package temperature

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

type Struct_MaxMinTemperature struct {
	Start_date      string      `json:"start_date"`                // example:`2006-01-02` วันที่เริ่มคำนวนค่าอุณหภูมิ
	Max_temperature interface{} `json:"max_temperature,omitempty"` // example:`38.5` อุณหภูมิสูงสุด  (°C)
	Min_temperature interface{} `json:"min_temperature,omitempty"` // example:`24` อุณหภูมิต่ำสุด (°C)

	Agency  *model_agency.Struct_Agency   `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"` // สถานี
}

type Struct_TeleStation struct {
	Id          int64           `json:"id"`                          // example:`1` รหัสสถานี
	Name        json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"คลองลาดพร้าว วัดบางบัว"}` ชื่อสถานี
	Lat         interface{}     `json:"tele_station_lat"`            // example:`13.854100` ละติจูดของสถานี
	Long        interface{}     `json:"tele_station_long"`           // example:`100.588000` ลองติจูดของสถานี
	OldCode     string          `json:"tele_station_oldcode"`        // example:`BKK021` รหัสสถานีเดิม
	Basin_id    int64           `json:"basin_id,omitempty"`          // example:`01` รหัสลุ่มน้ำ
	SubBasin_id int64           `json:"sub_basin_id,omitempty"`      // example:`21` รหัสลุ่มน้ำย่อย
}

type Struct_Temperature struct {
	Id          int64       `json:"id,omitempty"`          // example:`20753329` รหัสข้อมูลอุณหภูมิ
	Temperature interface{} `json:"temperature,omitempty"` // example:`7` อุณหภูมิ (°C)
	Time        string      `json:"temperature_datetime"`  // example:`2006-01-02 15:04`  วันที่เก็บค่าอุณหภูมิ

	Agency  *model_agency.Struct_Agency   `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"` // สถานี
}

type Struct_TemperatureLatest struct {
	SeasonMaxMin string                `json:"season_maxmin,omitempty"` // example:`อุณหภูมิสูงสุด`  แสดงข้อความอุณหภูมิสูงสุด/ต่ำสุด
	Data         []*Struct_Temperature `json:"data"`                    // ข้อมูลกราฟ
}

type Struct_TemperatureMaxMinByRegion struct {
	Struct_Temperature
	Humid    interface{} `json:"humid,omitempty"`    // example:`65.3`  ความชื้นสัมพัทธ์ (%RH)
	Pressure interface{} `json:"pressure,omitempty"` // example:`105.3` ความกดอากาศ หน่วย : mBar
}

// param for input query condition
type Param_TemperatureProvinces struct {
	Region_Code     string `json:"region_code"`     // required:false example:`1` รหัสภาค ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Region_Code_tmd string `json:"region_code_tmd"` // required:false example:`1` รหัสภาค 6 ภาคตามการแบ่งของกรมอุตุฯ ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Province_Code   string `json:"province_code"`   // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Data_Limit      int    `json:"data_limit"`      // จำนวนข้อมูล
	//IsIgnore       bool   `json:"is_ignore"`
	Order string `json:"order"`
}

type Struct_TemperatureGraphParam struct {
	Station_id string `json:"station_id"` // รหัสสถานี เช่น 10,23,413
	Start_date string `json:"start_date"` // เวลาเริ่มต้น 2006-01-02
	End_date   string `json:"end_date"`   // เวลาสิ้นสุด 2006-01-02
}

type Struct_TemperatureGraph struct {
	GraphData []*Struct_GraphByStationAndDate `json:"graph_data"` // ข้อมูลกราฟ
}

type Struct_GraphByStationAndDate struct {
	Datetime string      `json:"datetime"` // example:`2006-01-02 15:04` วันเวลา
	Value    interface{} `json:"value"`    // example:`137.91` อุณหภูมิ (°C)
}

type Struct_TemperatureMaxMin struct {
	Max *Struct_Temperature `json:"max,omitempty"` // อุณหภูมิสูงสุด
	Min *Struct_Temperature `json:"min,omitempty"` // อุณหภูมิต่ำสุด
}
