// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Werawan Prongpanom <werawan@haii.or.th>

package rainfall_month

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_rain24 "haii.or.th/api/thaiwater30/model/rainfall24hr"
)

// param for input query condition
type Param_RainfallMonth struct {
	Province_Code string `json:"province_code"`	// required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
}

type Struct_RainfallMonth struct {
	Start_date string      `json:"rainfall_start_date"`		// example:`2006-01-01 07:00:00` วันที่เริ่มต้น
	End_date   string      `json:"rainfall_end_date"`		// example:`2006-02-01 07:00:00` วันที่สิ้นสุด
	RainMonth  interface{} `json:"rain_month,omitempty"`	// example:`661.00` ฝน สะสมเดือนที่ผ่านมา  (mm.)

	Agency  *model_agency.Struct_Agency      `json:"agency"`	// หน่วยงาน
	Geocode *model_geocode.Struct_Geocode    `json:"geocode"`	// ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *model_rain24.Struct_TeleStation `json:"station"`	// สถานี
}
