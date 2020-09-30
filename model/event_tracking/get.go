// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_tracking is a model for api.event_log table. This table store eventlog information.
package event_tracking

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/zip"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// get event tracking select option by page
//  Parameters:
//		page
//			event tracking page
//  Return:
//		EventTrackingSelectOptionList
func GetEventTrackingSelectOption(page string) (*EventTrackingSelectOptionList, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	var q string

	if page == "solve_data" {
		q = getEventCode2
	} else {
		q = getEventCode
	}

	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectEventCodeList := make([]*EventTrackingSelectOptionCode, 0)
	eventCodeMap := make(map[int64][]*EventTrackingSelectOptionCode)
	for rows.Next() {
		opt := &EventTrackingSelectOptionCode{}
		var (
			id          sql.NullInt64
			name        sql.NullString
			category_id sql.NullInt64
		)
		rows.Scan(&id, &name, &category_id)
		opt.Text = json.RawMessage(name.String)
		opt.Value = id.Int64
		opt.Category = category_id.Int64
		if eventCodeMap[category_id.Int64] != nil {
			evetTrackOpt := eventCodeMap[category_id.Int64]
			evetTrackOpt = append(evetTrackOpt, opt)
			eventCodeMap[category_id.Int64] = evetTrackOpt
		} else {
			evetTrackOpt := make([]*EventTrackingSelectOptionCode, 0)
			evetTrackOpt = append(evetTrackOpt, opt)
			eventCodeMap[category_id.Int64] = evetTrackOpt
		}
		selectEventCodeList = append(selectEventCodeList, opt)
	}
	// get event log category
	q = getEventLogCategory
	p = []interface{}{}

	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectEventLogCategoryList := make([]*EventTrackingSelectOptionCategory, 0)

	for rows.Next() {
		opt := &EventTrackingSelectOptionCategory{}
		var (
			id   sql.NullInt64
			name sql.NullString
			code sql.NullString
		)
		rows.Scan(&id, &name, &code)
		//		opt.Text = json.RawMessage(name.String)
		if eventCodeMap[id.Int64] != nil {
			opt.Text = code.String
			opt.Value = id.Int64
			opt.Code = eventCodeMap[id.Int64]
			selectEventLogCategoryList = append(selectEventLogCategoryList, opt)
		}

	}

	// get agency list
	q = getAgencyList
	p = []interface{}{}

	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectAgencyList := make([]*EventTrackingSelectOptionAgency, 0)

	for rows.Next() {
		opt := &EventTrackingSelectOptionAgency{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		rows.Scan(&id, &name)
		opt.Text = json.RawMessage(name.String)
		opt.Value = id.Int64
		selectAgencyList = append(selectAgencyList, opt)
	}

	data := &EventTrackingSelectOptionList{}
	data.Category = selectEventLogCategoryList
	data.Code = selectEventCodeList
	data.Agency = selectAgencyList

	return data, nil
}

// get event tracking select option page download invalid data
//  Parameters:
//		None
//  Return:
//		Array EventTrackingSelectOptionInvalidData
func GetEventTrackingDownloadInvalidDataSelectOption() ([]*EventTrackingSelectOptionInvalidData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getEventInvalidDataSelectOption
	p := []interface{}{}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	selectInvalidDate := make([]*EventTrackingSelectOptionInvalidData, 0)

	var (
		id   sql.NullInt64
		name sql.NullString
		date sql.NullString
	)

	for rows.Next() {

		err = rows.Scan(&id, &name, &date)
		if err != nil {
			return nil, err
		}

		opt := &EventTrackingSelectOptionInvalidData{}
		opt.Value = id.Int64
		if !name.Valid || name.String == "" {
			name.String = "{}"
		}
		opt.Text = json.RawMessage(name.String)
		if date.Valid {
			opt.Date = strings.Split(date.String[1:len(date.String)-1], ",")
		}

		selectInvalidDate = append(selectInvalidDate, opt)
	}

	return selectInvalidDate, nil
}

// get event tracking by date range and agency and event cateory
//  Parameters:
//		startDate
//			start date
//		endDate
//			end date
//		agency
//			array agency id
//		eventCategory
//			array category id
//		eventChannel
//			array channel id
//		solveEvent
//			solve flag
//  Return:
//		Array EventTrackingData
func GetEventTrackingData(startDate, endDate string, agency, eventCategory, eventChannel []int64, solveEvent bool) ([]*EventTrackingData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getEventTracking
	p := []interface{}{startDate, endDate + ":59"}

	// make condition for get data
	if solveEvent {
		q += " AND el.solve_event_at IS NULL"
	}
	count := 3

	var agencyCondition string
	for i, v := range agency {
		if i > 0 {
			agencyCondition += " OR a.id=$" + strconv.Itoa(count)
		} else {
			agencyCondition += " a.id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	if agencyCondition != "" {
		q += " AND (" + agencyCondition + ")"
	}

	var categoryCondition string
	for i, v := range eventCategory {
		if i > 0 {
			categoryCondition += " OR elc.id=$" + strconv.Itoa(count)
		} else {
			categoryCondition += " elc.id=$" + strconv.Itoa(count)
		}

		p = append(p, v)
		count++
	}

	if categoryCondition != "" {
		q += " AND (" + categoryCondition + ")"
	}

	var channelCondition string
	for i, v := range eventChannel {
		if i > 0 {
			channelCondition += " OR ec.id=$" + strconv.Itoa(count)
		} else {
			channelCondition += " ec.id=$" + strconv.Itoa(count)
		}

		p = append(p, v)
		count++
	}

	if channelCondition != "" {
		q += " AND (" + channelCondition + ")"
	}

	//	q += " GROUP BY el.id,ec.description::jsonb,el.created_at,a.agency_name::jsonb,elc.description::jsonb,el.event_message,lsq.sent_at,el.solve_event_at,data_path,ec.code,md.metadataservice_name::jsonb ORDER BY id "
	q += " ORDER BY el.id"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*EventTrackingData, 0)
	for rows.Next() {
		eventData := &EventTrackingData{}
		var (
			eventLogID          int64
			eventCode           pqx.JSONRaw
			eventDate           time.Time
			agencyName          pqx.JSONRaw
			eventCategory       pqx.JSONRaw
			eventMessage        sql.NullString
			statusNotify        sql.NullString
			solveEventDate      sql.NullString
			dataPath            sql.NullString
			code                string
			metadataServiceName sql.NullString
		)
		rows.Scan(&eventLogID, &eventCode, &eventDate, &agencyName, &eventCategory, &eventMessage, &statusNotify, &solveEventDate, &dataPath, &code, &metadataServiceName)
		eventData.EventLogId = eventLogID
		eventData.EventCodeSubType = eventCode.JSON()
		eventData.EventCode = code
		eventData.EventDate = eventDate.Format("2006-01-02 15:04")
		eventData.Agency = agencyName.JSON()
		eventData.EventCategory = eventCategory.JSON()
		eventData.EventMessage = eventMessage.String
		//		eventData.MetadataServiceName = metadataServiceName.String
		if solveEventDate.Valid {
			solEventDate, _ := time.Parse(time.RFC3339, solveEventDate.String)
			eventData.SolveEventDate = solEventDate.Format("2006-01-02 15:04")
		} else {
			eventData.SolveEventDate = solveEventDate.String
		}

		if statusNotify.Valid {
			eventData.StatusNotify = "Yes"
		} else {
			eventData.StatusNotify = "No"
		}
		if dataPath.Valid {
			if len(dataPath.String) > 0 {
				eventData.Filepath = true
			} else {
				eventData.Filepath = false
			}
		} else {
			eventData.Filepath = false
		}
		mdsn := strings.Split(metadataServiceName.String, "|")
		for _, v := range mdsn {
			b, err := json.Marshal(v)
			if err != nil {
				eventData.MetadataServiceName = append(eventData.MetadataServiceName, json.RawMessage("{}"))
			} else {
				eventData.MetadataServiceName = append(eventData.MetadataServiceName, b)
			}
		}

		data = append(data, eventData)
	}

	return data, nil
}

// gent event tracking page solve data
//  Parameters:
//		startDate
//			start date
//		endDate
//			end date
//		agency
//			array agency id
//		eventCategory
//			array category id
//		eventChannel
//			array channel id
//  Return:
//		Array EventTrackingData
func GetEventTrackingSolveData(startDate, endDate string, agency, eventCategory, eventChannel []int64) ([]*EventTrackingData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getEventTrackingUpdate
	p := []interface{}{startDate, endDate + ":59"}
	count := 3

	// make condition get data
	var agencyCondition string
	for i, v := range agency {
		if i > 0 {
			agencyCondition += " OR a.id=$" + strconv.Itoa(count)
		} else {
			agencyCondition += " a.id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	if agencyCondition != "" {
		q += " AND (" + agencyCondition + ")"
	}

	var categoryCondition string
	for i, v := range eventCategory {
		if i > 0 {
			categoryCondition += " OR elc.id=$" + strconv.Itoa(count)
		} else {
			categoryCondition += " elc.id=$" + strconv.Itoa(count)
		}

		p = append(p, v)
		count++
	}

	if categoryCondition != "" {
		q += " AND (" + categoryCondition + ")"
	}

	var channelCondition string
	for i, v := range eventChannel {
		if i > 0 {
			channelCondition += " OR ec.id=$" + strconv.Itoa(count)
		} else {
			channelCondition += " ec.id=$" + strconv.Itoa(count)
		}

		p = append(p, v)
		count++
	}

	if channelCondition != "" {
		q += " AND (" + channelCondition + ")"
	}

	//	q += " GROUP BY el.id,ec.description::jsonb,el.created_at,a.agency_name::jsonb,el.event_message,lsq.sent_at,el.solve_event_at,elc.description::jsonb,ec.code,md.metadataservice_name::jsonb ORDER BY id "
	q += " ORDER BY el.id"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*EventTrackingData, 0)
	for rows.Next() {
		eventData := &EventTrackingData{}
		var (
			eventLogID          int64
			eventCode           string
			eventDate           time.Time
			agencyName          sql.NullString
			eventMessage        sql.NullString
			statusNotify        sql.NullString
			solveEventDate      sql.NullString
			eventCategory       string
			code                string
			metadataServiceName sql.NullString
		)
		rows.Scan(&eventLogID, &eventCode, &eventDate, &agencyName, &eventMessage, &statusNotify, &solveEventDate, &eventCategory, &code, &metadataServiceName)
		eventData.EventLogId = eventLogID
		eventData.EventCodeSubType = json.RawMessage(eventCode)
		eventData.EventCode = code
		eventData.EventDate = eventDate.Format("2006-01-02 15:04")
		eventData.Agency = json.RawMessage(agencyName.String)
		eventData.EventMessage = eventMessage.String
		eventData.SolveEventDate = solveEventDate.String
		//		eventData.MetadataServiceName = metadataServiceName.String
		if solveEventDate.Valid {
			solEventDate, _ := time.Parse(time.RFC3339, solveEventDate.String)
			eventData.SolveEventDate = solEventDate.Format("2006-01-02 15:04")
		} else {
			eventData.SolveEventDate = solveEventDate.String
		}
		eventData.EventCategory = json.RawMessage(eventCategory)
		if statusNotify.Valid {
			eventData.StatusNotify = "Yes"
		} else {
			eventData.StatusNotify = "No"
		}
		mdsn := strings.Split(metadataServiceName.String, "|")
		for _, v := range mdsn {
			b, err := json.Marshal(v)
			if err != nil {
				eventData.MetadataServiceName = append(eventData.MetadataServiceName, json.RawMessage("{}"))
			} else {
				eventData.MetadataServiceName = append(eventData.MetadataServiceName, b)
			}
		}
		data = append(data, eventData)
	}

	return data, nil
}

// get event tracking page download invalid data
//  Parameters:
//		date
//			date get event invalid data
//		agency
//			array agency id
//  Return:
//		Array EventInvalidData
func GetDownloadEventInvalidData(agency []int64, date string) ([]*EventInvalidData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getInvalidData
	p := []interface{}{date + " 00:00:00", date + " 23:59:59"}

	count := 3

	var agencyCondition string
	for i, v := range agency {
		if i > 0 {
			agencyCondition += " OR a.id=$" + strconv.Itoa(count)
		} else {
			agencyCondition += " a.id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	if agencyCondition != "" {
		q += " AND (" + agencyCondition + ")"
	}

	q += " GROUP BY el.id,ec.code,el.created_at,a.agency_name->>'th',dd.convert_name,md.metadata_method,result_file, " +
		"destination,source,output_filename,table_name,ddl.download_script,md.metadataservice_name->>'th' ORDER BY el.id"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*EventInvalidData, 0)

	for rows.Next() {
		eventData := &EventInvalidData{}
		var (
			eventLogID          int64
			eventLogCode        string
			eventLogDate        time.Time
			agencyName          sql.NullString
			scriptName          sql.NullString
			eventLogData        sql.NullString
			dataPath            sql.NullString
			metadataMethod      sql.NullString
			resultFile          sql.NullString
			destination         sql.NullString
			source              sql.NullString
			outputfilename      sql.NullString
			tablename           sql.NullString
			dlScript            sql.NullString
			metadataserviceName pqx.JSONRaw
		)
		rows.Scan(&eventLogID, &eventLogCode, &eventLogDate, &agencyName, &scriptName, &eventLogData, &dataPath, &metadataMethod, &resultFile, &destination, &source, &outputfilename, &tablename, &dlScript, &metadataserviceName)
		eventData.EventLogId = eventLogID
		eventData.EventCode = eventLogCode
		eventData.EventDate = eventLogDate.Format("2006-01-02 15:04")
		eventData.Agency = agencyName.String
		eventData.ScriptName = scriptName.String
		eventData.MetadataMethod = metadataMethod.String
		eventData.MetadataServiceName = metadataserviceName.JSON()
		if dataPath.Valid {
			if len(dataPath.String) > 0 {
				eventData.Filepath = true
			} else {
				eventData.Filepath = false
			}
		} else {
			eventData.Filepath = false
		}

		if resultFile.Valid {
			eventData.FileName = resultFile.String
		} else if destination.Valid {
			if len(destination.String) > 0 {
				eventData.FileName = destination.String
			} else if dlScript.String == "dl-basic" {
				eventData.FileName = source.String
			}
		} else if dlScript.Valid {
			if dlScript.String == "dl-basic" {
				eventData.FileName = source.String
			}
		}
		if outputfilename.Valid {
			eventData.FileName = outputfilename.String + ".csv"
		} else if tablename.Valid {
			eventData.FileName = tablename.String + ".csv"
		}

		if len(eventData.FileName) <= 0 {
			eventData.FileName = "filelist.json"
		}

		data = append(data, eventData)
	}

	return data, nil
}

// get event tracking page send invalid data
//  Parameters:
//		dateStart
//			start date for get data
//		dateEnd
//			end date for get data
//		agency
//			array agency id
//  Return:
//		Array EventSendInvalidData
func GetSendEventInvalidData(agency []int64, dateStart, dateEnd string) ([]*EventSendInvalidData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getSendInvalidData
	p := []interface{}{dateStart, dateEnd + ":59"}

	count := 3

	var agencyCondition string
	for i, v := range agency {
		if i > 0 {
			agencyCondition += " OR a.id=$" + strconv.Itoa(count)
		} else {
			agencyCondition += " a.id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	if agencyCondition != "" {
		q += " AND (" + agencyCondition + ")"
	}

	//	q += "GROUP BY el.id,lec.code,el.created_at,dd.convert_name,a.agency_name->>'th',el.event_message,md.metadataservice_name->>'th' ORDER BY el.created_at"
	q += " ORDER BY el.created_at"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := make([]*EventSendInvalidData, 0)

	for rows.Next() {
		eventData := &EventSendInvalidData{}
		var (
			eventLogID          int64
			eventLogCode        string
			eventLogDate        time.Time
			agencyName          sql.NullString
			scriptName          sql.NullString
			eventMsg            sql.NullString
			metadataServiceName sql.NullString
		)
		rows.Scan(&eventLogID, &eventLogCode, &eventLogDate, &scriptName, &agencyName, &eventMsg, &metadataServiceName)
		eventData.EventLogID = eventLogID
		eventData.EventCode = eventLogCode
		eventData.EventDate = eventLogDate.Format("2006-01-02 15:04")
		eventData.ScriptName = scriptName.String
		eventData.Agency = agencyName.String
		eventData.EventMessage = eventMsg.String

		mdsn := strings.Split(metadataServiceName.String, "|")
		for _, v := range mdsn {
			b, err := json.Marshal(v)
			if err != nil {
				eventData.MetadataServiceName = append(eventData.MetadataServiceName, json.RawMessage("{}"))
			} else {
				eventData.MetadataServiceName = append(eventData.MetadataServiceName, b)
			}
		}

		data = append(data, eventData)
	}

	return data, nil
}

// get event tarcking page tracking invalid data
//  Parameters:
//		dateStart
//			start date for get data
//		dateEnd
//			end date for get data
//		agency
//			array agency id
//  Return:
//		Array EventTrackingInvalidData
func GetTrackingEventInvalidData(agency []int64, dateStart, dateEnd string) ([]*EventTrackingInvalidData, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getTrackingInvalidData
	p := []interface{}{dateStart, dateEnd + ":00"}

	count := 3

	var agencyCondition string
	for i, v := range agency {
		if i > 0 {
			agencyCondition += " OR a.id=$" + strconv.Itoa(count)
		} else {
			agencyCondition += " a.id=$" + strconv.Itoa(count)
		}
		p = append(p, v)
		count++
	}

	if agencyCondition != "" {
		q += " AND (" + agencyCondition + ")"
	}

	//	q += " GROUP BY el.id,lec.code,el.created_at,a.agency_name->>'th',el.event_message,el.send_error_at,el.solve_event_at,el.solve_event,md.metadataservice_name->>'th' ORDER BY el.created_at"
	q += " ORDER BY el.created_at"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := make([]*EventTrackingInvalidData, 0)

	for rows.Next() {
		eventData := &EventTrackingInvalidData{}
		var (
			eventLogID          int64
			eventLogCode        sql.NullString
			eventLogDate        time.Time
			agencyName          sql.NullString
			eventMsg            sql.NullString
			sendErrAt           sql.NullString
			solveAt             sql.NullString
			solveMsg            sql.NullString
			metadataServiceName pqx.JSONRaw
		)
		rows.Scan(&eventLogID, &eventLogCode, &eventLogDate, &agencyName, &eventMsg, &sendErrAt, &solveAt, &solveMsg, &metadataServiceName)
		eventData.EventLogID = eventLogID
		eventData.EventCode = eventLogCode.String
		eventData.EventDate = eventLogDate.Format("2006-01-02 15:04")
		eventData.Agency = agencyName.String
		eventData.EventMessage = eventMsg.String
		eventData.MetadataServiceName = metadataServiceName.JSON()
		if sendErrAt.Valid {
			sendErrorDate, _ := time.Parse(time.RFC3339, sendErrAt.String)
			eventData.SendErrorAt = sendErrorDate.Format("2006-01-02 15:04")
		} else {
			eventData.SendErrorAt = sendErrAt.String
		}
		if solveAt.Valid {
			solEventDate, _ := time.Parse(time.RFC3339, solveAt.String)
			eventData.SolveEventDate = solEventDate.Format("2006-01-02 15:04")
		} else {
			eventData.SolveEventDate = solveAt.String
		}
		eventData.SolveEventMsg = solveMsg.String
		data = append(data, eventData)
	}

	return data, nil
}

// get file invalid data
//  Parameters:
//		eventLogID
//			eventlog id for get file
//  Return:
//		path file
func GetFileInvalidData(eventLogID int64) (string, error) {

	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	dataPathPrefix := setting.GetSystemSetting("server.service.dataimport.DataPathPrefix")

	q := getInvalidDataPath
	p := []interface{}{eventLogID}
	rows, err := db.Query(q, p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	defer rows.Close()
	var dataPath sql.NullString
	for rows.Next() {
		rows.Scan(&dataPath)
	}

	src := filepathx.JoinPath(dataPathPrefix, dataPath.String)

	zName := strings.Split(dataPath.String, "/")
	zName = append(zName[:0], zName[1:]...)
	zipName := strings.Join(zName, "_") + ".zip"
	zipName = filepathx.JoinPath(src, zipName)
	writer, err := zip.NewArchive(zipName)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	err = writer.AddPath(src)
	if err != nil {
		return "", err
	}
	return zipName, nil
}

// get file invalid data
//  Parameters:
//		uid
//			user id
//		eventLogID
//			eventlog id for get file
//  Return:
//		path file
func GetFileInvalidData2(uid, eventLogID int64) (string, error) {

	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	dataPathPrefix := setting.GetSystemSetting("server.service.dataimport.DataPathPrefix")

	q := getInvalidDataPath
	p := []interface{}{eventLogID}
	rows, err := db.Query(q, p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	defer rows.Close()
	var dataPath sql.NullString
	for rows.Next() {
		rows.Scan(&dataPath)
	}

	src := filepathx.JoinPath(dataPathPrefix, dataPath.String)

	var arrayFile []string
	r, _ := regexp.Compile(".*.zip")
	files, _ := ioutil.ReadDir(src)
	for _, f := range files {
		if !r.MatchString(f.Name()) {
			arrayFile = append(arrayFile, filepathx.JoinPath(src, f.Name()))
		}
	}

	zName := strings.Split(dataPath.String, "/")
	zName = append(zName[:0], zName[1:]...)
	zipName := strings.Join(zName, "_") + ".zip"
	zipName = filepathx.JoinPath(src, zipName)
	writer, err := zip.NewArchive(zipName)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	for _, v := range arrayFile {
		writer.AddPath(v)
	}
	err = IsDownloaded(uid, eventLogID)
	if err != nil {
		return "", errors.New("Can not stamp is downloaded")
	}
	return zipName, nil
}

// check is download event tracking
//  Parameters:
//		uid
//			user id
//		eventLogID
//			eventlog id
//  Return:
//		None
func IsDownloaded(uid, event_log_id int64) error {

	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := updateEventTrackingIsDownloaded
	p := []interface{}{uid, event_log_id}

	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}
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
