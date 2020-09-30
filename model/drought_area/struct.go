// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package drought_area is a model for public.drought_area table. This table store drought_area.
package drought_area

import (
	"encoding/json"
)

type Struct_DroughtArea struct {
	ProvinceCode string           `json:"province_code"` // example:`10` รหัสจังหวัดของประเทศไทย
	ProvinceName *json.RawMessage `json:"province_name"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
}
