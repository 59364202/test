// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package humid

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

type Struct_GraphParam struct {
	Station_id string `json:"station_id"` // รหัสสถานี เช่น 10,23,413
	Start_date string `json:"start_date"` // เวลาเริ่มต้น 2006-01-02
	End_date   string `json:"end_date"`   // เวลาสิ้นสุด 2006-01-02
}

type Struct_Graph struct {
	GraphData []*Struct_GraphByStationAndDate `json:"graph_data"` // ข้อมูลกราฟ
}

type Struct_GraphByStationAndDate struct {
	Datetime string      `json:"datetime"` // example:`2006-01-02 15:04` วันเวลา
	Value    interface{} `json:"value"`    // example:`64.91` ความชื้น
}

// param for input query condition
type Param_Provinces struct {
	Region_Code     string `json:"region_code"`     // required:false example:`1` รหัสภาค ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Region_Code_tmd string `json:"region_code_tmd"` // required:false example:`1` รหัสภาค 6 ภาคตามการแบ่งของกรมอุตุฯ ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Province_Code   string `json:"province_code"`   // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Data_Limit      int    `json:"data_limit"`      // จำนวนข้อมูล
	//IsIgnore       bool   `json:"is_ignore"`
	Order string `json:"order"`
}

type Struct_Humid struct {
	HumidDatetime string      `json:"humid_datetime,omitempty"` // example:`2006-01-02` วันที่
	Humid         interface{} `json:"humid,omitempty"`          // example:`14.8` ค่าความชื้น

	Agency  *model_agency.Struct_Agency   `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"` // สถานี
}

type Struct_TeleStation struct {
	Id      int64           `json:"id"`                          // example:`1` รหัสสถานี
	Name    json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"คลองลาดพร้าว วัดบางบัว"}` ชื่อสถานี
	Lat     interface{}     `json:"tele_station_lat"`            // example:`13.854100` ละติจูดของสถานี
	Long    interface{}     `json:"tele_station_long"`           // example:`100.588000` ลองติจูดของสถานี
	OldCode string          `json:"tele_station_oldcode"`        // example:`BKK021` รหัสสถานีเดิม
}
