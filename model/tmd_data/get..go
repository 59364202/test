package tmd_data

import (
	"database/sql"
	"errors"
	"haii.or.th/api/util/pqx"
//	"time"
)

type Struct_temperature struct {
	Temperature  string `json:"temperature"`
	Current_time string `json:"current_time"`
}

func Get_temperature(prov_id string) (*Struct_temperature, error) {
	if prov_id == "" {
		return nil, errors.New("No province id")
	}
//	current_date := ""
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// query
//	var q string = `SELECT temperature, to_char(public_date, 'HH24:MI')
//					FROM tmd_weather_report
//					WHERE prov_name = (select province_name ->> 'th' as province_name from lt_geocode where province_code = $1 limit 1)
//					and public_date >= $2
//					and deleted_by is NULL
//					`

	q := `SELECT t1.temperature, to_char(t1.public_date, 'HH24:MI') 
			FROM tmd_weather_report t1
			JOIN 
			(
			   SELECT guid,station_province, MAX(public_date) AS public_date
			   FROM (
				select * 
				from tmd_weather_report a1
				right join (
					select * from tmd_weather_station 
					where station_province = (select province_name ->> 'th' as province_name from lt_geocode where province_code = $1 limit 1)
				) a2
				on a1.guid = a2.station_id
			   ) x1
			   GROUP BY guid,station_province
			) t2
			ON t1.public_date = t2.public_date
			AND t1.guid = t2.guid
			limit 1`

//	current_time := time.Now().Local()
//	current_date = current_time.Format("2006-01-02")

	//	query result
	rows, err := db.Query(q, prov_id)
	if err != nil {
		return nil, err
	}
	//	result function
	var rs *Struct_temperature = &Struct_temperature{}
	rs.Temperature = "-"
	rs.Current_time = "-"

	for rows.Next() {
		var (
			_temp         sql.NullString
			_current_time sql.NullString
		)

		err = rows.Scan(&_temp, &_current_time)
		if err != nil {
			return nil, err
		}

		rs.Temperature = _temp.String
		rs.Current_time = _current_time.String + " น."

	}

	return rs, nil
}
