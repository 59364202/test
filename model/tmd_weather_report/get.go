package tmd_weather_report

import (
	"database/sql"

	"haii.or.th/api/util/pqx"
	//	"log"
	"errors"
)

//	get weather today
//	Returns:
//		array weather today
func Get_weather_today(year, month, day string) ([]*Struct_Weather_today, error) {
	if year == "" || month == "" || day == "" {
		return nil, errors.New("Invalid parameter!")
	}

	date := year + "-" + month + "-" + day
	dateEnd := date + " 23:59"

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// query
	var q string = `select id,
						wmo_code,
						to_char(date, 'YYYY-MM-DD'),
						mean_sea_level_pressure,
						temperature,
						max_temperature,
						diff_max_temperature,
						min_temperature,
						diff_min_temperature,
						relative_humidity,
						wind_direction,
						wind_speed,
						rainfall
					from tmd_weather_today
					--where date::date = $1
					where date BETWEEN $1 AND $2
					order by id asc`

	//	query result
	rows, err := db.Query(q, date, dateEnd)
	if err != nil {
		return nil, err
	}
	//	result function
	var rs []*Struct_Weather_today = make([]*Struct_Weather_today, 0)

	// loop
	for rows.Next() {
		var (
			_id                      sql.NullInt64
			_wmo_code                sql.NullInt64
			_import_date             sql.NullString
			_mean_sea_level_pressure sql.NullString
			_temperature             sql.NullString
			_max_temerature          sql.NullString
			_diff_max_temperature    sql.NullString
			_min_temperature         sql.NullString
			_diff_min_temoerature    sql.NullString
			_relative_humidity       sql.NullString
			_wind_direction          sql.NullString
			_wind_speed              sql.NullString
			_rainfall                sql.NullString
		)
		err = rows.Scan(
			&_id,
			&_wmo_code,
			&_import_date,
			&_mean_sea_level_pressure,
			&_temperature,
			&_max_temerature,
			&_diff_max_temperature,
			&_min_temperature,
			&_diff_min_temoerature,
			&_relative_humidity,
			&_wind_direction,
			&_wind_speed,
			&_rainfall,
		)
		if err != nil {
			return nil, err
		}

		// create struct and add data into struct
		s := &Struct_Weather_today{
			Id:                     _id.Int64,
			Wmo_code:               _wmo_code.Int64,
			Import_date:            _import_date.String,
			Mean_sea_level_presure: _mean_sea_level_pressure.String,
			Temperature:            _temperature.String,
			Max_temperature:        _max_temerature.String,
			Diff_max_temperature:   _diff_max_temperature.String,
			Min_temperature:        _min_temperature.String,
			Diff_min_temperature:   _diff_min_temoerature.String,
			Relative_humidity:      _relative_humidity.String,
			Wind_direction:         _wind_direction.String,
			Wind_speed:             _wind_speed.String,
			Rainfall:               _rainfall.String,
		}
		// append struct into array
		rs = append(rs, s)
	}

	return rs, nil
}
