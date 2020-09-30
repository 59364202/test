package event_code

import (
	"encoding/json"
	model_event_log_category "haii.or.th/api/thaiwater30/model/event_log_category"
)

type Struct_EventCode struct {
	EventCategory   *model_event_log_category.Struct_EventLogCategory `json:"event_log_category"`    // ประเภทเหตุการณ์
	ID              int64                                             `json:"id"`                    // example:`42` รหัส event code
	Code            string                                            `json:"code"`                  // example:`EventDataErrorInvalidData` โค้ด event code
	Description     json.RawMessage                                   `json:"description,omitempty"` // example:`{"en": "This record contains invalid data", "th": "ข้อผิดพลาดของข้อมูล"}` คำอธิบายเพิ่มเติม
	IsAutoclose     bool                                              `json:"is_autoclose"`          // example:`false` autoclose
	Troubleshoot    string                                            `json:"troubleshoot"`          // example:`how to fix it` วิธีแก้ปัญหา
	SubtypeCategory string                                            `json:"subtype_category"`      // example:`data` subtype category
}

type Struct_EventCode_SummaryEvent struct {
	Struct_EventCode
	SummaryEvent int64 `json:"summary_event,omitempty"` // example:`1` จำนวนเหตุการณ์
}

type Struct_EventCode_InputParamSwagger struct {
	EventCategoryID string          `json:"event_log_category_id"` // example:`1` event category id
	Code            string          `json:"code"`                  // example:`EventDataErrorInvalidData` โค้ด event code
	Description     json.RawMessage `json:"description,omitempty"` // example:`{"en": "This record contains invalid data", "th": "ข้อผิดพลาดของข้อมูล"}` คำอธิบายเพิ่มเติม
	IsAutoclose     bool            `json:"is_autoclose"`          // example:`false` autoclose
	Troubleshoot    string          `json:"troubleshoot"`          // example:`how to fix it` วิธีแก้ปัญหา
	SubtypeCategory string          `json:"subtype_category"`      // example:`data` subtype category
}

type Struct_EventCode_InputParamGetSwagger struct {
	EventCategoryID string `json:"event_log_category_id"` // example:`1`event category id
	ID              string `json:"id"`                    // example:`3` eventcode id
}

type Struct_EventCode_InputParam struct {
	EventCategoryID string          `json:"event_log_category_id"` // example:`1`event category id
	ID              string          `json:"id"`                    // example:`3` eventcode id
	Code            string          `json:"code"`                  // example:`EventDataErrorInvalidData` โค้ด event code
	Description     json.RawMessage `json:"description,omitempty"` // example:`{"en": "This record contains invalid data", "th": "ข้อผิดพลาดของข้อมูล"}` คำอธิบายเพิ่มเติม
	IsAutoclose     bool            `json:"is_autoclose"`          // example:`false` autoclose
	Troubleshoot    string          `json:"troubleshoot"`          // example:`how to fix it` วิธีแก้ปัญหา
	SubtypeCategory string          `json:"subtype_category"`      // example:`data` subtype category
}

type Struct_Event struct {
	Id       int64              `json:"id"`   // example:`1`event category id
	Code     string             `json:"code"` // example:`EventDataErrorInvalidData` โค้ด event code
	Subevent []*Struct_Subevent `json:"subevent"`
}

type Struct_Subevent struct {
	Id          int64           `json:"id"`          // example:`3` eventcode id
	Description json.RawMessage `json:"description"` // example:`{"en": "This record contains invalid data", "th": "ข้อผิดพลาดของข้อมูล"}` คำอธิบายเพิ่มเติม
}
