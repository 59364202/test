package metadata_provision

import (
	"encoding/json"
)

type MetadataProvisionOutput struct {
	ID           int64           `json:"id"`                             // example:`` รหัสบัญชีข้อมูล
	Name         json.RawMessage `json:"metadataservice_name,omitempty"` // example:`{"th":"ชื่อบัญชีข้อมูล"}` ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	DownloadName string          `json:"download_name"`                  // example:`ชื่อ download` ชื่อ download
	DatasetName  string          `json:"dataset_name"`                   // example:`ชื่อ dataset` ชื่อ dataset
}

type MetadataProvisionInput struct {
	ID   int64  `json:"id"`                   // example:`1` รหัสบัญชีข้อมูล
	Name string `json:"metadataservice_name"` // example:`{"th":"ชื่อบัญชีข้อมูล"}` ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
}
