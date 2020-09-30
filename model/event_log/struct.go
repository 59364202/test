package event_log

import (
	"encoding/json"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_event_code "haii.or.th/api/thaiwater30/model/event_code"
	model_event_log_category "haii.or.th/api/thaiwater30/model/event_log_category"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
)

type Struct_EventLog struct {
	Id                   int64           `json:"id"`                       // example:`642381` รหัส event log
	EventLogDate         string          `json:"event_log_date"`           // example:`2006-01-02 15:04` วันเวลา
	EventLogData         json.RawMessage `json:"event_log_data,omitempty"` // example:`{"dataset_id":308,"dataset_log_id":149538,"params":{"data_path":"dataset/aviso/media/image/ssh/global/2017/08/15/000000","duration":9639308,"config_name":"aviso-webextract-aviso","success_row":0,"error_row":1,"event_code_id":1},"node_ip":"192.168.12.189"}` ข้อมูล event log
	EventLogMessage      string          `json:"event_log_message"`        // exampel:`Partial Success` ข้อความแสดงเหตุการณ์
	EventLogDuration     int64           `json:"event_log_duration"`       // example:`9639392` ระยะเวลาการรัน script
	EventLogDurationTime int64           `json:"event_log_duration_time"`  // example:`9639392` ระยะเวลาการรัน script

	Metadata      *model_metadata.Struct_Metadata                   `json:"metadata"`       // บัญขีข้อมูล
	Agency        *model_agency.Struct_Agency                       `json:"agency"`         // หน่วยงาน
	EventCategory *model_event_log_category.Struct_EventLogCategory `json:"event_category"` // เหตุการณ์
	EventCode     *model_event_code.Struct_EventCode                `json:"event_code"`     // เหตุการณ์ย่อย
}

//Output Data => Used by BOF-DataIntegrationReport-EventLog Screen
type Struct_EventLogSummary_GroupByAgencyCategory struct {
	Agency              *model_agency.Struct_Agency                         `json:"agency"`             // หน่วยงาน
	ListOfEventCategory []*model_event_log_category.Struct_EventLogCategory `json:"event_log_category"` // เหตุการณ์ที่เกิดขึ้น
}

//Input Data => Used by BOF-DataIntegrationReport-Event Screen
type Struct_EventLog_InputParam struct {
	AgencyID        string `json:"agency_id"`             // example:`14` รหัสหน่วยงาน
	EventCategoryID string `json:"event_log_category_id"` // example:`5` รหัสเหตุการณ์
	EventCodeID     string `json:"event_code_id"`         // example:`1` รหัสเหตุการณ์ย่อย
	StartDate       string `json:"start_date"`            // example:`2017-08-01` วันที่เริ่มต้น
	EndDate         string `json:"end_date"`              // example:`2017-08-01` วันที่สิ้นสุด
}

type Struct_EventLogSummaryReport struct {
	ID            int64                            `json:"id"`                    // example:`2` รหัสเหตุการณ์
	Code          string                           `json:"code"`                  // example:`ERROR` โค้ดเหตุการณ์
	Description   json.RawMessage                  `json:"description,omitempty"` // example:`{"th": " Error conditions", "en": "Error conditions"}` คำอธิบายประเภทเหตุการณ์
	PercentEvent  float64                          `json:"percent_event"`         // example:`1.6355140186915889` คิดเป็น % จากทั้งหมด
	ListEventCode []*Struct_EventCodeSummaryReport `json:"event_code"`            // เหตุการณ์ย่อย
}

type Struct_EventCodeSummaryReport struct {
	ID          int64                         `json:"id"`                    // example:`19` รหัสเหตุการณ์ย่อย
	Code        string                        `json:"code"`                  // example:`EventDataNoticePartialSuccess` โค้ดเหตุการณ์ย่อย
	Description json.RawMessage               `json:"description,omitempty"` // example:`{"en": "This record contains invalid data"}` คำอธิบายเหตุการณ์ย่อย
	ListAgency  []*model_agency.Struct_Agency `json:"agency"`                // หน่วยงาน
}

type Struct_EventReport_Input struct {
	Date  string `json:"date"`  // exmaple:`2018-06-13` วันที่ของเหตุการณ์
	Agent string `json:"agent"` // example:`dataimport-haii` หน่วยงาน
	Event int    `json:"event"` // example:`3` รหัสเหตุการณ์
}
type Struct_EventReport struct {
	Type     string `json:"type"`     // example:`download` ขั้นตอนการทำงาน
	Id       string `json:"id"`       // example:`586` รหัส
	LogId    string `json:"log_id"`   // example:`1749940` log id
	Name     string `json:"name"`     // example:`hd-waterlevel-ws`
	Detail   string `json:"detail"`   // example:`` รายละเอียด
	DateTime string `json:"datetime"` // example:`2018-06-13 16:15` เวลาเริ่มทำงาน
}
