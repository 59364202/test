package lt_subcategory

import (
	"encoding/json"
	model_category "haii.or.th/api/thaiwater30/model/lt_category"
)

type Struct_subcategory struct {
	Id               int64                           `json:"id"`                         // example:`39` รหัสกลุ่มข้อมูลย่อย
	Subcategory_name json.RawMessage                 `json:"subcategory_name,omitempty"` // example:`{"th":"พท.ชลประทาน"}` ชื่อกลุ่มข้อมูลย่อย
	Category         *model_category.Struct_category `json:"category,omitempty"`         // กลุ่มข้อมูลหลัก
}

type SubCategory_struct struct {
	Id              int64           `json:"id"`                         // example:`39` รหัสกลุ่มข้อมูลย่อย
	SubCategoryName json.RawMessage `json:"subcategory_name,omitempty"` // example:`{"th":"พท.ชลประทาน"}` ชื่อกลุ่มข้อมูลย่อย
	CategoryId      interface{}     `json:"category_id"`                // example:`4` รหัสกลุ่มข้อมูลหลัก
	CategoryName    json.RawMessage `json:"category_name,omitempty"`    // example:`{"th": "พท.ชลประทาน"}` ชื่อกลุ่มข้อมูลหลัก
}
