//     Author: Thitiporn  Meeprasert <thitiporn@hii.or.th>
package wind

import (
	"database/sql"
	"encoding/json"
	//	"fmt"
	"strconv"
	"strings"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

//	get wind
//	Parameters:
//		p
//			Param_Provinces
//	Return:
//		Array Struct_Wind
func GetWind(p *Param_Provinces) ([]*Struct_Wind, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	r := []*Struct_Wind{}
	itf := []interface{}{}

	q := sqlWind

	//Check Filter province_id
	arrProvinceId := []string{}
	if p.Province_Code != "" {
		arrProvinceId = strings.Split(p.Province_Code, ",")
	}
	if len(arrProvinceId) > 0 {
		q += " AND "
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
		q += " AND "
		q += " area_code like '" + p.Region_Code + "' "
	}

	if p.Region_Code_tmd != "" {
		q += " AND "
		q += " tmd_area_code like '" + p.Region_Code_tmd + "' "
	}

	if p.Order != "" {
		q += " ORDER BY wind_speed " + p.Order
	}

	if p.Data_Limit == 0 {
		p.Data_Limit = 1
	}

	q += " LIMIT " + strconv.Itoa(p.Data_Limit)

	//	fmt.Println(q)
	row, err := db.Query(q, itf...)
	if err != nil {
		return nil, err
	}
	strDatetimeFormat := "2006-01-02 15:04"

	for row.Next() {
		var (
			_tele_station_id      int64
			_tele_station_lat     sql.NullFloat64
			_tele_station_long    sql.NullFloat64
			_tele_station_oldcode sql.NullString
			_name                 pqx.JSONRaw
			_amphoe_code          sql.NullString
			_amphoe_name          pqx.JSONRaw
			_tambon_code          sql.NullString
			_tumbon_name          pqx.JSONRaw
			_province_code        sql.NullString
			_province_name        pqx.JSONRaw
			_area_code            sql.NullString
			_area_name            sql.NullString
			_agency_id            sql.NullInt64
			_agency_name          pqx.JSONRaw
			_wind_datetime        sql.NullString
			_wind_speed           sql.NullFloat64
			_wind_dir_value       sql.NullFloat64
			_wind_dir             sql.NullString

			d       *Struct_Wind                  = &Struct_Wind{}
			station *Struct_TeleStation           = &Struct_TeleStation{}
			agency  *model_agency.Struct_Agency   = &model_agency.Struct_Agency{}
			geocode *model_geocode.Struct_Geocode = &model_geocode.Struct_Geocode{}
		)

		err = row.Scan(&_wind_datetime, &_wind_speed, &_wind_dir_value, &_wind_dir, &_tele_station_id, &_tele_station_oldcode, &_name, &_tele_station_lat, &_tele_station_long, &_area_code, &_province_code, &_amphoe_code, &_tambon_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name, &_agency_id, &_agency_name)

		if err != nil {
			return nil, err
		}

		d.WindSpeed = ValidData(_wind_speed.Valid, _wind_speed.Float64)
		d.WindDirValue = ValidData(_wind_dir_value.Valid, _wind_dir_value.Float64)
		d.WindDir = ValidData(_wind_dir.Valid, _wind_dir.String)
		if _wind_datetime.Valid {
			d.WindDatetime = pqx.NullStringToTime(_wind_datetime).Format(strDatetimeFormat)
		}

		station.Id = _tele_station_id
		station.Lat = ValidData(_tele_station_lat.Valid, _tele_station_lat.Float64)
		station.Long = ValidData(_tele_station_long.Valid, _tele_station_long.Float64)
		station.Name = _name.JSON()
		station.OldCode = _tele_station_oldcode.String

		agency.Agency_name = _agency_name.JSON()

		geocode.Tumbon_code = _tambon_code.String
		geocode.Tumbon_name = _tumbon_name.JSON()
		geocode.Amphoe_code = _amphoe_code.String
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

//ตรวจสอบค่าใน column float กรณีค่าเป็น null ให้ return type เป็น interface แทน flat64
//ตัวแปรที่รับค่าจาก db เป็น float64 ซึ่ง float64 เป็น null ไม่ได้ ถ้าใช้  column.float64 จะได้เป็น 0 ทั้งๆที่ใน db เป็น null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
