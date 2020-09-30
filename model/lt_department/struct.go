package lt_department

import (
	"encoding/json"
)

type Struct_Department struct {
	Id              int64           `json:"id"`                        // example:`1` ลำดับข้อมูลกรม
	Department_Name json.RawMessage `json:"department_name,omitempty"` // example:`{"th":"สำนักงานปลัดสำนักนายกรัฐมนตรี"}` ชื่อกรม
}
type Department_struct struct {
	Id                  int64           `json:"id"`                             // example:`1` ลำดับข้อมูลกรม
	DepartmentCode      string          `json:"department_code"`                // example:`1001` รหัสข้อมูลกรม
	DepartmentName      json.RawMessage `json:"department_name,omitempty"`      // example:`{"th":"สำนักงานปลัดสำนักนายกรัฐมนตรี"}` ชื่อกรม
	DepartmentShortName json.RawMessage `json:"department_shortname,omitempty"` // example:`{"th":""}` ชื่อย่อกรม
	MinistryId          int64           `json:"ministry_id"`                    // example:`1` ลำดับข้อมูลกระทรวง
	MinistryName        json.RawMessage `json:"ministry_name,omitempty"`        // example:`{"th":"สำนักนายกรัฐมนตรี","en":"Prime Minister's Office"}` ชื่อกระทรวง
}
