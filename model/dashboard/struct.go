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
	MetadataSName   json.RawMessage `json:"metadataservice_name"` // ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	AgencyID        interface{}     `json:"agency_id"`
	MetadataID      interface{}     `json:"metadata_id"`
	DataFrequency   interface{}     `json:"datafrequency"`
	AgencyShortName json.RawMessage `json:"agency_shortname"` // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
	AgencyName      json.RawMessage `json:"agency_name"`
	BeginDate       interface{}     `json:"import_begin_at"`
}

type MedadataAlertOutput struct {
	EventData    *result.Result          `json:"event_data"`    // สรุปเหตุการณ์นำเข้าข้อมูลที่เกิดขึ้นในระบบ 24 ชั่วโมงย้อนหลัง
	AlertOnline  *result.Result          `json:"alert_online"`  // รายการบัญชีข้อมูล Online
	AlertOffline *result.Result          `json:"alert_offline"` // รายการบัญชีข้อมูล Offline
	LastLogin    *result.Result          `json:"last_login"`    // Last login
	UserOnline   *result.Result          `json:"user_online"`   // user online
	Dataimport   *result.Result          `json:"dataimport"`    // Activity log (latest)
	Dataservice  *result.Result          `json:"dataservice"`   // การให้บริการข้อมูล
	ServerStatus *srvstatus.ServerStatus `json:"server_status"` // server status
}

type DataService struct {
	NoResult              int64 `json:"count_no_result"` // จำนวนรายการที่ยังค้างอยู่
	EnableWithoutDownload int64 `json:"count_service"`   // จำนวนรายการ ที่ไม่รวม download
	DownloadOnlyEnable    int64 `json:"count_download"`  // จำนวนรายการ ในส่วน download ที่ยังใช้การได้
}
