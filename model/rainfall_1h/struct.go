package rainfall_1h

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

type Struct_Rainfall1h_Graph struct {
	DateTime string      `json:"rainfall_datetime,omitempty"` // example:`2006-01-02 15:04` วันที่เก็บปริมาณน้ำฝน
	Rainfall interface{} `json:"rainfall_value"`              // example:`1.5` ค่าฝน (mm.)
}

type Param_Rainfall1h_Graph struct {
	Is24    bool
	IsDaily bool
	IsToday bool

	StationId int64  `json:"station_id"` // รหัสสถานี
	DateStart string `json:"date_start"` // วันที่เริ่มต้น
	DateEnd   string `json:"date_end"`   // วันที่สิ้นสุด
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

type Struct_MaxMin struct {
	Max *Struct_Rainfall1h `json:"max,omitempty"` // ฝนสูงสุด
	Min *Struct_Rainfall1h `json:"min,omitempty"` // ฝนต่ำสุด
}

type Struct_TeleStation struct {
	Id      int64           `json:"id"`                          // example:`1` รหัสสถานี
	Name    json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"คลองลาดพร้าว วัดบางบัว"}` ชื่อสถานี
	Lat     interface{}     `json:"tele_station_lat"`            // example:`13.854100` ละติจูดของสถานี
	Long    interface{}     `json:"tele_station_long"`           // example:`100.588000` ลองติจูดของสถานี
	OldCode string          `json:"tele_station_oldcode"`        // example:`BKK021` รหัสสถานีเดิม
}

type Struct_Rainfall1h struct {
	Id       int64       `json:"id,omitempty"`       // example:`20753329` รหัสข้อมูล
	Rainfall interface{} `json:"rainfall,omitempty"` // example:`23.5` ฝนรายชม.
	Time     string      `json:"rainfall_datetime"`  // example:`2006-01-02 15:04`  วันที่

	Agency  *model_agency.Struct_Agency   `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"` // สถานี
}
