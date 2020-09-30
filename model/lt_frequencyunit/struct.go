package lt_frequencyunit

import (
	"encoding/json"
)

type FrequencyUnit_struct struct {
	Id                int64           `json:"id"`                           // example:`2` รหัสหน่วยของความถี่การเชื่อมโยง
	ConvertMinute     int64           `json:"convert_minute"`               // example:`60` แปลงหน่วยเป็นนาที
	FrequencyUnitName json.RawMessage `json:"frequencyunit_name,omitempty"` // example:`{"th": "ชั่วโมง"}` ชื่อหน่วยของระยะเวลาของการเชื่อมโยง
}
