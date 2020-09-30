// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Werawan Prongpanom <werawan@haii.or.th>

package rainfall_month

import (
	"database/sql"
	"haii.or.th/api/util/pqx"
	_"haii.or.th/api/util/rest"
	"fmt"
	"encoding/json"
//	"os"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_rain24 "haii.or.th/api/thaiwater30/model/rainfall24hr"
)

func GetRainfallMonth(p *Param_RainfallMonth) ([]*Struct_RainfallMonth, error){
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	strSql, param := Gen_SQL_GetRainfallMonth(p)
	
//	fmt.Println("SQL : ", strSql)
//	os.Exit(1)
	
	row, err := db.Query(strSql, param...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := make([]*Struct_RainfallMonth, 0)

	strDatetimeFormat := "2006-01-02 07:00:00"

	for row.Next() {
		var (
			_tele_station_id		int64
			_tele_station_lat		sql.NullFloat64
			_tele_station_long		sql.NullFloat64
			_tele_station_oldcode	sql.NullString
			_name					pqx.JSONRaw
			_amphoe_name			pqx.JSONRaw
			_tumbon_name			pqx.JSONRaw
			_province_code			sql.NullString
			_province_name			pqx.JSONRaw
			_area_code				sql.NullString
			_area_name				sql.NullString
			_agency_id				sql.NullString
			_agency_name			pqx.JSONRaw
			_rainfallmonth			sql.NullFloat64
			_rainfall_datestart		sql.NullString
			_rainfall_dateend		sql.NullString

			d		*Struct_RainfallMonth				= &Struct_RainfallMonth{}
			station	*model_rain24.Struct_TeleStation	= &model_rain24.Struct_TeleStation{}
			agency	*model_agency.Struct_Agency			= &model_agency.Struct_Agency{}
			geocode	*model_geocode.Struct_Geocode		= &model_geocode.Struct_Geocode{}
		)

//		err = row.Scan(&_tele_station_id, &_rainfall_datestart, &_rainfall_dateend, &_rainfallmonth, &_tele_station_oldcode, &_tele_station_lat, &_tele_station_long, &_area_code, &_area_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name, &_name)
		err = row.Scan(&_tele_station_id, &_rainfall_datestart, &_rainfall_dateend, &_rainfallmonth, &_tele_station_oldcode, &_tele_station_lat, &_tele_station_long, &_area_code, &_province_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name, &_name)

		if err != nil {
			return nil, err
		}

		d.RainMonth = fmt.Sprintf("%.2f", _rainfallmonth.Float64)
		if _rainfall_datestart.Valid {
			d.Start_date = pqx.NullStringToTime(_rainfall_datestart).Format(strDatetimeFormat)
		}
		if _rainfall_dateend.Valid {
			d.End_date = pqx.NullStringToTime(_rainfall_dateend).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = model_rain24.ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = model_rain24.ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
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

		data = append(data, d)
	}

	return data, err
}