package event_tracking

import (
	"encoding/json"
)

type EventTrackingSelectOptionList struct {
	Category []*EventTrackingSelectOptionCategory `json:"event_type"`     // รหัสหมวดหลัก
	Code     []*EventTrackingSelectOptionCode     `json:"event_sub_type"` // รหัสหมวดย่อย
	Agency   []*EventTrackingSelectOptionAgency   `json:"agency"`         // รหัสหน่วยงาน
}

type EventTrackingSelectOptionInvalidData struct {
	Value interface{}     `json:"value"`          // example:`1` รหัสหน่วยงาน ที่เกิดปัญหา
	Text  json.RawMessage `json:"text,omitempty"` // example:`{"th":"สสนก."}` ชื่อหน่วยงาน
	Date  []string        `json:"date"`           // example:`["2006-01-02"]`
}

type EventTrackingSelectOptionCode struct {
	Value    interface{}     `json:"value"`          // example:`1` รหัสหมวดย่อย
	Text     json.RawMessage `json:"text,omitempty"` // example:`{"en":"Invalid Data"}` ชื่อหมวดย่อย
	Category int64           `json:"event_type"`     // example:`2` รหัส category
}

type EventTrackingSelectOptionAgency struct {
	Value interface{}     `json:"value"`          // example:`9` รหัสหน่วยงาน
	Text  json.RawMessage `json:"text,omitempty"` // example:`{"th":"สสนก.","en":"Hydro and Agro Informatics Institute","jp":""}` ชื่อหน่วยงาน
}

type EventTrackingSelectOptionCategory struct {
	Value interface{} `json:"value"` // example:`2` รหัสหมวดหลัก
	//	Text  json.RawMessage               `json:"text,omitempty"`
	Text interface{}                      `json:"text"`           // example:`{"en":"Critical"}` ชื่อหมวดหลัก
	Code []*EventTrackingSelectOptionCode `json:"event_sub_type"` // รายชื่อหมวดย่อย
}

type EventTrackingData struct {
	EventLogId          int64             `json:"event_log_id"`             // example:`2324` ลำดับเหตุการณ์
	EventCodeSubType    json.RawMessage   `json:"event_sub_type,omitempty"` // example:`{"th":"Invalid Data"}` ชื่อหมวดย่อย
	EventCode           string            `json:"event_code"`               // example:`1` หมวดย่อย
	EventDate           string            `json:"event_date"`               // example:`2006-01-02` วันที่เกิดเหตุการณ์
	Agency              json.RawMessage   `json:"agency_name,omitempty"`    // example:`{"th":"สสนก."}`  ชื่อหน่วยงาน
	EventCategory       json.RawMessage   `json:"event_type,omitempty"`     // example:`{"th":"Critical"}` ชื่อหมวดหลัก
	EventMessage        string            `json:"event_message"`            // example:`ข้อความเหตุการณ์` ข้อความเหตุการณ์
	StatusNotify        string            `json:"status_notify"`            // example:`true,false` สถานะแจ้งเตือน
	SolveEventDate      string            `json:"solve_event_at"`           // example:`2006-01-02 15:04` วันเวลาที่แก้ปัญหา
	Filepath            bool              `json:"filepath"`                 // ที่อยู่ไฟล์
	MetadataServiceName []json.RawMessage `json:"metadataservice_name"`     // example:`{"th":"metadataa"}` ชื่อ metadata
}

type EventTrackingUpdate struct {
	EventLogID []int64 `json:"event_log_id"`  // example:`[2245,423,4505]` ลำดับเหตุการณ์ เช่น [2245,423,4505]
	Message    string  `json:"event_message"` // example:`ข้อความ` ข้อความที่บันทึกเหตุการณ์ เช่น ข้อความ
}

type EventSendTrackingUpdate struct {
	EventLogID []int64 `json:"event_log_id"` // example:`[2245,423,4505]` ลำดับเหตุการณ์ เช่น [2245,423,4505]
	Date       string  `json:"event_date"`   // example:`2006-01-02` วันที่เกิดเหตุการณ์ เช่น  2006-01-02
}

type EventTrackingOption struct {
	DateStart    string  `json:"date_start"`     // example:`2017-08-24 00:00` วันที่เริ่มต้น เช่น  2017-08-24 00:00
	DateEnd      string  `json:"date_end"`       // example:`2017-08-24 20:00` วันที่สิ้นสุด เช่น 2017-08-24 20:00
	Agency       []int64 `json:"agency"`         // example:`[9]` รหัสหน่วยงาน เช่น [9]
	EventType    []int64 `json:"event_type"`     // example:`[2]` รหัสเหตุการณ์ เช่น [2]
	EventSubType []int64 `json:"event_sub_type"` // example:`[19]` รหัสเหตุการณ์ย่อย เช่น [9]
	SolveEvent   bool    `json:"solve_event"`    // example: false แก้ปัญหาเหตุการณ์
}

type EventInvalidData struct {
	EventLogId          int64           `json:"event_log_id"`         // example:`2324` ลำดับเหตุการณ์
	EventCode           string          `json:"event_code"`           // example:`1` หมวดย่อย
	EventDate           string          `json:"event_date"`           // example:`2006-01-02 15:04` วันที่เกิดเหตุการณ์
	Agency              string          `json:"agency_name"`          // example:`สสนก.` ชื่อหน่วยงาน
	ScriptName          string          `json:"script_name"`          // example:`convert name` ชื่อ script
	FileName            string          `json:"filename"`             // example:`filename` ชื่อไฟล์
	MetadataMethod      string          `json:"metadata_method"`      // example:`metadata method`
	Filepath            bool            `json:"filepath"`             // ที่อยู่ไฟล์
	MetadataServiceName json.RawMessage `json:"metadataservice_name"` // example:`{"th":"metadataa"}` ชื่อ metadata
}

type EventInvalidDataOption struct {
	Agency []int64 `json:"agency"` // example:`[9]` รหัสหน่วยงาน
	Date   string  `json:"date"`   // example:`2017-08-24` วันที่
}

type EventSendInvalidData struct {
	EventLogID          int64             `json:"event_log_id"`         // example:`2023` ลำดับบันทึกเหตุการณ์
	EventCode           string            `json:"event_code"`           // example:`1` รหัสเหตุการณ์
	EventDate           string            `json:"event_date"`           // example:`2006-01-02 15:04` วันที่บันทึกเหตุกาณณ์
	Agency              string            `json:"agency_name"`          // example:`สสนก.` ชื่อหน่วยงาน
	ScriptName          string            `json:"script_name"`          // example:`convert name` ชื่อ script
	EventMessage        string            `json:"event_message"`        // example:`error message` ข้อความเหตุการณ์
	MetadataServiceName []json.RawMessage `json:"metadataservice_name"` // example:`{"th":"metadataa"}` ชื่อ metadata
}

type EventSendInvalidDataOption struct {
	Agency    []int64 `json:"agency"`     // example:`[9]` รหัสหน่วยงาน เช่น  [1]
	DateStart string  `json:"date_start"` // example:`2017-06-25 00:00` เวลาเริ่มต้น เช่น 2017-06-25 00:00
	DateEnd   string  `json:"date_end"`   // example:`2017-06-27 23:59` เวลาสิ้นสุด เช่น 2017-06-27 23:59
}

type EventTrackingInvalidData struct {
	EventLogID          int64           `json:"event_log_id"`         // example:`2023` ลำดับบันทึกเหตุการณ์
	EventCode           string          `json:"event_code"`           // example:`1` รหัสเหตุการณ์
	EventDate           string          `json:"event_date"`           // example:`2006-01-02 15:04` วันที่บันทึกเหตุกาณณ์
	Agency              string          `json:"agency_name"`          // example:`สสนก.` ชื่อหน่วยงาน
	EventMessage        string          `json:"event_message"`        // example:`error message` ข้อความเหตุการณ์
	SendErrorAt         string          `json:"send_error_at"`        // example:`2006-01-02 15:04` วันเวลาที่ส่งปัญหา
	SolveEventDate      string          `json:"solve_event_at"`       // example:`2006-01-02 15:04` วันเวลาที่แก้ปัญหา
	SolveEventMsg       string          `json:"solve_event"`          // example:`แก้ไขแล้ว` บันทึกข้อความวิธีแกปัญหา
	MetadataServiceName json.RawMessage `json:"metadataservice_name"` // example:`{"th":"metadataa"}` ชื่อ metadata
}
