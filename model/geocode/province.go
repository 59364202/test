// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package geocode is a model for public.lt_geocode table. This table store geocode information.
package geocode

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

var SQL_GetAllProvince = "SELECT geocode , area_code, province_code , province_name " +
	" FROM lt_geocode " +
	" WHERE (amphoe_code is null OR trim(amphoe_code) = '') " +
	" AND  (tumbon_code is null OR trim(tumbon_code) = '') " +
	" AND province_code != '99' "
var SQL_GetAllProvince_OrderBy = " ORDER BY province_code "

//	get province
//	Parameters:
//		province_id
//			หาตาม id ของ geocode
//		province_code
//			หาตาม code ของ province
//		region_id
//			หาตาม area_code
//	Return:
//		[]Struct_Geocode
func getProvinces(province_id, province_code, region_id string) ([]*Struct_Geocode, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSql  string
		param   []interface{}
		_result *sql.Rows

		data     []*Struct_Geocode
		province *Struct_Geocode

		_geocode       string
		_area_code     string
		_province_code string
		_province_name sql.NullString
	)
	strSql = SQL_GetAllProvince
	if province_id != "" {
		//		strSql += " WHERE id = $1 "
		//		param = province_id
		strSql += " WHERE "
		province_id_array := strings.Split(province_id, ",")
		cp := len(province_id_array)

		for k, v := range province_id_array {
			param = append(param, v)
			strSql += " id = $" + strconv.Itoa(len(param))
			if k < cp-1 {
				strSql += " AND"
			}
		}
	} else if province_code != "" {
		strSql += " WHERE "
		province_code_array := strings.Split(province_code, ",")
		cp := len(province_code_array)

		for k, v := range province_code_array {
			param = append(param, v)
			strSql += " province_code = $" + strconv.Itoa(len(param))
			if k < cp-1 {
				strSql += " AND"
			}
		}

	} else if region_id != "" {
		//		strSql += " WHERE region_id = $1 "
		//		param = region_id
		strSql += " WHERE "
		region_id_array := strings.Split(region_id, ",")
		cp := len(region_id_array)

		for k, v := range region_id_array {
			param = append(param, v)
			strSql += " area_code = $" + strconv.Itoa(len(param))
			if k < cp-1 {
				strSql += " AND"
			}
		}
	}
	strSql += SQL_GetAllProvince_OrderBy
	if len(param) == 0 {
		_result, err = db.Query(strSql)
	} else {
		_result, err = db.Query(strSql, param...)
	}

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data = make([]*Struct_Geocode, 0)
	for _result.Next() {
		err = _result.Scan(&_geocode, &_area_code, &_province_code, &_province_name)
		if err != nil {
			return nil, err
		}

		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}

		province = &Struct_Geocode{}
		province.Geocode = _geocode
		province.Area_code = _area_code
		province.Province_code = _province_code
		province.Province_name = json.RawMessage(_province_name.String)
		data = append(data, province)
	}
	return data, nil
}

//func GetProvinceFromId(province_id string) ([]*ProvinceStruct, error) {
//	return getProvinces(province_id, "", "")
//}

//	get province from province code
//	Parameters:
//		province_code
//			หาตาม code ของ province
//	Return:
//		[]Struct_Geocode
func GetProvinceFromCode(province_code string) ([]*Struct_Geocode, error) {
	return getProvinces("", province_code, "")
}

//	get province from region id
//	Parameters:
//		region_id
//			หาตาม area_code
//	Return:
//		[]Struct_Geocode
func GetProvinceFromRegion(region_id string) ([]*Struct_Geocode, error) {
	return getProvinces("", "", region_id)
}

//	get all province
//	Return:
//		[]Struct_Geocode
func GetProvinceAll() ([]*Struct_Geocode, error) {
	return getProvinces("", "", "")
}
