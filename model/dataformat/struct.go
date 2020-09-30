package dataformat

import (
	"encoding/json"
)

type Dataformat_struct struct {
	Id                 int64           `json:"id"`                        // example:`1` รหัสรูปแบบของข้อมูล
	DataformatName     json.RawMessage `json:"dataformat_name,omitempty"` // example:`{"th": "XML", "en": "XML"}` ชื่อรูปแบบข้อมูล
	MetadataMethodID   interface{}     `json:"metadata_method_id"`        // example:`1` รหัสวิธีการได้มาของข้อมูล
	MetadataMethodName string          `json:"metadata_method_name"`      // example:`Web Service` ชื่อวิธีการได้มาของข้อมูล
}

type Struct_Option struct {
	Value int64  `json:"value"` // example:`1` รหัสวิธีการได้มาของข้อมูล
	Text  string `json:"text"`  // example:`Web Service` ชื่อวิธีการได้มาของข้อมูล
}
