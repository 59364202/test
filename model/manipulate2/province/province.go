package province

import (
	"database/sql"
	result "haii.or.th/api/thaiwater30/util/result"
	so "haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlGetProvince = "SELECT id , province_name FROM lt_province WHERE region_id = $1 ORDER BY province_code ASC"

func GetProvince(regionId string) (*result.Result, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		_id   sql.NullInt64
		_text sql.NullString
	)
	_result, err := db.Query(sqlGetProvince, regionId)
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
		data.Add(_id.Int64, _text.String)
	}

	return result.Result1(data.Option), nil
}
