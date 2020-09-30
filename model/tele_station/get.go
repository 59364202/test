// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package tele_station is a model for public.tele_station table. This table store tele_station information.
package tele_station

import (
	"database/sql"
	"encoding/json"
	//	"fmt"
	result "haii.or.th/api/thaiwater30/util/result"
	"strconv"
	"strings"

	"haii.or.th/api/util/errors"
	logx "haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
)

//	get tele station
//	Parameters:
//		area_code
//			รหัสภาค
//		province_code
//			รหัสจังหวัด
//		amphoe_code
//			รหัสอำเภอ
//		tumbon_code
//			รหัสตำบล
//	Return:
//		Array Struct_Station
func getTeleStation(area_code, province_code, amphoe_code, tumbon_code, station_type string) ([]*Struct_Station, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSql    string
		param     []interface{}
		SQL_where string = ""

		_result *sql.Rows

		data         []*Struct_Station
		tele_station *Struct_Station

		_id            int
		_name          sql.NullString
		_old_code      sql.NullString
		_lat           sql.NullFloat64
		_long          sql.NullFloat64
		_geocode       sql.NullString
		_province_name sql.NullString
		_province_code sql.NullString
		_agency_id     sql.NullInt64
		_station_type 	sql.NullString
	)
	strSql = SQL_GetTeleStaion

	if area_code != "" {
		param = append(param, area_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.area_code = $" + strconv.Itoa(len(param))
	}
	if province_code != "" {
		param = append(param, province_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.province_code = $" + strconv.Itoa(len(param))
	}
	if amphoe_code != "" {
		param = append(param, amphoe_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.amphoe_code = $" + strconv.Itoa(len(param))
	}
	if tumbon_code != "" {
		param = append(param, tumbon_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.tumbon_code = $" + strconv.Itoa(len(param))
	}
	if station_type != "" {
		if len(param) == 0 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		// ทำให้อยู่ในรูปแบบ $1, $2
		types := strings.Split(station_type, ",")
		tid := ""
		for _, v := range types {
			param = append(param, strings.TrimSpace(v))
			if tid != "" {
				tid += ","
			}
			tid += " $" + strconv.Itoa(len(param))
		}
		SQL_where += " ts.tele_station_type IN (" + tid + ") "
	}

	//	fmt.Println(strSql + SQL_where + SQL_GetTeleStaion_OrderBy)

	_result, err = db.Query(strSql+SQL_where+SQL_GetTeleStaion_OrderBy, param...)
	if err != nil {
		logx.Log(strSql+SQL_where+SQL_GetTeleStaion_OrderBy, param)
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_Station, 0)
	for _result.Next() {
		err = _result.Scan(&_id, &_name, &_old_code, &_lat, &_long, &_geocode, &_province_name, &_province_code, &_agency_id, &_station_type)
		if err != nil {
			return nil, err
		}

		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}
		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}

		tele_station = new(Struct_Station)
		tele_station.Station_id = _id
		tele_station.Station_name = json.RawMessage(_name.String)
		tele_station.Station_old_code = _old_code.String
		tele_station.Station_lat, _ = _lat.Value()
		tele_station.Station_long, _ = _long.Value()
		tele_station.Geocode = _geocode.String
		tele_station.Province_name = json.RawMessage(_province_name.String)
		tele_station.Province_code = _province_code.String
		tele_station.Agency_id = _agency_id.Int64
		tele_station.Station_type = _station_type.String

		data = append(data, tele_station)
	}

	return data, nil
}

//	get all tele station
//	Return:
//		Array Struct_Station
func GetTeleStation() ([]*Struct_Station, error) {
	return getTeleStation("", "", "", "", "")
}

//	get all tele station in province code
//	Parameters:
//		province_code
//			รหัสจังหวัด
//	Return:
//		Array Struct_Station
func GetTeleStationFromProvinceCode(province_code, station_type string) ([]*Struct_Station, error) {
	return getTeleStation("", province_code, "", "", station_type)
}

//	get weir station/watergate
//	Return:
//		Array Struct_WaterlevelCanalStationGroupByProvince
func GetWeirStationGroupByProvince() ([]*Struct_WaterlevelCanalStationGroupByProvince, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strProvinceCode string = "0"

		data       []*Struct_WaterlevelCanalStationGroupByProvince
		objStation *Struct_WaterlevelCanalStationGroupByProvince

		_province_code sql.NullString
		_province_name sql.NullString

		_station_geocode sql.NullString
		_station_id      sql.NullString
		_station_name    sql.NullString
		_station_lat     sql.NullFloat64
		_station_long    sql.NullFloat64
		_station_oldcode sql.NullString

		_result *sql.Rows
	)

	//	log.Println(sqlGetWeirStation)
	_result, err = db.Query(sqlGetWeirStation)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_WaterlevelCanalStationGroupByProvince, 0)
	for _result.Next() {
		err = _result.Scan(&_province_code, &_province_name, &_station_id, &_station_oldcode, &_station_name, &_station_lat, &_station_long, &_station_geocode)
		if err != nil {
			return nil, err
		}

		if strProvinceCode != _province_code.String {
			if strProvinceCode != "0" {
				data = append(data, objStation)
			}

			if !_province_name.Valid || _province_name.String == "" {
				_province_name.String = "{}"
			}

			objStation = &Struct_WaterlevelCanalStationGroupByProvince{}
			objStation.ProvinceCode = ValidData(_province_code.Valid, _province_code.String)
			objStation.ProvinceName = json.RawMessage(_province_name.String)
			objStation.TeleStation = nil

			strProvinceCode = _province_code.String
		}

		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}

		objWaterlevelCanal := &Struct_WaterlevelCanalStation{}
		objWaterlevelCanal.ID = _station_id.String
		objWaterlevelCanal.Name = json.RawMessage(_station_name.String)
		objWaterlevelCanal.OldCode = _station_oldcode.String
		objWaterlevelCanal.Lat = ValidData(_station_lat.Valid, _station_lat.Float64)
		objWaterlevelCanal.Long = ValidData(_station_long.Valid, _station_long.Float64)

		objStation.TeleStation = append(objStation.TeleStation, objWaterlevelCanal)
	}

	if strProvinceCode != "0" {
		data = append(data, objStation)
	}

	return data, nil
}

//	get	tele, canal station กรุ๊ปตามจังหวัด
//	Return:
//		Array Struct_WaterlevelCanalStationGroupByProvince
func GetWaterlevelCanalStationGroupByProvince() ([]*Struct_WaterlevelCanalStationGroupByProvince, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strProvinceCode string = "0"

		data       []*Struct_WaterlevelCanalStationGroupByProvince
		objStation *Struct_WaterlevelCanalStationGroupByProvince

		_province_code sql.NullString
		_province_name sql.NullString

		_station_id      sql.NullString
		_station_name    sql.NullString
		_station_lat     sql.NullFloat64
		_station_long    sql.NullFloat64
		_station_oldcode sql.NullString

		_result *sql.Rows
	)

	//log.Println(sqlGetWaterlevelCanalStation + sqlGetWaterlevelCanalStation_OrderBy)
	_result, err = db.Query(sqlGetWaterlevelCanalStation + sqlGetWaterlevelCanalStation_OrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_WaterlevelCanalStationGroupByProvince, 0)
	for _result.Next() {
		err = _result.Scan(&_province_code, &_province_name, &_station_id, &_station_oldcode, &_station_name, &_station_lat, &_station_long)
		if err != nil {
			return nil, err
		}

		if strProvinceCode != _province_code.String {
			if strProvinceCode != "0" {
				data = append(data, objStation)
			}

			if !_province_name.Valid || _province_name.String == "" {
				_province_name.String = "{}"
			}

			objStation = &Struct_WaterlevelCanalStationGroupByProvince{}
			objStation.ProvinceCode = ValidData(_province_code.Valid, _province_code.String)
			objStation.ProvinceName = json.RawMessage(_province_name.String)
			objStation.TeleStation = nil

			strProvinceCode = _province_code.String
		}

		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}

		objWaterlevelCanal := &Struct_WaterlevelCanalStation{}
		objWaterlevelCanal.ID = _station_id.String
		objWaterlevelCanal.Name = json.RawMessage(_station_name.String)
		objWaterlevelCanal.OldCode = _station_oldcode.String
		objWaterlevelCanal.Lat = ValidData(_station_lat.Valid, _station_lat.Float64)
		objWaterlevelCanal.Long = ValidData(_station_long.Valid, _station_long.Float64)

		objStation.TeleStation = append(objStation.TeleStation, objWaterlevelCanal)
	}

	if strProvinceCode != "0" {
		data = append(data, objStation)
	}

	return data, nil
}

//	get tele, canal station ตามจังหวัด
//	Parameters:
//		provinceCode
//			รหัสจังหวัด
//	Return:
//		Array Struct_WaterlevelCanalStation
func GetWaterlevelCanalStation(provinceCode string) ([]*Struct_WaterlevelCanalStation, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		data               []*Struct_WaterlevelCanalStation
		objWaterlevelCanal *Struct_WaterlevelCanalStation

		_province_code sql.NullString
		_province_name sql.NullString

		_station_id      sql.NullString
		_station_name    sql.NullString
		_station_lat     sql.NullFloat64
		_station_long    sql.NullFloat64
		_station_oldcode sql.NullString

		_result *sql.Rows

		sqlCmdWhere string = ""
		arrParam    []interface{}
	)

	if provinceCode != "" {
		arrParam = append(arrParam, provinceCode)
		sqlCmdWhere = " WHERE geo.province_code = $" + strconv.Itoa(len(arrParam))
	}

	//	log.Println(sqlGetWaterlevelCanalStation+sqlCmdWhere+" ORDER BY station_name->'th' ", arrParam)
	_result, err = db.Query(sqlGetWaterlevelCanalStation+sqlCmdWhere+" ORDER BY station_name->'th' ", arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_WaterlevelCanalStation, 0)
	for _result.Next() {
		err = _result.Scan(&_province_code, &_province_name, &_station_id, &_station_oldcode, &_station_name, &_station_lat, &_station_long)
		if err != nil {
			return nil, err
		}

		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}

		objWaterlevelCanal = &Struct_WaterlevelCanalStation{}
		objWaterlevelCanal.ID = _station_id.String
		objWaterlevelCanal.Name = json.RawMessage(_station_name.String)
		objWaterlevelCanal.OldCode = _station_oldcode.String
		objWaterlevelCanal.Lat = ValidData(_station_lat.Valid, _station_lat.Float64)
		objWaterlevelCanal.Long = ValidData(_station_long.Valid, _station_long.Float64)

		data = append(data, objWaterlevelCanal)
	}

	return data, nil
}

//	get station groupcy province
//	Parameters:
//		param
//			TeleStationParam
//	Return:
//		Array Struct_TeleStationGroupByProvince
func GetTeleStationGroupByProvince(param *TeleStationParam) ([]*Struct_TeleStationGroupByProvince, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}
	//Variables
	var (
		strProvinceCode string = "0"

		data []*Struct_TeleStationGroupByProvince
		obj  *Struct_TeleStationGroupByProvince

		_id            int
		_name          sql.NullString
		_old_code      sql.NullString
		_lat           sql.NullString
		_long          sql.NullString
		_geocode       sql.NullString
		_province_name sql.NullString
		_province_code sql.NullString

		_result *sql.Rows
	)
	//Set 'sqlCmdWhere' variables
	strSqlCmdWhere := ""

	switch param.DataType {
	case "waterlevel":
		strSqlCmdWhere = sqlConditionTeleStationByWaterLevel
	case "rainfall":
		strSqlCmdWhere = sqlConditionTeleStationByRainfall
	default:

	}

	if strSqlCmdWhere != "" {
		strSqlCmdWhere = " WHERE " + strSqlCmdWhere
	}

	_result, err = db.Query(SQL_GetTeleStaion + strSqlCmdWhere + SQL_GetTeleStation_OrderByProvince)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_TeleStationGroupByProvince, 0)
	for _result.Next() {
		err = _result.Scan(&_id, &_name, &_old_code, &_lat, &_long, &_geocode, &_province_name, &_province_code)
		if err != nil {
			return nil, err
		}

		if strProvinceCode != _province_code.String {

			if !_province_name.Valid || _province_name.String == "" {
				_province_name.String = "{}"
			}

			obj = new(Struct_TeleStationGroupByProvince)
			obj.ProvinceCode = _province_code.String
			obj.ProvinceName = json.RawMessage(_province_name.String)
			obj.TeleStation = make([]*TeleStationStruct, 0)

			data = append(data, obj)

			strProvinceCode = _province_code.String
		}
		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}

		teleStruct := new(TeleStationStruct)
		teleStruct.Station_id = _id
		teleStruct.Station_name = json.RawMessage(_name.String)
		teleStruct.Station_old_code = _old_code.String
		teleStruct.Station_lat = _lat.String
		teleStruct.Station_long = _long.String

		obj.TeleStation = append(obj.TeleStation, teleStruct)
	}

	return data, nil
}

//	get tele station by data type
//	Parameters:
//		param
//			TeleStationParam
//	Return:
//		result.Result
func GetTeleStationByDataType(param *TeleStationParam) (*result.Result, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		//data      		[]*TeleStationStruct
		//objTeleStation 	*TeleStationStruct

		data           []*GetTeleStation_OutputParam
		objTeleStation *GetTeleStation_OutputParam

		_id           sql.NullInt64
		_old_code     sql.NullString
		_station_name sql.NullString
		_agency_id    sql.NullInt64
		_geocode_id   sql.NullInt64
		_lat          sql.NullString
		_long         sql.NullString
		_type         sql.NullString
		_subbasin_id  sql.NullInt64

		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_result *sql.Rows
	)

	//Set 'sqlCmdWhere' variables
	arrSqlCmdWhere := []string{}
	strSqlCmdWhere := ""

	switch param.DataType {
	case "waterlevel":
		arrSqlCmdWhere = append(arrSqlCmdWhere, " tele_station_type IN ('W', 'A') ")
	case "rainfall":
		arrSqlCmdWhere = append(arrSqlCmdWhere, " tele_station_type IN ('R', 'A') ")
	default:

	}

	if len(arrSqlCmdWhere) > 0 {
		strSqlCmdWhere = " WHERE " + strings.Join(arrSqlCmdWhere, " AND ")
	}

	//	log.Printf(sqlGetTeleStationByDataType + strSqlCmdWhere + sqlGetTeleStationByDataTypeOrderBy)
	_result, err = db.Query(sqlGetTeleStationByDataType + strSqlCmdWhere + sqlGetTeleStationByDataTypeOrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	//data = make([]*TeleStationStruct, 0)
	data = make([]*GetTeleStation_OutputParam, 0)

	for _result.Next() {
		//Scan to execute query with variables

		err := _result.Scan(&_id, &_old_code, &_station_name, &_agency_id, &_geocode_id, &_lat, &_long, &_type, &_subbasin_id, &_agency_shortname, &_agency_name)
		if err != nil {
			return nil, err
		}

		//Generate TeleStation object
		/*
			objTeleStation = &TeleStationStruct{}
			objTeleStation.Station_id = _id
			objTeleStation.Station_oldcode = _old_code.String
			objTeleStation.Agency_id = _agency_id.Int64
			objTeleStation.Station_lat = _lat.String
			objTeleStation.Station_long = _long.String
			objTeleStation.Station_type = _type.String
			objTeleStation.Geocode_basin = _geocode_basin.String
			objTeleStation.Province_name = json.RawMessage("{}")

			if (_name.String == ""){_name.String = "{}"}
			objTeleStation.Station_name = json.RawMessage(_name.String)

			data = append(data, objTeleStation)
		*/

		objTeleStation = &GetTeleStation_OutputParam{}
		objTeleStation.Id = _id.Int64

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}
		objTeleStation.Station_name = json.RawMessage(_station_name.String)
		objTeleStation.Station_oldcode = _old_code.String
		objTeleStation.Agency_name = json.RawMessage(_agency_name.String)
		objTeleStation.Agency_shortname = json.RawMessage(_agency_shortname.String)

		data = append(data, objTeleStation)
	}

	return result.Result1(data), nil
}

//	get tele station
//	Parameters:
//		id
//			รหัสสถานี
//	Return:
//		Struct_TeleStation
func GetTeleStationById(id int64) (*Struct_TeleStation, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		station *Struct_TeleStation

		_id                   int64
		_subbasin_id          int64
		_agency_id            int64
		_geocode_id           int64
		_tele_station_name    sql.NullString
		_tele_station_lat     sql.NullFloat64
		_tele_station_long    sql.NullFloat64
		_tele_station_oldcode sql.NullString
		_tele_station_type    sql.NullString
		_left_bank            sql.NullFloat64
		_right_bank           sql.NullFloat64
		_ground_level         sql.NullFloat64
	)

	row, err := db.Query(SQL_GetTeleStationById, id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_subbasin_id, &_agency_id, &_geocode_id, &_tele_station_name, &_tele_station_lat, &_tele_station_long, &_tele_station_oldcode, &_tele_station_type,
			&_left_bank, &_right_bank, &_ground_level)
		if err != nil {
			return nil, err
		}

		if !_tele_station_name.Valid || _tele_station_name.String == "" {
			_tele_station_name.String = "{}"
		}

		station = &Struct_TeleStation{Id: _id, Subbasin_id: _subbasin_id, Agency_id: _agency_id, Geocode_id: _geocode_id}
		station.Tele_station_name = json.RawMessage(_tele_station_name.String)
		station.Tele_station_lat = _tele_station_lat.Float64
		station.Tele_station_long = _tele_station_long.Float64
		station.Tele_station_oldcode = _tele_station_oldcode.String
		station.Tele_station_type = _tele_station_type.String
		station.Left_bank = _left_bank.Float64
		station.Right_bank = _right_bank.Float64
		station.Ground_level = ValidData(_ground_level.Valid, _ground_level.Float64)

		if _left_bank.Float64 > _right_bank.Float64 {
			station.Min_bank = ValidData(_right_bank.Valid, _right_bank.Float64)
		} else {
			station.Min_bank = ValidData(_left_bank.Valid, _left_bank.Float64)
		}
	}

	return station, nil
}

func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

//	ข้อมูลสถานีโทรมาตรวัดระดับน้ำ ปตร. หรือฝาย ด้วยรหัสจังหวัด
//	Parameters:
//		province_code
//			รหัสจังหวัด
//	Return:
//		Array Struct_Station
func GetWeirStationFromProvinceCode(province_code string) ([]*Struct_Station, error) {
	return GetWeirStationByCondition("", province_code, "", "")
}

//	ข้อมูลสถานีโทรมาตรวัดระดับน้ำ ปตร. หรือฝาย ด้วยเงื่อนไข รหัสภาค หรือต. อ. จ.
//	Parameters:
//		area_code
//			รหัสภาค
//		province_code
//			รหัสจังหวัด
//		amphoe_code
//			รหัสอำเภอ
//		tumbon_code
//			รหัสตำบล
//	Return:
//		Array Struct_Station
func GetWeirStationByCondition(area_code, province_code, amphoe_code, tumbon_code string) ([]*Struct_Station, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSql    string
		param     []interface{}
		SQL_where string = ""

		_result *sql.Rows

		data         []*Struct_Station
		tele_station *Struct_Station

		_id            int
		_name          sql.NullString
		_old_code      sql.NullString
		_lat           sql.NullFloat64
		_long          sql.NullFloat64
		_geocode       sql.NullString
		_province_name sql.NullString
		_province_code sql.NullString
		_agency_id     sql.NullInt64
	)
	strSql = SQL_GetWeirStation

	if area_code != "" {
		param = append(param, area_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.area_code = $" + strconv.Itoa(len(param))
	}
	if province_code != "" {
		param = append(param, province_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.province_code = $" + strconv.Itoa(len(param))
	}
	if amphoe_code != "" {
		param = append(param, amphoe_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.amphoe_code = $" + strconv.Itoa(len(param))
	}
	if tumbon_code != "" {
		param = append(param, tumbon_code)
		if len(param) == 1 {
			SQL_where += " WHERE "
		} else {
			SQL_where += " AND "
		}
		SQL_where += " lg.tumbon_code = $" + strconv.Itoa(len(param))
	}

	_result, err = db.Query(strSql+SQL_where, param...)
	if err != nil {
		logx.Log(strSql+SQL_where+SQL_GetTeleStaion_OrderBy, param)
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_Station, 0)
	for _result.Next() {
		err = _result.Scan(&_id, &_name, &_old_code, &_lat, &_long, &_geocode, &_province_name, &_province_code, &_agency_id)
		if err != nil {
			return nil, err
		}

		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}
		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}

		tele_station = new(Struct_Station)
		tele_station.Station_id = _id
		tele_station.Station_name = json.RawMessage(_name.String)
		tele_station.Station_old_code = _old_code.String
		tele_station.Station_lat, _ = _lat.Value()
		tele_station.Station_long, _ = _long.Value()
		tele_station.Geocode = _geocode.String
		tele_station.Province_name = json.RawMessage(_province_name.String)
		tele_station.Province_code = _province_code.String
		tele_station.Agency_id = _agency_id.Int64

		data = append(data, tele_station)
	}

	return data, nil
}
