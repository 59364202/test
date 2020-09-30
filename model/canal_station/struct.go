// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package canal_station is a model for public.m_canal_station table. This table store m_canal_station information.
package canal_station

import (
	"encoding/json"
)

// โครงสร้างหลัก
type Struct_CanalStation struct {
	Id                 int64           `json:"id"`                           // example:`1` รหัสสถานีวัดระดับน้ำในคลอง
	Station_Name       json.RawMessage `json:"station_name,omitempty"`       // example:`{"th":"แม่น้ำเจ้าพระยา บริเวณปากคลองตลาด"}` ชื่อสถานีวัดระดับน้ำในคลอง
	Canal_station_name json.RawMessage `json:"canal_station_name,omitempty"` // example:`{"th":"แม่น้ำเจ้าพระยา บริเวณปากคลองตลาด"}` ชื่อสถานีวัดระดับน้ำในคลอง
	Canal_station_lat  interface{}     `json:"canal_station_lat,omitempty"`  // example:`13.742200` ละติจูดสถานีวัดระดับน้ำในคลอง
	Canal_station_long interface{}     `json:"canal_station_long,omitempty"` // example:`100.494640` ลองติจูดสถานีวัดระดับน้ำในคลอง
	Geocode            string          `json:"geocode,omitempty"`            // example:`100102` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Province_name      json.RawMessage `json:"province_name,omitempty"`      // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	Province_code      string          `json:"province_code,omitempty"`      // example:`10` รหัสจังหวัดของประเทศไทย
}

//type Struct_CanalStation_ForCheckMetadata struct {
//	Id              int64           `json:"id"`               // รหัสสถานีวัดระดับน้ำในคลอง
//	Name            json.RawMessage `json:"station_name"`     // ชื่อสถานีวัดระดับน้ำในคลอง
//	OldCode         string          `json:"station_oldcode"`  // รหัสเดิมของสถานีวัดระดับน้ำในคลอง
//	Lat             float64         `json:"lat"`              // ละติจูดสถานีวัดระดับน้ำในคลอง
//	Long            float64         `json:"long"`             // ลองติจูดสถานีวัดระดับน้ำในคลอง
//	Geocode         string          `json:"geocode"`          // ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
//	SubbasinName    json.RawMessage `json:"subbasin_name"`    // ชื่อลุ่มน้ำสาขา
//	AgencyShortname json.RawMessage `json:"agency_shortname"` // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
//	AgencyName      json.RawMessage `json:"agency_name"`      // ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ
//}

// โครงสร้างหลัก ที่กรุ๊ปตามจังหวัด
type Struct_CanalStationGroupByProvince struct {
	ProvinceName json.RawMessage        `json:"province_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	ProvinceCode string                 `json:"province_code"`           // example:`10` รหัสจังหวัดของประเทศไทย
	Station      []*Struct_CanalStation `json:"station"`                 // สถานี
}

//type Struct_CanalStation_InputParam struct {
//	ColumnName string `json:"column_name"` // ชื่อคอลั่ม
//	AgencyID   string `json:"agency_id"`   // รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ
//}
