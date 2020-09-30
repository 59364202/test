package impdata

import (
	"haii.or.th/api/thaiwater30/util/selectoption"
)

type ImportDataOptionListSwagger struct {
	Agency   *Agency               `json:"agency"`   // หน่วยงาน
	Metadata []*MetadataOptionList `json:"metadata"` // บัญชีข้อมูล
}

type Agency struct {
	Value interface{} `json:"value"` // example:`1` รหัส agency
	Text  interface{} `json:"text"`  // example:`{"th":"สสนก.","en":"","jp":""}`
}

type ImportDataOptionList struct {
	Agency   *selectoption.Option  `json:"agency"`   // หน่วยงาน
	Metadata []*MetadataOptionList `json:"metadata"` // บัญชีข้อมูล
}

type MetadataOptionList struct {
	Value           interface{} `json:"value"`            // example:`1` รหัสหน่วยงาน
	Text            interface{} `json:"text"`             // example:`{"th":"สสนก."}`ชื่อหน่วยงาน
	DownloadScript  interface{} `json:"download_script"`  // example:`dl-basic` กระบวนการดาวน์โหลด
	DownloadCommand interface{} `json:"download_command"` // example:`bin/rdl 1 dl-basic` คำสั่งดาวน์โหลด
	MetadataID      interface{} `json:"metadata_id"`      // example:`1` รหัส metadata
}
