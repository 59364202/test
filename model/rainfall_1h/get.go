// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainfall_1h is a model for public.rainfall_1h table. This table store rainfall_1h.
package rainfall_1h

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/sqltime"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

//ตรวจสอบค่าใน column float กรณีค่าเป็น null ให้ return type เป็น interface แทน flat64
//ตัวแปรที่รับค่าจาก db เป็น float64 ซึ่ง float64 เป็น null ไม่ได้ ถ้าใช้  column.float64 จะได้เป็น 0 ทั้งๆที่ใน db เป็น null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

//	get rainfall 1h graph
//	Parameters:
//		p
//			Param_Rainfall1h_Graph
//	Return:
//		Array Struct_Rainfall1h_Graph
func GetRainGraph(p *Param_Rainfall1h_Graph) ([]*Struct_Rainfall1h_Graph, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	if p.StationId <= 0 {
		return nil, rest.NewError(422, "invalid station id", nil)
	}

	var (
		obj *Struct_Rainfall1h_Graph

		_datetime sqltime.NullTime
		_rainfall sql.NullFloat64
	)
	strSql, param := Gen_SQL_GetRainfallGraph(p)
	//	fmt.Println(strSql)
	row, err := db.Query(strSql, param...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*Struct_Rainfall1h_Graph, 0)
	for row.Next() {
		err = row.Scan(&_datetime, &_rainfall)
		if err != nil {
			return nil, err
		}

		obj = &Struct_Rainfall1h_Graph{}
		obj.DateTime = _datetime.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		if _rainfall.Valid {
			obj.Rainfall = _rainfall.Float64
		}

		data = append(data, obj)
	}
	return data, nil
}

type Struct_bkk_rainfall1h struct {
	Tele_station_id   string `json:"tele_station_id"`   //example:`telebma0035`
	Tele_station_name string `json:"tele_station_name"` //example:`ประตูระบายน้ำคลองแสนแสบ ตอนถนนสังฆสันติสุข`
	Tele_station_lat  string `json:"tele_station_lat"`  //example:`13.855176`
	Tele_station_long string `json:"tele_station_long"` //example:`100.872918`
	Composite_pointid string `json:"composite_pointid"` //example:`226592`
	Rainfall_date     string `json:"rainfall_date"`     //example:`2017-05-22`
	Rainfall_time     string `json:"rainfall_time"`     //example:`16:00:00`
	Rainfall          string `json:"rainfall"`          //example:`26`
}

func Get_bkk_rainfall1h() ([]*Struct_bkk_rainfall1h, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var q string = `
	SELECT     t1.id, 
		t1.tele_station_name ->> 'th' AS tele_station_name, 
		t1.tele_station_lat, 
		t1.tele_station_long, 
		t3.rainfall_date, 
		t3.rainfall_time, 
		t3.rainfall1h AS rainfall 
	FROM       m_tele_station t1 
	inner join 
		( 
				SELECT   * 
				FROM     lt_geocode 
				WHERE    province_code = '10' 
				ORDER BY id) t2 
	ON         t1.geocode_id = t2.id 
	left join 
		( 
			SELECT s1.tele_station_id, 
					s1.rainfall1h, 
					s1.rainfall_datetime::timestamp::DATE AS rainfall_date, 
					s1.rainfall_datetime::timestamp::TIME AS rainfall_time 
			FROM   CACHE.latest_rainfall24h s1 
			WHERE  s1.rainfall1h IS NOT NULL 
			AND    s1.rainfall_datetime > $1
			AND    ( 
							qc_status IS NULL 
					OR     qc_status->>'is_pass' = 'true') ) t3 
	ON         t1.id = t3.tele_station_id 
	WHERE      t1.tele_station_name IS NOT NULL 
	AND        t3.rainfall1h IS NOT NULL 
	ORDER BY   t1.id`
	dt := time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04")
	rows, err := db.Query(q, dt)
	if err != nil {
		return nil, err
	}

	var rs []*Struct_bkk_rainfall1h = make([]*Struct_bkk_rainfall1h, 0)

	for rows.Next() {
		var (
			_tele_station_id   sql.NullString
			_tele_station_name sql.NullString
			_tele_station_lat  sql.NullString
			_tele_station_long sql.NullString
			_rainfall_date     sql.NullString
			_rainfall_time     sql.NullString
			_rainfall          sql.NullString
		)

		err = rows.Scan(&_tele_station_id, &_tele_station_name, &_tele_station_lat, &_tele_station_long, &_rainfall_date, &_rainfall_time, &_rainfall)

		if err != nil {
			return nil, err
		}

		s := &Struct_bkk_rainfall1h{
			Tele_station_id:   _tele_station_id.String,
			Tele_station_name: _tele_station_name.String,
			Tele_station_lat:  _tele_station_lat.String,
			Tele_station_long: _tele_station_long.String,
			Composite_pointid: "",
			Rainfall_date:     _rainfall_date.String,
			Rainfall_time:     _rainfall_time.String,
			Rainfall:          _rainfall.String,
		}

		rs = append(rs, s)
	}

	return rs, err
}

//	get max,min rainfall 1 hour
//	Parameters:
//		p
//			Param_Provinces
//	Return:
//		Array Struct_MaxMin
func GetMaxMinRainfall(p *Param_Provinces) ([]*Struct_MaxMin, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_MaxMin{}
	itf := []interface{}{}

	//Check Filter province_id
	province_condition := ""
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		province_condition += " AND "
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(p.Province_Code, " "))
			province_condition += " province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			province_condition += " province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	region_condition := ""
	if p.Region_Code != "" {
		region_condition += " AND "
		region_condition += " area_code like '" + p.Region_Code + "' "
	}

	if p.Region_Code_tmd != "" {
		region_condition += " AND "
		region_condition += " tmd_area_code like '" + p.Region_Code_tmd + "' "
	}

	q := ` WITH max AS (
select  
  dd.rainfall_datetime,
  rainfall1h as max_rainfall1h,
  d.id AS station_id,
  d.tele_station_oldcode,
  d.tele_station_name,
  d.tele_station_lat,
  d.tele_station_long,
  g.tmd_area_code AS area_code,
  g.tmd_area_name AS area_name,
  g.province_code,
  g.amphoe_code,
  g.tumbon_code,
  g.province_name,
  g.amphoe_name,
  g.tumbon_name,
  agency_name
from  latest.rainfall_1h dd
  LEFT JOIN m_tele_station d ON ((dd.tele_station_id = d.id))
  LEFT JOIN agency agt ON ((d.agency_id = agt.id))
  LEFT JOIN lt_geocode g ON ((d.geocode_id = g.id))
where tele_station_id = (select tele_station_id from  latest.rainfall_1h dd 
      LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id 
      LEFT JOIN lt_geocode g ON d.geocode_id = g.id 
    where 
      rainfall_datetime :: date = CURRENT_DATE  AND ((d.is_ignore = false) AND ((dd.qc_status IS NULL) OR ((dd.qc_status ->> 'is_pass' :: text) = 'true' :: text))` + province_condition + region_condition + `) 
    ORDER BY  rainfall1h DESC LIMIT  1)
),
min AS (
select  
  dd.rainfall_datetime as min_rainfall_datetime,
  rainfall1h as min_rainfall1h,
  d.id as min_station_id,
  d.tele_station_oldcode as min_tele_station_oldcode,
  d.tele_station_name as min_tele_station_name,
  d.tele_station_lat as min_tele_station_lat,
  d.tele_station_long as min_tele_station_long,
  g.tmd_area_code AS min_area_code,
  g.tmd_area_name AS min_area_name,
  g.province_code as min_province_code,
  g.amphoe_code as min_amphoe_code,
  g.tumbon_code as min_tumbon_code,
  g.province_name as min_province_name,
  g.amphoe_name as min_amphoe_name,
  g.tumbon_name as min_tumbon_name,
  agency_name as  min_agency_name
from  latest.rainfall_1h dd
  LEFT JOIN m_tele_station d ON ((dd.tele_station_id = d.id))
  LEFT JOIN agency agt ON ((d.agency_id = agt.id))
  LEFT JOIN lt_geocode g ON ((d.geocode_id = g.id))
where tele_station_id = (select tele_station_id from  latest.rainfall_1h dd 
      LEFT JOIN m_tele_station d ON dd.tele_station_id = d.id 
      LEFT JOIN lt_geocode g ON d.geocode_id = g.id 
    where 
      rainfall_datetime :: date = CURRENT_DATE  AND ((d.is_ignore = false) AND ((dd.qc_status IS NULL) OR ((dd.qc_status ->> 'is_pass' :: text) = 'true' :: text))` + province_condition + region_condition + `) 
    ORDER BY  rainfall1h LIMIT  1)
)
select * from max LEFT JOIN min ON 1 = 1`

	//	fmt.Println(q)
	row, err := db.Query(q)
	if err != nil {
		fmt.Println("query ", err)
		return nil, err
	}
	strDatetimeFormat := "2006-01-02 15:04"

	for row.Next() {
		var (
			station_id           int64
			tele_station_lat     sql.NullFloat64
			tele_station_long    sql.NullFloat64
			tele_station_oldcode sql.NullString
			tele_station_name    pqx.JSONRaw
			amphoe_code          sql.NullString
			amphoe_name          pqx.JSONRaw
			tumbon_code          sql.NullString
			tumbon_name          pqx.JSONRaw
			province_code        sql.NullString
			province_name        pqx.JSONRaw
			area_code            sql.NullString
			area_name            sql.NullString
			agency_name          pqx.JSONRaw
			rainfall_datetime    sql.NullString
			max_rainfall1h       sql.NullFloat64

			min_rainfall_datetime    sql.NullString
			min_rainfall1h           sql.NullFloat64
			min_station_id           int64
			min_tele_station_oldcode sql.NullString
			min_tele_station_name    pqx.JSONRaw
			min_tele_station_lat     sql.NullFloat64
			min_tele_station_long    sql.NullFloat64
			min_area_code            sql.NullString
			min_area_name            sql.NullString
			min_province_code        sql.NullString
			min_amphoe_code          sql.NullString
			min_tumbon_code          sql.NullString
			min_province_name        pqx.JSONRaw
			min_amphoe_name          pqx.JSONRaw
			min_tumbon_name          pqx.JSONRaw
			min_agency_name          pqx.JSONRaw

			d           *Struct_MaxMin                = &Struct_MaxMin{}
			t_max       *Struct_Rainfall1h            = &Struct_Rainfall1h{}
			t_min       *Struct_Rainfall1h            = &Struct_Rainfall1h{}
			station     *Struct_TeleStation           = &Struct_TeleStation{}
			agency      *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode     *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
			min_station *Struct_TeleStation           = &Struct_TeleStation{}
			min_agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			min_geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&rainfall_datetime, &max_rainfall1h, &station_id, &tele_station_oldcode, &tele_station_name, &tele_station_lat, &tele_station_long, &area_code, &area_name, &province_code, &amphoe_code, &tumbon_code, &province_name, &amphoe_name, &tumbon_name, &agency_name, &min_rainfall_datetime, &min_rainfall1h, &min_station_id, &min_tele_station_oldcode, &min_tele_station_name, &min_tele_station_lat, &min_tele_station_long, &min_area_code, &min_area_name, &min_province_code, &min_amphoe_code, &min_tumbon_code, &min_province_name, &min_amphoe_name, &min_tumbon_name, &min_agency_name)

		if err != nil {
			fmt.Println("scan error", err)
			return nil, err
		}

		// max
		t_max.Rainfall = ValidData(max_rainfall1h.Valid, max_rainfall1h.Float64)
		if rainfall_datetime.Valid {
			t_max.Time = pqx.NullStringToTime(rainfall_datetime).Format(strDatetimeFormat)
		}

		station.Id = station_id
		station.Lat = ValidData(tele_station_lat.Valid, tele_station_lat.Float64)
		station.Long = ValidData(tele_station_long.Valid, tele_station_long.Float64)
		station.Name = tele_station_name.JSON()
		station.OldCode = tele_station_oldcode.String

		agency.Agency_name = agency_name.JSON()

		geocode.Tumbon_code = tumbon_code.String
		geocode.Tumbon_name = tumbon_name.JSON()
		geocode.Amphoe_code = amphoe_code.String
		geocode.Amphoe_name = amphoe_name.JSON()
		geocode.Province_code = province_code.String
		geocode.Province_name = province_name.JSON()
		geocode.Area_code = area_code.String
		if area_name.String == "" {
			area_name.String = "{}"
		}
		geocode.Area_name = json.RawMessage(area_name.String)

		t_max.Station = station
		t_max.Agency = agency
		t_max.Geocode = geocode

		// min
		t_min.Rainfall = ValidData(min_rainfall1h.Valid, min_rainfall1h.Float64)
		if min_rainfall_datetime.Valid {
			t_min.Time = pqx.NullStringToTime(min_rainfall_datetime).Format(strDatetimeFormat)
		}

		min_station.Id = min_station_id
		min_station.Lat = ValidData(min_tele_station_lat.Valid, min_tele_station_lat.Float64)
		min_station.Long = ValidData(min_tele_station_long.Valid, min_tele_station_long.Float64)
		min_station.Name = min_tele_station_name.JSON()
		min_station.OldCode = min_tele_station_oldcode.String

		min_agency.Agency_name = min_agency_name.JSON()

		min_geocode.Tumbon_code = min_tumbon_code.String
		min_geocode.Tumbon_name = min_tumbon_name.JSON()
		min_geocode.Amphoe_code = min_amphoe_code.String
		min_geocode.Amphoe_name = min_amphoe_name.JSON()
		min_geocode.Province_code = min_province_code.String
		min_geocode.Province_name = min_province_name.JSON()
		min_geocode.Area_code = min_area_code.String
		if min_area_name.String == "" {
			min_area_name.String = "{}"
		}
		min_geocode.Area_name = json.RawMessage(min_area_name.String)

		t_min.Station = min_station
		t_min.Agency = min_agency
		t_min.Geocode = min_geocode

		d.Max = t_max
		d.Min = t_min
		r = append(r, d)
	}

	return r, nil

}
