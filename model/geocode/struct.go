// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package geocode is a model for public.lt_geocode table. This table store geocode information.
package geocode

import (
	"encoding/json"
)

type Struct_Geocode struct {
	Id            int64           `json:"id,omitempty"`            // example:`3` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Geocode       string          `json:"geocode,omitempty"`       // example:`100101` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Area_code     string          `json:"area_code,omitempty"`     // example:`1` รหัสภาคของประเทศไทย
	Area_name     json.RawMessage `json:"area_name,omitempty"`     // example:`{"th": "กรุงเทพมหานคร"}` ชื่อภาคของประเทศไทย
	Province_code string          `json:"province_code,omitempty"` // example:`10` รหัสจังหวัดของประเทศไทย
	Province_name json.RawMessage `json:"province_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	Amphoe_code   string          `json:"amphoe_code,omitempty"`   // example:`01` รหัสอำเภอของประเทศไทย
	Amphoe_name   json.RawMessage `json:"amphoe_name,omitempty"`   // example:`{"th": "พระนคร"}` ชื่ออำเภอของประเทศไทย
	Tumbon_code   string          `json:"tumbon_code,omitempty"`   // example:`01` รหัสตำบลของประเทศไทย
	Tumbon_name   json.RawMessage `json:"tumbon_name,omitempty"`   // example:`{"th": "พระบรมมหาราชวัง"}` ชื่อตำบลของประเทศไทย
}

type Param_Geocode struct {
	Geocode       string `json:"geocode"`       // รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Area_code     string `json:"area_code"`     // รหัสภาคของประเทศไทย
	Province_code string `json:"province_code"` // รหัสจังหวัดของประเทศไทย
	Amphoe_code   string `json:"amphoe_code"`   // รหัสอำเภอของประเทศไทย
	Tumbon_code   string `json:"tumbon_code"`   // รหัสตำบลของประเทศไทย
}
