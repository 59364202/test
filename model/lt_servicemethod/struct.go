package lt_servicemethod

import (
	"encoding/json"
)

type Struct_LtServicemethod struct {
	ID              int64           `json:"id"`                           // example:`1` รหัสวิธีการให้บริการข้อมูล
	ServiceMethodID string          `json:"servicemethod_id"`             // example:`zQAPyF_nUk1RmdOrAVHZa60udwS8VGdQ8b2xirnuhFqHNHylxmMtHgHhD5ZH8uSXUtB2BSxvl1EwQAWICe_KLw` รหัสวิธีการให้บริการข้อมูล
	Name            json.RawMessage `json:"servicemethod_name,omitempty"` // example:`{"th":"CD/DVD"}` ชื่อวิธีการให้บริการข้อมูล
	UserId          int64           `json:"-"`                            // example:`69` รหัสผู้ใช้
}

type Param_LtServicemethod struct {
	Name json.RawMessage `json:"servicemethod_name"` // example:`{"th":"CD/DVD"}` ชื่อวิธีการให้บริการข้อมูล
}
