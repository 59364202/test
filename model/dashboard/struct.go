package dashboard

import (
	"encoding/json"

	"haii.or.th/api/server/model/srvstatus"
	"haii.or.th/api/thaiwater30/util/result"
)

type Agency struct {
	ID              interface{}     `json:"id"`               // รหัสหน่วยงาน
	AgencyShortName json.RawMessage `json:"agency_shortname"` // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
	AgencyName      json.RawMessage `json:"agency_name"`      // ชื่อหน่วยงาน
}

type DataimportOutputData struct {
	MetadataSName   json.RawMessage `json:"metadataservice_name"`      // ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	AgencyID        interface{}     `json:"agency_id"`                 // รหัสของหน่วยงาน
	MetadataID      interface{}     `json:"metadata_id"`               // รหัสบัญชีข้อมูล
	DataFrequency   string          `json:"metadata_convertfrequency"` // ความถี่ของการนำเข้าข้อมูล
	AgencyShortName json.RawMessage `json:"agency_shortname"`          // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
	AgencyName      json.RawMessage `json:"agency_name"`               // รายชื่อเต็มหน่วยงาน
	LastUpdate      string          `json:"import_begin_at"`           // วันที่นำเข้าข้อมูลล่าสุด
	UpdatePlan      int64           `json:"metadata_update_plan"`      // ระยะเวลา(นาที) ที่ต้องถึงเวลลาอัพเดท
	OverMinute      int64           `json:"overdue_minute"`            // ระยะเวลา(นาที) นับจากวันที่ อัพเดทล่าสุด
}

type MedadataAlertOutput struct {
	DataimportOutputData *result.Result          `json:"data_import"`   // สรุปเหตุการณ์นำเข้าข้อมูลที่เกิดขึ้นในระบบ 24 ชั่วโมงย้อนหลัง
	Agency               *result.Result          `json:"agency_list"`   // สรุปรายการหน่วยงาน
	DataService          *result.Result          `json:"data_service"`  //สรุปรายการคำขอที่ค้างอยู่
	IgnoreCount          *result.Result          `json:"data_ignore"`   //จำนวนรายการที่ถูก ignore
	CountUser            *result.Result          `json:"count_user"`    //จำนวนผู้ใช้งานที่ active เป็น false
	ServerStatus         *srvstatus.ServerStatus `json:"server_status"` // server status
}

type DataService struct {
	NoResult int64 `json:"count_no_result"` // จำนวนรายการที่ยังค้างอยู่
}

type CountUser struct {
	CountUser int64 `json:"count_no_user"` // จำนวนรายการผู้ใช้ที่ที่ยังค้างอยู่
}

type DataIgnore struct {
	Rain  int64 `json:"ignore_count_rain"`  // จำนวนสถานีที่ถูก ignore (ฝน)
	Water int64 `json:"ignore_count_water"` // จำนวนสถานีที่ถูก ignore (น้ำ)
}

type FrequencyUnit2 struct {
	FrequencyName string `json:"datafrequency"` // ความถี่
}

type MetadataAlertStatus struct {
	MetadataID          int64           `json:"metadata_id"`          // รหัสบัญชีข้อมูล
	MetadataServiceName json.RawMessage `json:"metadataservice_name"` // ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	Agency              *Agency         `json:"agency"`               // หน่วยงาน
	LastUpdate          string          `json:"last_update"`          // วันที่อัพเดทล่าสุด
	OverMinute          int64           `json:"overdue_minute"`       // ระยะเวลา(นาที) นับจากวันที่ อัพเดทล่าสุด
	FrequencyUnitData   *FrequencyUnit2 `json:"frequency_unit"`       // ความถี่
}
