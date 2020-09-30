package dbamodule_history

import (
//	model_user "haii.or.th/api/server/model/user"
)

//Struct of dbamodule_history table
type Struct_DBAModuleHistory struct {
	ID        int64        `json:"id"`           // example:`1477` รหัส dbamodule_history
	TableName string       `json:"table_name"`   // example:`air` ชื่อตาราง
	Year      string       `json:"year"`         // example:`2006` ปี
	Month     string       `json:"month"`        // example:`01` เดือน
	Datetime  string       `json:"dba_datetime"` // example:`2006-01-02 15:04` วันเวลา
	Remark    string       `json:"dba_remark"`   // example:`CREATE public.air_y2006m01` remark
	User      *Struct_User `json:"user"`         // ผู้สร้างพาร์ทิชั่น
}

//Struct of input parameters
type Struct_DBAModuleHistory_InputParam struct {
	TableName string `json:"table_name"`
	Year      string `json:"year"`
	Month     string `json:"month"`
	Remarks   string `json:"remark"`
}

type Struct_User struct {
	ID       int64  `json:"ID"`       // example:`68` รหัสผู้ใช้
	FullName string `json:"FullName"` // example:`Kantamat Polsawang` ชื่อผู้ใช้
}
