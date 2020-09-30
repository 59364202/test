// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Peerapong Srisom <peerapong@haii.or.th>

package dam_uses_water

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"

	"database/sql"
	"encoding/json"
	"fmt"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	get dam uses water ของ  Province_code
//	Parameters:
//		param
//			ใช้ในส่วนของ Province_code
//	Return:
//		Struct_DamUsesWater_OutputParam
func GetDamUsesWater(param *Struct_DamUsesWater_InputParam) (*Struct_DamUsesWater_OutputParam, error) {

	//Find datetime default format
	strDateFormat := model_setting.GetSystemSetting("bof.Default.DateFormat")
	if strDateFormat == "" {
		strDateFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data            []*Struct_DamUsesWater
		objDamUsesWater *Struct_DamUsesWater

		_dam_uses_water         sql.NullFloat64
		_dam_uses_water_percent sql.NullFloat64
		_dam_released           sql.NullFloat64

		_start_date sql.NullString
		_end_date   sql.NullString

		_area_code     sql.NullString
		_area_name     sql.NullString
		_province_name sql.NullString
		_province_code sql.NullString

		_result *sql.Rows
	)
	//Query
	strSql, params := Gen_SQL_GetDamUsesWater(param)
	_result, err = db.Query(strSql, params...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	// Loop data result
	data = make([]*Struct_DamUsesWater, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_start_date, &_end_date, &_dam_uses_water, &_dam_uses_water_percent, &_dam_released,
			&_province_code, &_province_name, &_area_code, &_area_name)

		if err != nil {
			return nil, err
		}

		start_date := _start_date.String[:10]
		end_date := _end_date.String[:10]

		//Generate DamUsesWater object
		objDamUsesWater = &Struct_DamUsesWater{}

		objDamUsesWater.Dam_uses_water = fmt.Sprintf("%.0f", ValidData(_dam_uses_water.Valid, _dam_uses_water.Float64))
		objDamUsesWater.Dam_uses_water_percent = fmt.Sprintf("%.0f", ValidData(_dam_uses_water_percent.Valid, _dam_uses_water_percent.Float64))
		objDamUsesWater.Dam_released = fmt.Sprintf("%.0f", ValidData(_dam_released.Valid, _dam_released.Float64))

		objDamUsesWater.Start_date = start_date
		objDamUsesWater.End_date = end_date

		/*----- Geocode -----*/
		objDamUsesWater.Geocode = &model_geocode.Struct_Geocode{}
		objDamUsesWater.Geocode.Id = 0
		objDamUsesWater.Geocode.Geocode = ""
		objDamUsesWater.Geocode.Area_code = _area_code.String
		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		objDamUsesWater.Geocode.Area_name = json.RawMessage(_area_name.String)
		objDamUsesWater.Geocode.Province_code = _province_code.String
		if _province_name.String == "" {
			_province_name.String = "{}"
		}
		objDamUsesWater.Geocode.Province_name = json.RawMessage(_province_name.String)
		objDamUsesWater.Geocode.Amphoe_name = json.RawMessage("{}")
		objDamUsesWater.Geocode.Tumbon_name = json.RawMessage("{}")

		data = append(data, objDamUsesWater)
	}

	resultData := &Struct_DamUsesWater_OutputParam{}
	resultData.Data = data
	resultData.Result = "OK"

	return resultData, err
}

// check valid data return valid return data, invalid return null
//  Parameters:
//		valid
//			boolean check valid
//		value
//			value scan from db
//  Return:
//		if true return valid return data if not invalid return null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
