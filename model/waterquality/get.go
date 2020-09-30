// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package waterquality is a model for public.waterquality table. This table store waterquality.
package waterquality

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"haii.or.th/api/server/model/setting"
	model_waterquality_station "haii.or.th/api/thaiwater30/model/waterquality_station"
	//	udt "haii.or.th/api/thaiwater30/util/datetime"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/thaiwater30/util/validdata"
	//	"haii.or.th/api/thaiwater30/util/highchart"
	"sort"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

// get waterquality latetst
//  Parameters:
//		station_id
//			waterquality station id
//  Return:
//		Array Struct_WaterQuality
func Get_WaterQuanlityLatest(station_id int64) ([]*Struct_WaterQuality, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row          *sql.Rows
		data         []*Struct_WaterQuality
		waterquality *Struct_WaterQuality
		station      *model_waterquality_station.Struct_WaterQualityStation

		_waterquality_id           int64
		_waterquality_datetime     time.Time
		_waterquality_do           sql.NullFloat64
		_waterquality_ph           sql.NullFloat64
		_waterquality_temp         sql.NullFloat64
		_waterquality_turbid       sql.NullFloat64
		_waterquality_bod          sql.NullFloat64
		_waterquality_tcb          sql.NullFloat64
		_waterquality_fcb          sql.NullFloat64
		_waterquality_nh3n         sql.NullFloat64
		_waterquality_wqi          sql.NullFloat64
		_waterquality_ammonium     sql.NullFloat64
		_waterquality_nitrate      sql.NullFloat64
		_waterquality_colorstatus  sql.NullString
		_waterquality_status       sql.NullString
		_waterquality_salinity     sql.NullFloat64
		_waterquality_conductivity sql.NullFloat64
		_waterquality_tds          sql.NullFloat64
		_waterquality_chlorophyll  sql.NullFloat64
		_waterquality_station_name sql.NullString
		_waterquality_station_lat  sql.NullFloat64
		_waterquality_station_long sql.NullFloat64
		_province_name             sql.NullString
		_province_code             sql.NullString
		_amphoe_name               sql.NullString
		_tumbon_name               sql.NullString
		_agency_id                 sql.NullInt64
		_agency_name               sql.NullString
		_agency_shortname          sql.NullString
		_is_active                 sql.NullString
		_show_status               pqx.JSONRaw

		_data_id sql.NullInt64
		_oldcode sql.NullString
	)
	if station_id > 0 {
		row, err = db.Query(SQL_SelectWaterQuanlityLatest+" AND mws.id = $1 ", station_id)
	} else {
		row, err = db.Query(SQL_SelectWaterQuanlityLatest)
	}

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_WaterQuality, 0)
	for row.Next() {
		err = row.Scan(&_waterquality_id, &_waterquality_datetime, &_waterquality_do, &_waterquality_ph, &_waterquality_temp, &_waterquality_turbid,
			&_waterquality_bod, &_waterquality_tcb, &_waterquality_fcb, &_waterquality_nh3n, &_waterquality_wqi, &_waterquality_ammonium, &_waterquality_nitrate, &_waterquality_colorstatus,
			&_waterquality_status, &_waterquality_salinity, &_waterquality_conductivity, &_waterquality_tds, &_waterquality_chlorophyll, &_waterquality_station_lat,
			&_waterquality_station_long, &_waterquality_station_name, &_province_name, &_province_code, &_amphoe_name, &_tumbon_name, &_agency_name, &_agency_shortname, &_data_id, &_oldcode, &_agency_id, &_is_active, &_show_status)
		if err != nil {
			return nil, err
		}
		if !_waterquality_station_name.Valid || _waterquality_station_name.String == "" {
			_waterquality_station_name.String = "{}"
		}
		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}
		if !_amphoe_name.Valid || _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		if !_tumbon_name.Valid || _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}

		waterquality = &Struct_WaterQuality{Waterquality_Datetime: _waterquality_datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))}
		waterquality.Id = _data_id.Int64
		waterquality.Station_type = "waterquality"
		waterquality.Waterquality_Colorstatus = _waterquality_colorstatus.String
		waterquality.Waterquality_Status = _waterquality_status.String

		Do := validdata.ValidData(_waterquality_do.Valid, _waterquality_do.Float64)
		Ph := validdata.ValidData(_waterquality_ph.Valid, _waterquality_ph.Float64)
		Temp := validdata.ValidData(_waterquality_temp.Valid, _waterquality_temp.Float64)
		Turbid := validdata.ValidData(_waterquality_turbid.Valid, _waterquality_turbid.Float64)
		Bod := validdata.ValidData(_waterquality_bod.Valid, _waterquality_bod.Float64)
		Tcb := validdata.ValidData(_waterquality_tcb.Valid, _waterquality_tcb.Float64)
		Fcb := validdata.ValidData(_waterquality_fcb.Valid, _waterquality_fcb.Float64)
		Nh3n := validdata.ValidData(_waterquality_nh3n.Valid, _waterquality_nh3n.Float64)
		Wqi := validdata.ValidData(_waterquality_wqi.Valid, _waterquality_wqi.Float64)
		Ammonium := validdata.ValidData(_waterquality_ammonium.Valid, _waterquality_ammonium.Float64)
		Nitrate := validdata.ValidData(_waterquality_nitrate.Valid, _waterquality_nitrate.Float64)
		Salinity := validdata.ValidData(_waterquality_salinity.Valid, _waterquality_salinity.Float64)
		Conductivity := validdata.ValidData(_waterquality_conductivity.Valid, _waterquality_conductivity.Float64)
		Tds := validdata.ValidData(_waterquality_tds.Valid, _waterquality_tds.Float64)
		Chlorophyll := validdata.ValidData(_waterquality_chlorophyll.Valid, _waterquality_chlorophyll.Float64)

		// ตรวจสอบแต่ละ column ว่าจะแสดงผล column ไหนบ้าง ถ้าไม่แสดง column ไหน จะ return ค่าเป็น null
		//ข้อมูล คุณภาพน้ำจากสถานีฯ คพ. ที่ส่งไปยัง สสนก. เป็นข้อมูลดิบที่ยังไม่ได้กรอง ซึ่งทำให้บางสถานี ที่ไม่มีหัววัด ส่งค่าไปเป็น ค่า ศูนย์ เช่น หัววัด DO หรือบางสถานี ค่า error
		var show_status WaterQualityshowStatusStruct
		if err := json.Unmarshal(_show_status, &show_status); err != nil {
			return nil, err
		}
		waterquality.Waterquality_Do = validdata.ValidData(show_status.Do, Do)
		waterquality.Waterquality_Ph = validdata.ValidData(show_status.Ph, Ph)
		waterquality.Waterquality_Temp = validdata.ValidData(show_status.Temp, Temp)
		waterquality.Waterquality_Turbid = validdata.ValidData(show_status.Turbid, Turbid)
		waterquality.Waterquality_Bod = validdata.ValidData(show_status.Bod, Bod)
		waterquality.Waterquality_Tcb = validdata.ValidData(show_status.Tcb, Tcb)
		waterquality.Waterquality_Fcb = validdata.ValidData(show_status.Fcb, Fcb)
		waterquality.Waterquality_Nh3n = validdata.ValidData(show_status.Nh3n, Nh3n)
		waterquality.Waterquality_Wqi = validdata.ValidData(show_status.Wqi, Wqi)
		waterquality.Waterquality_Ammonium = validdata.ValidData(show_status.Ammonium, Ammonium)
		waterquality.Waterquality_Nitrate = validdata.ValidData(show_status.Nitrate, Nitrate)
		waterquality.Waterquality_Salinity = validdata.ValidData(show_status.Salinity, Salinity)
		waterquality.Waterquality_Conductivity = validdata.ValidData(show_status.Conductivity, Conductivity)
		waterquality.Waterquality_Tds = validdata.ValidData(show_status.Tds, Tds)
		waterquality.Waterquality_Chlorophyll = validdata.ValidData(show_status.Chlorophyll, Chlorophyll)

		station = &model_waterquality_station.Struct_WaterQualityStation{Id: _waterquality_id}
		station.Waterquality_Station_Name = json.RawMessage(_waterquality_station_name.String)
		station.Province_Name = json.RawMessage(_province_name.String)
		station.Amphoe_Name = json.RawMessage(_amphoe_name.String)
		station.Tumbon_Name = json.RawMessage(_tumbon_name.String)
		station.Province_Code = _province_code.String
		station.Waterquality_Station_Lat, _ = _waterquality_station_lat.Value()
		station.Waterquality_Station_Long, _ = _waterquality_station_long.Value()
		station.Agency_id = _agency_id.Int64
		station.Agency_Name = json.RawMessage(_agency_name.String)
		station.Agency_Shortname = json.RawMessage(_agency_shortname.String)
		station.Waterquality_Station_Oldcode = _oldcode.String
		station.Is_active = _is_active.String

		waterquality.Waterquality_Station = station

		//waterquality.ProvinceName = json.RawMessage(_province_name.String)
		//waterquality.Name = json.RawMessage(_waterquality_station_name.String)
		//waterquality.Datetime = _waterquality_datetime.Format("2006-01-02 15:04")
		//waterquality.Oldcode = _oldcode.String
		//waterquality.Value = ValidData(_waterquality_salinity.Valid, _waterquality_salinity.Float64)
		//waterquality.DataID = _data_id.Int64
		//waterquality.StationID = _waterquality_id

		data = append(data, waterquality)
	}

	return data, nil
}

//func Get_WaterQualityGraph(param *Param_WaterQualityGraph) (*highchart.HighChart, error) {
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, err
//	}
//	var (
//		_v        sql.NullFloat64
//		_datetime time.Time
//	)
//
//	strSQL := SQL_SelectWaterQualityGraph(param.DateType)
//	row, err := db.Query(strSQL, param.Id, param.DateStart, param.DateEnd)
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//	hc := highchart.NewChart()
//	series := hc.NewSerie("")
//	series.CurrentYear()
//	for row.Next() {
//		err = row.Scan(&_v, &_datetime)
//		if err != nil {
//			return nil, err
//		}
//		if !_v.Valid {
//			series.AddDateData(_datetime, nil)
//		} else {
//			series.AddDateData(_datetime, _v.Float64)
//		}
//	}
//
//	return hc, nil
//}

//  GetWaterquality compare analyst
//  Parameters:
//		inputData
//			WaterQualityGraphCompareAnalystInput
//  Return:
//		Array WaterQualityGraphCompareAnalystOutput2
func GetWaterQualityGraphCompareAnalyst(inputData *WaterQualityGraphCompareAnalystInput) ([]*WaterQualityGraphCompareAnalystOutput2, error) {

	inputStation := ""
	ids := strings.Split(inputData.DatetimeStart, "-")
	ide := strings.Split(inputData.DatetimeEnd, "-")
	if len(ids) > 0 {
		y, err := strconv.Atoi(ids[0])
		if err != nil {
			return nil, err
		}
		ids[0] = strconv.Itoa(y)
	}

	if len(ide) > 0 {
		y, err := strconv.Atoi(ide[0])
		if err != nil {
			return nil, err
		}
		ide[0] = strconv.Itoa(y)
	}

	p := []interface{}{strings.Join(ids, "-"), strings.Join(ide, "-")}
	for i, v := range inputData.WaterQualityStation {
		if i > 0 {
			inputStation += " OR wq.waterquality_id=$" + strconv.Itoa(i+3)
		} else {
			inputStation = "wq.waterquality_id=$" + strconv.Itoa(i+3)
		}
		p = append(p, v)
	}

	q := sqlSelectWaterQualityGraphCompare + ",wq.waterquality_" + inputData.Param
	q += sqlSelectWaterQualityGraphCompareFROM

	if inputStation != "" {
		q += "AND (" + inputStation + ")"
	}

	q += sqlSelectWaterQualityGraphCompareORDER
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*WaterQualityGraphCompareAnalystOutput2, 0)
	mapSeriesName := make(map[int64]*WaterQualityGraphCompareAnalystOutput)
	mapSeriesDataTime := make(map[int64]map[string]interface{})
	for rows.Next() {
		var (
			id          int64
			stationName pqx.JSONRaw
			datetime    time.Time
			agencyID    sql.NullInt64
			value       sql.NullFloat64
		)

		rows.Scan(&id, &stationName, &datetime, &agencyID, &value)

		msn := mapSeriesName[id]
		mst := mapSeriesDataTime[id]
		if msn != nil {
			mst[datetime.Format("2006-01-02 15:04")] = ValidData(value.Valid, value.Float64)
			msn.Data = mst
			mapSeriesDataTime[id] = mst
			mapSeriesName[id] = msn
		} else {
			seriesRow := &WaterQualityGraphCompareAnalystOutput{}
			dr, err := generateDateRange(strings.Join(ids, "-"), strings.Join(ide, "-"), getRangeTime(agencyID.Int64))
			if err != nil {
				return nil, err
			}
			mapSeriesDataTime[id] = dr
			seriesRow.SeriesName = json.RawMessage(stationName.JSON())
			dr[datetime.Format("2006-01-02 15:04")] = ValidData(value.Valid, value.Float64)
			mapSeriesDataTime[id] = dr
			seriesRow.Data = dr
			mapSeriesName[id] = seriesRow
		}
	}

	for k, v := range mapSeriesName {
		dataRow := &WaterQualityGraphCompareAnalystOutput2{}
		dataRow.SeriesName = v.SeriesName
		sdt := mapSeriesDataTime[k]
		datetimeOutput := make([]*WaterQualityGraphCompareAnalystData, 0)
		for d, i := range sdt {
			dto := &WaterQualityGraphCompareAnalystData{}
			dto.Datetime = d
			dto.Value = i
			datetimeOutput = append(datetimeOutput, dto)
		}
		dataSort := make(DataRange, 0)
		dataSort = append(dataSort, datetimeOutput...)
		sort.Sort(dataSort)
		datetimeOutput = dataSort
		dataRow.Data = datetimeOutput
		data = append(data, dataRow)
	}

	return data, nil
}

//  Parameters:
//		datetimeStart
//			start datetime for generate
//		datetimeEnd
//			end datetime for generate
//		timeRange
//			time range
//  Return:
//		Array information cctv
func generateDateRange(datetimeStart, datetimeEnd string, timeRange int) (map[string]interface{}, error) {
	t1, err := time.Parse("2006-01-02 15:04", datetimeStart)
	if err != nil {

		return nil, err
	}
	t2, err := time.Parse("2006-01-02 15:04", datetimeEnd)
	if err != nil {
		return nil, err
	}
	if float64(t1.Minute())/float64(timeRange) != 0 {
		iMin := t1.Minute() / timeRange
		fMin := float64(t1.Minute()) / float64(timeRange)
		tMin := fMin - float64(iMin)
		min := timeRange - int(tMin)
		t1 = t1.Add(time.Duration(min) * time.Minute)
	}
	data := make(map[string]interface{}, 0)
	for t1.Before(t2) {
		data[t1.Format("2006-01-02 15:04")] = nil
		t1 = t1.Add(time.Duration(timeRange) * time.Minute)
	}
	return data, nil
}

// get waterquality for graph
//  Parameters:
//		inputData
//			WaterQualityGraphParamsAnalystInput
//  Return:
//		Array WaterQualityGraphParamsAnalystOutput2
func GetWaterQualiyGraphParamsAnalyst(inputData *WaterQualityGraphParamsAnalystInput) ([]*WaterQualityGraphParamsAnalystOutput2, error) {

	if inputData.WaterQualityStation == 0 {
		return nil, rest.NewError(422, "No station", nil)
	}
	ids := strings.Split(inputData.DatetimeStart, "-")
	ide := strings.Split(inputData.DatetimeEnd, "-")
	if len(ids) > 0 {
		y, err := strconv.Atoi(ids[0])
		if err != nil {
			return nil, err
		}
		ids[0] = strconv.Itoa(y)
	}

	if len(ide) > 0 {
		y, err := strconv.Atoi(ide[0])
		if err != nil {
			return nil, err
		}
		ide[0] = strconv.Itoa(y)
	}

	p := []interface{}{strings.Join(ids, "-"), strings.Join(ide, "-"), inputData.WaterQualityStation}

	for i, v := range inputData.Param {
		if v != "" {
			inputData.Param[i] = "waterquality_" + v
		} else {
			return nil, rest.NewError(422, "Null input params", nil)
		}
	}

	q := sqlSelectWaterQualityGraphMultiParams + "," + strings.Join(inputData.Param, ",") + sqlSelectWaterQualityGraphMultiParamsFrom
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*WaterQualityGraphParamsAnalystOutput2, 0)
	mapSeriesName := make(map[string]*WaterQualityGraphParamsAnalystOutput)
	mapSeriesDataTime := make(map[string]map[string]interface{})
	for rows.Next() {
		var (
			id          int64
			stationName pqx.JSONRaw
			datetime    time.Time
			agencyID    sql.NullInt64
		)
		valuePtrs := make([]interface{}, len(inputData.Param)+4)
		valuePtrs[0] = &id
		valuePtrs[1] = &stationName
		valuePtrs[2] = &datetime
		valuePtrs[3] = &agencyID
		value := make([]sql.NullFloat64, len(inputData.Param))
		for i, _ := range inputData.Param {
			valuePtrs[i+4] = &value[i]
		}

		rows.Scan(valuePtrs...)

		for i, v := range inputData.Param {
			msn := mapSeriesName[v]
			mst := mapSeriesDataTime[v]
			if msn != nil {
				mst[datetime.Format("2006-01-02 15:04")] = ValidData(value[i].Valid, value[i].Float64)
				msn.Data = mst
				mapSeriesDataTime[v] = mst
				mapSeriesName[v] = msn
			} else {
				seriesRow := &WaterQualityGraphParamsAnalystOutput{}
				dr, err := generateDateRange(strings.Join(ids, "-"), strings.Join(ide, "-"), getRangeTime(agencyID.Int64))
				if err != nil {
					return nil, err
				}
				mapSeriesDataTime[v] = dr
				seriesRow.SeriesName = v
				dr[datetime.Format("2006-01-02 15:04")] = ValidData(value[i].Valid, value[i].Float64)
				mapSeriesDataTime[v] = dr
				seriesRow.Data = dr
				mapSeriesName[v] = seriesRow
			}
		}
	}

	for k, v := range mapSeriesName {
		dataRow := &WaterQualityGraphParamsAnalystOutput2{}
		dataRow.SeriesName = v.SeriesName
		sdt := mapSeriesDataTime[k]
		datetimeOutput := make([]*WaterQualityGraphCompareAnalystData, 0)
		for d, i := range sdt {
			dto := &WaterQualityGraphCompareAnalystData{}
			dto.Datetime = d
			dto.Value = i
			datetimeOutput = append(datetimeOutput, dto)
		}
		dataSort := make(DataRange, 0)
		dataSort = append(dataSort, datetimeOutput...)
		sort.Sort(dataSort)
		datetimeOutput = dataSort
		dataRow.Data = datetimeOutput
		data = append(data, dataRow)
	}

	return data, nil
}

//	Get waterqualiy graph waterlevel
//  Parameters:
//		inputData
//			WaterQualityGraphWaterlevelAnalystInput
//  Return:
//		Array WaterQualityWaterlevelOutput
func GetWaterQualiyGraphWaterlevel(inputData *WaterQualityGraphWaterlevelAnalystInput) ([]*WaterQualityWaterlevelOutput, error) {

	var q string
	var err error
	var dataRow *WaterQualityWaterlevelOutput
	// Waterquality
	data := make([]*WaterQualityWaterlevelOutput, 0)
	q = sqlSelectWaterquality + sqlSelectWaterqualityField + ",waterquality_" + inputData.ParamWQ + " " +
		sqlSelectWaterqualityEnd
	dataRow, err = getDataWaterQuality(inputData, q, "")
	if err != nil {
		return nil, err
	}
	data = append(data, dataRow)

	if inputData.WaterlevelStationType == "tele_waterlevel" {
		q = sqlSelectWaterlevel + ",waterlevel_msl " + sqlSelectWaterlevelFrom
		dataRow, err = getDataWaterQuality(inputData, q, inputData.WaterlevelStationType)
		data = append(data, dataRow)
	} else if inputData.WaterlevelStationType == "canal_waterlevel" {
		q = sqlSelectCanal
		dataRow, err = getDataWaterQuality(inputData, q, inputData.WaterlevelStationType)
		data = append(data, dataRow)
	}

	if err != nil {
		return nil, err
	}
	return data, nil
}

// 	Get WaterQuality Graph Compare DatetimeAndStation
//  Parameters:
//		inputData
//			WaterQualityGraphCompareDatetimeInput
//  Return:
//		Array WaterQualityCompareDatetimeOutput
func GetWaterQualityGraphCompareDatetimeAndStation(inputData *WaterQualityGraphCompareDatetimeInput) ([]*WaterQualityCompareDatetimeOutput, error) {

	if inputData.Param == "" {
		return nil, rest.NewError(422, "invalid field param", nil)
	}
	p := []interface{}{}

	station := []interface{}{}
	condition := ""
	for i, v := range inputData.WaterQualityStation {
		if i > 0 {
			condition += " OR waterquality_id=$" + strconv.Itoa(i+2)
		} else {
			condition = "waterquality_id=$" + strconv.Itoa(i+2)
		}
		station = append(station, v)
	}
	data := make([]*WaterQualityCompareDatetimeOutput, 0)
	for _, v := range inputData.Date {
		p = []interface{}{v}
		p = append(p, station...)
		q := sqlSelectWaterqualityCompare2 + ",waterquality_" + inputData.Param + " " +
			sqlSelectWaterqualityCompareFrom2 + " AND (" + condition + ")"

		db, err := pqx.Open()
		if err != nil {
			return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
		}

		rows, err := db.Query(q, p...)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}
		defer rows.Close()
		dataRow := &WaterQualityCompareDatetimeOutput{}
		dataRow.Datetime = v
		dataStation := make([]*WaterQualityCompareDatetimeOutputStation, 0)
		mapCompare := getStation(inputData.WaterQualityStation)
		for rows.Next() {
			var (
				stationID sql.NullInt64
				value     sql.NullFloat64
			)
			rows.Scan(&stationID, &value)

			stationData := mapCompare[stationID.Int64]
			stationData.Value = ValidData(value.Valid, value.Float64)
			mapCompare[stationID.Int64] = stationData
		}
		for _, v := range inputData.WaterQualityStation {
			dataStation = append(dataStation, mapCompare[v])
		}
		dataRow.Station = dataStation
		data = append(data, dataRow)
	}

	return data, nil
}

//  get waterquality salinity
//  Parameters:
//		None
//  Return:
//		Array MonitoringWaterqualityOutput
func GetWaterQualitySalinityAnalyst() ([]*MonitoringWaterqualityOutput, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlSelectWaterqualityMonitoring
	p := []interface{}{144, 143, 208}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	defer rows.Close()

	dataRow := make([]*MonitoringWaterqualityOutput, 0)
	for rows.Next() {
		var (
			station_date     time.Time
			station_name     pqx.JSONRaw
			station_salinity sql.NullFloat64
		)
		rows.Scan(&station_date, &station_name, &station_salinity)
		data := &MonitoringWaterqualityOutput{}
		data.StationDate = station_date.Format("2006-01-02 15:04")
		data.StationName = station_name.JSON()
		data.StationSalinity = ValidData(station_salinity.Valid, station_salinity.Float64)

		dataRow = append(dataRow, data)
	}

	return dataRow, nil
}

//  Parameters:
//		stationID
//			array station id
//  Return:
//		map station id with WaterQualityCompareDatetimeOutputStation
func getStation(stationID []int64) map[int64]*WaterQualityCompareDatetimeOutputStation {

	p := []interface{}{}
	q := "SELECT id, waterquality_station_name FROM m_waterquality_station WHERE "
	for i, v := range stationID {
		if i > 0 {
			q += " OR id=$" + strconv.Itoa(i+1)
		} else {
			q += "id=$" + strconv.Itoa(i+1)
		}
		p = append(p, v)
	}

	db, err := pqx.Open()
	if err != nil {
		return nil
	}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	mapCompare := make(map[int64]*WaterQualityCompareDatetimeOutputStation)

	for rows.Next() {
		var (
			stationID   sql.NullInt64
			stationName pqx.JSONRaw
		)
		dd := &WaterQualityCompareDatetimeOutputStation{}
		rows.Scan(&stationID, &stationName)
		dd.Station = stationName.JSON()
		mapCompare[stationID.Int64] = dd
	}

	return mapCompare
}

//  Parameters:
//		q
//			sql for get data
//		p
//			param for sql
//  Return:
//		WaterQualityWaterlevelOutput
func getDataWaterQualityCompareDatetime(q string, p []interface{}) (*WaterQualityWaterlevelOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dataRow := &WaterQualityWaterlevelOutput{}
	dateOutput := make([]*WaterQualityGraphOutputData, 0)
	for rows.Next() {
		var (
			stationName pqx.JSONRaw
			datetime    time.Time
			value       sql.NullFloat64
		)
		rows.Scan(&stationName, &datetime, &value)

		dd := &WaterQualityGraphOutputData{}
		dataRow.SeriesName = stationName.JSON()
		dd.Name = udt.DatetimeFormat(datetime, "datetime")
		dd.Value = ValidData(value.Valid, value.Float64)
		dateOutput = append(dateOutput, dd)
	}
	dataRow.Data = dateOutput
	return dataRow, nil
}

//  Parameters:
//		inputData
//			WaterQualityGraphWaterlevelAnalystInput
//		q
//			sql for get data
//		wqType
//			for check sql from waterquality or waterlevel
//  Return:
//		WaterQualityWaterlevelOutput
func getDataWaterQuality(inputData *WaterQualityGraphWaterlevelAnalystInput, q string, wqType string) (*WaterQualityWaterlevelOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	p := []interface{}{inputData.DatetimeStart, inputData.DatetimeEnd}
	if wqType != "" {
		p = append(p, inputData.WaterlevelStation)
	} else {
		p = append(p, inputData.WaterQualityStation)
	}
	fmt.Println(q)
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dataRow := &WaterQualityWaterlevelOutput{}
	dateOutput := make([]*WaterQualityGraphOutputData, 0)
	for rows.Next() {
		var (
			stationName pqx.JSONRaw
			datetime    time.Time
			value       sql.NullFloat64
		)
		rows.Scan(&stationName, &datetime, &value)

		dd := &WaterQualityGraphOutputData{}
		dataRow.SeriesName = stationName.JSON()
		dd.Name = udt.DatetimeFormat(datetime, "datetime")
		dd.Value = ValidData(value.Valid, value.Float64)
		dateOutput = append(dateOutput, dd)
	}
	dataRow.Data = dateOutput
	return dataRow, nil
}

//	get range tim by agency
//  Parameters:
//		agencyID
//  Return:
//		range time
func getRangeTime(agencyID int64) int {

	switch agencyID {
	case 14:
		return 30
	case 23:
		return 10
	default:
		return 10
	}
	return 10
}

func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

// func sort data
type DataRange []*WaterQualityGraphCompareAnalystData

func (a DataRange) Len() int      { return len(a) }
func (a DataRange) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DataRange) Less(i, j int) bool {
	t1, _ := time.Parse("2006-01-02 15:04", a[i].Datetime)
	t2, _ := time.Parse("2006-01-02 15:04", a[j].Datetime)
	return t1.Before(t2)
}

//	ข้อมูลคุณภาพน้ำ ล่าสุด
//	Return:
//		array Struct_WaterqualityLatest
func Get_WaterqualityLatest() ([]*Struct_WaterqualityLatest, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	q := `
SELECT mws.id as station_id, 
       mws.waterquality_station_name->>'th' as station_name,
       lg.tumbon_name->>'th' as station_district,
       lg.amphoe_name->>'th'as station_zone,
       lg.province_name->>'th' as station_province,
       '' as station_location,
       w.waterquality_datetime::DATE::text as dt_date,
       w.waterquality_datetime::TIME::text as dt_time,
       w.waterquality_ph as dt_ph, 
       w.waterquality_salinity as dt_salinity, 
       w.waterquality_turbid as dt_turbidity, 
       w.waterquality_conductivity as dt_conductivity, 
       w.waterquality_tds as dt_tds, 
       w.waterquality_chlorophyll as dt_chlorophyll, 
       w.waterquality_do as dt_do, 
       w.waterquality_temp as dt_temp, 
       mws.waterquality_station_lat as station_lat, 
       mws.waterquality_station_long as station_long, 
       a.agency_shortname->>'th' as data_source, 
       lg.province_code as province_id,
       mws.is_active as is_active
FROM   cache.latest_waterquality w 
       INNER JOIN m_waterquality_station mws 
               ON w.waterquality_id = mws.id 
       INNER JOIN lt_geocode lg 
               ON mws.geocode_id = lg.id 
       INNER JOIN agency a 
               ON mws.agency_id = a.id 
WHERE  mws.deleted_at = To_timestamp(0)
       AND mws.is_ignore = 'false' 
       AND mws.agency_id = 23
ORDER BY mws.sort_order ASC
	`
	row, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	rs := make([]*Struct_WaterqualityLatest, 0)
	for row.Next() {
		var (
			_station_id       sql.NullInt64
			_station_name     sql.NullString
			_station_district sql.NullString
			_station_zone     sql.NullString
			_station_province sql.NullString
			_station_location sql.NullString
			_dt_date          string
			_dt_time          string
			_dt_ph            sql.NullFloat64
			_dt_salinity      sql.NullFloat64
			_dt_turbidity     sql.NullFloat64
			_dt_conductivity  sql.NullFloat64
			_dt_tds           sql.NullFloat64
			_dt_chlorophyll   sql.NullFloat64
			_dt_do            sql.NullFloat64
			_dt_temp          sql.NullFloat64
			_station_lat      sql.NullFloat64
			_station_long     sql.NullFloat64
			_data_source      sql.NullString
			_province_id      sql.NullString
			_is_active        sql.NullString //add active status
		)
		err = row.Scan(&_station_id, &_station_name, &_station_zone, &_station_district, &_station_province, &_station_location, &_dt_date, &_dt_time, &_dt_ph, &_dt_salinity,
			&_dt_turbidity, &_dt_conductivity, &_dt_tds, &_dt_chlorophyll, &_dt_do, &_dt_temp, &_station_lat, &_station_long, &_data_source, &_province_id, &_is_active)

		if err != nil {
			return nil, err
		}

		rs = append(rs, &Struct_WaterqualityLatest{
			Station_id:       validdata.DataString(_station_id.Value()),
			Station_name:     validdata.DataString(_station_name.Value()),
			Station_district: validdata.DataString(_station_district.Value()),
			Station_zone:     validdata.DataString(_station_zone.Value()),
			Station_province: validdata.DataString(_station_province.Value()),
			Station_location: validdata.DataString(_station_location.Value()),
			Dt_date:          _dt_date,
			Dt_time:          _dt_time,
			Dt_ph:            validdata.DataString(_dt_ph.Value()),
			Dt_salinity:      validdata.DataString(_dt_salinity.Value()),
			Dt_turbidity:     validdata.DataString(_dt_turbidity.Value()),
			Dt_conductivity:  validdata.DataString(_dt_conductivity.Value()),
			Dt_tds:           validdata.DataString(_dt_tds.Value()),
			Dt_chlorophyll:   validdata.DataString(_dt_chlorophyll.Value()),
			Dt_do:            validdata.DataString(_dt_do.Value()),
			Dt_ec:            validdata.DataString(_dt_conductivity.Value()),
			Dt_temp:          validdata.DataString(_dt_temp.Value()),
			Station_lat:      validdata.DataString(_station_lat.Value()),
			Station_long:     validdata.DataString(_station_long.Value()),
			Data_source:      validdata.DataString(_data_source.Value()),
			Province_id:      validdata.DataString(_province_id.Value()),
			Is_active:        validdata.DataString(_is_active.Value()),
		})
	}

	return rs, nil
}
