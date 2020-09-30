package subcategory

import (
	"encoding/json"
)

type SubCategory_struct struct {
	Id              int64           `json:"id"`                         // example:`1` รหัสหมวดหมู่ย่อย
	SubCategoryName json.RawMessage `json:"subcategory_name,omitempty"` // example:`{"th": "พายุ", "en": "storm"}` ชื่อหมวดหมู่ย่อย
	CategoryId      interface{}     `json:"category_id"`                // example:`1` รหัสหมวดหมู่หลัก
	CategoryName    json.RawMessage `json:"category_name,omitempty"`    // example:`{"th": "อุตุนิยมวิทยา", "en": "meteorology"}` ชื่อหมวดหมู่หลัก
}
