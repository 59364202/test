// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package dam_uses_water

import (
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
)

type Struct_DamUsesWater_InputParam struct {
	Id            string `json:"id"`
	Dam_id        string `json:"dam_id"`
	Province_Code string `json:"province_code"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
}

type Struct_DamUsesWater_OutputParam struct {
	Result string                 `json:"result"` // สถานะ
	Data   []*Struct_DamUsesWater `json:"data"`   // ข้อมูล
}

type Struct_DamUsesWater struct {
	Start_date             string                        `json:"dam_start_date"`         // วันที่ีเริ่มต้น
	End_date               string                        `json:"dam_end_date"`           // วันที่ีเริ่มต้น
	Dam_uses_water         interface{}                   `json:"dam_uses_water"`         // example:`2803.49` ปริมาณน้ำที่ใช้ได้
	Dam_uses_water_percent interface{}                   `json:"dam_uses_water_percent"` // example:`20.83` เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.)
	Dam_released           interface{}                   `json:"dam_released"`           // example:`2599` ปริมาณน้ำระบาย (ล้าน.ลบ.ม.)
	Geocode                *model_geocode.Struct_Geocode `json:"geocode,omitempty"`      // ข้อมูลขอบเขตการปกครองของประเทศไทย
}
