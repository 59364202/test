// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show_system

import (
	"encoding/json"
)

type Struct_MetadataShowSystem struct {
	ID                 int64           `json:"id"`                             // example:`11111` รหัส
	MetadataShowSystem json.RawMessage `json:"metadata_show_system,omitempty"` // example:`{"th":"ระบบที่นำข้อมูลไปแสดง","en":"Thaiwater30"}` ระบบที่นำข้อมูลไปแสดง
}
