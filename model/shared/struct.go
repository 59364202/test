package shared

import ()

type TableStruct struct {
	TableName string `json:"table_name"` // ชื่อตาราง
}

type DeleteDataParam struct {
	TableName string `json:"table_name"`
	ID        string `json:"id"`
}

type PartitionParam struct {
	TableName string `json:"table_name"`
	Year      string `json:"year"`
}

type Struct_MetadataTable struct {
	ID               string   `json:"id"`                // example:`m_air_station` ชื่อตาราง
	TableName        string   `json:"table_name"`        // example:`m_air_station` ชื่อตาราง
	TableDescription string   `json:"table_description"` // example:`ข้อมูลพื้นฐานของเครื่องวัดอากาศ` คำอธิบาย
	ColumnList       []string `json:"columns"`           // example:`["air_station_lat", "air_station_long", "geocode_id"]` field
}
