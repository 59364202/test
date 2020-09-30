package swan

import (
	"haii.or.th/api/thaiwater30/model/swan_station"
)

type Struct_Swan struct {
	//	Id       int64       `json:"id"`                 // example:`1` รหัสข้อมูลคาดการณ์ความสูงคลื่น
	//	Datetime string      `json:"datetime"` // example:`2006-01-02 15:04` วันเวลาที่เก็บข้อมูลคาดการณ์ความสูงคลื่น
	Highsig interface{} `json:"highsig"` // example:`0.11235` ข้อมูลคาดการณ์ความสูงของคลื่น หน่วย m

	Swan_Station *swan_station.Struct_SwanStation `json:"swan_station,omitempty"` // สถานี
}

type Struct_WaveForecast struct {
	Date   string                        `json:"date"`   // example:`25 พ.ค. 2559` วันที่
	Region []*Struct_WaveForecast_Region `json:"region"` // array ของข้อมูลคลื่นในแต่ละภาค
}
type Struct_WaveForecast_Region struct {
	Region_id    string `json:"region_id"`    // example:`1` รหัสภาค
	Region_name  string `json:"region_name"`  // example:`อ่าวไทยฝั่งะวันออก` ภาค
	Level        string `json:"level"`        // exmaple:`3` รหัสสถานะ
	Status       string `json:"status"`       // example:`คลื่น สูงกว่า 2 เมตร` สถานะ
	Station_name string `json:"station_name"` // example:`ที่เกาะช้าง ตราด` สถานี
	Color        string `json:"color"`        // example:`#FF0000` สี
}
