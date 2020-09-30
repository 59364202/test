// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package latest_media is a model for cache.latest_media This table store latest_media information.
package latest_media

import (
	"database/sql"
	"time"
)

type Struct_Media struct {
	Path     string `json:"media_path"`     // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHNIfCEaLG3zGgZ2LEuTaHsMjGebDH8UN` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	Filename string `json:"filename"`       // example:`cri240_201612090100.gif` ชื่อไฟล์ของข้อมูลสื่อ
	FilePath string `json:"file_path"`      // example:`product/image/radar/cri/tmd/media/2016/12/09` ที่อยู่ของไฟล์ข้อมูลสื่อ
	Datetime string `json:"media_datetime"` // example:`2016-12-09 01:00` วันที่เก็บข้อมูลสื่อ

	PathThumb     interface{} `json:"media_path_thumb"` // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHMUFFJe41qS7RNBblHIQeFyeCkGeCBu9dWbFXBQC` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	FilenameThumb interface{} `json:"filename_thumb"`   // example:`thumb-cri240_201612090100.gif` ชื่อไฟล์ thumb ของข้อมูลสื่อ

	Dt time.Time `json:"-"` // วันที่ ใช้สำหรับนำไปใช้ต่อใน go
}

type Struct_UpdateCacheMedia struct {
	Id              sql.NullInt64  `json:"id"`
	AgencyId        sql.NullInt64  `json:"agency_id"`
	MediaTypeId     sql.NullInt64  `json:"media_type_id"`
	MediaDatetime   sql.NullString `json:"media_datetime"`
	MediaPath       sql.NullString `json:"media_path"`
	MediaDesc       sql.NullString `json:"media_desc"`
	Filename        sql.NullString `json:"filename"`
	ReferSource     sql.NullString `json:"refer_source"`
	Remark          sql.NullString `json:"remark"`
	DataId          sql.NullInt64  `json:"data_id"`
	Index           sql.NullString `json:"index"`
	DataimportLogId sql.NullInt64  `json:"dataimport_log_id"`
}

type Struct_Radar struct {
	RadarType          string      `json:"radar_type"`           // example:`cri240` ชนิดเรดาร์
	RadarName          string      `json:"radar_name"`           // example:`เรดาห์ที่เชียงราย รัศมี 240 กม.` ชื่อเรดาร์
	Agency	           string      `json:"agency"`		         // example:`tmd` ชื่อหน่วยงาน
	Timezone	       string      `json:"timezone"`       	 	 // example:`GMT` zone เวลา
	MediaDatetime      string      `json:"media_datetime"`       // example:`2016-12-09 01:00` วันที่เก็บข้อมูลสื่อ
	FilePath           string      `json:"file_path"`            // example:`product/image/radar/cri/tmd/media/2016/12/09` ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename           string      `json:"filename"`             // example:`cri240_201612090100.gif` ชื่อไฟล์ของข้อมูลสื่อ
	MediaPath          string      `json:"media_path"`           // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHNIfCEaLG3zGgZ2LEuTaHsMjGebDH8UN` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	FilenameThumb      interface{} `json:"filename_thumb"`       // example:`thumb-cri240_201612090100.gif` ชื่อไฟล์ thumb ของข้อมูลสื่อ
	MediaPathThumb     interface{} `json:"media_path_thumb"`     // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHMUFFJe41qS7RNBblHIQeFyeCkGeCBu9dWbFXBQC` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	FilePathAnimation  interface{} `json:"file_path_animation"`  // ที่อยู่ของไฟล์ animation ข้อมูลสื่อ
	FilenameAnimation  interface{} `json:"filename_animation"`   // ชื่อไฟล์ animation ของข้อมูลสื่อ
	MediaPathAnimation interface{} `json:"media_path_animation"` // ลิ้งค์ของไฟล์ animation ข้อมูลสื่อ
}

type RadarProvincesInput struct {
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด
}

type Param_RadarProvincesCache struct {
	StationId int64 `json:"station_id"`
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Data_Limit    int    `json:"data_limit"`    // จำนวนข้อมูล

	IsMain bool
}

type RainforcaseProvincesInput struct {
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด
}

type Struct_Media_Rainforcase struct {
	FilePath string `json:"file_path"`      // example:`product/image/radar/cri/tmd/media/2016/12/09` ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename string `json:"filename"`       // example:`cri240_201612090100.gif` ชื่อไฟล์ของข้อมูลสื่อ
	Path     string `json:"media_path"`     // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHNIfCEaLG3zGgZ2LEuTaHsMjGebDH8UN` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	Datetime string `json:"media_datetime"` // example:`2016-12-09 01:00` วันที่เก็บข้อมูลสื่อ
	ProvinceId string `json:"province_id"` // example:`10` รหัสจังหวัด
	}

type Struct_Media_Rainforcase7day struct {
	ProvinceId 			string `json:"province_id"` 		// example:`10` รหัสจังหวัด
	Province_name  		string `json:"province_name"`  		// example:`สระบุรี` ชื่อจังหวัด
	Forecast_day  		string `json:"forecast_day"`  		// example:`1` วัน
	Forecast_datetime  	string `json:"forecast_datetime"`  	// example:`2017-01-14 00:00:00` วันที่คาดการณ์
	Rainfall       		string `json:"rainfall"`       		// example:`0.0253` ปริมาณฝน
	Rainfall_text  		string `json:"rainfall_text"`  		// example:`ไม่มีฝน` สถานะ
	Path     			string `json:"media_path"`     		// example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHNIfCEaLG3zGgZ2LEuTaHsMjGebDH8UN` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	Filename 			string `json:"filename"`       		// example:`cri240_201612090100.gif` ชื่อไฟล์ของข้อมูลสื่อ
	FilePath 			string `json:"file_path"`      		// example:`product/image/radar/cri/tmd/media/2016/12/09` ที่อยู่ของไฟล์ข้อมูลสื่อ
	Datetime 			string `json:"media_datetime"` 		// example:`2016-12-09 01:00` วันที่เก็บข้อมูลสื่อ
}
