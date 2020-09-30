// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_ministry is a model for public.lt_ministry table. This table store lt_ministry information.
package lt_ministry

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//	"haii.or.th/api/util/rest"
)

//	get all ministry
//	Return:
//		[]Ministry_struct
func GetMinistry() ([]*Ministry_struct, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		data     []*Ministry_struct
		ministry *Ministry_struct

		_id        sql.NullInt64
		_code      sql.NullString
		_shorttext sql.NullString
		_text      sql.NullString
	)

	_result, err := db.Query(sqlGetMinistry)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Ministry_struct, 0)

	for _result.Next() {
		err := _result.Scan(&_id, &_code, &_shorttext, &_text)
		if err != nil {
			return nil, err
		}
		if !_shorttext.Valid || _shorttext.String == "" {
			_shorttext.String = "{}"
		}
		if !_text.Valid || _text.String == "" {
			_text.String = "{}"
		}

		ministry = &Ministry_struct{Id: _id.Int64, Code: _code.String}
		ministry.MinistryShortName = json.RawMessage(_shorttext.String)
		ministry.MinistryName = json.RawMessage(_text.String)

		data = append(data, ministry)
	}

	return data, nil
}

//	get data
//	Parameters:
//		id
//			ลำดับกระทรวง
//	Return:
//		[]Struct_Ministry
func getMinistry(id int64) ([]*Struct_Ministry, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		strSql string

		data    []*Struct_Ministry
		mnistry *Struct_Ministry

		_id            int64
		_ministry_name sql.NullString
	)
	strSql = SQL_selectMinistry
	if id != 0 {
		strSql += " WHERE id = " + strconv.FormatInt(id, 64)
	}

	row, err := db.Query(strSql)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_ministry_name)
		if err != nil {
			return nil, err
		}

		if !_ministry_name.Valid || _ministry_name.String == "" {
			_ministry_name.String = "{}"
		}

		mnistry = &Struct_Ministry{}
		mnistry.Id = _id
		mnistry.Ministry_Name = json.RawMessage(_ministry_name.String)

		data = append(data, mnistry)
	}

	return data, nil
}

//	get all ministry
//	Return:
//		[]Struct_Ministry
func GetAllMinistry() ([]*Struct_Ministry, error) {
	return getMinistry(0)
}

//	get ministry from id
//	Parameters:
//		id
//			ลำดับกระทรวง
//	Return:
//		[]Struct_Ministry
func GetMinistryFromId(id int64) ([]*Struct_Ministry, error) {
	return getMinistry(id)
}
