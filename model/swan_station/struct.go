package swan_station

import (
	"encoding/json"
)

type Struct_SwanStation struct {
	Id                int64           `json:"id"`                          // example:`1` รหัสสถานี
	Agency_Id         int64           `json:"agency_id,omitempty"`         // example:`9` รหัสหน่วยงาน
	Geocode_Id        int64           `json:"geocode_id,omitempty"`        // รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Name              json.RawMessage `json:"name,omitempty"`              // ชื่อสถานี
	Lat               interface{}     `json:"lat"`                         // ละติจูด
	Long              interface{}     `json:"long"`                        // ลองติจูด
	Oldcode           string          `json:"oldcode,omitempty"`           // รหัสสถานีเดิม
	Dataimport_Log_Id int64           `json:"dataimport_log_id,omitempty"` // รหัสอิมพอร์ตข้อมูล

	Province_Code string          `json:"province_code,omitempty"` // รหัสจังหวัด
	Province_Name json.RawMessage `json:"province_name,omitempty"` // ชื่อจังหวัด
}
