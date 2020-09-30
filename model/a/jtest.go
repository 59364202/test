package a

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	"database/sql"
	"encoding/json"
	"strconv"
)

type jtest_province struct {
	Province_id    interface{}     `json:"province_id"`    // example:`81` มันเป็นรหัสจังหวัดนะ
	Region_id      interface{}     `json:"region_id"`      // example:`5` มันเป็นรหัสภาคน
	Province_name  json.RawMessage `json:"province_name"`  // example:`{"th":"กระบี่"}` มันเป็นชื่อจังหวัดนะ
	Province_name2 json.RawMessage `json:"province_name2"` // example:`{"th":"กระบี่"}` มันเป็นชื่อจังหวัดนะ
}

type Param_handlerGetJuiss struct {
	Province_id string `json:"province_id"` // example:`81` required:false รหัสจังหวัดนะ
	Region_id   string `json:"region_id"`   // example:`5` มันเป็นรหัสภาคนะ
}

func Jtest_Province(param *Param_handlerGetJuiss) ([]*jtest_province, error) {
	db, err := pqx.Open()
	if err != nil {
		//	ใส่ error ลง log
		return nil, errors.Repack(err)
	}

//	if param.Province_id == "" || param.Region_id == "" {
//		return nil, errors.New("invalid param")
//	}

	// query
	p := []interface{}{}
	var q string = `
	SELECT province_code, area_code, province_name, province_name
	FROM public.lt_geocode
	WHERE 
	(
		(amphoe_code = '  ' AND tumbon_code = '  ') 
		OR (amphoe_name->>'th' = '' AND tumbon_name->>'th' = '')
	)`

	if param.Region_id != "" {
		p = append(p, param.Region_id)
		q += " AND area_code = $" + strconv.Itoa(len(p))

	}

	if param.Province_id != "" {
		p = append(p, param.Province_id)
		q += " AND province_code = $" + strconv.Itoa(len(p))
	}

	q += " ORDER BY province_name->>'th' ASC "

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//var result []*jtest_province = make([]*jtest_province,0)
	rs := make([]*jtest_province, 0)
	for rows.Next() {
		var (
			_province_code  sql.NullInt64
			_area_code      sql.NullInt64
			_province_name  pqx.JSONRaw
			_province_name2 sql.NullString
		)
		err = rows.Scan(&_province_code, &_area_code, &_province_name, &_province_name2)
		if err != nil {
			return nil, errors.Repack(err)
		}

		p := &jtest_province{
		//Province_id:   _province_code.Int64,
		//Region_id:     _area_code,
		//Province_name: _province_name.String,
		}

		p.Province_name = _province_name.JSON()

		if !_province_name2.Valid {
			_province_name2.String = "{}"
		}

		p.Province_name2 = json.RawMessage(_province_name2.String)

		p.Region_id, _ = _area_code.Value()
		p.Province_id, _ = _province_code.Value()

		//if _area_code.Valid {
		//	p.Region_id = _area_code.Int64
		//}
		rs = append(rs, p)
	}
	return rs, err
}
