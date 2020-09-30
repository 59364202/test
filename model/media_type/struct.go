package media_type

import ()

type Struct_MediaType struct {
	Id                int64  `json:"id"`                      // example:`18` รหัสแสดงชนิดข้อมูลสื่อ
	Type_name         string `json:"media_type_name"`         // example:`Precipitation` ชื่อแสดงชนิดข้อมูลสื่อ
	Subtype_name      string `json:"media_subtype_name"`      // example:`Asia` ชื่อแสดงชนิดย่อยข้อมูลสื่อ
	Type_subtype_name string `json:"media_type_subtype_name"` // example:`Precipitation - Asia` ชื่อแสดงชนิดข้อมูลสื่อ - ชื่อแสดงชนิดย่อยข้อมูลสื่อ
}
