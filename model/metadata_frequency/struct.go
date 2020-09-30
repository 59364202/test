package metadata_frequency

import ()

type Struct_MetadataFrequency struct {
	Id            int64  `json:"id"`                      // example:`507` รหัสความถี่ของข้อมูล
	Datafrequency string `json:"datafrequency,omitempty"` // example:`15 นาที` ความถี่ของข้อมูล
}
