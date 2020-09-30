// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

// model for public.
package humid

import (
	"database/sql"
	"time"
	//	"fmt"
	"encoding/json"
	"strconv"
	"strings"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

// get  graph by station and date
//  Parameters:
//    param
//      Struct_GraphParam
//  Return:
//    Struct_Graph
func GetGraph(param *Struct_GraphParam) (*Struct_Graph, error) {

	// check station id
	if param.Station_id == "" {
		return nil, rest.NewError(422, "No station id", nil)
	}
	data := &Struct_Graph{}
	// get data
	graphData, err := getGraphData(param.Station_id, param.Start_date, param.End_date)
	if err != nil {
		return nil, err
	}
	// add data
	data.GraphData = graphData
	return data, nil
}

// get  graph by station id and date range
//  Parameters:
//    stationID
//      tele station id
//    startDate
//      date start
//    endDate
//      date end
//  Return:
//    Array Struct_GraphByStationAndDate
func getGraphData(stationID, startDate, endDate string) ([]*Struct_GraphByStationAndDate, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlGraph
	p := []interface{}{stationID, startDate, endDate}

	//	fmt.Println(q)
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var (
		id       sql.NullInt64
		datetime time.Time
		value    sql.NullFloat64
	)

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02 15:04"
	}

	graphData := make([]*Struct_GraphByStationAndDate, 0)
	for rows.Next() {
		err := rows.Scan(&id, &datetime, &value)
		if err != nil {
			return nil, err
		}
		dataRow := &Struct_GraphByStationAndDate{}
		dataRow.Datetime = datetime.Format(strDatetimeFormat)
		dataRow.Value = ValidData(value.Valid, value.Float64)
		graphData = append(graphData, dataRow)
	}

	return graphData, nil
}

//	get humid
//	Parameters:
//		p
//			Param_Provinces
//	Return:
//		Array Struct_Humid
func GetHumid(p *Param_Provinces) ([]*Struct_Humid, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_Humid{}
	itf := []interface{}{}

	q := sqlHumid

	//Check Filter province_id
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		q += " AND "
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(p.Province_Code, " "))
			q += " province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			q += " province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	if p.Region_Code != "" {
		q += " AND "
		q += " area_code like '" + p.Region_Code + "' "
	}

	if p.Region_Code_tmd != "" {
		q += " AND "
		q += " tmd_area_code like '" + p.Region_Code_tmd + "' "
	}

	if p.Order != "" {
		q += " ORDER BY humid_value " + p.Order
	}

	if p.Data_Limit == 0 {
		p.Data_Limit = 1
	}

	q += " LIMIT " + strconv.Itoa(p.Data_Limit)

	//	fmt.Println(q)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}
	strDatetimeFormat := "2006-01-02 15:04"

	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_name                 pqx.JSONRaw
			_amphoe_code          sql.NullString
			_amphoe_name          pqx.JSONRaw
			_tambon_code          sql.NullString
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            sql.NullString
			_agency_id            sql.NullInt64
			_agency_name          pqx.JSONRaw
			_humid_datetime       sql.NullString
			_humid_value          sql.NullFloat64

			d       *Struct_Humid                 = &Struct_Humid{}
			station *Struct_TeleStation           = &Struct_TeleStation{}
			agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_humid_datetime, &_humid_value, &_tele_station_id, &_tele_station_oldcode, &_name, &_tele_station_lat, &_tele_station_long, &_area_code, &_area_name, &_province_code, &_amphoe_code, &_tambon_code, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name)

		if err != nil {
			return nil, err
		}

		d.Humid = ValidData(_humid_value.Valid, _humid_value.Float64)
		if _humid_datetime.Valid {
			d.HumidDatetime = pqx.NullStringToTime(_humid_datetime).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _name.JSON()
		station.OldCode = _tele_station_oldcode.String

		agency.Agency_name = _agency_name.JSON()

		geocode.Tumbon_code = _tambon_code.String
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_code = _amphoe_code.String
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code.String
		geocode.Province_name = _province_name.JSON()
		geocode.Area_code = _area_code.String
		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		geocode.Area_name = json.RawMessage(_area_name.String)

		d.Station = station
		d.Agency = agency
		d.Geocode = geocode

		r = append(r, d)
	}

	return r, nil
}

//ตรวจสอบค่าใน column float กรณีค่าเป็น null ให้ return type เป็น interface แทน flat64
//ตัวแปรที่รับค่าจาก db เป็น float64 ซึ่ง float64 เป็น null ไม่ได้ ถ้าใช้  column.float64 จะได้เป็น 0 ทั้งๆที่ใน db เป็น null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
