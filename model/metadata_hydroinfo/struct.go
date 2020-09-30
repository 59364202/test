package metadata_hydroinfo

import (
//model_metadata "haii.or.th/api/thaiwater30/model/metadata"
//model_hydroinfo "haii.or.th/api/thaiwater30/model/hydroinfo"
)

type Struct_MetadataHydroinfo_InputParam struct {
	MetadataID      int64   `json:"metadata_id"`
	ListHydroInfoID []int64 `json:"hydroinfo"`
}

type Struct_MetadataHydroinfo struct {
	//Metadata	*model_metadata.Struct_Metadata		`json:"metadata"`
	//Hydroinfo	*model_hydroinfo.Struct_Hydroinfo	`json:"hydroinfo"`
}
