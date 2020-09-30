package event_log_sink_template

import (
	"encoding/json"
)

type TemplateList struct {
	ID   int64  `json:"id"` // example:`1` รหัส template
	Name string `json:"name"` // example:`template name` ชื่อ template
}

type TemplateDetails struct {
	ID             int64           `json:"id"` // example:`1` รหัส template
	Name           string          `json:"name"` // example:`template name` ชื่อ template
	MessageSubject json.RawMessage `json:"message_subject,omitempty"` // example:`{"th":"สรุปรายการเหตุการณ์ที่เกิดขึ้นในระบบคลังข้อมูลน้ำฯ"}` หัวข้อ email
	MessageBody    json.RawMessage `json:"message_body,omitempty"` // example:`{"th":"<HTML><HEAD><meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\"><TITLE>สรุปรายการเหตุการณ์ที่เกิดขึ้นในระบบคลังข้อมูลน้ำฯ</TITLE></HEAD><BODY></BODY></HTML>"}` เนื้อหา email
}

type TemplatInputSwagger struct {
	Name           string      `json:"name"` // example:`template name` ชื่อ template
	MessageSubject interface{} `json:"message_subject"` // example:`{"th":"สรุปรายการเหตุการณ์ที่เกิดขึ้นในระบบคลังข้อมูลน้ำฯ"}` หัวข้อ email
	MessageBody    interface{} `json:"message_body"` // example:`{"th":"<HTML><HEAD><meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\"><TITLE>สรุปรายการเหตุการณ์ที่เกิดขึ้นในระบบคลังข้อมูลน้ำฯ</TITLE></HEAD><BODY></BODY></HTML>"}` เนื้อหา email
}

type TemplatInput struct {
	ID             int64       `json:"id"` // example:`1` รหัส template
	Name           string      `json:"name"` // example:`template name` ชื่อ template
	MessageSubject interface{} `json:"message_subject"` // example:`{"th":"สรุปรายการเหตุการณ์ที่เกิดขึ้นในระบบคลังข้อมูลน้ำฯ"}` หัวข้อ email
	MessageBody    interface{} `json:"message_body"` // example:`{"th":"<HTML><HEAD><meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\"><TITLE>สรุปรายการเหตุการณ์ที่เกิดขึ้นในระบบคลังข้อมูลน้ำฯ</TITLE></HEAD><BODY></BODY></HTML>"}` เนื้อหา email
}
