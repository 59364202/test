// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package basin is a model for public.basin table. This table store basin.
package basin

import (
	"encoding/json"
)

// โครงสร้างหลักของ basin
type Struct_Basin struct {
	Id         int64           `json:"id"`                   // example:`1` รหัสลุ่มน้ำ
	Agency_id  int64           `json:"agency_id,omitempty"`  // example:`3` รหัสหน่วยงาน
	Basin_code int64           `json:"basin_code,omitempty"` // example:`1` รหัสลุ่มน้ำ
	Basin_name json.RawMessage `json:"basin_name,omitempty"` // example:`{"th":"ลุ่มน้ำสาละวิน","en":"MAE NAM SALAWIN"}` ชื่อลุ่มน้ำ
}
