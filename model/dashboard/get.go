// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dashboard is a model for api.dataimport_download,api.dataimport_dataset,login_session table. This table store dataimport logs information.
package dashboard

import (
	"database/sql"
	"time"

	"haii.or.th/api/server/model/srvstatus"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/thaiwater30/util/result"
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
	frequencyunit := make(map[int64]int64)
	frequencyname := make(map[int64]string)
	data := &MedadataAlertOutput{}
	frequencyunit, frequencyname, _ = getConvertMinute()

	// get ignore number
	dataOnlineOutput, err := getCountIgnore()

	if err != nil {
		data.IgnoreCount = result.Result0(err.Error())
	} else {
		data.IgnoreCount = result.Result1(dataOnlineOutput)
	}

	// get agency
	agencylist, err := getAgency()
	if err != nil {
		data.Agency = result.Result0(err.Error())
	} else {
		data.Agency = result.Result1(agencylist)
	}

	// get dataimport
	dataimportLast, err := getDataimport(frequencyunit, frequencyname)

	if err != nil {
		data.DataimportOutputData = result.Result0(err.Error())
	} else {
		data.DataimportOutputData = result.Result1(dataimportLast)
	}

	// get number user
	countuser, err := getCountUser()

	if err != nil {
		data.CountUser = result.Result0(err.Error())
	} else {
		data.CountUser = result.Result1(countuser)
	}

	// // add data to struct
	cns, _ := model_order_detail.GetCountOrderDetailNoResult()
	dataservice := &DataService{}
	dataservice.NoResult = cns
	data.DataService = result.Result1(dataservice)

	data.ServerStatus = srvstatus.GetStatus()

	return data, nil
}

//  Get dataimport now
//  Parameters:
//		None
//  Return:
//		Add data DataimportOutput
func getDataimport(frequencyunit map[int64]int64, frequencyname map[int64]string) ([]*DataimportOutputData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	DataImportOutput := make([]*DataimportOutputData, 0)

	// sql downloader
	q := sqlLastDataImport
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			metadataservice_name pqx.JSONRaw
			agency_id            int64
			metadata_id          int64
			agency_shortname     pqx.JSONRaw
			agency_name          pqx.JSONRaw
			datetime             time.Time
			update_plan          int64
			datafrequency        string
		)
		rows.Scan(&metadataservice_name, &agency_id, &metadata_id, &agency_shortname, &agency_name, &datetime, &update_plan, &datafrequency)

		dl := &DataimportOutputData{}
		// add dataimport
		dl.MetadataID = metadata_id
		dl.MetadataSName = metadataservice_name.JSON()
		dl.AgencyID = agency_id
		dl.AgencyShortName = agency_shortname.JSON()
		dl.AgencyName = agency_name.JSON()
		dl.DataFrequency = datafrequency
		dl.LastUpdate = udt.DatetimeFormat(datetime, "datetime")
		dl.UpdatePlan = update_plan

		dtCal := datetime.Add(time.Duration(frequencyunit[metadata_id]) * time.Minute)
		dl.OverMinute = (time.Now().Add(time.Duration(-dtCal.UnixNano())).UnixNano() / 1000000000) / 60
		if dl.OverMinute > 0 {
			DataImportOutput = append(DataImportOutput, dl)
		}

	}
	return DataImportOutput, nil
}

//  Get Agency's list
//  Parameters:
//		None
//  Return:
//		Add data list of agency

func getAgency() ([]*Agency, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql agency
	q := sqlAgency
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// make array output Agency's list
	AgencyOutput := make([]*Agency, 0)

	for rows.Next() {
		var (
			agency_id        int64
			agency_shortname pqx.JSONRaw
			agency_name      pqx.JSONRaw
		)
		rows.Scan(&agency_id, &agency_shortname, &agency_name)

		dl := &Agency{}
		// add agency
		dl.ID = agency_id
		dl.AgencyShortName = agency_shortname.JSON()
		dl.AgencyName = agency_name.JSON()

		AgencyOutput = append(AgencyOutput, dl)
	}

	return AgencyOutput, nil
}

// get convert Minute
//  Parameters:
//		None
//  Return:
//		map data metadata with frequencyunit and frequencyname
func getConvertMinute() (map[int64]int64, map[int64]string, error) {
	frequencyunit := make(map[int64]int64)
	frequencyname := make(map[int64]string)
	db, err := pqx.Open()
	if err != nil {
		return nil, nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// select convert minute from database
	q := sqlSelectConvertMinute
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id                int64
			frequencyV        sql.NullInt64
			frequencyunitName sql.NullString
		)
		rows.Scan(&id, &frequencyV, &frequencyunitName)
		// add frequency to map[id]
		frequencyunit[id] = frequencyV.Int64
		frequencyname[id] = frequencyunitName.String
	}

	return frequencyunit, frequencyname, nil
}

//  Get Number of ignore
//  Parameters:
//		None
//  Return:
//		number rain and water

func getCountIgnore() (*DataIgnore, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql agency
	q := sqlIgnore
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dl := &DataIgnore{}
	if rows.Next() {
		var (
			rain_count  int64
			water_count int64
		)
		rows.Scan(&rain_count, &water_count)

		// add count
		dl.Rain = rain_count
		dl.Water = water_count
	}
	return dl, nil
}

//  Get Number of user
//  Parameters:
//		None
//  Return:
//		number user

func getCountUser() (*CountUser, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql agency
	q := sqlUser
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	dl := &CountUser{}
	if rows.Next() {
		var (
			user_count int64
		)
		rows.Scan(&user_count)

		// add count
		dl.CountUser = user_count
	}
	return dl, nil
}

// get dataimport metadata online type
//  Parameters:
//		frequencyunit
//					frequencyunit with metadata id
//		frequencyname
//					frequencyname with metadata id
//  Return:
//		Array information metadata status online
// func getLastOnline(frequencyunit map[int64]int64, frequencyname map[int64]string) ([]*MetadataAlertStatus, error) {

// 	db, err := pqx.Open()
// 	if err != nil {
// 		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
// 	}

// 	// sql last online metadata
// 	q := sqlOnline
// 	p := []interface{}{}

// 	// query
// 	rows, err := db.Query(q, p...)
// 	if err != nil {
// 		return nil, pqx.GetRESTError(err)
// 	}
// 	defer rows.Close()

// 	dataOnlineOutput := make([]*MetadataAlertStatus, 0)
// 	for rows.Next() {
// 		var (
// 			id           int64
// 			metadataName pqx.JSONRaw
// 			datetime     time.Time
// 			frequencyV   sql.NullInt64
// 			aid          sql.NullInt64
// 			aName        pqx.JSONRaw
// 			aShortName   pqx.JSONRaw
// 		)
// 		rows.Scan(&id, &metadataName, &aid, &aName, &aShortName, &datetime, &frequencyV)

// 		dd := &MetadataAlertStatus{}
// 		a := &Agency{}
// 		a.ID = validdata.ValidData(aid.Valid, aid.Int64)
// 		// add agency
// 		a.AgencyShortName = aShortName.JSON()
// 		a.AgencyName = aName.JSON()
// 		dd.Agency = a

// 		// add metadata
// 		dd.MetadataID = id
// 		dd.MetadataServiceName = metadataName.JSON()
// 		// last update
// 		dd.LastUpdate = udt.DatetimeFormat(datetime, "datetime")
// 		if frequencyunit[id] != 0 {
// 			fd := &FrequencyUnit2{}
// 			fd.FrequencyName = frequencyname[id]
// 			dd.FrequencyUnitData = fd
// 			dtCal := datetime.Add(time.Duration(frequencyV.Int64*frequencyunit[id]) * time.Minute)
// 			//			dd.OverMinute = int64(time.Now().Minute() - dtCal.Minute())
// 			dd.OverMinute = (time.Now().Add(time.Duration(-dtCal.UnixNano())).UnixNano() / 1000000000) / 60
// 			if dd.OverMinute > 0 {
// 				dataOnlineOutput = append(dataOnlineOutput, dd)
// 			}
// 		}
// 	}

// 	return dataOnlineOutput, nil
// }
