package event_log_sink_condition

import (
	"encoding/json"
	"haii.or.th/api/thaiwater30/util/selectoption"
)

type SinkConditionList struct {
	ID                int64           `json:"id"`                          // example:`1` หมายเลข เงื่อนไข
	Condition         string          `json:"name"`                        // example:`condition name` ชื่อเงื่อนไข
	Category          json.RawMessage `json:"event_log_type,omitempty"`    // example:`{"th":"category name"}` ชื่อหมวดหลัก
	Code              json.RawMessage `json:"event_log_subtype,omitempty"` // example:`{"th":"code name"}` ชื่อหมวดย่อย
	Template          string          `json:"event_log_sink_template"`     // example:`template name` ชื่อ template
	PostStartInterval string          `json:"post_start_interval"`         // example:`0 * * * *` เวลาเริ่มต้นของเงื่อนไข
}
type SinkCondition struct {
	ID                int64           `json:"id"`                          // example:`1` หมายเลข เงื่อนไข
	Condition         string          `json:"name"`                        // example:`condition name` ชื่อเงื่อนไข
	CategoryID        interface{}     `json:"event_log_type_id"`           // example:`1` รหัส หมวดหลัก
	Category          json.RawMessage `json:"event_log_type,omitempty"`    // example:`{"th":"category name"}` ชื่อหมวดหลัก
	CodeID            interface{}     `json:"event_log_subtype_id"`        // example:`2` รหัสหมวดย่อย
	Code              json.RawMessage `json:"event_log_subtype,omitempty"` /// example:`{"th":"code name"}` ชื่อหมวดย่อย
	TemplateID        int64           `json:"event_log_sink_template_id"`  // example:`4` รหัส template
	Template          string          `json:"event_log_sink_template"`     // example:`template name` ชื่อ template
	PostStartInterval string          `json:"post_start_interval"`         // example:`0 * * * *` เวลาเริ่มต้นของเงื่อนไข
}

type SinkConditionSelectOptionSwagger struct {
	Category []*CategoryOptionList `json:"event_log_type"`          // รายชื่อ category
	Template []*LTemplate          `json:"event_log_sink_template"` // รายชื่อ template
}

type LTemplate struct {
	Text  string `json:"texg"`  // example:`รายงานประจำวัน`
	Value int64  `json:"value"` // example:`1`
}

type SinkConditionSelectOption struct {
	Category []*CategoryOptionList  `json:"event_log_type"`          // รายชื่อ category
	Template []*selectoption.Option `json:"event_log_sink_template"` // รายชื่อ template
}

type CategoryOptionList struct {
	Category *SelectOptionJson   `json:"event_log_type"`    // รายชื่อ category
	Code     []*SelectOptionJson `json:"event_log_subtype"` // รายชื่อ code
}

type SelectOptionJson struct {
	Value interface{}     `json:"value"` // example:`1` ข้อมูล
	Text  json.RawMessage `json:"text"`  // example:`{"th":"category name"}` ชื่อ
}

type SinkConditionInputSwagger struct {
	Name              string      `json:"name"`                 // example:`condition name` ชื่อ condition เช่น Match All Condition
	Category          interface{} `json:"event_log_type_id"`    // example:`4` รหัส หมวดหลัก เช่น error condition
	Code              interface{} `json:"event_log_subtype_id"` // example:`20` รหัส หมวดย่อย เช่น Network Timeout
	Template          int64       `json:"template_id"`          // example:`4` รหัส template เช่น 20
	PostStartInterval string      `json:"post_start_interval"`  // example:`0 * * * *` ตั้งหน่วงเวลาในการส่ง เช่น 0 * * * *
}

type SinkConditionInputSwagger2 struct {
	ID                int64       `json:"id"`                   // example:`1` รหัส condition เช่น 3
	Name              string      `json:"name"`                 // example:`condition name` ชื่อ condition เช่น Match All Condition
	Category          interface{} `json:"event_log_type_id"`    // example:`4` รหัส หมวดหลัก เช่น error condition
	Code              interface{} `json:"event_log_subtype_id"` // example:`20` รหัส หมวดย่อย เช่น Network Timeout
	Template          int64       `json:"template_id"`          // example:`4` รหัส template เช่น 20
	PostStartInterval string      `json:"post_start_interval"`  // example:`0 * * * *` ตั้งหน่วงเวลาในการส่ง เช่น 0 * * * *
}

type SinkConditionInput struct {
	ID                int64       `json:"id"`                   // example:`1` รหัส condition
	Name              string      `json:"name"`                 // example:`condition name` ชื่อ condition
	Channel           interface{} `json:"event_log_channel"`    // example:`3` รหัส eventlog channel
	Category          interface{} `json:"event_log_type_id"`    // example:`4` รหัส หมวดหลัก
	Code              interface{} `json:"event_log_subtype_id"` // example:`20` รหัส หมวดย่อย
	Service           interface{} `json:"service_id"`           // example:`10` รหัส service
	Agent             interface{} `json:"agent_user_id"`        // example:`8` รหัส agent
	User              interface{} `json:"user_id"`              // example:`56` รหัส user
	Template          int64       `json:"template_id"`          // example:`4` รหัส template
	PostStartInterval string      `json:"post_start_interval"`  // example:`0 * * * *` ตั้งหน่วงเวลาในการส่ง
}
