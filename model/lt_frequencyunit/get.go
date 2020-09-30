// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_frequencyunit is a model for public.lt_frequencyunit table. This table store lt_frequencyunit information.
package lt_frequencyunit

import (
	"database/sql"
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlGetFrequencyUnit = " SELECT f.id, f.frequencyunit_name, f.convert_minute FROM lt_frequencyunit f "
var sqlGetFrequencyUnit_Where = " WHERE f.deleted_by IS NULL "
var sqlGetFrequencyUnit_OrderBy = " ORDER BY f.frequencyunit_name->>'th' "

//	get frequencyunit
//	Parameters:
//		frequencyUnitId
//			รหัสหน่วยความถี่
//	Return:
//		[]FrequencyUnit_struct
func GetFrequencyUnit(frequencyUnitId string) ([]*FrequencyUnit_struct, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Where Condition
	sqlCmdWhere := sqlGetFrequencyUnit_Where
	if frequencyUnitId != "" {
		sqlCmdWhere += " AND f.id = " + frequencyUnitId
	}

	//Execute query
	_result, err := db.Query(sqlGetFrequencyUnit + sqlCmdWhere + sqlGetFrequencyUnit_OrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	var (
		data          []*FrequencyUnit_struct
		frequencyUnit *FrequencyUnit_struct

		_id            sql.NullInt64
		_name          sql.NullString
		_convertMinute sql.NullInt64
	)

	data = make([]*FrequencyUnit_struct, 0)

	for _result.Next() {
		err := _result.Scan(&_id, &_name, &_convertMinute)
		if err != nil {
			return nil, err
		}
		frequencyUnit = &FrequencyUnit_struct{}
		frequencyUnit.Id = _id.Int64
		frequencyUnit.ConvertMinute = _convertMinute.Int64

		if _name.String == "" {
			_name.String = "{}"
		}
		frequencyUnit.FrequencyUnitName = json.RawMessage(_name.String)

		data = append(data, frequencyUnit)
	}

	return data, nil
}
