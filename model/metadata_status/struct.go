package metadata_status

import (
	"encoding/json"
)

type MetadataStatusParams struct {
	ID         int64           `json:"id"`                   // example:`2`  สถานะของบัญชีข้อมูล
	StatusID   string          `json:"metadata_status_id"`   // example:`IJPRn6cYcJZ3KFrDxEEnnJ4d8_FlXM1X_l5oN4AnPxY3MvRzkxDEZsx-ZZ48swiGmrUxi40TGQp0T-VlwmtN-g` รหัสสถานะของบัญชีข้อมูลแบบเข้ารหัส
	StatusName json.RawMessage `json:"metadata_status_name"` // example:`{"en": "", "th": "รอการเชื่อมโยง"}` ชื่อสถานะของบัญชีข้อมูล
	IsDeleted  bool            `json:"is_deleted"`           // example:`true` สามารถลบได้
}

type DataMetadataStatusPostParams struct {
	StatusName json.RawMessage `json:"metadata_status_name"` // example:`{"th":"รอการเชื่อมโยง", "en":"english" }` ชื่อสถานะของบัญชีข้อมูล
}
type DataMetadataStatusParams struct {
	StatusID string `json:"id"` // example:`IJPRn6cYcJZ3KFrDxEEnnJ4d8_FlXM1X_l5oN4AnPxY3MvRzkxDEZsx-ZZ48swiGmrUxi40TGQp0T-VlwmtN-g` รหัสสถานะของบัญชีข้อมูลแบบเข้ารหัส
	DataMetadataStatusPostParams
}
