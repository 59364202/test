package water_map

import (
	"encoding/json"
)

type WaterMapOutput struct {
	DamDaily []*DamDailyOutput `json:"dam"`             // เขื่อน
	TeleWL   []*TeleWaterlevel `json:"tele_waterlevel"` // ระดับน้ำ
}

//type BasinDam struct {
//	BasinName json.RawMessage   `json:"basin_name"`
//	DamDaily  []*DamDailyOutput `json:"data"`
//}
//
//type BasinWaterlevel struct {
//	BasinName json.RawMessage   `json:"basin_name"`
//	TeleWL    []*TeleWaterlevel `json:"data"`
//}

type DamDailyOutput struct {
	DamName          json.RawMessage `json:"dam_name"`           // example:`{"th":"ภูมิพล"}` ชื่อเขื่อน
	AgencyName       json.RawMessage `json:"agency_name"`        // example:`{"th":"กฟผ"}` ชื่อหน่วยงาน
	BasinName        json.RawMessage `json:"basin_name"`         // example:`{"th":"ลุ่มน้ำ"}` ชื่อลุ่มน้ำ
	SubbasinName     json.RawMessage `json:"subbasin_name"`      // example:`{"th":"ลุ่มน้ำสาขา"}` ชื่อลุ่มน้ำสาขา
	AreaName         json.RawMessage `json:"area_name"`          // example:`{"th":"ภาคเหนือ"}` ชื่อภาค
	ProvinceName     json.RawMessage `json:"province_name"`      // example:`{"th":"ตาก"}` ชื่อจังหวัด
	AmphoeName       json.RawMessage `json:"amphoe_name"`        // example:`{"th":"สามเงา"}` ชื่ออำเภอ
	TumbonName       json.RawMessage `json:"tumbon_name"`        // example:`{"th":"ย่านรี"}` ชื่อตำบล
	Lat              interface{}     `json:"lat"`                // example:`17.241944` ละติจูด
	Long             interface{}     `json:"long"`               // example:`98.975278` ลองติจูด
	DamOldCode       string          `json:"dam_oldcode"`        // example:`1` รหัสสถานีเดิม
	MaxWaterlevel    interface{}     `json:"max_water_level"`    // example:`260` ระดับกักเก็บสูงสุด [ม.(รทก.)]
	NormalWaterlevel interface{}     `json:"normal_water_level"` // example:`260` ระดับกักเก็บปกติ [ม.(รทก.)]
	MinWaterlevel    interface{}     `json:"min_water_level"`    // example:`213` ระดับกักเก็บต่ำสุด [ม.(รทก.)]
	DamDate          string          `json:"dam_date"`           // example:`2006-01-02` วันที่เก็บข้อมูล
	DamInflow        interface{}     `json:"dam_inflow"`         // example:`10` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	DamReleased      interface{}     `json:"dam_released"`       // example:`11` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	DamStorage       interface{}     `json:"dam_storage"`        // example:`200` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
}

type TeleWaterlevel struct {
	TeleStationName    json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 8","th":"คลองแสนแสบ บางกะปิ","jp":"バンコク8"}` ชื่อสถานี
	AgencyName         json.RawMessage `json:"agency_name"`          // example:`{"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)","en":"Hydro and Agro Informatics Institute","jp":""}` ชื่อหน่วยงาน
	BasinName          json.RawMessage `json:"basin_name"`           // example:`{"th":"ลุ่มน้ำเจ้าพระยา","en":"MAE NAM CHAO PHRAYA"}` ชื่อลุ่มน้ำ
	SubbasinName       json.RawMessage `json:"subbasin_name"`        // example:`{"th":"ที่ราบแม่น้ำเจ้าพระยา","en":"MAE NAM CHAO PRAYA  PLAIN AREA"}` ชื่อลุ่มน้ำสาขา
	AreaName           json.RawMessage `json:"area_name"`            // example:`{"th":"กรุงเทพมหานคร"}` ชื่อภาค
	ProvinceName       json.RawMessage `json:"province_name"`        // example:`{"th":"กรุงเทพมหานคร"}` ชื่อจังหวัด
	AmphoeName         json.RawMessage `json:"amphoe_name"`          // example:`{"th":"บางกะปิ"}` ชื่ออำเภอ
	TumbonName         json.RawMessage `json:"tumbon_name"`          // example:`{"th":"หัวหมาก"}` ชื่อตำบล
	Lat                interface{}     `json:"lat"`                  // example:`13.761400` ละติจูด
	Long               interface{}     `json:"long"`                 // example:`100.616060` ลองติจูด
	TeleStationOldCode string          `json:"tele_station_oldcode"` // example:`BKK008` รหัสสถานีเดิม
	RightBank          interface{}     `json:"right_bank"`           // example:`10.8`  ระดับตลิ่ง (ขวา)
	LeftBank           interface{}     `json:"left_bank"`            // example:`32.9`  ระดับตลิ่ง (ซ้าย)
	WaterlevelDatetime string          `json:"waterlevel_datetime"`  // example:`2006-01-02 15:04` วันเวลาที่ตรวจสอบค่าระดับน้ำ
	WaterlevelM        interface{}     `json:"waterlevel_m"`         // example:`23.3` ระดับน้ำ เมตร
	WaterlevelMSL      interface{}     `json:"waterlevel_msl"`       // example:`22.1` ระดับน้ำ รทก
	WaterlevelIn       interface{}     `json:"waterlevel_in"`        // example:`1.2` ระดับน้ำด้านในประตูระบายน้ำ
	WaterlevelOut      interface{}     `json:"waterlevel_out"`       // example:`2.3` ระดับน้ำด้านนอกประตูระบายน้ำ
	WaterlevelOut2     interface{}     `json:"waterlevel_out2"`      // example:`4.3` ระดับน้ำด้านนอกประตูระบายน้ำ
	FlowRate           interface{}     `json:"flow_rate"`            // example:`3.3` อัตราการไหล
	Discharge          interface{}     `json:"discharge"`            // example:`2.3` ปริมาณการระบายน้ำ
}
