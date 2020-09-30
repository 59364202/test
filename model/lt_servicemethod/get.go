// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_servicemethod is a model for public.lt_servicemethod table. This table store lt_servicemethod information.
package lt_servicemethod

import (
	"database/sql"
	"encoding/json"
	model "haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

func GetAllServiceMethod() ([]*Struct_LtServicemethod, error) {
	return getServiceMethod([]int64{})
}

//	get service method
//	Parameters:
//		serviceMethodID
//			[]รหัสวีธีการให้บริการข้อมูล
//	Return:
//		[]Struct_LtServicemethod
func getServiceMethod(serviceMethodID []int64) ([]*Struct_LtServicemethod, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""

	//Check Filter agency_id
	if len(serviceMethodID) > 0 {
		if len(serviceMethodID) == 1 {
			arrParam = append(arrParam, serviceMethodID[0])
			sqlCmdWhere += " AND s.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, intId := range serviceMethodID {
				arrParam = append(arrParam, intId)
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND s.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	var (
		data             []*Struct_LtServicemethod
		objServiceMethod *Struct_LtServicemethod

		_id   sql.NullInt64
		_name sql.NullString

		_result *sql.Rows
	)

	//Query data
	_result, err = db.Query(sqlGetServiceMethod+sqlCmdWhere+sqlGetServiceMethodOrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_LtServicemethod, 0)
	for _result.Next() {
		err := _result.Scan(&_id, &_name)
		if err != nil {
			return nil, err
		}

		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}

		objServiceMethod = &Struct_LtServicemethod{}
		objServiceMethod.ID = _id.Int64
		objServiceMethod.Name = json.RawMessage(_name.String)
		objServiceMethod.ServiceMethodID, _ = model.GetCipher().EncryptText(strconv.FormatInt(_id.Int64, 10))

		data = append(data, objServiceMethod)
	}

	return data, nil
}
