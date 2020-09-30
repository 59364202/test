package a

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx" // db connection

	"database/sql"
	"encoding/json"
	"strconv"
)

// define struct, Capital letter is public, Normal letter is private
type m_province struct {
	Province_id int64       `json:"province_id"`	// example:`81` Province ID
	Region_id   interface{} `json:"region_code"` // example:`5` Region ID
	// Province_name string      `json:"province_name"`
	Province_name  json.RawMessage `json:"province_name"`	// example:`{"th": "กระบี่"}` Province Name
	Province_name2 json.RawMessage
}

// parameter by querystring, require struct
// ex, ?region_id=5&province_id=10
type Param_handlerGetManorot struct {
	Province_id string `json:"province_id"`	// example:`81` required:false Province ID
	Region_id   string `json:"region_id"`	// example:`5` required:false Region ID
}

// return array of province structure
// func M_Province(prov_code string) ([]*m_province, error) {
func M_Province(param *Param_handlerGetManorot) ([]*m_province, error) {
	db, err := pqx.Open()

	if err != nil {
		// return err
		// log error and display error
		return nil, errors.Repack(err)
	}

	// query
	p := []interface{}{}
	//	var q string = `
	//	SELECT province_code, area_code, province_name->>'th'
	//	FROM public.lt_geocode
	//	WHERE (
	//		(amphoe_code = ' ' AND tumbon_code = ' ')
	//		 OR
	//		(amphoe_name->>'th' = '' AND tumbon_name->>'th' = '')
	//	)`

	var q string = `
	SELECT province_code, area_code, province_name, province_name
	FROM public.lt_geocode
	WHERE (
		(amphoe_code = ' ' AND tumbon_code = ' ')
		 OR 
		(amphoe_name->>'th' = '' AND tumbon_name->>'th' = '')
	)`

	if param.Region_id != "" {
		p = append(p, param.Region_id)
		q += " AND area_code = $" + strconv.Itoa(len(p))

	}
	if param.Province_id != "" {
		p = append(p, param.Province_id)
		q += " AND province_code = $" + strconv.Itoa(len(p))
	}

	q += "ORDER BY province_name->>'th' ASC"

	// query data from sql
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, errors.Repack(err)
	}

	// array of struct
	// var result []*province = make([]*province, 0)
	rs := make([]*m_province, 0)

	// loop throught output
	for rows.Next() {
		// declare variables
		var (
			_province_code sql.NullInt64
			_area_code     sql.NullInt64
			// _province_name sql.NullString
			_province_name  pqx.JSONRaw    // custom class
			_province_name2 sql.NullString // origin from go
		)

		// loop through row and kept data to variables
		err = rows.Scan(&_province_code, &_area_code, &_province_name, &_province_name2)
		if err != nil {
			return nil, errors.Repack(err)
		}

		// assign data
		p := &m_province{
			Province_id: _province_code.Int64,
			// Region_id:     _area_code,
			// Province_name: _province_name.String, // must has comma at the end
			Province_name: _province_name.JSON(),
		}

		// recieve json from pql.JSONRaw (custom function)
		p.Province_name = _province_name.JSON()

		// recieve json (original function)
		if !_province_name2.Valid {
			_province_name2.String = "{}"
		}
		p.Province_name2 = json.RawMessage(_province_name2.String)

		// variable that can recieve null (interface) can do 2 ways
		// 1. assign by Value function
		p.Region_id, _ = _area_code.Value()

		// 2. check if not null then assign value
		// if _area_code.Valid {
		//		p.Region_id, _ = _area_code.Int64
		// }

		// append data object to array
		rs = append(rs, p)

		// pr := new(province)
		// pr.Province_id = _province_code
	}
	return rs, nil
}
