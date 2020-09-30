package api_service

import ()

type Metadata_Media struct {
	Agency_id      int64  `json:"agency_id"`      // example:`50` รหัสหน่วยงาน agency's serial number
	Media_type_id  int64  `json:"media_type_id"`  // example:`141` รหัสแสดงชนิดข้อมูลสื่อ
	Media_datetime string `json:"media_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บข้อมูลสื่อ record date
	Media_path     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/image?image=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJBMDrGmg4sd2EqTleGoMGzbfRwOOn9GdVm9blDTBH42TRFzCh4Sws-QyEtcntJfW62mIK0Q==` ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	Filename       string `json:"filename"`       // example:`00Latest.jpg` ชื่อไฟล์ของข้อมูลสื่อ file name
	Media_desc     string `json:"media_desc"`     // example:`ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi` รายละเอียดของข้อมูลสื่อ description
	Refer_source   string `json:"refer_source"`   // example:`http://weather.is.kochi-u.ac.jp/SE/00Latest.jpg` แหล่งข้อมูลสำหรับอ้างอิง reference source
}

type Metadata_Tele_waterlevel struct {
	Tele_station_id     int64   `json:"tele_station_id"`     // example:123 รหัสสถานีโทรมาตร
	Waterlevel_datetime string  `json:"waterlevel_datetime"` // วันที่ตรวจสอบค่าระดับน้ำ
	Waterlevel_msl      float64 `json:"waterlevel_msl"`      // ระดับน้ำ รทก
}
type Metadata_Canal_waterlevel struct {
	Canal_station_id          int64   `json:"canal_station_id"`          // รหัสสถานีวัดระดับน้ำในคลอง
	Canal_waterlevel_datetime string  `json:"canal_waterlevel_datetime"` // วันที่วัดระดับน้ำในคลอง
	Canal_waterlevel_value    float64 `json:"canal_waterlevel_value"`    // ค่าระดับน้ำในคลอง
	Comm_status               string  `json:"comm_status"`               // สถานะของเครื่องวัด
}

type Metadata_Temperature struct {
	Tele_station_id int64   `json:"tele_station_id"` // รหัสสถานีโทรมาตร
	Temp_datetime   string  `json:"temp_datetime"`   // วันที่เก็บค่าอุณหภูมิ
	Temp_value      float64 `json:"temp_value"`      // ค่าอุณหภูมิ
}

type Metadata_Pressure struct {
	Tele_station_id   int64   `json:"tele_station_id"`   // รหัสสถานีโทรมาตร
	Pressure_datetime string  `json:"pressure_datetime"` // วันที่เก็บค่าความกดอากาศ
	Pressure_value    float64 `json:"pressure_value"`    // ค่าความกดอากาศ
}

type Metadata_Humid struct {
	Tele_station_id int64   `json:"tele_station_id"` // รหัสสถานีโทรมาตร
	Humid_datetime  string  `json:"humid_datetime"`  // วันที่เก็บค่าความชื้นสัมพัทธ์
	Humid_value     float64 `json:"humid_value"`     // ค่าความชื้นสัมพัทธ์
}

type Metadata_Solar struct {
	Tele_station_id float64 `json:"tele_station_id"` // รหัสสถานีโทรมาตร
	Solar_datetime  string  `json:"solar_datetime"`  // วันที่เก็บค่าความเข้มแสง
	Solar_value     float64 `json:"solar_value"`     // ค่าความเข้มแสง
}

type Metadata_Floodforecast struct {
	Floodforecast_station_id int64   `json:"floodforecast_station_id"` // รหัสสถานีคาดการณ์น้ำท่วม
	Floodforecast_datetime   string  `json:"floodforecast_datetime"`   // วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	Floodforecast_value      float64 `json:"floodforecast_value"`      // ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}
