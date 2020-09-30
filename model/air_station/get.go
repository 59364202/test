// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package air_station is a model for public.air_station table. This table store air_station.
package air_station

import (
	//model_agency "haii.or.th/api/thaiwater30/model/agency"
	//model_subbasin "haii.or.th/api/thaiwater30/model/subbasin"
	//model_geocode "haii.or.th/api/thaiwater30/model/geocode"
	//result "haii.or.th/api/thaiwater30/util/result"
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//		"log"
	"strconv"
)

func GetAirStationForCheckMetadata(param *Struct_AirStation_InputParam) ([]*Struct_AirStation_ForCheckMetadata, error) {

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
		data          []*Struct_AirStation_ForCheckMetadata
		objAirStation *Struct_AirStation_ForCheckMetadata

		_id               sql.NullInt64
		_station_oldcode  sql.NullString
		_station_name     sql.NullString
		_agency_name      sql.NullString
		_agency_shortname sql.NullString
		_lat              sql.NullFloat64
		_long             sql.NullFloat64
		_geocode          sql.NullString

		_result *sql.Rows
	)

	//Where Condition
	var sqlCmdWhere string = " WHERE a." + param.ColumnName + " IS NULL AND a.agency_id = $1 "

	//Query
	//	log.Printf(sqlGetAirStation+sqlCmdWhere+sqlGetAirStationOrderBy, intAgencyID)
	_result, err = db.Query(sqlGetAirStation+sqlCmdWhere+sqlGetAirStationOrderBy, intAgencyID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_AirStation_ForCheckMetadata, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_station_oldcode, &_station_name, &_agency_name, &_agency_shortname, &_lat, &_long, &_geocode)
		if err != nil {
			return nil, err
		}

		if _station_name.String == "" || !_station_name.Valid {
			_station_name.String = "{}"
		}
		if _agency_shortname.String == "" || !_agency_shortname.Valid {
			_agency_shortname.String = "{}"
		}
		if _agency_name.String == "" || !_agency_name.Valid {
			_agency_name.String = "{}"
		}

		//Generate AirStation object
		objAirStation = &Struct_AirStation_ForCheckMetadata{}
		objAirStation.Id = _id.Int64
		objAirStation.OldCode = _station_oldcode.String
		objAirStation.Name = json.RawMessage(_station_name.String)
		objAirStation.Lat, _ = _lat.Value()
		objAirStation.Long, _ = _long.Value()
		objAirStation.Geocode = _geocode.String

		objAirStation.AgencyName = json.RawMessage(_agency_name.String)
		objAirStation.AgencyShortname = json.RawMessage(_agency_shortname.String)

		data = append(data, objAirStation)
	}

	return data, nil
}

func checkInputParam(param *Struct_AirStation_InputParam) error {

	//Check Parameters
	if param.AgencyID == "" {
		return errors.New("'agency_id' is not null.")
	}
	if param.ColumnName == "" {
		return errors.New("'column_name' is not null.")
	}

	return nil
}
