// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package rainfall24hr is a model for public.rainfall24hr table. This table store rainfall24hr.
package rainfall24hr

import (
	"fmt"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"

	"database/sql"
	"encoding/json"

	//	"fmt"
	"strconv"
	"strings"
	"time"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

const cacheTableName = "cache.latest_rainfall24h"

//	get rainfall24h from cache.latest_rainfall24h
//	Parameters:
//		p
//			Param_Rainfall24
//	Return:
//		Array Struct_Rainfall24H
func GetRainfallThailandDataCache(p *Param_Rainfall24) ([]*Struct_Rainfall24H, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_Rainfall24H{}
	itf := []interface{}{}

	//	change r.rainfall24h <= 500
	//	<= 600 ตามที่ผู้ใช้แจ้งกรณีพายุเข้าภาคใต

	// agency_id : 3=>dwr, 8=>egat, 9=>haii, 10=>bma, 13=>tmd
	//	q := `
	//	WITH m_m AS (
	//		SELECT
	//			m.id,
	//			m.tele_station_name,
	//			m.tele_station_lat,
	//			m.tele_station_long,
	//			ig.is_ignore,
	//			m.tele_station_oldcode,
	//			lg.area_code,
	//			lg.area_name,
	//			lg.province_code AS province_code,
	//			lg.province_name,
	//			lg.amphoe_name,
	//			lg.tumbon_name,
	//			a.agency_name,
	//			a.agency_shortname,
	//			lg.warning_zone,
	//			basin_id,
	//			subbasin_id
	//		FROM
	//			m_tele_station m
	//			LEFT  JOIN ignore ig ON m.id = ig.station_id ::int
	//				AND ig.data_category = 'rainfall_24h'
	//			INNER JOIN lt_geocode lg ON m.geocode_id = lg.id
	//				AND m.agency_id IN (3, 8, 9, 10, 13)
	//				AND lg.geocode <> '999999'
	//			INNER JOIN agency a ON m.agency_id = a.id
	//		WHERE ig.is_ignore = false OR ig.is_ignore IS NULL
	//		ORDER BY
	//			m.id
	//	)
	//	SELECT
	//		m.id,
	//		m.tele_station_name,
	//		m.tele_station_lat,
	//		m.tele_station_long,
	//		m.tele_station_oldcode,
	//		m.amphoe_name,
	//		m.tumbon_name,
	//		m.province_code,
	//		m.province_name,
	//		m.area_code,
	//		m.area_name,
	//		m.warning_zone,
	//		m.agency_name,
	//		m.agency_shortname,
	//		r.data_id as data_id,
	//		r.rainfall24h,
	//		CASE
	//			WHEN r.rainfall1h = 999999 THEN NULL
	//			ELSE r.rainfall1h
	//		END AS rainfall1h,
	//		r.rainfall_datetime,
	//		b.id AS basin_id, b.basin_code, b.basin_name,m.subbasin_id
	//	FROM
	//		m_m m
	//		LEFT JOIN cache.latest_rainfall24h r ON m.id = r.tele_station_id
	//		LEFT  JOIN basin b ON m.basin_id = b.id

	q := `
	WITH m_m AS (
		SELECT
			m.id,
			m.tele_station_name,
			m.tele_station_lat,
			m.tele_station_long,
			ig.is_ignore,
			m.tele_station_oldcode,
			lg.area_code,
			lg.area_name,
			lg.province_code AS province_code,
			lg.province_name,
			lg.amphoe_name,
			lg.tumbon_name,
			a.agency_name,
			a.agency_shortname,
			lg.warning_zone,
			m.basin_id,
			sb.subbasin_code AS subbasin_id
		FROM
			m_tele_station m
			LEFT  JOIN ignore ig ON m.id = ig.station_id ::int
				AND ig.data_category = 'rainfall_24h'
			INNER JOIN lt_geocode lg ON m.geocode_id = lg.id
				AND m.agency_id IN (3, 8, 9, 10, 13)
				AND lg.geocode <> '999999'
			INNER JOIN agency a ON m.agency_id = a.id
			INNER JOIN subbasin sb ON m.subbasin_id = sb.id
		WHERE ig.is_ignore = false 
		OR ig.is_ignore IS NULL
		AND (M.tele_station_type = 'A' OR M.tele_station_type = 'R' OR M.tele_station_type IS NULL)
		ORDER BY
			m.id
	)
	SELECT
		m.id,
		m.tele_station_name,
		m.tele_station_lat,
		m.tele_station_long,
		m.tele_station_oldcode,
		m.amphoe_name,
		m.tumbon_name,
		m.province_code,
		m.province_name,
		m.area_code,
		m.area_name,
		m.warning_zone,
		m.agency_name,
		m.agency_shortname,
		r.data_id as data_id,
		r.rainfall24h,
		CASE
			WHEN r.rainfall1h = 999999 THEN NULL
			ELSE r.rainfall1h
		END AS rainfall1h,
		r.rainfall_datetime,
		b.id AS basin_id,
		b.basin_code,
		b.basin_name,
		m.subbasin_id
	FROM
		m_m m
		LEFT JOIN cache.latest_rainfall24h r ON m.id = r.tele_station_id 
		LEFT JOIN basin b ON m.basin_id = b.id
		`

	//ตรวจสอบเงื่อนไ ว่าต้องการให้แสดงสถานีที่มีค่าฝน =0 หรือไม่ ถ้า p1.Include_zero = 1 หมายถึงให้แสดงสถานีที่มีค่าฝนเป็น 0 ด้วย
	if p.Include_zero == "" {
		q += "AND ( r.rainfall24h <> 0  OR r.rainfall1h <> 0 ) "
	}

	q += ` AND r.rainfall24h <> 999999 
	                AND r.rainfall24h <= 600
	WHERE   r.rainfall_datetime IS NOT NULL 
			AND ( qc_status->>'is_pass' = 'true' OR qc_status IS NULL)
	`

	//-- Check Filter by parameters --//
	if p.Station_id != "" {
		itf = append(itf, p.Station_id)
		q += " AND m.tele_station_id = $" + strconv.Itoa(len(itf))
	}
	if p.Start_Date != "" || p.End_Date != "" {
		if p.Start_Date != "" {
			itf = append(itf, p.Start_Date)
			q += " AND r.rainfall_datetime >= $" + strconv.Itoa(len(itf))
		}
		if p.End_Date != "" {
			itf = append(itf, p.End_Date)
			q += " AND r.rainfall_datetime <= $" + strconv.Itoa(len(itf))
		}
	} else {
		now := time.Now()
		if p.IsHourly { // ยัอนหลัง 3  ชม.
			// q += " AND r.rainfall_datetime BETWEEN date_trunc('hour', Now() - interval '3 hour') AND date_trunc('hour', now()) "
			q += " AND r.rainfall_datetime BETWEEN '" + now.Add(-3*time.Hour).Format("2006-01-02 15:00") + "' AND '" + now.Format("2006-01-02 15:00") + "' "
		} else { // ย้อนหลัง 24 ชม.
			//-- update 27/06/2018 - ย้อนหลัง  12 hr --//
			// q += " AND r.rainfall_datetime >= NOW() - interval '12 hour' "
			q += " AND r.rainfall_datetime >= '" + now.Add(-12*time.Hour).Format("2006-01-02 15:00") + "' "
		}
	}

	//Check Filter province_id
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(p.Province_Code, " "))
			q += " AND m.province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			q += " AND m.province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	if p.Basin_id != "" {
		itf = append(itf, p.Basin_id)
		q += " AND basin_id = $" + strconv.Itoa(len(itf))
	}

	if p.Region_Code != "" {
		q += " AND m.area_code like '" + p.Region_Code + "' "
	}

	if p.Order == "asc" {
		q += ` ORDER BY r.rainfall24h ASC NULLS LAST `
	} else {
		q += ` ORDER  BY r.rainfall24h DESC nulls last `
	}

	if p.Data_Limit > 0 {
		q += " LIMIT " + strconv.Itoa(p.Data_Limit)
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02 15:04"
	}

	fmt.Println(q)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_name    pqx.JSONRaw
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_amphoe_name          pqx.JSONRaw
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            sql.NullString
			_warning_zone         sql.NullString
			_agency_name          pqx.JSONRaw
			_agency_shortname     pqx.JSONRaw
			_data_id              sql.NullInt64
			_rainfall24h          sql.NullFloat64
			_rainfall1h           sql.NullFloat64
			_rainfall_datetime    sql.NullString

			_basin_id   sql.NullInt64
			_basin_code sql.NullInt64
			_basin_name sql.NullString

			_sub_basin_id sql.NullInt64

			d       *Struct_Rainfall24H           = &Struct_Rainfall24H{}
			station *Struct_TeleStation           = &Struct_TeleStation{}
			agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_tele_station_id, &_tele_station_name, &_tele_station_lat, &_tele_station_long, &_tele_station_oldcode,
			&_amphoe_name, &_tumbon_name, &_province_code, &_province_name, &_area_code, &_area_name, &_warning_zone, &_agency_name, &_agency_shortname,
			&_data_id, &_rainfall24h, &_rainfall1h, &_rainfall_datetime,
			&_basin_id, &_basin_code, &_basin_name, &_sub_basin_id)

		if err != nil {
			return nil, err
		}

		if _rainfall1h.Float64 > 200 { // rainfall 1h ต้องน้อยกว่า 200
			_rainfall1h.Valid = false
		}
		// ฝ่ายน้ำแจ้งปรับจาก 500 เป็น 600 กรณีพายุเข้าภาคใต้
		if _rainfall24h.Float64 > 600 { // rainfall 24h ต้องน้อยกว่า 600
			_rainfall24h.Valid = false
		}

		d.Id = _data_id.Int64
		d.StationType = "rainfall_24h"
		d.Rain1H = ValidData(_rainfall1h.Valid, _rainfall1h.Float64)
		d.Rain24H = ValidData(_rainfall24h.Valid, _rainfall24h.Float64)
		if _rainfall_datetime.Valid {
			d.Time = pqx.NullStringToTime(_rainfall_datetime).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _tele_station_name.JSON()
		station.OldCode = _tele_station_oldcode.String
		station.Type = "rainfall_24h"
		station.SubBasin_id = _sub_basin_id.Int64

		agency.Agency_name = _agency_name.JSON()
		agency.Agency_shortname = _agency_shortname.JSON()

		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code.String
		geocode.Province_name = _province_name.JSON()
		geocode.WarningZone = _warning_zone.String
		geocode.Area_code = _area_code.String
		if _area_name.String == "" {
			_area_name.String = "{}"
		}
		geocode.Area_name = json.RawMessage(_area_name.String)

		/*----- Basin -----*/
		d.Basin = &model_basin.Struct_Basin{}
		d.Basin.Id = _basin_id.Int64
		d.Basin.Basin_code = _basin_code.Int64

		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		d.Basin.Basin_name = json.RawMessage(_basin_name.String)

		d.Station = station
		d.Agency = agency
		d.Geocode = geocode

		r = append(r, d)
	}

	return r, nil
}

type rainfallThailandDataCacheResult struct {
	TeleStationId int64           // รหัสสถานี
	DataId        int64           // รหัสข้อมูล
	Rainfall24H   sql.NullFloat64 // rainfall 24h
	Rainfall1H    sql.NullFloat64 // rainfall 1h
	DateTime      sql.NullString  // วันเวลา
	QC            sql.NullString  // qc_status
}

//	update cache.latest_rainfall24h
func UpdateRainfallThailandDataCache() error {
	db, err := pqx.Open()
	if err != nil {
		return errors.New("err pqx.open : " + err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		return errors.New("err begin : " + err.Error())
	}
	defer tx.Rollback()

	lx := pqx.NewAdvisoryLock(tx, cacheTableName)
	b, err := lx.Wait(time.Minute * 10)
	if err != nil {
		return errors.New("err NewAdvisoryLock : " + err.Error())
	}
	if !b {
		return errors.Newf("Can not get %s update lock", cacheTableName)
	}
	defer lx.Unlock()
	// หา datetime ที่มากที่สุด
	q := "SELECT MAX(rainfall_datetime) FROM " + cacheTableName
	var ts sql.NullString
	var lastupd time.Time

	if err = tx.QueryRow(q).Scan(&ts); err == nil {
		lastupd = pqx.NullStringToTime(ts)
		lastupd = lastupd.Add(time.Hour * -1) // ดูย้อนหลังไป 1 ชม เผื่อข้อมูลแต่ละหน่วยงานมันย้อนหลังมาต่างกัน
	}

	//	q = `
	//	SELECT r24.tele_station_id,
	//	       r24.id,
	//	       r24.rainfall24h,
	//	       r1.rainfall1h,
	//	       r24.rainfall_datetime
	//	FROM   (SELECT tele_station_id,
	//	               Max(rainfall_datetime) AS rainfall_datetime
	//	        FROM   rainfall_24h
	//	        WHERE  deleted_at = To_timestamp(0)
	//	               AND rainfall_datetime >= $1
	//	        GROUP  BY tele_station_id) r24d
	//	       INNER JOIN rainfall_24h r24
	//	               ON r24d.tele_station_id = r24.tele_station_id
	//	                  AND r24d.rainfall_datetime = r24.rainfall_datetime
	//	       LEFT JOIN rainfall_1h r1
	//	              ON r24.tele_station_id = r1.tele_station_id
	//	                 AND r24.rainfall_datetime = r1.rainfall_datetime
	//	                 AND r1.deleted_at = To_timestamp(0)
	//	ORDER  BY r24d.tele_station_id
	//	`

	q = `
SELECT r24.tele_station_id 
       , r24.id 
       , r24.rainfall24h 
       , r1.rainfall1h 
	   , r24.rainfall_datetime 
	   , r24.qc_status
FROM   (
		SELECT tele_station_id 
               , rainfall_datetime 
               , Rank() 
                   OVER ( 
                     partition BY tele_station_id 
                     ORDER BY rainfall_datetime DESC
                   ) AS o 
        FROM   rainfall_24h 
        WHERE  rainfall_datetime >= $1 
               AND deleted_at = To_timestamp(0)
       ) r24d 
       INNER JOIN rainfall_24h r24 
               ON r24d.tele_station_id = r24.tele_station_id 
                  AND r24d.rainfall_datetime = r24.rainfall_datetime 
       LEFT JOIN rainfall_1h r1 
              ON r24.tele_station_id = r1.tele_station_id 
                 AND r24.rainfall_datetime = r1.rainfall_datetime 
                 AND r1.deleted_at = To_timestamp(0) 
WHERE  r24d.o = 1 
ORDER  BY r24d.tele_station_id 
	`

	// หาข้อมูลที่ datetime มากกว่าที่เรามี
	rows, err := tx.Query(q, lastupd)
	if err != nil {
		return errors.New(" err tx.Query : " + tx.Commit().Error())
	}
	newData := []*rainfallThailandDataCacheResult{}
	for rows.Next() {
		d := new(rainfallThailandDataCacheResult)
		if err := rows.Scan(&d.TeleStationId, &d.DataId, &d.Rainfall24H, &d.Rainfall1H, &d.DateTime, &d.QC); err != nil {
			return errors.Repack(err)
		}
		newData = append(newData, d)
	}
	if err := rows.Close(); err != nil {
		return errors.Repack(err)
	}
	// no new data
	if len(newData) == 0 {
		log.Logf("UpdateRainfallThailandDataCache no new data : %s", lastupd)
		return nil
	}
	q = `
	INSERT INTO ` + cacheTableName + `
	            ( 
	                        tele_station_id, 
	                        rainfall_datetime, 
	                        rainfall24h, 
	                        rainfall1h, 
							data_id, 
							qc_status 
	            ) 
	            VALUES 
	            ( 
	                        $1,$2,$3,$4,$5,$6 
	            ) 
	on conflict 
	            ( 
	                        tele_station_id 
	            ) 
	            do UPDATE 
	set    rainfall_datetime = excluded.rainfall_datetime, 
	    	rainfall24h = excluded.rainfall24h, 
	    	rainfall1h = excluded.rainfall1h, 
			data_id = excluded.data_id,
			qc_status = excluded.qc_status
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return errors.Repack(err)
	}
	// upsert
	for _, d := range newData {
		if _, err = stmt.Exec(d.TeleStationId, d.DateTime.String,
			d.Rainfall24H,
			d.Rainfall1H,
			d.DataId,
			d.QC); err != nil {
			return errors.Repack(err)
		}
	}
	stmt.Close()

	return errors.Repack(tx.Commit())
}
