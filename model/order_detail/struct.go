// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"database/sql"
	"encoding/json"
	"time"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_category "haii.or.th/api/thaiwater30/model/lt_category"
	model_lt_department "haii.or.th/api/thaiwater30/model/lt_department"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_lt_ministry "haii.or.th/api/thaiwater30/model/lt_ministry"
	model_lt_servicemethod "haii.or.th/api/thaiwater30/model/lt_servicemethod"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_order_detail_status "haii.or.th/api/thaiwater30/model/order_detail_status"
	//	model_order_header "haii.or.th/api/thaiwater30/model/order_header"
)

type Struct_OrderDetail struct {
	Agency   *model_agency.Struct_A             `json:"agency"`   // หน่วยงาน
	Basin    []*model_basin.Struct_Basin        `json:"basin"`    // ลุ่มน้ำ
	Category *model_lt_category.Struct_category `json:"category"` // กลุ่มข้อมูลหลัก

	CountLog int64 `json:"count_log"` // example:`1` จำนวนครั้งของการเรียกใช้เซอวิส

	Department *model_lt_department.Struct_Department `json:"department"` // กรม

	Detail_Frequency              string      `json:"detail_frequency"`     // example:`1 วัน` ความถี่ของข้อมูล
	Detail_Fromdate               string      `json:"detail_fromdate"`      // example:`2006-01-02` ระยะการขอข้อมูลตั้งแต่
	Detail_Letterdate             string      `json:"detail_letterdate"`    // example:`2006-01-02` วันที่พิมพ์จดหมายให้กับหน่วยงานเจ้าของข้อมูล
	Detail_Letterno               string      `json:"detail_letterno"`      // example:`QWRT` เลขที่จดหมายส่งให้กับหน่วยงานเจ้าของข้อมูล
	Detail_Letterpath             string      `json:"detail_letterpath"`    // example:`bTNG-JDd9PXJzop2-tTU9Nj-cWZE0_jPqSNPkZ3v6-pX_HKGb7GXZcvfsSjUTuSaA1tGKBpf-9z9dYmH2095Uw` ที่อยู่ของเอกสารตอบรับจากหน่วยงานเจ้าของข้อมูล
	Detail_Remark                 string      `json:"detail_remark"`        // example:`ทำแอพคาดการณ์ฝน` หมายเหตุสำหรับผู้ขอใช้บริการกรอกข้อมูลเพิ่มเติม
	Detail_Source_Result          string      `json:"detail_source_result"` // example:`AP` ผลอนุมัติจากหน่วยงานเจ้าของข้อมูล
	Detail_Source_Result_Itf      interface{} `json:"-"`
	Detail_Source_Result_Date     string      `json:"detail_source_result_date"` // example:`2006-01-02` วันที่บันทึกผลการพิจารณา
	Detail_Source_Result_Date_Itf interface{} `json:"-"`
	Detail_Todate                 string      `json:"detail_todate"` // example:`2006-01-02` ระยะเวลาการขอข้อมูลถึง
	E_Id                          string      `json:"e_id"`          // example:`bTNG-JDd9PXJzop2-tTU9Nj-cWZE0_jPqSNPkZ3v6-pX_HKGb7GXZcvfsSjUTuSaA1tGKBpf-9z9dYmH2095Uw` รหัสของการเรียกใช้เซอวิส
	Id                            int64       `json:"id"`            // example:`55` รหัส order detail
	IsEnabled                     bool        `json:"is_enabled"`    // example:`true` สถานะของการเรียกใช้เซอวิส

	IsSuccess  bool   `json:"is_success"`  // example:`true` (ใช้สำหรับส่ง email) สถานะการอัพโหลดไปยังเครื่อง laravel
	FolderPath string `json:"folder_path"` // example:`` (ใช้สำหรับส่ง email) ที่เก็บไฟล์
	ErrMsg     string `json:"err_msg"`     // example:`` (ใช้สำหรับส่ง email) err string

	LatestAccessTime string                          `json:"latest_access_time"` // example:`2006-01-02 15:04` วันเวลาล่าสุดของการเรียกใช้
	Metadata         *model_metadata.Struct_Metadata `json:"metadata,omiempty"`  // บัญชีข้อมูล

	Ministry                    *model_lt_ministry.Struct_Ministry                  `json:"ministry"`                    // กระทรวง
	Order_Detail_Status         *model_order_detail_status.Struct_OrderDetailStatus `json:"order_detail_status"`         // สถานะของ order detail
	Order_Header_Id             int64                                               `json:"order_header_id"`             // example:`45` รหัส order header
	Order_Header_Order_Datetime string                                              `json:"order_header_order_datetime"` // example:`2006-01-02` วันที่ขอข้อมูล
	Order_Purpose               string                                              `json:"order_purpose"`               // example:`ทำแอพคาดการณ์ฝน` วัตถุประสงค์ของการขอข้อมูล
	Order_Expire_Datetime       string                                              `json:"order_expire_date"`           // example:`2006-01-0ุ` วันที่หมดอายุ

	Province    []*model_lt_geocode.Struct_Geocode             `json:"province"`    // จังหวัด
	Service     *model_lt_servicemethod.Struct_LtServicemethod `json:"service"`     // เซอวิส
	Service_Id  int64                                          `json:"service_id"`  // example:`3` รหัสเซอร์วิส
	Service_Url string                                         `json:"service_url"` // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/api_service?mid=1&eid=bTNG-JDd9PXJzop2-tTU9Nj-cWZE0_jPqSNPkZ3v6-pX_HKGb7GXZcvfsSjUTuSaA1tGKBpf-9z9dYmH2095Uw` ลิงค์ของการเรียกใช้เซอวิส
	Status_Id   int64                                          `json:"status_id"`   // example:`1` รหัสของสถานะการขอข้อมูล

	User_AgencyName     json.RawMessage `json:"user_agency_name,omitempty"`     // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงานของผู้ขอใช้บริการ
	User_ContractPhone  string          `json:"user_contract_phone"`            // example:`0865669906` เบอร์ติดต่อของผู้ขอใช้บริการ
	User_DepartmentName json.RawMessage `json:"user_department_name,omitempty"` // example:`{"th": "กรมเจ้าท่า"}` ชื่อกรมของผู้ขอใช้บริการ
	User_Fullname       string          `json:"user_fullname"`                  // example:`คำแก้ว ดีต่อใจ` ชื่อผู้ขอใช้บริการ
	User_Id             int64           `json:"user_id"`                        // example:`69` รหัสผู้ขอใช้บริการ
	User_MinistryName   json.RawMessage `json:"user_ministry_name,omitempty"`   // example:`{"th": "กระทรวงคมนาคม", "en": "Ministry of Transport"}` ชื่อกระทรวงของผู้ขอใช้บริการ
	User_OfficeName     string          `json:"user_offilce_name"`              // example:`Cim System (Thailand)` ชื่อสำนักงานของผู้ขอใช้บริการ
}

type Param_OrderDetail struct {
	Metadata_Id                   int64       `json:"metadata_id"` // example:`1` รหัสบัญชีข้อมูล
	Order_Header_Id               int64       `json:"-"`
	Status_Id                     int64       `json:"-"`
	Service_Id                    int64       `json:"service_id"`       // example:`1` รหัสเซอร์วิส
	Detail_Todate                 string      `json:"detail_todate"`    // example:`2006-01-02` ระยะเวลาการขอข้อมูลถึง
	Detail_Fromdate               string      `json:"detail_fromdate"`  // example:`2006-01-02` ระยะการขอข้อมูลตั้งแต่
	Detail_Frequency              string      `json:"detail_frequency"` // example:`1 ชั่วโมง` ความถี่ของข้อมูล
	Detail_Remark                 string      `json:"detail_remark"`    // example:`ทดสอบ` หมายเหตุสำหรับผู้ขอใช้บริการกรอกข้อมูลเพิ่มเติม
	Detail_Province               string      `json:"detail_province"`  // example:`10` รหัสจังหวัด
	Detail_Basin                  string      `json:"detail_basin"`     // example:`1,2,3` รหัสลุ่มน้ำ
	Detail_Source_Result_Itf      interface{} `json:"-"`
	Detail_Source_Result_Date_Itf interface{} `json:"-"`
}

type Struct_CountOrderDetailByAgencyId struct {
	AgencyId   int64           `json:"agency_id"`             // example:`9` รหัสหน่วยงาน
	AgencyName json.RawMessage `json:"agency_name,omitempty"` // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงาน
	Count      int64           `json:"count,omitempty"`       // example:`3` จำนวนครั้งที่ให้บริการ
}

type Param_OrderLetter struct {
	Order_Header_Id int64    `json:"order_header_id"` // example:`1` รหัสคำขอ
	Date            []string `json:"date"`            // example:`2006-01-02` วันที่
	LetterNo        []string `json:"letterno"`        // example:`ABC123` เลขที่เอกสาร
	Agency          []int64  `json:"agency"`          // example:`1` รหัสหน่วยงาน
	UserId          int64    `json:"-"`
}

type Param_OrderLetterPath struct {
	Order_Header_Id   int64  `json:"order_header_id"` // รหัสคำขอ
	Agency_Id         int64  `json:"agency_id"`       // รหัสหน่วยงาน
	Detail_Letterpath string `json:"-"`
}
type Param_OrderApprove struct {
	Service_Id      int64  `json:"-"`
	Agency_Id       int64  `json:"agency_id"`       // example:`1` รหัสหน่วยงาน
	Detail_Letterno string `json:"detail_letterno"` // example:`AB322` เลขที่เอกสารร้องขอข้อมูล
}
type Pram_OrderApprove_Put struct {
	Id                   int64  `json:"id"`                   // รหัสรายการ order detail id
	Detail_Source_Result string `json:'detail_source_result'` // ผลอนุมัติ AP อนุมัติ ,DA ไม่อนุมัติ
}
type Param_OrderExpireDate_Put struct {
	Id                 int64  `json:"id"`          // example:`55` รหัส order detail
	Detail_Expire_Date string `json:"expire_date"` // example:`1970-01-12 07:00:00+07` วันหมดอายุ
}
type Param struct {
	Date_Start      string `json:"datestart"` // example:`2006-01-02` วันที่เริ่มต้น
	Date_End        string `json:"dateend"`   // example:`2006-01-02` วันที่สิ้นสุด
	Agency_Id       int64  `json:"agency_id"`
	UserAgency_Id   int64  `json:"user_agency_id"`
	Statud_Id       int64  `json:"status_id"`
	Order_Header_Id int64  `json:"order_header_id"`
	Order_Detail_id int64  `json:"order_detail_id"`
	Date            string `json:"date"`
	User_Id         int64  `json:"user_id"`
	Id              int64  `json:"id"`
	Field           string `json:"field"`
	Eid             string `json:"eid"`
}
type Strct_Data struct {
	OrderDetail_id    int64
	Service_id        int64
	Detail_fromdate   sql.NullString // ข้อมูลตั้งแต่วันที่..
	Detail_todate     sql.NullString // ข้อมูลจนถึงวันที่
	Metadata_id       int64
	Agency_id         sql.NullString // รหัสหน่วยงาน
	Connection_format sql.NullString
	Sql               string
	Table_name        sql.NullString // ชื่อตาราง
	Lookup_table      sql.NullString
	Fields            sql.NullString
	SelectFields      string         // select field ดึงมาจาก import_setting
	Province          sql.NullString // รหัสจังหวัด
	Basin             sql.NullString // รหัสลุ่มน้ำ
	IsEnabled         sql.NullBool
	Frequency         sql.NullString
	MediaTypeId       string // รหัสประเภทข้อมูลสื่อ
	CreateAt          time.Time
	Dataset_id        int64          // dataset
	HasQC             bool           // มี qc_status
	AdditionalDataset sql.NullString // dataset ที่เกี่ยวข้อง
}

type MailData struct {
	UserName    string                `json:"username"`
	UserId      int64                 `json:"user_id"`
	Date        string                `json:"date"`
	Data        []*Struct_OrderDetail `json:"data"`
	ServiceId   int64                 `json:""`
	AgentUserId int64                 `json:""`
	IsInit      bool                  `json:"isinit"`
}

// type MailDataErr struct {
// 	MailData
// 	Error []error `json:"error"`
// }
