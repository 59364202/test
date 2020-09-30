package media_animation

import (
	"time"
)

type Struct_Media struct {
	Path     string `json:"media_path"`     // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7l00s__U-WBd0SNWAl6wOcp7-aUuLdWN050DeakpkhzQma1fIGW_MQgIjamwFchbEN` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	Filename string `json:"filename"`       // example:`d03_day01.jpg` ชื่อไฟล์ของข้อมูลสื่อ
	FilePath string `json:"file_path"`      // example:`product/precipitation/thailand/latest/haii` ที่อยู่ของไฟล์ข้อมูลสื่อ
	Datetime string `json:"media_datetime"` // example:`2016-12-09 01:00` วันที่เก็บข้อมูลสื่อ

	PathThumb     interface{} `json:"media_path_thumb"` // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7l00s__U-WBd0SNWAl6wOcp7-aUuLdWN050DeakpkhzQma1fIWA7UihoiTHRti1M1lOxsoQwf2` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	FilenameThumb interface{} `json:"filename_thumb"`   // example:`thumb-d03_day01.jpg` ชื่อไฟล์ thumb ของข้อมูลสื่อ

	Dt time.Time `json:"-"` // วันที่ ใช้สำหรับนำไปใช้ต่อใน go
}
