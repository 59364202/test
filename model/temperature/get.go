// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
// model for public.temperature
package temperature

//ฤดูหนาว 20 ตค. - 20 กพ.
//ฤดูหนาว
//เริ่มประมาณกลางเดือนตุลาคมถึงประมาณกลางเดือนกุมภาพันธ์
//ฤดูร้อน
//เริ่มประมาณกลางเดือนกุมภาพันธ์ถึงประมาณกลางเดือนพฤษภาคม
//https://www.tmd.go.th/info/info.php?FileID=23

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"

	util_datetime "haii.or.th/api/thaiwater30/util/datetime"
)

//	get Max Min Temperature in current week from  public.temperature
//	query ที่ต้องมีใส่เงื่อนไข qc_status ที่เป็น field json ไม่ว่าใส่ index หรือไม่ใส่ query time จะเพิ่มไปอีก 1 เท่า
//	น่าจะต้องแยก qc status กับ detail ออกจากกันใส่ index ถึงจะ query ได้เร็ว
//	Parameters:
//		p
//			Param_TemperatureProvinces
//	Return:
//		Array Struct_MaxMinTemperature
func GetMaxMinTemperatureThisWeek(p *Param_TemperatureProvinces) ([]*Struct_MaxMinTemperature, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_MaxMinTemperature{}
	itf := []interface{}{}

	tn := time.Now().UTC()
	year, week := tn.ISOWeek()
	date := util_datetime.FirstDayOfISOWeek(year, week, time.UTC)
	start_date := date.Format("2006-01-02")

	q := `WITH maxmin_query AS (
	SELECT tele_station_id,max(temp_value) as max,min(temp_value) as min
	FROM temperature
	WHERE qc_status ->> 'is_pass'::text like 'true' AND temp_value < 999999 AND temp_datetime >= '`
	q += start_date + "' "

	q += ` GROUP BY tele_station_id  
) SELECT tele_station_id,'`
	q += start_date + "'::date as start_date,"

	q += `max,min,
	tele_station_oldcode,tele_station_lat,tele_station_long,
	area_code,province_code,area_name,province_name,amphoe_name,tumbon_name,
	agency_id,agency_name,tele_station_name
FROM maxmin_query 
INNER JOIN m_tele_station m  ON maxmin_query.tele_station_id = m.id
INNER JOIN lt_geocode lg  ON m.geocode_id = lg.id  AND m.is_ignore = 'false' AND lg.geocode <> '999999'
INNER JOIN agency a ON m.agency_id = a.id`

	//Check Filter province_id
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		if len(itf) == 0 {
			q += " WHERE "
		} else {
			q += " AND "
		}
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
		if len(itf) == 0 {
			q += " WHERE "
		} else {
			q += " AND "
		}
		q += " area_code like '" + p.Region_Code + "' "
	}

	q += ` ORDER BY max desc,min `

	if p.Data_Limit == 0 {
		p.Data_Limit = 1
	}

	q += "LIMIT " + strconv.Itoa(p.Data_Limit)

	//	fmt.Println(q)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}
	strDatetimeFormat := "2006-01-02"

	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_name                 pqx.JSONRaw
			_amphoe_name          pqx.JSONRaw
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            sql.NullString
			_agency_id            sql.NullInt64
			_agency_name          pqx.JSONRaw
			_max                  sql.NullFloat64
			_min                  sql.NullFloat64
			_start_date           sql.NullString

			d       *Struct_MaxMinTemperature     = &Struct_MaxMinTemperature{}
			station *Struct_TeleStation           = &Struct_TeleStation{}
			agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_tele_station_id, &_start_date, &_max, &_min, &_tele_station_oldcode, &_tele_station_lat, &_tele_station_long, &_area_code, &_province_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name, &_name)

		if err != nil {
			return nil, err
		}

		d.Max_temperature = fmt.Sprintf("%.2f", _max.Float64)
		d.Min_temperature = fmt.Sprintf("%.2f", _min.Float64)
		if _start_date.Valid {
			d.Start_date = pqx.NullStringToTime(_start_date).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _name.JSON()
		station.OldCode = _tele_station_oldcode.String

		agency.Agency_name = _agency_name.JSON()
		geocode.Tumbon_name = _tumbon_name.JSON()
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

//	get latest temperature max,min by region from latest.temperature for thaiwater main page
//  latest.v_main_temperature_maxmin_by_tmd_region
//	Return:
//		Array Struct_TemperatureMaxMinByRegion
func GetTemperatureMaxMinByRegion() ([]*Struct_TemperatureMaxMinByRegion, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_TemperatureMaxMinByRegion{}
	q := sqlTemperatureMaxMinByRegion

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02 15:04"
	}

	row, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_name    pqx.JSONRaw
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_amphoe_code          sql.NullString
			_amphoe_name          pqx.JSONRaw
			_tumbon_code          sql.NullString
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            pqx.JSONRaw
			_agency_id            sql.NullInt64
			_agency_name          pqx.JSONRaw
			_temperature_datetime sql.NullString
			_temperature          sql.NullFloat64
			_humid                sql.NullFloat64
			_pressure             sql.NullFloat64
			_rank                 int64

			d       *Struct_TemperatureMaxMinByRegion = &Struct_TemperatureMaxMinByRegion{}
			station *Struct_TeleStation               = &Struct_TeleStation{}
			agency  *model_agency.Struct_Agency       = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode     = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_humid, &_pressure, &_temperature_datetime, &_temperature, &_tele_station_id, &_tele_station_oldcode, &_tele_station_name,
			&_tele_station_lat, &_tele_station_long, &_area_code, &_province_code, &_amphoe_code, &_tumbon_code,
			&_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_rank, &_agency_name)
		if err != nil {
			return nil, err
		}

		d.Temperature = ValidData(_temperature.Valid, _temperature.Float64)
		d.Humid = ValidData(_humid.Valid, _humid.Float64)
		d.Pressure = ValidData(_pressure.Valid, _pressure.Float64)
		if _temperature_datetime.Valid {
			d.Time = pqx.NullStringToTime(_temperature_datetime).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _tele_station_name.JSON()
		station.OldCode = _tele_station_oldcode.String

		agency.Id = _agency_id.Int64
		agency.Agency_name = _agency_name.JSON()

		geocode.Tumbon_code = _tumbon_code.String
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_code = _amphoe_code.String
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code.String
		geocode.Province_name = _province_name.JSON()
		geocode.Area_code = _area_code.String
		geocode.Area_name = _area_name.JSON()

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

// get temperature graph by station and date
//  Parameters:
//    param
//      Struct_TemperatureGraphParam
//  Return:
//    Struct_TemperatureGraph
func GetTemperatureGraph(param *Struct_TemperatureGraphParam) (*Struct_TemperatureGraph, error) {

	// check station id
	if param.Station_id == "" {
		return nil, rest.NewError(422, "No station id", nil)
	}
	data := &Struct_TemperatureGraph{}
	// get data
	graphData, err := getTemperatureGraphData(param.Station_id, param.Start_date, param.End_date)
	if err != nil {
		return nil, err
	}
	// add data
	data.GraphData = graphData
	return data, nil
}

// get temperature graph by station id and date range
//  Parameters:
//    stationID
//      tele station id
//    startDate
//      date start
//    endDate
//      date end
//  Return:
//    Array Struct_GraphByStationAndDate
func getTemperatureGraphData(stationID, startDate, endDate string) ([]*Struct_GraphByStationAndDate, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlTemperatureGraph
	p := []interface{}{stationID, startDate, endDate}

	fmt.Println(q)
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

//	get latest temperature from latest.temperature
//	Parameters:
//		p
//			Param_TemperatureProvinces
//	Return:
//		Struct_TemperatureLatest
func GetTemperatureLatest(p *Param_TemperatureProvinces) (*Struct_TemperatureLatest, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02 15:04"
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := &Struct_TemperatureLatest{}
	itf := []interface{}{}

	m := time.Now().Month()
	fmt.Println(m, int(m))

	month := int(time.Now().Month())
	orderBy := ""
	seasonMaxMin := ""
	if (month >= 11) && (month <= 3) {
		orderBy = "ASC"
		seasonMaxMin = "อุณหภูมิต่ำสุด"
	} else {
		orderBy = "DESC"
		seasonMaxMin = "อุณหภูมิสูงสุด"
	}
	r.SeasonMaxMin = seasonMaxMin

	q := sqlTemperature

	//-- Check Filter by parameters --//
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

	q += " ORDER BY temp_datetime DESC, temp_value " + orderBy

	if p.Data_Limit != 0 {
		q += " LIMIT " + strconv.Itoa(p.Data_Limit)
	}

	fmt.Println(q)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_name    pqx.JSONRaw
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_amphoe_code          sql.NullString
			_amphoe_name          pqx.JSONRaw
			_tumbon_code          sql.NullString
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            pqx.JSONRaw
			_agency_id            sql.NullInt64
			_agency_name          pqx.JSONRaw
			_data_id              sql.NullInt64
			_temperature          sql.NullFloat64
			_temperature_datetime sql.NullString
			_basin_id             sql.NullInt64
			_subbasin_id          sql.NullInt64
			_tmd_area_code        sql.NullString
			_tmd_area_name        pqx.JSONRaw

			d       *Struct_Temperature           = &Struct_Temperature{}
			station *Struct_TeleStation           = &Struct_TeleStation{}
			agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_temperature_datetime, &_temperature, &_data_id, &_tele_station_id,
			&_tele_station_oldcode, &_tele_station_name,
			&_tele_station_lat, &_tele_station_long, &_area_code, &_province_code, &_amphoe_code, &_tumbon_code,
			&_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name, &_basin_id,
			&_subbasin_id, &_tmd_area_code, &_tmd_area_name)
		if err != nil {
			return nil, err
		}

		d.Id = _data_id.Int64
		d.Temperature = ValidData(_temperature.Valid, _temperature.Float64)
		if _temperature_datetime.Valid {
			d.Time = pqx.NullStringToTime(_temperature_datetime).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _tele_station_name.JSON()
		station.OldCode = _tele_station_oldcode.String
		station.Basin_id = _basin_id.Int64
		station.SubBasin_id = _subbasin_id.Int64

		agency.Id = _agency_id.Int64
		agency.Agency_name = _agency_name.JSON()

		geocode.Tumbon_code = _tumbon_code.String
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_code = _amphoe_code.String
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code.String
		geocode.Province_name = _province_name.JSON()
		geocode.Area_code = _area_code.String
		geocode.Area_name = _area_name.JSON()

		d.Station = station
		d.Agency = agency
		d.Geocode = geocode

		r.Data = append(r.Data, d)
	}

	return r, nil
}

//	get max,min temperature
//	Parameters:
//		p
//			Param_TemperatureProvinces
//	Return:
//		Array Struct_TemperatureMaxMin
func GetMaxMinTemperature(p *Param_TemperatureProvinces) ([]*Struct_TemperatureMaxMin, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_TemperatureMaxMin{}
	itf := []interface{}{}

	//Check Filter province_id
	province_condition := ""
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		province_condition += " AND "
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(p.Province_Code, " "))
			province_condition += " province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			province_condition += " province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	region_condition := ""
	if p.Region_Code != "" {
		region_condition += " AND "
		region_condition += " area_code like '" + p.Region_Code + "' "
	}

	if p.Region_Code_tmd != "" {
		region_condition += " AND "
		region_condition += " tmd_area_code like '" + p.Region_Code_tmd + "' "
	}

	q := ` WITH max_temperature AS (
select  
	row_number() OVER (ORDER BY  temp_value DESC) as number,
	dd.temp_datetime,
	temp_value as max_temp,
	d.id AS station_id,
	d.tele_station_oldcode,
	d.tele_station_name,
	d.tele_station_lat,
	d.tele_station_long,
	g.tmd_area_code AS area_code,
	g.tmd_area_name AS area_name,
	g.province_code,
	g.amphoe_code,
	g.tumbon_code,
	g.province_name,
	g.amphoe_name,
	g.tumbon_name,
	agency_name
from  latest.temperature  dd
  LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id
  LEFT JOIN agency agt ON d.agency_id = agt.id
  LEFT JOIN lt_geocode g ON d.geocode_id = g.id
	WHERE temp_datetime :: date = CURRENT_DATE  AND d.geocode_id is not null AND ((d.is_ignore = false) AND ((dd.qc_status IS NULL) OR (dd.qc_status ->> 'is_pass' :: text) = 'true' :: text)` + province_condition + region_condition + `) 
    ORDER BY  temp_value DESC LIMIT  10
),
min_temperature AS (
select  
	row_number() OVER (ORDER BY  temp_value) as number,
	dd.temp_datetime as min_temp_datetime,
	temp_value as min_temp,
	d.id as min_station_id,
	d.tele_station_oldcode as min_tele_station_oldcode,
	d.tele_station_name as min_tele_station_name,
	d.tele_station_lat as min_tele_station_lat,
	d.tele_station_long as min_tele_station_long,
	g.tmd_area_code AS min_area_code,
	g.tmd_area_name AS min_area_name,
	g.province_code as min_province_code,
	g.amphoe_code as min_amphoe_code,
	g.tumbon_code as min_tumbon_code,
	g.province_name as min_province_name,
	g.amphoe_name as min_amphoe_name,
	g.tumbon_name as min_tumbon_name,
	agency_name as  min_agency_name
from  latest.temperature  dd
  LEFT JOIN m_tele_station d ON ((dd.tele_station_id = d.id))
  LEFT JOIN agency agt ON ((d.agency_id = agt.id))
  LEFT JOIN lt_geocode g ON ((d.geocode_id = g.id))
  WHERE temp_datetime :: date = CURRENT_DATE AND d.geocode_id is not null  AND ((d.is_ignore = false) AND ((dd.qc_status IS NULL) OR (dd.qc_status ->> 'is_pass' :: text) = 'true' :: text)` + province_condition + region_condition + `)
    ORDER BY  temp_value LIMIT  10
)
select * from max_temperature max LEFT JOIN min_temperature min ON max.number = min.number`

	//	fmt.Println(q)
	row, err := db.Query(q)
	if err != nil {
		fmt.Println("query ", err)
		return nil, err
	}
	strDatetimeFormat := "2006-01-02 15:04"

	for row.Next() {
		var (
			station_id           int64
			tele_station_lat     sql.NullFloat64
			tele_station_long    sql.NullFloat64
			tele_station_oldcode sql.NullString
			tele_station_name    pqx.JSONRaw
			amphoe_code          sql.NullString
			amphoe_name          pqx.JSONRaw
			tumbon_code          sql.NullString
			tumbon_name          pqx.JSONRaw
			province_code        sql.NullString
			province_name        pqx.JSONRaw
			area_code            sql.NullString
			area_name            sql.NullString
			agency_name          pqx.JSONRaw
			temp_datetime        sql.NullString
			max_temp             sql.NullFloat64

			min_temp_datetime        sql.NullString
			min_temp                 sql.NullFloat64
			min_station_id           int64
			min_tele_station_oldcode sql.NullString
			min_tele_station_name    pqx.JSONRaw
			min_tele_station_lat     sql.NullFloat64
			min_tele_station_long    sql.NullFloat64
			min_area_code            sql.NullString
			min_area_name            sql.NullString
			min_province_code        sql.NullString
			min_amphoe_code          sql.NullString
			min_tumbon_code          sql.NullString
			min_province_name        pqx.JSONRaw
			min_amphoe_name          pqx.JSONRaw
			min_tumbon_name          pqx.JSONRaw
			min_agency_name          pqx.JSONRaw

			d           *Struct_TemperatureMaxMin     = &Struct_TemperatureMaxMin{}
			t_max       *Struct_Temperature           = &Struct_Temperature{}
			t_min       *Struct_Temperature           = &Struct_Temperature{}
			station     *Struct_TeleStation           = &Struct_TeleStation{}
			agency      *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode     *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
			min_station *Struct_TeleStation           = &Struct_TeleStation{}
			min_agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			min_geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}

			max_row_number int64
			min_row_number int64
		)

		err = row.Scan(&max_row_number, &temp_datetime, &max_temp, &station_id, &tele_station_oldcode, &tele_station_name, &tele_station_lat,
			&tele_station_long, &area_code, &area_name, &province_code, &amphoe_code, &tumbon_code, &province_name, &amphoe_name,
			&tumbon_name, &agency_name, &min_row_number, &min_temp_datetime, &min_temp, &min_station_id, &min_tele_station_oldcode, &min_tele_station_name,
			&min_tele_station_lat, &min_tele_station_long, &min_area_code, &min_area_name, &min_province_code, &min_amphoe_code,
			&min_tumbon_code, &min_province_name, &min_amphoe_name, &min_tumbon_name, &min_agency_name)

		if err != nil {
			fmt.Println("scan error", err)
			return nil, err
		}

		// max temperature
		t_max.Temperature = ValidData(max_temp.Valid, max_temp.Float64)
		if temp_datetime.Valid {
			t_max.Time = pqx.NullStringToTime(temp_datetime).Format(strDatetimeFormat)
		}

		station.Id = station_id
		station.Lat = ValidData(tele_station_lat.Valid, tele_station_lat.Float64)
		station.Long = ValidData(tele_station_long.Valid, tele_station_long.Float64)
		station.Name = tele_station_name.JSON()
		station.OldCode = tele_station_oldcode.String

		agency.Agency_name = agency_name.JSON()

		geocode.Tumbon_code = tumbon_code.String
		geocode.Tumbon_name = tumbon_name.JSON()
		geocode.Amphoe_code = amphoe_code.String
		geocode.Amphoe_name = amphoe_name.JSON()
		geocode.Province_code = province_code.String
		geocode.Province_name = province_name.JSON()
		geocode.Area_code = area_code.String
		if area_name.String == "" {
			area_name.String = "{}"
		}
		geocode.Area_name = json.RawMessage(area_name.String)

		t_max.Station = station
		t_max.Agency = agency
		t_max.Geocode = geocode

		// min temperature
		t_min.Temperature = ValidData(min_temp.Valid, min_temp.Float64)
		if min_temp_datetime.Valid {
			t_min.Time = pqx.NullStringToTime(min_temp_datetime).Format(strDatetimeFormat)
		}

		min_station.Id = min_station_id
		min_station.Lat = ValidData(min_tele_station_lat.Valid, min_tele_station_lat.Float64)
		min_station.Long = ValidData(min_tele_station_long.Valid, min_tele_station_long.Float64)
		min_station.Name = min_tele_station_name.JSON()
		min_station.OldCode = min_tele_station_oldcode.String

		min_agency.Agency_name = min_agency_name.JSON()

		min_geocode.Tumbon_code = min_tumbon_code.String
		min_geocode.Tumbon_name = min_tumbon_name.JSON()
		min_geocode.Amphoe_code = min_amphoe_code.String
		min_geocode.Amphoe_name = min_amphoe_name.JSON()
		min_geocode.Province_code = min_province_code.String
		min_geocode.Province_name = min_province_name.JSON()
		min_geocode.Area_code = min_area_code.String
		if min_area_name.String == "" {
			min_area_name.String = "{}"
		}
		min_geocode.Area_name = json.RawMessage(min_area_name.String)

		t_min.Station = min_station
		t_min.Agency = min_agency
		t_min.Geocode = min_geocode

		d.Max = t_max
		d.Min = t_min
		r = append(r, d)
	}

	return r, nil

}
