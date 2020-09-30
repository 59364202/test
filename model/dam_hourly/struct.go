package dam_hourly

import (
	"encoding/json"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
)

type Struct_Dam_Houly struct {
	Dam_id                 int64       `json:"dam_id,omitempty"`       // example:`17` รหัสเขื่อน
	Id                     int64       `json:"id"`                     // example:`574848` รหัสข้อมูล
	Dam_date               string      `json:"dam_date"`               // example:`2006-01-02 15:04` วันที่เก็บข้อมูล
	Dam_storage            interface{} `json:"dam_storage"`            // example:`40.9297` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	Dam_storage_percent    interface{} `json:"dam_storage_percent"`    // example:`null` เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.)
	Dam_inflow             interface{} `json:"dam_inflow"`             // example:`0.1859` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	Dam_inflow_acc_percent interface{} `json:"dam_inflow_acc_percent"` // example:`null` เปอร์เซนต์ปริมาณน้ำไหลเทียบกับปริมาณน้ำไหลลงเขื่อนขนาดใหญ่เฉลี่ยรวมทั้งปี (%)
	Dam_uses_water         interface{} `json:"dam_uses_water"`         // example:`null` ปริมาณน้ำที่ใช้ได้
	Dam_uses_water_percent interface{} `json:"dam_uses_water_percent"` // example:`null` เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.)
	Dam_level              interface{} `json:"dam_level"`              // example:`57.74` ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.)
	Dam_released           interface{} `json:"dam_released"`           // example:`0.468` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	Dam_spilled            interface{} `json:"dam_spilled"`            // example:`0` ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง
	Dam_losses             interface{} `json:"dam_losses"`             // example:`null` ปริมาณน้ำที่สูญเสีย
	Dam_evap               interface{} `json:"dam_evap"`               // example:`null` ปริมาณน้ำที่ระเหย

	Station_type string `json:"station_type,omitempty"` // example:`dam_hourly` ชนิดของสถานี
}

type Struct_DamHourly struct {
	Struct_Dam_Houly
	Dam     *model_dam.Struct_Dam         `json:"dam,omitempty"`     // เขื่อน
	Agency  *model_agency.Struct_Agency   `json:"agency,omitempty"`  // หน่วยงาน
	Basin   *model_basin.Struct_Basin     `json:"basin,omitempty"`   // ลุ่มน้ำ
	Geocode *model_geocode.Struct_Geocode `json:"geocode,omitempty"` // ข้อมูลขอบเขตการปกครองของประเทศไทย

}

type Struct_DamHourly_ErrorData struct {
	ID                  int64           `json:"id"`                     // รหัสข้อมูล
	StationID           int64           `json:"station_id"`             // รหัสเขื่อน
	StationOldCode      string          `json:"station_oldcode"`        // รหัสเดิมของเขื่อน
	Datetime            string          `json:"datetime"`               // วันที่เก็บข้อมูล
	StationName         json.RawMessage `json:"station_name"`           // ชื่อเขื่อน
	ProvinceName        json.RawMessage `json:"station_province_name"`  // ชื่อจังหวัดของประเทศไทย
	AgencyName          json.RawMessage `json:"agency_name"`            // ชื่อหน่วยงาน
	AgencyShortName     json.RawMessage `json:"agency_shortname"`       // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
	DamStorage          interface{}     `json:"dam_storage"`            // ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	DamStoragePercent   interface{}     `json:"dam_storage_percent"`    // เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.)
	DamInflow           interface{}     `json:"dam_inflow"`             // ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	DamInflowAccPercent interface{}     `json:"dam_inflow_acc_percent"` // เปอร์เซนต์ปริมาณน้ำไหลเทียบกับปริมาณน้ำไหลลงเขื่อนขนาดใหญ่เฉลี่ยรวมทั้งปี (%)
	DamUsesWater        interface{}     `json:"dam_uses_water"`         // ปริมาณน้ำที่ใช้ได้
	DamUsesWaterPercent interface{}     `json:"dam_uses_water_percent"` // เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.)
	DamLevel            interface{}     `json:"dam_level"`              // ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.)
	DamReleased         interface{}     `json:"dam_released"`           // ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	DamSpilled          interface{}     `json:"dam_spilled"`            // ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง
	DamLosses           interface{}     `json:"dam_losses"`             // ปริมาณน้ำที่สูญเสีย
	DamEvap             interface{}     `json:"dam_evap"`               // ปริมาณน้ำที่ระเหย
}

type Struct_DamHourly_InputParam struct {
	Id         string `json:"id"`
	Dam_id     string `json:"station_id"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

type Struct_DamHourlyLastest_InputParam struct {
	Region_id   string `json:"region_id"`   // รหัสภาค
	Province_id string `json:"province_id"` // รหัสจังหวัด
	Basin_id    string `json:"basin_id"`    // รหัสลุ่มน้ำ
	Dam_date    string `json:"dam_date"`    // วันที่
	Agency_id   string `json:"agency_id"`   // รหัสหน่วยงาน
	Start_date  string `json:"start_date"`  // วันที่ีเริ่มต้น
	End_date    string `json:"end_date"`    // วันที่สิ้นสุด
}

type DamHourlyLastest_OutputParam struct {
	Header []string            `json:"header"` // example:`["id","dam_date","dam_level","dam_storage","dam_storage_percent","dam_inflow","dam_inflow_acc_percent","dam_uses_water","dam_uses_water_percent","dam_released","dam_spilled","dam_losses","dam_evap"]` หัวตาราง
	Data   []*Struct_Dam_Houly `json:"data"`   // ข้อมูล
}

var DamHourlyColumnForCheckErrordata = []string{
	"id",
	"station_oldcode",
	"datetime",
	"station_name",
	"station_province_name",
	"agency_name",
	"agency_shortname",
	"dam_level",
	"dam_storage",
	"dam_storage_percent",
	"dam_inflow",
	"dam_inflow_acc_percent",
	"dam_uses_water",
	"dam_uses_water_percent",
	"dam_released",
	"dam_spilled",
	"dam_losses",
	"dam_evap",
}
