package setting

//
// bof.Tool.Ignore.TableList
//
type Struct_Inore_TableList struct {
	Name                string `json:"name"`                // example:`ฝนย้อนหลัง 24 ชั่วโมง` ข้อความที่แสดง
	Data_type           string `json:"data_type"`           // example:`rainfall_24h` ชื่อตาราง
	Station_table       string `json:"station_table"`       // example:`m_tele_station` ชื่อตาราง master
	Station_column_name string `json:"station_column_name"` // example:`tele_station_id` ชื่อ ฟิลด์ที่เป็น foreign key ของ master
}

type Struct_Ignore_Station struct {
	ID   string `json:"station_id"`
	Name string `json:"station_name"`
}

//
// bof.Tool.Ignore.TableList ที่ใส่ lastest_ignore_datetime
//
type Struct_Inore_TableList_LatestIgnoreDatetime struct {
	Struct_Inore_TableList
	Struct_Ignore_Station
	Lastest_ignore_datetime string `json:"lastest_ignore_datetime"` // example:`2006-01-02 15:04` วันที่ ignore ครั้งล่าสุด
}
