// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainfall_daily is a model for public.rainfall_daily table. This table store rainfall_daily.
package rainfall_daily

import (
	"database/sql"
	"encoding/json"
	"time"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/sqltime"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	
	model_basin "haii.or.th/api/thaiwater30/model/basin"
)

//	get rain_daily
//	Parameters:
//		p
//			Param_RainfallDaily
//	Return:
//		Array Struct_RainfallDaily
func GetRain(p *Param_RainfallDaily) ([]*Struct_RainfallDaily, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows
		obj *Struct_RainfallDaily

		_tele_station_id   int64
		_tele_station_name sql.NullString
		_rainfall_value    sql.NullFloat64
		_rainfall_datetime sqltime.NullTime
		_province_code     sql.NullString
		_province_name     sql.NullString
		_tele_station_lat  sql.NullFloat64
		_tele_station_long sql.NullFloat64
		_amphoe_name       sql.NullString
		_tumbon_name       sql.NullString
		_agency_name       sql.NullString
		_warning_zone      sql.NullString
		_basin_id   		sql.NullInt64
		_basin_code 		sql.NullInt64
		_basin_name 		sql.NullString
		_sub_basin_id 		sql.NullInt64
	)
	strSql, param := Gen_SQL_SelectRain(p)
	row, err = db.Query(strSql, param...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*Struct_RainfallDaily, 0)
	for row.Next() {
		err = row.Scan(&_tele_station_id, &_tele_station_name, &_rainfall_value, &_rainfall_datetime, &_province_code, &_province_name, &_tele_station_lat, &_tele_station_long,
			&_amphoe_name, &_tumbon_name, &_agency_name, &_warning_zone, &_basin_id, &_basin_code, &_basin_name, &_sub_basin_id)
		if err != nil {
			return nil, err
		}
		if !_tele_station_name.Valid || _tele_station_name.String == "" {
			_tele_station_name.String = "{}"
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

		obj = &Struct_RainfallDaily{}
		obj.StationId = _tele_station_id
		obj.StationName = json.RawMessage(_tele_station_name.String)
		obj.Rainfall = _rainfall_value.Float64
		obj.DateTime = _rainfall_datetime.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		obj.ProvinceCode = _province_code.String
		obj.ProvinceName = json.RawMessage(_province_name.String)
		obj.StationLat, _ = _tele_station_lat.Value()
		obj.StationLong, _ = _tele_station_long.Value()
		obj.AmphoeName = json.RawMessage(_amphoe_name.String)
		obj.TumbonName = json.RawMessage(_tumbon_name.String)
		obj.AgencyName = json.RawMessage(_agency_name.String)
		obj.WarningZone = _warning_zone.String
		
		/*----- Basin -----*/
		obj.Basin = &model_basin.Struct_Basin{}
		obj.Basin.Id = _basin_id.Int64
		obj.Basin.Basin_code = _basin_code.Int64
		obj.SubBasin_id = _sub_basin_id.Int64

		if !_basin_name.Valid || _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		obj.Basin.Basin_name = json.RawMessage(_basin_name.String)

		data = append(data, obj)
	}
	return data, nil
}

//	get rain_daily graph
//	Parameters:
//		p
//			Param_RainfallDaily
//	Return:
//		Array Struct_RainfallDaily_Graph
func GetRainGraph(p *Param_RainfallDaily) ([]*Struct_RainfallDaily_Graph, error) {
	layout := setting.GetSystemSetting("setting.Default.DateFormat")
	if p.StationId == 0 {
		return nil, rest.NewError(422, "no station_id ", nil)
	}

	if p.IsDaily {
		if p.StratDate != "" && p.EndDate != "" { // ใส่ start_date , end_date มา
			tStartDate, err := time.Parse(layout, p.StratDate)
			if err != nil {
				return nil, rest.NewError(422, "invalid start date", err)
			}
			tEndDate, err := time.Parse(layout, p.EndDate)
			if err != nil {
				return nil, rest.NewError(422, "invalid end date", err)
			}
			if !tStartDate.Equal(tEndDate) { // เช็ค start_date != end_date
				if !tStartDate.Before(tEndDate) { // เช็ค start_date ต้องน้อยกว่า end_date
					return nil, rest.NewError(422, "invalid date range", nil)
				}
			}

			if tEndDate.Sub(tStartDate) > 31*24*time.Hour { // เช็ค start_date -> end_date ไม่เกิน 31 วัน
				return nil, rest.NewError(422, "limit date range", nil)
			}
		} else { // ไม่ใส่ start_date หรือ end_date บังคับเป็น ย้อนหลัง 7 วัน
			p.StratDate = time.Now().AddDate(0, 0, -7).Format(layout)
			p.EndDate = time.Now().AddDate(0, 0, -1).Format(layout)
		}

	} else if p.IsMonthly {
		if p.Month <= 0 {
			return nil, rest.NewError(422, "invalid month", nil)
		}
		if p.Year <= 0 {
			return nil, rest.NewError(422, "invalid year", nil)
		}
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	strSql, param := Gen_SQL_SelectRainGraph(p)
	row, err := db.Query(strSql, param...)
	if err != nil {
		return nil, err
	}
	var (
		obj       *Struct_RainfallDaily_Graph
		_date     sqltime.NullTime
		_rainfall sql.NullFloat64
	)
	data := make([]*Struct_RainfallDaily_Graph, 0)
	for row.Next() {
		err = row.Scan(&_date, &_rainfall)
		if err != nil {
			return nil, err
		}

		obj = &Struct_RainfallDaily_Graph{}
		obj.DateTime = _date.Time.Format(setting.GetSystemSetting("setting.Default.DateFormat"))
		if _rainfall.Valid {
			obj.Rainfall = _rainfall.Float64
		}

		data = append(data, obj)
	}

	return data, nil
}
