package lt_ministry

import (
	"encoding/json"
)

type Struct_Ministry struct {
	Id            int64           `json:"id"`                       // example:`1` ลำดับข้อมูลกระทรวง
	Ministry_Name json.RawMessage `json:"ministry_name,omitenmpty"` // example:`{"th":"สำนักนายกรัฐมนตรี","en":"Prime Minister's Office"}` ชื่อกระทรวง
}

type Ministry_struct struct {
	Id                int64           `json:"id"`                           // example:`1` ลำดับข้อมูลกระทรวง
	Code              string          `json:"ministry_code"`                // example:`1000` รหัสข้อมูลกระทรวง
	MinistryName      json.RawMessage `json:"ministry_name,omitempty"`      // example:`{"th":"สำนักนายกรัฐมนตรี","en":"Prime Minister's Office"}` ชื่อกระทรวง
	MinistryShortName json.RawMessage `json:"ministry_shortname,omitempty"` // example:`{"th":"นรนายกรัฐมนตรี"}` ชื่อย่อกระทรวง
}
