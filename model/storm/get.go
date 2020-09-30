// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package storm is a model for public.storm table. This table store storm information.
package storm

import (
	"database/sql"
	"strings"
	"fmt"
	"strconv"
	"time"
	// "encoding/json"
	"haii.or.th/api/server/model/setting"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/thaiwater30/util/sqltime"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	
	// model_setting "haii.or.th/api/server/model/setting"
	//	"encoding/json"
	//	"haii.or.th/api/server/model/setting"
)

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

//	get storm ใช้ใน service public
//หาข้อมูลพายุ ย้อนหลัง 1 วัน
//	Return:
//		Array Struct_Strom
func GetStormCurrentDate() ([]*Struct_Strom, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//	fmt.Println(SQL_GetStormCurrentDate())
	row, err := db.Query(SQL_GetStormCurrentDate())
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	result := make([]*Struct_Strom, 0)
	tempString := ""
	var Storm *Struct_Strom
	for row.Next() {
		var (
			_storm_datetime       sqltime.NullTime
			_storm_lat            sql.NullFloat64
			_storm_direction_lat  sql.NullString
			_storm_long           sql.NullFloat64
			_storm_direction_long sql.NullString
			_storm_name           sql.NullString
			_storm_pressure       sql.NullString
			_storm_wind           sql.NullString
		)
		err = row.Scan(&_storm_datetime, &_storm_lat, &_storm_direction_lat, &_storm_long, &_storm_direction_long, &_storm_name, &_storm_pressure, &_storm_wind)
		if err != nil {
			return nil, errors.Repack(err)
		}

		if tempString != _storm_name.String {
			tempString = _storm_name.String

			Storm = &Struct_Strom{}
			Storm.Storm_Name = _storm_name.String
			result = append(result, Storm)
		}

		obj := &Struct_Strom_Data{}
		if _storm_datetime.Valid {
			obj.Storm_Datetime = _storm_datetime.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		}

		obj.Storm_Lat, _ = _storm_lat.Value()
		obj.Storm_Direction_Lat, _ = _storm_direction_lat.Value()
		obj.Storm_Long, _ = _storm_long.Value()
		obj.Storm_Direction_Long, _ = _storm_direction_long.Value()
		obj.Storm_Name = _storm_name.String
		obj.Storm_Pressure = _storm_pressure.String
		obj.Storm_Wind = _storm_wind.String

		Storm.Storm_Data = append(Storm.Storm_Data, obj)
	}

	return result, nil
}

type Struct_storm_latest struct {
	Status_flag         bool   `json:"status_flag"`         //example:`false`
	Storm_name          string `json:"storm_name"`          //example:`ไม่มีพายุ`
	Storm_wind          string `json:"storm_wind"`          //example:`0`
	Storm_warning_level int    `json:"storm_warning_level"` //example:0
	Storm_warning_color string `json:"storm_warning_color"` //example:`#663300`
	Storm_type          string `json:"storm_type"`          //example:``
	Storm_distant		string `json:"storm_distant"`       // example:`100` ระยะห่างจากประเทศ หน่วย กม.
}

//type Struct_storm_scale struct {
//	Category   string `json:"category"`
//	Color      string `json:"color"`
//	Kmh_text   string `json:"kmh_text"`
//	Knots_text string `json:"knots_text"`
//	Mph_text   string `json:"mph_text"`
//	Operator   string `josn:"operator"`
//	Scale_text string `josn:"scale_text"`
//	Strength   string `json:"strength"`
//	term       string `json:"term"`
//}

// ใช้ใน mobile service
func Get_storm_latest() ([]*Struct_storm_latest, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	dt := time.Now().AddDate(0, 0, -1) // CURRENT_TIMESTAMP - interval '1 day'

	var q string = `
	SELECT 	
		 storm_name,
		 storm_wind,
		 CASE WHEN storm_long > 106 THEN GREATEST(ROUND((storm_long - 106)*108), 0)
			  WHEN storm_long <= 97 THEN GREATEST(ROUND((97 - storm_long)*108), 0)
				ELSE 0
		  END AS distant
FROM (
	SELECT 
		 ROW_NUMBER() OVER(partition by storm_name ORDER BY storm_datetime DESC) as row_number,*
	FROM storm
	WHERE deleted_at = To_timestamp(0)
		AND storm_name :: text IN (
									SELECT DISTINCT( storm_name :: text )
									FROM   storm
									WHERE  storm_datetime > '` + dt.Format(time.RFC3339) + `' AND deleted_at = To_timestamp(0)
						)
		AND (storm_long >= 91 AND storm_long <= 125.3)
		AND (storm_lat >= 5.4 AND storm_lat <= 21.9)
		AND extract(year from storm_datetime) = '` + strconv.Itoa(dt.Year()) + `'
		AND storm_datetime > '` + dt.Format(time.RFC3339) + `'
		AND deleted_at = To_timestamp(0)
	ORDER BY storm_name :: text,storm_datetime
) storm_all
WHERE row_number = 1 
ORDER BY distant
LIMIT 1
`

	//	a := make([]*uSetting.Struct_StormSetting, 0)
	storm_setting := &uSetting.Arr_Struct_StormSetting{}
	err = setting.GetSystemSettingPtr("Frontend.public.storm_setting", &storm_setting)
	if err != nil {
		return nil, err
	}

	//	storm_setting := &uSetting.Arr_Struct_StormSetting{}
	//	storm_setting.Setting = a

	rows, err := db.Query(q)

	var rs []*Struct_storm_latest = make([]*Struct_storm_latest, 0)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			_storm_name sql.NullString
			_storm_wind sql.NullString
			_storm_distant sql.NullString
		)

		err := rows.Scan(&_storm_name, &_storm_wind, &_storm_distant)
		if err != nil {
			return nil, err
		}

		s := &Struct_storm_latest{
			Status_flag: true,
			Storm_name:  _storm_name.String,
			Storm_wind:  _storm_wind.String,
			Storm_distant:	_storm_distant.String,
		}

		//			data_storm := storm_setting.CompareSetting(_storm_wind)
		data_storm_setting := storm_setting.CompareSetting(_storm_wind.String)
		if data_storm_setting != nil {
			s.Storm_warning_level = data_storm_setting.Level
			s.Storm_warning_color = data_storm_setting.Color
			s.Storm_type = data_storm_setting.Category
		}
		rs = append(rs, s)
	}
	if len(rs) == 0 {
		s := &Struct_storm_latest{
			Status_flag:         false,
			Storm_name:          "ไม่มีพายุ",
			Storm_wind:          "0",
			Storm_warning_level: 0,
			Storm_warning_color: "#663300",
			Storm_type:          " ",
			Storm_distant:		 "0",
		}
		rs = append(rs, s)
	}

	return rs, err
}

// ------ storm history -----
// @Summary			get storm name
// @Description		get strom name from criteria (name, month [1-12], year 4 digits)
// @Parameter		StructStormHistoryParam
// @Response		200	[]string
func GetStormName(param *StructStormHistoryParam) ([]string, error) {
	// connection
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// fmt.Println(param.Year)
	// append sql criteria in slice
	sqlLimit := ""
	where := ""
	var sqlWhereArray []string
	
	if param.Name != "" {
		where = " storm_name LIKE '%" + strings.ToUpper(param.Name) + "%'"
		sqlWhereArray = append(sqlWhereArray, where)
	}
	if param.Month != 0 {
		where = " EXTRACT(month FROM start_date) = " + strconv.Itoa(param.Month)
		sqlWhereArray = append(sqlWhereArray, where)
	}
	if param.Year != 0 {
		where = " EXTRACT(year FROM start_date) = " + strconv.Itoa(param.Year)
		sqlWhereArray = append(sqlWhereArray, where)
	}
	
	// add AND between each criteria
	// fmt.Println(strings.Join(sqlWhereArray, " AND"))
	sqlWhere := strings.Join(sqlWhereArray, " AND")
	
	// add where clause
	if len(sqlWhere) > 0 {
		sqlWhere = " WHERE " + sqlWhere
	} else {
		sqlLimit = " LIMIT 10"
	}
	
	// order by
	sqlOrderBy := " ORDER BY start_date DESC"
	
	// get storm name from criteria
	rows, err := db.Query(sqlGetStormPeriod + sqlWhere + sqlOrderBy + sqlLimit)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	
	// create data structure
	// var (
	//	rs 		[]*StructStormPeriod
	//	storm  	*StructStormPeriod
	// )
	// rs = make([]*StructStormPeriod, 0)
	
	// loop through data
	var stromNameList []string
	for rows.Next() {
		var _storm_name sql.NullString
		
		// scan query result to variable
		err := rows.Scan(&_storm_name)
		if err != nil {
			return nil, err
		}
		
		// add data rows to struct
		stromNameList = append(stromNameList, _storm_name.String)
	}
	
	return stromNameList, err
}

// @Summary			get storm history data for give storm
// @Description		get strom data from array of storm name
// @Parameter		[]string
// @Response		200	StructStormHistory
func GetStormHistoryData(stormName []string) ([]*StructStormHistory, error) {
	// datetime format for this data
	strDatetimeFormat := "2006-01-02 15:04"
	
	// connection
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	
	// create output structure
	var (
		stormObj	[]*StructStormHistory
		stormHistory *StructStormHistory
	)
	stormObj = make([]*StructStormHistory, 0)
	
	// loop through each storm
	for _, name := range stormName {
	    fmt.Println("Storm -", name)
	    
	    // set where criteria
	    sqlWhere := "WHERE storm_name_alias = '" + name + "'"
	    
	    // get storm data for each storm, add cache if possible
	    // fmt.Println(sqlGetStormHistory + sqlWhere)
		_rows, err := db.Query(sqlGetStormHistory + sqlWhere)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}
		
	    // create data structure
	    var (
			stormDataObj  	*StructStormHistoryData
			
			_id 					sql.NullInt64
			_storm_datetime         time.Time
			_storm_lat              sql.NullFloat64
			_storm_long             sql.NullFloat64
			_storm_directionlat     sql.NullString
			_storm_directionlong    sql.NullString
			_storm_pressure			sql.NullString
			_strom_wind				sql.NullFloat64
			_storm_color			sql.NullString
			_storm_to_lat           sql.NullFloat64
			_storm_to_long          sql.NullFloat64
		 )
	 
		 rs := make([]*StructStormHistoryData, 0)

	    // loop through data
	    for _rows.Next() {
	    	// scan query result to variable
			err := _rows.Scan(&_id, &_storm_datetime, &_storm_lat, &_storm_long,
				&_storm_directionlat, &_storm_directionlong, 
				&_storm_pressure, &_strom_wind, &_storm_color,
				&_storm_to_lat, &_storm_to_long)
			
			if err != nil {
				return nil, err
			}
			
			stormDataObj = &StructStormHistoryData{}
			stormDataObj.Id = _id.Int64
			stormDataObj.StormDatetime = _storm_datetime.Format(strDatetimeFormat)
			stormDataObj.StormLat = ValidData(_storm_lat.Valid, _storm_lat.Float64)
			stormDataObj.StormLong = ValidData(_storm_long.Valid, _storm_long.Float64)
			stormDataObj.StormDirectionLat = _storm_directionlat.String
			stormDataObj.StormDirectionLong = _storm_directionlong.String
			stormDataObj.StormPressure = _storm_pressure.String
			stormDataObj.StormWind = ValidData(_strom_wind.Valid, _strom_wind.Float64)
			stormDataObj.StormColor = _storm_color.String
			stormDataObj.StormToLat = ValidData(_storm_to_lat.Valid, _storm_to_lat.Float64)
			stormDataObj.StormToLong = ValidData(_storm_to_long.Valid, _storm_to_long.Float64)
			
			// push data rows to struct
			rs = append(rs, stormDataObj)

			// debug ok
			// fmt.Printf("%+v\n", stormDataObj)
	    }
		
	    // append data to structure
	    stormHistory = &StructStormHistory{}
	    stormHistory.StormName = name
	    stormHistory.Data = rs
		
	    stormObj = append(stormObj, stormHistory)
	}
	
	return stormObj, err
}

// @Summary			get storm history main
// @Description		get all storm name and history data from given criteria
// @Parameter		StructStormHistoryParam
// @Response		200	[]StructStormHistory
func GetStormHistory(param *StructStormHistoryParam) ([]*StructStormHistory, error) {
	// get storm name
	stormNameArray, err := GetStormName(param);
	if err != nil {
		return nil, err
	}
	
	// get strom history data by storm name
	stormHistory, err := GetStormHistoryData(stormNameArray)
	if err != nil {
		return nil, err
	}
	
	return stormHistory, err
}
// ------ END storm history -----