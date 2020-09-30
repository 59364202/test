package setting

import ()

//
// bof.Shared.DataTypeOption
//

type Struct_DataTypeOption struct {
	Name                string `json:"name"`                // example:`ฝน` ข้อความที่แสดง
	Data_type           string `json:"data_type"`           // example:`rainfall` ชื่อตาราง
	Data_type_cache     string `json:"data_type_cache"`     // example:`rainfall`  ชื่อตาราง
	Station_table       string `json:"station_table"`       // example:`m_tele_station` ชื่อตาราง master
	Station_column_name string `json:"station_column_name"` // example:`tele_station_id` ชื่อ ฟิลด์ที่เป็น foreign key ของ master
}
