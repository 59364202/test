package ignore

import (
	//model_user "haii.or.th/api/server/model/user"
	"encoding/json"
)

//	Struct of input parameters
type Struct_IgnoreStation_InputParam struct {
	ID        string `json:"data_id"`    // example:`69137` รหัส data
	IsIgnore  bool   `json:"is_ignore"`  // exampel:`true` สถานะ ignore
	StationID string `json:"station_id"` // example:`37` รหัสสถานี
	TableName string `json:"table_name"` // example:`dam_daily` ชื่อชนิดข้อมูล
	StartDate string `json:"date_start"`
	EndDate   string `json:"date_end"`
}

//Struct of ignore parameters
type Struct_IgnoreStation struct {
	StationID       int64           `json:"station_id"`                 // รหัสสถานี
	StationName     json.RawMessage `json:"station_name,omitempty"`     // ชื่อของสถานี
	ProvinceName    json.RawMessage `json:"province_name,omitempty"`    // ชื่อจังหวัดของสถานีที่ติดตั้ง
	AgencyShortName json.RawMessage `json:"agency_shortname,omitempty"` // ชื่อย่อหน่วยงานของสถานีที่ติดตั้ง
	Datetime        string          `json:"data_datetime,omitempty"`    // วันเวลาของข้อมูล
}

type Struct_IgnoreDescription struct {
	Description string `json:"description"` // Description ของ table
}
