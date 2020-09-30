// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//
package station

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	//	"log"
	"strconv"
	"strings"
)

//	column สำหรับ check metadata
//	Parameters:
//		tableName
//			ชื่อตาราง
//	Return:
//		Array string
func GetColumnForCheckMetadata(tableName string) []string {

	var ColumnNameForCheckMetadata = []string{"id", "station_oldcode", "agency_name", "lat", "long", "geocode"} //"subbasin_name",

	if tableName != "m_ground" {
		ColumnNameForCheckMetadata = append(ColumnNameForCheckMetadata, "station_name")
	}

	if tableName == "m_dam" || tableName == "m_canal_station" || tableName == "m_tele_station" || tableName == "m_floodforecast_station" {
		ColumnNameForCheckMetadata = append(ColumnNameForCheckMetadata, "subbasin_name")
	}

	return ColumnNameForCheckMetadata
}

//	ข้อมูลสำหรับ check metadata
//	Parameters:
//		param
//			Struct_Station_InputParam
//	Return:
//		Array Struct_Station_ForCheckMetadata
func GetStationInfoForCheckMetadata(param *Struct_Station_InputParam) ([]*Struct_Station_ForCheckMetadata, error) {

	err := checkInputParam(param)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Convert AgencyID type from string to int64
	intAgencyID, err := strconv.ParseInt(param.AgencyID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data            []*Struct_Station_ForCheckMetadata
		objCanalStation *Struct_Station_ForCheckMetadata

		_id               sql.NullInt64
		_station_oldcode  sql.NullString
		_station_name     sql.NullString
		_agency_name      sql.NullString
		_agency_shortname sql.NullString
		_lat              sql.NullFloat64
		_long             sql.NullFloat64
		_geocode          sql.NullInt64
		_subbasin_name    sql.NullString

		_result *sql.Rows
	)

	//SQL Condition
	var sqlCmdQuery string = ""
	switch param.TableName {
	case "m_air_station":
		sqlCmdQuery = sqlGetAirStation
	case "m_canal_station":
		sqlCmdQuery = sqlGetCanalStation
	case "m_dam":
		sqlCmdQuery = sqlGetDam
	case "m_medium_dam":
		sqlCmdQuery = sqlGetMediumDam
	case "m_ground":
		sqlCmdQuery = sqlGetGround
	case "m_tele_station":
		sqlCmdQuery = sqlGetTeleStation
	case "m_floodforecast_station":
		sqlCmdQuery = sqlGetFloodForecastStation
	case "m_ford_station":
		sqlCmdQuery = sqlGetFordStation
	case "m_waterquality_station":
		sqlCmdQuery = sqlGetWaterQualityStation
	case "m_swan_station":
		sqlCmdQuery = sqlGetSwanStation
	default:
		return nil, rest.NewError(422, "'"+param.TableName+"' table not found (model/station/get.go:106)", nil)
	}

	//	log.Printf(strings.Replace(sqlCmdQuery, "#{ColumnName}", param.ColumnName, -1), intAgencyID)
	_result, err = db.Query(strings.Replace(sqlCmdQuery, "#{ColumnName}", param.ColumnName, -1), intAgencyID)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_Station_ForCheckMetadata, 0)

	for _result.Next() {
		objCanalStation = &Struct_Station_ForCheckMetadata{}
		//Scan to execute query with variables
		if param.TableName == "m_ground" {
			err := _result.Scan(&_id, &_station_name, &_agency_name, &_agency_shortname, &_lat, &_long, &_geocode)
			if err != nil {
				return nil, err
			}
			
			objCanalStation.Name = json.RawMessage("{}")
			objCanalStation.SubbasinName = json.RawMessage("{}")
		} else {
			err := _result.Scan(&_id, &_station_oldcode, &_station_name, &_agency_name, &_agency_shortname, &_lat, &_long, &_geocode, &_subbasin_name)
			if err != nil {
				return nil, err
			}

			if _station_name.String == "" || !_station_name.Valid {
				_station_name.String = "{}"
			}
			
			if _subbasin_name.String == "" || !_subbasin_name.Valid {
				_subbasin_name.String = "{}"
			}

			objCanalStation.Name = json.RawMessage(_station_name.String)
			objCanalStation.SubbasinName = json.RawMessage(_subbasin_name.String)

		}

		if _agency_shortname.String == "" || !_agency_shortname.Valid {
			_agency_shortname.String = "{}"
		}
		if _agency_name.String == "" || !_agency_name.Valid {
			_agency_name.String = "{}"
		}

		//Generate AirStation object
		
		objCanalStation.Id = _id.Int64
		objCanalStation.OldCode = _station_oldcode.String
		objCanalStation.Lat = ValidData(_lat.Valid, _lat.Float64)
		objCanalStation.Long = ValidData(_long.Valid, _long.Float64)
		objCanalStation.Geocode = ValidData(_geocode.Valid, _geocode.Int64)
		//objCanalStation.Lat = _lat.Float64
		//objCanalStation.Long = _long.Float64
		//objCanalStation.Geocode = _geocode.Int64

		objCanalStation.AgencyName = json.RawMessage(_agency_name.String)
		objCanalStation.AgencyShortname = json.RawMessage(_agency_shortname.String)

		data = append(data, objCanalStation)
	}

	return data, nil
}

//	validate parameter
//	Parameters:
//		param
//			Struct_Station_InputParam
//	Return:
//		nil, error
func checkInputParam(param *Struct_Station_InputParam) error {

	//Check Parameters
	if param.AgencyID == "" {
		return errors.New("'agency_id' is not null.")
	}
	if param.ColumnName == "" {
		return errors.New("'column_name' is not null.")
	}

	return nil
}

func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
