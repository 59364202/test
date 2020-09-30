package dam

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"

	"encoding/json"
)

// โครงสร้างแบบใหญ่
type Struct_Dam struct {
	Agency  *model_agency.Struct_Agency   `json:"agency,omitempty"`  // หน่วยงาน
	Basin   *model_basin.Struct_Basin     `json:"basin,omitempty"`   // ลุ่มน้ำ
	Geocode *model_geocode.Struct_Geocode `json:"geocode,omitempty"` // ข้อมูลขอบเขตการปกครองของประเทศไทย

	Max_storage     float64 `json:"max_storage,omitempty"`     // example:`13462` ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.]
	Min_storage     float64 `json:"min_storage,omitempty"`     // example:`3800` ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด  [ล้าน ลบ.ม.]
	Max_water_level float64 `json:"max_water_level,omitempty"` // ระดับกักเก็บสูงสุด [ม.(รทก.)]
	Min_water_level float64 `json:"min_water_level,omitempty"` // ระดับกักเก็บต่ำสุด [ม.(รทก.)]
	SubBasin_id     int64   `json:"sub_basin_id,omitempty"`    // example:`21` รหัสลุ่มน้ำย่อย

	Struct_D
}

// โครงสร้างแบบเล็ก
type Struct_D struct {
	Id          int64           `json:"id"`                     // example:`1` รหัสเขื่อน
	Dam_oldcode string          `json:"dam_oldcode,omitempty"`  // example:`1` รหัสเดิมของเขื่อน
	Dam_name    json.RawMessage `json:"dam_name"`               // example:`{"th":"ภูมิพล"}` ชื่อเขื่อน
	Dam_lat     interface{}     `json:"dam_lat"`                // example:`17.242000` พิกัดละติจูด
	Dam_long    interface{}     `json:"dam_long"`               // example:`98.972800` พิกัดลองติจูด
	Max_storage interface{}     `json:"max_storage,omitempty"`  // example:`13462` ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.]
	SubBasin_id int64           `json:"sub_basin_id,omitempty"` // example:`21` รหัสลุ่มน้ำย่อย
}

type Struct_GetDam struct {
	Struct_D
	Max_storage float64 `json:"max_storage,omitempty"` // example:`13462` ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.]
	Min_storage float64 `json:"min_storage,omitempty"` // example:`3800` ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด  [ล้าน ลบ.ม.]

	Agency *model_agency.Struct_A `json:"agency,omitempty"` // หน่วยงาน
}

type Struct_Dam_ForCheckMetadata struct {
	Id              int64           `json:"id"`               // รหัสเขื่อน
	Name            json.RawMessage `json:"station_name"`     // ชื่อเขื่อน
	OldCode         string          `json:"station_oldcode"`  // รหัสเดิมของเขื่อน
	Lat             float64         `json:"lat"`              // พิกัดละติจูด
	Long            float64         `json:"long"`             // พิกัดลองติจูด
	Geocode         string          `json:"geocode"`          // รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	SubbasinName    json.RawMessage `json:"subbasin_name"`    // ชื่อลุ่มน้ำสาขา
	AgencyShortname json.RawMessage `json:"agency_shortname"` // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
	AgencyName      json.RawMessage `json:"agency_name"`      // ชื่อหน่วยงาน
}

type Struct_GetDam_InputParam struct {
	AgencyId   string `json:"agency_id"` // รหัสหน่วยงาน
	Date       string `json:"date"`
	Basin      string `json:"basin"`
	DataType   string `json:"data_type"`
	ColumnName string `json:"column_name"`
}

type Struct_GetDam_OutputParam struct {
	Id               int64           `json:"id"`               // รหัสเขื่อน
	Station_name     json.RawMessage `json:"station_name"`     // ชื่อเขื่อน
	Station_oldcode  string          `json:"station_oldcode"`  // รหัสเดิมของเขื่อน
	Agency_name      json.RawMessage `json:"agency_name"`      // ชื่อหน่วยงาน
	Agency_shortname json.RawMessage `json:"agency_shortname"` // ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)
}

type Struct_DamGroupByAgency struct {
	Agency *model_agency.Struct_Agency `json:"agency,omitempty"` // หน่วยงาน
	Dam    []*Struct_Dam               `json:"dam,omitempty"`    // เขื่อน
}
