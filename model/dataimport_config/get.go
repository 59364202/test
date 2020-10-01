// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataimport_config is a model for api.dataimport_download table. This table store dataimport config information.
package dataimport_config

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/cronjob"
	"haii.or.th/api/server/model/dataimport"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
)

// get dataimport download config by id
//  Parameters:
//		dataimportDownloadID
//				download id for get information
//  Return:
//		dataimport download config
func GetDataImportDownloadConfig(dataimportDownloadID string) (*DataImportDownloadConfig, error) {
	downloadID, err := strconv.Atoi(dataimportDownloadID)

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get download config
	q := getDownloadConfig
	p := []interface{}{downloadID}

	// process sql get dataimport download
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	datadlcfg := &DataImportDownloadConfig{}
	// loop dataimport download config
	for rows.Next() {
		// declare value data from database
		var downloadSetting pqx.JSONRaw
		var downloadName sql.NullString
		var crontabSetting sql.NullString
		var downloadScript sql.NullString
		var isCronenabled sql.NullBool
		var description sql.NullString
		var node sql.NullString
		var max_process sql.NullInt64
		rows.Scan(&datadlcfg.ID, &datadlcfg.AgentUserID, &downloadSetting, &downloadScript, &crontabSetting, &downloadName, &isCronenabled, &description, &node, &max_process)

		b, err := dataimport.CipherDownloadPasswordJSON(downloadSetting, false)
		if err != nil {
			log.Logf("Can not descrypt download password ...%v", err)
		}
		datadlcfg.DownloadSetting = b.JSON()

		datadlcfg.DownloadName = downloadName.String
		datadlcfg.CrontabSetting = crontabSetting.String
		datadlcfg.DownloadScript = downloadScript.String
		datadlcfg.IsCronenabled = isCronenabled.Bool
		datadlcfg.Description = description.String
		datadlcfg.Node = node.String
		datadlcfg.MaxProcess = max_process.Int64
	}

	// return dataimport download config
	return datadlcfg, err
}

// get dataimport download list
//  Parameters:
//		page
//			 param for check call from page migrate
//  Return:
//		All dataimport download
func GetDataImportDownloadList(page string) (*DataDetailsDownloadList, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql get download config
	q := getDownloadConfigList
	p := []interface{}{}

	// process dataimport download config list
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dlcfglist := make([]*DataImportDownloadConfigList, 0)
	// loop for dataimport download list
	for rows.Next() {
		opt := &DataImportDownloadConfigList{}
		// declare value for get
		var (
			dlName        sql.NullString
			dlScript      sql.NullString
			isCronenabled sql.NullBool
			dlHost        sql.NullString
			description   sql.NullString
			dlDataset     sql.NullInt64
			node          sql.NullString
		)
		rows.Scan(&opt.ID, &dlScript, &dlName, &isCronenabled, &dlHost, &description, &dlDataset, &node)
		// check pqtranfrom
		if dlName.Valid {
			opt.DownloadName = dlName.String
		}
		if matchPqDownloadList(page, dlHost.String) {
			opt.DownloadScript = dlScript.String
			opt.DownloadName = dlName.String
			opt.IsCronenabled = isCronenabled.Bool
			opt.Description = description.String
			opt.Node = node.String
			if dlDataset.Valid {
				opt.IsRun = true
			}
			dlcfglist = append(dlcfglist, opt)
		}
	}

	// sql agent dataimport
	q = getAgentDataimport
	p = []interface{}{}

	// process data
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	selectAgent := make([]*selectoption.Option, 0)
	// loop for get agent dataimport
	for rows.Next() {
		sAgent := &selectoption.Option{}
		var account string
		var id int64
		rows.Scan(&id, &account)
		sAgent.Text = account
		sAgent.Value = id
		selectAgent = append(selectAgent, sAgent)
	}

	selectOpt := &SelectOptionDownloadList{}

	// get config from table api.system_setting
	selectOpt.DownloadScript = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.dl.DownloadScript"))

	if page == "migrate" {
		selectOpt.DownloadDriver = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.dl.DownloadDriverMigrate"))
	} else {
		selectOpt.DownloadDriver = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.dl.DownloadDriver"))
	}

	// add agent to select option
	selectOpt.AgentUser = result.Result1(selectAgent)

	//	add rdl nodes to select option
	rdlNodes, _ := GetAllRDLNodeFromSetting()
	selectOpt.RDLNodes = result.Result1(rdlNodes)

	// add download list to json
	dlList := &DataDetailsDownloadList{}
	dlList.DataImportDownloaaList = dlcfglist
	dlList.SelectOpt = selectOpt

	// return download list
	return dlList, nil
}

// check pqtranfrom download list
//  Parameters:
//		page
//			 param for check call from page migrate
//		value
//			url value for check with regex migrate
//  Return:
//		boolean for check migrate download list
func matchPqDownloadList(page string, value string) bool {
	re := regexp.MustCompile(`pq.*`)
	uri, _ := url.Parse(value)
	if page == "migrate" {
		return re.MatchString(uri.Scheme)
	}
	return !re.MatchString(uri.Scheme)
}

//GetAllRDLNodeFromSetting get all rdl node from system setting
func GetAllRDLNodeFromSetting() ([]*selectoption.Option, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := getRDLNodesFromSystemSetting
	rows, err := db.Query(q)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	rdlNodes := make([]*selectoption.Option, 0)
	for rows.Next() {
		var (
			_name sql.NullString
			_id   sql.NullInt64
		)
		rows.Scan(&_name, &_id)
		if _name.Valid {
			rdlNodes = append(rdlNodes, &selectoption.Option{Text: _name.String, Value: _id.Int64})
		}
	}
	return rdlNodes, nil
}

//GetRDLNodeFromSetting get rdl node from system setting by id
func GetRDLNodeFromSetting(nodeId int64) (*DLName, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := getRDLNodesFromSystemSetting + " AND id = $1"
	var (
		_name sql.NullString
		_id   sql.NullInt64
	)
	err = db.QueryRow(q, nodeId).Scan(&_name, &_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	return &DLName{Text: _name.String, Value: _id.Int64}, nil
}

// get dataimport dataset config
//  Parameters:
//		dataimportDatasetID
//			 dataset id for get information
//  Return:
//		json convert setting, import setting, lookup table, import table
func GetDataImportDatasetConfig(dataimportDatasetID string) (*DataImportDataSetConfig1, error) {
	datasetID, err := strconv.Atoi(dataimportDatasetID)

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get dataset config
	q := getDatasetConfig
	p := []interface{}{datasetID}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	datadscfg := &DataImportDataSetConfig1{}

	// declare value
	var (
		convertSetting string
		importSetting  string
		lookupTable    string
		importTable    string
		convertScript  sql.NullString
		importScript   sql.NullString
		convertName    sql.NullString
	)

	// loop data to struct
	for rows.Next() {
		rows.Scan(&datadscfg.DataSetConfigID, &datadscfg.AgentUserID, &datadscfg.DownloadConfigID, &convertSetting, &importSetting, &lookupTable, &importTable, &convertScript, &importScript, &convertName)
		datadscfg.ConvertSetting = json.RawMessage(convertSetting)
		datadscfg.ImportSetting = json.RawMessage(importSetting)
		datadscfg.LookupTable = json.RawMessage(lookupTable)
		datadscfg.ImportTable = json.RawMessage(importTable)
		datadscfg.ConvertScript = convertScript.String
		datadscfg.ImportScript = importScript.String
		datadscfg.ConvertName = convertName.String

	}

	// read json and get desctination
	cfgsImsTableSetting := map[string]interface{}{}
	//	b,_ := json.Marshal(importSetting)
	err = json.Unmarshal([]byte(importSetting), &cfgsImsTableSetting)
	cfgs := cfgsImsTableSetting["configs"].([]interface{})[0].(map[string]interface{})["imports"].([]interface{})
	ims := cfgs[0].(map[string]interface{})
	datadscfg.ImportDestination = ims["destination"].(string)

	type ImportTableConfig struct {
		Table map[string]interface{} `json:"tables"`
	}

	// read json and get partition , unique
	cfgsImsTable := map[string]interface{}{}
	err = json.Unmarshal([]byte(importTable), &cfgsImsTable)
	for _, v := range cfgsImsTable["tables"].(map[string]interface{}) {
		tableName := v.(map[string]interface{})
		if tableName["unique_constraint"] != nil {
			datadscfg.UniqueConstraint = tableName["unique_constraint"].(string)
		}
		if tableName["partition_field"] != nil {
			datadscfg.PartitionField = tableName["partition_field"].(string)
		}
	}
	type ConfigSetting struct {
		DataFolder string        `json:"data_folder"`
		Configs    []interface{} `json:"configs"`
	}

	// read json and get row_validator
	cfgsConfigSetting := &ConfigSetting{}
	err = json.Unmarshal([]byte(convertSetting), &cfgsConfigSetting)

	for _, v := range cfgsConfigSetting.Configs {
		tableName := v.(map[string]interface{})
		if tableName["row_validator"] != nil {
			datadscfg.RowValidator = tableName["row_validator"].(string)
		}
	}

	return datadscfg, err
}

func GetDataImportDatasetList(page, download_name string) (*DataDetailsDatasetList, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	downloadIDMigrate, err := matchPqDatasetList(page)

	// sql get dataset config list
	q := getDatasetConfigList
	p := []interface{}{}
	if len(downloadIDMigrate) > 0 {
		var condition string
		for i, v := range downloadIDMigrate {
			if i != 0 {
				condition += " OR dataimport_download_id=$" + strconv.Itoa(i+1)
			} else {
				condition = "dataimport_download_id=$" + strconv.Itoa(i+1)
			}
			p = append(p, v)
		}
		q += " AND (" + condition + ") "
	}

	if download_name == "migrate" {
		q += " AND dl.download_name like 'migrate-%'"
	} else if download_name == "notmigrate" {
		q += " AND dl.download_name not like 'migrate-%'"
	} else {
		q += ""
	}

	q += getDatasetConfigListOrderBy

	//fmt.Println(a)
	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dscfglist := make([]*DataImportDatasetConfigList, 0)

	// loop for get dataset config list
	for rows.Next() {
		opt := &DataImportDatasetConfigList{}
		var (
			dlName    sql.NullString
			cvName    sql.NullString
			dlID      sql.NullInt64
			dlScript  sql.NullString
			tableName sql.NullString
		)
		rows.Scan(&opt.ID, &dlName, &cvName, &dlID, &dlScript, &tableName)
		opt.DownloadName = dlName.String
		opt.ConvertName = cvName.String
		opt.DownloadScript = dlScript.String
		opt.TableName = tableName.String
		if dlID.Valid {
			opt.DownloadID = dlID.Int64
		}
		dscfglist = append(dscfglist, opt)
	}

	// sql get download config name
	q = getDownloadConfigNameForDataset
	p = []interface{}{}
	if len(downloadIDMigrate) > 0 {
		var condition string
		for i, v := range downloadIDMigrate {
			if i != 0 {
				condition += " OR id=$" + strconv.Itoa(i+1)
			} else {
				condition = "id=$" + strconv.Itoa(i+1)
			}
			p = append(p, v)
		}
		q += " AND (" + condition + ") "
	}
	q += getDownloadConfigNameForDatasetOrderBy

	// process data
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectDownloadList := make([]*DownloadNameOptionAndFilename, 0)
	// loop dataset config list
	for rows.Next() {
		opt := &DownloadNameOptionAndFilename{}
		var (
			dlName         sql.NullString
			resultFile     sql.NullString
			destination    sql.NullString
			source         sql.NullString
			outputfilename sql.NullString
			tablename      sql.NullString
			oldtablename   sql.NullString
			dlScript       sql.NullString
		)
		rows.Scan(&opt.Value, &dlName, &resultFile, &destination, &source, &outputfilename, &tablename, &oldtablename, &dlScript)
		opt.Text = dlName.String

		// find file name from download config
		if resultFile.String != "" {
			opt.FileName = resultFile.String
		} else if destination.Valid {
			if len(destination.String) > 0 {
				opt.FileName = destination.String
			} else if dlScript.String == "dl-basic" {
				opt.FileName = source.String
			}
		} else if dlScript.Valid {
			if dlScript.String == "dl-basic" {
				opt.FileName = source.String
			}
		}
		if outputfilename.Valid {
			opt.FileName = outputfilename.String + ".csv"
		} else if oldtablename.Valid {
			opt.FileName = oldtablename.String + ".csv"
		} else if tablename.Valid {
			opt.FileName = tablename.String + ".csv"
		}

		if len(opt.FileName) <= 0 && dlScript.String == "dl-basic" {
			opt.FileName = "result.json"
		}

		if len(opt.FileName) <= 0 && dlScript.String == "dl-collector" {
			opt.FileName = "filelist.json"
		}

		selectDownloadList = append(selectDownloadList, opt)
	}

	// sql agent dataimport
	q = getAgentDataimport
	p = []interface{}{}

	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	// dropdown json agent list
	selectAgent := make([]*selectoption.Option, 0)
	for rows.Next() {
		sAgent := &selectoption.Option{}
		var account string
		var id int64
		rows.Scan(&id, &account)
		sAgent.Text = account
		sAgent.Value = id
		selectAgent = append(selectAgent, sAgent)
	}

	// sql get table for import
	q = getMasterTable
	p = []interface{}{}

	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	selectOptMaster := make([]*SelectOptionRelated, 0)
	OptRelated := make([]*RelatedOption, 0)
	sopt := &SelectOptionRelated{}
	soptrelated := &RelatedOption{}
	// loop for get table and column
	for rows.Next() {
		var parent string
		var column string
		rows.Scan(&parent, &column)
		// condition get column
		if sopt.Value != "" {
			if sopt.Value != parent {
				// new table
				sopt.Related = OptRelated
				selectOptMaster = append(selectOptMaster, sopt)

				// add column
				OptRelated = make([]*RelatedOption, 0)
				sopt = &SelectOptionRelated{}
				sopt.Text = parent
				sopt.Value = parent
				soptrelated = &RelatedOption{}
				soptrelated.Text = column
				soptrelated.Value = column
				OptRelated = append(OptRelated, soptrelated)
			} else {
				// add column
				soptrelated := &RelatedOption{}
				soptrelated.Text = column
				soptrelated.Value = column
				OptRelated = append(OptRelated, soptrelated)
			}
		} else {
			// first row
			sopt.Text = parent
			sopt.Value = parent
			soptrelated.Text = column
			soptrelated.Value = column
			OptRelated = append(OptRelated, soptrelated)
		}
	}

	if sopt.Value != "" {
		// last table
		sopt.Related = OptRelated
		selectOptMaster = append(selectOptMaster, sopt)
	}

	// sql get transection table
	q = getParentTable
	p = []interface{}{}

	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//	selectOptParent := make([]*selectoption.Option, 0)
	selectOptParent := make([]*SelectOptionRelated, 0)
	OptRelated = make([]*RelatedOption, 0)
	sopt = &SelectOptionRelated{}
	soptrelated = &RelatedOption{}
	for rows.Next() {
		var parent string
		var column string
		rows.Scan(&parent, &column)
		if sopt.Value != "" {
			if sopt.Value != parent {
				// new table
				sopt.Related = OptRelated
				selectOptParent = append(selectOptParent, sopt)

				OptRelated = make([]*RelatedOption, 0)
				sopt = &SelectOptionRelated{}
				sopt.Text = parent
				sopt.Value = parent
				soptrelated = &RelatedOption{}
				soptrelated.Text = "#"
				soptrelated.Value = "#"
				OptRelated = append(OptRelated, soptrelated)
				soptrelated = &RelatedOption{}
				soptrelated.Text = column
				soptrelated.Value = column
				OptRelated = append(OptRelated, soptrelated)
			} else {
				// add column
				soptrelated := &RelatedOption{}
				soptrelated.Text = column
				soptrelated.Value = column
				OptRelated = append(OptRelated, soptrelated)
			}
		} else {
			// first row
			sopt.Text = parent
			sopt.Value = parent
			soptrelated = &RelatedOption{}
			soptrelated.Text = "#"
			soptrelated.Value = "#"
			OptRelated = append(OptRelated, soptrelated)
			soptrelated = &RelatedOption{}
			soptrelated.Text = column
			soptrelated.Value = column
			OptRelated = append(OptRelated, soptrelated)
		}
	}

	if sopt.Value != "" {
		sopt.Related = OptRelated
		selectOptParent = append(selectOptParent, sopt)
	}

	selectOpt := &SelectOptionDatasetList{}
	// get config from api.system_setting
	selectOpt.ConvertScript = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.ci.ConvertScript"))
	selectOpt.ImportScript = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.ci.ImportScript"))
	selectOpt.Type = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.ci.Type"))
	selectOpt.TransformMethod = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.ci.TransformMethod"))
	selectOpt.AddMissing = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.ci.AddMissing"))
	selectOpt.InputFormatDatetime = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.ci.InputFormatDatetime"))

	// add data to json
	selectOpt.ParentTable = result.Result1(selectOptParent)
	selectOpt.MasterTable = result.Result1(selectOptMaster)
	selectOpt.AgentUser = result.Result1(selectAgent)
	selectOpt.DownloadList = result.Result1(selectDownloadList)

	dsList := &DataDetailsDatasetList{}
	dsList.SelectOpt = selectOpt
	dsList.DataImportDatasetList = dscfglist

	// return dataimport dataset list
	return dsList, nil
}

// download pqtranfrom
//  Parameters:
//		page
//			 param for check call from page migrate
//  Return:
//		array download id
func matchPqDatasetList(page string) ([]int64, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql dataimport download config list
	q := getDownloadConfigList
	p := []interface{}{}

	var downloadIDMigrate []int64
	// process sql
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	// loop download
	for rows.Next() {
		var (
			dlID          int64
			dlName        sql.NullString
			dlScript      sql.NullString
			isCronenabled sql.NullBool
			dlHost        sql.NullString
			description   sql.NullString
			dlDataset     sql.NullInt64
			node          sql.NullString
		)
		rows.Scan(&dlID, &dlScript, &dlName, &isCronenabled, &dlHost, &description, &dlDataset, &node)
		if matchPqDownloadList(page, dlHost.String) {
			downloadIDMigrate = append(downloadIDMigrate, dlID)
		}
	}

	// return data
	return downloadIDMigrate, nil
}

// Get metadata list
//  Parameters:
//		page
//			 param for check call from page migrate
//  Return:
//		metadata list with select option
func GetDataImportConfigMetadataList(page string) (*DataDetails, error) {

	// conncet database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	var q string
	if page == "migrate" {
		q = getMetaDataListProvision
	} else {
		q = getMetaDataList
	}

	p := []interface{}{}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	metadataList := make([]*MetadataListName, 0)
	// loop metadata list
	for rows.Next() {
		// declare value for get data from table
		metadata := &MetadataListName{}
		var (
			//			monitorScript sql.NullString
			downloadID        sql.NullInt64
			downloadName      sql.NullString
			datasetID         sql.NullInt64
			datasetName       sql.NullString
			metadataID        int64
			metadataName      sql.NullString
			agencyID          sql.NullInt64
			agencyName        sql.NullString
			additionalDataset sql.NullString
		)
		rows.Scan(&agencyID, &agencyName, &metadataID, &metadataName, &downloadID, &downloadName, &datasetID, &datasetName, &additionalDataset)

		// add metadata details
		metadataDesc := &MetadataDetails{}
		metadataDesc.Value = metadataID
		metadataDesc.Text = metadataName.String

		// add download description
		downloadDesc := &MetadataDetails{}
		downloadDesc.Value = downloadID.Int64
		downloadDesc.Text = downloadName.String

		// add dataset description
		datasetDesc := &MetadataDetails{}
		datasetDesc.Value = datasetID.Int64
		datasetDesc.Text = datasetName.String
		// add monitor script
		//		metadata.MonitorScript = monitorScript.String
		metadata.DownloadName = downloadDesc
		metadata.DatasetName = datasetDesc
		metadata.MetadataName = metadataDesc
		// agency description
		agencyDesc := &MetadataDetails{}
		agencyDesc.Value = agencyID.Int64
		agencyDesc.Text = agencyName.String
		metadata.AgencyName = agencyDesc

		if additionalDataset.Valid {
			metadata.AdditionalDataset = strings.Split(additionalDataset.String, ",")
		}

		metadataList = append(metadataList, metadata)
	}

	// sql dataimport download name list
	q = getDownloadNameList
	p = []interface{}{}
	// process data
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dlList := map[int64]string{}
	allDlList := map[int64]string{}
	// loop for get download list
	for rows.Next() {
		var dlName sql.NullString
		var dlID int64
		var dlHost sql.NullString
		rows.Scan(&dlID, &dlName, &dlHost)
		if matchPqDownloadList(page, dlHost.String) {
			dlList[dlID] = dlName.String
		}
		allDlList[dlID] = dlName.String
	}

	// sql get dataset name list
	q = getDatasetNameList
	p = []interface{}{}

	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectDatasetList := make([]*MetadataDetails, 0)
	optList := map[int64][]*MetadataDetails{}
	var listID []int64
	// loop metadata list for dataset
	for rows.Next() {
		opt := &MetadataDetails{}
		var dlName sql.NullString
		var dlID sql.NullInt64
		rows.Scan(&opt.Value, &dlName, &dlID)
		opt.Text = dlName.String
		selectDatasetList = append(selectDatasetList, opt)

		if dlID.Int64 != 0 {
			dsList := optList[dlID.Int64]
			if dsList != nil {
				dsList = append(dsList, opt)
				optList[dlID.Int64] = dsList
			} else {
				datasetListID := make([]*MetadataDetails, 0)
				datasetListID = append(datasetListID, opt)
				listID = append(listID, dlID.Int64)
				optList[dlID.Int64] = datasetListID
			}
		}
	}

	selectDownloadDatasetList := make([]*DownloadDatasetListID, 0)
	selectAllDownloadDatasetList := make([]*DownloadDatasetListID, 0)
	// loop add data metadata ,download,dataset to json
	for _, v := range listID {
		if dlList[v] != "" {
			opt := &DownloadDatasetListID{}
			opt.DownloadID = v
			opt.DownloadName = dlList[v]
			opt.DatasetListID = optList[v]
			selectDownloadDatasetList = append(selectDownloadDatasetList, opt)
		}
		if allDlList[v] != "" {
			opt := &DownloadDatasetListID{}
			opt.DownloadID = v
			opt.DownloadName = allDlList[v]
			opt.DatasetListID = optList[v]
			selectAllDownloadDatasetList = append(selectAllDownloadDatasetList, opt)
		}
	}

	selectOpt := &SelectOption{}
	selectOpt.DownloadDatasetList = result.Result1(selectDownloadDatasetList)
	selectOpt.AllDownloadDatasetList = result.Result1(selectAllDownloadDatasetList)

	data := &DataDetails{}
	data.SelectOpt = selectOpt
	data.MetaData = metadataList
	// return data details
	return data, nil
}

// get Select option list for history
//  Parameters:
//		page
//			 param for check call from page migrate
//  Return:
//		metadata option list
func GetDataimportHistorySelectOptList(page string) (*HistorySelectOptList, error) {

	// connect database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	var q string
	if page == "migrate" {
		q = getMetadataListProvision
	} else {
		q = getMetadataList
	}

	p := []interface{}{}
	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectMetadataList1 := make([]*selectoption.Option, 0)
	agencyIDmap := map[int64][]*selectoption.Option{}
	// loop agency map
	for rows.Next() {
		opt := &selectoption.Option{}
		var (
			id        sql.NullInt64
			name      sql.NullString
			agency_id sql.NullInt64
		)
		rows.Scan(&id, &name, &agency_id)
		opt.Text = name.String
		opt.Value = id.Int64
		selectMetadataList1 = append(selectMetadataList1, opt)
		// map agency
		if agencyIDmap[agency_id.Int64] != nil {
			ml := agencyIDmap[agency_id.Int64]
			ml = append(ml, opt)
			agencyIDmap[agency_id.Int64] = ml
		} else {
			selectMetadataList := make([]*selectoption.Option, 0)
			selectMetadataList = append(selectMetadataList, opt)
			agencyIDmap[agency_id.Int64] = selectMetadataList
		}
	}

	// sql agency list
	q = getAgencyList
	p = []interface{}{}
	// process data
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectAgencyList := make([]*MetadataNameInAgency, 0)

	// loop metadata to agency
	for rows.Next() {
		opt := &MetadataNameInAgency{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		opt.Text = name.String
		opt.Value = id.Int64
		opt.Metadata = agencyIDmap[id.Int64]
		selectAgencyList = append(selectAgencyList, opt)
	}

	selectOpt := &SelectOptionHistory{}

	// get config from table api.system_setting
	selectOpt.Status = result.ResultJson1(setting.GetSystemSettingJSON("bof.DataIntegration.hs.Status"))
	selectOpt.Agency = result.Result1(selectAgencyList)
	selectOpt.Metadata = result.Result1(selectMetadataList1)
	hl := &HistorySelectOptList{}
	hl.SelectOpt = selectOpt

	dateRange, _ := strconv.Atoi(setting.GetSystemSetting("setting.Default.DateRange"))
	hl.DateRange = int64(dateRange)
	// return data
	return hl, nil
}

// dataimport history list
//  Parameters:
//		agencyID
//			 agency id import data
//		metadataID
//			 metadata id import data
//		processStatus
//			 process status dataimport
//		beginAt
//			 start date dataimport
//		endAt
//			 end date dataimport
//		page
//			 param for check call from page migrate
//  Return:
//		dataimport history
func GetDataimportHistoryList(agencyID, metadataID, processStatus []int64, beginAt, endAt, page string) ([]*HistoryMetadata, error) {

	var q string
	if page == "migrate" {
		q = getHistoryListProvision
	} else {
		q = getHistoryList
	}

	// created condition
	downloadEventCode := ""
	conditionCount := 3
	p := []interface{}{agencyID, metadataID, processStatus, beginAt, endAt}
	if beginAt == "" || endAt == "" {
		current_date := time.Now().AddDate(0, 0, -30).Local().Format("2006-01-02")
		q += " WHERE ddl.download_begin_at >= $1"
		p = []interface{}{current_date}
		conditionCount = 2
	} else {
		q += " WHERE (ddl.download_begin_at BETWEEN $1 AND $2) "
		//			" AND (api.dataimport_dataset_log.import_begin_at BETWEEN $1 AND $2) "
		p = []interface{}{beginAt + ":00", endAt + ":59"}
		conditionCount = 3
	}

	agencyCondition := ""
	for i, v := range agencyID {
		if i > 0 {
			agencyCondition += " OR a.id = $" + strconv.Itoa(conditionCount)
		} else {
			agencyCondition = " a.id = $" + strconv.Itoa(conditionCount)
		}
		p = append(p, v)
		conditionCount++
	}

	metadataCondition := ""
	if page == "migrate" {
		for i, v := range metadataID {
			if i > 0 {
				metadataCondition += " OR md.id = $" + strconv.Itoa(conditionCount)
			} else {
				metadataCondition = " md.id = $" + strconv.Itoa(conditionCount)
			}
			p = append(p, v)
			conditionCount++
		}
	} else {
		for i, v := range metadataID {
			if i > 0 {
				metadataCondition += " OR md.id = $" + strconv.Itoa(conditionCount)
			} else {
				metadataCondition = " md.id = $" + strconv.Itoa(conditionCount)
			}
			p = append(p, v)
			conditionCount++
		}
	}

	processStatusCondition := ""
	var processFirst bool
	for i, v := range processStatus {

		if v == -1 {
			if downloadEventCode == "" {
				downloadEventCode = "download_event_code_id = 2"
			} else {
				downloadEventCode += " OR download_event_code_id = 2"
			}
		} else if v == 0 {
			if downloadEventCode == "" {
				downloadEventCode = "download_event_code_id != 2"
			} else {
				downloadEventCode += " OR download_event_code_id != 2"
			}

			if i > 0 && processFirst {
				processStatusCondition += " OR ddsl.process_status IS NULL "
			} else {
				processFirst = true
				processStatusCondition = " ddsl.process_status IS NULL "
			}
		} else {
			if downloadEventCode == "" {
				downloadEventCode = " download_event_code_id = 1"
			} else {
				downloadEventCode += " OR download_event_code_id = 1"
			}

			if i > 0 && processFirst {
				processStatusCondition += " OR ddsl.process_status = $" + strconv.Itoa(conditionCount)
			} else {
				processFirst = true
				processStatusCondition = " ddsl.process_status = $" + strconv.Itoa(conditionCount)
			}
			p = append(p, v)
			conditionCount++
		}
	}

	if agencyCondition != "" {
		agencyCondition = "(" + agencyCondition + ")"
		q += " AND " + agencyCondition
	}
	if metadataCondition != "" {
		metadataCondition = "(" + metadataCondition + ")"
		q += " AND " + metadataCondition
	}
	if processStatusCondition != "" {
		processStatusCondition = "(" + processStatusCondition + ")"
		q += " AND " + processStatusCondition
	}
	if downloadEventCode != "" {
		downloadEventCode = "(" + downloadEventCode + ")"
		q += " AND " + downloadEventCode
	}

	q += " ORDER BY ddl.download_begin_at"

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dataHistoryList := make([]*HistoryMetadata, 0)
	// loop history list
	for rows.Next() {
		htmd := &HistoryMetadata{}
		var (
			downloadID        sql.NullInt64
			downloadLogID     sql.NullInt64
			datasetID         sql.NullInt64
			datasetLogID      sql.NullInt64
			dlBeginAt         sql.NullString
			dlDuration        sql.NullInt64
			cvBeginAt         sql.NullString
			cvDuration        sql.NullInt64
			imBeginAt         sql.NullString
			imDuration        sql.NullInt64
			processStatus     sql.NullInt64
			filesize          sql.NullInt64
			metadataName      pqx.JSONRaw
			metadataFrequency sql.NullString
			metadataChannel   sql.NullString
			downloadScript    sql.NullString
			agencyName        pqx.JSONRaw
			eventCodeID       sql.NullInt64
			eventCodeIDcv     sql.NullInt64
			eventCodeIDim     sql.NullInt64
			codeDescDL        pqx.JSONRaw
			codeDescCV        pqx.JSONRaw
			codeDescIM        pqx.JSONRaw
		)
		// scan sql migrate
		if page == "migrate" {
			rows.Scan(&downloadID, &downloadLogID, &datasetID, &datasetLogID, &dlBeginAt, &dlDuration, &cvBeginAt,
				&cvDuration, &imBeginAt, &imDuration, &processStatus, &filesize, &metadataName,
				&downloadScript, &agencyName, &eventCodeID,
				&eventCodeIDcv, &eventCodeIDim, &codeDescDL, &codeDescCV, &codeDescIM)
		} else {
			// scan sql
			rows.Scan(&downloadID, &downloadLogID, &datasetID, &datasetLogID, &dlBeginAt, &dlDuration, &cvBeginAt,
				&cvDuration, &imBeginAt, &imDuration, &processStatus, &filesize, &metadataName,
				&metadataFrequency, &metadataChannel, &downloadScript, &agencyName, &eventCodeID,
				&eventCodeIDcv, &eventCodeIDim, &codeDescDL, &codeDescCV, &codeDescIM)
		}

		htmd.DownloadID = downloadID.Int64
		htmd.DownloadLogID = downloadLogID.Int64
		htmd.DatasetID = datasetID.Int64
		htmd.DatasetLogID = datasetLogID.Int64
		htmd.AgencyName = agencyName.JSON()

		// check process status
		if processStatus.Valid {
			if processStatus.Int64 == 3 && imBeginAt.Valid {
				htmd.ReRunFlag = "null"
			} else {
				htmd.ReRunFlag = "cv"
			}

			if processStatus.Int64 == 3 {
				htmd.EventCode = codeDescIM.JSON()
			} else if processStatus.Int64 == 2 && eventCodeIDim.Valid {
				htmd.EventCode = codeDescIM.JSON()
			} else if processStatus.Int64 == 2 && !eventCodeIDim.Valid {
				htmd.EventCode = codeDescCV.JSON()
			} else if processStatus.Int64 == 1 && eventCodeIDcv.Valid {
				htmd.EventCode = codeDescCV.JSON()
			} else {
				htmd.EventCode = codeDescDL.JSON()
			}
		} else {
			if eventCodeID.Int64 == 2 {
				htmd.ReRunFlag = "No New Data"
			} else {
				htmd.ReRunFlag = "dl"
			}
			htmd.EventCode = codeDescDL.JSON()
		}

		t, err := time.Parse(time.RFC3339, dlBeginAt.String)
		if err == nil {
			htmd.DownloadBeginAt = t.Format("2006-01-02 15:04")
		}
		t, err = time.Parse(time.RFC3339, cvBeginAt.String)
		if err == nil {
			htmd.ConvertBeginAt = t.Format("2006-01-02 15:04")
		}

		t, err = time.Parse(time.RFC3339, imBeginAt.String)
		if err == nil {
			htmd.ImportBeginAt = t.Format("2006-01-02 15:04")
		}

		htmd.ProcessStatus = processStatus.Int64
		htmd.MetadataServicName = metadataName.JSON()
		htmd.MetadataConvertFrequency = metadataFrequency.String
		htmd.MetadataChannel = metadataChannel.String
		htmd.Filesize = filesize.Int64
		htmd.DownloadScript = downloadScript.String
		htmd.Duration = (dlDuration.Int64 + cvDuration.Int64 + imDuration.Int64) / 1000000000
		if dlBeginAt.String != "" {
			t, _ := time.Parse(time.RFC3339, dlBeginAt.String)
			t.Add(time.Duration(dlDuration.Int64/1000000000) * time.Second)
			htmd.DownloadEndAt = t.Format("2006-01-02 15:04")
		}

		if cvBeginAt.String != "" {
			t, _ := time.Parse(time.RFC3339, cvBeginAt.String)
			t.Add(time.Duration(cvDuration.Int64/1000000000) * time.Second)
			htmd.ConvertEndAt = t.Format("2006-01-02 15:04")
		}

		if imBeginAt.String != "" {
			t, _ := time.Parse(time.RFC3339, imBeginAt.String)
			t.Add(time.Duration(imDuration.Int64/1000000000) * time.Second)
			htmd.ImportEndAt = t.Format("2006-01-02 15:04")
		}

		// add data from table to struct
		dataHistoryList = append(dataHistoryList, htmd)
	}

	// return history list
	return dataHistoryList, nil
}

func GetDownloadCronList() ([]*DownloadCronSetting, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getDownloadCronList
	p := []interface{}{}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	data := make([]*DownloadCronSetting, 0)
	for rows.Next() {
		var (
			id             int64
			downloadName   sql.NullString
			crontabSetting sql.NullString
			downloadScript sql.NullString
			isCronenabled  bool
			node           sql.NullString
			account        sql.NullString
			description    sql.NullString
		)
		d := &DownloadCronSetting{}
		rows.Scan(&id, &downloadName, &crontabSetting, &downloadScript, &isCronenabled, &node, &account, &description)
		d.DownloadID = id
		d.DownloadName = downloadName.String
		d.DownloadScript = downloadScript.String
		d.CrontabSetting = crontabSetting.String
		d.IsCronenabled = isCronenabled
		d.Description = description.String
		d.Node = node.String
		u := strings.Split(account.String, "-")
		if len(u) >= 2 {
			d.Agency = strings.ToUpper(u[1])
		}
		data = append(data, d)
	}
	return data, nil
}

func GetServerCronList() []*cronjob.JobInfo {
	return cronjob.ListJobs()
}

func GetConfigVariable() ([]*ConfigVariable, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getConfigVariable
	p := []interface{}{}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	data := make([]*ConfigVariable, 0)
	for rows.Next() {
		var (
			id           int64
			category     sql.NullString
			configName   sql.NullString
			variableName sql.NullString
			value        sql.NullString
		)
		d := &ConfigVariable{}
		rows.Scan(&id, &category, &configName, &variableName, &value)
		d.ID = id
		d.Category = category.String
		d.ConfigName = configName.String
		d.VariableName = variableName.String
		d.Value = value.String
		data = append(data, d)
	}
	return data, nil
}

func GetListCatVariable() (interface{}, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getListVariable
	p := []interface{}{}

	// process data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*ListVariable, 0)
	for rows.Next() {
		d := &ListVariable{}
		var (
			id            sql.NullInt64
			name_category sql.NullString
		)
		err = rows.Scan(&id, &name_category)
		if err != nil {
			return 0, errors.Repack(err)
		}

		d.ID = id.Int64
		d.NameCategory = name_category.String
		data = append(data, d)
	}

	return data, err
}
