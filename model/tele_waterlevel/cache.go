// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package tele_waterlevel is a model for public.tele_waterlevel table. This table store tele_waterlevel.
package tele_waterlevel

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
	"haii.or.th/api/thaiwater30/util/float"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

const cacheTableName = "cache.latest_waterlevel"

var sqlCacheWhere = "--WHERE"
var sqlCache string = `
WITH m_t AS (
	SELECT
	ig.remark,
        m.id,
        m.tele_station_name,
        m.tele_station_lat,
        m.tele_station_long,
        m.tele_station_oldcode,
        m.ground_level,
        COALESCE(m.offset, 0) as offset,
        CASE
			WHEN (M.station_type_msl = 0 AND M.agency_id = 12) THEN M.riverbank + "offset"
			WHEN ((M.station_type_msl != 0 OR M.station_type_msl IS NULL) AND M.agency_id = 12) THEN M.riverbank
			ELSE LEAST(left_bank, right_bank)
			END AS min_bank,
        m.left_bank,
        m.right_bank,
        m.sort_order,
        m.subbasin_id,
        m.agency_id,
        a.agency_name,
        a.agency_shortname,
        lg.id as geocode_id,
        lg.tumbon_code,
        lg.tumbon_name,
        lg.amphoe_code,
        lg.amphoe_name,
        lg.province_code,
        lg.province_name,
        lg.area_code,
        lg.area_name,
        b.id AS basin_id,
        b.basin_name
    FROM
		m_tele_station m
		LEFT  JOIN ignore ig ON m.id = ig.station_id ::int
				AND ig.data_category = 'tele_waterlevel'
        INNER JOIN lt_geocode lg ON m.geocode_id = lg.id
        AND m.agency_id IN (3, 8, 9, 10, 21)
        INNER JOIN subbasin sb ON m.subbasin_id = sb.id
        INNER JOIN basin b ON sb.basin_id = b.id
		INNER JOIN agency a ON m.agency_id = a.id
	WHERE (ig.is_ignore = false OR ig.is_ignore IS NULL)
)
SELECT
    ct.data_id_current,
    ct.datetime_current,
    ct.value_current,
    ct.value_previous,
    CASE
        WHEN value_current = 999999 THEN NULL
        WHEN (m.min_bank - m.ground_level) != 0 THEN (
            (ct.value_current - m.ground_level) * 100
        ) / (m.min_bank - m.ground_level)
        ELSE NULL
    END AS storage_percent,
    ct.type || '_waterlevel' as tele_station_type,m.*,flow_rate,discharge 
FROM       m_t m 
INNER JOIN  ` + cacheTableName + ` ct
ON m.id = ct.tele_station_id 
AND ct.type = 'tele'
AND value_current != 999999
`

// Remove this sql, not union data with canal wl
// ` + sqlCacheWhere + `
// UNION ALL
// SELECT    ct.data_id_current,
//        	 ct.datetime_current,
//        	 ct.value_current,
//        	 ct.value_previous,
//        	NULL,
//        	ct.type || '_waterlevel' as tele_station_type,m.*,flow_rate,discharge
// FROM       m_c m
// LEFT JOIN  ` + cacheTableName + ` ct
// ON         m.id = ct.tele_station_id
// AND        ct.type = 'canal'
// ` + sqlCacheWhere + `

//	get waterlevel from cache.latest_waterlevel
//	Parameters:
//		p
//			Waterlevel_InputParam
//	Return:
//		Array Struct_Waterlevel
func GetWaterLevelThailandDataCache(p *Waterlevel_InputParam) ([]*Struct_Waterlevel, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_Waterlevel{}

	q := sqlCache
	qWhere := ""
	itf := []interface{}{}
	fmt.Println(p.Station_id)
	if p.Station_id != "" {
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
		itf = append(itf, p.Station_id)
		qWhere += " tele_station_id = $" + strconv.Itoa(len(itf))
	}
	if p.Agency_id != "" {
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
		itf = append(itf, p.Agency_id)
		qWhere += " agency_id = $" + strconv.Itoa(len(itf))
	}
	if p.Basin_id != "" {
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
		itf = append(itf, p.Basin_id)
		qWhere += " basin_id = $" + strconv.Itoa(len(itf))
	}
	if p.IsHourly { // ชั่วโมงล่าสุด
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}

		// qWhere += " datetime_current BETWEEN date_trunc('hour', Now() - interval '1 hour') AND date_trunc('hour', now()) "

		// Fix criteria to match main page
		// ds := time.Now().Add(-1 * time.Hour).Format("2006-01-02 15:00")
		// de := time.Now().Format("2006-01-02 15:00")
		// qWhere += " datetime_current BETWEEN '" + ds + "' AND '" + de + "' "

		dt := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
		qWhere += " datetime_current >= '" + dt + "' "

	}

	if p.Start_date != "" && p.End_date != "" {
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
		itf = append(itf, p.Start_date)
		qWhere += " datetime_current >= $" + strconv.Itoa(len(itf))
		itf = append(itf, p.End_date)
		qWhere += " AND datetime_current <= $" + strconv.Itoa(len(itf)) //+ "23:59"
		fmt.Println(qWhere)
	}

	//Check Filter province_id
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
	}
	if len(arrProvinceId) > 0 {
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(p.Province_Code, " "))
			qWhere += " province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			qWhere += " province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	if p.Region_Code != "" {
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
		qWhere += " area_code like '" + p.Region_Code + "' "
	}

	if p.IsMain { // จากหน้าหลัก ย้อนหลังแค่ไม่เกิน 3 วัน
		if len(itf) == 0 {
			qWhere += " WHERE "
		} else {
			qWhere += " AND "
		}
		fmt.Println("p.IsMain", p.IsMain)
		// qWhere += " datetime_current::date >= NOW() - interval '3 day' "
		dt := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
		qWhere += " datetime_current >= '" + dt + "' "
	}

	// qc status
	if qWhere == "" {
		qWhere += " WHERE "
	} else {
		qWhere += " AND "
	}
	qWhere += " ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') "

	limit := ""
	if p.Order == "asc" {
		limit += ` ORDER BY storage_percent ASC NULLS LAST `
	} else {
		limit += ` ORDER BY storage_percent DESC NULLS LAST `
	}

	if p.Data_Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(p.Data_Limit)
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = "2006-01-02 15:04"
	}

	// debug sql
	// q = strings.Replace(q, sqlCacheWhere, qWhere, -1)
	fmt.Println(q + qWhere + limit)
	row, err := db.Query(q+qWhere+limit, itf...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var (
			_id                  int64
			_remark              sql.NullString
			_waterlevel_datetime time.Time
			_waterlevel_msl      sql.NullFloat64
			_pre_waterlevel_msl  sql.NullFloat64
			_storage_percent     sql.NullFloat64
			_table               string
			_station_id          int64
			_station_name        pqx.JSONRaw
			_station_lat         sql.NullFloat64
			_station_long        sql.NullFloat64
			_station_oldcode     sql.NullString
			_ground_level        sql.NullFloat64
			_offset              sql.NullFloat64
			_min_bank            sql.NullFloat64
			_left_bank           sql.NullFloat64
			_right_bank          sql.NullFloat64
			_sort_order          sql.NullInt64
			_subbasin_id         sql.NullInt64
			_agency_id           int64
			_agency_name         pqx.JSONRaw
			_agency_shortname    pqx.JSONRaw

			_gecode_id     int64
			_tumbon_code   string
			_tumbon_name   pqx.JSONRaw
			_amphoe_code   string
			_amphoe_name   pqx.JSONRaw
			_province_code string
			_province_name pqx.JSONRaw
			_area_code     string
			_area_name     pqx.JSONRaw
			_basin_id      int64
			_basin_name    pqx.JSONRaw

			_flow_rate sql.NullFloat64
			_discharge sql.NullFloat64
		)
		err = row.Scan(&_id, &_waterlevel_datetime, &_waterlevel_msl, &_pre_waterlevel_msl, &_storage_percent, &_table, &_remark, &_station_id, &_station_name, &_station_lat,
			&_station_long, &_station_oldcode, &_ground_level, &_offset, &_min_bank, &_left_bank, &_right_bank, &_sort_order, &_subbasin_id, &_agency_id, &_agency_name, &_agency_shortname, &_gecode_id, &_tumbon_code, &_tumbon_name,
			&_amphoe_code, &_amphoe_name, &_province_code, &_province_name, &_area_code, &_area_name, &_basin_id, &_basin_name,
			&_flow_rate, &_discharge,
		)

		teleStation := &model_tele_station.Struct_TeleStation{
			Id:                   _station_id,
			Remark:               _remark,
			Tele_station_name:    _station_name.JSON(),
			Tele_station_oldcode: _station_oldcode.String,
			Tele_station_type:    _table,
			Ground_level:         _ground_level.Float64,
			Min_bank:             float.ToFixed(_min_bank.Float64, 2),
			Left_bank:            ValidData(_left_bank.Valid, _left_bank.Float64),
			Right_bank:           ValidData(_right_bank.Valid, _right_bank.Float64),
			SubBasin_id:          _subbasin_id.Int64,
			Agency_id:            _agency_id,
			Geocode_id:           _gecode_id,
		}
		log.Println(teleStation.Remark)
		teleStation.Tele_station_lat, _ = _station_lat.Value()
		teleStation.Tele_station_long, _ = _station_long.Value()

		geocode := &model_lt_geocode.Struct_Geocode{}
		geocode.Tumbon_code = _tumbon_code
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_code = _amphoe_code
		geocode.Amphoe_name = _amphoe_name.JSON()
		geocode.Province_code = _province_code
		geocode.Province_name = _province_name.JSON()
		geocode.Area_code = _area_code
		geocode.Area_name = _area_name.JSON()

		//		Tumbon_code:   _tumbon_code,
		//			Tumbon_name:   _tumbon_name.JSON(),
		//			Amphoe_code:   _amphoe_code,
		//			Amphoe_name:   _amphoe_name.JSON(),
		//			Province_code: _province_code,
		//			Province_name: _province_name.JSON(),

		//		agency := &model_agency.Struct_Agency{
		//			Id:               _agency_id,
		//			Agency_name:      _agency_name.JSON(),
		//			Agency_shortname: _agency_shortname.JSON(),
		//		}
		agency := &model_agency.Struct_Agency{
			Agency_shortname: _agency_shortname.JSON(),
		}
		agency.Id = _agency_id
		agency.Agency_name = _agency_name.JSON()
		agency.Agency_shortname = _agency_shortname.JSON()

		basin := &model_basin.Struct_Basin{
			Id:         _basin_id,
			Basin_name: _basin_name.JSON(),
		}

		waterlevel := &Struct_Waterlevel{
			Id:                  _id,
			Remark:              _remark.String,
			Waterlevel_datetime: _waterlevel_datetime.Format(strDatetimeFormat),
			Waterlevel_msl:      ValidData(_waterlevel_msl.Valid, fmt.Sprintf("%.2f", _waterlevel_msl.Float64)),
			Pre_Waterlevel_msl:  ValidData(_pre_waterlevel_msl.Valid, fmt.Sprintf("%.2f", _pre_waterlevel_msl.Float64)),
			Storage_percent:     ValidData(_storage_percent.Valid, fmt.Sprintf("%.2f", _storage_percent.Float64)),
			FlowRate:            ValidData(_flow_rate.Valid, fmt.Sprintf("%.2f", _flow_rate.Float64)),
			Discharge:           ValidData(_discharge.Valid, fmt.Sprintf("%.2f", _discharge.Float64)),

			Table: _table,

			Station: teleStation,
			Geocode: geocode,
			Agency:  agency,
			Basin:   basin,
		}

		r = append(r, waterlevel)
	}

	return r, nil
}

type waterLevelThailandDataCacheResult struct {
	TeleStationID int64
	DataId_C      sql.NullInt64
	Datetime_C    sql.NullString
	Value_C       sql.NullFloat64
	DataId_P      sql.NullInt64
	Datetime_P    sql.NullString
	Value_P       sql.NullFloat64
	Type          sql.NullString
	QC            sql.NullString
	FlowRate      sql.NullString
	Discharge     sql.NullString
}

//	update cache.latest_waterlevel
func UpdateWaterLevelThailandDataCache() error {
	db, err := pqx.Open()
	if err != nil {
		return errors.New("err begin : " + err.Error())
	}
	tx, err := db.Begin()
	if err != nil {
		return errors.New("err begin")
	}
	defer tx.Rollback()

	lx := pqx.NewAdvisoryLock(tx, cacheTableName)
	b, err := lx.Wait(time.Minute * 10)
	if err != nil {
		//		return errors.Repack(err)
		return errors.New("err NewAdvisoryLock")
	}
	if !b {
		return errors.Newf("Can not get %s update lock", cacheTableName)
	}
	defer lx.Unlock()

	q := "SELECT MAX(datetime_current) FROM " + cacheTableName + " WHERE (qc_status IS NULL OR qc_status ->> 'is_pass' = 'true')"
	var ts sql.NullString
	var lastupd time.Time

	if err = tx.QueryRow(q).Scan(&ts); err == nil {
		// วันเวลาที่มากที่สุดจาก cache tablename -1 hour
		lastupd = pqx.NullStringToTime(ts).Add(-1 * time.Hour)
	}

	q = `
	SELECT  tw.tele_station_id, 
	        tw.id, 
	        tw.waterlevel_datetime,
			CASE
              WHEN ((station_type_msl = 0) AND ("offset" IS NOT NULL) AND (agency_id = 12))
		        THEN (tw.waterlevel_m + COALESCE("offset", (0)::real))
		      WHEN (((station_type_msl = 1) OR (station_type_msl IS NULL)) AND (agency_id = 12))
				THEN tw.waterlevel_m
			  WHEN ((agency_id = 9) AND (tw.waterlevel_msl != 999999))
			    THEN (tw.waterlevel_msl + COALESCE("offset", (0)::real))
              ELSE tw.waterlevel_msl
            END AS waterlevel_msl, 
			tw.qc_status,
	        'tele',
					flow_rate,
					discharge	            	                             
	FROM       ( 
	                    SELECT   id, 
	                             tele_station_id, 
	                             waterlevel_m, 
	                             waterlevel_msl AS waterlevel_msl, 
	                             waterlevel_datetime,
															 flow_rate,
															 discharge,
															 qc_status
	                    FROM     tele_waterlevel 
	                    WHERE    deleted_at = To_timestamp(0) 
	                    AND      waterlevel_datetime >= $1 ) tw 
	INNER JOIN m_tele_station m 
	ON         tw.tele_station_id = m.id 
	AND        m.agency_id IN (3,8,9,10,12,21)
	-- UNION ALL 
	-- SELECT  tw.tele_station_id, 
	--         tw.id, 
	--         tw.waterlevel_datetime, 
	--         tw.waterlevel_msl, 
	--         tw.qc_status,
	-- 				'tele',
	-- 				flow_rate,
	-- 				discharge	            	                               
	-- FROM       ( 
	--                     SELECT   id, 
	--                              tele_station_id, 
	--                              waterlevel_m, 
	--                              waterlevel_msl AS waterlevel_msl, 
	--                              waterlevel_datetime, 
	-- 														 flow_rate,
	-- 														 discharge,	                             
	--                              qc_status
	--                     FROM     tele_waterlevel 
	--                     WHERE    deleted_at = To_timestamp(0) 
	--                     -- AND      date_part('minute'::text, waterlevel_datetime) = '0'::DOUBLE PRECISION
	--                     AND      waterlevel_datetime >= $1 ) tw 
	-- INNER JOIN m_tele_station m 
	-- ON         tw.tele_station_id = m.id 
	-- AND        m.agency_id IN (9,21) 
	-- UNION ALL 
	-- SELECT     cw.canal_station_id, 
	--            cw.id, 
	--            cw.canal_waterlevel_datetime, 
	--            cw.canal_waterlevel_value, 
	--            cw.qc_status,
	--            'canal',
	--					 null as flow_rate,
	--					 null as discharge	            
	-- FROM       ( 
	--                    SELECT   id, 
	--                             canal_station_id, 
	--                             canal_waterlevel_value, 
	--                             canal_waterlevel_datetime, 
	--                             qc_status
	--                    FROM     canal_waterlevel 
	--                    WHERE    deleted_at = '1970-01-01 07:00:00+07' 
	--                    AND      canal_waterlevel_datetime >= $1 ) cw 
	-- INNER JOIN m_canal_station m 
	-- ON         cw.canal_station_id = m.id 
	-- ORDER BY 6, 1, 3
	`

	fmt.Println(q)
	fmt.Println(lastupd)

	rows, err := tx.Query(q, lastupd)
	if err != nil {
		return errors.New(" err tx.Query : " + tx.Commit().Error())
		//		return errors.Repack(tx.Commit())
	}
	// newdata := make(map[string]*waterLevelThailandDataCacheResult)
	newdata := make([]*waterLevelThailandDataCacheResult, 0)
	// var ids string
	for rows.Next() {
		d := new(waterLevelThailandDataCacheResult)
		if err := rows.Scan(&d.TeleStationID,
			&d.DataId_C, &d.Datetime_C, &d.Value_C,
			&d.QC,
			&d.Type,
			&d.FlowRate,
			&d.Discharge,
		); err != nil {
			return errors.Repack(err)
		}
		newdata = append(newdata, d)
		// _id := strconv.FormatInt(d.TeleStationID, 10)
		// newdata[d.Type.String+"."+_id] = d
		// if ids != "" {
		// 	ids += ","
		// }
		// ids += _id
	}
	if err := rows.Close(); err != nil {
		return errors.New("err rows.Close()")
		//		return errors.Repack(err)
	}

	if len(newdata) == 0 {
		return nil
	}

	// Do we need to update old record?
	// q = `SELECT tele_station_id,
	// 	  data_id_current, datetime_current, value_current,
	// 	  data_id_previous, datetime_previous, value_previous,
	// 	  type
	// 	  FROM ` + cacheTableName + " WHERE tele_station_id IN (" + ids + ")"

	// if rows, err = tx.Query(q); err != nil {
	// 	return errors.Repack(err)
	// }

	// olddata := make(map[string]*waterLevelThailandDataCacheResult)
	// for rows.Next() {
	// 	d := new(waterLevelThailandDataCacheResult)
	// 	if err := rows.Scan(&d.TeleStationID,
	// 		&d.DataId_C, &d.Datetime_C, &d.Value_C,
	// 		&d.DataId_P, &d.Datetime_P, &d.Value_P,
	// 		&d.Type,
	// 	); err != nil {
	// 		return errors.Repack(err)
	// 	}
	// 	_id := strconv.FormatInt(d.TeleStationID, 10)
	// 	olddata[d.Type.String+"."+_id] = d
	// }
	// if err := rows.Close(); err != nil {
	// 	return errors.Repack(err)
	// }

	// for k, d := range newdata {
	// 	_id := strconv.FormatInt(d.TeleStationID, 10)
	// 	od := olddata[d.Type.String+"."+_id]
	// 	// No old data to check
	// 	if od == nil {
	// 		continue
	// 	}

	// 	// Data was not changed, no need to update
	// 	if *od == *d {
	// 		delete(newdata, k)
	// 		continue
	// 	}

	// 	// Do we got data1?
	// 	if d.DataId_P.Valid {
	// 		continue
	// 	}
	// 	// ข้อมูลเก่า กับ ข้อมูลใหม่ คนละอันกัน
	// 	if od.DataId_C != d.DataId_C && od.Type == d.Type && od.Datetime_C != d.Datetime_C {
	// 		d.DataId_P = od.DataId_C
	// 		d.Datetime_P = od.Datetime_C
	// 		d.Value_P = od.Value_C
	// 	}
	// }

	// q = `INSERT INTO ` + cacheTableName + ` (tele_station_id,
	// 	  data_id_current, datetime_current, value_current,
	// 	  data_id_previous, datetime_previous, value_previous,
	// 	  type) VALUES ($1,$2,
	// 	  $3,$4,$5,$6,$7,$8) ON CONFLICT (tele_station_id, type) DO UPDATE SET
	// 	  data_id_current = EXCLUDED.data_id_current,
	// 	  datetime_current =  EXCLUDED.datetime_current,
	// 	  value_current = EXCLUDED.value_current,
	// 	  data_id_previous = EXCLUDED.data_id_previous,
	// 	  datetime_previous =  EXCLUDED.datetime_previous,
	// 	  value_previous = EXCLUDED.value_previous`

	// q = `INSERT INTO ` + cacheTableName + ` (tele_station_id,
	// 	  data_id_current, datetime_current, value_current,
	// 	  qc_status,
	// 	  type) VALUES ($1, $2, $3, $4, $5, $6)
	// 	  ON CONFLICT (tele_station_id, type) DO UPDATE SET
	// 	  data_id_current = EXCLUDED.data_id_current,
	// 	  datetime_current =  EXCLUDED.datetime_current,
	// 	  qc_status = EXCLUDED.qc_status,
	// 	  value_current = EXCLUDED.value_current,
	// 	  data_id_previous = ` + cacheTableName + `.data_id_current,
	// 	  datetime_previous =  ` + cacheTableName + `.datetime_current,
	// 	  value_previous = ` + cacheTableName + `.value_current`
	// qUpdate := `
	// UPDATE ` + cacheTableName + `
	// SET data_id_current = $2,
	// 	datetime_current = $3,
	// 	value_current = $4,
	// 	qc_status = $5,
	// 	data_id_previous = data_id_current,
	// 	datetime_previous = datetime_current,
	// 	value_previous = value_current
	// WHERE tele_station_id = $1 AND type = $6
	// AND datetime_current > $3
	// `
	qUpdate := `
	UPDATE ` + cacheTableName + `
	SET    data_id_current = CASE 
							WHEN $3 > datetime_current THEN  $2 
							ELSE data_id_current 
							END, 
		datetime_current = CASE 
								WHEN $3 > datetime_current  THEN  $3 
								ELSE datetime_current 
							END, 
		value_current = CASE 
							WHEN $3 > datetime_current THEN  $4
							ELSE value_current 
						END, 
		qc_status = CASE 
						WHEN $3 > datetime_current THEN  $5
						ELSE qc_status 
					END, 
		flow_rate = CASE 
						WHEN $3 > datetime_current THEN  $7
						ELSE flow_rate 
					END,
		discharge = CASE 
						WHEN $3 > datetime_current THEN  $8
						ELSE discharge 
					END,					
		data_id_previous = CASE 
								WHEN $3 > datetime_current  THEN  data_id_current 
								ELSE data_id_previous 
							END, 
		datetime_previous = CASE 
								WHEN $3 > datetime_current  THEN  datetime_current 
								ELSE datetime_previous 
							END, 
		value_previous = CASE 
							WHEN $3 > datetime_current THEN  value_current 
							ELSE value_previous 
							END 
	WHERE  tele_station_id = $1 
		AND type = $6 
	`

	qInsert := `
	INSERT INTO ` + cacheTableName + `
	(tele_station_id, data_id_current, datetime_current, value_current, qc_status, type, flow_rate, discharge)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	fmt.Println(qUpdate)
	fmt.Println(qInsert)
	// stmt, err := tx.Prepare(q)
	// if err != nil {
	// 	return errors.Repack(err)
	// }

	for _, d := range newdata {
		//		fmt.Printf("%v %v %v %v %v %v %v %v\n", d.TeleStationID, d.DataId_C, d.Datetime_C, d.Value_C, d.DataId_P, d.Datetime_P, d.Value_P, d.Type)
		// if _, err = stmt.Exec(d.TeleStationID,
		// 	d.DataId_C, d.Datetime_C, d.Value_C,
		// 	d.DataId_P, d.Datetime_P, d.Value_P,
		// 	d.Type); err != nil {
		// 	return errors.Repack(err)
		// }
		// if _, err = stmt.Exec(d.TeleStationID,
		// 	d.DataId_C, d.Datetime_C, d.Value_C,
		// 	d.QC,
		// 	d.Type); err != nil {
		// 	return errors.Repack(err)
		// }
		param := []interface{}{d.TeleStationID, d.DataId_C, d.Datetime_C, d.Value_C, d.QC, d.Type, d.FlowRate, d.Discharge}
		res, err := db.Exec(qUpdate, param...)
		if err != nil {
			return errors.Repack(err)
		}
		count, err2 := res.RowsAffected()
		if err2 != nil {
			return errors.Repack(err)
		}
		//update ไม่เกิดผล แสดงว่าเป็น สถานีใหม่ ให้ insert
		if count < 1 {
			if _, err := db.Exec(qInsert, param...); err != nil {
				return errors.Repack(err)
			}
		}

	}
	// stmt.Close()
	// return errors.Repack(tx.Commit())
	return nil
}

//	ข้อมูลระดับน้ำ ล่าสุด ในจังหวัด
//	Parameters:
//		prov_id
//			รหัสจังหวัด
//	Return:
//		 ข้อมูลระดับน้ำ ล่าสุด ในจังหวัด
func GetWaterLevelMinMax(prov_id string) (*Struct_WaterlevelMinMax, error) {
	if prov_id == "" {
		return nil, errors.New("no prov_id")
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	wl_setting, err := uSetting.New_Struct_WaterlevelSetting()
	if err != nil {
		return nil, errors.Repack(err)
	}

	// query ตามจังหวัดเฉพาะวันนี้ มาหาค่า min, max เอาเอง
	qWhere := " WHERE ( qc_status IS NULL OR qc_status->>'is_pass' = 'true') "
	q := strings.Replace(sqlCache, sqlCacheWhere, qWhere, -1)
	q = " SELECT datetime_current, value_current, value_previous, storage_percent, tele_station_name->>'th', province_name->>'th' FROM (" + q + ") a WHERE province_code = $1 AND storage_percent IS NOT NULL"
	fmt.Println(q)
	row, err := db.Query(q, prov_id)
	if err != nil {
		return nil, errors.Repack(err)
	}

	var maxStruct, minStruct *temp_WaterLevelMinMax
	for row.Next() {
		d, err := newTemp_WaterLevelMinMax(row)
		if err != nil {
			return nil, errors.Repack(err)
		}
		// เริ่มแรกมาให้เก็บค่าของ min
		if minStruct == nil {
			minStruct = d
			continue
		}
		if maxStruct == nil {
			if d.Value > minStruct.Value {
				// ค่า ใหม่ มากกว่าต่า min ให้ max = ค่าใหม่
				maxStruct = d
			} else {
				// ค่า ใหม่ น้อยกว่าค่า min ให้ สลับค่าน้อย ไปเป็นค่ามาก แล้ว ค่าใหม่ไปเป็นค่าเก่า
				maxStruct = minStruct
				minStruct = d
			}
		}
		// ค่าใหม่น้อยสุด
		if d.Value < minStruct.Value {
			minStruct = d
		}
		// ค่าใหม่มากสุด
		if d.Value > maxStruct.Value {
			maxStruct = d
		}
	}

	rs := &Struct_WaterlevelMinMax{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"}
	if minStruct != nil {
		rs.Current_time = minStruct.Time
		rs.Current_date_thai = minStruct.Date
		rs.Station_name_min = minStruct.Name
		rs.Percent_min = minStruct.Percent
		rs.Wl_msl_min = minStruct.Wl
		rs.Wl_min_before_level = minStruct.Level
		// เปรี่ยบเทียบค่ากับเกณฑ์จาก settign
		mn := wl_setting.CompareScale(minStruct.Value)
		if mn != nil {
			rs.Wl_min_color = mn.Color
		}
	}
	if maxStruct != nil {
		rs.Station_name_max = maxStruct.Name
		rs.Percent_max = maxStruct.Percent
		rs.Wl_msl_max = maxStruct.Wl
		rs.Wl_max_before_level = maxStruct.Level
		// เปรี่ยบเทียบค่ากับเกณฑ์จาก settign
		mx := wl_setting.CompareScale(maxStruct.Value)
		if mx != nil {
			rs.Wl_max_color = mx.Color
		}
	}

	return rs, nil
}
