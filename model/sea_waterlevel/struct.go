package sea_waterlevel

import (
	"encoding/json"
)

type SeaWaterlevelLatestSwagger struct {
	Name json.RawMessage `json:"name"`                // example:`{"th":"สถานีวัดระดับน้ำทะเล"}`ชื่อสถานี
	Data interface{}     `json:"sea_waterlevel_data"` // example:`145`ระดับน้ำ
}

type SeaWaterlevelLatest struct {
	Name *SeaTagName `json:"name"`                // ชื่อสถานี
	Data interface{} `json:"sea_waterlevel_data"` // example:`145`ระดับน้ำ
}

type SeaWaterlevelLatestStation struct {
	Data []*SeaWaterlevelData `json:"data"` // ข้อมูล
	//	Data2 []*SeaWaterlevelData2 `json:"data,ommitempty"`
}

type SeaWaterlevelData2 struct {
	Station      *SeaStation    `json:"sea_station"`              // สถานี
	Agency       *Agency        `json:"agency"`                   // หน่วยงาน
	Geocode      *Geocode       `json:"geocode"`                  // ข้อมูลขอบเขตการปกครองของประเทศไทย
	MaxForecasat interface{}    `json:"max_seaforecast_datetime"` // วันเวลาข้อมูลที่มากที่สุด
	MinForecast  interface{}    `json:"min_seaforecast_datetime"` // วันเวลาข้อมูลที่น้อยที่สุด
	Seaforecast  []*SeaForecast `json:"sea_water_forecast"`       // ข้อมูลทำนายระดับน้ำ
}

type SeaForecast struct {
	Datetime string      `json:"seaforecast_datetime"` // วันเวลาเก็บข้อมูลคาดการณ์น้ำท่วม
	Value    interface{} `json:"seaforecast_value"`    // ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s)
}

type SeaWaterlevelData struct {
	Station  *SeaStation `json:"sea_station"`         // สถานี
	Agency   *Agency     `json:"agency"`              // หน่วยงาน
	Geocode  *Geocode    `json:"geocode"`             // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Datetime string      `json:"waterlevel_datetime"` // วันเวลาของค่าระดับน้ำ
	Value    interface{} `json:"waterlevel_value"`    // ระดับน้ำ
}

type SeaWaterlevelData3 struct {
	Datetime string      `json:"datetime"` // example:`2006-01-02 15:04`วันเวลาของค่าระดับน้ำ
	Value    interface{} `json:"value"`    // example:`2`ระดับน้ำ
}

type SeaStation struct {
	ID      interface{}      `json:"id"`                  // รหัสสถานี
	Name    *json.RawMessage `json:"sea_station_name"`    // ชื่อสถานี
	Lat     interface{}      `json:"sea_station_lat"`     // ละติจูดของสถานี
	Long    interface{}      `json:"sea_station_long"`    // ลองติจูดของสถานี
	OldCode string           `json:"sea_station_oldcode"` // รหัสสถานีเดิม
}

type Agency struct {
	ID        interface{}     `json:"id"`               // example:`9`รหัสหน่วยงาน
	Name      json.RawMessage `json:"agency_name"`      // example:`{"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)","en":"Hydro and Agro Informatics Institute","jp":""}`ชื่อหน่วยงาน
	ShortName json.RawMessage `json:"agency_shortname"` // example:`{"th":"สสนก.","en":"HAII","jp":""}`ชื่อย่อหน่วยงาน
}

type Geocode struct {
	ID           interface{}     `json:"id"`            // รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Area         interface{}     `json:"area_code"`     // รหัสภาค
	AreaName     json.RawMessage `json:"area_name"`     // ชื่อภาค
	Province     interface{}     `json:"province_code"` // รหัสจังหวัด
	ProvinceName json.RawMessage `json:"province_name"` // ชื่อจังหวัด
	Amphoe       interface{}     `json:"amphoe_code"`   // รหัสอำเภอ
	AmphoeName   json.RawMessage `json:"amphoe_name"`   // ชื่ออำเภอ
	Tumbon       interface{}     `json:"tumbon_code"`   // รหัสตำบล
	TumbonName   json.RawMessage `json:"tumbon_name"`   // ชื่อตำบล
}

type SeaTagName struct {
	TH string `json:"th"` // example:`ชื่อภาษาไทย` ภาษาไทย
	EN string `json:"en"` // example:`english` english
	JP string `json:"jp"` // example:`japan` japan
}

type SeaWaterlevelInput struct {
	StationID []int64 `json:"station_id"` // รหัสสถานีระดับน้ำทะเล [23,413,243]
	StartDate string  `json:"start_date"` // วันที่เริ่มต้น 2006-01-02
	EndDate   string  `json:"end_date"`   // วันที่สิ้นสุด 2006-01-02
}

type SeaWaterlevelOutput struct {
	SeriesName *json.RawMessage      `json:"sea_station_name,omitempty"` // example:`{"th":"สถานีระดับน้ำทะเล"}`ชื่อสถานี
	Data       []*SeaWaterlevelData3 `json:"sea_waterlevel_data"`        // ข้อมูลระดับน้ำ
}
