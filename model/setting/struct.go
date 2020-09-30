package setting

import ()

type Struct_SystemSetting struct {
	Id          int64  `json:"id"`          // รหัส system setting
	User_id     int64  `json:"user_id"`     // รหัสผู้ใช้
	Name        string `json:"name"`        // ชื่อ system setting
	Is_public   bool   `json:"is_public"`   //
	Value       string `json:"data"`        // ค่า setting
	Description string `json:"description"` // คำอธิบาย
}
