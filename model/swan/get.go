// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package swan is a model for public.swan table. This table store swan information.
package swan

import (
	"database/sql"
	"encoding/json"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"

	uDatetime "haii.or.th/api/thaiwater30/util/datetime"
	uSetting "haii.or.th/api/thaiwater30/util/setting"

	model_swan_station "haii.or.th/api/thaiwater30/model/swan_station"
)

//	get swan
//	Return:
//		Array Struct_Swan
func GetSwanCurrentDate() ([]*Struct_Swan, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	var (
		data         []*Struct_Swan
		swan         *Struct_Swan
		swan_station *model_swan_station.Struct_SwanStation

		_id               sql.NullInt64
		_swan_name        sql.NullString
		_lat              sql.NullFloat64
		_long             sql.NullFloat64
		_province_code    sql.NullString
		_province_name    sql.NullString
		_max_swan_highsig sql.NullFloat64
	)
	now := time.Now().Format("2006-01-02")
	next1Day := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	rows, err := db.Query(SQL_SelectSwanCurrentDate, now, next1Day)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data = make([]*Struct_Swan, 0)
	for rows.Next() {
		err = rows.Scan(&_id, &_swan_name, &_lat, &_long, &_province_code, &_province_name, &_max_swan_highsig)
		if err != nil {
			return nil, errors.Repack(err)
		}

		swan = new(Struct_Swan)
		swan_station = new(model_swan_station.Struct_SwanStation)
		data = append(data, swan)

		if !_swan_name.Valid || _swan_name.String == "" {
			_swan_name.String = "{}"
		}
		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}
		swan.Highsig, _ = _max_swan_highsig.Value()

		swan.Swan_Station = swan_station
		swan_station.Id = _id.Int64
		swan_station.Name = json.RawMessage(_swan_name.String)
		swan_station.Lat, _ = _lat.Value()
		swan_station.Long, _ = _long.Value()
		swan_station.Province_Code = _province_code.String
		swan_station.Province_Name = json.RawMessage(_province_name.String)
	}

	return data, nil
}

//	คาดการณ์คลื่น 3 วัน ล่าสุด
//	Return:
//		array คาดการณ์คลื่น 3 วัน ล่าสุด
func Get_WaveForecast() ([]*Struct_WaveForecast, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	q := `
WITH max_date AS 
( 
       SELECT Max(swan_datetime) AS c_date 
       FROM   swan 
       WHERE  swan_datetime BETWEEN $1 AND $2
       AND    deleted_at = To_timestamp(0) 
)
, data AS 
( 
           SELECT     m.id, 
                      m.swan_name->>'th' AS swan_name, 
                      s.swan_datetime::     date, 
                      s.swan_highsig, 
                      m.station_area, 
                      CASE 
                                 WHEN m.id IN (1,2,3,4) then 3 
                                 WHEN m.id IN (5,6,7,8,9) THEN 4 
                                 WHEN m.id IN (10,11,12,13,14,15) THEN 5 
                                 WHEN m.id IN (16,17,18,19) THEN 6 
                                 ELSE 0 
                      END AS region_id, 
                      CASE 
                                 WHEN m.id >= 1 
                                 AND        m.id <= 15 THEN 1 
                                 WHEN m.id >= 16 
                                 AND        m.id <= 19 THEN 2 
                                 ELSE 7 
                      END AS zone_id 
           FROM       swan s 
           INNER JOIN m_swan_station m 
           ON         s.swan_station_id = m.id 
           LEFT  JOIN max_date d 
           ON         d.c_date = s.swan_datetime 
           WHERE      s.swan_datetime BETWEEN $1 AND $2
           AND        s.deleted_at = to_timestamp(0) 
           AND        m.id < 20 
)
, sortable AS 
( 
         SELECT   d.id, 
                  d.swan_name, 
                  d.swan_datetime, 
                  d.swan_highsig, 
                  d.station_area, 
                  d.region_id, 
                  d.zone_id, 
                  rank() OVER (partition BY d.region_id, d.swan_datetime ORDER BY d.swan_highsig DESC) AS max_region,
                  rank() OVER (partition BY d.zone_id, d.swan_datetime ORDER BY d.swan_highsig DESC)   AS max_zone,
                  dense_rank() OVER (ORDER BY d.swan_datetime )                                        AS orders
         FROM     data d 
         WHERE    d.region_id <> 0 
         AND      d.zone_id <> 0 
         ORDER BY d.swan_datetime, 
                  d.region_id, 
                  d.zone_id 
)
, combined AS 
( 
       SELECT id, 
              swan_name, 
              swan_datetime, 
              swan_highsig, 
              station_area, 
              region_id as map_region, 
              orders
       FROM   sortable 
       WHERE  max_region = 1 
       UNION ALL 
       SELECT id, 
              swan_name, 
              swan_datetime, 
              swan_highsig, 
              station_area, 
              zone_id as map_region, 
              orders as odate
       FROM   sortable 
       WHERE  max_zone = 1 
) 
SELECT   swan_name, 
         swan_datetime, 
         swan_highsig, 
         station_area, 
         map_region, 
         orders  
FROM     combined 
ORDER BY swan_datetime, 
         map_region
`
	now := time.Now().Format("2006-01-02")
	next2Day := time.Now().AddDate(0, 0, 2).Format("2006-01-02")
	row, err := db.Query(q, now, next2Day)

	if err != nil {
		return nil, errors.Repack(err)
	}
	var sw *Struct_WaveForecast
	var arrD []*Struct_WaveForecast_Region
	var ss uSetting.Arr_Struct_WaveSetting // wave setting จาก db
	ss, err = uSetting.New_Struct_WaveSetting()
	if err != nil {
		return nil, errors.Repack(err)
	}
	var rs []*Struct_WaveForecast = make([]*Struct_WaveForecast, 0) // result
	var temp_odate int64 = 0                                        // เก็บ odate เพื่อแยก array
	for row.Next() {
		var (
			_swan_name     sql.NullString
			_swan_datetime sql.NullString
			_swan_highsig  sql.NullFloat64
			_station_area  sql.NullString
			_map_region    sql.NullString
			_odate         sql.NullInt64

			d *Struct_WaveForecast_Region
		)
		err = row.Scan(&_swan_name, &_swan_datetime, &_swan_highsig, &_station_area, &_map_region, &_odate)
		if err != nil {
			return nil, errors.Repack(err)
		}
		if temp_odate != _odate.Int64 {
			temp_odate = _odate.Int64

			arrD = make([]*Struct_WaveForecast_Region, 0)

			dt := pqx.NullStringToTime(_swan_datetime)
			sw = &Struct_WaveForecast{
				Date:   dt.Format("02 ") + uDatetime.MonthTHShort(dt.Month()) + dt.AddDate(543, 0, 0).Format(" 2006"),
				Region: arrD,
			}
			rs = append(rs, sw)
		}
		ss_w := ss.CompareSetting(_swan_highsig.Float64)
		if ss_w == nil {
			continue // ความสูงคลื่นไม่เข้าเกณฑ์
		}
		if _map_region.String == "1" {
			_station_area.String = "อ่าวไทยตอนบน"
		} else if _map_region.String == "2" {
			_station_area.String = "อ่าวไทยตอนล่าง"
		}

		d = &Struct_WaveForecast_Region{
			Region_id:    _map_region.String,
			Region_name:  _station_area.String,
			Level:        ss_w.Level,
			Status:       ss_w.Status,
			Station_name: "ที่" + _swan_name.String,
			Color:        ss_w.Color,
		}
		arrD = append(arrD, d)
		sw.Region = arrD
	}
	return rs, nil
}
