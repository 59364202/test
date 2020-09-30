package hydroinfo

import (
	"encoding/json"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
)

type Struct_Hydroinfo struct {
	ID              int64                         `json:"id"`                       // example:`5` รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
	HydroinfoName   json.RawMessage               `json:"hydroinfo_name,omitempty"` // example:`{"th":"การรักษาระบบนิเวศและคุณภาพน้ำ","en":"Ecosystem and Water Quality Preservation"}` ชื่อด้าน
	Agency          []*model_agency.Struct_Agency `json:"agency"`                   // หน่วยงาน
	HydroinfoNumber interface{}                   `json:"hydroinfo_number"`         // example:`5` ลำดับของชือ เพื่อใช้ในการนำเสนอ
}

type Struct_Hydroinfo_InputParam struct {
	ID              int64           `json:"id"`               // example:`1` รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
	HydroInfoName   json.RawMessage `json:"hydroinfo_name"`   // example:`{"th":"ภาษาไทย", "en":"english"}` ชื่อด้าน
	HydroinfoNumber int64           `json:"hydroinfo_number"` // example:`1` ลำดับของชือ เพื่อใช้ในการนำเสนอ
	AgencyID        string          `json:"agency_id"`        // example:`1` รหัสหน่วยงาน
}
