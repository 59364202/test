package station

import (
	"encoding/json"
)

type Struct_Station_ForCheckMetadata struct {
	Id              int64           `json:"id"`                         // example:`5` รหัสสถานี
	Name            json.RawMessage `json:"station_name"`               // example:`{"en":"Nonsi Witthaya School","th":"โรงเรียนนนทรีวิทยา"}` ชื่อสถานี
	OldCode         string          `json:"station_oldcode"`            // example:`12t` รหัสสถานีเดิม
	Lat             interface{}     `json:"lat"`                        // example:`13.708038` ละติจูด
	Long            interface{}     `json:"long"`                       // example:`100.547345` ลองติจูด
	Geocode         interface{}     `json:"geocode,omitempty"`          // example:`null` ข้อมูลขอบเขตการปกครองของประเทศไทย
	SubbasinName    json.RawMessage `json:"subbasin_name,omitempty"`    // example:`{"th":"น้ำของ","en":"NAM KHONG"}` ลุ่มน้ำสาขา
	AgencyShortname json.RawMessage `json:"agency_shortname,omitempty"` // example:`{"en":"PCD"}` ชื่อย่อหน่วยงาน
	AgencyName      json.RawMessage `json:"agency_name,omitempty"`      // example:`{"th":"กรมควบคุมมลพิษ ","en":"Pollution Control Department","jp":""}` ชื่อหน่วยงาน
}

type Struct_Station_InputParam struct {
	ColumnName string `json:"column_name"` // example:`geocode_id` ชื่อ field
	AgencyID   string `json:"agency_id"`   // example:`14` รหัสหน่วยงาน
	TableName  string `json:"table_name"`  // example:`m_air_station` ชื่อตาราง
}
