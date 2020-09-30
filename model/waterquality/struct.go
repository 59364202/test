package waterquality

import (
	"encoding/json"
	"time"

	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_waterquality_station "haii.or.th/api/thaiwater30/model/waterquality_station"
)

type Struct_WaterQuality struct {
	Id                        int64       `json:"id,omitempty"`              // example:`13` รหัสข้อมูลคุณภาพน้ำ
	Waterquality_Id           int64       `json:"waterquality_id"`           // example:`109` รหัสสถานี
	Waterquality_Datetime     string      `json:"waterquality_datetime"`     // example:`2006-01-02 15:04` วันเวลาที่วัดคุณภาพน้ำ
	Waterquality_Do           interface{} `json:"waterquality_do"`           // example:`3.1` ออกซิเจนละลายในน้ำ หน่วย mg/l
	Waterquality_Ph           interface{} `json:"waterquality_ph"`           // example:`6.2` ความเป็นกรด-ด่าง
	Waterquality_Temp         interface{} `json:"waterquality_temp"`         // example:`null` อุณหภูมิน้ำ หน่วย C
	Waterquality_Turbid       interface{} `json:"waterquality_turbid"`       // example:`null` ค่าความขุ่นในน้ำ หน่วย NTU
	Waterquality_Bod          interface{} `json:"waterquality_bod"`          // example:`null` ค่าความสกปรกในรูปสารอินทรีย์ หน่วย mg/l
	Waterquality_Tcb          interface{} `json:"waterquality_tcb"`          // example:`null` ปริมาณแบคทีเรียในรูปโคลิฟอร์มทั้งหมด หน่วย MPN/100 ml
	Waterquality_Fcb          interface{} `json:"waterquality_fcb"`          // example:`null` ปริมาณแบคทีเรียในรูปฟีคลอโคลิฟอร์ม หน่วย MPN/100 ml
	Waterquality_Nh3n         interface{} `json:"waterquality_nh3n"`         // example:`null` ปริมาณแอมโมเนีย-ไนโตรเจน หน่วย mg/l
	Waterquality_Wqi          interface{} `json:"waterquality_wqi"`          // example:`null` ช่วงคะแนน WQI
	Waterquality_Ammonium     interface{} `json:"waterquality_ammonium"`     // example:`null` ปริมาณแอมโมเนีย
	Waterquality_Nitrate      interface{} `json:"waterquality_nitrate"`      // example:`null` ไนโตรเจน
	Waterquality_Colorstatus  string      `json:"waterquality_colorstatus"`  // example:`` สถานะของสี
	Waterquality_Status       string      `json:"waterquality_status"`       // example:`ปกติ` สถานะของคุณภาพน้ำ
	Waterquality_Salinity     interface{} `json:"waterquality_salinity"`     // example:`null` ค่าความเค็ม
	Waterquality_Conductivity interface{} `json:"waterquality_conductivity"` // example:`229.2` ความนำไฟฟ้าในน้ำ หน่วย uS/cm ชื่อเต็ม The Electrical Conductivity (ec)
	Waterquality_Tds          interface{} `json:"waterquality_tds"`          // example:`null` ค่า tds
	Waterquality_Chlorophyll  interface{} `json:"waterquality_chlorophyll"`  // example:`null` คลอโรฟิลด์

	Waterquality_Station *model_waterquality_station.Struct_WaterQualityStation `json:"waterquality_station,omitempty"` // สถานี

	Station_type string                    `json:"station_type,omitempty"` // example:`waterquality` ชนิดของสถานี
	Basin        *model_basin.Struct_Basin `json:"basin,omitempty"`        // ลุ่มน้ำ
}
type Param_WaterQualityCache struct {
	StationId     int64  `json:"station_id"`
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Data_Limit    int    `json:"data_limit"`    // จำนวนข้อมูล
	Basin_id      string `json:"basin_id"`      // required:false example:`1` รหัสลุ่มน้ำ ไม่ใส่ = ทุกลุ่มน้ำ เลือกได้หลายลุ่มน้ำ เช่น 1,2,4

	IsMain bool
}

type Param_WaterQualityGraph struct {
	Id       int64  `json:"id"`
	station  string `json:"station"`
	DateType string `json:"datatype"`
	//	DateStart     string `json:"datestart"`
	//	DateEnd       string `json:"dateend"`
	DateStart time.Time `json:"datestart"`
	DateEnd   time.Time `json:"dateend"`
}

type WaterQualityGraphCompareAnalystInput struct {
	WaterQualityStation []int64 `json:"waterquality_station_id"` // รหัสสถานีคุณภาพน้ำ เช่น [202,132]
	Param               string  `json:"param"`                   // ชื่อ field ที่ต้องการ เช่น ph
	DatetimeStart       string  `json:"start_datetime"`          // เวลาเริ่มต้น เช่น 2017-08-20 00:04
	DatetimeEnd         string  `json:"end_datetime"`            // เวลาสิ้นสุด เช่น 2017-08-20 15:04
}

type WaterQualityGraphCompareAnalystOutput struct {
	SeriesName json.RawMessage        `json:"series_name"` // ชื่อ
	Data       map[string]interface{} `json:"data"`        // ข้อมูล
}

type WaterQualityGraphCompareAnalystOutput2 struct {
	SeriesName json.RawMessage                        `json:"series_name"` // example:`{"th": "สถานีบ้านห้วยซัน"}` ชื่อ
	Data       []*WaterQualityGraphCompareAnalystData `json:"data"`        // ข้อมูล
}

type WaterQualityGraphCompareAnalystData struct {
	Datetime string      `json:"datetime"` // example:`2006-01-02 15:04` วันเวลา
	Value    interface{} `json:"value"`    // example:`23` ค่าคุณภาพน้ำ
}

type WaterQualityGraphOutputData struct {
	Name  string      `json:"name"`  // example:`2006-01-02 15:04` วันเวลา
	Value interface{} `json:"value"` // example:`23` ค่าคุณภาพน้ำ
}

type WaterQualityGraphParamsAnalystInput struct {
	WaterQualityStation int64    `json:"waterquality_station_id"` // รหัสสถานีคุณภาพน้ำ เช่น 202
	Param               []string `json:"param"`                   // ชื่อ field ที่ต้องการ เช่น [ph]
	DatetimeStart       string   `json:"start_datetime"`          // เวลาเริ่มต้น เช่น 2017-08-20 00:04
	DatetimeEnd         string   `json:"end_datetime"`            // เวลาสิ้นสุด เช่น 2017-08-20 15:04
}

type WaterQualityGraphParamsAnalystOutput struct {
	SeriesName string                 `json:"series_name"` // example:`series name`ชื่อ
	Data       map[string]interface{} `json:"data"`        // ข้อมูล
}

type WaterQualityGraphParamsAnalystOutput2 struct {
	SeriesName string                                 `json:"series_name"` // example:`series name`ชื่อ
	Data       []*WaterQualityGraphCompareAnalystData `json:"data"`        // ข้อมูล
}

type WaterQualityGraphWaterlevelAnalystInput struct {
	WaterQualityStation   int64  `json:"waterquality_station_id"` // รหัสสถานีคุณภาพ เช่น 132
	WaterlevelStation     int64  `json:"waterlevel_station_id"`   // รหัสสถานีระดับน้ำ เช่น 100
	DatetimeStart         string `json:"start_datetime"`          // เวลาเริ่มต้น 2017-08-20 00:04
	DatetimeEnd           string `json:"end_datetime"`            // เวลาเริ่มสิ้นสุด 2017-08-20 15:04
	ParamWQ               string `json:"param"`                   // ชื่อ field ที่ต้องการ เช่น ph
	WaterlevelStationType string `json:"waterlevel_station_type"` // ชื่อฐานข้อมูลระดับน้ำ tele_waterlevel หรือ หcanal_waterlevel
}

type WaterQualityWaterlevelOutput struct {
	SeriesName json.RawMessage                `json:"series_name"` // example:`series name`ชื่อ
	Data       []*WaterQualityGraphOutputData `json:"data"`        // ข้อมูล
}

type WaterQualityGraphCompareDatetimeInput struct {
	WaterQualityStation []int64  `json:"waterquality_station_id"` // รหัสสถานีคุณภาพน้ำ [23,53,6,7]
	Param               string   `json:"param"`                   // ชื่อ field ที่ต้องการ เช่น ph,do
	Date                []string `json:"date"`                    // วันที่ [2006-01-02,2007-01-02]
}

type WaterQualityCompareDatetimeAnalystOutput struct {
	SeriesName string                 `json:"series_name"` // ชื่อ
	Data       map[string]interface{} `json:"data"`        // ข้อมูล
}

type WaterQualityGraphAnalystOutputDataLang struct {
	TH string `json:"th"`
	EN string `json:"en"`
	JP string `json:"jp"`
}

type WaterQualityCompareDatetimeOutput struct {
	Datetime string                                      `json:"datetime"` // example:`2016-01-02 15:04` วันเวลา
	Station  []*WaterQualityCompareDatetimeOutputStation `json:"station"`  // สถานี
}

type WaterQualityCompareDatetimeOutputStation struct {
	Station json.RawMessage `json:"station_name"` // example:`{"th":"สถานีคุณภาพน้ำ"}`ชื่อสถานี
	Value   interface{}     `json:"value"`        // example:`8.7` ค่าคุณภาพน้ำ
}

type MonitoringWaterqualityOutput struct {
	StationName     json.RawMessage `json:"station_name"`          // example:`{"th":"สถานีคุณภาพน้ำ"}`ชื่อสถานี
	StationDate     string          `json:"waterquality_datetime"` // example:`2016-01-02 15:04` วันเวลาที่วัดคุณภาพน้ำ
	StationSalinity interface{}     `json:"waterquality_salinity"` // example:`8.7`ค่าความเค็ม
}

type Struct_WaterqualityLatest struct {
	Station_id       string `json:"station_id"`       // example:`206` รหัสสถานี
	Station_name     string `json:"station_name"`     // example:`วัดมะขาม` ชื่อสถานี
	Station_district string `json:"station_district"` // example:`เมือง` ชื่ออำเภอ
	Station_zone     string `json:"station_zone"`     // example:`ปทุมธานี` ชื่อจังหวัด
	Station_province string `json:"station_province"` // example:`แม่น้ำเจ้าพระยา` ชื่อจังหวัด
	Station_location string `json:"station_location"` // example:`แม่น้ำเจ้าพระยา` ชื่อแม่น้ำ
	Dt_date          string `json:"dt_date"`          // example:`2017-07-20` วันที่
	Dt_time          string `json:"dt_time"`          // example:`16:00:00` เวลา
	Dt_ph            string `json:"dt_ph"`            // example:`7.3` ความเป็นกรด-ด่าง
	Dt_salinity      string `json:"dt_salinity"`      // example:`0.09` ค่าความเค็ม (ppt)
	Dt_turbidity     string `json:"dt_turbidity"`     // example:`124` ความขุ่น
	Dt_conductivity  string `json:"dt_conductivity"`  // example:`224` ความนำไฟฟ้า
	Dt_tds           string `json:"dt_tds"`           // example:`130`
	Dt_chlorophyll   string `json:"dt_chlorophyll"`   // example:`9.3` คลอโรฟิลล์
	Dt_do            string `json:"dt_do"`            // example:`3.94` ปริมาณออกซิเจนละลาย (mg/l)
	Dt_ec            string `json:"dt_ec"`            // example:`224` ค่าความนำไฟฟ้า (µS/cm)
	Dt_temp          string `json:"dt_temp"`          // example:`30.23` ค่าอุณหภูมิ (oC)
	Dt_depth         string `json:"dt_depth"`         // example:`1.5` ระดับน้ำ
	Station_lat      string `json:"station_lat"`      // example:`14.00409444` ละติจูด
	Station_long     string `json:"station_long"`     // example:`100.5406944` ลองติจูด
	Data_source      string `json:"data_source"`      // example:`กปน.` แหล่งข้อมูล
	Province_id      string `json:"province_id"`      // example:`13` รหัสจังหวัด
	Is_active        string `json:"is_active"`        // example:'Y'
}

//struct สำหรับ เช็คการแสดงผลว่า สถานีไหน ให้ซ่อนหรือแสดงผลข้อมูลไหน
type WaterQualityshowStatusStruct struct {
	Do           bool `json:"do"`
	Conductivity bool `json:"conductivity"`
	Ph           bool `json:"ph"`
	Temp         bool `json:"temp"`
	Turbid       bool `json:"turbid"`
	Bod          bool `json:"bod"`
	Tcb          bool `json:"tcb"`
	Fcb          bool `json:"fcb"`
	Nh3n         bool `json:"nh3n"`
	Wqi          bool `json:"wqi"`
	Ammonium     bool `json:"ammonium"`
	Nitrate      bool `json:"nitrate"`
	Salinity     bool `json:"salinity"`
	Tds          bool `json:"tds"`
	Chlorophyll  bool `json:"chlorophyll"`
	Colorstatus  bool `json:"colorstatus"`
}
