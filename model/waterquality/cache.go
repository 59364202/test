// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package waterquality is a model for public.waterquality table. This table store waterquality.
package waterquality

import (
	"database/sql"
	"encoding/json"
	//	"fmt"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	model_setting "haii.or.th/api/server/model/setting"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_waterquality_station "haii.or.th/api/thaiwater30/model/waterquality_station"

	"haii.or.th/api/thaiwater30/util/validdata"
)

const cacheTableName = "cache.latest_waterquality"

//	get waterquality from cache.latest_waterquality
//	Parameters:
//		station_id
//			รหัสสถานี
//	Return:
//		Array Struct_WaterQuality
func GetWaterQualityThailandDataCache(param *Param_WaterQualityCache) ([]*Struct_WaterQuality, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_WaterQuality{}
	itf := []interface{}{}

	q := `
	SELECT w.waterquality_id,
	       w.waterquality_datetime,
	       w.waterquality_do,
	       w.waterquality_ph,
	       w.waterquality_temp,
	       w.waterquality_turbid,
	       w.waterquality_bod,
	       w.waterquality_tcb,
	       w.waterquality_fcb,
	       w.waterquality_nh3n,
	       w.waterquality_wqi,
	       w.waterquality_ammonium,
	       w.waterquality_nitrate,
	       w.waterquality_colorstatus,
	       w.waterquality_status,
	       w.waterquality_salinity,
	       w.waterquality_conductivity,
	       w.waterquality_tds,
	       w.waterquality_chlorophyll,
	       mws.waterquality_station_lat,
	       mws.waterquality_station_long,
	       mws.waterquality_station_name,
	       lg.id,
	       lg.province_name,
	       lg.province_code,
	       lg.amphoe_name,
	       lg.tumbon_name,
	       a.id,
	       a.agency_name,
	       a.agency_shortname,
	       w.data_id                        AS data_id,
	       mws.waterquality_station_oldcode AS oldcode,
	       mws.is_active,
				 mws.show_status,
				 b.id AS basin_id, b.basin_code, b.basin_name,subbasin_id	       
	FROM   cache.latest_waterquality w
	       INNER JOIN m_waterquality_station mws
	               ON w.waterquality_id = mws.id
	       INNER JOIN lt_geocode lg
	               ON mws.geocode_id = lg.id
	       INNER JOIN agency a
	               ON mws.agency_id = a.id 
				LEFT  JOIN basin b ON mws.basin_id = b.id
	WHERE  mws.deleted_at = To_timestamp(0)
		   AND (mws.is_active = 'Y' OR mws.is_active IS NULL)
		   AND mws.is_ignore = 'false'
		   AND ( w.qc_status IS NULL OR w.qc_status->>'is_pass' = 'true') 
	`
	arrProvinceId := []string{}
	if param.Province_Code != "" {
		arrProvinceId = strings.Split(param.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		if len(arrProvinceId) == 1 {
			itf = append(itf, strings.Trim(param.Province_Code, " "))
			q += " AND lg.province_code = $" + strconv.Itoa(len(itf))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrProvinceId {
				itf = append(itf, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(itf)))
			}
			q += " AND lg.province_code IN (" + strings.Join(arrSqlCmd, ",") + ") "
		}
	}

	if param.IsMain { // หน้าหลัก ย้อนหลังแค่ไม่เกิน 3 วัน
		dt := time.Now().AddDate(0, 0, -3).Format("2006-01-02 15:04")
		// q += " AND w.waterquality_datetime > current_timestamp - interval '3 day' "
		q += " AND w.waterquality_datetime > '" + dt + "' "
	}
	//	 พบค่า DO เป็น 0 ทางหน่วยงานแจ้งว่า มีโอกาสมีค่าเป็น 0 ได้ทั้งหัววัดเสีย และวัดค่าได้ 0 จริง
	//		q += "AND waterquality_do <> 0 "
	if param.StationId > 0 {
		q += " AND mws.id = $1"
		itf = append(itf, param.StationId)
	}

	if param.Basin_id != "" {
		itf = append(itf, param.Basin_id)
		q += " AND basin_id = $" + strconv.Itoa(len(itf))
	}

	if param.Data_Limit > 0 {
		q += " LIMIT " + strconv.Itoa(param.Data_Limit)
	}

	//	fmt.Println(q)
	rows, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			_waterquality_id           int64
			_waterquality_datetime     time.Time
			_waterquality_do           sql.NullFloat64
			_waterquality_ph           sql.NullFloat64
			_waterquality_temp         sql.NullFloat64
			_waterquality_turbid       sql.NullFloat64
			_waterquality_bod          sql.NullFloat64
			_waterquality_tcb          sql.NullFloat64
			_waterquality_fcb          sql.NullFloat64
			_waterquality_nh3n         sql.NullFloat64
			_waterquality_wqi          sql.NullFloat64
			_waterquality_ammonium     sql.NullFloat64
			_waterquality_nitrate      sql.NullFloat64
			_waterquality_colorstatus  sql.NullString
			_waterquality_status       sql.NullString
			_waterquality_salinity     sql.NullFloat64
			_waterquality_conductivity sql.NullFloat64
			_waterquality_tds          sql.NullFloat64
			_waterquality_chlorophyll  sql.NullFloat64
			_waterquality_station_name pqx.JSONRaw
			_waterquality_station_lat  sql.NullFloat64
			_waterquality_station_long sql.NullFloat64
			_geocode_id                int64
			_province_name             pqx.JSONRaw
			_province_code             sql.NullString
			_amphoe_name               pqx.JSONRaw
			_tumbon_name               pqx.JSONRaw
			_agency_id                 int64
			_agency_name               pqx.JSONRaw
			_agency_shortname          pqx.JSONRaw
			_data_id                   sql.NullInt64
			_oldcode                   sql.NullString
			_is_active                 sql.NullString
			_show_status               pqx.JSONRaw
			_basin_id                  sql.NullInt64
			_basin_code                sql.NullInt64
			_basin_name                sql.NullString

			_sub_basin_id sql.NullInt64

			d       *Struct_WaterQuality                                   = &Struct_WaterQuality{}
			station *model_waterquality_station.Struct_WaterQualityStation = &model_waterquality_station.Struct_WaterQualityStation{}
		)
		if err := rows.Scan(&_waterquality_id, &_waterquality_datetime, &_waterquality_do, &_waterquality_ph, &_waterquality_temp,
			&_waterquality_turbid, &_waterquality_bod, &_waterquality_tcb, &_waterquality_fcb, &_waterquality_nh3n, &_waterquality_wqi,
			&_waterquality_ammonium, &_waterquality_nitrate, &_waterquality_colorstatus, &_waterquality_status, &_waterquality_salinity,
			&_waterquality_conductivity, &_waterquality_tds, &_waterquality_chlorophyll, &_waterquality_station_lat, &_waterquality_station_long,
			&_waterquality_station_name, &_geocode_id, &_province_name, &_province_code, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name,
			&_agency_shortname, &_data_id, &_oldcode, &_is_active, &_show_status, &_basin_id, &_basin_code, &_basin_name, &_sub_basin_id); err != nil {
			return nil, err
		}
		d.Waterquality_Datetime = _waterquality_datetime.Format(model_setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		d.Id = _data_id.Int64
		d.Waterquality_Id = _waterquality_id
		d.Station_type = "waterquality"
		d.Waterquality_Colorstatus = _waterquality_colorstatus.String
		d.Waterquality_Status = _waterquality_status.String

		Do := validdata.ValidData(_waterquality_do.Valid, _waterquality_do.Float64)
		Ph := validdata.ValidData(_waterquality_ph.Valid, _waterquality_ph.Float64)
		Temp := validdata.ValidData(_waterquality_temp.Valid, _waterquality_temp.Float64)
		Turbid := validdata.ValidData(_waterquality_turbid.Valid, _waterquality_turbid.Float64)
		Bod := validdata.ValidData(_waterquality_bod.Valid, _waterquality_bod.Float64)
		Tcb := validdata.ValidData(_waterquality_tcb.Valid, _waterquality_tcb.Float64)
		Fcb := validdata.ValidData(_waterquality_fcb.Valid, _waterquality_fcb.Float64)
		Nh3n := validdata.ValidData(_waterquality_nh3n.Valid, _waterquality_nh3n.Float64)
		Wqi := validdata.ValidData(_waterquality_wqi.Valid, _waterquality_wqi.Float64)
		Ammonium := validdata.ValidData(_waterquality_ammonium.Valid, _waterquality_ammonium.Float64)
		Nitrate := validdata.ValidData(_waterquality_nitrate.Valid, _waterquality_nitrate.Float64)
		Salinity := validdata.ValidData(_waterquality_salinity.Valid, _waterquality_salinity.Float64)
		Conductivity := validdata.ValidData(_waterquality_conductivity.Valid, _waterquality_conductivity.Float64)
		Tds := validdata.ValidData(_waterquality_tds.Valid, _waterquality_tds.Float64)
		Chlorophyll := validdata.ValidData(_waterquality_chlorophyll.Valid, _waterquality_chlorophyll.Float64)

		// ตรวจสอบแต่ละ column ว่าจะแสดงผล column ไหนบ้าง ถ้าไม่แสดง column ไหน จะ return ค่าเป็น null
		//ข้อมูล คุณภาพน้ำจากสถานีฯ คพ. ที่ส่งไปยัง สสนก. เป็นข้อมูลดิบที่ยังไม่ได้กรอง ซึ่งทำให้บางสถานี ที่ไม่มีหัววัด ส่งค่าไปเป็น ค่า ศูนย์ เช่น หัววัด DO หรือบางสถานี ค่า error
		var show_status WaterQualityshowStatusStruct
		if err := json.Unmarshal(_show_status, &show_status); err != nil {
			return nil, err
		}
		d.Waterquality_Do = validdata.ValidData(show_status.Do, Do)
		d.Waterquality_Ph = validdata.ValidData(show_status.Ph, Ph)
		d.Waterquality_Temp = validdata.ValidData(show_status.Temp, Temp)
		d.Waterquality_Turbid = validdata.ValidData(show_status.Turbid, Turbid)
		d.Waterquality_Bod = validdata.ValidData(show_status.Bod, Bod)
		d.Waterquality_Tcb = validdata.ValidData(show_status.Tcb, Tcb)
		d.Waterquality_Fcb = validdata.ValidData(show_status.Fcb, Fcb)
		d.Waterquality_Nh3n = validdata.ValidData(show_status.Nh3n, Nh3n)
		d.Waterquality_Wqi = validdata.ValidData(show_status.Wqi, Wqi)
		d.Waterquality_Ammonium = validdata.ValidData(show_status.Ammonium, Ammonium)
		d.Waterquality_Nitrate = validdata.ValidData(show_status.Nitrate, Nitrate)
		d.Waterquality_Salinity = validdata.ValidData(show_status.Salinity, Salinity)
		d.Waterquality_Conductivity = validdata.ValidData(show_status.Conductivity, Conductivity)
		d.Waterquality_Tds = validdata.ValidData(show_status.Tds, Tds)
		d.Waterquality_Chlorophyll = validdata.ValidData(show_status.Chlorophyll, Chlorophyll)

		station.Id = _waterquality_id
		station.Waterquality_Station_Name = _waterquality_station_name.JSON()
		station.Geocode_id = _geocode_id
		station.Province_Code = _province_code.String
		station.Province_Name = _province_name.JSON()
		station.Amphoe_Name = _amphoe_name.JSON()
		station.Tumbon_Name = _tumbon_name.JSON()
		station.Waterquality_Station_Lat = validdata.ValidData(_waterquality_station_lat.Valid, _waterquality_station_lat.Float64)
		station.Waterquality_Station_Long = validdata.ValidData(_waterquality_station_long.Valid, _waterquality_station_long.Float64)
		station.Agency_id = _agency_id
		station.Agency_Name = _agency_name.JSON()
		station.Agency_Shortname = _agency_shortname.JSON()
		station.Waterquality_Station_Oldcode = _oldcode.String
		station.Is_active = _is_active.String
		station.SubBasin_id = _sub_basin_id.Int64

		/*----- Basin -----*/
		d.Basin = &model_basin.Struct_Basin{}
		d.Basin.Id = _basin_id.Int64
		d.Basin.Basin_code = _basin_code.Int64

		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		d.Basin.Basin_name = json.RawMessage(_basin_name.String)

		d.Waterquality_Station = station

		r = append(r, d)
	}
	return r, nil
}

type DataCacheResult struct {
	WaterQualityId int64
	DataId         int64
	DateTime       sql.NullString
	_do            sql.NullFloat64
	_conductivity  sql.NullFloat64
	_ph            sql.NullFloat64
	_temp          sql.NullFloat64
	_turbid        sql.NullFloat64
	_bod           sql.NullFloat64
	_tcb           sql.NullFloat64
	_fcb           sql.NullFloat64
	_nh3n          sql.NullFloat64
	_wqi           sql.NullFloat64
	_ammonium      sql.NullFloat64
	_nitrate       sql.NullFloat64
	_salinity      sql.NullFloat64
	_tds           sql.NullFloat64
	_chlorophyll   sql.NullFloat64
	_colorstatus   sql.NullString
	_status        sql.NullString
	_is_active     sql.NullString
	_qc            sql.NullString
}

//	update cache.latest_waterquality
func UpdateWaterQualityThailandDataCache() error {
	db, err := pqx.Open()
	if err != nil {
		return errors.New("err open : " + err.Error())
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

	dt := time.Now().Add(-1 * time.Hour).Format("2006-01-02 15:04:05")
	q := `
	SELECT w.id,
	       w.waterquality_id,
	       w.waterquality_datetime,
	       w.waterquality_do,
	       w.waterquality_conductivity,
	       w.waterquality_ph,
	       w.waterquality_temp,
	       w.waterquality_turbid,
	       w.waterquality_bod,
	       w.waterquality_tcb,
	       w.waterquality_fcb,
	       w.waterquality_nh3n,
	       w.waterquality_wqi,
	       w.waterquality_ammonium,
	       w.waterquality_nitrate,
	       w.waterquality_salinity,
	       w.waterquality_tds,
	       w.waterquality_chlorophyll,
	       w.waterquality_colorstatus,
		   w.waterquality_status,
		   w.qc_status
	FROM   waterquality w
	       INNER JOIN (SELECT waterquality_id,
	                          Max(waterquality_datetime) AS waterquality_datetime
	                   FROM   waterquality
	                   WHERE  deleted_at = To_timestamp(0)
	                   AND    waterquality_datetime >= '` + dt + `'
	                   GROUP  BY waterquality_id) tw
	               ON w.waterquality_id = tw.waterquality_id
	                  AND w.waterquality_datetime = tw.waterquality_datetime
	                  AND w.deleted_at = To_timestamp(0)
	`
	// หาข้อมูลที่ datetime มากกว่าที่เรามี เวลา xx.00 ของทุกชั่วโมง
	rows, err := tx.Query(q)
	if err != nil {
		return errors.New(" err tx.Query : " + tx.Commit().Error())
	}
	newData := []*DataCacheResult{}
	for rows.Next() {
		d := &DataCacheResult{}
		err = rows.Scan(&d.DataId, &d.WaterQualityId, &d.DateTime, &d._do, &d._conductivity, &d._ph, &d._temp, &d._turbid, &d._bod, &d._tcb, &d._fcb,
			&d._nh3n, &d._wqi, &d._ammonium, &d._nitrate, &d._salinity, &d._tds, &d._chlorophyll, &d._colorstatus, &d._status, &d._qc)
		if err != nil {
			return err
		}
		newData = append(newData, d)
	}
	if err := rows.Close(); err != nil {
		return errors.Repack(err)
	}
	// no new data
	if len(newData) == 0 {
		return nil
	}
	q = `
	INSERT INTO ` + cacheTableName + `
	            (
	                        data_id,
	                        waterquality_id,
	                        waterquality_datetime,
	                        waterquality_do,
	                        waterquality_conductivity,
	                        waterquality_ph,
	                        waterquality_temp,
	                        waterquality_turbid,
	                        waterquality_bod,
	                        waterquality_tcb,
	                        waterquality_fcb,
	                        waterquality_nh3n,
	                        waterquality_wqi,
	                        waterquality_ammonium,
	                        waterquality_nitrate,
	                        waterquality_salinity,
	                        waterquality_tds,
	                        waterquality_chlorophyll,
	                        waterquality_colorstatus,
							waterquality_status,
							qc_status
	            )
	            VALUES
	            (
	                        $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
							$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
							$21
	            )
	on conflict
	            (
	                        waterquality_id
	            )
	            do UPDATE
	set    data_id = excluded.data_id,
	       waterquality_datetime = excluded.waterquality_datetime,
	       waterquality_do = excluded.waterquality_do,
	       waterquality_conductivity = excluded.waterquality_conductivity,
	       waterquality_ph = excluded.waterquality_ph,
	       waterquality_temp = excluded.waterquality_temp,
	       waterquality_turbid = excluded.waterquality_turbid,
	       waterquality_bod = excluded.waterquality_bod,
	       waterquality_tcb = excluded.waterquality_tcb,
	       waterquality_fcb = excluded.waterquality_fcb,
	       waterquality_nh3n = excluded.waterquality_nh3n,
	       waterquality_wqi = excluded.waterquality_wqi,
	       waterquality_ammonium = excluded.waterquality_ammonium,
	       waterquality_nitrate = excluded.waterquality_nitrate,
	       waterquality_salinity = excluded.waterquality_salinity,
	       waterquality_tds = excluded.waterquality_tds,
	       waterquality_chlorophyll = excluded.waterquality_chlorophyll,
	       waterquality_colorstatus = excluded.waterquality_colorstatus,
		   waterquality_status = excluded.waterquality_status,
		   qc_status = excluded.qc_status
	`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return errors.Repack(err)
	}
	for _, d := range newData {
		if _, err = stmt.Exec(d.DataId, d.WaterQualityId, d.DateTime, d._do, d._conductivity, d._ph, d._temp,
			d._turbid, d._bod, d._tcb, d._fcb, d._nh3n, d._wqi, d._ammonium,
			d._nitrate, d._salinity, d._tds, d._chlorophyll, d._colorstatus, d._status, d._qc); err != nil {
			return errors.Repack(err)
		}
	}
	stmt.Close()

	return errors.Repack(tx.Commit())
}
