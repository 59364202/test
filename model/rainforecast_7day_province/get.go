// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Permporn Kuibumrung <permporn@haii.or.th>
//

// Package rainforecast_7day_province is a model for public.rainforecast_7day table. This table store rainforecast.
package rainforecast_7day_province

import (
	"database/sql"
	_"encoding/json"
	"time"
	"fmt"
	"haii.or.th/api/server/model/setting"

	"haii.or.th/api/thaiwater30/util/float"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	ข้อมูลคาดการณ์ฝน 3 วัน
//	Parameters:
//		prov_id
//			รหัสจังหวัด
//	Return:
//		array ข้อมูลคาดการณ์ฝน 3 วัน
func GetRainForecast3Day(prov_id string) ([]*Struct_RainForecastData, error) {
	if prov_id == "" {
		return nil, errors.New("no prov_id")
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := SQL_SelectRainforecast
	row, err := db.Query(q, prov_id)
	if err != nil {
		return nil, err
	}
	fmt.Println(row)
	rs := make([]*Struct_RainForecastData, 0)
	
	for row.Next() {
		var (
			_day					sql.NullString
			_rainforecast_datetime  time.Time
			_province_code          sql.NullString
			_province_name          sql.NullString
			_rainforecast_value     sql.NullFloat64
			_rainforecast_leveltext sql.NullString
		)
		err = row.Scan(&_day,&_rainforecast_datetime, &_province_code, &_province_name, &_rainforecast_value, &_rainforecast_leveltext)
		if err != nil {
			return nil, err
		}
		d := &Struct_RainForecastData{
			Forecast_day:    	_day.String,
			Forecast_datetime:  _rainforecast_datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat")),
			Province_id:    	_province_code.String,
			Province_name:  	_province_name.String,
			Rainfall:       	float.String(_rainforecast_value.Float64, 4),
			Rainfall_text:  	_rainforecast_leveltext.String,
		}
		rs = append(rs, d)
	}
	return rs, nil
}