package rainfall_yearly

import (
	"database/sql"
	"encoding/json"
	//	"time"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/sqltime"
	"haii.or.th/api/util/pqx"
	//	"haii.or.th/api/util/rest"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
)

//	get rainfall yearly
//	Parameters:
//		p
//			Param_RainfallYearly
//	Return:
//		Array Struct_RainfallYearly
func GetRain(p *Param_RainfallYearly) ([]*Struct_RainfallYearly, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows
		obj *Struct_RainfallYearly

		_tele_station_id      int64
		_tele_station_name    sql.NullString
		_rainfall_value       sql.NullFloat64
		_rainfall_datetime    sqltime.NullTime
		_province_code        sql.NullString
		_province_name        sql.NullString
		_tele_station_lat     sql.NullFloat64
		_tele_station_long    sql.NullFloat64
		_amphoe_name          sql.NullString
		_tumbon_name          sql.NullString
		_agency_name          sql.NullString
		_agency_shortname     sql.NullString
		_data_id              int64
		_tele_station_oldcode sql.NullString
		_basin_id   		sql.NullInt64
		_basin_code 		sql.NullInt64
		_basin_name 		sql.NullString
		_sub_basin_id 		sql.NullInt64
	)
	strSql, param := Gen_SQL_SelectRain(p)
	row, err = db.Query(strSql, param...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*Struct_RainfallYearly, 0)
	for row.Next() {
		err = row.Scan(&_tele_station_id, &_tele_station_name, &_rainfall_value, &_rainfall_datetime, &_province_code, &_province_name, &_tele_station_lat, &_tele_station_long,
			&_amphoe_name, &_tumbon_name, &_agency_name, &_agency_shortname, &_data_id, &_tele_station_oldcode, &_basin_id, &_basin_code, &_basin_name, &_sub_basin_id)
		if err != nil {
			return nil, err
		}
		if !_tele_station_name.Valid || _tele_station_name.String == "" {
			_tele_station_name.String = "{}"
		}
		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}
		if !_amphoe_name.Valid || _amphoe_name.String == "" {
			_amphoe_name.String = "{}"
		}
		if !_tumbon_name.Valid || _tumbon_name.String == "" {
			_tumbon_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		obj = &Struct_RainfallYearly{}
		obj.Id = _data_id
		obj.StationType = "rainfall_yearly"
		obj.Time = _rainfall_datetime.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		if _rainfall_value.Valid {
			obj.RainfallValue = _rainfall_value.Float64
		}

		station := &Struct_TeleStation{Id: _tele_station_id}
		station.Name = json.RawMessage(_tele_station_name.String)
		station.Lat, _ = _tele_station_lat.Value()
		station.Long, _ = _tele_station_long.Value()
		station.OldCode = _tele_station_oldcode.String
		station.SubBasin_id = _sub_basin_id.Int64
		obj.Station = station

		agency := &model_agency.Struct_Agency{}
		agency.Agency_name = json.RawMessage(_agency_name.String)
		agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		obj.Agency = agency

		geocode := &model_geocode.Struct_Geocode{}
		geocode.Province_code = _province_code.String
		geocode.Province_name = json.RawMessage(_province_name.String)
		geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)
		geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)
		obj.Geocode = geocode
		
		obj.Basin = &model_basin.Struct_Basin{}
		obj.Basin.Id = _basin_id.Int64
		obj.Basin.Basin_code = _basin_code.Int64

		if !_basin_name.Valid || _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		obj.Basin.Basin_name = json.RawMessage(_basin_name.String)

		//		obj.StationId = _tele_station_id
		//		obj.StationName = json.RawMessage(_tele_station_name.String)
		//		obj.Rainfall = _rainfall_value.Float64
		//		obj.DateTime = _rainfall_datetime.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		//		obj.ProvinceCode = _province_code.String
		//		obj.ProvinceName = json.RawMessage(_province_name.String)
		//		obj.StationLat = _tele_station_lat.String
		//		obj.StationLong = _tele_station_long.String
		//		obj.AmphoeName = json.RawMessage(_amphoe_name.String)
		//		obj.TumbonName = json.RawMessage(_tumbon_name.String)
		//		obj.AgencyName = json.RawMessage(_agency_name.String)

		data = append(data, obj)
	}
	return data, nil
}
