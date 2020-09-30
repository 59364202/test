package waterquality_station

import (
	"encoding/json"
)

type Struct_WaterQualityStation struct {
	Id                           int64           `json:"id"`                                     // example:`109` รหัสถานี
	Agency_id                    int64           `json:"agency_id,omitempty"`                    // example:`14` รหัสหน่วยงาน
	Agency_Name                  json.RawMessage `json:"agency_name,omitempty"`                  // example:`{"th": "กรมควบคุมมลพิษ ", "en": "Pollution Control Department"}` ชื่อหน่วยงาน
	Agency_Shortname             json.RawMessage `json:"agency_shortname,omitempty"`             // example:`{"en": "PCD"}` ชื่อย่อหน่วยงาน
	Geocode_id                   int64           `json:"geocode_id,omitempty"`                   // example:`8372` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Waterquality_Station_Name    json.RawMessage `json:"waterquality_station_name,omitempty"`    // example:`{"th": "สถานีบ้านห้วยซัน"}` ชื่อสถานี
	Waterquality_Station_Lat     interface{}     `json:"waterquality_station_lat,omitempty"`     // example:`16.518361` ละติจูด
	Waterquality_Station_Long    interface{}     `json:"waterquality_station_long,omitempty"`    // example:`102.898361` ลองติจูด
	Waterquality_Station_Oldcode string          `json:"waterquality_station_oldcode,omitempty"` // example:`212` รหัสสถานีเดิม
	Is_active                    string          `json:"is_active,omitempty"`
	SubBasin_id                  int64           `json:"sub_basin_id,omitempty"` // example:`21` รหัสลุ่มน้ำย่อย

	Province_Name json.RawMessage `json:"province_name,omitempty"` // example:`{"th": "นราธิวาส"}` ชื่อจังหวัด
	Amphoe_Name   json.RawMessage `json:"amphoe_name,omitempty"`   // example:`{"th": "รือเสาะ"}` ชื่ออำเภอ
	Tumbon_Name   json.RawMessage `json:"tumbon_name,omitempty"`   // example:`{"th": "สาวอ"}` ชื่อตำบล
	Province_Code string          `json:"province_code,omitempty"` // example:`96` รหัสจังหวัด
}

type Struct_Station struct {
	Id                           int64           `json:"id"`                                     // example:`109` รหัสถานี
	Agency_id                    int64           `json:"agency_id,omitempty"`                    // example:`14` รหัสหน่วยงาน
	Geocode_id                   int64           `json:"geocode_id,omitempty"`                   // example:`8372` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Waterquality_Station_Name    json.RawMessage `json:"waterquality_station_name,omitempty"`    // example:`{"th": "สถานีบ้านห้วยซัน"}` ชื่อสถานี
	Waterquality_Station_Lat     interface{}     `json:"waterquality_station_lat,omitempty"`     // example:`16.518361` ละติจูด
	Waterquality_Station_Long    interface{}     `json:"waterquality_station_long,omitempty"`    // example:`102.898361` ลองติจูด
	Waterquality_Station_Oldcode string          `json:"waterquality_station_oldcode,omitempty"` // example:`212` รหัสสถานีเดิม
	Is_active                    string          `json:"is_active,omitempty"`
	SubBasin_id                  int64           `json:"sub_basin_id,omitempty"` // example:`21` รหัสลุ่มน้ำย่อย
}
