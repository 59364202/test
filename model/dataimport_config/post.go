// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataimport_config is a model for api.dataimport_download table. This table store dataimport config information.
package dataimport_config

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"haii.or.th/api/server/model/dataimport"
	map_struct "haii.or.th/api/thaiwater30/model/metadata"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
)

// func add dataimport download config
//  Parameters:
//		downloadConfig
//			 DataDownloadConfig
//		uid
//			user id
//  Return:
//		NewRdlID
func AddDataImportDownloadConfig(downloadConfig *DataDownloadConfig, uid int64) (*NewRdlID, error) {
	runDownloadJson, err := addJsonDownloadConfig(downloadConfig, uid, insDownloadConfig)
	if err != nil {
		return nil, err
	}
	return runDownloadJson, nil
}

// func json download config for insert
//  Parameters:
//		downloadConfig
//			 DataDownloadConfig
//		uid
//			user id
//		q
//			sql for dataimport download config
//  Return:
//		NewRdlID
func addJsonDownloadConfig(downloadConfig *DataDownloadConfig, uid int64, q string) (*NewRdlID, error) {
	dss, err := dataimport.CipherDownloadPassword(&downloadConfig.DownloadSetting, true)
	if err != nil {
		log.Logf("Can not encrypt download password ...%v", err)
	}

	b, err := json.Marshal(dss)
	if err != nil {
		return nil, err
	}
	dst := string(b[:])
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	if downloadConfig.Node == "" {
		downloadConfig.Node = "276"
	}
	// value for insert
	p := []interface{}{downloadConfig.AgentUserID, dst, uid, downloadConfig.DownloadScript, downloadConfig.CrontabSetting,
		downloadConfig.DownloadName, downloadConfig.Description, downloadConfig.Node, downloadConfig.MaxProcess, downloadConfig.IsCronEnabled}

	var download_id int64
	// insert data and return id
	err = db.QueryRow(q, p...).Scan(&download_id)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	// created json for new rundownload
	data, err := NewRunDownloadID(downloadConfig.AgentUserID, download_id)
	if err != nil {
		return nil, err
	}

	// return data
	return data, nil
}

// func dataimport dataset config
//  Parameters:
//		datasetConfig
//			 DataImportDataSetConfig
//		uid
//			user id
//  Return:
//		dataimport dataset id
func AddDataImportDatasetConfig(datasetConfig *DataImportDataSetConfig, uid int64) (int64, error) {
	dataimportDownloadID, err := addJsonDatasetConfig(datasetConfig, uid, insDatasetConfig)
	if err != nil {
		return 0, err
	}

	return dataimportDownloadID, nil
}

// func json dataimport dataset config
//  Parameters:
//		datasetConfig
//			 DataImportDataSetConfig
//		uid
//			user id
//		q
//			sql for dataimport dataset config
//  Return:
//		dataimport dataset id
func addJsonDatasetConfig(CnvImConfig *DataImportDataSetConfig, uid int64, q string) (int64, error) {

	//	CnvImConfig.ConvertSetting.Configs[0].RowValidator = CnvImConfig.RowValidator
	b, err := json.Marshal(CnvImConfig.ConvertSetting)
	if err != nil {
		return 0, err
	}
	convertSetting := string(b[:])
	// end convert setting

	//make data to json for import setting
	ims := makeImportSetting(CnvImConfig.ConvertSetting, CnvImConfig.ImportDestination)

	b, err = json.Marshal(ims)
	if err != nil {
		return 0, err
	}
	importSetting := string(b[:])
	// end import table

	// make data to json for import table
	imt := makeImportTable(CnvImConfig.ConvertSetting, CnvImConfig.ImportDestination, CnvImConfig.UniqueConstraint, CnvImConfig.PartitionField)

	b, err = json.Marshal(imt)
	if err != nil {
		return 0, err
	}
	importTable := string(b[:])
	// end import table

	// make data to json for lookup table
	//lt := makeLookupTable(CnvImConfig.ConvertSetting) // function เก่า  ไม่ได้ใช้งาน
	lt := setLookupTable(CnvImConfig.ImportDestination, CnvImConfig.ConvertSetting) //function ใหม่ , Support lookup table for QC rule

	b, err = json.Marshal(lt)
	if err != nil {
		return 0, err
	}
	lookupTable := string(b[:])
	// end lookup table

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// importSetting, lookupTable, importTable
	p := []interface{}{CnvImConfig.AgentUserID, CnvImConfig.DownloadConfigID, convertSetting, importSetting, lookupTable, importTable, uid, CnvImConfig.ConvertScript, CnvImConfig.ImportScript, CnvImConfig.ConvertName}

	var dataset_id int64
	// insert data and return id
	err = db.QueryRow(q, p...).Scan(&dataset_id)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	// return id
	return dataset_id, nil
}

// func make data to import table
//  Parameters:
//		ConvertSetting
//			 ConvertSettingStruct
//		importDestination
//			 import table
//		UniqueConstraint
//			 unique constraint table import
//		PartitionField
//			partition field example canal_waterlevel_datetime
//  Return:
//		json for field import_table
func makeImportTable(ConvertSetting ConvertSettingStruct, importDestination, UniqueConstraint, PartitionField string) interface{} {

	type Tables struct {
		Tables map[string]interface{} `json:"tables"`
	}

	var fields []string
	fields = append(fields, "#row")
	for _, v := range ConvertSetting.Configs {
		tif := v.Fields
		for _, k := range tif {
			fields = append(fields, k.Name)
		}
	}
	tableValue := map[string]interface{}{}
	tableValue["unique_constraint"] = UniqueConstraint
	tableValue["partition_field"] = PartitionField
	tableValue["fields"] = fields

	tables := map[string]interface{}{}
	tables[importDestination] = tableValue

	tableCfg := &Tables{}
	tableCfg.Tables = tables
	return tableCfg
}

// func make data to lookup table
//  Parameters:
//		ConvertSetting
//			 ConvertSettingStruct
//  Return:
//		json for field lookup_table
func makeLookupTable(ConvertSetting ConvertSettingStruct) interface{} {

	tables := map[string]interface{}{}
	evalLookup, _ := regexp.Compile("lookup(.*)")
	// check tranform method
	for _, v := range ConvertSetting.Configs {
		tif := v.Fields
		for _, k := range tif {
			if k.TransformMethod != "" {
				// transform method mapping find field name
				if k.TransformMethod == "mapping" {
					tfp := k.TransformParams.(map[string]interface{})
					var field []string
					field = append(field, tfp["to"].(string))
					for _, o := range tfp["from"].([]interface{}) {
						field = append(field, o.(string))
					}
					fields := map[string]interface{}{}
					if tfp["add_missing"] != nil {
						fields["allow_add_missing"] = tfp["add_missing"].(bool)
					}
					fields["fields"] = field
					tables[tfp["table"].(string)] = fields
					// transform method evalute find func loopup
				} else if k.TransformMethod == "evaluate" {
					sParams := k.TransformParams.(string)
					// find field name from func lookup
					if len(evalLookup.FindString(sParams)) > 0 {
						//						s := strings.Split(sParams, ",")
						//						s[0] = strings.TrimLeft(s[0], "lookup(")
						//						s[len(s)-1] = strings.TrimLeft(s[len(s)-1], ")")
						//						var field []string
						//						field = append(field, s[1])
						//						fields := map[string]interface{}{}
						//						for i := 2; i < len(s); i++ {
						//							if (float64(i) / 2) != 0 {
						//								field = append(field, s[i])
						//							}
						//						}
						var field []string
						fields := map[string]interface{}{}
						sPrms := strings.TrimRight(strings.TrimLeft(sParams, "lookup("), ")") // ตัดคำว่า "lookup(" , ")" ออกจากทั้งซ้าย และ ขวา
						s := strings.Split(sPrms, ",")
						field = append(field, TrimSingleQuote(s[1]))
						for i := 2; i < len(s); i = i + 2 {
							field = append(field, TrimSingleQuote(s[i]))
						}

						fields["allow_add_missing"] = false
						fields["fields"] = field
						tables[TrimSingleQuote(s[0])] = fields
					}
				}
			}
		}
	}
	// add field to struct
	type Tables struct {
		Tables map[string]interface{} `json:"tables"`
	}

	tableCfg := &Tables{}
	tableCfg.Tables = tables

	return tableCfg
}

// New func make data to lookup table replace func makeLookuptable
//  Parameters:
//		table_name
//  Return:
//		json for field lookup_table
func setLookupTable(table_name string, ConvertSetting ConvertSettingStruct) interface{} {

	type Tables struct {
		Tables map[string]interface{} `json:"tables"`
	}

	tableCfg := &Tables{}

	st_table := map_struct.GetTable(table_name) // รับ struct ของ table ที่จะนำเข้าข้อมูล  ->> model/metadata/map.go
	if st_table == nil {
		return tableCfg
	}

	tables := map[string]interface{}{}

	//get struct of master table
	if st_table.IsMaster == false && st_table.MasterTable != "" {
		table_name = st_table.MasterTable

		if st_table = map_struct.GetTable(table_name); st_table == nil {
			return tableCfg
		}
	}

	//table master

	table_master := get_field_lookup(table_name)
	if table_master != nil {
		tables[table_name] = get_field_lookup(table_name)
	}

	//lt_geocode
	if st_table.HasProvince {
		tables["lt_geocode"] = get_field_lookup("lt_geocode")
	}

	//basin
	if st_table.HasBasin {
		tables["basin"] = get_field_lookup("basin")
	}

	//master table
	if st_table.MasterTable != "" {
		tables[st_table.MasterTable] = get_field_lookup(st_table.MasterTable)
	}

	/*  Setting lookup table from mapping  */
	evalLookup, _ := regexp.Compile("mapping.*")
	var field []string
	//	fmt.Println("Test -> " + st_cv.)
	// check tranform method
	for _, v := range ConvertSetting.Configs {
		tif := v.Fields
		for _, k := range tif {
			if k.TransformMethod != "" {
				// transform method mapping find field name
				if k.TransformMethod == "mapping" || k.TransformMethod == "mappingnil" {
					tfp := k.TransformParams.(map[string]interface{})

					if tables[tfp["table"].(string)] == nil {
						tables[tfp["table"].(string)] = map[string]interface{}{
							"fields": []string{},
						}
						//						fields["fields"] = []string{}
					}

					fields := tables[tfp["table"].(string)].(map[string]interface{})
					field = fields["fields"].([]string)

					//					fmt.Println("table:", tables["m_dam"])

					field = append(field, tfp["to"].(string))
					for _, o := range tfp["from"].([]interface{}) {
						field = append(field, o.(string))
					}

					cleaned := []string{}

					for _, value := range field {
						//						fmt.Println(value)
						if !stringInSlice(value, cleaned) {
							cleaned = append(cleaned, value)
						}
					}

					fields["fields"] = cleaned
					tables[tfp["table"].(string)] = fields
					//					fmt.Println("table:", tables["m_dam"])

					// transform method evalute find func loopup
				} else if k.TransformMethod == "evaluate" {
					sParams := k.TransformParams.(string)
					// find field name from func lookup
					if len(evalLookup.FindString(sParams)) > 0 {
						var field []string
						fields := map[string]interface{}{}
						sPrms := strings.TrimRight(strings.TrimLeft(sParams, "lookup("), ")") // ตัดคำว่า "lookup(" , ")" ออกจากทั้งซ้าย และ ขวา
						s := strings.Split(sPrms, ",")
						field = append(field, TrimSingleQuote(s[1]))
						for i := 2; i < len(s); i = i + 2 {
							field = append(field, TrimSingleQuote(s[i]))
						}

						fields["allow_add_missing"] = false
						fields["fields"] = field
						tables[TrimSingleQuote(s[0])] = fields
					}
				}
			}
		}
	}
	/*  setting lookup table from mapping End  */

	//	tables["m_tables"] = "master"
	tableCfg.Tables = tables

	// make data to json for lookup table
	//	b, _ := json.Marshal(tableCfg)
	//	lookupTable := string(b[:])
	// end lookup table

	//fmt.Println(lookupTable)

	return tableCfg

}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		v = strings.Trim(v, " ")
		str = strings.Trim(str, " ")
		if v == str {
			return true
		}
	}
	return false
}

func get_field_lookup(table_name string) interface{} {

	fields := map[string]interface{}{}
	st_table := map_struct.GetTable(table_name) // รับ struct ของ table นั้นๆ  ->> model/metadata/map.go

	if st_table.Fields == "" {
		return nil
	}

	if st_table == nil {
		return fields
	}
	strF := strings.Replace(st_table.Fields, " ", "", -1) // ลบช่องว่างออก
	arr_field_name := strings.Split(strF, ",")
	fields["fields"] = arr_field_name

	//เพิ่ม allow_add_missing สำหรับ table ที่ขึ้นต้นด้วย "m" เช่น m_dam
	fst_cha := string(table_name[0]) //รับตัวอักษรตัวแรก

	if fst_cha == "m" {
		fields["allow_add_missing"] = true
	}

	return fields
}

//	ตัด ' ออกจากทางซ้าย, ขวา ของ string
//	Parameters:
//		s
//			string
//	Return
//		'str' -> str
func TrimSingleQuote(s string) string {
	return strings.TrimRight(strings.TrimLeft(s, "'"), "'")
}

// func make data to import setting
//  Parameters:
//		ConvertSetting
//			 ConvertSettingStruct
//		importDestination
//			 import table
//  Return:
//		json for field import_setting
func makeImportSetting(ConvertSetting ConvertSettingStruct, importDestination string) interface{} {

	type Imports struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
	}

	type Config struct {
		Name    string     `json:"name"`
		Imports []*Imports `json:"imports"`
	}

	type ConfigList struct {
		Configs []*Config `json:"configs"`
	}

	ims := &Imports{}
	cfgs := &Config{}
	// add config to struct
	cfgSetting := ConvertSetting.Configs
	cfg := cfgSetting[0]

	// get source filename
	ims.Source = cfg.InputName
	ims.Destination = importDestination

	imsList := make([]*Imports, 0)
	imsList = append(imsList, ims)
	// config name and import list to struct
	cfgs.Name = cfg.Name
	cfgs.Imports = imsList

	cfgsList := make([]*Config, 0)
	cfgsList = append(cfgsList, cfgs)

	cfgl := &ConfigList{}

	cfgl.Configs = cfgsList

	return cfgl
}

// func new rundownload id to json
//  Parameters:
//		uid
//			 user id for download config
//		downloadID
//			 dataimport download id
//  Return:
//		NewRdlID
func NewRunDownloadID(uid, downloadID int64) (*NewRdlID, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql agent name
	q := getAgentName
	p := []interface{}{uid}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, err
	}
	var user_account string
	for rows.Next() {

		rows.Scan(&user_account)
	}

	// split data get agency name
	agency := strings.Split(user_account, "-")
	if len(agency) < 2 {
		return nil, nil
	}

	data := &NewRdlID{}
	data.Agency = agency[len(agency)-1]
	data.DownloadID = strconv.Itoa(int(downloadID))

	// return data
	return data, nil
}

// copy dataimport download config
//  Parameters:
//		downloadID
//			 dataimport download id
//  Return:
//		new dataimport download id
func CopyDataimportDownloadConfig(downloadID int64) (int64, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql  download config
	q := cpDownloadConfig
	p := []interface{}{downloadID}

	var download_id int64
	// insert data
	err = db.QueryRow(q, p...).Scan(&download_id)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	// return download id
	return download_id, nil
}

// copy dataset config
//  Parameters:
//		datasetID
//			 dataimport dataset id
//  Return:
//		new dataimport dataset id
func CopyDataimportDatasetConfig(datasetID int64) (int64, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql copy dataset config
	q := cpDatasetConfig
	p := []interface{}{datasetID}

	var dataset_id int64

	// insert dataset
	err = db.QueryRow(q, p...).Scan(&dataset_id)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	// return dataset id
	return dataset_id, nil
}
