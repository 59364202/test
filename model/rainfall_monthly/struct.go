package rainfall_monthly

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
)

type Struct_RainfallMonthly struct {
	//	StationId    int64           `json:"station_id,omitempty"`
	//	StationName  json.RawMessage `json:"station_name,omitempty"`
	//	DateTime     string          `json:"date_time,omitempty"`
	//	Rainfall     float64         `json:"rainfall"`
	//	ProvinceCode string          `json:"province_code,omitempty"`
	//	ProvinceName json.RawMessage `json:"province_name,omitempty"`
	//	StationLat   string          `json:"station_lat,omitempty"`
	//	StationLong  string          `json:"station_long,omitempty"`
	//	AmphoeName   json.RawMessage `json:"amphoe_name,omitempty"`
	//	TumbonName   json.RawMessage `json:"tumbon_name,omitempty"`
	//	AgencyName   json.RawMessage `json:"agency_name,omitempty"`

	Id            int64       `json:"id"`                // example:`120732` รหัสข้อมูลปริมาณน้ำฝน
	RainfallValue interface{} `json:"rainfall_value"`    // example:`1.5` ค่าฝน (mm.)
	Time          string      `json:"rainfall_datetime"` // example:`2006-01-02 15:04` วันเวลาที่เก็บปริมาณน้ำฝน
	StationType   string      `json:"station_type"`      // example:`rainfall_monthly` ชนิดของสถานี

	Agency  *model_agency.Struct_Agency   `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"` // สถานี\
	Basin   *model_basin.Struct_Basin   `json:"basin,omitempty"` // ลุ่มน้ำ
}

type Struct_TeleStation struct {
	Id      int64           `json:"id"`                          // example:`1051` รหัสสถานี
	Name    json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"บ้านคลองพร้าว"}` ชื่อสถานี
	Lat     interface{}     `json:"tele_station_lat"`            // example:`12.050653` ละติจูดของสถานี
	Long    interface{}     `json:"tele_station_long"`           // example:`102.296239` ลองติจูดของสถานี
	OldCode string          `json:"tele_station_oldcode"`        // example:`STN0787` รหัสสถานีเดิม
	Type    string          `json:"tele_station_type"`           // example:`rainfall_monthly` ชนิดของสถานี
	SubBasin_id 	int64       `json:"sub_basin_id,omitempty"`      // example:`21` รหัสลุ่มน้ำย่อย
}

type Struct_RainfallMonthly_Graph struct {
	DateTime string      `json:"date_time"` // example:`2006-01-02 15:04` วันเวลาที่เก็บปริมาณน้ำฝน
	Rainfall interface{} `json:"rainfall"`  // example:`1.5` ค่าฝน (mm.)
	DayCount interface{} `json:"day_count"` // example:`30` จำนวนวัน
}

type Param_RainfallMonthly struct {
	IsYearly  bool   // เป็นฝนรายปี
	Order     string `json:"order"`      // order by
	DataLimit int    `json:"data_limit"` // จำนวนข้อมูล
	StationId int64  `json:"station_id"` // รหัสสถานี
	StratDate string `json:"start_date"` // วันที่เริ่มต้น
	EndDate   string `json:"end_date"`   // วันที่สิ้นสุด
	Year      int    `json:"year"`       // ปี
}
