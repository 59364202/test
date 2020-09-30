// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataformat is a model for public.lt_dataformat This table store dataformat information.
package dataformat

import (
	"database/sql"
	"encoding/json"
	//	"fmt"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlGetDataFormat = " SELECT d.id, d.dataformat_name,lmm.id,lmm.metadata_method_name FROM lt_dataformat d LEFT JOIN lt_metadata_method lmm ON d.metadata_method_id=lmm.id "
var sqlGetDataformatWhere = " WHERE d.deleted_at=to_timestamp(0) "
var sqlGetDataFormatOrderBy = " ORDER BY d.dataformat_name->>'en' "

//  get data
//	Parameters:
//		dataformatId
//			รหัสรูปแบบข้อมูล
//		metdata_method_id
//			รหัสวิธีการได้มาซึ่งข้อมูล
//	Return:
//		[]Dataformat_struct
func GetDataformat(dataformatId string, metdata_method_id []int64) ([]*Dataformat_struct, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Check dataformatId
	sqlCmdWhere := sqlGetDataformatWhere
	if dataformatId != "" {
		sqlCmdWhere += " AND d.id = " + dataformatId + " "
	}
	//	fmt.Println(metdata_method_id)
	p := []interface{}{}
	if len(metdata_method_id) > 0 {
		var condition_method_id string
		for i, v := range metdata_method_id {
			if i != 0 {
				condition_method_id += " OR d.metadata_method_id=$" + strconv.Itoa(i+1)
			} else {
				condition_method_id = " d.metadata_method_id=$" + strconv.Itoa(i+1)
			}
			p = append(p, v)
		}
		sqlCmdWhere += " AND (" + condition_method_id + ")"
	}
	//	fmt.Println(sqlGetDataFormat + sqlCmdWhere + sqlGetDataFormatOrderBy)
	//Execute Query
	_result, err := db.Query(sqlGetDataFormat+sqlCmdWhere+sqlGetDataFormatOrderBy, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Variables
	var (
		data       []*Dataformat_struct
		dataformat *Dataformat_struct

		_id                   sql.NullInt64
		_dataformat_name      sql.NullString
		_metadata_method_id   sql.NullInt64
		_metadata_method_name sql.NullString
	)

	data = make([]*Dataformat_struct, 0)

	// Loop data result
	for _result.Next() {
		err := _result.Scan(&_id, &_dataformat_name, &_metadata_method_id, &_metadata_method_name)
		if err != nil {
			return nil, err
		}

		//Generate Dataformat object
		dataformat = &Dataformat_struct{}
		dataformat.Id = _id.Int64
		if _metadata_method_id.Valid {
			dataformat.MetadataMethodID = _metadata_method_id.Int64
		}
		dataformat.MetadataMethodName = _metadata_method_name.String
		if _dataformat_name.String == "" {
			_dataformat_name.String = "{}"
		}
		dataformat.DataformatName = json.RawMessage(_dataformat_name.String)

		data = append(data, dataformat)
	}

	return data, nil
}

var sqlGetMetadataMethodOption = "SELECT id,metadata_method_name FROM public.lt_metadata_method WHERE deleted_at=to_timestamp(0)"

//  get วิธีการได้มาซึ่งข้อมูล
//	Return:
//		[]Struct_Option
func GetSelectOption() ([]*Struct_Option, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// sql get download config
	q := sqlGetMetadataMethodOption
	p := []interface{}{}

	// process sql get dataimport download
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	sopt := make([]*Struct_Option, 0)
	for rows.Next() {
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		dataRow := &Struct_Option{}
		rows.Scan(&id, &name)
		dataRow.Text = name.String
		dataRow.Value = id.Int64

		sopt = append(sopt, dataRow)
	}

	return sopt, nil
}
