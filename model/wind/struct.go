//     Author: Thitiporn  Meeprasert <thitiporn@hii.or.th>
package wind

import (
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

// param for input query condition
type Param_Provinces struct {
	Region_Code     string `json:"region_code"`     // required:false example:`1` รหัสภาค ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Region_Code_tmd string `json:"region_code_tmd"` // required:false example:`1` รหัสภาค 6 ภาคตามการแบ่งของกรมอุตุฯ ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Province_Code   string `json:"province_code"`   // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Data_Limit      int    `json:"data_limit"`      // จำนวนข้อมูล
	//IsIgnore       bool   `json:"is_ignore"`
	Order string `json:"order"`
}

type Struct_Wind struct {
	WindDatetime string      `json:"wind_datetime,omitempty"`  // example:`2006-01-02` วันที่
	WindSpeed    interface{} `json:"wind_speed,omitempty"`     // example:`14.8` ค่าความเร็วลม
	WindDirValue interface{} `json:"wind_dir_value,omitempty"` // example:`270` ค่าองศาของทิศทางลม (degree)
	WindDir      interface{} `json:"wind_dir,omitempty"`       // example:`NE` ทิศทางลม

	Agency  *model_agency.Struct_Agency   `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *Struct_TeleStation           `json:"station"` // สถานี
}

type Struct_TeleStation struct {
	Id      int64           `json:"id"`                          // example:`1` รหัสสถานี
	Name    json.RawMessage `json:"tele_station_name,omitempty"` // example:`{"th":"คลองลาดพร้าว วัดบางบัว"}` ชื่อสถานี
	Lat     interface{}     `json:"tele_station_lat"`            // example:`13.854100` ละติจูดของสถานี
	Long    interface{}     `json:"tele_station_long"`           // example:`100.588000` ลองติจูดของสถานี
	OldCode string          `json:"tele_station_oldcode"`        // example:`BKK021` รหัสสถานีเดิม
}
