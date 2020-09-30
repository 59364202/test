package medium_dam

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
)

type Struct_Dam struct {
	Id              int64           `json:"id"`                    // example:`1` รหัสข้อมูลเขื่อน
	Dam_oldcode     string          `json:"dam_oldcode,omitempty"` // example:`1` รหัสเดิมของเขื่อน
	Dam_lat         float64         `json:"dam_lat,omitempty"`     // example:`17.242000` พิกัดของเขื่อน latitude
	Dam_long        float64         `json:"dam_long,omitempty"`    // example:`98.972800` พิกัดของเขื่อน longitude
	Dam_name        json.RawMessage `json:"dam_name"`              // example:`{"th":"ภูมิพล"}` ชื่อเขื่อน
	Max_storage     float64         `json:"max_storage,omitempty"` // example:`13462` ปริมาตรน้ำที่ระดับเก็บกักสูงสุด [ล้าน ลบ.ม.]
	Min_storage     float64         `json:"min_storage,omitempty"` // example:`13462` ปริมาตรน้ำที่ระดับเก็บกักต่ำงสุด [ล้าน ลบ.ม.]
	Max_water_level float64         `json:"max_water_level,omitempty"`
	Min_water_level float64         `json:"min_water_level,omitempty"`
}

type Struct_DamGroupByAgency struct {
	Agency *model_agency.Struct_Agency `json:"agency,omitempty"` // หน่วยงาน
	Dam    []*Struct_Dam               `json:"dam,omitempty"`    // เขื่อน
}

type Struct_Medium_Dam struct {
	Id                  int64       `json:"id"`                  // example:`574848` รหัสข้อมูล
	Dam_date            string      `json:"dam_date"`            // example:`2006-01-02 15:04` วันที่เก็บข้อมูล
	Dam_storage         interface{} `json:"dam_storage"`         // example:`40.9297` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	Dam_storage_percent interface{} `json:"dam_storage_percent"` // example:`null` เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.)
	Dam_inflow          interface{} `json:"dam_inflow"`          // example:`0.1859` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	Dam_uses_water      interface{} `json:"dam_uses_water"`      // example:`null` ปริมาณน้ำที่ใช้ได้
	Dam_released        interface{} `json:"dam_released"`        // example:`0.468` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
}

type Struct_MediumDamLatest struct {
	Struct_Medium_Dam
	Dam     *model_dam.Struct_Dam         `json:"dam,omitempty"`     // เขื่อน
	Agency  *model_agency.Struct_Agency   `json:"agency,omitempty"`  // หน่วยงาน
	Basin   *model_basin.Struct_Basin     `json:"basin,omitempty"`   // ลุ่มน้ำ
	Geocode *model_geocode.Struct_Geocode `json:"geocode,omitempty"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
}

type Struct_DamHourlyLastest_InputParam struct {
	Region_id   string `json:"region_id"`   // รหัสภาค
	Basin_id    string `json:"basin_id"`    // รหัสลุ่มน้ำ
	Province_id string `json:"province_id"` // รหัสจังหวัด
	Dam_date    string `json:"dam_date"`    // วันที่
	Agency_id   string `json:"agency_id"`   // รหัสหน่วยงาน
}
