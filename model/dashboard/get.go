// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dashboard is a model for api.dataimport_download,api.dataimport_dataset,login_session table. This table store dataimport logs information.
package dashboard

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// Get data for dashboard return evant data, alert dataimport, last login, user online, dataservice, server status
//  Parameters:
//		None
//  Return:
//		Information Dashboard
func GetDashboard() (*MedadataAlertOutput, error) {

	// frequencyunit := make(map[int64]int64)
	// frequencyname := make(map[int64]string)
	// data := &MedadataAlertOutput{}
	// frequencyunit, frequencyname, _ = getConvertMinute()
	// // get last online metadata
	// dataOnlineOutput, err := getLastOnline(frequencyunit, frequencyname)

	// if err != nil {
	// 	data.AlertOnline = result.Result0(err.Error())
	// } else {
	// 	data.AlertOnline = result.Result1(dataOnlineOutput)
	// }

	// // get last offline metadata
	// dataOfflineOutput, err := getLastOffline(frequencyunit, frequencyname)
	// if err != nil {
	// 	data.AlertOffline = result.Result0(err.Error())
	// } else {
	// 	data.AlertOffline = result.Result1(dataOfflineOutput)
	// }

	// param := &model_event_log.Struct_EventLog_InputParam{}
	// param.StartDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04")
	// param.EndDate = time.Now().Format("2006-01-02 15:04")
	// // get event log summary
	// eventData, err := model_event_log.GetEventLogSummaryGroupByCategory(param)
	// if err != nil {
	// 	data.EventData = result.Result0(err.Error())
	// } else {
	// 	data.EventData = result.Result1(eventData)
	// }

	// // get last login
	// dataLastLoginOutput, showLimitLastLogin, err := getLastLogin()

	// if err != nil {
	// 	data.LastLogin = result.Result0(err.Error())
	// } else {
	// 	data.LastLogin = result.Result1(showLimitLastLogin)
	// 	data.UserOnline = result.Result1(len(dataLastLoginOutput))
	// }

	// // get dataimport
	// dataimportLast, err := getDataimport()

	// if err != nil {
	// 	data.Dataimport = result.Result0(err.Error())
	// } else {
	// 	data.Dataimport = result.Result1(dataimportLast)
	// }

	// // add data to struct
	// cns, _ := model_order_detail.GetCountOrderDetailNoResult()
	// cwd, _ := model_order_detail.GetCountOrderDetailEnableWithoutDownload()
	// cdoe, _ := model_order_detail.GetCountOrderDetailDownloadOnlyEnable()
	// dataservice := &DataService{}
	// dataservice.NoResult = cns
	// dataservice.EnableWithoutDownload = cwd
	// dataservice.DownloadOnlyEnable = cdoe
	//data.Dataservice = result.Result1(dataservice)

	// data.ServerStatus = srvstatus.GetStatus()
	// getAgency()
	return nil, nil
}

func GetAgency() ([]*DataimportOutputData, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql agency
	q := sqlLastDataImport
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// make array output MetadataAlertStatusOffline
	AgencyOutput := make([]*DataimportOutputData, 0)

	for rows.Next() {
		var (
			metadataservice_name pqx.JSONRaw
			agency_id            int64
			metadata_id          int64
			datafrequency        string
			agency_shortname     pqx.JSONRaw
			agency_name          pqx.JSONRaw
			import_begin         string
		)
		rows.Scan(&metadataservice_name, &agency_id, &metadata_id, &datafrequency, &agency_shortname, &agency_name, &import_begin)

		dl := &DataimportOutputData{}
		// add agency
		dl.MetadataID = metadata_id
		dl.MetadataSName = metadataservice_name.JSON()
		dl.AgencyID = agency_id
		dl.AgencyShortName = agency_shortname.JSON()
		dl.AgencyName = agency_name.JSON()
		dl.DataFrequency = datafrequency
		dl.BeginDate = import_begin

		AgencyOutput = append(AgencyOutput, dl)

	}

	return AgencyOutput, nil
}
