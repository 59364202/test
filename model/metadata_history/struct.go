package metadata_history

import (
	model_user "haii.or.th/api/server/model/user"
)

type Struct_MetadataHistory struct {
	ID           int64  `json:"id"`                  // รหัสประวัติการแก้ไขบัญชีข้อมูล
	Description  string `json:"history_description"` // รายละเอียดของการแก้ไขบัญชีข้อมูล
	Datetime     string `json:"history_datetime"`    // วันที่แก้ไขบัญชีข้อมูล
	UserFullname string `json:"created_by"`          // ชื่อผู้แก้ไข

	User *model_user.User `json:"user"` // ผู้แก้ไข
}

type Struct_MetadataHistory_InputParam struct {
	MetadataID  int64  `json:"metadata_id"`         // รหัสบัญชีข้อมูล
	Description string `json:"history_description"` // รายละเอียดของการแก้ไขบัญชีข้อมูล
}
