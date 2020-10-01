package ignore_history

import (
	"encoding/json"
)

//Struct of ignore_history table
type Struct_IgnoreHistory struct {
	ID              int64           `json:"id"`               // example:`44` รหัสประวัติของการ ignore ข้อมูล
	IgnoreDatetime  string          `json:"ignore_datetime"`  // example:`2006-01-02 15:04` วันเวลาที่ Ignore ข้อมูล
	DataCategory    string          `json:"data_category"`    // example:`rainfall_24h` ประเภขข้อมูล คือ ฝน เขื่อน ระดับน้ำ และคุณภาพน้ำ
	StationID       int64           `json:"station_id"`       // example:`308` รหัสสถานี
	StationOldCode  string          `json:"station_oldcode"`  // example:`WEI012` รหัสสถานีเดิม
	StationName     json.RawMessage `json:"station_name"`     // example:`{"th": "ท้ายฝายยางกะฮาด"}` ชื่อของสถานี
	StationProvince json.RawMessage `json:"station_province"` // example:`{"th": "ชัยภูมิ"}` ชื่อจังหวัดของสถานีที่ติดตั้ง
	//AgencyName      json.RawMessage `json:"agency_name"`
	AgencyShortname json.RawMessage `json:"agency_shortname"` // example:`{"th": "", "en": "HAII"}` ชื่อย่อหน่วยงานของสถานีที่ติดตั้ง
	DataID          int64           `json:"data_id"`          // example:`676639` รหัสของข้อมูลที่ ignore
	DataDate        string          `json:"data_datetime"`    // example:`2006-01-02 15:04` วันวลาของข้อมูล
	Remark          string          `json:"remark"`           // example:`Unlock Ignore` หมายเหตุ
	DataValue       interface{}     `json:"value"`            // example:`274.600006` ค่าของข้อมูล

	User *Struct_User `json:"user"` // ผู้เปลี่ยนสถานะ ignore
}

//Struct of input parameters
type Struct_IgnoreHistory_InputParam struct {
	ID               string `json:"id"`              // รหัสข้อมูล
	StationID        string `json:"station_id"`      // รหัสสถานี
	TableName        string `json:"table_name"`      // enum:[rainfall_24h,dam_daily,dam_hourly,tele_waterlevel] example:`tele_waterlevel` ชื่อชนิดข้อมูล
	IsIgnore         bool   `json:"is_ignore"`       // สถานะ ignore
	SQLCmdIgnoreInfo string `json:"sql_ignore_info"` //
}

type Struct_User struct {
	ID       int64  // example:`68` รหัสผู้ใช้
	FullName string // example:`Kantamat Polsawang` ชื่อผู้ใช้
}
