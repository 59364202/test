package lt_geocode

import (
	"encoding/json"
)

type Struct_Geocode_Id struct {
	Id          	int64  `json:"id,omitempty"`        // example:`3` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Geocode     string `json:"geocode,omitempty"`      // example:`100101` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Province_name	string `json:"province_name"`		// example:`กรุงเทพมหานคร` ชื่อจังหวัด
	Amphoe_name		string `json:"amphoe_name"`			// example:`พระนคร` ชื่ออำเภอ
	Tumbon_name		string `json:"tumbon_name"`			// example:`พระบรมมหาราชวัง` ชื่อตำบล
}

type Struct_Geocode struct {
	Id          int64  `json:"id,omitempty"`           // example:`3` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Geocode     string `json:"geocode,omitempty"`      // example:`100101` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	WarningZone string `json:"warning_zone,omitempty"` // example:`06` โซนเตือนภัยพิจารณาตามพื้นที่

	Struct_Area
	Struct_Amphoe
	Struct_Tumbon
	Struct_Province
}

type Struct_Area struct {
	Area_code string          `json:"area_code,omitempty"` // example:`1` รหัสภาคของประเทศไทย
	Area_name json.RawMessage `json:"area_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อภาคของประเทศไทย
}

type Struct_Region struct {
	Area_code string          `json:"region_code"`           // example:`1` รหัสภาคของประเทศไทย
	Area_name json.RawMessage `json:"region_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อภาคของประเทศไทย
}

type Struct_Amphoe struct {
	Amphoe_code string          `json:"amphoe_code,omitempty"` // example:`01` รหัสอำเภอของประเทศไทย
	Amphoe_name json.RawMessage `json:"amphoe_name,omitempty"` // example:`{"th": "พระนคร"}` ชื่ออำเภอของประเทศไทย
}
type Struct_Tumbon struct {
	Tumbon_code string          `json:"tumbon_code,omitempty"` // example:`01` รหัสตำบลของประเทศไทย
	Tumbon_name json.RawMessage `json:"tumbon_name,omitempty"` // example:`{"th": "พระบรมมหาราชวัง"}` ชื่อตำบลของประเทศไทย
}
type Struct_Province struct {
	Province_code string          `json:"province_code,omitempty"` // example:`10` รหัสจังหวัดของประเทศไทย
	Province_name json.RawMessage `json:"province_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
}

type Struct_Geocode_P struct {
	Id      int64  `json:"id,omitempty"`      // example:`3` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Geocode string `json:"geocode,omitempty"` // example:`100101` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Struct_Province
}

type Province struct {
	ProvinceCode string          `json:"province_code"` // example:`10` รหัสจังหวัดของประเทศไทย
	ProvinceName json.RawMessage `json:"province_name"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	TeleStation  []*Station      `json:"tele_station"`  // สถานี
}

type Station struct {
	StationID   int64           `json:"tele_station_id"`      // example:`2212` รหัสสถานีโทรมาตร
	StationCode string          `json:"tele_station_oldcode"` // example:`566201` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน
	StationName json.RawMessage `json:"tele_station_name"`    // example:`{"th":"เกาะลันตา"}` ชื่อสถานีโทรมาตร
	MaxYear     interface{}     `json:"max_year"`             // example:`2006` ปีที่มากที่สุดของข้อมูลล
	MinYear     interface{}     `json:"min_year"`             // example:`2006` ปีที่น้อยที่สุดของข้อมูลล
}
