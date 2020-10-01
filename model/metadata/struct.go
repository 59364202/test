// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for public.metadata table. This table store metadata.
package metadata

import (
	"database/sql"
	"encoding/json"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_hydroinfo "haii.or.th/api/thaiwater30/model/hydroinfo"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_servicemethod "haii.or.th/api/thaiwater30/model/lt_servicemethod"
	model_subcategory "haii.or.th/api/thaiwater30/model/lt_subcategory"
	model_metadata_frequency "haii.or.th/api/thaiwater30/model/metadata_frequency"
	model_metadata_history "haii.or.th/api/thaiwater30/model/metadata_history"
)

type Struct_M struct {
	Id                        int64           `json:"id"`                             // example:`516` รหัสบัญชีข้อมูล
	Metadata_Convertfrequency string          `json:"metadata_convertfrequency"`      // example:`1 วัน` ความถึ่การเชื่อมโยงของบัญชีข้อมูล
	Metadataservice_Name      json.RawMessage `json:"metadataservice_name,omitempty"` // example:`{"th": "แผนภาพค่าเบี่ยงเบนความสูงระดับน้ำทะเล"}` ชื่อบัญชีข้อมูล
	Connection_Format         string          `json:"connection_format,omitempty"`    // example:`Online` รูปแบบการเชื่อมโยง
	Dataimport_Download_Id    interface{}     `json:"dataimport_download_id"`         // example:`305` รหัส download
	Dataimport_Dataset_Id     interface{}     `json:"dataimport_dataset_id"`          // example:`308` รหัส dataset
}

type Struct_Metadata struct {
	Agency *model_agency.Struct_Agency `json:"agency"`          // หน่วยงาน
	Basin  []*model_basin.Struct_Basin `json:"basin,omitempty"` // ลุ่มน้ำ
	Struct_M
	MetadataId string `json:"metadata_id"` // example:`6ISkJmlyE41mdGl574O2j64ockelAp8NjbNPg3Q76FYeMKDpwO4iZpiXjjscylEBbJpnCVHbIRwzdDzI3Wecsg` รหัสบัญชีข้อมูลแบบเข้ารหัส

	Metadataagency_Name  json.RawMessage `json:"metadataagency_name,omitempty"`  // example:`{"th": "แผนที่แม่น้ำสำคัญในประเทศที่กรมเจ้าท่าดูแล "}` ชื่อบัญชีข้อมูลสำหรับให้หน่วยงานตรวจสอบ
	Metadata_Description json.RawMessage `json:"metadata_description,omitempty"` // example:`{"th": "ภาพแผนที่แม่น้ำสำคัญในประเทศ" }` คำอธิบายเพิ่มเติม

	FormDate            string `json:"fromdate,omitempty"`            // example:`2006-01-02` ช่วงข้อมูลตั้งแต่
	ToDate              string `json:"todate,omitempty"`              // example:`2006-01-02` ช่วงข้อมูลสิ้นสุด
	Total_Import_Record int64  `json:"total_import_record,omitempty"` // example:`25` จำนวนแถวที่นำเข้า
	Last_Import_Date    string `json:"last_import_date,omitempty"`    // example:`2006-01-02 15:04` วันเวลาที่อิมพอร์ตล่าสุด

	Subcategory *model_subcategory.Struct_subcategory `json:"subcategory"` // กลุ่มข้อมูลย่อย

	Hydroinfo []*model_hydroinfo.Struct_Hydroinfo `json:"hydroinfo"` // ข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ

	Province []*model_geocode.Struct_Province `json:"province,omitempty"` // จังหวัด

	Servicemethod []*model_servicemethod.Struct_LtServicemethod        `json:"servicemethod,omitempty"` // กระบวนการให้บริการ
	Frequency     []*model_metadata_frequency.Struct_MetadataFrequency `json:"frequency,omitempty"`     // ความถี่

	Table string `json:"-"`
}

type Param_Metadata struct {
	Id          int64  `json:"id"`
	Data_name   string `json:"data_name"`
	Category    int64  `json:"category"`
	Subcategory int64  `json:"subcategory"`
	Ministry    int64  `json:"ministry"`
	Department  int64  `json:"department"`
	Agency_id   int64  `json:"agency_id"`
}

type Struct_MetadataSummary struct {
	Summary_Total  int64                                 `json:"summary_total"`         // example:`2` จำนวนทั้งหมด
	Summary_Import int64                                 `json:"summary_import"`        // example:`2` จำนวนทั้งหมดที่เชื่อมโยง
	Count_Metadata int64                                 `json:"count_metadata"`        // example:`0` จำนวนบัญชีข้อมูล
	Subcategory    *model_subcategory.Struct_subcategory `json:"subcategory,omitempty"` // กลุ่มข้อมูลย่อย
	Agency         *model_agency.Struct_Agency           `json:"agency,omitempty"`      // หน่วยงาน
}

type Struct_Metadata_InputParam struct {
	Agency_Id   string `json:"agency_id"`   //  example:`57` รหัสหน่วยงาน
	Category_Id string `json:"category_id"` //  example:`1` รหัสหมวดหมู่หลัก
}

type Struct_MetadataImportByAgency_Table struct {
	Columns   []string                 `json:"columns"`    // example:`["id", "waterquality_id", "waterquality_datetime"]` หัวตาราง
	Data      []map[string]interface{} `json:"data"`       // example:`[{"id":121308, "waterquality_id": 132, "waterquality_id": "2017-08-13T17:00:00+07:00"}]` ข้อมูลในรูปแบบ json
	DataArray [][]interface{}          `json:"data_array"` // example:`[[121308, 132, "2017-08-13T17:00:00+07:00"]]` ข้อมูลในรูปแบบ array
}
type Struct_Data_Media struct {
	Agency_id      interface{} `json:"agency_id"`      // example:`9` รหัสหน่วยงาน
	Media_Type_id  interface{} `json:"media_type_id"`  // example:`15` รหัสชนิดข้อมูลสื้อ
	Media_Datetime interface{} `json:"media_datetime"` // example:`2006-01-02 15:04` วันเวลาที่นำเข้าข้อมูล
	Media_Path     interface{} `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=1K9V5PxLSCUKnWhUI_7UACVrh9KWOvgatwBG0nKVsKNevW-FVz4mswY9QdFbqfeQJRWuwBV67dwuMxETACLAnH9VNfMW66F01t1ioHllyMimOxIThK9fEERbVdbg2V0OazTIipyVn7LpLyOMFMJH6w
	//` ลิ้งค์ข้อมูลสื่่อ (ใช้ได้เลย)
	Media_Desc interface{} `json:"media_desc"` // example:`THAILAND 168 HRS FORECASTED WAVE HEIGHT WITH DIRECTION (WRF-SWAN MODEL)` คำอธิบายเพิ่มเติม
	Filename   interface{} `json:"filename"`   // example:`wave_168hr.gif` ชื่อไฟล์ข้อมูลสื้อ

	// ส่วนที่ไม่ต้อง return
	Refer_Source interface{} `json:"-"` // example:`https://api.haii.or.th/tiservice/v1/ws/MBtOTp6IUXbjaCxhQoFQNrFgZUCzNgbo/model/swan/antimation/lastest` แหล่งข้อมูลสำหรับอ้างอิง
	Path         string      `json:"-"`
}

type Struct_MetadataImportByAgency struct {
	Img       []*Struct_Data_Media                 `json:"img"`               // ข้อมูลสื่อ
	Table     *Struct_MetadataImportByAgency_Table `json:"table"`             // ตาราง
	Weather   *Struct_MetadataImportByAgency_Table `json:"weather"`           // ตาราง + แผนที่
	TableName string                               `json:"tb_name,omitempty"` // example:`waterquality` ชื่อตาราง
}

type Struct_MetadataImportByAgency_Img struct {
	Name     string `json:"filename"`            // example:`I.png` ชื่อไฟล์ของข้อมูลสื่อ
	Path     string `json:"media_path"`          // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJH92K2RGFR_Qk7s7SzdWePNk75UQ33K_OgfjqbiJh1KaLWi24TPfyMliOZo_s` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	FilePath string `json:"file_path,omitempty"` // example:`product/image/strom_map/thailand/ucl/media/2017/08/11` ที่อยู่ของไฟล์ข้อมูลสื่อ
}
type Struct_MetadataStatus struct {
	Id                int64           `json:"id"`                           // example:`52` order header id
	UserName          string          `json:"user_name,omitempty"`          // example:`คำแก้ว ดีต่อใจ` ชื่อผู้ใช้
	MetadataName      json.RawMessage `json:"metadata_name,omitempty"`      // example:`{"th":"ฝน 1 วัน จากข้อมูลโทรมาตร" }` ชื่อบัญชีข้อมูล
	ServicemethodName json.RawMessage `json:"servicemethod_name,omitempty"` // example:`{"th":"เว็บเซอร์วิสข้อมูลล่าสุด" }` ชื่อวิธีการให้บริการข้อมูล
	CreateDate        string          `json:"create_date,omitempty"`        // example:`2006-01-02 15:04` วันที่ยื่นคำขอ
	ResultDate        string          `json:"result_date,omitempty"`        // example:`2006-01-02 15:04` วันที่แจ้งผล
	Result            string          `json:"result,omitempty"`             // example:`AP` ผลคำขอ
}

type Struct_Metadata_Data struct {
	Id                       int64  `json:"id"`                        // example:`1` รหัสบัญชีข้อมูล
	MetadataId               string `json:"metadata_id"`               // example:`VTml3cSy3t7jojxiLEfnFcEPnDlYmvEtddbgVrN_Jecjj_OoMneWYkhvKkguYsNR4DVoReZJAYQWTfPqoQ5Y5A` รหัสบัญชีข้อมูลแบบเข้ารหัส
	SubcategoryId            int64  `json:"subcategory_id"`            // example:`22` รหัสกลุ่มข้อมูลย่อย
	AgencyId                 int64  `json:"agency_id"`                 // example:`1` รหัสหน่วยงาน
	DataunitId               int64  `json:"dataunit_id"`               // example:`2` รหัสหน่วยของข้อมูล
	DataformatId             int64  `json:"dataformat_id"`             // example:`11` รหัสรูปแบบของข้อมูล
	ConnectionFormat         string `json:"connection_format"`         // example:`Offline` รูปแบบการเชื่อมโยง
	MetadataChannel          string `json:"metadata_channel"`          // example:`ftp://nhc2md.haii.or.th` ช่องทางการเชื่อมโยงข้อมูล
	MetadataConvertfrequency string `json:"metadata_convertfrequency"` // example:`1 เดือน` ความถึ่การเชื่อมโยงของบัญชีข้อมูล
	MetadataContact          string `json:"metadata_contact"`          // example:`คุณปัญจรัตน์ ปรุงเจริญ` ติดต่อเจ้าของข้อมูล
	MetadataAgencystoredate  string `json:"metadata_agencystoredate"`  // example:`2016-11-29` วันที่หน่วยงานเริ่มจัดเก็บข้อมูล
	MetadataStartdatadate    string `json:"metadata_startdatadate"`    // example:`2016-11-29` วันที่เริ่มใช้งานข้อมูล
	MetadataReceiveDate      string `json:"metadata_receive_date"`     // example:`2016-11-29` วันที่ได้รับข้อมูลจากหน่วยงานเข้าสู่คลังฯ
	MetadataUpdatePlan       int64  `json:"metadata_update_plan"`      // example:`43200` ระยะเวลาการปรับปรุงข้อมูล หน่วยเป็นนาที
	MetadataLaws             string `json:"metadata_laws"`             // example:`` ข้อจำกัดทางกฎหมาย
	MetadataRemark           string `json:"metadata_remark"`           // example:`` หมายเหตุ
	MetadataStatusID         int64  `json:"metadata_status_id"`        // example:`1` รหัสสถานะการเชื่อมโยง

	MetadataAgencyName  json.RawMessage `json:"metadataagency_name,omitempty"`  // example:`{"th":"ระดับน้ำ (**ระบบออนไลน์**)", "en":"" }` ชือบัญชีข้อมูลที่ยืนยันกับหน่วยงาน
	MetadataServiceName json.RawMessage `json:"metadataservice_name,omitempty"` // example:`{"th":"ระดับน้ำ", "en":"" }` ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	MetadataTag         json.RawMessage `json:"metadata_tag,omitempty"`         // example:`{"th":"", "en":"" }` tag ของบัญชีข้อมูล
	MetadataDescription json.RawMessage `json:"metadata_description,omitempty"` // example:`{"th":"ประกอบไปด้วยสถานีตรวจวัด จำนวน 23 สถานีตรวจวัด โดยวาดกราฟ บนกระดาษอ่านข้อมูลทุก ชม. ส่งข้อมูลมา อัพเดตที่กรมเจ้าท่า ทุกเดือน (ช่วงต้นเดือน ประมาณ วันที่ 1-5 ของทุกเดือน) มีข้อมูลตั้งแต่ปี 2516 (ในบางสถานี) (water level เฉพาะสถานีนครหลวงและนครสวรรค์ส่วนที่เหลือเป็นระดับน้ำแบบ Tide (เป็นระดับน้ำที่มีน้ำขึ้น-ลงภายในเวลา 24 ชม.)", "en":"" }` รายละเอียดของบัญชีข้อมูล

	Servicemethod []*model_servicemethod.Struct_LtServicemethod        `json:"servicemethod,omitempty"` // วิธีการให้บริการข้อมูล
	Hydroinfo     []*model_hydroinfo.Struct_Hydroinfo                  `json:"hydroinfo,omitempty"`     // ข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
	Frequency     []*model_metadata_frequency.Struct_MetadataFrequency `json:"frequency,omitempty"`     // ความถี่ของข้อมูล
	History       []*model_metadata_history.Struct_MetadataHistory     `json:"history,omitempty"`       // ประวัติการแก้ไขบัญชีข้อมูล

	CategoryID int64 `json:"category_id"`        // example:`2` รหัสกลุ่มข้อมูลหลัก
	MethodID   int64 `json:"metadata_method_id"` // example:`3` รหัสวิธีการได้มาของข้อมูล
}
type Struct_Metadata_Data_Post_InputParam struct {
	ID                       int64           `json:"-"`
	SubcategoryID            int64           `json:"subcategory_id"`                 // example:`22` รหัสหมวดย่อยของข้อมูล
	AgencyID                 int64           `json:"agency_id"`                      // example:`1` รหัสหน่วยงานเจ้าของข้อมูล
	DataunitID               int64           `json:"dataunit_id"`                    // example:`2` รหัสหน่วยของข้อมูล
	DataformatID             int64           `json:"dataformat_id"`                  // example:`11` รหัสรูปแบบของข้อมูล
	ConnectionFormat         string          `json:"connection_format"`              // example:`Offline` รูปแบบการเชื่อมโยง
	MetadataContact          string          `json:"metadata_contact"`               // example:`คุณปัญจรัตน์ ปรุงเจริญ` ติดต่อเจ้าของข้อมูล
	MetadataAgencystoredate  string          `json:"metadata_agencystoredate"`       // example:`2016-11-29` วันที่หน่วยงานเริ่มจัดเก็บข้อมูล
	MetadataStartdatadate    string          `json:"metadata_startdatadate"`         // example:`2016-11-29` วันที่เริ่มใช้งานข้อมูล
	MetadataReceiveDate      string          `json:"metadata_receive_date"`          // example:`2016-11-29` วันที่ได้รับข้อมูลจากหน่วยงานเข้าสู่คลังฯ
	MetadataUpdatePlan       int64           `json:"metadata_update_plan"`           // example:`43200` ระยะเวลาการปรับปรุงข้อมูล หน่วยเป็นนาที
	MetadataLaws             string          `json:"metadata_laws"`                  // example:`` ข้อจำกัดทางกฎหมาย
	MetadataRemark           string          `json:"metadata_remark"`                // example:`` หมายเหตุ
	MetadataStatusID         int64           `json:"metadata_status_id"`             // example:`1` รหัสสถานะการเชื่อมโยง
	MetadataConvertFrequency string          `json:"metadata_convertfrequency"`      // example:`1 เดือน` ความถึ่การเชื่อมโยงของบัญชีข้อมูล
	ImportCount              int64           `json:"import_count"`                   // example:`1` จำนวนที่ใช้ในการคำนวณเปอร์เซนต์นำเข้าข้อมูล หน่วย : ครั้งต่อวัน
	MetadataAgencyName       json.RawMessage `json:"metadataagency_name,omitempty"`  // example:`{"th":"ระดับน้ำ (**ระบบออนไลน์**)", "en":"" }` ชือบัญชีข้อมูลที่ยืนยันกับหน่วยงาน
	MetadataServiceName      json.RawMessage `json:"metadataservice_name,omitempty"` // example:`{"th":"ระดับน้ำ", "en":"" }` ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	MetadataTag              json.RawMessage `json:"metadata_tag,omitempty"`         // example:`{"th":"", "en":"" }` tag ของบัญชีข้อมูล
	MetadataDescription      json.RawMessage `json:"metadata_description,omitempty"` // example:`{"th":"ประกอบไปด้วยสถานีตรวจวัด จำนวน 23 สถานีตรวจวัด โดยวาดกราฟ บนกระดาษอ่านข้อมูลทุก ชม. ส่งข้อมูลมา อัพเดตที่กรมเจ้าท่า ทุกเดือน (ช่วงต้นเดือน ประมาณ วันที่ 1-5 ของทุกเดือน) มีข้อมูลตั้งแต่ปี 2516 (ในบางสถานี) (water level เฉพาะสถานีนครหลวงและนครสวรรค์ส่วนที่เหลือเป็นระดับน้ำแบบ Tide (เป็นระดับน้ำที่มีน้ำขึ้น-ลงภายในเวลา 24 ชม.)", "en":"" }` รายละเอียดของบัญชีข้อมูล

	ServiceMethod []int64  `json:"servicemethod"` // example:`[1,2,3,4]`รหัสวิธีการให้บริการข้อมูล
	Hydroinfo     []int64  `json:"hydroinfo"`     // example:`[2,3,6]`รหัสกลุ่มข้อมูลด้านน้ำและภูมิอากาศ
	Frequency     []string `json:"frequency"`     // example:`1 ชั่วโมง` รหัสความถี่ของข้อมูล

	HistoryDescription string `json:"history_description"` // example:`เปลี่ยนรูปแบบการเชื่อมโยง` รายละเอียดของการแก้ไขบัญชีข้อมูล
}

type Struct_Metadata_Data_InputParam struct {
	Struct_Metadata_Data_Post_InputParam
	MetadataID string `json:"metadata_id"` //	example:`VTml3cSy3t7jojxiLEfnFcEPnDlYmvEtddbgVrN_Jecjj_OoMneWYkhvKkguYsNR4DVoReZJAYQWTfPqoQ5Y5A` รหัสบัญชีข้อมูลแบบเข้ารหัส
}

type Struct_Metadata_Table_InputParam struct {
	MetadataID    string  `json:"metadata_id"`
	SubcategoryID []int64 `json:"subcategory_id"`
	AgencyID      []int64 `json:"agency_id"`
	Hydroinfo     []int64 `json:"hydroinfo_id"`
	CategoryID    int64   `json:"category_id"`
}

//	map struct table ที่เตรียวไว้สำหรับ shopping
type Struct_Table struct {
	Table          string // ชื่อตาราง
	PartitionField string // ชื่อ partition field
	MasterId       string // ชื่อ master id field
	MasterTable    string // ชื่อ master table
	SelectColumn   string //
	Fields         string //field lookup table
	WhereHAII      string // เงื่อนไขถ้าเป็น สสนก
	Where          string // เพิ่มเงื่อนไข ให้ value <> 999999
	WhereHydro     string // เพิ่มเงื่อนไข  eget hydro1-8
	

	IsMaster    bool // เป็นตาราง master ?
	HasProvince bool // มี geocode_id ในตาราง?
	HasBasin    bool //	มี subbasin_id ในตาราง?
}
type Strct_Data struct {
	OrderDetail_id    int64
	Service_id        int64
	Detail_fromdate   sql.NullString
	Detail_todate     sql.NullString
	Metadata_id       int64
	Agency_id         sql.NullString
	Connection_format sql.NullString
	Sql               sql.NullString
	Table_name        sql.NullString
	Lookup_table      sql.NullString
	Fields            sql.NullString
	Province          sql.NullString
	Basin             sql.NullString
	IsEnabled         sql.NullBool
	Frequency         sql.NullString
}

type MetadataOfflineDate struct {
	MetadataID int64 `json:"metadata_id"` // รหัสบัญชีข้อมูล
}
