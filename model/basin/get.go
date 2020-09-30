// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package basin is a model for public.basin table. This table store basin.
package basin

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/pqx"
)

//	get basin ทั้งหมดจากตารางที่ บัญชีข้อมูลนี้ผูกอยู่
//	Parameters:
//		meta_id
//			metadata id
//	Return:
//		[]Struct_Basin
func GetAllBasinFromMetaId(meta_id int64) ([]*Struct_Basin, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// query
	row, err := db.Query(SQL_selectAllBasinFromMetaId, meta_id)
	if err != nil {
		return nil, err
	}

	data, err := scanRowToBasin(row)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//	get basin from basin code
//	Parameters:
//		code
//			basin code
//	Return:
//		[]Struct_Basin
func GetBasinFromCode(code string) ([]*Struct_Basin, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	row, err := db.Query(SQL_selectBasinFromCode + code + SQL_selectBasinFromCode_end)
	if err != nil {
		return nil, err
	}
	data, err := scanRowToBasin(row)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//	scan query result to struct
//	Parameters:
//		row
//			sql.rows จากการ query
//	Return:
//		[]Struct_Basin
func scanRowToBasin(row *sql.Rows) ([]*Struct_Basin, error) {
	var (
		data  []*Struct_Basin
		basin *Struct_Basin

		_id         int64
		_basin_code int64
		_basin_name sql.NullString
	)

	for row.Next() {
		err := row.Scan(&_id, &_basin_code, &_basin_name)
		if err != nil {
			return nil, err
		}

		if !_basin_name.Valid || _basin_name.String == "" {
			_basin_name.String = "{}"
		}

		basin = &Struct_Basin{}
		basin.Id = _id
		basin.Basin_code = _basin_code
		basin.Basin_name = json.RawMessage(_basin_name.String)

		data = append(data, basin)
	}
	return data, nil
}

//	get all basin
//	Return:
//		[]Struct_Basin
func GetAllBasin() ([]*Struct_Basin, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	row, err := db.Query(SQL_selectAllBasin + " ORDER BY basin_name->>'th' ")
	if err != nil {
		return nil, err
	}

	data, err := scanRowToBasin(row)
	if err != nil {
		return nil, err
	}

	return data, nil
}
