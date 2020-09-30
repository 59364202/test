// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Peerapong Srisom <peerapong@haii.co.th>
//

// Package forecast is a model for provinces.floodforecast_waterlevel table. This table store forecast data.
package floodforecast_waterlevel

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/server/model/setting"
	//waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"time"
	//	"fmt"
	//	"os"
	"strings"
)

// 	get cpy forecast latest
//	Return:
//		FloodforecastOutputWithScale
func CpyFloodForecastLatest(param *FloodforecastInput) (*FloodforecastOutputWithScale, error) {
	var err error
	oData := &FloodforecastOutputWithScale{}
	oData.Scale = setting.GetSystemSettingJson("Frontend.public.waterlevel_setting")
	oData.FloodforecastLoad, err = floodForecastLatest(param)
	//oData.WaterlevelObserve, err = waterlevel.GetWaterlevelLatestForFloodforecast()
	return oData, err
}

//	get forecast latest by province_code
//	Parameters:
//		cpyStation
//			 หา ตาม province_code
//	Return:
//		[]FloodForecastOutPutCpy
func floodForecastLatest(param *FloodforecastInput) ([]*FloodForecastOutPutCpy, error) {

	p := []interface{}{}

	var (
		sqlSelectMaxDatetime string
		sqlSelectInterval    string
		strWhere             string
	)

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	// Province Fillter
	arrProvinceId := []string{}
	if param.Province_Code != "" {
		arrProvinceId = strings.Split(param.Province_Code, ",")
	}

	if len(arrProvinceId) > 0 {
		strWhere += " AND "
		if len(arrProvinceId) == 1 {
			p = append(p, strings.Trim(param.Province_Code, " "))
			strWhere += " geo.province_code = $" + strconv.Itoa(len(p))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				p = append(p, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(p)))
			}
			strWhere += " geo.province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	sqlSelectMaxDatetime = getMaxDateFloodForecast + getMaxDateFloodForecastOrderBy
	sqlSelectInterval = getCpyInterval + " WHERE floodforecast_datetime > (" + sqlSelectMaxDatetime + ") - interval '8 day' " + strWhere + " " + getCpyIntervalOrderBy

	//fmt.Print(sqlSelectInterval)
	//	os.Exit(3)
	q := sqlSelectInterval

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
			mCritical       sql.NullFloat64
			aid            sql.NullInt64
			aNmae          pqx.JSONRaw
			aShortName     pqx.JSONRaw
			tumbon_name    pqx.JSONRaw
			amphoe_name    pqx.JSONRaw
			province_code  sql.NullString
			province_name  pqx.JSONRaw
			area_code      sql.NullString
			area_name      sql.NullString
			geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		// scan data from database
		rows.Scan(&stationOldCode, &id, &datetime, &value, &mid, &mName, &lat, &long, &mAlarm, &mWarning, &mCritical, &unit, &aid, &aNmae, &aShortName, &area_name, &area_code, &province_name, &province_code, &amphoe_name, &tumbon_name)

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
				station.Critical = ValidData(mCritical.Valid, mCritical.Float64)
				dataRow.Station = station

				agency := &Agency{}
				agency.ID = aid.Int64
				agency.Name = aNmae.JSON()
				agency.ShortName = aShortName.JSON()
				dataRow.Agency = agency

				geocode.Tumbon_name = tumbon_name.JSON()
				geocode.Amphoe_name = amphoe_name.JSON()
				geocode.Province_code = province_code.String
				geocode.Province_name = province_name.JSON()
				geocode.Area_code = area_code.String
				if area_name.String == "" {
					area_name.String = "{}"
				}
				geocode.Area_name = json.RawMessage(area_name.String)

				dataRow.Geocode = geocode

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
			//	dataRow.Station.FloodForecastOldCode = stationOldCode.String
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
			station.Critical = ValidData(mCritical.Valid, mCritical.Float64)
			dataRow.Station = station

			agency := &Agency{}
			agency.ID = aid.Int64
			agency.Name = aNmae.JSON()
			agency.ShortName = aShortName.JSON()
			dataRow.Agency = agency

			geocode.Tumbon_name = tumbon_name.JSON()
			geocode.Amphoe_name = amphoe_name.JSON()
			geocode.Province_code = province_code.String
			geocode.Province_name = province_name.JSON()
			geocode.Area_code = area_code.String
			if area_name.String == "" {
				area_name.String = "{}"
			}
			geocode.Area_name = json.RawMessage(area_name.String)

			dataRow.Geocode = geocode

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
 
// check valid data
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}

	return nil
}
