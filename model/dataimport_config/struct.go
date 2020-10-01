// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataimport_config is a model for api.dataimport_download table. This table store dataimport config information.
package dataimport_config

import (
	"encoding/json"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/dataimport"
)

type DataImportConfig struct {
	DataImportDownload *DataImportDownloadConfig  `json:"dataimport_download_config"` // download config
	DataImportDataset  []*DataImportDatasetConfig `json:"dataimport_dataset_config"`  // dataset config
	//	MonitorScript      string                     `json:"monitor_script"`
	CrontabSetting string `json:"crontab_setting"` // example:`* * * * *`การตั้งค่าเวลาของ crontab
}

type DataImportDownloadConfig struct {
	ID              int64           `json:"id"`                         // example:`1` รหัส download config
	AgentUserID     int64           `json:"agent_user_id"`              // example:`2`  รหัสผู้ใช้ของหน่วยงาน
	DownloadSetting json.RawMessage `json:"download_setting,omitempty"` // example:`{"archive_folder":"","result_file":"","data_folder":"download/bma/canal","source_options":[{"name":"canal","retry_count":3,"details":[{"driver":"","host":"ftp://bma1:vz5IjIkXOS84Q8HLhK5va8EE795ASh6NkdpfUBBqgz-WilcmDK1vPyhF0K8xNdp3WRQkzjecvnzZJ4B7W90zZA@nhc2bma.haii.or.th","timeout_seconds":0,"tracking_key":"","files":[{"source":"Canal.txt","destination":"canal.txt"}]}]}]}`  download setting
	DownloadScript  string          `json:"download_method"`            // example:`dl-basic`  กระบวนการดาวน์โหลด
	CrontabSetting  string          `json:"crontab_setting"`            // example:`* * * * *` การตั้งค่าเวลาของ crontab
	DownloadName    string          `json:"download_name"`              // example:`download name`  ชื่อการดาวน์โหลด
	IsCronenabled   bool            `json:"is_cronenabled"`             // example:`true` สถานะ crontab
	Description     string          `json:"description"`                // example:`download description` คำอธิบายเพิ่มเติม
	Node            string          `json:"node"`                       // example:`node0` node
	MaxProcess      int64           `json:"max_process"`                // example:`10`
}

type DataImportDownloadConfigList struct {
	ID             int64  `json:"id"`              // example:`1` รหัส download config
	DownloadName   string `json:"download_name"`   // example:`download_name`  ชื่อการดาวน์โหลด
	DownloadScript string `json:"download_method"` // example:`dl-basic`  กระบวนการดาวน์โหลด
	IsCronenabled  bool   `json:"is_cronenabled"`  // example:`false`  สถานะ crontab
	Description    string `json:"description"`     // example:`description`  คำอธิบายเพิ่มเติม
	IsRun          bool   `json:"is_run"`          // example:`false` สถานะการทำงาน
	Node           string `json:"node"`            // example:`node0` node
}

type DataDetailsDownloadListSwagger struct {
	SelectOpt              *SelectOptionDownloadListSwagger `json:"select_option"`            // dropdown
	DataImportDownloaaList []*DataImportDownloadConfigList  `json:"dataimport_download_list"` // download config
}

type SelectOptionDownloadListSwagger struct {
	DownloadScript []*DLScript `json:"download_method"` // กระบวนการดาวน์โหลด
	DownloadDriver []*DLDriver `json:"download_driver"` // ชื่อ driver
	AgentUser      []*AUser    `json:"agent_user"`      // ชื่อผู้ใช้ของหน่วยงาน
}

type AUser struct {
	Text  string `json:"text"`  // example:`dataimport-bma`
	Value string `json:"value"` // example:`9`
}

type DLDriver struct {
	Text  string `json:"text"`  // example:`http://`
	Value string `json:"value"` // example:`http://`
}

type DLScript struct {
	Text   string `json:"text"`   // example:`dl-collector`
	Value  string `json:"value"`  // example:`dl-collector`
	Enable string `json:"enable"` // example:`collector`
}

type SelectOptionDownloadList struct {
	DownloadScript *result.ResultJson `json:"download_method"` // กระบวนการดาวน์โหลด
	DownloadDriver *result.ResultJson `json:"download_driver"` // ชื่อ driver
	AgentUser      *result.Result     `json:"agent_user"`      // ชื่อผู้ใช้ของหน่วยงาน
	RDLNodes       *result.Result     `json:"rdl_nodes"`       // rdl node ทั้งหมดที่มีในระบบ
}

type DataDetailsDownloadList struct {
	SelectOpt              *SelectOptionDownloadList       `json:"select_option"`            // dropdown
	DataImportDownloaaList []*DataImportDownloadConfigList `json:"dataimport_download_list"` // download config
}

type SelectOptionDatasetList struct {
	//	xDownloadScript      *result.ResultJson `json:"download_method"`       // กระบวนการดาวน์โหลด
	//	xDownloadDriver      *result.ResultJson `json:"download_driver"`       // ชื่อ driver
	ConvertScript       *result.ResultJson `json:"convert_method"`        // กระบวนการคอนเวิร์ท
	ImportScript        *result.ResultJson `json:"import_method"`         // กระบวนการอิมพอร์ต
	Type                *result.ResultJson `json:"type"`                  // ชนิดของฟิลดิ์ข้อมูล
	TransformMethod     *result.ResultJson `json:"transform_method"`      // กระบวนการแปลง
	AddMissing          *result.ResultJson `json:"add_missing"`           // ฟังก์ชันการเพิ่มข้อมูลสถานี
	InputFormatDatetime *result.ResultJson `json:"input_format_datetime"` // รูปแบบของวันเวลา
	MasterTable         *result.Result     `json:"master_table"`          // master table
	AgentUser           *result.Result     `json:"agent_user"`            // ผู้ใช้ของหน่วยงาน
	DownloadList        *result.Result     `json:"download_list"`         // ชุดข้อมูล
	ParentTable         *result.Result     `json:"import_table"`          // ตารางที่จะนำข้อมูลเข้า
}

type SelectOptionDatasetListSwagger struct {
	ConvertScript       *ConvertMethod        `json:"convert_method"`        // กระบวนการคอนเวิร์ท
	ImportScript        *ImportMethod         `json:"import_method"`         // กระบวนการอิมพอร์ต
	Type                *DataimportFieldType  `json:"type"`                  // ชนิดของฟิลดิ์ข้อมูล
	TransformMethod     *TransformMethod      `json:"transform_method"`      // กระบวนการแปลง
	AddMissing          *AddMissing           `json:"add_missing"`           // ฟังก์ชันการเพิ่มข้อมูลสถานี
	InputFormatDatetime *InputFormatDatetime  `json:"input_format_datetime"` // รูปแบบของวันเวลา
	MasterTable         *RRTable              `json:"master_table"`          // master table
	AgentUser           *AgentUser            `json:"agent_user"`            // ผู้ใช้ของหน่วยงาน
	DownloadList        *DownloadListFilename `json:"download_list"`         // ชุดข้อมูล
	ParentTable         *RRTable              `json:"import_table"`          // ตารางที่จะนำข้อมูลเข้า
}

type RRTable struct {
	Result string                 `json:"result"` //example:`OK`
	Data   []*SelectOptionRelated `json:"data"`
}

type AgentUser struct {
	Result string   `json:"result"` //example:`OK`
	Data   []*AUser `json:"data"`
}

type DownloadListFilename struct {
	Result string                           `json:"result"` //example:`OK`
	Data   []*DownloadNameOptionAndFilename `json:"data"`
}

type InputFormatDatetime struct {
	Result string `json:"result"` //example:`OK`
	Data   []*FDT `json:"data"`   // example:`[{"text": "%Y-%m-%d %H:%M:%S","value": "%Y-%m-%d %H:%M:%S"}]`
}

type AddMissing struct {
	Result string      `json:"result"` //example:`OK`
	Data   []*AMissing `json:"data"`   // example:`[{"text": "true","value": "true"}]`
}

type TransformMethod struct {
	Result string `json:"result"` //example:`OK`
	Data   []*TFM `json:"data"`   // example:`[{"text": "constant","value": "constant","enable": "constant"},{ "text": "custom","value": "evaluate","enable": "evaluate"}]`
}

type ConvertMethod struct {
	Result string      `json:"result"` //example:`OK`
	Data   []*CVScript `json:"data"`   // example:`[{"text": "cv-std","value": "cv-std"}]`
}

type ImportMethod struct {
	Result string      `json:"result"` //example:`OK`
	Data   []*ImScript `json:"data"`   // example:`[{"text": "im-std","value": "im-std"}]`
}

type DataimportFieldType struct {
	Result string    `json:"result"` //example:`OK`
	Data   []*DSType `json:"data"`   // example:`[{"text": "string","value": "string"},{"text": "int","value": "int"}]`
}

type FDT struct {
	Text  string `json:"text"`  // example:`%Y-%m-%d %H:%M:%S"` // format text
	Value string `json:"value"` // example:`%Y-%m-%d %H:%M:%S` // format value
}

type AMissing struct {
	Text  string `json:"text"`  // example:`true` // missing text
	Value string `json:"value"` // example:`true` // missing value
}

type CVScript struct {
	Text  string `json:"text"`  // example:`cv-std` // convert script text
	Value string `json:"value"` // example:`cv-std` // convert script value
}

type ImScript struct {
	Text  string `json:"text"`  // example:`im-std` // import script text
	Value string `json:"value"` // example:`im-std` // import script value
}

type DSType struct {
	Text  string `json:"text"`  // example:`string` // datatype text
	Value string `json:"value"` // example:`string` // datatype value
}

type TFM struct {
	Text   string `json:"text"`   // example:`constant` // tranform text
	Value  string `json:"value"`  // example:`constant` // tranform value
	Enable string `json:"enable"` // example:`constant` // tranform enable func
}

type DataImportDatasetConfigList struct {
	ID             int64       `json:"id"`              // example:`1`รหัส dataset
	DownloadName   string      `json:"download_name"`   // example:`"dl-bma-canal"`ชื่อการดาวน์โหลด
	ConvertName    string      `json:"convert_name"`    // example:`"cv-bma-canal"`ชื่อการคอนเวิร์ท
	TableName      string      `json:"table_name"`      // example:`"tele_waterlevel"`ชื่อตารางที่นำเข้าข้อมูล
	DownloadID     interface{} `json:"download_id"`     // example:`1`รหัส download
	DownloadScript string      `json:"download_script"` // example:`"dl-basic"`กระบวนการดาวน์โหลด
}

type SelectOptionRelated struct {
	Value   string           `json:"value"`   // example:`m_canal_station`ค่า
	Text    string           `json:"text"`    // example:`m_canal_station`ข้อความ
	Related []*RelatedOption `json:"related"` // related selectoption
}

type RelatedOption struct {
	Value string `json:"value"` // example:`canal_station_oldcode` ค่า
	Text  string `json:"text"`  // example:`canal_station_oldcode` ข้อความ
}

type DownloadNameOptionAndFilename struct {
	Value    interface{} `json:"value"`    // example:`1` รหัส download
	Text     interface{} `json:"text"`     // example:`bma-canal` ชื่อ download
	FileName string      `json:"filename"` // example:`canal.txt` ชื่อไฟล์
}

type DataDetailsDatasetList struct {
	SelectOpt             *SelectOptionDatasetList       `json:"select_option"`           // selectoption
	DataImportDatasetList []*DataImportDatasetConfigList `json:"dataimport_dataset_list"` // dataset
}

type DataDetailsDatasetListSwagger struct {
	SelectOpt             *SelectOptionDatasetListSwagger `json:"select_option"`           // selectoption
	DataImportDatasetList []*DataImportDatasetConfigList  `json:"dataimport_dataset_list"` // dataset
}

type DataImportDatasetConfig struct {
	ID                   int64           `json:"id"`                        // รหัส dataset
	AgentUserID          int64           `json:"agent_user_id"`             // รหัสผู้ใช้ของหน่วยงาน
	DataImportDownloadID int64           `json:"dataimport_download_id"`    // รหัส download
	ConvertSetting       json.RawMessage `json:"convert_setting,omitempty"` // ตั้งค่าการคอนเวิร์ท
	ImportSetting        json.RawMessage `json:"import_setting,omitempty"`  // ตั้งค่าการอิมพอร์ต
	LookupTable          json.RawMessage `json:"lookup_table,omitempty"`    // ชื่อตารางที่ใช้ในการ map ข้อมูล
	ImportTable          json.RawMessage `json:"import_table,omitempty"`    // ชื่อตารางที่จะนำข้อมูลเข้า
	ConvertScript        string          `json:"convert_method"`            // กระบวนการคอนเวิร์ท
	ImportScript         string          `json:"import_method"`             // กระบวนการอิมพอร์ต
}

type AddDataimportConfig struct {
	DownloadConfig      DataDownloadConfig        `json:"download_config"` // ตั้งค่าการ download
	ConvertImportConfig []DataImportDataSetConfig `json:"cv_im_configs"`   // ตั้งค่า dataset
}
type DataDownloadConfigSwagger struct {
	AgentUserID     int64                     `json:"agent_user_id"`    // รหัสผู้ใช้ของหน่วยงาน
	DownloadScript  string                    `json:"download_method"`  // กระบวนการดาวน์โหลด
	DownloadName    string                    `json:"download_name"`    // ชื่อการดาวน์โหลด
	CrontabSetting  string                    `json:"crontab_setting"`  // การตั้งค่าเวลาของ crontab
	DownloadType    string                    `json:"download_type"`    // ชนิดของการดาวน์โหลด
	DownloadSetting dataimport.DownloadConfig `json:"download_setting"` // ตั้งค่าการ download
	Description     string                    `json:"description"`      // คำอธิบายเพิ่มเติม
	Node            string                    `json:"node"`             // node
}

type DataDownloadConfig struct {
	DownloadConfigID int64                     `json:"id"`               // รหัส download
	AgentUserID      int64                     `json:"agent_user_id"`    // รหัสผู้ใช้ของหน่วยงาน
	DownloadScript   string                    `json:"download_method"`  // กระบวนการดาวน์โหลด
	DownloadName     string                    `json:"download_name"`    // ชื่อการดาวน์โหลด
	CrontabSetting   string                    `json:"crontab_setting"`  // การตั้งค่าเวลาของ crontab
	DownloadType     string                    `json:"download_type"`    // ชนิดของการดาวน์โหลด
	DownloadSetting  dataimport.DownloadConfig `json:"download_setting"` // ตั้งค่าการ download
	Description      string                    `json:"description"`      // คำอธิบายเพิ่มเติม
	Node             string                    `json:"node"`             // node
	MaxProcess       int64                     `json:"max_process"`      // max process
	IsCronEnabled    bool                      `json:"is_cronenabled"`   // cron enabled
}

type DataImportDataSetConfigSwagger struct {
	DownloadConfigID  int64                `json:"download_id"`       // example:`1`รหัส download
	AgentUserID       int64                `json:"agent_user_id"`     // example:`7` รหัสผู้ใช้ของหน่วยงาน
	ConvertScript     string               `json:"convert_method"`    // example:`cv-std` กระบวนการคอนเวิร์ท
	ImportScript      string               `json:"import_method"`     // example:`im-std` กระบวนการอิมพอร์ต
	ImportDestination string               `json:"import_table"`      // example:`rainfall` ชื่อตารางที่จะนำข้อมูลเข้า
	UniqueConstraint  string               `json:"unique_constraint"` // example:`uk_rainfall` ชื่อ unique key ของตาราง
	PartitionField    string               `json:"partition_field"`   // example:`rainfall_datetime` ชื่อคอลัมน์ที่ใช้ในการตรวจสอบพาทิชัน
	RowValidator      string               `json:"row_validator"`     // example:`!is_nil(rainfall_value)` การตั้งค่า reject null ของข้อมูล
	ConvertName       string               `json:"convert_name"`      // example:`ฝน`ชื่อชุดข้อมูล
	ConvertSetting    ConvertSettingStruct `json:"convert_setting"`   // ตั้งค่าการคอนเวิร์ท
}

type DataImportDataSetConfig struct {
	DownloadConfigID  int64                `json:"download_id"`       // รหัส download
	DataSetConfigID   int64                `json:"dataset_config_id"` // รหัสการตั้งค่า
	AgentUserID       int64                `json:"agent_user_id"`     // รหัสผู้ใช้ของหน่วยงาน
	ConvertScript     string               `json:"convert_method"`    // กระบวนการคอนเวิร์ท
	ImportScript      string               `json:"import_method"`     // กระบวนการอิมพอร์ต
	ImportDestination string               `json:"import_table"`      // ชื่อตารางที่จะนำข้อมูลเข้า
	UniqueConstraint  string               `json:"unique_constraint"` // ชื่อ unique key ของตาราง
	PartitionField    string               `json:"partition_field"`   // ชื่อคอลัมน์ที่ใช้ในการตรวจสอบพาทิชัน
	RowValidator      string               `json:"row_validator"`     // การตั้งค่า reject null ของข้อมูล
	ConvertName       string               `json:"convert_name"`      // ชื่อชุดข้อมูล
	ConvertSetting    ConvertSettingStruct `json:"convert_setting"`   // ตั้งค่าการคอนเวิร์ท
}

type ConvertSettingStruct struct {
	DownloadFolder string    `json:"data_folder"` // example:`dataset/haii/rainfall`โฟลเดอร์สำหรับวางไฟล์คอนเวิร์ทและอิมพอร์ต
	Configs        []Configs `json:"configs"`     // ตั้งค่า
}

type Configs struct {
	Name         string   `json:"name"`          // example:`rainfall` ชื่อของการดาวน์โหลดใช้สำหรับการประมวลผล Go ต้องตั้งชื่อภาษาอังกฤษเท่านั้น
	InputName    string   `json:"input_name"`    // example:`rainfall.json` ชื่อไฟล์ที่ใช้แปลงข้อมูล
	HeaderRow    int64    `json:"header_row"`    // example:`0` จำนวนบรรทัดส่วนหัว
	DataTag      string   `json:"data_tag"`      // example:`value` tag json
	RowValidator string   `json:"row_validator"` // example:`!is_nil(rainfall_value)` การตั้งค่า reject null ของข้อมูล
	Fields       []Fields `json:"fields"`        // field
}

type Fields struct {
	Name            string      `json:"name"`             // example:`rainfall_value`ชื่อคอลัมน์ที่จะนำข้อมูลเข้า
	Type            string      `json:"type"`             // example:`float`ชนิดของฟิลดิ์ข้อมูล
	InputFields     []string    `json:"input_fields"`     // example:`1`ชื่อฟิลดิ์ input สำหรับการคอนเวิร์ท
	TransformMethod string      `json:"transform_method"` // example:`custom`กระบวนการแปลง
	TransformParams interface{} `json:"transform_params"` // example:`input("1")`ตั้งค่าการแปลง
}

type DataImportDataSetConfig1 struct {
	DownloadConfigID  int64           `json:"download_id"`       // example:`1`รหัส download
	DataSetConfigID   int64           `json:"dataset_config_id"` // example:`1`รหัสการตั้งค่า
	AgentUserID       int64           `json:"agent_user_id"`     // example:`8`รหัสผู้ใช้ของหน่วยงาน
	ConvertScript     string          `json:"convert_method"`    // example:`cv-std`กระบวนการคอนเวิร์ท
	ImportScript      string          `json:"import_method"`     // example:`im-std`กระบวนการอิมพอร์ต
	ImportDestination string          `json:"import_table"`      // example:`canal_waterlevel`ชื่อตารางที่จะนำข้อมูลเข้า
	UniqueConstraint  string          `json:"unique_constraint"` // example:`uk_canal_waterlevel`ชื่อ unique key ของตาราง
	PartitionField    string          `json:"partition_field"`   // example:`canal_waterlevel_datetime`ชื่อคอลัมน์ที่ใช้ในการตรวจสอบพาทิชัน
	RowValidator      string          `json:"row_validator"`     // example:`!is_nil(canal_waterlevel_value)`การตั้งค่า reject null ของข้อมูล
	ConvertName       string          `json:"convert_name"`      // example:`ข้อมูลคลอง`ชื่อชุดข้อมูล
	ConvertSetting    json.RawMessage `json:"convert_setting"`   // example:`{"data_folder":"dataset/bma/canal","configs":[{"name":"canal","input_name":"canal.txt","header_row":1,"data_tag":"","row_validator":"","fields":[{"name":"canal_station_id","type":"int","input_fields":["0"],"transform_method":"mapping","transform_params":{"add_missing":false,"from":["canal_station_oldcode"],"table":"m_canal_station","to":"id"}},{"name":"canal_waterlevel_datetime","type":"string","input_fields":["1","2"],"transform_method":"datetime","transform_params":{"input_format":"%d/%m/%Y %H:%M:%S"}},{"name":"canal_waterlevel_value","type":"float","input_fields":["3"],"transform_method":"","transform_params":""},{"name":"comm_status","type":"string","input_fields":["4"],"transform_method":"","transform_params":""},{"name":"qc_status","type":"","transform_method":"qc","transform_params":""}]}]}`ตั้งค่าการคอนเวิร์ท
	ImportSetting     json.RawMessage `json:"import_setting"`    // example:`{"configs":[{"name":"canal","imports":[{"source":"canal.txt","destination":"canal_waterlevel"}]}]}`ตั้งค่าการอิมพอร์ต
	LookupTable       json.RawMessage `json:"lookup_table"`      // example:`{"tables":{"m_canal_station":{"allow_add_missing":false,"fields":["id","canal_station_oldcode"]}}}`ชื่อตารางที่ใช้ในการ map ข้อมูล
	ImportTable       json.RawMessage `json:"import_table_json"` // example:`{"tables":{"canal_waterlevel":{"fields":["#row","canal_station_id","canal_waterlevel_datetime","canal_waterlevel_value","comm_status","qc_status"],"partition_field":"canal_waterlevel_datetime","unique_constraint":"uk_canal_waterlevel"}}}`ชื่อตารางที่ใช้ในการอิมพอร์ตข้อมูล
}

type DataDetailsSwagger struct {
	SelectOpt []*DownloadDatasetListIDSwagger `json:"select_option"` // selectoption
	MetaData  []*MetadataListNameSwagger      `json:"metadata_list"` // metadata list
}
type DownloadDatasetListIDSwagger struct {
	DownloadID    int64     `json:"id"`           // example:`18`รหัส download
	DownloadName  string    `json:"name"`         // example:`bma-canal`ชื่อ download
	DatasetListID []*DSName `json:"dataset_list"` // dataset list
}

type MetadataListNameSwagger struct {
	//	MonitorScript string           `json:"monitor_script"`
	DownloadName *DLName  `json:"download_name"` // ชื่อ download
	DatasetName  *DSName  `json:"dataset_name"`  // ชื่อ dataset
	MetadataName *MDName  `json:"metadata_name"` // ชื่อ metadata
	AgencyName   *AGCName `json:"agency_name"`   // ชื่อหน่วยงาน
}

type AGCName struct {
	Text  string `json:"text"`  // example:`HAII` agency name
	Value int64  `json:"value"` // example:`1` agency id
}

type DLName struct {
	Text  string `json:"text"`  // example:`dl-canal-bma` download name
	Value int64  `json:"value"` // example:`1` download id
}

type DSName struct {
	Text  string `json:"text"`  // example:`cv-canal-bma` dataset name
	Value int64  `json:"value"` // example:`1` dataset id
}
type MDName struct {
	Text  string `json:"text"`  // example:`ข้อมูลคลอง` metadata name
	Value int64  `json:"value"` // example:`1` metadata id
}

type DataDetails struct {
	SelectOpt *SelectOption       `json:"select_option"` // selectoption
	MetaData  []*MetadataListName `json:"metadata_list"` // metadata list
}

type SelectOption struct {
	DownloadDatasetList    *result.Result `json:"download_list"`     // download list
	AllDownloadDatasetList *result.Result `json:"all_download_list"` // download list
}

type MetadataListName struct {
	//	MonitorScript string           `json:"monitor_script"`
	DownloadName      *MetadataDetails `json:"download_name"`      // ชื่อ download
	DatasetName       *MetadataDetails `json:"dataset_name"`       // ชื่อ dataset
	MetadataName      *MetadataDetails `json:"metadata_name"`      // ชื่อ metadata
	AgencyName        *MetadataDetails `json:"agency_name"`        // ชื่อหน่วยงาน
	AdditionalDataset []string         `json:"additional_dataset"` // รหัส dataset ที่เกี่ยวข้อง (ถ้ามี)
}

type MetadataDetails struct {
	Text  interface{} `json:"text"`  // ข้อความ
	Value interface{} `json:"value"` // ค่า
}

type MetadataDescription struct {
	MetadataID        int64    `json:"metadata_id"` // รหัส metadata
	DownloadID        int64    `json:"download_id"` // รหัส download
	DatasetID         int64    `json:"dataset_id"`  // รหัส dataset
	MonitorScript     string   `json:"monitor_script"`
	AdditionalDataset []string `json:"additional_dataset"` // รหัส dataset ที่เกี่ยวข้อง (ถ้ามี)
}

type DownloadDatasetListID struct {
	DownloadID    int64              `json:"id"`           // example:`18`รหัส download
	DownloadName  string             `json:"name"`         // example:`bma-canal`ชื่อ download
	DatasetListID []*MetadataDetails `json:"dataset_list"` // dataset list
}

type HistorySelectOptList struct {
	SelectOpt *SelectOptionHistory `json:"select_option"` // selectoption
	DateRange int64                `json:"date_range"`    // example:`30`ช่วงวันที่
}

type HistorySelectOptListSwagger struct {
	SelectOpt *SelectOptionHistorySwagger `json:"select_option"` // selectoption
	DateRange int64                       `json:"date_range"`    // example:`30`ช่วงวันที่
}

type SelectOptionHistorySwagger struct {
	Agency   []*AgencyWithMetadata `json:"agency_id"`      // หน่วยงาน
	Status   []*STSwagger          `json:"process_status"` // สถานะ
	Metadata []*MDName             `json:"metadata"`       // บัญชีข้อมูล
}

type AgencyWithMetadata struct {
	Text     string `json:"text"`        // example:`HAII`ข้อความ
	Value    int64  `json:"value"`       // example:`9`ค่า
	Metadata int64  `json:"metadata_id"` // example:`1`รหัส metadata
}

type STSwagger struct {
	Text  string `json:"text"`  // example:`Download Failed` ข้อความ
	Value string `json:"value"` // example:`0` ค่า
}

type SelectOptionHistory struct {
	Agency   *result.Result     `json:"agency_id"`      // หน่วยงาน
	Status   *result.ResultJson `json:"process_status"` // สถานะ
	Metadata *result.Result     `json:"metadata"`       // บัญชีข้อมูล
}

type MetadataNameInAgency struct {
	Text     interface{} `json:"text"`        // example:`HAII`ข้อความ
	Value    interface{} `json:"value"`       // example:`9`ค่า
	Metadata interface{} `json:"metadata_id"` // example:`1`รหัส metadata
}

type HistoryList struct {
	MetaData []*HistoryMetadata `json:"history_list"` // ประวัติการรัน Script
}

type HistoryMetadata struct {
	DownloadID               int64           `json:"dataimport_download_id"`     // example:`1` รหัส download
	DownloadLogID            int64           `json:"dataimport_download_log_id"` // example:`1321`  รหัส download log
	DatasetID                int64           `json:"dataimport_dataset_id"`      // example:`1`  รหัส dataset
	DatasetLogID             int64           `json:"dataimport_dataset_log_id"`  // example:`843`  รหัส dataset log
	DownloadBeginAt          string          `json:"download_begin_at"`          // example:`2006-01-02 15:14`  วันเวลาเริ่มต้น download
	DownloadEndAt            string          `json:"download_end_at"`            // example:`2006-01-02 15:14`  วันเวลาสิ้นสุด download
	ConvertBeginAt           string          `json:"convert_begin_at"`           // example:`2006-01-02 15:14`  วันเวลาเริ่มต้น convert
	ConvertEndAt             string          `json:"convert_end_at"`             // example:`2006-01-02 15:14`  วันเวลาสิ้นสุด convert
	ImportBeginAt            string          `json:"import_begin_at"`            // example:`2006-01-02 15:14`  วันเวลาเริ่มต้น import
	ImportEndAt              string          `json:"import_end_at"`              // example:`2006-01-02 15:14`  วันเวลาสิ้นสุด import
	ProcessStatus            int64           `json:"process_status"`             // example:`3`  สถานะ
	MetadataServicName       json.RawMessage `json:"metadataservice_name"`       // example:`{"th":"ข้อมูลระดับน้ำในคลอง"}`  ชื่อบัญชีข้อมูล
	MetadataConvertFrequency string          `json:"metadata_convert_frequency"` // example:`15 นาที`  ความถี่การ convert
	MetadataChannel          string          `json:"metadata_channel"`           // example:`web service` ช่องทางการเชื่อมโยง
	AgencyName               json.RawMessage `json:"agency_name"`                // example:`{"en":"bma"}`  ชื่อหน่วยงาน
	Filesize                 int64           `json:"filesize"`                   // example:`5543`  ขนาดไฟล์
	ReRunFlag                string          `json:"rerun_flag"`                 // example:`cv`  สถานะล่าสุด
	DownloadScript           string          `json:"download_method"`            // example:`dl-basic`  กระบวนการ download
	Duration                 int64           `json:"duration"`                   // example:`532134`  ระยะเวลา(วินาที)
	EventCode                json.RawMessage `json:"event_code"`                 // example:`1`  เหตุการณ์
}

type HistoryDataSelect struct {
	AgencyID      []int64 `json:"agency_id"`      // example:[9] รหัสหน่วยงาน
	MetadataID    []int64 `json:"metadata_id"`    // example:[512] รหัสบัญชีข้อมูล
	ProcessStatus []int64 `json:"process_status"` // example:[-1,0,1,2,3] สถานะ
	BeginAt       string  `json:"begin_at"`       // example:`2017-08-25 00:04` วันเวลาเริ่มต้น
	EndAt         string  `json:"end_at"`         // example:`2017-08-25 15:04` วันเวลาสิ้นสุด
}

type CronEnabled struct {
	DownloadID    string `json:"dataimport_download_id"` // รหัส download
	IsCronenabled bool   `json:"is_cronenabled"`         // สถานะเปิดปิด cron
}

type NewRdlID struct {
	DownloadID string `json:"download_id"` // example:`1`รหัส download
	Agency     string `json:"agency"`      // example:`haii`รหัสหน่วยงาน
}

type DownloadCronSetting struct {
	DownloadID     int64  `json:"download_id"`     // example:`1`รหัส download
	IsCronenabled  bool   `json:"is_cronenabled"`  // สถานะเปิดปิด cron
	DownloadScript string `json:"download_method"` // กระบวนการดาวน์โหลด
	DownloadName   string `json:"download_name"`   // ชื่อการดาวน์โหลด
	Description    string `json:"description"`     // คำอธิบาย
	CrontabSetting string `json:"crontab_setting"` // การตั้งค่าเวลาของ crontab
	Node           string `json:"node"`            // node
	Agency         string `json:"agency"`          // Agency
}

type ConfigVariable struct {
	ID           int64  `json:"id"`
	Category     string `json:"category"`
	ConfigName   string `json:"name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
}

type ListVariable struct {
	ID           int64  `json:"id"`
	NameCategory string `json:"name_category"`
}
