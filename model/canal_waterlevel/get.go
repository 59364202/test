// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package canal_waterlevel is a model for public.canal_waterlevel table. This table store canal_waterlevel.
package canal_waterlevel

import (
	"database/sql"
	model_setting "haii.or.th/api/server/model/setting"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	//	"haii.or.th/api/thaiwater30/util/highchart"
	"haii.or.th/api/thaiwater30/util/validdata"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"strconv"
	"time"
)

// Get Canal Waterlevel Graph By station and detetime
//  Parameters:
//		param 	
//				Param_CanalWaterlevel
//  Return:
//		Data canal waterlevel for graph condition station and datetime
func GetCanalWaterlevelGraphByStationAndDateAnalyst(param *Param_CanalWaterlevel) (*model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystOutput, error) {

	if param.Station_id == "" {
		return nil, rest.NewError(422, "No station id", nil)
	}
	
	// define output
	data := &model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystOutput{}
	// call func get data
	graphData, err := getCanalWaterlevelGraphAnalyst(param.Station_id, param.Start_date, param.End_date)
	if err != nil {
		return nil, err
	}
	// add data to struct
	data.GraphData = graphData

	return data, nil
}

// Get Canal Waterlevel Graph By station and year
//  Parameters:
//		param 	
//				GetCanalWaterlevelYearlyGraphInput
//  Return:
//		Data canal waterlevel for graph condition station and year
func GetWaterlevelYearlyGraphAnalyst(param *GetCanalWaterlevelYearlyGraphInput) (*model_tele_waterlevel.GetWaterlevelYearlyGraphAnalystOutput, error) {

	if param.StationID == "" {
		return nil, rest.NewError(422, "No station id", nil)
	}

	// define output
	data := &model_tele_waterlevel.GetWaterlevelYearlyGraphAnalystOutput{}
	// make array for output
	aGraphData := make([]*model_tele_waterlevel.GetWaterlevelYearlyGraphAnalystOutputYear, 0)
	for _, v := range param.Year {
		gData := &model_tele_waterlevel.GetWaterlevelYearlyGraphAnalystOutputYear{}
		gData.Year = v
		// get data by year
		graphData, err := getCanalWaterlevelGraphAnalyst(param.StationID, strconv.Itoa(v)+"-01-01 00:00:00", strconv.Itoa(v)+"-12-31 23:59:59")
		if err != nil {
			return nil, err
		}
		// add data to array
		gData.GraphData = graphData
		aGraphData = append(aGraphData, gData)
	}
	data.GraphData = aGraphData

	return data, nil
}

// Get canal waterlevel graph by station and detetime
//  Parameters:
//		stationID 	
//				canal station id 
//		startDate 	
//				start date 
//		endDate 	
//				end date
//  Return:
//		Data canal waterlevel for graph condition station and year
func getCanalWaterlevelGraphAnalyst(stationID, startDate, endDate string) ([]*model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystDataOutput, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get waterlevel by station and date
	q := sqlSelectCanal
	p := []interface{}{stationID, startDate, endDate}

	// process sql get waterlevel by station and date
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	var (
		id             sql.NullInt64
		canal_datetime time.Time
		canal_value    sql.NullFloat64
	)

	// make array for output 
	graphData := make([]*model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystDataOutput, 0)
	for rows.Next() {
		err := rows.Scan(&id, &canal_datetime, &canal_value)
		if err != nil {
			return nil, err
		}
		dataRow := &model_tele_waterlevel.GetWaterlevelGraphByStationAndDateAnalystDataOutput{}
		//dataRow.Datetime = canal_datetime.Format(time.RFC3339)
		dataRow.Datetime = canal_datetime.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		dataRow.Value = validdata.ValidData(canal_value.Valid, canal_value.Float64)
		// add data to array
		graphData = append(graphData, dataRow)
	}

	return graphData, nil
}
