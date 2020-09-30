// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_dataunit is a model for public.lt_dataunit table. This table store lt_dataunit information.
package lt_dataunit

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strings"
)

var sqlGetDataunit = " SELECT u.id, u.dataunit_name FROM lt_dataunit u "
var sqlGetDataunit_Where = " WHERE u.deleted_by IS NULL "
var sqlGetDataunit_OrderBy = " ORDER BY u.dataunit_name->>'th' "

//	get data unit
//	Parameters:
//		dataunitId
//			รหัสหน่วยข้อมูล
//	Return:
//		[]Dataunit_struct
func GetDataunit(dataunitId string) ([]*Dataunit_struct, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Get data
	strSQLCmd := sqlGetDataunit_Where
	if dataunitId != "" {
		if len(strings.Split(dataunitId, ",")) == 1 {
			strSQLCmd += " AND u.id = " + dataunitId
		} else {
			strSQLCmd += " AND u.id IN (" + dataunitId + ") "
		}
	}

	_result, err := db.Query(sqlGetDataunit + strSQLCmd + sqlGetDataunit_OrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	var (
		data     []*Dataunit_struct
		dataunit *Dataunit_struct

		_id   sql.NullInt64
		_name sql.NullString
	)

	data = make([]*Dataunit_struct, 0)

	for _result.Next() {
		err := _result.Scan(&_id, &_name)
		if err != nil {
			return nil, err
		}

		dataunit = &Dataunit_struct{}
		dataunit.Id = _id.Int64

		if _name.String == "" {
			_name.String = "{}"
		}
		dataunit.DataunitName = json.RawMessage(_name.String)

		data = append(data, dataunit)
	}

	return data, nil
}
