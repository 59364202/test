// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package waterquality_station is a model for public.waterquality_station table. This table store waterquality_station.
package waterquality_station

import (
	"database/sql"
	"encoding/json"

	"haii.or.th/api/util/pqx"
	//	"log"
	"strconv"
	"strings"
)

//	gen sql, param
//	Parameters:
//		p
//			ค่า
//		field
//			ชื่อฟิลด์
//		param
//			parameter ทั้งหมด
//	Return:
//		string
//		param
func genSQLWhere(p, field string, param []interface{}) (string, []interface{}) {
	strSQLWhere := ""
	strSQLWhere += " AND ("
	arr := strings.Split(p, ",")
	lenArr := len(arr)
	for i, v := range arr {
		param = append(param, v)
		strSQLWhere += " " + field + " = $" + strconv.Itoa(len(param))
		if i < lenArr-1 {
			strSQLWhere += " OR "
		}
	}
	strSQLWhere += ")"

	return strSQLWhere, param
}

//	get waterquality station
//	Parameters:
//		area_code
//			รหัสภาค
//		province_code
//			รหัสจังหัวด
//		amphoe_code
//			รหัสอำเภอ
//		tumbon_code
//			รหัสตำบล
//	Return:
//		Array Struct_Station
func get_WaterQualityStaion(area_code, province_code, amphoe_code, tumbon_code string) ([]*Struct_Station, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		strSql    string
		param     []interface{}
		SQL_where string = ""
		s         string

		data []*Struct_Station
		ws   *Struct_Station

		row                           *sql.Rows
		_id                           int64
		_agency_id                    int64
		_geocode_id                   int64
		_waterquality_station_name    sql.NullString
		_waterquality_station_lat     sql.NullFloat64
		_waterquality_station_long    sql.NullFloat64
		_waterquality_station_oldcode sql.NullString
		_waterquality_is_active       sql.NullString
	)

	strSql = SQL_GetAllWaterQualityStation

	if area_code != "" {
		s, param = genSQLWhere(area_code, "lg.area_code", param)
		SQL_where += s
		//		SQL_where += "("
		//		arr := strings.Split(area_code, ",")
		//		lenArr := len(arr)
		//		for i, v := range arr {
		//			param = append(param, v)
		//			SQL_where += " lg.area_code = $" + strconv.Itoa(len(param))
		//			if i < lenArr-1 {
		//				strSQL_Where += " OR "
		//			}
		//		}
		//		strSQL_Where += ") AND"
	}
	if province_code != "" {
		s, param = genSQLWhere(province_code, "lg.province_code", param)
		SQL_where += s
		//		SQL_where += "("
		//		arr := strings.Split(province_code, ",")
		//		lenArr := len(arr)
		//		for i, v := range arr {
		//			param = append(param, v)
		//			SQL_where += " lg.province_code = $" + strconv.Itoa(len(param))
		//			if i < lenArr-1 {
		//				strSQL_Where += " OR "
		//			}
		//		}
		//		strSQL_Where += ") AND"
	}
	if amphoe_code != "" {
		s, param = genSQLWhere(amphoe_code, "lg.amphoe_code", param)
		SQL_where += s
		//		param = append(param, amphoe_code)
		//		SQL_where += " lg.amphoe_code = $" + strconv.Itoa(len(param)) + " AND"
	}
	if tumbon_code != "" {
		s, param = genSQLWhere(tumbon_code, "lg.tumbon_code", param)
		SQL_where += s
		//		param = append(param, tumbon_code)
		//		SQL_where += " lg.tumbon_code = $" + strconv.Itoa(len(param)) + " AND"
	}
	row, err = db.Query(strSql+SQL_where, param...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err := row.Scan(&_id, &_agency_id, &_geocode_id, &_waterquality_station_name, &_waterquality_station_lat, &_waterquality_station_long, &_waterquality_station_oldcode, &_waterquality_is_active)
		if err != nil {
			return nil, err
		}

		if !_waterquality_station_name.Valid || _waterquality_station_name.String == "" {
			_waterquality_station_name.String = "{}"
		}

		ws = &Struct_Station{Id: _id, Agency_id: _agency_id, Geocode_id: _geocode_id}
		ws.Waterquality_Station_Name = json.RawMessage(_waterquality_station_name.String)
		ws.Waterquality_Station_Lat, _ = _waterquality_station_lat.Value()
		ws.Waterquality_Station_Long, _ = _waterquality_station_long.Value()
		ws.Waterquality_Station_Oldcode = _waterquality_station_oldcode.String
		ws.Is_active = _waterquality_is_active.String

		data = append(data, ws)
	}

	return data, nil
}

//	get all waterquality station
//	Return:
//		Array Struct_Station
func Get_AllWaterQualityStaion() ([]*Struct_Station, error) {
	return get_WaterQualityStaion("", "", "", "")
}

//	get waterquality station from province code
//	Parameters:
//		province_code
//			รหัสจังหวัด
//	Return:
//		Array Struct_Station
func Get_AllWaterQualityStaion_From_ProvinceCode(province_code string) ([]*Struct_Station, error) {
	return get_WaterQualityStaion("", province_code, "", "")
}
