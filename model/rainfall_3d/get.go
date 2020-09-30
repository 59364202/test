// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

// model for public.rainfall_daily table this table store rainfall_yesterday for calculate rainfall 3 day
package rainfall_3d

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_rain24 "haii.or.th/api/thaiwater30/model/rainfall24hr"
)

type Struct_Rainfall3d struct {
	Start_date string      `json:"rainfall_start_date"` // example:`2006-01-02` วันที่เริ่มต้น
	End_date   string      `json:"rainfall_end_date"`   // example:`2006-01-02` วันที่สิ้นสุด
	Rain3D     interface{} `json:"rain_3d,omitempty"`   // example:`235` ฝน 3 วัน (mm.)

	Agency  *model_agency.Struct_Agency      `json:"agency"`  // หน่วยงาน
	Geocode *model_geocode.Struct_Geocode    `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Station *model_rain24.Struct_TeleStation `json:"station"` // สถานี
}

//	get GetRainfall3d from  public.rainfall_daily
//	Parameters:
//		p
//			Param_Rainfall24
//	Return:
//		Array Struct_Rainfall3d
func GetRainfall3d(p *model_rain24.Param_Rainfall24) ([]*Struct_Rainfall3d, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_Rainfall3d{}
	itf := []interface{}{}
	start_date := time.Now().AddDate(0, 0, -3).Format("2006-01-02")

	//	ฝนวันไม่เกิน <= 600 ตามที่ผู้ใช้แจ้งกรณีพายุเข้าภาคใต
	// agency_id : 3=>dwr, 8=>egat, 9=>haii, 10=>bma, 13=>tmd
	q := `WITH sum_query AS (
	SELECT tele_station_id,sum( rainfall_value) AS rainfall3d
	FROM rainfall_daily
	WHERE rainfall_datetime >= '`

	q += start_date + "' "

	q += `GROUP BY tele_station_id HAVING sum(rainfall_value) < 2000::double precision) 
SELECT  
	tele_station_id, date_trunc('day'::text, now() - '3 days'::interval) AS start_date, date_trunc('day'::text, now() - '1 day'::interval) AS end_date,rainfall3d,
	tele_station_oldcode,tele_station_lat,tele_station_long,
	area_code,province_code,area_name,province_name,amphoe_name,tumbon_name,
	agency_id,agency_name,tele_station_name
FROM sum_query 
INNER JOIN m_tele_station m  ON sum_query.tele_station_id = m.id
INNER JOIN lt_geocode lg  ON m.geocode_id = lg.id  AND m.is_ignore = 'false' AND lg.geocode <> '999999'
INNER JOIN agency a ON m.agency_id = a.id`

	//Check Filter province_id
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		if len(itf) == 0 {
			q += " WHERE "
		} else {
			q += " AND "
		}
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(p.Province_Code, " "))
			q += " province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			q += " province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	if p.Region_Code != "" {
		if len(itf) == 0 {
			q += " WHERE "
		} else {
			q += " AND "
		}
		q += " area_code like '" + p.Region_Code + "' "
	}

	q += ` ORDER BY rainfall3d DESC `
	if p.Data_Limit > 0 {
		q += " LIMIT " + strconv.Itoa(p.Data_Limit)
	}

	fmt.Println(q)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}
	strDatetimeFormat := "2006-01-02"

	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_name                 pqx.JSONRaw
			_amphoe_name          pqx.JSONRaw
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            sql.NullString
			_agency_id            sql.NullInt64
			_agency_name          pqx.JSONRaw
			_rainfall3d           sql.NullFloat64
			_rainfall_datestart   sql.NullString
			_rainfall_dateend     sql.NullString

			d       *Struct_Rainfall3d               = &Struct_Rainfall3d{}
			station *model_rain24.Struct_TeleStation = &model_rain24.Struct_TeleStation{}
			agency  *model_agency.Struct_Agency      = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode    = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_tele_station_id, &_rainfall_datestart, &_rainfall_dateend, &_rainfall3d, &_tele_station_oldcode, &_tele_station_lat, &_tele_station_long, &_area_code, &_province_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name, &_name)

		if err != nil {
			return nil, err
		}

		d.Rain3D = fmt.Sprintf("%.2f", _rainfall3d.Float64)
		if _rainfall_datestart.Valid {
			d.Start_date = pqx.NullStringToTime(_rainfall_datestart).Format(strDatetimeFormat)
		}
		if _rainfall_dateend.Valid {
			d.End_date = pqx.NullStringToTime(_rainfall_dateend).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = model_rain24.ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = model_rain24.ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _name.JSON()
		station.OldCode = _tele_station_oldcode.String

		agency.Agency_name = _agency_name.JSON()
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code.String
		geocode.Province_name = _province_name.JSON()
		geocode.Area_code = _area_code.String
		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		geocode.Area_name = json.RawMessage(_area_name.String)

		d.Station = station
		d.Agency = agency
		d.Geocode = geocode

		r = append(r, d)
	}

	return r, nil
}
