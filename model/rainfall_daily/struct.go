package rainfall_daily

import (
	"encoding/json"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
)

type Struct_RainfallDaily struct {
	StationId    int64           `json:"tele_station_id,omitempty"`   // example:`1051` รหัสสถานี
	StationName  json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"บ้านคลองพร้าว"}` ชื่อสถานี
	DateTime     string          `json:"rainfall_datetime"`           // example:`2006-01-02 15:04` วันที่เก็บปริมาณน้ำฝน
	Rainfall     interface{}     `json:"rainfall_value"`              // example:`1.5` ค่าฝน (mm.)
	ProvinceCode string          `json:"province_code,omitempty"`     // example:`10` รหัสจังหวัด
	ProvinceName json.RawMessage `json:"province_name,omitempty"`     // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัด
	StationLat   interface{}     `json:"tele_station_lat,omitempty"`  // example:`12.050653` ละติจูดของสถานี
	StationLong  interface{}     `json:"tele_station_long,omitempty"` // example:`102.296239` ลองจิจูดของสถานี
	AmphoeName   json.RawMessage `json:"amphoe_name,omitempty"`       // example:`{"th": "พระนคร"}` ชื่ออำเภอ
	TumbonName   json.RawMessage `json:"tumbon_name,omitempty"`       // example:`{"th": "พระบรมมหาราชวัง"}` ชื่อตำบล
	AgencyName   json.RawMessage `json:"agency_name,omitempty"`       // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงาน
	WarningZone  string          `json:"warning_zone"`                // example:`06` โซนเตือนภัยพิจารณาตามพื้นที่
	Basin   *model_basin.Struct_Basin   `json:"basin,omitempty"` // ลุ่มน้ำ
	SubBasin_id 	int64       `json:"sub_basin_id,omitempty"`      // example:`21` รหัสลุ่มน้ำย่อย
}

type Struct_RainfallDaily_Graph struct {
	DateTime string      `json:"rainfall_datetime"` // example:`2006-01-02 15:04` วันที่เก็บปริมาณน้ำฝน
	Rainfall interface{} `json:"rainfall_value"`    // example:`3.5` ค่าฝน (mm.)
}

type Param_RainfallDaily struct {
	IsDaily   bool   // เป็นฝนรายวัน
	IsMonthly bool   // เป็นฝนรายเดือน
	Order     string `json:"order"`      // order by
	DataLimit int    `json:"data_limit"` // จำนวนข้อมูล
	StationId int64  `json:"station_id"` // รหัสสถานี
	StratDate string `json:"start_date"` // วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // วันที่สิ้นสุด
	Year      int    `json:"year"`       // ปี
	Month     int    `json:"month"`      // เดือน
}
