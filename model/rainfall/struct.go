package rainfall

import (
	"encoding/json"
)

type GetRainGraphParam struct {
	Id           int64  `json:"id"`            // รหัสสถานี
	Datatype     string `json:"datatype"`      // ประเภทข้อมูล
	Provice_code string `json:"province_code"` // จังหวัด
}

type RainfallStruct struct {
	Id              int64       `json:"id"`              // example:`114497665` รหัสข้อมูลปริมาณน้ำฝน
	Tele_station_id int64       `json:"tele_station_id"` // example:`17` รหัสสถานี
	Rainfall_date   string      `json:"rainfall_date"`   // example:`2006-01-02 15:04` วันที่เก็บปริมาณน้ำฝน
	Rainfall5m      interface{} `json:"rainfall5m"`      // example:`null` ปริมาณน้ำฝนทุก 5 นาที (mm.)
	Rainfall10m     interface{} `json:"rainfall10m"`     // example:`0.2` ปริมาณน้ำฝนทุก 10 นาที (mm.)
	Rainfall15m     interface{} `json:"rainfall15m"`     // example:`null` ปริมาณน้ำฝนทุก 15 นาที (mm.)
	Rainfall30m     interface{} `json:"rainfall30m"`     // example:`null` ปริมาณน้ำฝนทุก 30 นาที (mm.)
	Rainfall1h      interface{} `json:"rainfall1h"`      // example:`0.2` ปริมาณน้ำฝนทุก 1 ชั่วโมง (mm.)
	Rainfall3h      interface{} `json:"rainfall3h"`      // example:`null` ปริมาณน้ำฝนทุก 3 ชั่วโมง (mm.)
	Rainfall6h      interface{} `json:"rainfall6"`       // example:`null` ปริมาณน้ำฝนทุก 6 ชั่วโมง (mm.)
	Rainfall12h     interface{} `json:"rainfall12h"`     // example:`null` ปริมาณน้ำฝนทุก 12 ชั่วโมง (mm.)
	Rainfall24h     interface{} `json:"rainfall24h"`     // example:`0.2` ปริมาณน้ำฝนทุก 24 ชั่วโมง (mm.)
	Rainfall_acc    interface{} `json:"rainfall_acc"`    // example:`613.400024` ปริมาณน้ำฝนสะสม
}

type Struct_Rainfall_ErrorData struct {
	ID              int64           `json:"id"`                    // example:`63466119` รหัสข้อมูลปริมาณน้ำฝน
	StationID       int64           `json:"station_id"`            // example:`634` รหัสสถานี
	StationOldCode  string          `json:"station_oldcode"`       // example:`VLGE05` รหัสสถานีเดิม
	Datetime        string          `json:"datetime"`              // example:`2006-01-02` วันที่เก็บปริมาณน้ำฝน
	StationName     json.RawMessage `json:"station_name"`          // example:`{"th":"ห้วยปลาหลด"}` ชื่อสถานี
	ProvinceName    json.RawMessage `json:"station_province_name"` // example:`{"th":"ตาก"}`ชื่อจังหวัด
	AgencyName      json.RawMessage `json:"agency_name"`           // example:`{"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)","en":"Hydro and Agro Informatics Institute"}` ชื่อหน่วยงาน
	AgencyShortName json.RawMessage `json:"agency_shortname"`      // example:`{"en":"HAII"}` ชื่อย่อของหน่วยงาน
	Rainfall5m      interface{}     `json:"rainfall5m"`            // example:`null` ปริมาณน้ำฝนทุก 5 นาที (mm.)
	Rainfall10m     interface{}     `json:"rainfall10m"`           // example:`999999` ปริมาณน้ำฝนทุก 10 นาที (mm.)
	Rainfall15m     interface{}     `json:"rainfall15m"`           // example:`null` ปริมาณน้ำฝนทุก 15 นาที (mm.)
	Rainfall30m     interface{}     `json:"rainfall30m"`           // example:`null` ปริมาณน้ำฝนทุก 30 นาที (mm.)
	Rainfall1h      interface{}     `json:"rainfall1h"`            // example:`999999` ปริมาณน้ำฝนทุก 1 ชั่วโมง (mm.)
	Rainfall3h      interface{}     `json:"rainfall3h"`            // example:`null` ปริมาณน้ำฝนทุก 3 ชั่วโมง (mm.)
	Rainfall6h      interface{}     `json:"rainfall6h"`            // example:`null` ปริมาณน้ำฝนทุก 6 ชั่วโมง (mm.)
	Rainfall12h     interface{}     `json:"rainfall12h"`           // example:`null` ปริมาณน้ำฝนทุก 12 ชั่วโมง (mm.)
	Rainfall24h     interface{}     `json:"rainfall24h"`           // example:`999999` ปริมาณน้ำฝนทุก 24 ชั่วโมง (mm.)
	RainfallAcc     interface{}     `json:"rainfall_acc"`          // example:`999999` ปริมาณน้ำฝนสะสม
}

type Rainfall_InputParam struct {
	Id         string `json:"id"`         //
	Station_id string `json:"station_id"` // รหัสสถานี
	Start_date string `json:"start_date"` // วันที่เริ่มต้น
	End_date   string `json:"end_date"`   // วันที่สิ้นสุด
	Agency_id  string `json:"agency_id"`  // รหัสหน่วยงาน
}

type GetRainfallLastest_OutputParam struct {
	Data   []*RainfallStruct `json:"data"`   // ข้อมูล
	Header []string          `json:"header"` // example:`["id","rainfall_date","rainfall5m","rainfall10m","rainfall15m","rainfall30m","rainfall1h","rainfall3h","rainfall6h","rainfall12h","rainfall24h","rainfall_acc"]` หัวตาราง
}

type GetAdvRainMonthStationGraphInput struct {
	StationID int64   `json:"tele_station_id"` // รหัสสถานีโทรมาตร เช่น 1,23,5,67
	Year      []int64 `json:"year"`            //ปี เช่น [2015,2016]
}

type GetAdvRainMonthAreaGraphInput struct {
	AreaCode string  `json:"region_code"` // รหัส ภาค 02
	Year     []int64 `json:"year"`        // ปี [2014,2015]
	BaseLine string  `json:"normal"`      // baseline เช่น 48,30
}

type GetAdvRainAreaOutput struct {
	GraphData []*GetAdvRainMonthStationGraphOutput `json:"graph_data"` // ข้อมูล
	Baseline  []*GetAdvRainMonthStationGraphOutput `json:"normal"`     // เส้น baseline
}

type GetAdvRainMonthStationGraphOutput struct {
	Year    int64                                `json:"year"`         // example:`2015`ปี
	GSeries []*GetAdvRainMonthStationMonthOutput `json:"graph_series"` // ข้อมูล
}

type GraphData struct {
	GData []*GetAdvRainMonthStationGraphOutput `json:"graph_data"` // ข้อมูล
}

type GraphDataYearly struct {
	GData    []*GetAdvRainYearlyOutput `json:"graph_data"` // ข้อมูล
	Normal30 interface{}               `json:"normal_30"`  // example:`1466.93` เปรียบเทียบ 30 ปี
	Normal48 interface{}               `json:"normal_48"`  // example:`1374` เปรียบเทียบ 48 ปี
}

type GetAdvRainMonthStationMonthOutput struct {
	Month int64       `json:"month"` // example:`1` เดือน
	Value interface{} `json:"value"` // example:`41.2` ค่า
}

type GetAdvRainMonthGraphInput struct {
	StationID int64   `json:"tele_station_id"` // รหัสสถานี เช่น 3450
	Month     []int64 `json:"month"`           // เดือน [1,2,3,4]
	StartYear int64   `json:"start_year"`      // ปีเริ่มต้น 2014
	EndYear   int64   `json:"end_year"`        // ปีสิ้นสุด 2015
}

type GetAdvRainYearlyGraphInput struct {
	StationID int64 `json:"tele_station_id"` // รหัสสถานี เช่น 2340
	StartYear int64 `json:"start_year"`      // ปีเริ่มต้น 2014
	EndYear   int64 `json:"end_year"`        // ปีสิ้นสุด 2015
}

type GetAdvRainYearlyOutput struct {
	Year  int64       `json:"year"`  // example:`2006` ปี
	Value interface{} `json:"value"` // example:`200` ค่า
}

var RainfallColumnForCheckErrordata = []string{
	"id",
	"station_oldcode",
	"datetime",
	"station_name",
	"station_province_name",
	"agency_name",
	"agency_shortname",
	"rainfall5m",
	"rainfall10m",
	"rainfall15m",
	"rainfall30m",
	"rainfall1h",
	"rainfall3h",
	"rainfall6h",
	"rainfall12h",
	"rainfall24h",
	"rainfall_acc",
}
