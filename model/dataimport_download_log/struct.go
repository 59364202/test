package dataimport_download_log

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
)

type Field_NumberOfDownloadSize struct {
	NumberOfDownloadSize float64 `json:"download_size"` // example:`0` ขนาดไฟล์
}
type Field_NumberOfFileDownload struct {
	NumberOfFileDownload int64 `json:"number_of_file"` // example:`15` จำนวนไฟล์
}
type Field_NumberOfDataRecord struct {
	NumberOfDataRecord int64 `json:"number_of_record"` // example:`15` จำนวน record
}
type Field_DownloadDate struct {
	DownloadDate string `json:"download_date"` // example:`2006-01-02` วันที่ download
}

type Struct_DownloadLog struct {
	Field_DownloadDate
	Field_NumberOfDataRecord
	Field_NumberOfDownloadSize
	Field_NumberOfFileDownload
}

type Struct_DownloadLog_Summary struct {
	NumberOfExpectedDownload    float64 `json:"download_count_expected"` // example:`2` ประมาณการนำเข้าข้อมูล (ครั้ง/1วัน)
	NumberOfDownloadCountActual int64   `json:"download_count_actual"`   // example:`15` จำนวนการดาวน์โหลดที่เกิดขึ้นจริง
	NumberOfDownloadCount       int64   `json:"download_count"`          // example:`0` นำเข้าข้อมูลแล้ว
	PercentDownloadCount        float64 `json:"download_count_percent"`  // example:`0` จำนวน download คิดเป็น %

	DownloadLastestDate string `json:"download_lastest_date"` // example:`2006-01-02` วันที่ download ล่าสุด

	Field_DownloadDate
	Field_NumberOfDataRecord
	Field_NumberOfFileDownload
}
type Struct_DownloadLog_Summary_Agency struct {
	Struct_DownloadLog_Summary
	Agency *model_agency.Struct_Agency `json:"agency,omitempty"` // หน่วยงาน
}

type Struct_DownloadLog_Summary_Metadata struct {
	Struct_DownloadLog_Summary
	Metadata *model_metadata.Struct_Metadata `json:"metadata,omitempty"` // บัญชีข้อมูล
}

type Struct_DownloadLog_YearlyCompare struct {
	Year                 int64     `json:"year"`                   // example:`2006` ปี
	PercentDownloadCount []float64 `json:"download_count_percent"` // example:`[0,1,2,3]` จำนวนที่ download คิดเป็น %
	NumberOfFileDownload []int64   `json:"number_of_file"`         // example:`[0,1,2,3]` จำนวนไฟล์
	NumberOfDataRecord   []int64   `json:"number_of_record"`       // example:`[0,1,2,3]` จำนวน record
}

//Input Data => Used by BOF-DataIntegrationReport-Overall Screen
type Struct_DownloadLog_Inputparam struct {
	AgencyID         string `json:"agency_id"`         // example:57 รหัสหน่วยงาน
	Month            string `json:"month"`             // required:false example:1 เดือน
	Year             string `json:"year"`              // example:2006 ปี
	StartDate        string `json:"start_date"`        // '2016-12-31'
	EndDate          string `json:"end_date"`          // '2016-12-31'
	ConnectionFormat string `json:"connection_format"` // example:online ประเภทการเชื่อมต่อ online, offline
}
