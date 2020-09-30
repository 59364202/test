// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package tele_waterlevel is a model for public.tele_waterlevel table. This table store tele waterlevel data.
package tele_waterlevel

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"
	//  float "haii.or.th/api/thaiwater30/util/float"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	//"haii.or.th/api/thaiwater30/util/month"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
)

//
//func GetWaterLevelThailand(p *Waterlevel_InputParam) ([]*Struct_Waterlevel, error) {
//  //Open Database
//  db, err := pqx.Open()
//  if err != nil {
//    return nil, errors.Repack(err)
//  }
//
//  //Variables
//  var (
//    data       []*Struct_Waterlevel
//    waterlevel *Struct_Waterlevel
//
//    strSQL       string = ""
//    strSQL_Where string = ""
//    param        []interface{}
//    row          *sql.Rows
//
//    _id                   int64
//    _tele_station_id      int64
//    _waterlevel_datetime  time.Time
//    _waterlevel_m         sql.NullFloat64
//    _waterlevel_msl       sql.NullFloat64
//    _pre_waterlevel_msl   sql.NullFloat64
//    _subbasin_id          int64
//    _agency_id            int64
//    _tele_station_name    sql.NullString
//    _tele_station_lat     sql.NullFloat64
//    _tele_station_long    sql.NullFloat64
//    _tele_station_oldcode sql.NullString
//    _ground_level         sql.NullFloat64
//    _min_bank             sql.NullFloat64
//    _storage_percent      sql.NullFloat64
//    _sort_order           sql.NullInt64
//    _table                sql.NullString
//    _tumbon_name          sql.NullString
//    _amphoe_name          sql.NullString
//    _province_code        sql.NullString
//    _province_name        sql.NullString
//    _basin_name           sql.NullString
//    _tumbon_code          sql.NullString
//    _amphoe_code          sql.NullString
//    _basin_id             sql.NullInt64
//    _agency_shortname     sql.NullString
//  )
//  // gen sql
//  strSQL = SQL_SelectWaterLevel
//  if p.Province_code != "" {
//    arr := strings.Split(p.Province_code, ",")
//    lenArr := len(arr)
//    for i, v := range arr {
//      param = append(param, v)
//      strSQL_Where += " province_code = $" + strconv.Itoa(len(param))
//      if i < lenArr-1 {
//        strSQL_Where += " OR "
//      }
//    }
//    strSQL_Where += " AND"
//  }
//  if p.Basin_id != "" {
//    arr := strings.Split(p.Basin_id, ",")
//    lenArr := len(arr)
//    for i, v := range arr {
//      param = append(param, v)
//      strSQL_Where += " basin_id = $" + strconv.Itoa(len(param))
//      if i < lenArr-1 {
//        strSQL_Where += " OR "
//      }
//    }
//    strSQL_Where += " AND"
//  }
//  if p.Subbasin_id != "" {
//    arr := strings.Split(p.Subbasin_id, ",")
//    lenArr := len(arr)
//    for i, v := range arr {
//      param = append(param, v)
//      strSQL_Where += " subbasin_id = $" + strconv.Itoa(len(param))
//      if i < lenArr-1 {
//        strSQL_Where += " OR "
//      }
//    }
//    strSQL_Where += " AND"
//  }
//  if p.Agency_id != "" {
//    arr := strings.Split(p.Agency_id, ",")
//    lenArr := len(arr)
//    for i, v := range arr {
//      param = append(param, v)
//      strSQL_Where += " ct.agency_id = $" + strconv.Itoa(len(param))
//      if i < lenArr-1 {
//        strSQL_Where += " OR "
//      }
//    }
//    strSQL_Where += " AND"
//  }
//  // query
//  if len(param) != 0 {
//    row, err = db.Query(strSQL+" WHERE "+strSQL_Where[0:len(strSQL_Where)-3]+SQL_SelectWaterLevel_OrderBy, param...)
//  } else {
//    log.Printf(strSQL + SQL_SelectWaterLevel_OrderBy)
//    row, err = db.Query(strSQL + SQL_SelectWaterLevel_OrderBy)
//  }
//  if err != nil {
//    return nil, err
//  }
//  // loop data result
//  for row.Next() {
//    err = row.Scan(&_id, &_tele_station_id, &_waterlevel_datetime, &_waterlevel_m, &_waterlevel_msl, &_pre_waterlevel_msl, &_subbasin_id, &_agency_id, &_tele_station_name,
//      &_tele_station_lat, &_tele_station_long, &_tele_station_oldcode, &_ground_level, &_min_bank, &_storage_percent, &_sort_order, &_table, &_tumbon_name, &_amphoe_name,
//      &_province_code, &_province_name, &_basin_name, &_tumbon_code, &_amphoe_code, &_basin_id, &_agency_shortname)
//    if err != nil {
//      return nil, err
//    }
//
//    if !_tumbon_name.Valid || _tumbon_name.String == "" {
//      _tumbon_name.String = "{}"
//    }
//    if !_amphoe_name.Valid || _amphoe_name.String == "" {
//      _amphoe_name.String = "{}"
//    }
//    if !_province_name.Valid || _province_name.String == "" {
//      _province_name.String = "{}"
//    }
//    if !_tele_station_name.Valid || _tele_station_name.String == "" {
//      _tele_station_name.String = "{}"
//    }
//    if !_basin_name.Valid || _basin_name.String == "" {
//      _basin_name.String = "{}"
//    }
//    if !_agency_shortname.Valid || _agency_shortname.String == "" {
//      _agency_shortname.String = "{}"
//    }
//
//    waterlevel = &Struct_Waterlevel{Id: _id, Waterlevel_datetime: _waterlevel_datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))}
//
//    //if _waterlevel_m.Valid {
//    //  waterlevel.Waterlevel_m = _waterlevel_m.Float64
//    //}
//    //if _waterlevel_msl.Valid {
//    //  waterlevel.Waterlevel_msl = _waterlevel_msl.Float64
//    //}
//    waterlevel.Storage_percent = ValidData(_storage_percent.Valid, _storage_percent.Float64)
//    waterlevel.Waterlevel_m = ValidData(_waterlevel_m.Valid, _waterlevel_m.Float64)
//    waterlevel.Waterlevel_msl = ValidData(_waterlevel_msl.Valid, _waterlevel_msl.Float64)
//    waterlevel.Pre_Waterlevel_msl = ValidData(_pre_waterlevel_msl.Valid, _pre_waterlevel_msl.Float64)
//
//    teleStation := &model_tele_station.Struct_TeleStation{Id: _tele_station_id, Subbasin_id: _subbasin_id}
//    teleStation.Tele_station_lat = _tele_station_lat.Float64
//    teleStation.Tele_station_long = _tele_station_long.Float64
//    teleStation.Tele_station_oldcode = _tele_station_oldcode.String
//    teleStation.Min_bank = _min_bank.Float64
//    teleStation.Ground_level = _ground_level.Float64
//    teleStation.Tele_station_name = json.RawMessage(_tele_station_name.String)
//    teleStation.Tele_station_oldcode = _tele_station_oldcode.String
//    waterlevel.Station = teleStation
//
//    geocode := &model_lt_geocode.Struct_Geocode{}
//    geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)
//    geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)
//    geocode.Province_name = json.RawMessage(_province_name.String)
//    geocode.Province_code = _province_code.String
//    waterlevel.Geocode = geocode
//
//    agency := &model_agency.Struct_Agency{Id: _agency_id}
//    agency.Agency_name = json.RawMessage(_agency_shortname.String)
//    agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
//    waterlevel.Agency = agency
//
//    basin := &model_basin.Struct_Basin{Id: _basin_id.Int64}
//    basin.Basin_name = json.RawMessage(_basin_name.String)
//    waterlevel.Basin = basin
//
//    //waterlevel.ProvinceName = json.RawMessage(_province_name.String)
//    //waterlevel.Name = json.RawMessage(_tele_station_name.String)
//    //waterlevel.Datetime = _waterlevel_datetime.Format("2006-01-02 15:04")
//    //waterlevel.Oldcode = _tele_station_oldcode.String
//    //waterlevel.Value = ValidData(_waterlevel_msl.Valid, _waterlevel_msl.Float64)
//    //waterlevel.DataID = _id
//    //waterlevel.StationID = _tele_station_id
//    waterlevel.SortOrder = _sort_order.Int64
//    waterlevel.Table = _table.String
//
//    data = append(data, waterlevel)
//  }
//
//  return data, nil
//}
//func GetWaterLevelAnHourBefore() ([]*Struct_Waterlevel, error) {
//  //Open Database
//  db, err := pqx.Open()
//  if err != nil {
//    return nil, errors.Repack(err)
//  }
//
//  //Variables
//  var (
//    data       []*Struct_Waterlevel
//    waterlevel *Struct_Waterlevel
//
//    row *sql.Rows
//
//    _id                   int64
//    _waterlevel_datetime  time.Time
//    _waterlevel_m         sql.NullFloat64
//    _waterlevel_msl       sql.NullFloat64
//    _subbasin_id          int64
//    _agency_id            int64
//    _tele_station_name    sql.NullString
//    _tele_station_lat     sql.NullFloat64
//    _tele_station_long    sql.NullFloat64
//    _tele_station_oldcode sql.NullString
//    _ground_level         sql.NullFloat64
//    _min_bank             sql.NullFloat64
//    _tumbon_name          sql.NullString
//    _amphoe_name          sql.NullString
//    _province_code        sql.NullString
//    _province_name        sql.NullString
//    _storage_percent      sql.NullFloat64
//    _basin_name           sql.NullString
//  )
//  // query
//  row, err = db.Query(SQL_SelectWaterLevelAnHourBefore)
//  if err != nil {
//    return nil, pqx.GetRESTError(err)
//  }
//  // loop data result
//  for row.Next() {
//    err = row.Scan(&_id, &_waterlevel_datetime, &_waterlevel_m, &_waterlevel_msl, &_subbasin_id, &_agency_id, &_tele_station_name, &_tele_station_lat, &_tele_station_long, &_tele_station_oldcode,
//      &_ground_level, &_min_bank, &_tumbon_name, &_amphoe_name, &_province_code, &_province_name, &_storage_percent, &_basin_name)
//    if err != nil {
//      return nil, err
//    }
//
//    if !_tumbon_name.Valid || _tumbon_name.String == "" {
//      _tumbon_name.String = "{}"
//    }
//    if !_amphoe_name.Valid || _amphoe_name.String == "" {
//      _amphoe_name.String = "{}"
//    }
//    if !_province_name.Valid || _province_name.String == "" {
//      _province_name.String = "{}"
//    }
//    if !_tele_station_name.Valid || _tele_station_name.String == "" {
//      _tele_station_name.String = "{}"
//    }
//    if !_basin_name.Valid || _basin_name.String == "" {
//      _basin_name.String = "{}"
//    }
//
//    waterlevel = &Struct_Waterlevel{Waterlevel_m: _waterlevel_m.Float64, Waterlevel_msl: _waterlevel_msl.Float64, Storage_percent: _storage_percent.Float64,
//      Waterlevel_datetime: _waterlevel_datetime.Format("2006-01-02 15:04")}
//
//    teleStation := &model_tele_station.Struct_TeleStation{Id: _id, Subbasin_id: _subbasin_id, Agency_id: _agency_id}
//    teleStation.Tele_station_lat = _tele_station_lat.Float64
//    teleStation.Tele_station_long = _tele_station_long.Float64
//    teleStation.Tele_station_oldcode = _tele_station_oldcode.String
//    teleStation.Min_bank = _min_bank.Float64
//    teleStation.Ground_level = _ground_level.Float64
//    teleStation.Tele_station_name = json.RawMessage(_tele_station_name.String)
//    teleStation.Basin_name = json.RawMessage(_basin_name.String)
//    waterlevel.Station = teleStation
//
//    geocode := &model_lt_geocode.Struct_Geocode{}
//    geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)
//    geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)
//    geocode.Province_name = json.RawMessage(_province_name.String)
//    geocode.Province_code = _province_code.String
//    teleStation.Geocode = geocode
//
//    data = append(data, waterlevel)
//  }
//
//  return data, nil
//}

//  Get waterlevel by station and date
//  Parameters:
//    param
//      Waterlevel_InputParam
//  Return:
//    Array GetWaterlevelLastest_OutputParam
func GetWaterlevelByStationAndDate(param *Waterlevel_InputParam) (*GetWaterlevelLastest_OutputParam, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data          []*WaterlevelLastestStruct
		objWaterlevel *WaterlevelLastestStruct

		_id                  sql.NullInt64
		_waterlevel_datetime time.Time
		_waterlevel_m        sql.NullFloat64
		_waterlevel_msl      sql.NullFloat64
		_flow_rate           sql.NullFloat64
		_discharge           sql.NullFloat64

		_result *sql.Rows
	)

	//Query
	//log.Printf(sqlGetWaterlevelByStationAndDate + sqlGetWaterlevelByStationAndDateOrderBy, param.Station_id, param.Start_date, param.End_date)
	_result, err = db.Query(sqlGetWaterlevelByStationAndDate+sqlGetWaterlevelByStationAndDateOrderBy,
		param.Station_id, param.Start_date, param.End_date)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*WaterlevelLastestStruct, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_waterlevel_datetime,
			&_waterlevel_m, &_waterlevel_msl, &_flow_rate, &_discharge)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily object
		objWaterlevel = &WaterlevelLastestStruct{}
		objWaterlevel.Id = _id.Int64
		objWaterlevel.Datetime = _waterlevel_datetime.Format(strDatetimeFormat)
		objWaterlevel.Waterlevel_m = ValidData(_waterlevel_m.Valid, _waterlevel_m.Float64)
		objWaterlevel.Waterlevel_msl = ValidData(_waterlevel_msl.Valid, _waterlevel_msl.Float64)
		objWaterlevel.Flow_rate = ValidData(_flow_rate.Valid, _flow_rate.Float64)
		objWaterlevel.Discharge = ValidData(_discharge.Valid, _discharge.Float64)

		data = append(data, objWaterlevel)
	}

	resultData := &GetWaterlevelLastest_OutputParam{}
	resultData.Data = data
	resultData.Header = arrWaterlevelLastestHeader

	return resultData, nil
}

//func GetWaterlevelGraphByStationAndDate(param *Waterlevel_InputParam) (*highchart.HighChart, error) {
//
//  //Open Database
//  db, err := pqx.Open()
//  if err != nil {
//    return nil, errors.Repack(err)
//  }
//
//  //Variables
//  var (
//    _hasRow              bool = false
//    _id                  sql.NullInt64
//    _waterlevel_datetime time.Time
//    _waterlevel_m        sql.NullFloat64
//    _waterlevel_msl      sql.NullFloat64
//    _waterlevel_in       sql.NullFloat64
//    _waterlevel_out      sql.NullFloat64
//    _waterlevel_out2     sql.NullFloat64
//    _flow_rate           sql.NullFloat64
//    _discharge           sql.NullFloat64
//
//    _result *sql.Rows
//  )
//  //  log.Println(_waterlevel_datetime)
//  //  log.Printf(sqlGetWaterlevelByStationAndDate+sqlGetWaterlevelByStationAndDateOrderBy, param.Station_id, param.Start_date, param.End_date)
//  _result, err = db.Query(sqlGetWaterlevelByStationAndDate+" ORDER BY waterlevel_datetime ASC ",
//    param.Station_id, param.Start_date, param.End_date)
//  if err != nil {
//    return nil, err
//  }
//  hc := highchart.NewChart()
//  series := hc.NewLineSerie("")
//  for _result.Next() {
//    _hasRow = true
//    err := _result.Scan(&_id, &_waterlevel_datetime, &_waterlevel_m, &_waterlevel_msl, &_waterlevel_in, &_waterlevel_out, &_waterlevel_out2, &_flow_rate, &_discharge)
//    if err != nil {
//      return nil, err
//    }
//    if _waterlevel_msl.Valid {
//      series.AddDateData(_waterlevel_datetime, _waterlevel_msl.Float64)
//    } else {
//      series.AddDateData(nil, nil)
//    }
//  }
//  i, err := strconv.ParseInt(param.Station_id, 10, 64)
//  if err != nil {
//    return nil, err
//  }
//  tele_station, err := model_tele_station.GetTeleStationById(i)
//  if err != nil {
//    return hc, nil
//  }
//
//  min_bank_str := strconv.FormatFloat(float.TwoDigit(tele_station.Min_bank), 'f', -1, 64)
//  pl, ts := hc.NewPlotLines("ตลิ่งต่ำสุด " + min_bank_str + " ")
//  pl.Value = tele_station.Min_bank
//  if _hasRow {
//    ts.AddDateData(_waterlevel_datetime, tele_station.Min_bank)
//  }
//
//  ground_level_str := strconv.FormatFloat(float.TwoDigit(tele_station.Ground_level), 'f', -1, 64)
//  pl, ts = hc.NewPlotLines("ท้องน้ำ " + ground_level_str + " ")
//  pl.Value = tele_station.Ground_level
//  if _hasRow {
//    ts.AddDateData(_waterlevel_datetime, tele_station.Ground_level)
//  }
//
//  return hc, nil
//}

// Get error waterlevel
//  Parameters:
//    param
//      Waterlevel_InputParam
//  Return:
//    Array Struct_Waterlevel_ErrorData
func GetErrorWaterlevel(param *Waterlevel_InputParam) ([]*Struct_Waterlevel_ErrorData, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data          []*Struct_Waterlevel_ErrorData
		objWaterlevel *Struct_Waterlevel_ErrorData

		_id                    sql.NullInt64
		_station_id            sql.NullInt64
		_oldcode               sql.NullString
		_date                  time.Time
		_station_name          sql.NullString
		_station_province_name sql.NullString
		_agency_name           sql.NullString
		_agency_shortname      sql.NullString

		_waterlevel_m   sql.NullFloat64
		_waterlevel_msl sql.NullFloat64
		//_waterlevel_in   sql.NullFloat64
		//_waterlevel_out  sql.NullFloat64
		//_waterlevel_out2 sql.NullFloat64
		_flow_rate sql.NullFloat64
		_discharge sql.NullFloat64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""
	var arrAgencyID = []string{}

	if param.Agency_id != "" {
		arrAgencyID = strings.Split(param.Agency_id, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyID) > 0 {
		if len(arrAgencyID) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Agency_id, " "))
			sqlCmdWhere = " AND d.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strID := range arrAgencyID {
				arrParam = append(arrParam, strings.Trim(strID, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere = " AND d.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	if param.Start_date != "" && param.End_date != "" {
		arrParam = append(arrParam, param.Start_date)
		sqlCmdWhere += " AND waterlevel_datetime >= $" + strconv.Itoa(len(arrParam))

		arrParam = append(arrParam, param.End_date+" 23:59")
		sqlCmdWhere += " AND waterlevel_datetime <= $" + strconv.Itoa(len(arrParam))
	}

	//Query
	//  log.Println(sqlGetErrorWaterlevel + sqlCmdWhere)
	//  log.Println(arrParam)
	_result, err = db.Query(sqlGetErrorWaterlevel+sqlCmdWhere, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Loop data result
	data = make([]*Struct_Waterlevel_ErrorData, 0)
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_oldcode, &_date, &_station_name, &_station_province_name, &_agency_name, &_agency_shortname,
			&_waterlevel_m, &_waterlevel_msl, &_flow_rate, &_discharge, &_station_id)
		//&_waterlevel_m, &_waterlevel_msl, &_waterlevel_in, &_waterlevel_out, &_waterlevel_out2, &_flow_rate, &_discharge, &_station_id)
		if err != nil {
			return nil, err
		}

		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}
		if !_station_province_name.Valid || _station_province_name.String == "" {
			_station_province_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}

		//Generate DamDaily object
		objWaterlevel = &Struct_Waterlevel_ErrorData{}
		objWaterlevel.ID = _id.Int64
		objWaterlevel.StationID = _station_id.Int64
		objWaterlevel.StationOldCode = _oldcode.String
		objWaterlevel.Datetime = _date.Format(strDatetimeFormat)
		objWaterlevel.StationName = json.RawMessage(_station_name.String)
		objWaterlevel.ProvinceName = json.RawMessage(_station_province_name.String)
		objWaterlevel.AgencyName = json.RawMessage(_agency_name.String)
		objWaterlevel.AgencyShortName = json.RawMessage(_agency_shortname.String)

		//objWaterlevel.WaterlevelIn = ValidData(_waterlevel_in.Valid, _waterlevel_in.Float64)
		objWaterlevel.WaterlevelM = ValidData(_waterlevel_m.Valid, _waterlevel_m.Float64)
		objWaterlevel.WaterlevelMsl = ValidData(_waterlevel_msl.Valid, _waterlevel_msl.Float64)
		//objWaterlevel.WaterlevelOut = ValidData(_waterlevel_out.Valid, _waterlevel_out.Float64)
		//objWaterlevel.WaterlevelOut2 = ValidData(_waterlevel_out2.Valid, _waterlevel_out2.Float64)
		objWaterlevel.FlowRate = ValidData(_flow_rate.Valid, _flow_rate.Float64)
		objWaterlevel.Discharge = ValidData(_discharge.Valid, _discharge.Float64)

		data = append(data, objWaterlevel)
	}

	return data, nil
}

// get tele waterlevel by station and date for page analyst
//  Parameters:
//    param
//      Waterlevel_InputParam
//  Return:
//    GetWaterlevelGraphByStationAndDateAnalystOutput
func GetWaterlevelGraphByStationAndDateAnalyst(param *Waterlevel_InputParam) (*GetWaterlevelGraphByStationAndDateAnalystOutput, error) {

	// check station id
	if param.Station_id == "" {
		return nil, rest.NewError(422, "No station id", nil)
	}
	data := &GetWaterlevelGraphByStationAndDateAnalystOutput{}
	// get data
	graphData, err := getWaterlevelGraphAnalyst(param.Station_id, param.Start_date, param.End_date)
	if err != nil {
		return nil, err
	}
	// add data
	data.GraphData = graphData
	i, err := strconv.ParseInt(param.Station_id, 10, 64)
	if err != nil {
		return nil, err
	}
	tele_station, err := model_tele_station.GetTeleStationById(i)
	if err != nil {
		return data, nil
	}
	data.MinBank = tele_station.Min_bank
	data.GroundLevel = tele_station.Ground_level
	return data, nil
}

// get tele waterlevel yearly for graph page analyst
//  Parameters:
//    param
//      GetWaterlevelYearlyGraphInput
//  Return:
//    GetWaterlevelYearlyGraphAnalystOutput
func GetWaterlevelYearlyGraphAnalyst(param *GetWaterlevelYearlyGraphInput) (*GetWaterlevelYearlyGraphAnalystOutput, error) {

	// check station id
	if param.StationID == "" {
		return nil, rest.NewError(422, "No station id", nil)
	}

	data := &GetWaterlevelYearlyGraphAnalystOutput{}
	// get data
	aGraphData := make([]*GetWaterlevelYearlyGraphAnalystOutputYear, 0)
	for _, v := range param.Year {
		gData := &GetWaterlevelYearlyGraphAnalystOutputYear{}
		gData.Year = v
		// get data by year
		graphData, err := getWaterlevelGraphAnalyst(param.StationID, strconv.Itoa(v)+"-01-01 00:00:00", strconv.Itoa(v)+"-12-31 23:59:59")
		if err != nil {
			return nil, err
		}
		// add data
		gData.GraphData = graphData
		aGraphData = append(aGraphData, gData)
	}
	data.GraphData = aGraphData

	i, err := strconv.ParseInt(param.StationID, 10, 64)
	if err != nil {
		return nil, err
	}
	// get info. tele station
	tele_station, err := model_tele_station.GetTeleStationById(i)
	if err != nil {
		return data, nil
	}
	data.MinBank = tele_station.Min_bank
	data.GroundLevel = tele_station.Ground_level
	return data, nil
}

// watergate ปตร./ฝาย
// get tele waterlevel field in,out for page analyst
//  Parameters:
//    None
//  Return:
//    Array GetWaterlevelInOutLatestAnalystOutput
func GetWaterlevelInOutLatestAnalyst() ([]*GetWaterlevelInOutLatestAnalystOutput, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	r := []*GetWaterlevelInOutLatestAnalystOutput{}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}
	// test at local
	//	strDatetimeFormat = time.RFC3339

	q := sqlSelectWaterlevelInOutLatest
	p := []interface{}{}
	//	fmt.Println(q)

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			_waterlevel_datetime time.Time
			_waterlevel_in       sql.NullFloat64
			_waterlevel_out      sql.NullFloat64
			_waterlevel_out2     sql.NullFloat64
			_pump                sql.NullInt64
			_floodgate           sql.NullInt64
			_pump_on             sql.NullInt64
			_floodgate_open      sql.NullInt64
			_floodgate_height    sql.NullFloat64

			_station_id      int64
			_station_name    pqx.JSONRaw
			_station_lat     sql.NullFloat64
			_station_long    sql.NullFloat64
			_station_oldcode sql.NullString

			_agency_id        int64
			_agency_name      pqx.JSONRaw
			_agency_shortname pqx.JSONRaw

			_gecode_id     int64
			_geocode       sql.NullString
			_tumbon_code   string
			_tumbon_name   pqx.JSONRaw
			_amphoe_code   string
			_amphoe_name   pqx.JSONRaw
			_province_code string
			_province_name pqx.JSONRaw
			_area_code     string
			_area_name     pqx.JSONRaw

			_subbasin_id            sql.NullInt64
			_basin_id               int64
			_basin_name             pqx.JSONRaw
			_watergate_out_datetime time.Time
		)
		err = rows.Scan(&_station_name, &_waterlevel_datetime, &_waterlevel_in,
			&_waterlevel_out, &_waterlevel_out2, &_station_id, &_station_lat, &_station_long, &_station_oldcode,
			&_gecode_id, &_geocode, &_province_code, &_province_name, &_amphoe_code,
			&_amphoe_name, &_tumbon_code, &_tumbon_name,
			&_pump, &_floodgate, &_pump_on, &_floodgate_open, &_floodgate_height,
			&_agency_id, &_agency_name, &_agency_shortname, &_area_code, &_area_name, &_basin_id, &_basin_name, &_watergate_out_datetime, &_subbasin_id)

		teleStation := &model_tele_station.Struct_TeleStation{
			Id:                   _station_id,
			Tele_station_name:    _station_name.JSON(),
			Tele_station_oldcode: _station_oldcode.String,
			Subbasin_id:          _subbasin_id.Int64,
			Agency_id:            _agency_id,
			Geocode_id:           _gecode_id,
		}
		teleStation.Tele_station_lat, _ = _station_lat.Value()
		teleStation.Tele_station_long, _ = _station_long.Value()

		geocode := &model_lt_geocode.Struct_Geocode{}
		geocode.Tumbon_code = _tumbon_code
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_code = _amphoe_code
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code
		geocode.Province_name = _province_name.JSON()
		geocode.Area_code = _area_code
		geocode.Area_name = _area_name.JSON()

		agency := &model_agency.Struct_Agency{
			Agency_shortname: _agency_shortname.JSON(),
		}
		agency.Id = _agency_id
		agency.Agency_name = _agency_name.JSON()
		agency.Agency_shortname = _agency_shortname.JSON()

		basin := &model_basin.Struct_Basin{
			Id:         _basin_id,
			Basin_name: _basin_name.JSON(),
		}

		DatetimeOut := ""
		if _watergate_out_datetime.Format(strDatetimeFormat) == "0001-01-01 00:00" {
			DatetimeOut = _waterlevel_datetime.Format(strDatetimeFormat)
		} else {
			DatetimeOut = _watergate_out_datetime.Format(strDatetimeFormat)
		}

		data := &GetWaterlevelInOutLatestAnalystOutput{
			DatetimeIn:   _waterlevel_datetime.Format(strDatetimeFormat),
			WaterlevelIn: ValidData(_waterlevel_in.Valid, _waterlevel_in.Float64),

			// ข้อมูลโทรมาตร สสนก. วัดระดับน้ำที่ปตร. บางสถานีจะมี 2 หัววัด หรือแยกเป็น 2 สุถานี ทำให้วันที่ in/out อาจจะไม่เท่ากัน
			DatetimeOut:    DatetimeOut,
			WaterlevelOut:  ValidData(_waterlevel_out.Valid, _waterlevel_out.Float64),
			WaterlevelOut2: ValidData(_waterlevel_out2.Valid, _waterlevel_out2.Float64),

			Pump:            ValidData(_pump.Valid, _pump.Int64),
			PumpOn:          ValidData(_pump_on.Valid, _pump_on.Int64),
			Floodgate:       ValidData(_floodgate.Valid, _floodgate.Int64),
			FloodgateHeight: ValidData(_floodgate_height.Valid, _floodgate_height.Float64),
			FloodgateOpen:   ValidData(_floodgate_open.Valid, _floodgate_open.Int64),

			Station: teleStation,
			Geocode: geocode,
			Agency:  agency,
			Basin:   basin,
		}

		r = append(r, data)
	}

	return r, nil
}

// water gate ปตร./ฝาย
// get tele waterlevel in out for graph
//  Parameters:
//    inputData
//      GetWaterlevelInOutGrapthAnalystInput
//  Return:
//    Array GetWaterlevelInOutGrapthAnalystOutput
func GetWaterlevelInOutGraphAnalyst(inputData *GetWaterlevelInOutGrapthAnalystInput) ([]*GetWaterlevelInOutGrapthAnalystOutput, error) {

	// check station and date range
	if inputData.StationID == 0 || inputData.EndDate == "" || inputData.StartDate == "" {
		return nil, rest.NewError(422, "Invalid Parameter", nil)
	}

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	groundLevel, criticallevel := getStationInOut(inputData.StationID)

	// sql get waterlevel by station and date
	q := sqlSelectWatergateByStationAndDateAnalyst
	p := []interface{}{inputData.StationID, inputData.StartDate, inputData.EndDate + " 23:59:59"}
	//	fmt.Println(q)

	// process sql get waterlevel by station and date
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// define value
	data := make([]*GetWaterlevelInOutGrapthAnalystOutput, 0)
	dataIn := &GetWaterlevelInOutGrapthAnalystOutput{}
	dataIn.Name = "เหนือ ปตร."
	dataIn.GroundLevel = groundLevel
	dataIn.CriticalLevel = criticallevel

	dataOut := &GetWaterlevelInOutGrapthAnalystOutput{}
	dataOut.Name = "ท้าย ปตร."
	dataOut.GroundLevel = groundLevel
	dataOut.CriticalLevel = criticallevel
	graphDataIn := make([]*GetWaterlevelGraphByStationAndDateAnalystDataOutput, 0)
	graphDataOut := make([]*GetWaterlevelGraphByStationAndDateAnalystDataOutput, 0)

	for rows.Next() {
		var (
			id                  sql.NullInt64
			waterlevel_datetime time.Time
			_waterlevel_in      sql.NullFloat64
			_waterlevel_out     sql.NullFloat64
		)
		rows.Scan(&id, &waterlevel_datetime, &_waterlevel_in, &_waterlevel_out)

		dataRowIn := &GetWaterlevelGraphByStationAndDateAnalystDataOutput{}
		dataRowOut := &GetWaterlevelGraphByStationAndDateAnalystDataOutput{}
		//		dataRowIn.Datetime = waterlevel_datetime.Format(time.RFC3339)
		dataRowIn.Datetime = waterlevel_datetime.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		dataRowIn.Value = ValidData(_waterlevel_in.Valid, _waterlevel_in.Float64)

		//		dataRowOut.Datetime = waterlevel_datetime.Format(time.RFC3339)
		dataRowOut.Datetime = waterlevel_datetime.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		dataRowOut.Value = ValidData(_waterlevel_out.Valid, _waterlevel_out.Float64)

		fmt.Println(waterlevel_datetime.Format(time.RFC3339))
		fmt.Println(_waterlevel_in)
		fmt.Println(_waterlevel_out)
		fmt.Println(ValidData(_waterlevel_out.Valid, _waterlevel_out.Float64))

		graphDataIn = append(graphDataIn, dataRowIn)
		graphDataOut = append(graphDataOut, dataRowOut)
	}
	dataIn.Data = graphDataIn
	dataOut.Data = graphDataOut
	data = append(data, dataIn)
	data = append(data, dataOut)

	return data, nil
}

// get station information for in and out value
//  Parameters:
//    stationID
//      tele_station_id
//  Return:
//    ground_level, critical_level
func getStationInOut(stationID int64) (interface{}, interface{}) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, nil
	}
	p := []interface{}{stationID}
	q := "SELECT ground_level,critical_level FROM public.m_tele_station WHERE deleted_at=to_timestamp(0) AND id=$1"
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	var (
		_groundLevel   sql.NullFloat64
		_criticalLevel sql.NullFloat64
	)
	groundLevel := make([]interface{}, 0)
	criticalLevel := make([]interface{}, 0)
	for rows.Next() {
		rows.Scan(&_groundLevel, &_criticalLevel)

		if _groundLevel.Valid {
			groundLevel = append(groundLevel, _groundLevel.Float64)
		} else {
			groundLevel = append(groundLevel, nil)
		}
		if _criticalLevel.Valid {
			criticalLevel = append(criticalLevel, _criticalLevel.Float64)
		} else {
			criticalLevel = append(criticalLevel, nil)
		}
	}
	// return data
	return groundLevel, criticalLevel
}

// get tele waterlevel graph analyst by station id and date range
//กราฟนักวิเคราะห์ รหัสสถานี วันที่เริ่มต้น วันที่สิ้นสุด
//  Parameters:
//    stationID
//      tele station id
//    startDate
//      date start
//    endDate
//      date end
//  Return:
//    Array GetWaterlevelGraphByStationAndDateAnalystDataOutput
func getWaterlevelGraphAnalyst(stationID, startDate, endDate string) ([]*GetWaterlevelGraphByStationAndDateAnalystDataOutput, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get waterlevel by station and date
	q := sqlSelectWaterlevelByStationAndDateAnalyst
	p := []interface{}{stationID, startDate, endDate}

	fmt.Println(q)

	// process sql get waterlevel by station and date
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var (
		id                  sql.NullInt64
		waterlevel_datetime time.Time
		waterlevel_msl      sql.NullFloat64
	)

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02 15:04"
	}

	graphData := make([]*GetWaterlevelGraphByStationAndDateAnalystDataOutput, 0)
	for rows.Next() {
		err := rows.Scan(&id, &waterlevel_datetime, &waterlevel_msl)
		if err != nil {
			return nil, err
		}
		dataRow := &GetWaterlevelGraphByStationAndDateAnalystDataOutput{}
		//dataRow.Datetime = waterlevel_datetime.Format(time.RFC3339)
		dataRow.Datetime = waterlevel_datetime.Format(strDatetimeFormat)
		dataRow.Value = ValidData(waterlevel_msl.Valid, waterlevel_msl.Float64)
		graphData = append(graphData, dataRow)
	}

	return graphData, nil
}

//  Get Waterlevel group by basin for graph
//  Parameters:
//    subbasin_id
//      subbasin id
//    datetime
//      datetime for get data
//  Return:
//    Array Struct_WaterlevelBasinGraphAnalystAdvance
func GetWaterlevelBasinGraphAnalystAdvance(subbasin_id int64, datetime string) ([]*Struct_WaterlevelBasinGraphAnalystAdvance, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		data []*Struct_WaterlevelBasinGraphAnalystAdvance
		obj  *Struct_WaterlevelBasinGraphAnalystAdvance
	)

	row, err := db.Query(SQL_WaterlevelBasinGraphAnalystAdvance(subbasin_id, datetime))
	if err != nil {
		return nil, err
	}
	data = make([]*Struct_WaterlevelBasinGraphAnalystAdvance, 0)
	for row.Next() {
		var (
			_id           int64
			_name         pqx.JSONRaw
			_agency_id    int64
			_ground_level sql.NullFloat64
			_min_bank     sql.NullFloat64
			_sort_order   sql.NullInt64
			_value        sql.NullFloat64
			_type         string
		)
		err = row.Scan(&_id, &_name, &_agency_id, &_ground_level, &_min_bank, &_sort_order, &_value, &_type)
		if err != nil {
			return nil, err
		}
		obj = &Struct_WaterlevelBasinGraphAnalystAdvance{}
		obj.Name = _name.JSONPtr()
		obj.GroundLevel = ValidData(_ground_level.Valid, _ground_level.Float64)
		obj.MinBank = ValidData(_min_bank.Valid, _min_bank.Float64)
		obj.Value = ValidData(_value.Valid, _value.Float64)
		data = append(data, obj)
	}

	return data, nil
}

//  Parameters:
//    subbasin_id
//      subbasin id
//  Return:
//    Array Struct_WaterlevelBasin24HGraphAnalystAdvance
func GetWaterlevelBasinGraph24HAnalystAdvance(subbasin_id int64) ([]*Struct_WaterlevelBasin24HGraphAnalystAdvance, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		data  []*Struct_WaterlevelBasin24HGraphAnalystAdvance
		obj   *Struct_WaterlevelBasin24HGraphAnalystAdvance
		gData *Struct_WaterlevelBasin24HGraphAnalystAdvance_GraphData
	)

	row, err := db.Query(SQL_WaterlevelBasinGraph24HAnalystAdvance(), subbasin_id)
	if err != nil {
		return nil, err
	}
	data = make([]*Struct_WaterlevelBasin24HGraphAnalystAdvance, 0)
	tempStr := ""
	for row.Next() {
		var (
			_id           int64
			_name         *pqx.JSONRaw
			_lat          sql.NullFloat64
			_long         sql.NullFloat64
			_ground_level sql.NullFloat64
			_left_bank    sql.NullFloat64
			_right_bank   sql.NullFloat64
			_distance     sql.NullFloat64
			_sort_order   sql.NullFloat64
			_datetime     time.Time
			_value        sql.NullFloat64
		)
		err = row.Scan(&_id, &_name, &_lat, &_long, &_ground_level, &_left_bank, &_right_bank, &_distance, &_sort_order, &_datetime, &_value)
		if err != nil {
			return nil, err
		}
		datetime := _datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))

		if tempStr != datetime {
			tempStr = datetime
			obj = &Struct_WaterlevelBasin24HGraphAnalystAdvance{}
			obj.Datetime = datetime
			obj.GraphData = make([]*Struct_WaterlevelBasin24HGraphAnalystAdvance_GraphData, 0)
			data = append(data, obj)
		}
		gData = &Struct_WaterlevelBasin24HGraphAnalystAdvance_GraphData{}
		gData.Name = _name.JSONPtr()
		//    gData.Lat = ValidData(_lat.Valid, _lat.String)
		//    gData.Long = ValidData(_long.Valid, _long.String)
		gData.GroundLevel = ValidData(_ground_level.Valid, _ground_level.Float64)
		gData.LeftBank = ValidData(_left_bank.Valid, _left_bank.Float64)
		gData.RightBank = ValidData(_right_bank.Valid, _right_bank.Float64)
		gData.Value = ValidData(_value.Valid, _value.Float64)
		gData.Distance = ValidData(_distance.Valid, _distance.Float64)
		obj.GraphData = append(obj.GraphData, gData)
	}

	return data, nil
}

// get waterlevel latest for flood forecast
//  Parameters:
//    None
//  Return:
//    Array ObserveWaterlevelOutput
func GetWaterlevelLatestForFloodforecast() ([]*ObserveWaterlevelOutput, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	q := sqlSelectWaterlevelLatestForFloodforecast
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, errors.Repack(err)
	}

	data := make([]*ObserveWaterlevelOutput, 0)
	for rows.Next() {
		var (
			id         sql.NullInt64
			mOldcode   sql.NullString
			datetime   time.Time
			value      sql.NullFloat64
			lat        sql.NullFloat64
			long       sql.NullFloat64
			mid        sql.NullInt64
			mName      pqx.JSONRaw
			aid        sql.NullInt64
			aNmae      pqx.JSONRaw
			aShortName pqx.JSONRaw
		)
		rows.Scan(&id, &datetime, &value, &mid, &mName, &lat, &long, &mOldcode, &aid, &aNmae, &aShortName)

		dataRow := &ObserveWaterlevelOutput{}
		dataRow.ID = id.Int64
		dataRow.TeleWaterlevelValue = ValidData(value.Valid, value.Float64)
		dataRow.TeleWaterlevelDatetime = datetime.Format("2006-01-02 15:04")

		station := &WaterlevelTeleStation{}
		station.ID = mid.Int64
		station.Name = mName.JSON()
		station.Lat = ValidData(lat.Valid, lat.Float64)
		station.Long = ValidData(long.Valid, long.Float64)
		station.OldCode = mOldcode.String
		dataRow.Station = station

		agency := &Agency{}
		agency.ID = aid.Int64
		agency.Name = aNmae.JSON()
		agency.ShortName = aShortName.JSON()
		dataRow.Agency = agency

		data = append(data, dataRow)
	}

	return data, nil
}

// add func check valid data
//ตรวจสอบค่าที่ได้จากฐานข้อมูล ว่าเป็น null หรือไม่ ถ้าใช่ ให้ return null ใช้กับ table ที่มี column เป็นตัวเลข
//ตัวแปรที่เราไปรับค่าจาก db เป็น float64 ซึ่ง float64 เป็น null ไม่ได้
//ถ้าาใช้  val.float64 จะได้เป็น 0 ทั้งๆที่ใน db เป็น null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

type Struct_wl_station_graph struct {
	Date_timpestamp string `json:"date_timestamp"` //example:`2016-05-23 14:00:00`
	Values          string `json:"value"`          //example:`2.16`
}

func Get_Wl_station_graph(station_id string) ([]*Struct_wl_station_graph, error) {
	if station_id == "" {
		return nil, errors.New("No station_id")
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	currentTimestamp := now.Format("2006-01-02 15:04")
	currentTimestamp7D := now.AddDate(0, 0, -7).Format("2006-01-02 15:04")
	var q string = `
  SELECT waterlevel_datetime,
       ROUND((waterlevel_msl + COALESCE(mt.offset,0))::numeric, 2) AS waterlevel_msl
  FROM tele_waterlevel tw
        INNER JOIN m_tele_station mt ON mt.id = tw.tele_station_id
  WHERE tele_station_id = $1
    AND waterlevel_datetime BETWEEN '` + currentTimestamp7D + `' AND '` + currentTimestamp + `'
    -- AND waterlevel_datetime = date_trunc('hour', waterlevel_datetime)
    AND tw.deleted_at = to_timestamp(0)
    AND ( tw.qc_status IS NULL OR tw.qc_status->>'is_pass' = 'true' )
  ORDER BY waterlevel_datetime::date,
           waterlevel_datetime::TIME`

	rows, err := db.Query(q, station_id)
	if err != nil {
		return nil, err
	}

	var rs []*Struct_wl_station_graph = make([]*Struct_wl_station_graph, 0)

	for rows.Next() {
		var (
			_waterelevel_date sql.NullString
			_waterlevel_msl   sql.NullString
		)

		err := rows.Scan(&_waterelevel_date, &_waterlevel_msl)
		if err != nil {
			return nil, err
		}

		//    z := pqx.NullStringToTime(_waterelevel_date)
		wl_datetime := pqx.NullStringToTime(_waterelevel_date)

		s := &Struct_wl_station_graph{
			//Date_timpestamp: wl_datetime.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat")),
			Date_timpestamp: wl_datetime.Format("2006-01-02 15:04:05"),
			Values:          _waterlevel_msl.String,
		}
		rs = append(rs, s)
	}

	return rs, err
}

type Struct_wl_latest_list_prov struct {
	Wl_datas []*Struct_waterlevel_datas `json:"wl_datas"`
	Map      []*Struct_waterlevel_map   `json:"map"`
}

type Struct_waterlevel_map struct {
	Image_url   string `json:"image_url"`   //example:"http://www.nhc.in.th/product/latest/img/rain24.jpg"
	Image_time  string `json:"image_time"`  //example:"15:00 น."
	Detail_link string `json:"detail_link"` //example:"http://www.nhc.in.th/web/index.php?model=telemetering&view=weather"
}

type Struct_waterlevel_datas struct {
	Station_name      string `json:"station_name"`      //example:`สถานีห้วยแร้ง จ.ตราด`
	Tele_station_id   string `json:"tele_station_id"`   //example:`telehaii0912`
	Province_name     string `json:"province_name"`     //example:`ตราด`
	Wl_msl            string `json:"wl_msl"`            //example:`2.21`
	Percent           string `json:"percent"`           //example:`92.20 %`
	Situation         string `json:"situation"`         //example:`น้ำมาก`
	Situation_color   string `json:"situation_color"`   //example:`#003CFA`
	Tele_station_lat  string `json:"tele_station_lat"`  //example:`12.325200`
	Tele_station_long string `json:"tele_station_long"` //example:`102.500000`
	Wl_tele_date      string `json:"wl_tele_date"`      //example:`2017-11-14`
	Wl_tele_time      string `json:"wl_tele_time"`      //example:`22:20:00`
	Wl_tele_date_thai string `json:"wl_tele_date_thai"` //example:`14 พ.ย. 2560`
	Province_id       string `json:"province_id"`       //example:`23`
	Ground_level      string `json:"ground_level"`      //example:`-14.627`
	Water_bank        string `json:"water_bank"`        //example:`3.634`
	Warn_level        string `json:"warn_level"`        //example:`9999.99`
	Wl_tele_datetime  string `json:"wl_tele_datetime"`  //example:`2017-11-14 22:20 น.`
	Amphoe_name       string `json:"amphoe_name"`       //example:`เมืองตราด`
	Before_level      string `json:"before_level"`      //example:`fa fa-circle-o`
}

func Get_wl_latest_list_prov(prov_id string) (*Struct_wl_latest_list_prov, error) {
	if prov_id == "" {
		return nil, errors.New("no prov_id")
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var q string = `
    SELECT concat('สถานี',st.tele_station_name ->> 'th',' จ.',pn.province_name ->> 'th') AS station_name,
       lr.tele_station_id,
       pn.province_name ->> 'th' AS province_name,
	    lr.value_current + COALESCE(st.offset,0) AS wl_msl,
	    CASE
	        WHEN lr.value_current = 999999 THEN NULL
	        WHEN (st2.water_bank - st.ground_level) != 0 THEN (
	            (lr.value_current - st.ground_level) * 100
	        ) / (st2.water_bank - st.ground_level)
	        ELSE NULL
	    END AS percent,      
	    st.tele_station_lat,
	    st.tele_station_long,
	    to_char(lr.datetime_current::date, 'YYYY-MM-DD') AS wl_tele_date,
	    to_char(lr.datetime_current::TIME, 'HH24:MM:SS') AS wl_tele_time,
	    pn.province_code,
	    st.ground_level,
	    st2.water_bank,
	    to_char(lr.datetime_current,'YYYY-MM-DD HH24:MI') AS waterlevel_datetime,
	    pn.amphoe_name ->> 'th' AS amphoe_name
    FROM cache.latest_waterlevel lr
    LEFT JOIN public.m_tele_station st ON lr.tele_station_id = st.id
    LEFT JOIN public.lt_geocode pn ON st.geocode_id = pn.id
       LEFT  JOIN ignore ig ON st.id = ig.station_id ::int AND ig.data_category = 'tele_waterlevel'    
    LEFT JOIN
      ( SELECT m_tele_station.id,
               left_bank,
               right_bank,
               CASE
                   WHEN left_bank > right_bank THEN right_bank
                   ELSE left_bank
               END AS water_bank
       FROM m_tele_station	         
       WHERE left_bank IS NOT NULL AND right_bank IS NOT NULL
    ) st2 ON lr.tele_station_id = st2.id
    WHERE pn.province_code = $1
      AND st.ground_level IS NOT NULL
      AND lr.value_current IS NOT NULL
      AND st.left_bank IS NOT NULL
      AND st.right_bank IS NOT NULL
      AND st.geocode_id IS NOT NULL
      --AND lr.datetime_current::date = current_date
      AND lr.datetime_current >= $2
      AND ( lr.qc_status IS NULL OR lr.qc_status->>'is_pass' = 'true' )
      AND (ig.is_ignore = false OR ig.is_ignore IS NULL)
    ORDER BY lr.datetime_current::TIME , percent DESC
    LIMIT 20
  `
	dt := time.Now().Add(-4 * time.Hour).Format("2006-01-02 15:00")

	waterlevel_setting := &uSetting.Struct_WaterlevelSetting{}
	err = setting.GetSystemSettingPtr("Frontend.public.waterlevel_setting", &waterlevel_setting)

	rows, err := db.Query(q, prov_id, dt)
	if err != nil {
		return nil, err
	}

	var rs *Struct_wl_latest_list_prov = &Struct_wl_latest_list_prov{}
	var lt_datas []*Struct_waterlevel_datas = make([]*Struct_waterlevel_datas, 0)
	var lt_map []*Struct_waterlevel_map = make([]*Struct_waterlevel_map, 0)

	for rows.Next() {
		var (
			_Station_name      sql.NullString
			_Tele_station_id   sql.NullString
			_Province_name     sql.NullString
			_Wl_msl            sql.NullString
			_Percent           sql.NullFloat64
			_Tele_station_lat  sql.NullString
			_Tele_station_long sql.NullString
			_Wl_tele_date      sql.NullString
			_Wl_tele_time      sql.NullString
			_Province_id       sql.NullString
			_Ground_level      sql.NullString
			_Water_bank        sql.NullString
			_Wl_tele_datetime  sql.NullString
			_Amphoe_name       sql.NullString
		)

		err := rows.Scan(&_Station_name, &_Tele_station_id, &_Province_name, &_Wl_msl, &_Percent, &_Tele_station_lat, &_Tele_station_long, &_Wl_tele_date, &_Wl_tele_time, &_Province_id, &_Ground_level, &_Water_bank, &_Wl_tele_datetime, &_Amphoe_name)
		if err != nil {
			return nil, err
		}

		month_thai := []string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."}

		d := strings.Split(_Wl_tele_date.String, "-")
		m, err := strconv.ParseInt(d[1], 10, 64)
		if err != nil {
			return nil, err
		}

		y, err := strconv.ParseInt(d[0], 10, 32)
		if err != nil {
			return nil, err
		}

		result := strconv.FormatInt(int64(y+543), 10)
		m = m - 1

		strDate := d[2] + " " + month_thai[m] + " " + result

		waterlevel_scale := waterlevel_setting.CompareScale(_Percent.Float64)

		//Setfomat number for percent data
		_p := fmt.Sprintf("%.2f", _Percent.Float64)

		s := &Struct_waterlevel_datas{
			Station_name:      _Station_name.String,
			Tele_station_id:   _Tele_station_id.String,
			Province_name:     _Province_name.String,
			Wl_msl:            _Wl_msl.String,
			Percent:           _p + " %",
			Situation:         waterlevel_scale.Situation,
			Situation_color:   waterlevel_scale.Color,
			Tele_station_lat:  _Tele_station_lat.String,
			Tele_station_long: _Tele_station_long.String,
			Wl_tele_date:      _Wl_tele_date.String,
			Wl_tele_time:      _Wl_tele_time.String,
			Wl_tele_date_thai: strDate,
			Province_id:       _Province_id.String,
			Ground_level:      _Ground_level.String,
			Water_bank:        _Water_bank.String,
			Warn_level:        "9999.99",
			Wl_tele_datetime:  _Wl_tele_datetime.String + " น.",
			Amphoe_name:       _Amphoe_name.String,
			Before_level:      "fa fa-arrow-up",
		}

		lt_datas = append(lt_datas, s)

	}

	current_date := time.Now().Format("01/02/2006 15:04")

	map_data := &Struct_waterlevel_map{
		Image_url:   "http://www.nhc.in.th/product/latest/img/rain24.jpg",
		Image_time:  current_date + " น.",
		Detail_link: "http://www.thaiwater.net",
	}

	lt_map = append(lt_map, map_data)

	rs = &Struct_wl_latest_list_prov{
		Wl_datas: lt_datas,
		Map:      lt_map,
	}

	return rs, err

}

type Struct_wl_latest_list struct {
	Max []*Struct_waterlevel_list_datas `json:"max"`
	Min []*Struct_waterlevel_list_datas `json:"min"`
	Map []*Struct_waterlevel_list_map   `json:"map"`
}

type Struct_waterlevel_list_map struct {
	Image_url   string `json:"image_url"`   //example:"http://www.nhc.in.th/product/latest/img/rain24.jpg"
	Image_time  string `json:"image_time"`  //example:"15:00 น."
	Detail_link string `json:"detail_link"` //example:"http://www.nhc.in.th/web/index.php?model=telemetering&view=weather"
}

type Struct_waterlevel_list_datas struct {
	Data_type         string `json:"data_type"`         //example:"Min"
	Station_name      string `json:"station_name"`      //example:`สถานีห้วยแร้ง จ.ตราด`
	Tele_station_id   string `json:"tele_station_id"`   //example:`telehaii0912`
	Province_name     string `json:"province_name"`     //example:`ตราด`
	Wl_msl            string `json:"wl_msl"`            //example:`2.21`
	Percent           string `json:"percent"`           //example:`92.20 %`
	Situation         string `json:"situation"`         //example:`น้ำมาก`
	Situation_color   string `json:"situation_color"`   //example:`#003CFA`
	Tele_station_lat  string `json:"tele_station_lat"`  //example:`12.325200`
	Tele_station_long string `json:"tele_station_long"` //example:`102.500000`
	Wl_tele_date      string `json:"wl_tele_date"`      //example:`2017-11-14`
	Wl_tele_time      string `json:"wl_tele_time"`      //example:`22:20:00`
	Wl_tele_date_thai string `json:"wl_tele_date_thai"` //example:`14 พ.ย. 2560`
	Province_id       string `json:"province_id"`       //example:`23`
	Ground_level      string `json:"ground_level"`      //example:`-14.627`
	Water_bank        string `json:"water_bank"`        //example:`3.634`
	Warn_level        string `json:"warn_level"`        //example:`9999.99`
	Wl_tele_datetime  string `json:"wl_tele_datetime"`  //example:`2017-11-14 22:20 น.`
}

func Get_waterlevel_latest_list() (*Struct_wl_latest_list, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// ใน table cache.latest_waterlevel ได้คำนวน waterlevel_msl + offset มาให้แล้ว
	var q string = `
    SELECT concat('สถานี',st.tele_station_name ->> 'th',' จ.',pn.province_name ->> 'th') AS station_name,
       lr.tele_station_id,
       pn.province_name ->> 'th' AS province_name,
	    lr.value_current AS wl_msl,	    
	    CASE
	        WHEN lr.value_current = 999999 THEN NULL
	        WHEN (st2.water_bank - st.ground_level) != 0 THEN (
	            (lr.value_current - st.ground_level) * 100
	        ) / (st2.water_bank - st.ground_level)
	        ELSE NULL
	    END AS percent,
	    st.tele_station_lat,
	    st.tele_station_long,
	    to_char(lr.datetime_current::date, 'YYYY-MM-DD') AS wl_tele_date,
	    to_char(lr.datetime_current::TIME, 'HH24:MM:SS') AS wl_tele_time,
	    pn.province_code,
	    st.ground_level,
	    st2.water_bank,
	    to_char(lr.datetime_current,'YYYY-MM-DD HH24:MI') AS waterlevel_datetime,
	    pn.amphoe_name ->> 'th' AS amphoe_name
    FROM cache.latest_waterlevel lr
    LEFT JOIN public.m_tele_station st ON lr.tele_station_id = st.id
    LEFT JOIN public.lt_geocode pn ON st.geocode_id = pn.id
    LEFT  JOIN ignore ig ON st.id = ig.station_id ::int AND ig.data_category = 'tele_waterlevel'
    LEFT JOIN
      ( SELECT m_tele_station.id,
               left_bank,
               right_bank,
               CASE
                   WHEN left_bank > right_bank THEN right_bank
                   ELSE left_bank
               END AS water_bank
       FROM m_tele_station	         
       WHERE left_bank IS NOT NULL AND right_bank IS NOT NULL 
          
    ) st2 ON lr.tele_station_id = st2.id
    WHERE st.ground_level IS NOT NULL
      AND lr.value_current IS NOT NULL
      AND st.left_bank IS NOT NULL
      AND st.right_bank IS NOT NULL
      AND st.geocode_id IS NOT NULL
      AND lr.value_current <> 999999
      AND lr.datetime_current >= $1
      AND ( lr.qc_status IS NULL OR lr.qc_status->>'is_pass' = 'true' )
      AND (ig.is_ignore = false OR ig.is_ignore IS NULL)      
	  AND lr.type = 'tele'
    ORDER BY percent`

	waterlevel_setting := &uSetting.Struct_WaterlevelSetting{}
	err = setting.GetSystemSettingPtr("Frontend.public.waterlevel_setting", &waterlevel_setting)

	var rs *Struct_wl_latest_list = &Struct_wl_latest_list{}
	var lt_max_datas []*Struct_waterlevel_list_datas = make([]*Struct_waterlevel_list_datas, 0)
	var lt_min_datas []*Struct_waterlevel_list_datas = make([]*Struct_waterlevel_list_datas, 0)
	var lt_map []*Struct_waterlevel_list_map = make([]*Struct_waterlevel_list_map, 0)

	q1 := q + ` DESC LIMIT 20 `

	dt := time.Now().Add(-12 * time.Hour).Format("2006-01-02 15:00")

	rows, err := db.Query(q1, dt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			_Station_name      sql.NullString
			_Tele_station_id   sql.NullString
			_Province_name     sql.NullString
			_Wl_msl            sql.NullString
			_Percent           sql.NullFloat64
			_Tele_station_lat  sql.NullString
			_Tele_station_long sql.NullString
			_Wl_tele_date      sql.NullString
			_Wl_tele_time      sql.NullString
			_Province_id       sql.NullString
			_Ground_level      sql.NullString
			_Water_bank        sql.NullString
			_Wl_tele_datetime  sql.NullString
			_Amphoe_name       sql.NullString
		)

		err := rows.Scan(&_Station_name, &_Tele_station_id, &_Province_name, &_Wl_msl, &_Percent, &_Tele_station_lat, &_Tele_station_long, &_Wl_tele_date, &_Wl_tele_time, &_Province_id, &_Ground_level, &_Water_bank, &_Wl_tele_datetime, &_Amphoe_name)
		if err != nil {
			return nil, err
		}

		strDate := ConvThaiYear(_Wl_tele_date.String)

		waterlevel_scale := waterlevel_setting.CompareScale(_Percent.Float64)

		//Setfomat number for percent data
		_p := fmt.Sprintf("%.2f", _Percent.Float64)

		s := &Struct_waterlevel_list_datas{
			Data_type:         "Max",
			Station_name:      _Station_name.String,
			Tele_station_id:   _Tele_station_id.String,
			Province_name:     _Province_name.String,
			Wl_msl:            _Wl_msl.String,
			Percent:           _p + " %",
			Situation:         waterlevel_scale.Situation,
			Situation_color:   waterlevel_scale.Color,
			Tele_station_lat:  _Tele_station_lat.String,
			Tele_station_long: _Tele_station_long.String,
			Wl_tele_date:      _Wl_tele_date.String,
			Wl_tele_time:      _Wl_tele_time.String,
			Wl_tele_date_thai: strDate,
			Province_id:       _Province_id.String,
			Ground_level:      _Ground_level.String,
			Water_bank:        _Water_bank.String,
			Warn_level:        "9999.99",
			Wl_tele_datetime:  _Wl_tele_datetime.String + " น.",
		}

		lt_max_datas = append(lt_max_datas, s)

	}

	q2 := q + ` ASC LIMIT 20 `

	rows_min, err := db.Query(q2, dt)
	if err != nil {
		return nil, err
	}

	for rows_min.Next() {
		var (
			_Station_name      sql.NullString
			_Tele_station_id   sql.NullString
			_Province_name     sql.NullString
			_Wl_msl            sql.NullString
			_Percent           sql.NullFloat64
			_Tele_station_lat  sql.NullString
			_Tele_station_long sql.NullString
			_Wl_tele_date      sql.NullString
			_Wl_tele_time      sql.NullString
			_Province_id       sql.NullString
			_Ground_level      sql.NullString
			_Water_bank        sql.NullString
			_Wl_tele_datetime  sql.NullString
			_Amphoe_name       sql.NullString
		)

		err := rows_min.Scan(&_Station_name, &_Tele_station_id, &_Province_name, &_Wl_msl, &_Percent, &_Tele_station_lat, &_Tele_station_long, &_Wl_tele_date, &_Wl_tele_time, &_Province_id, &_Ground_level, &_Water_bank, &_Wl_tele_datetime, &_Amphoe_name)
		if err != nil {
			return nil, err
		}

		strDate := ConvThaiYear(_Wl_tele_date.String)

		waterlevel_scale := waterlevel_setting.CompareScale(_Percent.Float64)

		//Setfomat number for percent data
		_p := fmt.Sprintf("%.2f", _Percent.Float64)

		s := &Struct_waterlevel_list_datas{
			Data_type:         "Min",
			Station_name:      _Station_name.String,
			Tele_station_id:   _Tele_station_id.String,
			Province_name:     _Province_name.String,
			Wl_msl:            _Wl_msl.String,
			Percent:           _p + " %",
			Situation:         waterlevel_scale.Situation,
			Situation_color:   waterlevel_scale.Color,
			Tele_station_lat:  _Tele_station_lat.String,
			Tele_station_long: _Tele_station_long.String,
			Wl_tele_date:      _Wl_tele_date.String,
			Wl_tele_time:      _Wl_tele_time.String,
			Wl_tele_date_thai: strDate,
			Province_id:       _Province_id.String,
			Ground_level:      _Ground_level.String,
			Water_bank:        _Water_bank.String,
			Warn_level:        "9999.99",
			Wl_tele_datetime:  _Wl_tele_datetime.String + " น.",
		}

		lt_min_datas = append(lt_min_datas, s)

	}

	current_date := time.Now()
	cur_time := current_date.Format("02/01/2006 15.04")

	map_data := &Struct_waterlevel_list_map{
		Image_url:   "http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/image?image=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJAMg9OoFxEIkUxSEAjaioopkeZ02BlQ==",
		Image_time:  cur_time + " น.",
		Detail_link: "http://www.nhc.in.th/web/index.php?model=telemetering&view=weather",
	}

	lt_map = append(lt_map, map_data)

	rs = &Struct_wl_latest_list{
		Max: lt_max_datas,
		Min: lt_min_datas,
		Map: lt_map,
	}

	return rs, err
}

func ConvThaiYear(date string) (s string) {

	month_thai := []string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."}
	d := strings.Split(date, "-")
	m, err := strconv.ParseInt(d[1], 10, 64)
	if err != nil {
		return
	}
	y, err := strconv.ParseInt(d[0], 10, 32)
	if err != nil {
		return
	}

	result := strconv.FormatInt(int64(y+543), 10)
	m = m - 1
	var thai_year string = d[2] + " " + month_thai[m] + " " + result

	return thai_year
}

type Struct_wl_basiN_latest struct {
	Wl_basin  []*Struct_Get_wl_basin_data `json:"wl_basin"`
	Wl_region []*Struct_Get_wl_basin_data `json:"wl_region"`
}

type Struct_Get_wl_basin_data struct {
	Pos_id          int64  `json:"pos_id"`
	Name            string `json:"name"`
	Situation       string `json:"situation"`
	Situation_color string `json:"situation_color"`
}

func Get_wl_basin_data() (*Struct_wl_basiN_latest, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	cur_date := time.Now().Add(-4 * time.Hour).Format("2006-01-02 15:00")
	//  cur_date = cur_date.Format("2006-01-02")

	q := `WITH sortable AS
  ( SELECT unnest( ARRAY[ '430'::text, '610'::text, '616'::text, '630'::text, '454'::text, '653'::text, '664'::text, '683'::text, '573'::text, '302'::text, '291'::text, '359'::text, '228'::text, '238'::text, '280'::text, '89'::text, '68'::text, '71'::text ] ) AS s_key,
           unnest( ARRAY[ '0'::text, '1'::text, '2'::text, '3'::text, '4'::text, '5'::text, '6'::text, '7'::text, '8'::text, '9'::text, '10'::text, '11'::text, '12'::text, '13'::text, '14'::text, '15'::text, '16'::text, '17'::text ] ) AS s_index),
     wl_data_today AS
  ( SELECT wl_data.waterlevel_datetime,
           wl_data.tele_station_id,
           wl_data.waterlevel_m,
           wl_data.waterlevel_msl,
           wl_data.flow_rate,
           wl_data.discharge,
           wl_data.tele_basin,
           wl_data.tele_region,
           wl_data.percent,
           sortable.s_index,
           rank() OVER (PARTITION BY wl_data.tele_station_id
                        ORDER BY wl_data.waterlevel_datetime DESC) AS r
   FROM
     ( SELECT waterlevel_datetime,
              tele_waterlevel.tele_station_id,
              waterlevel_m,
              waterlevel_msl,
              flow_rate,
              discharge,
              CASE
                  WHEN tele_waterlevel.tele_station_id IN ('430',
                                                           '610',
                                                           '616') THEN 'PING'
                  WHEN tele_waterlevel.tele_station_id IN ('630',
                                                           '454') THEN 'WAN'
                  WHEN tele_waterlevel.tele_station_id IN ('653',
                                                           '664') THEN 'YOM'
                  WHEN tele_waterlevel.tele_station_id IN ('683',
                                                           '573') THEN 'NAN'
                  WHEN tele_waterlevel.tele_station_id IN ('89',
                                                           '68',
                                                           '71') THEN 'CHAOPRAYA'
                  WHEN tele_waterlevel.tele_station_id IN ('302') THEN 'CHI_UPPER'
                  WHEN tele_waterlevel.tele_station_id IN ('359',
                                                           '291') THEN 'CHI_LOWER'
                  WHEN tele_waterlevel.tele_station_id IN ('228',
                                                           '238') THEN 'MUN_UPPER'
                  WHEN tele_waterlevel.tele_station_id IN ('280') THEN 'MUN_LOWER'
                  ELSE 'OTHER'
              END AS tele_basin,
              CASE
                  WHEN tele_waterlevel.tele_station_id IN ('430',
                                                           '610',
                                                           '616',
                                                           '630',
                                                           '454',
                                                           '653',
                                                           '664',
                                                           '683',
                                                           '573') THEN 'NORTH'
                  WHEN tele_waterlevel.tele_station_id IN ('89',
                                                           '68',
                                                           '71') THEN 'CENTRAL'
                  WHEN tele_waterlevel.tele_station_id IN ('302',
                                                           '359',
                                                           '291',
                                                           '228',
                                                           '238',
                                                           '280') THEN 'NORTHEASTERN'
                  ELSE 'OTHER'
              END AS tele_region,
              (((tele_waterlevel.waterlevel_msl - m_tele_station.ground_level) / (LEAST(m_tele_station.left_bank, m_tele_station.right_bank) - m_tele_station.ground_level)) * (100)::double precision) AS percent
      FROM tele_waterlevel
      LEFT OUTER JOIN m_tele_station ON m_tele_station.id = tele_waterlevel.tele_station_id
      WHERE waterlevel_datetime >= $1
        AND waterlevel_msl IS NOT NULL
    AND waterlevel_msl < 9999 
    AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' ) ) wl_data
   LEFT JOIN sortable ON wl_data.tele_station_id::text = sortable.s_key
   WHERE(wl_data.tele_station_id ) in (430,610, 616, 630, 454, 653, 664,
                                                                    683, 573, 302, 359, 291, 228,
                                                                   238, 280, 89, 68, 71)  ),
     filters AS
  ( SELECT wl_data_today.waterlevel_datetime,
           wl_data_today.tele_station_id,
           wl_data_today.waterlevel_msl,
           wl_data_today.waterlevel_m,
           wl_data_today.flow_rate,
           wl_data_today.discharge,
           wl_data_today.tele_basin,
           wl_data_today.tele_region,
           wl_data_today.percent,
           rank() OVER (PARTITION BY wl_data_today.tele_basin
                        ORDER BY wl_data_today.s_index) AS r_basin,
                       rank() OVER (PARTITION BY wl_data_today.tele_region
                                    ORDER BY wl_data_today.s_index) AS r_region
   FROM wl_data_today
   WHERE wl_data_today.r = 1 ),
     gbasin AS
  ( SELECT unnest(ARRAY['PING'::text, 'WAN'::text, 'YOM'::text, 'NAN'::text, 'CHAOPRAYA'::text, 'CHI_UPPER'::text, 'CHI_LOWER'::text, 'MUN_UPPER'::text, 'MUN_LOWER'::text]) AS basin_name,
           unnest(ARRAY['ลุ่มน้ำปิง'::text, 'ลุ่มน้ำวัง'::text, 'ลุ่มน้ำยม'::text, 'ลุ่มน้ำน่าน'::text, 'ลุ่มน้ำเจ้าพระยา'::text, 'ลุ่มน้ำชีตอนบน'::text, 'ลุ่มน้ำชีตอนล่าง'::text, 'ลุ่มน้ำมูลตอนบน'::text, 'ลุ่มน้ำมูลตอนล่าง'::text]) AS basin_name_th ),
     gregion AS
  ( SELECT unnest(ARRAY['NORTH'::text, 'NORTHEASTERN'::text, 'CENTRAL'::text]) AS region_name,
           unnest(ARRAY['ลุ่มน้ำภาคเหนือ'::text, 'ลุ่มน้ำภาคอีสาน'::text, 'ลุ่มน้ำเจ้าพระยา'::text]) AS region_name_th )
SELECT row_number() OVER () AS pos_id,
                         gbasin.basin_name_th AS name, -- wl_text_scale(filters.percent) AS situation,
 -- wl_color_scale(filters.percent) AS situation_color,
 filters.percent
FROM gbasin
LEFT JOIN filters ON gbasin.basin_name = filters.tele_basin
AND (filters.r_basin = 1
     OR filters.r_basin = NULL::bigint)
UNION ALL
SELECT row_number() OVER () AS pos_id,
                         gregion.region_name_th AS name, -- wl_text_scale(filters.percent) AS situation,
 -- wl_color_scale(filters.percent) AS situation_color,
 filters.percent
FROM gregion
LEFT JOIN filters ON gregion.region_name = filters.tele_region
AND (filters.r_region = 1
     OR filters.r_region = NULL::bigint)`

	waterlevel_setting := &uSetting.Struct_WaterlevelSetting{}
	err = setting.GetSystemSettingPtr("Frontend.public.waterlevel_setting", &waterlevel_setting)

	rs := &Struct_wl_basiN_latest{}
	wl_basin := make([]*Struct_Get_wl_basin_data, 0)
	wl_region := make([]*Struct_Get_wl_basin_data, 0)

	//var rs *Struct_wl_latest_list = &Struct_wl_latest_list{}
	//var lt_max_datas []*Struct_waterlevel_list_datas = make([]*Struct_waterlevel_list_datas, 0)
	//var lt_min_datas []*Struct_waterlevel_list_datas = make([]*Struct_waterlevel_list_datas, 0)
	//var lt_map []*Struct_waterlevel_list_map = make([]*Struct_waterlevel_list_map, 0)

	rows, err := db.Query(q, cur_date)
	if err != nil {
		return nil, err
	}

	count := 0

	for rows.Next() {
		var (
			_Pos_id  sql.NullInt64
			_Name    sql.NullString
			_Percent sql.NullFloat64
		)

		err := rows.Scan(&_Pos_id, &_Name, &_Percent)
		if err != nil {
			return nil, err
		}

		waterlevel_scale := waterlevel_setting.CompareScale(_Percent.Float64)

		s := &Struct_Get_wl_basin_data{
			Pos_id: _Pos_id.Int64,
			Name:   _Name.String,
		}

		//Check percent data is nil
		if _Percent.Valid {
			s.Situation = waterlevel_scale.Situation
			s.Situation_color = waterlevel_scale.Color
		} else {
			s.Situation_color = "#BDBDBD"
			s.Situation = "ไม่มีข้อมูล"
		}

		if count < 9 {
			wl_basin = append(wl_basin, s)
			rs.Wl_basin = wl_basin
		} else {
			wl_region = append(wl_region, s)
			rs.Wl_region = wl_region
		}

		count += 1
	}

	return rs, nil
}
