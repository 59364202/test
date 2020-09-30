package amphoe

import (
	"database/sql"
	result "haii.or.th/api/thaiwater30/util/result"
	so "haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlGetAmphoe = "SELECT id , amphoe_name FROM lt_amphoe WHERE is_deleted is null OR is_deleted != TRUE AND province_id = $1 ORDER BY amphoe_code ASC"

func GetAmphoe(provinceId string) (*result.Result, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		_id   sql.NullInt64
		_text sql.NullString
	)
	_result, err := db.Query(sqlGetAmphoe, provinceId)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	var data = so.NewSelect()
	for _result.Next() {
		err := _result.Scan(&_id, &_text)
		if err != nil {
			return nil, err
		}
		data.Add(_text.String, _id.Int64)
	}

	return result.Result1(data.Option), nil
}
