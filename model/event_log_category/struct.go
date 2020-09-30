package event_log_category

import (
	"encoding/json"
)

type Struct_ELC struct {
	ID          int64           `json:"id"`                    // example:`1` รหัสประเภทเหตุการณ์
	Code        string          `json:"code"`                  // example:`CRITICAL` โค้ดประเภทเหตุการณ์
	Description json.RawMessage `json:"description,omitempty"` // example:`{"th": "วิกฤต", "en": "Critical"}` ชื่อประเภทเหตุการณ์
	Color       string          `json:"color"`                 // example:`FFB777` สีประเภทเหตุการณ์
}

type Struct_EventLogCategory struct {
	Struct_ELC
	SummaryEvent int64 `json:"summary_event"` // example:`107` จำนวนเหตุการณ์ที่เกิดขึ้น
}

type Struct_EventLogCategory_InputParamSwagger struct {
	Code        string          `json:"code"` // example:`CRITICAL` เหตุการณ์
	Description json.RawMessage `json:"description,omitempty"` // example:`{"en":"Program Critical"}` ข้อความอธิบายเหตุการณ์
	Color       string          `json:"color"` // example:`FFB777` รหัสสี
}

type Struct_EventLogCategory_InputParam struct {
	ID          string          `json:"id"` //example:`1` รหัสเหตุการณ์
	Code        string          `json:"code"` // example:`CRITICAL` เหตุการณ์
	Description json.RawMessage `json:"description,omitempty"` // example:`{"en":"Program Critical"}` ข้อความอธิบายเหตุการณ์
	Color       string          `json:"color"` // example:`FFB777` รหัสสี
}
