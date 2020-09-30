// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package medium_dam is a model for public.medium_dam table. This table store medium_dam information.
package medium_dam

import (
	"database/sql"
	"encoding/json"
	"log"

	//	"log"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	float "haii.or.th/api/thaiwater30/util/float"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
)

// get mediumdam group by agency
//  Parameters:
//		None
//  Return:
//		Array Struct_DamGroupByAgency
func GetDamGroupByAgency() ([]*Struct_DamGroupByAgency, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data   []*Struct_DamGroupByAgency
		objDam *Struct_DamGroupByAgency

		_dam sql.NullString

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_result *sql.Rows
	)

	//	log.Println(sqlGetDamGroupbyAgency)
	_result, err = db.Query(sqlGetDamGroupbyAgency)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_DamGroupByAgency, 0)
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_agency_id, &_agency_shortname, &_agency_name, &_dam)
		if err != nil {
			return nil, err
		}

		//Generate DamStation object
		objDam = &Struct_DamGroupByAgency{}

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		objDam.Agency = &model_agency.Struct_Agency{}
		objDam.Agency.Id = _agency_id.Int64
		objDam.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		objDam.Agency.Agency_name = json.RawMessage(_agency_name.String)

		if _dam.String == "" {
			objDam.Dam = nil
		} else {
			arrDamStationStruct := []*Struct_Dam{}
			arrDamStation := strings.Split(_dam.String, "|")
			for _, dam := range arrDamStation {
				arrDam := strings.Split(dam, "##")
				intID, err := strconv.ParseInt(arrDam[0], 10, 64)
				if err != nil {
					return nil, err
				}
				strName := arrDam[1]
				if strName == "" {
					strName = "{}"
				}
				objDamStation := &Struct_Dam{}
				objDamStation.Id = intID
				objDamStation.Dam_name = json.RawMessage(strName)
				arrDamStationStruct = append(arrDamStationStruct, objDamStation)
			}
			objDam.Dam = arrDamStationStruct
		}

		data = append(data, objDam)
	}

	return data, nil
}

func GetDamLatest(param *Struct_DamHourlyLastest_InputParam) ([]*Struct_MediumDamLatest, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02"
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data []*Struct_MediumDamLatest
		obj  *Struct_MediumDamLatest

		_id                  sql.NullInt64
		_dam_date            time.Time
		_dam_storage         sql.NullFloat64
		_dam_storage_percent sql.NullFloat64
		_dam_inflow          sql.NullFloat64
		_dam_uses_water      sql.NullFloat64
		_dam_released        sql.NullFloat64

		_dam_id   sql.NullInt64
		_dam_name sql.NullString
		_dam_lat  sql.NullFloat64
		_dam_long sql.NullFloat64

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_basin_id   sql.NullInt64
		_basin_code sql.NullInt64
		_basin_name sql.NullString

		_geocode_id    sql.NullInt64
		_geocode       sql.NullString
		_area_code     sql.NullString
		_area_name     sql.NullString
		_province_name sql.NullString
		_amphoe_name   sql.NullString
		_tumbon_name   sql.NullString
		_province_code sql.NullString

		_dam_oldcode sql.NullString

		_subbasin_id sql.NullInt64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""
	arrBasinID := []string{}
	if param.Basin_id != "" {
		arrBasinID = strings.Split(param.Basin_id, ",")
	}

	//Check Filter basin_id
	if len(arrBasinID) > 0 {
		if len(arrBasinID) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Basin_id, " "))
			sqlCmdWhere += " AND b.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrBasinID {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND b.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	arrProvinceID := []string{}
	if param.Province_id != "" {
		arrProvinceID = strings.Split(param.Province_id, ",")
	}
	//Check Filter province_id
	if len(arrProvinceID) > 0 {
		if len(arrProvinceID) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Province_id, " "))
			sqlCmdWhere += " AND province_code = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceID {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND province_code IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	if param.Region_id != "" {
		sqlCmdWhere += " AND area_code like '" + param.Region_id + "' "
	}

	if param.Dam_date != "" {
		sqlCmdWhere += " AND mediumdam_date = '" + param.Dam_date + "' "
	}

	//Query
	log.Println(SQL_GetDamLastest + SQL_GetDamLastest2 + sqlCmdWhere + SQL_GetDamLastest_OrderBy)
	_result, err = db.Query(SQL_GetDamLastest+SQL_GetDamLastest2+sqlCmdWhere+SQL_GetDamLastest_OrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_dam_date,
			&_dam_id, &_dam_name, &_dam_lat, &_dam_long,
			&_dam_inflow,
			&_dam_storage, &_dam_storage_percent,
			&_dam_uses_water,
			&_dam_released,
			&_agency_id, &_agency_shortname, &_agency_name,
			&_basin_id, &_basin_code, &_basin_name,
			&_geocode_id, &_geocode, &_area_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_province_code, &_dam_oldcode, &_subbasin_id)

		if err != nil {
			return nil, err
		}

		//Generate DamDaily Object
		/*----- DamDaily -----*/
		obj = &Struct_MediumDamLatest{}

		obj.Id = _id.Int64
		//obj.Dam_date = _dam_date.Format(model_setting.GetSystemSetting("setting.Default.DatetFormat"))
		obj.Dam_date = _dam_date.Format(strDatetimeFormat)
		obj.Dam_storage = float.TwoDigit(_dam_storage.Float64)
		obj.Dam_storage_percent, _ = _dam_storage_percent.Value()
		obj.Dam_uses_water, _ = _dam_uses_water.Value()
		obj.Dam_inflow = float.TwoDigit(_dam_inflow.Float64)
		obj.Dam_released = float.TwoDigit(_dam_released.Float64)

		/*----- Dam -----*/
		obj.Dam = &model_dam.Struct_Dam{}
		obj.Dam.Id = _dam_id.Int64
		obj.Dam.Dam_lat, _ = _dam_lat.Value()
		obj.Dam.Dam_long, _ = _dam_long.Value()
		obj.Dam.Dam_oldcode = _dam_oldcode.String
		obj.Dam.SubBasin_id = _subbasin_id.Int64

		if _dam_name.String == "" {
			_dam_name.String = "{}"
		}
		obj.Dam.Dam_name = json.RawMessage(_dam_name.String)

		/*----- Agency -----*/
		obj.Agency = &model_agency.Struct_Agency{}
		obj.Agency.Id = _agency_id.Int64

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		obj.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)

		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		obj.Agency.Agency_name = json.RawMessage(_agency_name.String)

		/*----- Basin -----*/
		obj.Basin = &model_basin.Struct_Basin{}
		obj.Basin.Id = _basin_id.Int64
		obj.Basin.Basin_code = _basin_code.Int64

		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		obj.Basin.Basin_name = json.RawMessage(_basin_name.String)

		/*----- Geocode -----*/
		obj.Geocode = &model_geocode.Struct_Geocode{}
		obj.Geocode.Id = _geocode_id.Int64
		obj.Geocode.Geocode = _geocode.String
		obj.Geocode.Area_code = _area_code.String

		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		obj.Geocode.Area_name = json.RawMessage(_area_name.String)

		obj.Geocode.Province_code = _province_code.String

		if _province_name.String == "" {
			_province_name.String = "{}"
		}
		obj.Geocode.Province_name = json.RawMessage(_province_name.String)

		if _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		obj.Geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)

		if _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		obj.Geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)

		//objDamHourly.ProvinceName = json.RawMessage(_province_name.String)
		//objDamHourly.Name = json.RawMessage(_dam_name.String)
		//objDamHourly.Datetime = _dam_date.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		//objDamHourly.Oldcode = _dam_oldcode.String
		//objDamHourly.Value = ValidData(_dam_storage.Valid, _dam_storage.Float64)
		//objDamHourly.DataID = _id.Int64
		//objDamHourly.StationID = _dam_id.Int64

		data = append(data, obj)
	}

	//Return Data
	return data, nil
}
