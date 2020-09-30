// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package forecast is a model for public.floodforecast table. This table store forecast data.
package forecast

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"
	waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

// 	get cpy forecast latest
//	Return:
//		FloodforecastOutputWithScale
func CpyFloodForecastLatest() (*FloodforecastOutputWithScale, error) {
	cpyStation := []int64{}
	var err error
	oData := &FloodforecastOutputWithScale{}
	oData.Scale = setting.GetSystemSettingJson("Frontend.public.waterlevel_setting")
	oData.FloodforecastLoad, err = floodForecastLatest(cpyStation)
	oData.WaterlevelObserve, err = waterlevel.GetWaterlevelLatestForFloodforecast()
	return oData, err
}

//	get forecast latest by station id
//	Parameters:
//		cpyStation
//			 หา ตาม station_id
//	Return:
//		[]FloodForecastOutPutCpy
func floodForecastLatest(cpyStation []int64) ([]*FloodForecastOutPutCpy, error) {

	p := []interface{}{}
	var condition string
	var sqlSelectMaxDatetime string
	var sqlSelectInterval string
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	// loop add condition station id
	_getMaxDateFloodForecast := strings.Replace(getMaxDateFloodForecast, "--replace", "- interval '14 day'", 1)
	if len(cpyStation) > 0 {
		var station string
		for i, v := range cpyStation {
			if i != 0 {
				station += " OR floodforecast_station_id=$" + strconv.Itoa(i+1)
			} else {
				station = "floodforecast_station_id=$" + strconv.Itoa(i+1)
			}
			p = append(p, v)
		}
		condition = " WHERE (" + station + ") "
		sqlSelectMaxDatetime = _getMaxDateFloodForecast + condition + getMaxDateFloodForecastOrderBy
		// sqlSelectInterval = getCpyInterval + condition + " AND floodforecast_datetime > (" + sqlSelectMaxDatetime + ") - interval '14 day'" + getCpyIntervalOrderBy
		sqlSelectInterval = getCpyInterval + condition + " AND floodforecast_datetime > $" + strconv.Itoa(len(p)+1) + getCpyIntervalOrderBy

	} else {
		sqlSelectMaxDatetime = _getMaxDateFloodForecast + getMaxDateFloodForecastOrderBy
		sqlSelectInterval = getCpyInterval + " WHERE floodforecast_datetime > $1" + getCpyIntervalOrderBy
	}
	var t time.Time
	err = db.QueryRow(sqlSelectMaxDatetime, p...).Scan(&t)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	q := sqlSelectInterval
	p = append(p, t)
	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// define value
	data := make([]*FloodForecastOutPutCpy, 0)
	valueIns := make([]*FloodForecastCpyValue, 0)
	dataRow := &FloodForecastOutPutCpy{}
	for rows.Next() {
		var (
			id             sql.NullInt64
			stationOldCode sql.NullString
			datetime       time.Time
			value          sql.NullFloat64
			lat            sql.NullFloat64
			long           sql.NullFloat64
			unit           sql.NullString
			mid            sql.NullInt64
			mName          pqx.JSONRaw
			mAlarm         sql.NullFloat64
			mWarning       sql.NullFloat64
			aid            sql.NullInt64
			aNmae          pqx.JSONRaw
			aShortName     pqx.JSONRaw
		)

		// scan data from database
		rows.Scan(&stationOldCode, &id, &datetime, &value, &mid, &mName, &lat, &long, &mAlarm, &mWarning, &unit, &aid, &aNmae, &aShortName)

		if dataRow.Station != nil {
			// check duplicate oldcode for add array data
			if dataRow.Station.FloodForecastOldCode != stationOldCode.String {
				dataRow.FloodForecastData = valueIns
				data = append(data, dataRow)
				dataRow = &FloodForecastOutPutCpy{}
				dataRow.ID = id.Int64
				valueIns = make([]*FloodForecastCpyValue, 0)

				//				dataRow.Station.FloodForecastOldCode = stationOldCode.String
				station := &FloodForecastStation{}
				station.ID = mid.Int64
				station.FloodforecastName = mName.JSON()
				station.FloodForecastOldCode = stationOldCode.String
				if lat.Valid {
					station.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
				}
				if long.Valid {
					station.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
				}

				station.Alarm = ValidData(mAlarm.Valid, mAlarm.Float64)
				station.Warning = ValidData(mWarning.Valid, mWarning.Float64)
				dataRow.Station = station

				agency := &Agency{}
				agency.ID = aid.Int64
				agency.Name = aNmae.JSON()
				agency.ShortName = aShortName.JSON()
				dataRow.Agency = agency

				dataIn := &FloodForecastCpyValue{}
				dataIn.FloodForecastDatetime = udt.DatetimeFormat(datetime, "datetime")
				if value.Valid {
					dataIn.FloodForecastValue = strconv.FormatFloat(value.Float64, 'f', 6, 64)
					dataRow.FloodForecastValue = strconv.FormatFloat(value.Float64, 'f', 6, 64)
				}
				dataRow.FloodForecastDatetime = udt.DatetimeFormat(datetime, "datetime")
				valueIns = append(valueIns, dataIn)
				dataRow.FloodForecastData = valueIns
			} else {
				dataIn := &FloodForecastCpyValue{}
				dataIn.FloodForecastDatetime = udt.DatetimeFormat(datetime, "datetime")
				if value.Valid {
					dataIn.FloodForecastValue = strconv.FormatFloat(value.Float64, 'f', 6, 64)
				}
				valueIns = append(valueIns, dataIn)
				dataRow.FloodForecastData = valueIns
			}
		} else {
			// add data first row
			//			dataRow.Station.FloodForecastOldCode = stationOldCode.String
			dataRow.ID = id.Int64
			station := &FloodForecastStation{}
			station.ID = mid.Int64
			station.FloodforecastName = mName.JSON()
			station.FloodForecastOldCode = stationOldCode.String
			if lat.Valid {
				station.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
			}
			if long.Valid {
				station.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
			}
			station.Alarm = ValidData(mAlarm.Valid, mAlarm.Float64)
			station.Warning = ValidData(mWarning.Valid, mWarning.Float64)
			dataRow.Station = station

			agency := &Agency{}
			agency.ID = aid.Int64
			agency.Name = aNmae.JSON()
			agency.ShortName = aShortName.JSON()
			dataRow.Agency = agency

			dataIn := &FloodForecastCpyValue{}
			dataIn.FloodForecastDatetime = udt.DatetimeFormat(datetime, "datetime")
			if value.Valid {
				dataIn.FloodForecastValue = strconv.FormatFloat(value.Float64, 'f', 6, 64)
				dataRow.FloodForecastValue = strconv.FormatFloat(value.Float64, 'f', 6, 64)
			}
			dataRow.FloodForecastDatetime = udt.DatetimeFormat(datetime, "datetime")
			valueIns = append(valueIns, dataIn)
			dataRow.FloodForecastData = valueIns
		}
	}

	// add array data to output
	if dataRow.Station != nil {
		if dataRow.Station.FloodForecastOldCode != "" {
			dataRow.FloodForecastData = valueIns
			data = append(data, dataRow)
		}
	}

	return data, nil
}

//	get swan station information
//  Parameters:
//		None
//	Return:
//		[]SwanStationOutput
func SwanStationDetails() ([]*SwanStationOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	q := getSwanStation
	p := []interface{}{}

	//query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*SwanStationOutput, 0)

	for rows.Next() {
		var (
			stationName sql.NullString
			lat         sql.NullFloat64
			long        sql.NullFloat64
		)
		// scan data
		rows.Scan(&stationName, &lat, &long)
		dataRow := &SwanStationOutput{}
		dataRow.SwanName = json.RawMessage(stationName.String)
		dataRow.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
		dataRow.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
		data = append(data, dataRow)
	}
	return data, nil
}

//	get swan forecsat ล่าสุด
//	Parameters:
//		param
//			ใช้ในส่วนของ Station_id
//	Return:
//		[]SwanForecastOutput
func SwanForecastLatest(param *Struct_Swan_Forecast_InputParam) ([]*SwanForecastOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}
	
	// querystring station_id
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""
	arrStationId := []string{}
	if param.StationID != "" {
		arrStationId = strings.Split(param.StationID, ",")
	}
	
	if len(arrStationId) > 0 {
		if len(arrStationId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.StationID, " "))
			sqlCmdWhere += " AND s.swan_station_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrStationId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND s.swan_station_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}
	
	// get sql
	q := getSwanForecast
		
	// query
	rows, err := db.Query(q+sqlCmdWhere+getSwanForecastOrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*SwanForecastOutput, 0)
	valueIns := make([]*SwanForecastValue, 0)
	dataRow := &SwanForecastOutput{}
	var nameCp string
	for rows.Next() {
		var (
			swanName  sql.NullString
			datetime  time.Time
			depth     sql.NullFloat64
			highsig   sql.NullFloat64
			direction sql.NullFloat64
			periodTop sql.NullFloat64
			periodAvg sql.NullFloat64
			windx     sql.NullFloat64
			windy     sql.NullFloat64
			id        sql.NullInt64
			mid       sql.NullInt64
			lat       sql.NullFloat64
			long      sql.NullFloat64
		)
		
		// scan data
		rows.Scan(&swanName, &id, &datetime, &depth, &highsig, &direction, &periodTop, &periodAvg, &windx, &windy, &mid, &lat, &long)

		// check first row
		if nameCp != "" {
			// check duplicate station for add data to array
			if nameCp != swanName.String {
				dataRow.SwanForecastData = valueIns
				data = append(data, dataRow)
				dataRow = &SwanForecastOutput{}
				valueIns = make([]*SwanForecastValue, 0)

				nameCp = swanName.String
				station := &SwanStationOutput2{}
				station.ID = mid.Int64
				station.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
				station.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
				station.SwanName = json.RawMessage(swanName.String)
				dataRow.Station = station
				dataIn := &SwanForecastValue{}
				dataIn.ID = id.Int64
				dataIn.Datetime = udt.DatetimeFormat(datetime, "datetime")
				dataIn.Depth = ValidData(depth.Valid, depth.Float64)
				dataIn.Highsig = ValidData(highsig.Valid, highsig.Float64)
				dataIn.Direction = ValidData(direction.Valid, direction.Float64)
				dataIn.PeriodTop = ValidData(periodTop.Valid, periodTop.Float64)
				dataIn.PeriodAverage = ValidData(periodAvg.Valid, periodAvg.Float64)
				dataIn.WindX = ValidData(windx.Valid, windx.Float64)
				dataIn.WindY = ValidData(windy.Valid, windy.Float64)
				valueIns = append(valueIns, dataIn)
				dataRow.SwanForecastData = valueIns
			} else {
				// add data if station duplicate
				dataIn := &SwanForecastValue{}
				dataIn.ID = id.Int64
				dataIn.Datetime = udt.DatetimeFormat(datetime, "datetime")
				dataIn.Depth = ValidData(depth.Valid, depth.Float64)
				dataIn.Highsig = ValidData(highsig.Valid, highsig.Float64)
				dataIn.Direction = ValidData(direction.Valid, direction.Float64)
				dataIn.PeriodTop = ValidData(periodTop.Valid, periodTop.Float64)
				dataIn.PeriodAverage = ValidData(periodAvg.Valid, periodAvg.Float64)
				dataIn.WindX = ValidData(windx.Valid, windx.Float64)
				dataIn.WindY = ValidData(windy.Valid, windy.Float64)
				valueIns = append(valueIns, dataIn)
				dataRow.SwanForecastData = valueIns
			}
		} else {
			// add data first row
			nameCp = swanName.String
			station := &SwanStationOutput2{}
			station.ID = mid.Int64
			station.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
			station.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
			station.SwanName = json.RawMessage(swanName.String)
			dataRow.Station = station
			dataIn := &SwanForecastValue{}
			dataIn.ID = id.Int64
			dataIn.Datetime = udt.DatetimeFormat(datetime, "datetime")
			dataIn.Depth = ValidData(depth.Valid, depth.Float64)
			dataIn.Highsig = ValidData(highsig.Valid, highsig.Float64)
			dataIn.Direction = ValidData(direction.Valid, direction.Float64)
			dataIn.PeriodTop = ValidData(periodTop.Valid, periodTop.Float64)
			dataIn.PeriodAverage = ValidData(periodAvg.Valid, periodAvg.Float64)
			dataIn.WindX = ValidData(windx.Valid, windx.Float64)
			dataIn.WindY = ValidData(windy.Valid, windy.Float64)
			valueIns = append(valueIns, dataIn)
			dataRow.SwanForecastData = valueIns
		}
	}

	if nameCp != "" {
		dataRow.SwanForecastData = valueIns
		data = append(data, dataRow)
	}

	return data, nil
}

func FloodForecastSubbasin() ([]*FloodForecastMonitoring, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	q := getFloodForecast
	p := []interface{}{10}
	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*FloodForecastMonitoring, 0)

	for rows.Next() {
		var (
			subbasin_name              pqx.JSONRaw
			floodforecast_station_name pqx.JSONRaw
			floodforecast_datetime     time.Time
			floodforecast_value        sql.NullFloat64
		)
		// scan data
		rows.Scan(&subbasin_name, &floodforecast_station_name, &floodforecast_datetime, &floodforecast_value)
		dataRow := &FloodForecastMonitoring{}
		dataRow.SubbasinName = subbasin_name.JSON()
		dataRow.FloodforecastName = floodforecast_station_name.JSON()
		dataRow.Datetime = floodforecast_datetime.Format("2006-01-02 15:04")
		dataRow.Value = ValidData(floodforecast_value.Valid, floodforecast_value.Float64)
		data = append(data, dataRow)
	}

	return data, nil
}

// check valid data
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}

	return nil
}
