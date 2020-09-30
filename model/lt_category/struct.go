// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package category is a model for public.lt_category table. This table store lt_category information.
package lt_category

import (
	"encoding/json"
)

type Struct_category struct {
	Id            int64           `json:"id"`                      // example:`4` รหัสกลุ่มข้อมูลหลัก
	Category_name json.RawMessage `json:"category_name,omitempty"` // example:`{"th":"แผนที่","en":"Map","jp":"マップ"}` ชื่อกลุ่มข้อมูลหลัก
}
