package rainfall24hr

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
)

type Rainfall24HrStruct struct {
	model_tele_station.TeleStationStruct
	Rain        float64 `json:"rain,omitempty"`         // example:`19` ฝน 24 ชม (mm.)
	Rain1H      float64 `json:"rain1h,omitempty"`       // example:`7` ฝน 1 ชม (mm.)
	Time        string  `json:"time,omitempty"`         // example:`2006-01-02 15:04` วันที่เก็บปริมาณน้ำฝน
	WarningZone string  `json:"warning_zone,omitempty"` // โซนเตือนภัยพิจารณาตามพื้นที่

	StationType string `json:"station_type,omitempty"` // example:`rainfall_24h` ชนิดของสถานี
}

type Struct_Rainfall24H struct {
	Id          int64       `json:"id"`                // example:`20753329` รหัสข้อมูลปริมาณน้ำฝน
	Rain24H     interface{} `json:"rain_24h"`          // example:`19` ฝน 24 ชม (mm.)
	Rain1H      interface{} `json:"rain_1h,omitempty"` // example:`7` ฝน 1 ชม (mm.)
	Time        string      `json:"rainfall_datetime"` // example:`2006-01-02 15:04` วันที่เก็บปริมาณน้ำฝน
	StationType string      `json:"station_type"`      // example:`rainfall_24h` ชนิดของสถานี

	Agency  *model_agency.Struct_Agency   `json:"agency"`          // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"`         // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"`         // สถานี
	Basin   *model_basin.Struct_Basin     `json:"basin,omitempty"` // ลุ่มน้ำ
}

type Struct_TeleStation struct {
	Id          int64           `json:"id"`                          // example:`1` รหัสสถานี
	Name        json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"คลองลาดพร้าว วัดบางบัว"}` ชื่อสถานี
	Lat         interface{}     `json:"tele_station_lat"`            // example:`13.854100` ละติจูดของสถานี
	Long        interface{}     `json:"tele_station_long"`           // example:`100.588000` ลองติจูดของสถานี
	OldCode     string          `json:"tele_station_oldcode"`        // example:`BKK021` รหัสสถานีเดิม
	Type        string          `json:"tele_station_type"`           // example:`R` ชนิดของสถานี
	SubBasin_id int64           `json:"sub_basin_id,omitempty"`      // example:`21` รหัสลุ่มน้ำย่อย
}

// param for input query condition
type Param_Rainfall24 struct {
	Region_Code   string `json:"region_code"`   // required:false example:`1` รหัสภาค ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Start_Date    string `json:"start_date"`    // วันที่เริ่มต้น
	End_Date      string `json:"end_date"`      // วันที่สิ้นสุด
	Data_Limit    int    `json:"data_limit"`    // จำนวนข้อมูล
	IsDaily       bool   `json:"isDaily"`       // เป็นรายวัน
	Is24          bool
	IsHourly      bool
	Include_zero  string `json:"include_zero"` // required:false example:`1` ต้องการให้แสดงข้อมูลสถานนีที่มีค่าฝนเป็น 0 ด้วยหรือไม่ 1=แสดง
	//IsIgnore       bool   `json:"is_ignore"`
	Order    string `json:"order"`
	Basin_id string `json:"basin_id"` // required:false example:`1` รหัสลุ่มน้ำ ไม่ใส่ = ทุกลุ่มน้ำ เลือกได้หลายลุ่มน้ำ เช่น 1,2,4
}

type Struct_AdvRainDiagram struct {
	Name  *json.RawMessage `json:"name"`  // example:`อ.ตรอน (N.60)` ชื่อสถานี
	Lat   interface{}      `json:"lat"`   // example:`17.41439` ละติจูดของสถานี
	Long  interface{}      `json:"long"`  // example:`100.1238` ลองติจูดของสถานี
	Value interface{}      `json:"value"` // example:`19` ค่า
}

type Struct_AdvOnload_Agency struct {
	AgencyId   int64           `json:"agency_id"`   // example:`1` รหัสหน่วยงาน
	AgencyName json.RawMessage `json:"agency_name"` // example:`{"th":"กรมเจ้าท่า","en":"Marine Department","jp":""}` ชื่อหน่วยงาน
	MaxYear    interface{}     `json:"max_year"`    // example:`2006` ปีที่มากที่สุดของข้อมูล
	MinYear    interface{}     `json:"min_year"`    // example:`2006` ปีที่น้อยที่สุดของข้อมูล
}

type Param_AdvRainSum struct {
	AgencyId  int64  `json:"agency_id"`  // example:`9` รหัสหน่วยงาน
	DateStart string `json:"date_start"` // example:`2017-01-01` วันที่เริ่มต้น
	DateEnd   string `json:"date_end"`   // example:`2017-12-31` วันที่สิ้นสุด
	Boundary  string `json:"boundary"`   // boundary จาก db เก่า
}

type Struct_AdvRainSum struct {
	X     float64 `json:"x"`     // St_x
	Y     float64 `json:"y"`     // St_y
	Sum   float64 `json:"sum"`   // example:`10.3`  ค่าฝนรวม
	Count float64 `json:"count"` // example:`310` จำนวนค่าฝนที่คำนวณ
}

type Struct_RainfallMaxMin struct {
	Current_time      string `json:"current_time"`      // example:`09:30` เวลาของฝนตกมากที่สุด
	Station_name_max  string `json:"station_name_max"`  // exmaple:`สถานีบ้านด่านดู่ จ.เลย` สถานีฝนตกมากที่สุด
	Rainfall_max      string `json:"rainfall_max"`      // example:`23.50` ค่าฝนตกมากที่สุด
	Current_datetime  string `json:"current_datetime"`  // example:`2018-01-19 09:30` วันเวลาของฝนตกมากที่สุด
	Rainfall_max_text string `json:"rainfall_max_text"` // example:`ฝนตกปานกลาง` สถานการณ์ของฝนตกมากที่สุด
	Station_name_min  string `json:"station_name_min"`  // example:`สถานีนาซ่าว จ.เลย` สถานีฝนตกน้อยที่สุด
	Rainfall_min      string `json:"rainfall_min"`      // example:`0.20` ค่าฝนตกน้อยที่สุด
	Rainfall_min_text string `json:"rainfall_min_text"` // example:`ฝนตกเล็กน้อย` สถานการณ์ของฝนตกน้อยที่สุด
}

//	initial default value
func (s *Struct_RainfallMaxMin) Init() {
	s.Current_time = "-"
	s.Station_name_max = "-"
	s.Rainfall_max = "-"
	s.Current_datetime = "-"
	s.Rainfall_max_text = "-"
	s.Station_name_min = "-"
	s.Rainfall_min = "-"
	s.Rainfall_min_text = "-"
}
