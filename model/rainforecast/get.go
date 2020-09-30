// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainforecast is a model for public.rainforecast table. This table store rainforecast.
package rainforecast

import (
	"database/sql"
	"encoding/json"
	"time"

	"haii.or.th/api/server/model/setting"

	"haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/thaiwater30/util/float"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	คาดการณ์ฝนล่วงหน้า
//	Return:
//		Array Struct_RainforecastCurrentDate
func GetRainforecastCurrentDate() ([]*Struct_RainforecastCurrentDate, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		obj *Struct_RainforecastCurrentDate

		_province_code sql.NullString
		_province_name sql.NullString
		_value         sql.NullFloat64
		_zone_code     sql.NullString
	)
	data := make([]*Struct_RainforecastCurrentDate, 0)

	row, err := db.Query(SQL_SelectRainforecastCurrentDate())
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_province_code, &_province_name, &_value, &_zone_code)
		if err != nil {
			return nil, err
		}

		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}

		obj = &Struct_RainforecastCurrentDate{Province_Code: _province_code.String, Province_Name: json.RawMessage(_province_name.String),
			Value: _value.Float64, Zone_Code: _zone_code.String}
		data = append(data, obj)
	}

	return data, nil
}

//	เตือนภัยคาดการณ์ฝนล่วงหน้า
//	Return:
//		Array Struct_Rainforecast
func GetRainforecastWarning() ([]*Struct_Rainforecast, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	data := make([]*Struct_Rainforecast, 0)

	row, err := db.Query(SQL_SelectRainforecast, time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var (
			obj *Struct_Rainforecast

			provinceCode sql.NullString
			provinceName pqx.JSONRaw
			level        sql.NullInt64
		)
		err = row.Scan(&provinceCode, &provinceName, &level)
		if err != nil {
			return nil, err
		}

		obj = &Struct_Rainforecast{
			Province_Code:      provinceCode.String,
			Province_Name:      provinceName.JSONPtr(),
			Rainforecast_Level: level.Int64,
		}
		data = append(data, obj)
	}

	return data, nil
}

//	ข้อมูลคาดการณ์ฝน 3 วัน
//	Parameters:
//		prov_id
//			รหัสจังหวัด
//	Return:
//		array ข้อมูลคาดการณ์ฝน 3 วัน
func GetRainForecast3Day(prov_id string) ([]*Struct_RainForecast3Day, error) {
	if prov_id == "" {
		return nil, errors.New("no prov_id")
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := `
SELECT r.rainforecast_datetime 
       , r.rainforecast_leveltext 
       , r.rainforecast_level 
FROM   rainforecast r 
       INNER JOIN lt_geocode g 
               ON r.geocode_id = g.id 
WHERE  r.rainforecast_datetime >= $2 
       AND r.deleted_at = To_timestamp(0) 
       AND g.province_code = $1 
ORDER  BY r.rainforecast_datetime 
LIMIT  3 
	`
	row, err := db.Query(q, prov_id, time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	next3Day := time.Now().AddDate(0, 0, 2)
	result := make([]*Struct_RainForecast3Day, 0)
	result = append(result, &Struct_RainForecast3Day{"วันนี้", "ไม่มีข้อมูล", 0})
	result = append(result, &Struct_RainForecast3Day{"วันพรุ่งนี้", "ไม่มีข้อมูล", 0})
	result = append(result, &Struct_RainForecast3Day{next3Day.Format("02 ") + datetime.MonthTHShort(next3Day.Month()), "ไม่มีข้อมูล", 0})
	var i int = 0
	for row.Next() {
		var (
			_rainforecast_datetime  time.Time
			_rainforecast_leveltext sql.NullString
			_rainforecast_level     sql.NullInt64
		)
		err = row.Scan(&_rainforecast_datetime, &_rainforecast_leveltext, &_rainforecast_level)
		if err != nil {
			return nil, err
		}

		d := result[i]
		d.Status = _rainforecast_leveltext.String
		d.Status_id = _rainforecast_level.Int64

		result[i] = d
		i++
	}
	return result, nil
}

func GetRainForcastRegion() ([]*Struct_RainForcastRegion, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	//	q := `
	//SELECT r.rainforecast_datetime     AS date
	//       , g.tmd_area_code           AS region_id
	//       , Max(r.rainforecast_level) AS status
	//FROM   rainforecast r
	//       inner join lt_geocode g
	//               ON r.geocode_id = g.id
	//WHERE  r.rainforecast_datetime >= current_date
	//       AND r.deleted_at = To_timestamp(0)
	//GROUP  BY g.tmd_area_code
	//          , r.rainforecast_datetime
	//ORDER  BY r.rainforecast_datetime
	//          , g.tmd_area_code :: INT
	//	`
	q := `
SELECT * 
FROM   (SELECT rainforecast_datetime       AS date 
               , 1                         AS region_id 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.tmd_area_code = '3' 
        GROUP  BY rainforecast_datetime 
        UNION ALL 
        SELECT rainforecast_datetime 
               , 2 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.tmd_area_code = '4' 
        GROUP  BY rainforecast_datetime 
        UNION ALL 
        SELECT rainforecast_datetime 
               , 3 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.tmd_area_code = '1' 
        GROUP  BY rainforecast_datetime 
        UNION ALL 
        SELECT rainforecast_datetime 
               , 4 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.tmd_area_code = '2' 
        GROUP  BY rainforecast_datetime 
        UNION ALL 
        SELECT rainforecast_datetime 
               , 5 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.tmd_area_code = '5' 
        GROUP  BY rainforecast_datetime 
        UNION ALL 
        SELECT rainforecast_datetime 
               , 6 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.tmd_area_code = '6' 
        GROUP  BY rainforecast_datetime 
        UNION ALL 
        SELECT rainforecast_datetime 
               , 7 
               , Max(r.rainforecast_level) AS status 
        FROM   rainforecast r 
               inner join lt_geocode g 
                       ON r.geocode_id = g.id 
        WHERE  r.rainforecast_datetime::date >= $1 
               AND r.deleted_at = To_timestamp(0) 
               AND g.zone_detail = '{"th":"พื้นที่จังหวัดกรุงเทพฯ และจังหวัดใกล้เคียง"}' 
GROUP  BY rainforecast_datetime)a 
ORDER  BY DATE 
          , region_id 
	`
	row, err := db.Query(q, time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	rs := make([]*Struct_RainForcastRegion, 0)
	tempDate := ""
	st := &Struct_RainForcastRegion{}
	var d *Struct_RainForcastRegion
	for row.Next() {
		var (
			_date      sql.NullString
			_region_id sql.NullInt64
			_status    sql.NullInt64
		)

		err = row.Scan(&_date, &_region_id, &_status)
		if err != nil {
			return nil, err
		}
		date := pqx.NullStringToTime(_date)
		strDate := datetime.WeekDayTH(date.Weekday()) + date.Format(" 02 ") + datetime.MonthTHShort(date.Month()) // พุธ 25 พ.ค.
		if tempDate != strDate {
			tempDate = strDate
			d = st.New(strDate)
			rs = append(rs, d)
		}
		d.AddRegion(_region_id.Int64, _status.Int64)
	}

	return rs, nil
}

//	ข้อมูลคาดการณ์ฝน 3 วัน ล่าสุด
//	Return:
//		array ข้อมูลคาดการณ์ฝน 3 วัน ล่าสุด
func GetRainForecastData() ([]*Struct_RainForecastData, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	q := `
SELECT r.rainforecast_datetime 
       , g.province_code 
       , g.province_name ->> 'th' 
       , r.rainforecast_value
       , r.rainforecast_leveltext 
       , r.rainforecast_level 
FROM   rainforecast r 
       inner join lt_geocode g 
               ON r.geocode_id = g.id 
WHERE  r.rainforecast_datetime >= $1 
       AND r.deleted_at = To_timestamp(0) 
ORDER  BY g.province_code :: INT 
          , r.rainforecast_datetime 
	`
	row, err := db.Query(q, time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	rs := make([]*Struct_RainForecastData, 0)
	for row.Next() {
		var (
			_rainforecast_datetime  sql.NullString
			_province_code          sql.NullString
			_province_name          sql.NullString
			_rainforecast_value     sql.NullFloat64
			_rainforecast_leveltext sql.NullString
			_rainforecast_level     sql.NullInt64
		)
		err = row.Scan(&_rainforecast_datetime, &_province_code, &_province_name, &_rainforecast_value, &_rainforecast_leveltext, &_rainforecast_level)
		if err != nil {
			return nil, err
		}
		d := &Struct_RainForecastData{
			Forecast_date:  pqx.NullStringToTime(_rainforecast_datetime).Format(setting.GetSystemSetting("setting.Default.DateFormat")),
			Province_id:    _province_code.String,
			Province_name:  _province_name.String,
			Rainfall:       float.String(_rainforecast_value.Float64, 4),
			Rainfall_text:  _rainforecast_leveltext.String,
			Rainfall_level: _rainforecast_level.Int64,
		}
		rs = append(rs, d)
	}

	return rs, nil
}
