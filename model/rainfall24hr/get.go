// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainfall24hr is a model for public.rainfall24hr table. This table store rainfall24hr.
package rainfall24hr

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	//	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	uDatetime "haii.or.th/api/thaiwater30/util/datetime"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
)

//	get rainfall from province code
//	Parameters:
//		province_code
//			รหัสจังหวัด
//	Return:
//		Array Struct_Rainfall24H
func GetRainfallThailandFromProvinceCode(province_code string) ([]*Struct_Rainfall24H, error) {
	return GetRainfallThailandDataCache(&Param_Rainfall24{Province_Code: province_code, IsDaily: true})
}

//	get rainfall
//	Return:
//		Array Struct_Rainfall24H
func GetRainfallThailand24Hr() ([]*Struct_Rainfall24H, error) {
	return GetRainfallThailandDataCache(&Param_Rainfall24{IsDaily: true})
}

//	การกระจายฝนล่วงหน้า
//	Parameters:
//		agency_id
//			array รหัสหน่วยงาน
//		date
//			วันที่
//	Return:
//		Array Struct_AdvRainDiagram
func GetAdvRainDiagram(agency_id []int64, date string) ([]*Struct_AdvRainDiagram, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		data []*Struct_AdvRainDiagram
		obj  *Struct_AdvRainDiagram
	)
	var q, p = SQL_AdvRainDiagram(agency_id, date)
	row, err := db.Query(q, p...)
	if err != nil {
		return nil, err
	}
	data = make([]*Struct_AdvRainDiagram, 0)
	for row.Next() {
		var (
			_name  pqx.JSONRaw
			_lat   sql.NullString
			_long  sql.NullString
			_value sql.NullFloat64
		)
		err = row.Scan(&_name, &_lat, &_long, &_value)
		if err != nil {
			return nil, err
		}
		obj = &Struct_AdvRainDiagram{}
		obj.Name = _name.JSONPtr()
		obj.Lat = ValidData(_lat.Valid, _lat.String)
		obj.Long = ValidData(_long.Valid, _long.String)
		obj.Value = ValidData(_value.Valid, _value.Float64)
		data = append(data, obj)
	}

	return data, nil
}

//	get หน่วยงาน, ปีที่มากที่ ปีที่น้อยที่สุด ของข้อมูล
//	Return:
//		Array Struct_AdvOnload_Agency
func GetAdvOnload_Agency() ([]*Struct_AdvOnload_Agency, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		data []*Struct_AdvOnload_Agency
		obj  *Struct_AdvOnload_Agency
	)
	row, err := db.Query(SQL_AdvOnload_Agency)
	if err != nil {
		return nil, err
	}
	data = make([]*Struct_AdvOnload_Agency, 0)
	for row.Next() {
		var (
			_id       int64
			_name     sql.NullString
			_max_year sql.NullInt64
			_min_year sql.NullInt64
		)
		err = row.Scan(&_id, &_name, &_max_year, &_min_year)
		if err != nil {
			return nil, err
		}
		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}

		obj = &Struct_AdvOnload_Agency{}
		obj.AgencyId = _id
		obj.MaxYear = ValidData(_max_year.Valid, _max_year.Int64)
		obj.MinYear = ValidData(_min_year.Valid, _min_year.Int64)
		obj.AgencyName = json.RawMessage(_name.String)
		data = append(data, obj)
	}

	return data, nil
}

//	การกระจายตัวของฝน
//	Parameters:
//		p
//			Param_AdvRainSum
//	Return:
//		Array Struct_AdvRainSum
func GetAdvRainSum(p *Param_AdvRainSum) ([]*Struct_AdvRainSum, error) {
	if p.AgencyId <= 0 { // valid agency_id
		return nil, rest.NewError(422, "invalid agency_id", nil)
	}
	if p.Boundary == "" { // valid boundary
		return nil, rest.NewError(422, "invalid boundary", nil)
	}
	if p.DateStart == "" { // valid date_start
		return nil, rest.NewError(422, "invalid date_start", nil)
	}
	if p.DateEnd == "" { // valid date_end
		return nil, rest.NewError(422, "invalid date_end", nil)
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// gen sql
	q, itf := Gen_SQLAdvRainSum(p)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}

	data := []*Struct_AdvRainSum{}
	for row.Next() {
		var (
			_x     float64
			_y     float64
			_sum   float64
			_count float64
		)
		err = row.Scan(&_x, &_y, &_sum, &_count)
		if err != nil {
			return nil, err
		}

		o := &Struct_AdvRainSum{
			X:     _x,
			Y:     _y,
			Sum:   _sum,
			Count: _count,
		}

		data = append(data, o)
	}

	return data, nil
}

//ตรวจสอบค่าใน column float กรณีค่าเป็น null ให้ return type เป็น interface แทน flat64
//ตัวแปรที่รับค่าจาก db เป็น float64 ซึ่ง float64 เป็น null ไม่ได้ ถ้าใช้  column.float64 จะได้เป็น 0 ทั้งๆที่ใน db เป็น null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

type Struct_rainfall24h struct {
	Tele_station_id   string `json:"tele_station_id"`   //example:`telebma0035`
	Tele_station_name string `json:"tele_station_name"` //example:`ประตูระบายน้ำคลองแสนแสบ ตอนถนนสังฆสันติสุข`
	Tele_station_lat  string `json:"tele_station_lat"`  //example:`13.855176`
	Tele_station_long string `json:"tele_station_long"` //example:`100.872918`
	Composite_pointid string `json:"composite_pointid"` //example:`226592`
	Rainfall_date     string `json:"rainfall_date"`     //example:`2017-05-22`
	Rainfall_time     string `json:"rainfall_time"`     //example:`16:00:00`
	Rainfall          string `json:"rainfall"`          //example:`26`
}

func Get_BKK_Rainfall24h() ([]*Struct_rainfall24h, error) {
	//Conect DB
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Query statement
	var q string = `
	SELECT     t1.id, 
			t1.tele_station_name ->> 'th' AS tele_station_name, 
			t1.tele_station_lat, 
			t1.tele_station_long, 
			t3.rainfall_date, 
			t3.rainfall_time, 
			t3.rainfall24h 
	FROM       m_tele_station t1 
	INNER JOIN 
			( 
						SELECT   * 
						FROM     lt_geocode 
						WHERE    province_code = '10' 
						ORDER BY id) t2 
	ON         t1.geocode_id = t2.id 
	LEFT JOIN 
			( 
					SELECT s1.tele_station_id, 
							s1.rainfall24h, 
							s1.rainfall_datetime::timestamp::DATE AS rainfall_date, 
							s1.rainfall_datetime::timestamp::TIME AS rainfall_time 
					FROM   CACHE.latest_rainfall24h s1 
					WHERE  s1.rainfall24h IS NOT NULL 
					--AND    s1.rainfall_datetime > current_timestamp - interval '1 day' 
					AND    s1.rainfall_datetime > $1 
					AND    ( 
									qc_status IS NULL 
							OR    	qc_status->>'is_pass' = 'true') ) t3 
	ON         t1.id = t3.tele_station_id 
	WHERE      t1.tele_station_name IS NOT NULL 
	AND        t3.rainfall24h IS NOT NULL 
	ORDER BY   t1.id
	`
	dt := time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04")
	//Query result
	rows, err := db.Query(q, dt)
	if err != nil {
		return nil, err
	}

	var rs []*Struct_rainfall24h = make([]*Struct_rainfall24h, 0)

	//Add data into struct
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

		s := &Struct_rainfall24h{
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

//	ข้อมูลฝน ล่าสุด ในจังหวัด
//	Parameters:
//		prov_id
//			รหัสจังหวัด
func GetRainfallMaxMin(prov_id string) (*Struct_RainfallMaxMin, error) {
	s := &Struct_RainfallMaxMin{}
	s.Init()
	if prov_id == "" {
		return s, nil
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	q := `
WITH zxc 
    AS (SELECT lg.province_code 
                , lg.province_name ->> 'th'         AS province_name 
                , r.rainfall24h 
                , r.rainfall_datetime 
                , m.tele_station_name ->> 'th'      AS tele_station_name 
                , m.id 
                , Rank() 
                    over ( 
                    	PARTITION BY lg.province_code 
                    	ORDER BY r.rainfall24h )      AS r_min 
                , Rank() 
                    over ( 
                    	PARTITION BY lg.province_code 
                    	ORDER BY r.rainfall24h DESC ) AS r_max 
        FROM   latest.rainfall_24h r 
                INNER JOIN m_tele_station m 
						ON r.tele_station_id = m.id 
				LEFT  JOIN ignore ig ON m.id = ig.station_id ::int
						AND ig.data_category = 'rainfall_24h'
						AND (ig.is_ignore = false OR ig.is_ignore IS NULL) 
                INNER JOIN lt_geocode lg 
                        ON m.geocode_id = lg.id 
        --WHERE  rainfall_datetime >= current_date 
        WHERE  rainfall_datetime >= $2 
                AND r.rainfall24h > 0 
                --AND m.is_ignore = FALSE 
				AND lg.province_code = $1 
				AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true')
        ORDER  BY lg.province_code) 
SELECT rainfall24h 
    	, rainfall_datetime 
    	, tele_station_name || ' จ.' || province_name
    	, id 
FROM   zxc 
WHERE  r_min = 1 
        OR r_max = 1 
ORDER  BY rainfall24h 
	`
	dt := time.Now().Format("2006-01-02")
	row, err := db.Query(q, prov_id, dt)
	if err != nil {
		return nil, err
	}

	scAll, err := uSetting.New_Struct_RainSetting()
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var (
			_rainfall24h       sql.NullFloat64
			_rainfall_datetime sql.NullString
			_tele_station_name sql.NullString
			_id                int64
		)
		err := row.Scan(&_rainfall24h, &_rainfall_datetime, &_tele_station_name, &_id)
		if err != nil {
			return nil, err
		}

		sc := scAll.CompareRain(_rainfall24h.Float64)

		if s.Rainfall_min == "-" {
			// เข้ามาครั้งแรกให้ทำฝนตกน้อยที่สุดก่อน
			s.Rainfall_min = fmt.Sprintf("%.2f", _rainfall24h.Float64)
			s.Station_name_min = _tele_station_name.String
			if sc != nil {
				s.Rainfall_min_text = sc.Criterion_text
			}
		} else {
			rainfall_datetime := pqx.NullStringToTime(_rainfall_datetime)
			s.Current_time = rainfall_datetime.Format("15:04")
			s.Current_datetime = uDatetime.DatetimeFormat(rainfall_datetime, uDatetime.DATETIME)
			s.Rainfall_max = fmt.Sprintf("%.2f", _rainfall24h.Float64)
			s.Station_name_max = _tele_station_name.String
			if sc != nil {
				s.Rainfall_max_text = sc.Criterion_text
			}
		}
	}
	return s, nil
}
