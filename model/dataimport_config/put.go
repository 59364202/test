// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataimport_config is a model for api.dataimport_download table. This table store dataimport config information.
package dataimport_config

import (
	"encoding/json"
	"strings"

	"haii.or.th/api/server/model/dataimport"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"

	//	"haii.or.th/api/util/shell"
	"strconv"
)

// updata dataimport download config
//  Parameters:
//		downloadConfig
//			 DataDownloadConfig
//		uid
//			user id
//  Return:
//		NewRdlID
func UpdateDataImportConfig(downloadConfig *DataDownloadConfig, uid int64) (*NewRdlID, error) {
	runDownloadJson, err := editJsonDownloadConfig(downloadConfig, uid, updateDownloadConfig)
	if err != nil {
		return nil, err
	}

	return runDownloadJson, nil
}

// update json dataimport download config
//  Parameters:
//		downloadConfig
//			 DataDownloadConfig
//		uid
//			user id
//		q
//			sql for dataimport download config
//  Return:
//		NewRdlID
func editJsonDownloadConfig(downloadConfig *DataDownloadConfig, uid int64, q string) (*NewRdlID, error) {

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

	// value for edit
	if downloadConfig.Node == "" {
		downloadConfig.Node = "276"
	}

	p := []interface{}{downloadConfig.AgentUserID, dst, uid, downloadConfig.DownloadScript, downloadConfig.DownloadConfigID, downloadConfig.CrontabSetting,
		downloadConfig.DownloadName, downloadConfig.Description, downloadConfig.Node, downloadConfig.MaxProcess, downloadConfig.IsCronEnabled}

	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	// new download id
	data, err := NewRunDownloadID(downloadConfig.AgentUserID, downloadConfig.DownloadConfigID)
	if err != nil {
		return nil, err
	}

	// return data
	return data, nil
}

// update dataimport config dataset
//  Parameters:
//		datasetConfig
//			 DataImportDataSetConfig
//		uid
//			user id
//  Return:
//		dataimport dataset id
func UpdateDataImportConfigDataset(datasetConfig *DataImportDataSetConfig, uid int64) (int64, error) {
	dataimportDownloadID, err := editJsonDatasetConfig(datasetConfig, uid, updateDatasetConfig)
	if err != nil {
		return 0, err
	}

	return dataimportDownloadID, nil
}

// update  dataset config
//  Parameters:
//		datasetConfig
//			 DataImportDataSetConfig
//		uid
//			user id
//		q
//			sql for dataimport dataset config
//  Return:
//		dataimport dataset id
func editJsonDatasetConfig(CnvImConfig *DataImportDataSetConfig, uid int64, q string) (int64, error) {

	//	CnvImConfig.ConvertSetting.Configs[0].RowValidator = CnvImConfig.RowValidator
	b, err := json.Marshal(CnvImConfig.ConvertSetting)
	if err != nil {
		return 0, err
	}
	convertSetting := string(b[:])

	// end convert setting

	// make data to json for import setting
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

	// make data to json for  lookup table
	//lt := makeLookupTable(CnvImConfig.ConvertSetting) // function เก่า  ไม่ได้ใช้งาน
	lt := setLookupTable(CnvImConfig.ImportDestination, CnvImConfig.ConvertSetting) //function ใหม่ , Support lookup table for QC rule

	b, err = json.Marshal(lt)
	if err != nil {
		return 0, err
	}
	lookupTable := string(b[:])
	// end lookup table

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// importSetting, lookupTable, importTable
	p := []interface{}{CnvImConfig.AgentUserID, CnvImConfig.DownloadConfigID, convertSetting, importSetting, lookupTable, importTable, uid, CnvImConfig.ConvertScript, CnvImConfig.ImportScript, CnvImConfig.ConvertName, CnvImConfig.DataSetConfigID}

	// connect database server
	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return 0, nil
}

// func update metadata
//  Parameters:
//		metaDesc
//			 MetadataDescription
//		page
//			for page migrate or not
//		uid
//			user id
//  Return:
//		dataimport dataset id
func UpdateMetadata(metaDesc *MetadataDescription, page string, uid int64) error {
	var q string
	p := []interface{}{metaDesc.MetadataID, uid, metaDesc.DownloadID, metaDesc.DatasetID}
	// prepare query, data to update
	if page == "migrate" {
		q = updateMetaDataProvision
	} else {
		q = updateMetaData
		if len(metaDesc.AdditionalDataset) > 0 {
			p = append(p, strings.Join(metaDesc.AdditionalDataset, ",")) // ["1","2"] => "1,2"
		} else {
			p = append(p, nil)
		}
	}

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

// func update cron field
//  Parameters:
//		dataimport_download_id
//			 dataimport download id for update
//		uid
//			user id
//		is_cronenabled
//			is cronenabled change true or false
//  Return:
//		None
func UpdateIsCronEnabled(dataimport_download_id string, uid int64, is_cronenabled bool) error {
	downloadID, _ := strconv.Atoi(dataimport_download_id)
	q := updateIsCronenabled

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	// add value for update
	p := []interface{}{uid, downloadID, is_cronenabled}

	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func UpdateDownloadCronList(dataimport_download_id string, uid int64, crontab_setting string) error {
	downloadID, _ := strconv.Atoi(dataimport_download_id)
	q := updateDownloadCron

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// prepare sql
	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	// add value for update
	p := []interface{}{uid, downloadID, crontab_setting}

	// execute data
	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func UpdateApiCronList(name, crontab_setting string) error {
	s := map[string]interface{}{name: crontab_setting}
	setting.SetSystemSetting(setting.SystemUserID, s, true)
	return nil
}

type Struct_category2 struct {
	VariableID   int64  `json:"variable_id"`
	Category     int64  `json:"category"`
	Name         string `json:"name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
}

func UpdateConfigVariable(category int64, user, name, value string, uid, vid int64) (*Struct_category2, error) {
	q := `UPDATE api.config_variable SET category = $1, config_name = $2, variable_name = $3 , value = $4,updated_by = $5, updated_at = NOW() WHERE id = $6`

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(q)
	if err != nil {
		return nil, err
	}

	// ctime := time.Now()
	userid := uid

	_, err = stmt.Exec(category, user, name, value, userid, vid)
	if err != nil {
		return nil, errors.Repack(err)
	}

	tx.Commit()
	if err != nil {
		return nil, errors.Repack(err)
	}

	return nil, nil
}
