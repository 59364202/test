package lt_dataunit

import (
	"encoding/json"
)

type Dataunit_struct struct {
	Id           int64           `json:"id"`                      // example:`4` รหัสหน่วยของข้อมูล
	DataunitName json.RawMessage `json:"dataunit_name,omitempty"` // example:`{"th": "กิโลเมตรต่อชั่วโมง", "en": "Kilometer per hour"}` ชื่อหน่วยข้อมูล
}
