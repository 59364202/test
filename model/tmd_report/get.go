package tmd_report

import (
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	"haii.or.th/api/thaiwater30/model/lt_geocode"

	"database/sql"
	"encoding/json"
)

//	get Temperature
//	Parameters:
//		prov_id
//			รหัสจังหวัด
func GetTemperature(prov_id string) (*Struct_Temperature, error) {
	s := &Struct_Temperature{}
	s.Init()

	if prov_id == "" {
		return s, nil
	}

	p, err := lt_geocode.GetProvince(prov_id)
	if err != nil {
		return s, errors.Repack(err)
	}
	if p == nil {
		return s, errors.New("invalid prov_id")
	}
	var objmap map[string]string
	err = json.Unmarshal(p.Province_name, &objmap)
	if err != nil {
		return s, errors.Repack(err)
	}
	prov_name := datatype.MakeString(objmap["th"]) // province_name->>'th'

	db, err := pqx.Open()
	if err != nil {
		return s, errors.Repack(err)
	}

	q := `
SELECT public_date 
       , temperature 
FROM   latest.tmd_weather_report 
WHERE  prov_name = $1 
	`

	rows, err := db.Query(q, prov_name)
	if err != nil {
		return s, nil
	}
	for rows.Next() {
		var (
			_date sql.NullString
			_temp sql.NullFloat64
		)
		err = rows.Scan(&_date, &_temp)
		if err != nil {
			return s, errors.Repack(err)
		}
		date := pqx.NullStringToTime(_date)
		s.Temperature = datatype.MakeString(_temp.Float64)
		s.Current_time = date.Format("15:04") + " น."
	}

	return s, nil
}
