package tele_station

import (
	"encoding/json"
)

type Struct_TeleStation struct {
	Id                   int64           `json:"id"`                             // example:`3516` รหัสสถานี
	Subbasin_id          int64           `json:"subbasin_id,omitempty"`          // example:`255` รหัสลุ่มน้ำสาขา
	Tele_station_name    json.RawMessage `json:"tele_station_name,omitempty"`    // example:`{"th": "อ.ตรอน (N.60)"}` ชื่อสถานี
	Tele_station_lat     interface{}     `json:"tele_station_lat,omitempty"`     // example:`17.41439` ละติจูด
	Tele_station_long    interface{}     `json:"tele_station_long,omitempty"`    // example:`100.1238` ลองติจูด
	Tele_station_oldcode string          `json:"tele_station_oldcode,omitempty"` // example:`skTD04` รหัสสถานีเดิม
	Tele_station_type    string          `json:"tele_station_type,omitempty"`    // example:`tele_waterlevel` ชนิดของสถานี
	Left_bank            interface{}     `json:"left_bank"`                      // example:`58.45` ระดับตลิ่ง (ซ้าย)
	Right_bank           interface{}     `json:"right_bank"`                     // example:`56.55` ระดับตลิ่ง (ขวา)
	Min_bank             interface{}     `json:"min_bank,omitempty"`             // example:`56.55` ระดับตลิ่งต่ำสุด
	Ground_level         interface{}     `json:"ground_level,omitempty"`         // example:`47.46` ระดับท้องน้ำ ม.รทก
	SubBasin_id          int64           `json:"sub_basin_id,omitempty"`         // example:`21` รหัสลุ่มน้ำย่อย

	Agency_id  int64 `json:"agency_id,omitempty"`  // example:`8` รหัสหน่วยงาน
	Geocode_id int64 `json:"geocode_id,omitempty"` // example:`5115` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
}

type Struct_Station struct {
	Station_id       int             `json:"station_id"`              // example:`3516` รหัสสถานี
	Station_name     json.RawMessage `json:"station_name,omitempty"`  // example:`{"th": "อ.ตรอน (N.60)"}` ชื่อสถานี
	Station_old_code string          `json:"station_old_code"`        // example:`skTD04` รหัสสถานีเดิม
	Station_lat      interface{}     `json:"station_lat,omitempty"`   // example:`17.41439` ละติจูด
	Station_long     interface{}     `json:"station_long,omitempty"`  // example:`100.1238` ลองติจูด
	Geocode          string          `json:"geocode,omitempty"`       // example:`100101` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Province_name    json.RawMessage `json:"province_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัด
	Province_code    string          `json:"province_code,omitempty"` // example:`10` รหัสจังหวัด
	Agency_id        int64           `json:"agency_id,omitempty"`     // example:`9` รหัสหน่วยงาน
	Station_type     string          `json:"station_type,omitempty"`  // ชนิดของสถานี
}

type TeleStationStruct struct {
	Station_id       int             `json:"station_id"`                 // example:`3516` รหัสสถานี
	Station_name     json.RawMessage `json:"station_name,omitempty"`     // example:`{"th": "อ.ตรอน (N.60)"}` ชื่อสถานี
	Station_old_code string          `json:"station_old_code"`           // example:`skTD04` รหัสสถานีเดิม
	Station_lat      interface{}     `json:"station_lat,omitempty"`      // example:`17.41439` ละติจูด
	Station_long     interface{}     `json:"station_long,omitempty"`     // example:`100.1238` ลองติจูด
	Province_id      int             `json:"province_id,omitempty"`      // example:`10` รหัสจังหวัด
	Geocode          string          `json:"geocode,omitempty"`          // example:`5115` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Tumbon_name      json.RawMessage `json:"tumbon_name,omitempty"`      // example:`{"th": "พระบรมมหาราชวัง"}` ชื่อตำบล
	Amphoe_name      json.RawMessage `json:"amphoe_name,omitempty"`      // example:`{"th": "พระนคร"}` ชื่ออำเภอ
	Province_name    json.RawMessage `json:"province_name,omitempty"`    // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัด
	Province_code    string          `json:"province_code,omitempty"`    // example:`10` รหัสจังหวัด
	Station_oldcode  string          `json:"station_oldcode,omitempty"`  // example:`skTD04` รหัสสถานีเดิม
	Agency_id        int64           `json:"agency_id,omitempty"`        // example:`9` รหัสหน่วยงาน
	Agency_name      json.RawMessage `json:"agency_name,omitempty"`      // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงาน
	Agency_shortname json.RawMessage `json:"agency_shortname,omitempty"` // example:`{"en": "HAII"}` ชื่อย่อหน่วยงาน
	Station_type     string          `json:"station_type,omitempty"`     // ชนิดของสถานี
	Geocode_basin    string          `json:"geocode_basin,omitempty"`    //
}

type Struct_TeleStationGroupByProvince struct {
	ProvinceCode interface{}          `json:"province_code"`           // example:`10` รหัสจังหวัด
	ProvinceName json.RawMessage      `json:"province_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัด
	TeleStation  []*TeleStationStruct `json:"station"`                 // สถานี
}

type TeleStationParam struct {
	DataType   string `json:"data_type"`
	AgencyId   string `json:"agency_id"`
	ProvinceId string `json:"province_id"`
}

type GetTeleStation_OutputParam struct {
	Id               int64           `json:"id"`               // รหัสสถานี
	Station_name     json.RawMessage `json:"station_name"`     // ชื่อสถานี
	Station_oldcode  string          `json:"station_oldcode"`  // รหัสสถานีเดิม
	Agency_name      json.RawMessage `json:"agency_name"`      // ชื่อหน่วยงาน
	Agency_shortname json.RawMessage `json:"agency_shortname"` // ชื่อย่อหน่วยงาน
}

type Struct_WaterlevelCanalStationGroupByProvince struct {
	ProvinceCode interface{}                      `json:"province_code"`           // example:`01`รหัสจังหวัด
	ProvinceName json.RawMessage                  `json:"province_name,omitempty"` // example:`{"th":"กรุงเทพมหานคร"}`ชื่อจังหวัด
	TeleStation  []*Struct_WaterlevelCanalStation `json:"station"`                 // สถานี
}

type Struct_WaterlevelCanalStation struct {
	ID      string          `json:"station_id"`             // example:`56`รหัสสถานี
	OldCode string          `json:"station_oldcode"`        // example:`BTG093`รหัสสถานีเดิม
	Name    json.RawMessage `json:"station_name,omitempty"` // example:`"th":"สถานีวัดระดับน้ำกรุงเทพฯ"`ชื่อสถานี
	Lat     interface{}     `json:"station_lat"`            // example:`9.451456`ละติจูด
	Long    interface{}     `json:"station_long"`           // example:`103.456785`ลองติจูด
}
