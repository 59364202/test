package dam_daily

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"

	"encoding/json"
	//	"haii.or.th/api/thaiwater30/util/highchart"
)

type Struct_DamDaily_InputParam struct {
	Id         string `json:"id"`
	Dam_id     string `json:"station_id"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

//type Struct_DamDailyGraph_Output struct {
//	Dam           *model_dam.Struct_Dam `json:"dam"`
//	Dam_data_type json.RawMessage       `json:"dam_data_type"`
//	Dam_graph     *highchart.HighChart  `json:"dam_graph"`
//}

type Struct_DamDailyLastest_InputParam struct {
	Region_id   string `json:"region_id"`   // รหัสภาค  null=All
	Province_id string `json:"province_id"` // null=All
	Basin_id    string `json:"basin_id"`    // null=All
	Dam_date    string `json:"dam_date"`    // null=Lastest
	Agency_id   string `json:"agency_id"`   // agency id
	Start_date  string `json:"start_date"`  // วันที่ีเริ่มต้น
	End_date    string `json:"end_date"`    // วันที่สิ้นสุด
}

//type Struct_DamDailyLastest_Output struct {
//	Agency *model_agency.Struct_Agency    `json:"agency"`
//	Data   []*Struct_DamDaily             `json:"data"`
//	Graph  []*Struct_DamDailyGraph_Output `json:"graph"`
//}

type Struct_DamDailyLastest_OutputParam struct {
	Header []string            `json:"header"` // example:`["id","dam_date","dam_level","dam_storage","dam_storage_percent","dam_inflow","dam_inflow_acc_percent","dam_uses_water","dam_uses_water_percent","dam_released","dam_spilled","dam_losses","dam_evap"]` หัวตาราง
	Data   []*Struct_Dam_Daily `json:"data"`   // ข้อมูล
}

type Struct_Dam_Daily struct {
	Id                     int64       `json:"id"`                     // example:`67826` รหัสข้อมูล
	Dam_date               string      `json:"dam_date"`               // example:`2006-01-02` วันที่เก็บข้อมูล
	Dam_storage            interface{} `json:"dam_storage"`            // example:`6603.49` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	Dam_storage_percent    interface{} `json:"dam_storage_percent"`    // example:`49.05` เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.)
	Dam_inflow             interface{} `json:"dam_inflow"`             // example:`14.79` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	Dam_inflow_acc_percent interface{} `json:"dam_inflow_acc_percent"` // example:`30.83` เปอร์เซนต์ปริมาณน้ำไหลเทียบกับปริมาณน้ำไหลลงเขื่อนขนาดใหญ่เฉลี่ยรวมทั้งปี (%)
	Dam_uses_water         interface{} `json:"dam_uses_water"`         // example:`2803.49` ปริมาณน้ำที่ใช้ได้
	Dam_uses_water_percent interface{} `json:"dam_uses_water_percent"` // example:`20.83` เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.)
	Dam_level              interface{} `json:"dam_level"`              // example:`0` ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.)
	Dam_released           interface{} `json:"dam_released"`           // example:`3` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	Dam_spilled            interface{} `json:"dam_spilled"`            // example:`0` ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง
	Dam_losses             interface{} `json:"dam_losses"`             // example:`0` ปริมาณน้ำที่สูญเสีย
	Dam_evap               interface{} `json:"dam_evap"`               // example:`0` ปริมาณน้ำที่ระเหย
	SubBasin_id            int64       `json:"sub_basin_id,omitempty"` // example:`21` รหัสลุ่มน้ำย่อย

	Station_type string `json:"station_type,omitempty"` // example:`dam_daily` ชนิดของสถานี
}

type Struct_DamCctv struct {
	Cctv_id  int64  `json:"id"`  // example:`1` รหัสข้อมูล
	Cctv_url string `json:"url"` // example:`http://cctv1bb.egat.co.th/axis-cgi/mjpg/video.cgi` url ของ cctv
}

type Struct_DamDaily struct {
	Struct_Dam_Daily
	Dam     *model_dam.Struct_D           `json:"dam,omitempty"`     // เขื่อน
	Agency  *model_agency.Struct_Agency   `json:"agency,omitempty"`  // หน่วยงาน
	Basin   *model_basin.Struct_Basin     `json:"basin,omitempty"`   // ลุ่มน้ำ
	Geocode *model_geocode.Struct_Geocode `json:"geocode,omitempty"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Cctv    Struct_DamCctv                `json:"cctv,omitempty"`    // cctv อ่างเก็บน้ำ
}

// สำหรับหน้า thaiwater main ข้อมูลเขื่อนรวม 6 ภาค
type Struct_DamDailySummary struct {
	Dam_date               string                        `json:"dam_date"`               // example:`2006-01-02` วันที่เก็บข้อมูล
	Region_name            string                        `json:"region_name"`            // example:`ภาคเหนือ ภาคตะวันออกเฉพียงเหนือ ฯลฯ` ภาค
	Dam_storage            interface{}                   `json:"dam_storage"`            // example:`6603.49` ปริมาณน้ำกักเก็บปัจจุบัน สะสม (ล้าน ลบ.ม.)
	Dam_storage_percent    interface{}                   `json:"dam_storage_percent"`    // example:`49.05` เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  สะสม  (% รนก.)
	Dam_uses_water         interface{}                   `json:"dam_uses_water"`         // example:`2803.49` ปริมาณน้ำที่ใช้ได้ สะสม
	Dam_uses_water_percent interface{}                   `json:"dam_uses_water_percent"` // example:`20.83` เปอร์เซนต์ปริมาตรน้ำใช้การได้ สะสม  (% รนก.)
	Geocode                *model_geocode.Struct_Geocode `json:"geocode,omitempty"`      // ข้อมูลขอบเขตการปกครองของประเทศไทย
}

// สำหรับหน้า thaiwater main ข้อมูลน้ำใช้การของ 4 เขื่อนหลัก รายวัน
type Struct_DamDailySummary4Dam struct {
	Dam_date       string      `json:"dam_date"`               // example:`2006-01-02` วันที่เก็บข้อมูล
	Dam_storage    interface{} `json:"dam_storage"`            // example:`6603.49` ปริมาณน้ำกักเก็บปัจจุบัน สะสม (ล้าน ลบ.ม.)
	Dam_inflow     interface{} `json:"dam_storage_percent"`    // example:`49.05` ปริมาณน้ำไหลเข้าอ่าง (ล้าน ลบ.ม)
	Dam_uses_water interface{} `json:"dam_uses_water"`         // example:`2803.49` ปริมาณน้ำที่ใช้ได้ สะสม
	Dam_released   interface{} `json:"dam_uses_water_percent"` // example:`20.83` ปริมาณการระบาย (ล้าน ลบ.ม.)
}

type Struct_DamDailyNear struct {
	Struct_Dam_Daily
	Dam     *model_dam.Struct_D           `json:"dam,omitempty"`     // เขื่อน
	Agency  *model_agency.Struct_Agency   `json:"agency,omitempty"`  // หน่วยงาน
	Basin   *model_basin.Struct_Basin     `json:"basin,omitempty"`   // ลุ่มน้ำ
	Geocode *model_geocode.Struct_Geocode `json:"geocode,omitempty"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
}

type Struct_DamDaily_ErrorData struct {
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

type Struct_GetDamGraph_InputParam struct {
	Dam_id    string `json:"dam_id"`
	Dam_type  string `json:"dam_type"`
	Year      string `json:"year"`
	From_date string `json:"from_date"`
	To_date   string `json:"to_date"`
}

var DamDailyColumnForCheckErrordata = []string{
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

type GraphAnalystDamDailyInput struct {
	DataType string  `json:"data_type"` // ชื่อ field ในฐานข้อมูล เช่น dam_storage
	DamID    []int64 `json:"dam_id"`    // รหัสเขื่อน เช่น [1,2,3,4]
	Year     []int64 `json:"year"`      // ปีสำหรับดึงข้อมูล เช่น [2015,2016]
	Month    int64   `json:"month"`     // เดือน เช่น 3
	Day      int64   `json:"day"`       // วัน เช่น  24
}

type GraphAnalystDamDailyOutput struct {
	Year int64       `json:"year"` // example:`2016` ปี
	Date string      `json:"date"` // exmaple:`2016-03-24` วันที่
	Data interface{} `json:"data"` // example:`140` ข้อมูล
}

type GraphDamDailyInput struct {
	DataType string  `json:"data_type"` // ชื่อ field ในฐานข้อมูล เช่น dam_storage
	DamID    []int64 `json:"dam_id"`    // รหัสเขื่อน เช่น [1,2,3,4]
	Year     []int64 `json:"year"`      // ปีสำหรับดึงข้อมูล เช่น [2015,2016]
	Month    int64   `json:"month"`     // เดือน เช่น 3
	Day      int64   `json:"day"`       // วัน เช่น  24
}

type GraphDamDailyOutput struct {
	Year int64       `json:"year"` // example:`2016` ปี
	Date string      `json:"date"` // exmaple:`2016-03-24` วันที่
	Data interface{} `json:"data"` // example:`140` ข้อมูล
}

type MonitoringDamOutput struct {
	DamId             string          `json:"dam_id"`              // id เขื่อน
	DamName           json.RawMessage `json:"dam_name"`            // example:`{"th":"เขื่อนภูมิพล"}` ชื่อเขื่อน
	DamDate           string          `json:"dam_date"`            // example:`2006-01-02` วันที่เก็บข้อมูล
	DamStorage        interface{}     `json:"dam_storage"`         // example:`140` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	DamInflow         interface{}     `json:"dam_inflow"`          // example:`10` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	DamReleased       interface{}     `json:"dam_released"`        // example:`11` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	DamStoragePercent interface{}     `json:"dam_storage_percent"` // เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.)
	DamUsesWater      interface{}     `json:"dam_uses_water"`      // ปริมาณน้ำที่ใช้ได้
}

type Struct_DamGraph struct {
	Data        []*Struct_GraphData `json:"data"`                  // ข้อมูลกราฟ
	Max_storage float64             `json:"max_storage,omitempty"` // example:`13462` ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.]
	Min_storage float64             `json:"min_storage,omitempty"` // example:`3800` ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด  [ล้าน ลบ.ม.]
}

type Struct_GraphData struct {
	Date  string      `json:"date"`  // example:`2006-01-02` วันที่เก็บข้อมูล
	Value interface{} `json:"value"` // example:`140` ข้อมูล
}

type Struct_WaterInformation struct {
	Dam_id                 string `json:"dam_id"`                 // exmaple:`1` รหัสเขื่อน
	Dam_name               string `json:"dam_name"`               // example:`เขื่อนภูมิพล` ชื่อเขื่อน
	Dam_lat                string `json:"dam_lat"`                // example:`17.242` พิกัดละติจูดของเขื่อน
	Dam_long               string `json:"dam_long"`               // exmaple:`98.9728` พิกัดลองจิจูดของเขื่อน
	Dam_uses_water         string `json:"dam_uses_water"`         // exmaple:`247` ปริมาณน้ำที่ใช้การได้
	Dam_uses_water_percent string `json:"dam_uses_water_percent"` // exmaple:`2` ปริมาณน้ำที่ใช้การได้ คิดเป็น %
	Status_color           string `json:"status_color"`           // exmaple:`#FFC000` โค้ดสี
}

type Struct_WaterInformation_DamList struct {
	Dam_id                   string `json:"dam_id"`                   // example:`67826` รหัสเขื่อน
	Dam_name                 string `json:"dam_name"`                 // example:`เขื่อนภูมิพล` ชื่อเขื่อน
	Province_id              int64  `json:"province_id"`              // exmaple:`63` รหัสจังหวัด
	Province_name            string `json:"province_name"`            // example:`ตาก` ชื่อจังหวัด
	Dam_date                 string `json:"dam_date"`                 // example:`2016-05-25` วันที่
	Dam_storage              string `json:"dam_storage"`              // example:`4,047` ปริมาณน้ำเก็บกัก
	Dam_storage_percent      string `json:"dam_storage_percent"`      // example:`30%` ปริมาณน้ำเก็บกัก (%)
	Dam_storage_status_color string `json:"dam_storage_status_color"` // example:`#FFC000` โค้ดสี แสดงสถานะของเขื่อน
	Dam_uses_water           string `json:"dam_uses_water"`           // exmaple:`247` ปริมาณน้ำใช้การ
	Dam_uses_water_percent   string `json:"dam_uses_water_percent"`   // example:`2%` ปริมาณน้ำใช้การ (%)
	Dam_uses_status_color    string `json:"dam_uses_status_color"`    // exmaple:`#FFC000` โค้ดสี แสดงสถานะนำ้ใช้การของเขื่อน
	Dam_inflow               string `json:"dam_inflow"`               // exmaple:`0` ปริมาณน้ำไหลลงอ่าง
	Dam_inflow_acc_percent   string `json:"dam_inflow_acc_percent"`   // example:`0.95%` ปริมาณน้ำไหลลงอ่าง (%)
	Dam_released             string `json:"dam_released"`             // example:`4` ปริมาณน้ำระบาย
	Dam_lat                  string `json:"dam_lat"`                  // exmaple:`17.242000` พิกัดละติจูด
	Dam_long                 string `json:"dam_long"`                 // example:`98.972800` พิกัดลองจิจูด
	Normal_storage           string `json:"normal_storage"`           // example:`13,462` ปริมาตรน้ำที่ระดับเก็บกักปกติ [ล้าน ลบ.ม.]
	Dam_storage_status_level int64  `json:"dam_storage_status_level"` // example:`1` สถานะของเขื่อน
	Uses_water_percent       string `json:"uses_water_percent"`       // example:`2` ปริมาณน้ำใช้การ (%)
	Amphoe_name              string `json:"amphoe_name"`              // example:`สามเงา` อำเภอ
	District_name            string `json:"district_name"`            // example:`วังหมัน` ตำบล
}

type Struct_DamGraphHistory struct {
	Dt_date     string `json:"dt_date"`     // example:`2016-01-01` วันที่
	Dam_storage string `json:"dam_storage"` // example:`4927` ปริมาณน้ำกักเก็บ
}

type Struct_InputDamNear struct {
	Province_id string `json:"province_id"` // exmaple:`63` รหัสจังหวัด
}

type Struct_DamNear struct {
	Id_dam string `json:"id_name_near"` // exmaple:`1,12,11,36` เขื่อนใกล้เคียง
}
