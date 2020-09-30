package servicemethod

import (
	"database/sql"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strings"
)

var sqlGetServiceMethod = " SELECT s.id, l.code, st.servicemethod_name " +
	" FROM lt_servicemethod s " +
	" LEFT JOIN lt_servicemethod_translate st ON s.id = st.method_id " +
	" LEFT JOIN language l ON l.id = st.language_id " +
	" WHERE s.deleted_by IS NULL "
var sqlGetServiceMethodOrderBy = " ORDER BY s.id ,l.id "

func GetServiceMethod(serviceMethodId string) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Get data
	strSQLCmd := sqlGetServiceMethod
	if serviceMethodId != "" {
		if len(strings.Split(serviceMethodId, ",")) == 1 {
			strSQLCmd += " AND s.id = " + serviceMethodId
		} else {
			strSQLCmd += " AND s.id IN (" + serviceMethodId + ") "
		}
	}

	_result, err := db.Query(strSQLCmd + sqlGetServiceMethodOrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	var (
		id            int64
		text          map[string]string
		data          []*ServiceMethod_struct
		serviceMethod *ServiceMethod_struct

		_id   sql.NullInt64
		_code sql.NullString
		_text sql.NullString
	)

	data = make([]*ServiceMethod_struct, 0)
	id = 0

	for _result.Next() {
		err := _result.Scan(&_id, &_code, &_text)
		if err != nil {
			return nil, err
		}

		if id != _id.Int64 {
			if id != 0 {
				serviceMethod.ServiceMethodName = text
				data = append(data, serviceMethod)
			}
			id = _id.Int64
			serviceMethod = &ServiceMethod_struct{}
			serviceMethod.Id = id
			text = make(map[string]string)
		}

		if _text.Valid {
			text[_code.String] = _text.String
		}
	}

	if id != 0 {
		serviceMethod.ServiceMethodName = text
		data = append(data, serviceMethod)
	}

	return result.Result1(data), nil
}
