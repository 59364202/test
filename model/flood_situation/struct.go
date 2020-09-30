// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package flood_situation is a model for public.flood_situation table. This table store flood_situation.
package flood_situation

import (
	"encoding/json"
)

type Struct_FloodSituation struct {
	ProvinceCode     string           `json:"province_code"`     // รหัสจังหวัดของประเทศไทย
	ProvinceName     *json.RawMessage `json:"province_name"`     // ชื่อจังหวัดของประเทศไทย
	FloodDatetime    string           `json:"flood_datetime"`    // วันและเวลาที่ประกาศสถานการณ์น้ำ
	FloodName        string           `json:"flood_name"`        // ชื่อสถานการณ์น้ำ
	FloodLink        string           `json:"flood_link"`        // ลิ้งที่แสดงสถานการณ์น้ำ
	FloodDescription string           `json:"flood_description"` // รายละเอียดสถานการณ์น้ำ
	FloodAuthor      string           `json:"flood_author"`      // ผู้รายงานสถานการณ์
	FloodColorlevel  string           `json:"flood_colorlevel"`  // ระดับสีเกณฑืเตือนภัย
}
