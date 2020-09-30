package a

import (
	"database/sql"
	"encoding/json"
	
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

type fern_province struct {
	Province_id    interface{}     `json:"province_id"`    // example:`81` รหัสจังหวัด
	Region_id      interface{}     `json:"region_id"`      // example:`1` รหัสภาค
	Province_name  json.RawMessage `json:"province_name"`  // example:`{"th":"กระบี่"}` ชื่อจังหวัด
	Province_name2 json.RawMessage `json:"province_name2"` // example:`{"th":"กระบี่"}` ชื่อจังหวัด2
}

type Param_handlerGetFern struct {
	Province_id string `json:"province_id"` // example:`81` required:false รหัสจังหวัด
	Region_id   string `json:"region_id"` // example:`5` required:false รหัสภาค
}

//func Fern_Province(prov_code string) ([]*fern_province, error) { // func รับค่าธรรมดา
func Fern_Province(param *Param_handlerGetFern) ([]*fern_province, error) {
	db, err := pqx.Open()
	if err != nil {
		// ใส่  error ลง log
		return nil, errors.Repack(err)
	}

	/*if param.Province_id == "" || param.Region_id == "" {
		return nil, errors.New("invalid param")
	}*/

	p := []interface{}{}
	//query
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

	//var result []*cim_province = make([]*cim_province,0)
	rs := make([]*fern_province, 0)

	for rows.Next() {
		var (
			_province_code  sql.NullInt64
			_area_code      sql.NullInt64
			_province_name2 sql.NullString //ตัวเดิม
			_province_name  pqx.JSONRaw    //ทำขึ้นเอง
		)
		err = rows.Scan(&_province_code, &_area_code, &_province_name, &_province_name2)
		if err != nil {
			return nil, errors.Repack(err)
		}
		// Region_id เป็น  interface{}
		p := &fern_province{
		//Province_id: _province_code.Int64,
		//Province_id:   _province_code.Int64,
		//Region_id:     _area_code.Int64,
		//Province_name: _province_name.String,
		}
		// start JSON ทำขึ้นเอง
		p.Province_name = _province_name.JSON()
		if !_province_name2.Valid {
			_province_name2.String = "{}"
		}

		p.Province_name2 = json.RawMessage(_province_name2.String)
		// ทำได้สองวิธี
		// 1. ใช้ฟังก์ชั่น  Value()
		p.Region_id, _ = _area_code.Value()
		p.Province_id, _ = _province_code.Value()

		// 2. เช็คว่าถ้าไม่ใช่  null ให้ใส่ค่า
		/*if _area_code.Valid{
			p.Region_id = _area_code.Int64
		}*/

		//pr := new(cim_province)
		//pr.Province_id = _province_code

		rs = append(rs, p)

	}
	return rs, err
}
