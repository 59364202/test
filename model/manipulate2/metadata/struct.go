package metadata

import ()

type MetadataData_struct struct {
	Id                       int64  `json:"id"`
	SubcategoryId            int64  `json:"subcategory_id"`
	AgencyId                 int64  `json:"agency_id"`
	DataunitId               int64  `json:"dataunit_id"`
	DataformatId             int64  `json:"dataformat_id"`
	ServiceMethodId          int64  `json:"servicemethod_id"`
	ConnectionFormat         string `json:"connection_format"`
	MetadataMethod           string `json:"metadata_method"`
	MetadataChannel          string `json:"metadata_channel"`
	MetadaDatafrequency      string `json:"metadata_datafrequency"`
	MetadataConvertfrequenct string `json:"metadata_convertfrequency"`
	MetadataContact          string `json:"metadata_contact"`
	MetadataAgencystoredate  string `json:"metadata_agencystoredate"`
	MetadataStartdatadate    string `json:"metadata_startdatadate"`
	Scriptname               string `json:"scriptname"`
	MetadataUpdatePlan       string `json:"metadata_update_plan"`
	MetadataLaws             string `json:"metadata_laws"`
	MetadataRemark           string `json:"metadata_remark"`
	MetadataStatus           string `json:"metadata_status"`
}

type MetadataTable_struct struct {
	Id                  int64             `json:"id"`
	SubcategoryId       int64             `json:"subcategory_id"`
	AgencyId            int64             `json:"agency_id"`
	MetadataServiceName map[string]string `json:"metadataservice_name"`
	MetadataAgencyName  map[string]string `json:"metadataagency_name"`
}

type Metadata_struct struct {
	Data    interface{}       `json:"metadata"`
	Hydro   []int64           `json:"hydro"`
	History []*History_struct `json:"history"`
}

type Hydro_struct struct {
	Id         int64  `json:"id"`
	MetadataId int64  `json:"metadata_id"`
	CreatedBy  string `json:"create_by"`
}

type History_struct struct {
	HistoryDatetime    string `json:"history_datetime"`
	HistoryDescription string `json:"history_description"`
	CreatedBy          string `json:"created_by"`
}
