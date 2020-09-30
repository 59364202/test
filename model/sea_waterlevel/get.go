// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package sea_waterlevel is a model for public.sea_waterlevel table. This table store sea waterlevel data.
package sea_waterlevel

import (
	"database/sql"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/thaiwater30/util/validdata"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"time"
)

// get sea waterlevel list
//  Parameters:
//		None
//  Return:
//		Array SeaWaterlevelLatest
func GetSeaWaterlevel() ([]*SeaWaterlevelLatest, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get sea waterlevel by station and dates
	q := sqlSelectSeaWaterlevelLatest
	p := []interface{}{}

	// process sql get sea waterlevel by DISTINCT station and max date
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var (
		id                  sql.NullInt64
		stationName         pqx.JSONRaw
		waterlevel_datetime time.Time
		waterlevel_value    sql.NullFloat64
		sea_waterlevel_lat  sql.NullFloat64
		sea_waterlevel_long sql.NullFloat64
		sea_station_oldcode sql.NullString
		aID                 sql.NullInt64
		aName               pqx.JSONRaw
		aShortname          pqx.JSONRaw
		gID                 sql.NullInt64
		area_code           sql.NullString
		area_name           pqx.JSONRaw
		province_code       sql.NullString
		province_name       pqx.JSONRaw
		amphoe_code         sql.NullString
		amphoe_name         pqx.JSONRaw
		tumbon_code         sql.NullString
		tumbon_name         pqx.JSONRaw
	)
	tableData := make([]*SeaWaterlevelLatest, 0)
	dataReal := make([]*SeaWaterlevelData, 0)
	sID := make([]int64, 0)
	pDatetime := make([]time.Time, 0)
	// data observe
	for rows.Next() {
		// scan data
		rows.Scan(&id, &stationName, &waterlevel_datetime, &waterlevel_value, &sea_waterlevel_lat, &sea_waterlevel_long,
			&sea_station_oldcode, &aID, &aName, &aShortname, &gID, &area_code, &area_name, &province_code, &province_name,
			&amphoe_code, &amphoe_name, &tumbon_code, &tumbon_name)

		dd := &SeaWaterlevelData{}

		// add agency
		a := &Agency{}
		a.ID = validdata.ValidData(aID.Valid, aID.Int64)
		a.Name = aName.JSON()
		a.ShortName = aShortname.JSON()

		dd.Agency = a

		// add geocode
		gc := &Geocode{}
		gc.ID = validdata.ValidData(gID.Valid, gID.Int64)
		gc.Area = validdata.ValidData(area_code.Valid, area_code.String)
		gc.AreaName = area_name.JSON()
		gc.Province = validdata.ValidData(province_code.Valid, province_code.String)
		gc.ProvinceName = province_name.JSON()
		gc.Amphoe = validdata.ValidData(amphoe_code.Valid, amphoe_code.String)
		gc.AmphoeName = amphoe_name.JSON()
		gc.Tumbon = validdata.ValidData(tumbon_code.Valid, tumbon_code.String)
		gc.TumbonName = tumbon_name.JSON()

		dd.Geocode = gc
		// add station
		st := &SeaStation{}
		st.ID = id.Int64
		st.Name = stationName.JSONPtr()
		st.Lat, _ = sea_waterlevel_lat.Value()
		st.Long, _ = sea_waterlevel_long.Value()
		st.OldCode = sea_station_oldcode.String

		dd.Station = st
		// add value
		dd.Datetime = udt.DatetimeFormat(waterlevel_datetime, "datetime")
		dd.Value = validdata.ValidData(waterlevel_value.Valid, waterlevel_value.Float64)
		dataReal = append(dataReal, dd)

		sID = append(sID, id.Int64)
		pDatetime = append(pDatetime, waterlevel_datetime)
	}

	// value for forecast
	dataLatest := &SeaWaterlevelLatest{}
	tName := &SeaTagName{}
	tName.TH = "ตรวจวัดจริง"
	dataLatest.Name = tName
	dataLatest.Data = dataReal
	tableData = append(tableData, dataLatest)

	// data forecast
	q = sqlSelectSeaForecast
	dataForecastStation := make([]*SeaWaterlevelData2, 0)
	for _, v := range sID {
		// define date start and date end
		t := time.Now()
		t = t.Add(time.Duration(-t.Minute()) * time.Minute)
		dts := udt.DatetimeFormat(t, "datetime")
		dte := udt.DatetimeFormat(t.Add(20*time.Hour), "datetime")

		p := []interface{}{v, dts, dte}

		// query
		rows, err := db.Query(q, p...)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}
		// get seawater level forecast
		dd, err := getSealWaterlevelForecast(rows)

		rows.Close()
		if err != nil {
			return nil, err
		}
		dataForecastStation = append(dataForecastStation, dd)
	}

	dataForecast := &SeaWaterlevelLatest{}
	tName = &SeaTagName{}
	tName.TH = "ทำนาย"
	dataForecast.Name = tName
	dataForecast.Data = dataForecastStation
	tableData = append(tableData, dataForecast)
	return tableData, nil
}

// func read row forecast
//  Parameters:
//		row
//			sql.Rows
//  Return:
//		SeaWaterlevelData2
func getSealWaterlevelForecast(rows *sql.Rows) (*SeaWaterlevelData2, error) {
	dataRow := &SeaWaterlevelData2{}
	var maxF float64
	var minF float64
	minF = 999999
	validMaxMin := true
	seaforecast := make([]*SeaForecast, 0)
	for rows.Next() {
		// define value
		var (
			id                           sql.NullInt64
			stationName                  pqx.JSONRaw
			forecast_waterlevel_datetime time.Time
			forecast_waterlevel_value    sql.NullFloat64
			sea_waterlevel_lat           sql.NullFloat64
			sea_waterlevel_long          sql.NullFloat64
			sea_station_oldcode          sql.NullString
			aID                          sql.NullInt64
			aName                        pqx.JSONRaw
			aShortname                   pqx.JSONRaw
			gID                          sql.NullInt64
			area_code                    sql.NullString
			area_name                    pqx.JSONRaw
			province_code                sql.NullString
			province_name                pqx.JSONRaw
			amphoe_code                  sql.NullString
			amphoe_name                  pqx.JSONRaw
			tumbon_code                  sql.NullString
			tumbon_name                  pqx.JSONRaw
		)
		// scan data
		rows.Scan(&id, &stationName, &forecast_waterlevel_datetime, &forecast_waterlevel_value, &sea_waterlevel_lat, &sea_waterlevel_long,
			&sea_station_oldcode, &aID, &aName, &aShortname, &gID, &area_code, &area_name, &province_code, &province_name,
			&amphoe_code, &amphoe_name, &tumbon_code, &tumbon_name)
		dd := &SeaWaterlevelData2{}

		// add agency
		a := &Agency{}
		a.ID = validdata.ValidData(aID.Valid, aID.Int64)
		a.Name = aName.JSON()
		a.ShortName = aShortname.JSON()

		dd.Agency = a

		// add geocode
		gc := &Geocode{}
		gc.ID = validdata.ValidData(gID.Valid, gID.Int64)
		gc.Area = validdata.ValidData(area_code.Valid, area_code.String)
		gc.AreaName = area_name.JSON()
		gc.Province = validdata.ValidData(province_code.Valid, province_code.String)
		gc.ProvinceName = province_name.JSON()
		gc.Amphoe = validdata.ValidData(amphoe_code.Valid, amphoe_code.String)
		gc.AmphoeName = amphoe_name.JSON()
		gc.Tumbon = validdata.ValidData(tumbon_code.Valid, tumbon_code.String)
		gc.TumbonName = tumbon_name.JSON()

		dd.Geocode = gc

		// add station
		st := &SeaStation{}
		st.ID = validdata.ValidData(id.Valid, id.Int64)
		st.Name = stationName.JSONPtr()
		st.Lat, _ = sea_waterlevel_lat.Value()
		st.Long, _ = sea_waterlevel_long.Value()
		st.OldCode = sea_station_oldcode.String

		dd.Station = st
		dforecast := &SeaForecast{}
		// add value
		dforecast.Datetime = udt.DatetimeFormat(forecast_waterlevel_datetime, "datetime")
		dforecast.Value = validdata.ValidData(forecast_waterlevel_value.Valid, forecast_waterlevel_value.Float64)
		dataRow = dd
		seaforecast = append(seaforecast, dforecast)
		// check max min
		if validMaxMin {
			if forecast_waterlevel_value.Valid {
				if maxF < forecast_waterlevel_value.Float64 {
					maxF = forecast_waterlevel_value.Float64
				}
				if minF > forecast_waterlevel_value.Float64 {
					minF = forecast_waterlevel_value.Float64
				}
			} else {
				validMaxMin = false
			}
		}
	}
	// add max min to struct
	if validMaxMin {
		dataRow.MaxForecasat = maxF
		dataRow.MinForecast = minF
	}

	dataRow.Seaforecast = seaforecast
	return dataRow, nil
}

// sea waterlevel observe
//  Parameters:
//		inputData
//			SeaWaterlevelInput
//  Return:
//		Array SeaWaterlevelOutput
func GetSeaWaterlevelReal(inputData *SeaWaterlevelInput) ([]*SeaWaterlevelOutput, error) {
	data := make([]*SeaWaterlevelOutput, 0)
	if inputData.StartDate == "" || inputData.EndDate == "" {
		return nil, rest.NewError(422, "Invalid Input Datetime", nil)
	}
	// loop station
	for _, v := range inputData.StationID {
		dataRow, err := GetSeaWaterlevelByStation(v, sqlSelectSeaWaterlevelRealByStation, inputData.StartDate, inputData.EndDate+" 23:59:59")
		if err != nil {
			return nil, err
		}
		if len(dataRow.Data) > 0 {
			data = append(data, dataRow)
		}
	}
	return data, nil
}

// get sea waterlevel forecast
//  Parameters:
//		SeaWaterlevelInput
//  Return:
//		Array SeaWaterlevelOutput
func GetSeaWaterlevelForecast(inputData *SeaWaterlevelInput) (interface{}, error) {
	data := make([]*SeaWaterlevelOutput, 0)
	if inputData.StartDate == "" || inputData.EndDate == "" {
		return nil, rest.NewError(422, "Invalid Input Datetime", nil)
	}
	// loop condition station
	for _, v := range inputData.StationID {
		dataRow, err := GetSeaWaterlevelByStation(v, sqlSelectSeaWaterlevelForecastByStation, inputData.StartDate, inputData.EndDate+" 23:59:59")
		if err != nil {
			return nil, err
		}
		if len(dataRow.Data) > 0 {
			data = append(data, dataRow)
		}
	}
	return data, nil
}

// Get sea waterlevel station
//  Parameters:
//		stationID
//			sea station id
//		q
//			sql get data
//		startDate
//			start date get data
//		endDate
//			end date get data
//  Return:
//		Array SeaWaterlevelOutput
func GetSeaWaterlevelByStation(stationID int64, q, startDate, endDate string) (*SeaWaterlevelOutput, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql get sea waterlevel by station and date
	p := []interface{}{stationID, startDate, endDate}

	// process sql get sea waterlevel by DISTINCT station and max date
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dataOutput := &SeaWaterlevelOutput{}
	dataRow := make([]*SeaWaterlevelData3, 0)

	// get data by station
	for rows.Next() {
		var (
			stationName         pqx.JSONRaw
			waterlevel_datetime time.Time
			waterlevel_value    sql.NullFloat64
		)
		err := rows.Scan(&stationName, &waterlevel_datetime, &waterlevel_value)
		if err != nil {
			return nil, err
		}
		dataOutput.SeriesName = stationName.JSONPtr()
		dd := &SeaWaterlevelData3{}
		// add data to array
		dd.Datetime = udt.DatetimeFormat(waterlevel_datetime, "datetime")
		dd.Value = validdata.ValidData(waterlevel_value.Valid, waterlevel_value.Float64)
		dataRow = append(dataRow, dd)
	}
	dataOutput.Data = dataRow

	return dataOutput, nil
}
