// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package canal_station is a model for public.m_canal_station table. This table store m_canal_station information.
package canal_station

import (
	"database/sql"
	"encoding/json"
	"strconv"

	//	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	get canal_station from canal_station id
//	Parameters:
//		id
//			canal station id
//	Return:
//		[]Struct_CanalStation
func GetCanalStationById(id int64) ([]*Struct_CanalStation, error) {
	return getCanalStation(id, "", "", "", "")
}

//	get canal_station
//	Parameters:
//		station_id
//			canal station id
//		area_code
//			area code
//		province_code
//			province code
//		amphoe_code
//			amphoe code
//		tumbon_code
//			tumbon code
//	Return:
//		[]Struct_CanalStation
func getCanalStation(station_id int64, area_code, province_code, amphoe_code, tumbon_code string) ([]*Struct_CanalStation, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSql    string
		param     []interface{}
		SQL_where string = ""

		_result *sql.Rows

		data          []*Struct_CanalStation
		canal_station *Struct_CanalStation

		_id            int64
		_name          sql.NullString
		_lat           sql.NullFloat64
		_long          sql.NullFloat64
		_geocode       sql.NullString
		_province_name sql.NullString
		_province_code sql.NullString
	)
	strSql = SQL_GetCanalStaion

	if station_id != 0 {
		param = append(param, station_id)
		SQL_where += " mcs.id = $" + strconv.Itoa(len(param)) + " AND"
	}
	if area_code != "" {
		param = append(param, area_code)
		SQL_where += " lg.area_code = $" + strconv.Itoa(len(param)) + " AND"
	}
	if province_code != "" {
		param = append(param, province_code)
		SQL_where += " lg.province_code = $" + strconv.Itoa(len(param)) + " AND"
	}
	if amphoe_code != "" {
		param = append(param, amphoe_code)
		SQL_where += " lg.amphoe_code = $" + strconv.Itoa(len(param)) + " AND"
	}
	if tumbon_code != "" {
		param = append(param, tumbon_code)
		SQL_where += " lg.tumbon_code = $" + strconv.Itoa(len(param)) + " AND"
	}

	// query
	if len(param) != 0 {
		strSql += " WHERE " + SQL_where[0:len(SQL_where)-3]
		_result, err = db.Query(strSql+SQL_GetCanalStaion_OrderBy, param...)
	} else {
		_result, err = db.Query(strSql + SQL_GetCanalStaion_OrderBy)
	}
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_CanalStation, 0)
	for _result.Next() {
		err = _result.Scan(&_id, &_name, &_lat, &_long, &_geocode, &_province_name, &_province_code)
		if err != nil {
			return nil, err
		}

		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}
		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}

		canal_station = new(Struct_CanalStation)
		canal_station.Id = _id
		canal_station.Station_Name = json.RawMessage(_name.String)
		canal_station.Canal_station_name = json.RawMessage(_name.String)
		canal_station.Canal_station_lat, _ = _lat.Value()
		canal_station.Canal_station_long, _ = _long.Value()
		canal_station.Geocode = _geocode.String
		canal_station.Province_name = json.RawMessage(_province_name.String)
		canal_station.Province_code = _province_code.String

		data = append(data, canal_station)
	}

	return data, nil
}

//	get all canal_station
//	Return:
//		[]Struct_CanalStation
func GetCanalStation() ([]*Struct_CanalStation, error) {
	return getCanalStation(0, "", "", "", "")
}

//	get canal_station from province code
//	Parameters:
//		province_code
//			province code
//	Return:
//		[]Struct_CanalStation
func GetCanalStationFromProvinceCode(province_code string) ([]*Struct_CanalStation, error) {
	return getCanalStation(0, "", province_code, "", "")
}

//	get canal_station เป็นกรุ๊ปๆแยกตาม province
//	Return:
//		[]Struct_CanalStationGroupByProvince
func GetCanalStationGroupByProvince() ([]*Struct_CanalStationGroupByProvince, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strProvinceCode string = "0"

		data []*Struct_CanalStationGroupByProvince
		obj  *Struct_CanalStationGroupByProvince

		_id            int64
		_name          sql.NullString
		_lat           sql.NullFloat64
		_long          sql.NullFloat64
		_geocode       sql.NullString
		_province_name sql.NullString
		_province_code sql.NullString
	)
	// query
	_row, err := db.Query(SQL_GetCanalStaion + SQL_GetCanalStaion_OrderByProvince)
	if err != nil {
		return nil, err
	}

	data = make([]*Struct_CanalStationGroupByProvince, 0)
	for _row.Next() {
		err = _row.Scan(&_id, &_name, &_lat, &_long, &_geocode, &_province_name, &_province_code)
		if err != nil {
			return nil, err
		}

		if strProvinceCode != _province_code.String {
			if !_province_name.Valid || _province_name.String == "" {
				_province_name.String = "{}"
			}

			obj = new(Struct_CanalStationGroupByProvince)
			obj.ProvinceCode = _province_code.String
			obj.ProvinceName = json.RawMessage(_province_name.String)
			obj.Station = make([]*Struct_CanalStation, 0)

			data = append(data, obj)

			strProvinceCode = _province_code.String
		}

		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}
		station := new(Struct_CanalStation)
		station.Id = _id
		station.Canal_station_name = json.RawMessage(_name.String)
		station.Canal_station_lat, _ = _lat.Value()
		station.Canal_station_long, _ = _long.Value()

		obj.Station = append(obj.Station, station)
	}

	return data, err
}

//
//func GetCanalStationForCheckMetadata(param *Struct_CanalStation_InputParam) ([]*Struct_CanalStation_ForCheckMetadata, error) {
//
//	err := checkInputParam(param)
//	if err != nil {
//		return nil, errors.Repack(err)
//	}
//
//	//Convert AgencyID type from string to int64
//	intAgencyID, err := strconv.ParseInt(param.AgencyID, 10, 64)
//	if err != nil {
//		return nil, errors.Repack(err)
//	}
//
//	//Open Database
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, errors.Repack(err)
//	}
//
//	//Variables
//	var (
//		data            []*Struct_CanalStation_ForCheckMetadata
//		objCanalStation *Struct_CanalStation_ForCheckMetadata
//
//		_id               sql.NullInt64
//		_station_oldcode  sql.NullString
//		_station_name     sql.NullString
//		_agency_name      sql.NullString
//		_agency_shortname sql.NullString
//		_lat              sql.NullFloat64
//		_long             sql.NullFloat64
//		_geocode          sql.NullString
//		_subbasin_name    sql.NullString
//
//		_result *sql.Rows
//	)
//
//	//Where Condition
//	var sqlCmdWhere string = " WHERE a." + param.ColumnName + " IS NULL AND a.agency_id = $1 "
//
//	//Query
//	log.Printf(sqlGetCanalStation+sqlCmdWhere+sqlGetCanalStationOrderBy, intAgencyID)
//	_result, err = db.Query(sqlGetCanalStation+sqlCmdWhere+sqlGetCanalStationOrderBy, intAgencyID)
//
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//	defer _result.Close()
//
//	// Loop data result
//	data = make([]*Struct_CanalStation_ForCheckMetadata, 0)
//
//	for _result.Next() {
//		//Scan to execute query with variables
//		err := _result.Scan(&_id, &_station_oldcode, &_station_name, &_agency_name, &_agency_shortname, &_lat, &_long, &_geocode, &_subbasin_name)
//		if err != nil {
//			return nil, err
//		}
//
//		if _station_name.String == "" || !_station_name.Valid {
//			_station_name.String = "{}"
//		}
//		if _agency_shortname.String == "" || !_agency_shortname.Valid {
//			_agency_shortname.String = "{}"
//		}
//		if _agency_name.String == "" || !_agency_name.Valid {
//			_agency_name.String = "{}"
//		}
//		if _subbasin_name.String == "" || !_subbasin_name.Valid {
//			_subbasin_name.String = "{}"
//		}
//
//		//Generate AirStation object
//		objCanalStation = &Struct_CanalStation_ForCheckMetadata{}
//		objCanalStation.Id = _id.Int64
//		objCanalStation.OldCode = _station_oldcode.String
//		objCanalStation.Name = json.RawMessage(_station_name.String)
//		objCanalStation.Lat = _lat.Float64
//		objCanalStation.Long = _long.Float64
//		objCanalStation.Geocode = _geocode.String
//		objCanalStation.SubbasinName = json.RawMessage(_subbasin_name.String)
//
//		objCanalStation.AgencyName = json.RawMessage(_agency_name.String)
//		objCanalStation.AgencyShortname = json.RawMessage(_agency_shortname.String)
//
//		data = append(data, objCanalStation)
//	}
//
//	return data, nil
//}
//
//func checkInputParam(param *Struct_CanalStation_InputParam) error {
//
//	//Check Parameters
//	if param.AgencyID == "" {
//		return errors.New("'agency_id' is not null.")
//	}
//	if param.ColumnName == "" {
//		return errors.New("'column_name' is not null.")
//	}
//
//	return nil
//}
