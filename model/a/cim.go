package a

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	"database/sql"
	"encoding/json"
	"strconv"
)

type cim_province struct {
	Province_id   int64           `json:"province_id"`   // example:`81` มันเป็นรหัสจังหวัดนะ
	Region_id     interface{}     `json:"region_id"`     // example:`5` มันเป็นรหัสภาคนะ
	Province_name json.RawMessage `json:"province_name"` // example:`{"th":"กระบี่"}` มันเป็นชื่อจังหวัดนะ
}
type Param_handlerGetCim struct {
	Province_id string `json:"province_id"` // example:`81` required:false รหัสจังหวัดนะ
	Region_id   string `json:"region_id"`   // example:`5` required:false มันเป็นรหัสภาคนะ
}

func Cim_Province(param *Param_handlerGetCim) ([]*cim_province, error) {
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
	SELECT province_code, null, province_name
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

	rs := make([]*cim_province, 0)

	for rows.Next() {
		// import "database/sql"
		// import "encoding/json"
		var (
			_province_code sql.NullInt64
			_area_code     sql.NullInt64
			_province_name pqx.JSONRaw // ทำขึ้นเอง
		)
		err = rows.Scan(&_province_code, &_area_code, &_province_name)
		if err != nil {
			return nil, errors.Repack(err)
		}
		// Region_id เป็น interface{}
		p := &cim_province{
			Province_id: _province_code.Int64,
		}
		// start json
		// ทำขึ้นเอง
		p.Province_name = _province_name.JSON()

		// end json

		// start null
		// ทำได้ 2 วิธี
		// 1. ใช้ฟังค์ชั่น Value()
		p.Region_id, _ = _area_code.Value()

		// 2. เช็คว่าถ้าไม่ใช่ null ให้ใส่ค่า
		if _area_code.Valid {
			p.Region_id = _area_code.Int64
		}
		// end null
		rs = append(rs, p)
	}
	return rs, err
}
