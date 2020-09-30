// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package rainfall_7d

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_rain24 "haii.or.th/api/thaiwater30/model/rainfall24hr"
)

// param for input query condition
type Param_Rainfall7d struct {
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
}

type Struct_Rainfall7d struct {
	Start_date string      `json:"rainfall_start_date"` // example:`2006-01-02` วันที่เริ่มต้น
	End_date   string      `json:"rainfall_end_date"`   // example:`2006-01-02` วันที่สิ้นสุด
	Rain7D     interface{} `json:"rain_7d,omitempty"`   // example:`235` ฝน 3 วัน (mm.)

	Agency  *model_agency.Struct_Agency      `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode    `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *model_rain24.Struct_TeleStation `json:"station"` // สถานี
}