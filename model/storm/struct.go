package storm

import ()

type Struct_Strom struct {
	Storm_Name string               `json:"storm_name,omitempty"` // example:`BANYAN-17` ชื่อพายุ
	Storm_Data []*Struct_Strom_Data `json:"storm_data"`           // ข้อมูลพายุ
}

type Struct_Strom_Data struct {
	Storm_Datetime       string      `json:"storm_datetime"`       // example:`2006-01-02 15:04` วันเวลาที่เก็บข้อมูลเส้นทางพายุ
	Storm_Lat            interface{} `json:"storm_lat"`            // example:`19.6` ละติจูด
	Storm_Direction_Lat  interface{} `json:"storm_direction_lat"`  // example:`N` ทิศทาง N หรือ S
	Storm_Long           interface{} `json:"storm_long"`           // example:`165.2` ลองติจูด
	Storm_Direction_Long interface{} `json:"storm_direction_long"` // example:`E` ทิศทาง E หรือ W
	Storm_Name           string      `json:"storm_name,omitempty"` // example:`BANYAN-17` ชื่อพายุ
	Storm_Pressure       string      `json:"storm_pressure"`       // example:`982` ความกดอากาศ mb
	Storm_Wind           string      `json:"storm_wind"`           // example:`34 kt` ลมพายุที่หนักที่สุดในระยะ (kt)
}

// ----- storm history -----
// api params
type StructStormHistoryParam struct {
	Name	string 		`json:"name"`   // example:`BANYAN-17` คำค้นหา
	Month	int 		`json:"month"`	// example:`5` เดือน 1-12
	Year	int 		`json:"year"` 	// example:`2019` ปี ค.ศ 4 หลัก
}

// storm name that occur in given criteria
type StructStormPeriod struct {
	StormName           string      `json:"name"` // example:`BANYAN-17` ชื่อพายุ
}

// storm history for single storm
type StructStormHistory struct {
	StormName	string      				`json:"name"`	// example:`BANYAN-17` ชื่อพายุ
	Data   		[]*StructStormHistoryData 	`json:"data"`   // ข้อมูลเส้นทาง
}

// storm history data
type StructStormHistoryData struct {
	Id              	int64       	`json:"id"`             // example:`114497665` รหัสข้อมูลปริมาณน้ำฝน
	StormDatetime       string      	`json:"datetime"`       // example:`2006-01-02 15:04` วันเวลาที่เก็บข้อมูลเส้นทางพายุ
	StormLat            interface{} 	`json:"lat"`            // example:`19.6` ละติจูด
	StormLong           interface{} 	`json:"long"`           // example:`165.2` ลองติจูด
	StormDirectionLat   interface{} 	`json:"direction_lat"`  // example:`N` ทิศทาง N หรือ S
	StormDirectionLong 	interface{} 	`json:"direction_long"` // example:`E` ทิศทาง E หรือ W
	StormPressure       interface{}     `json:"pressure"`       // example:`982` ความกดอากาศ mb
	StormWind           interface{}     `json:"wind"`           // example:`34 kt` ลมพายุที่หนักที่สุดในระยะ (kt)
	StormColor          string     		`json:"color"`          // example:`#000` รหัสสีตามความแรงพายุ
	StormToLat          interface{} 	`json:"to_lat"`         // example:`19.6` ละติจูดปลายทาง
	StormToLong         interface{} 	`json:"to_long"`        // example:`165.2` ลองติจูดปลายทาง
}
// ----- end storm history -----