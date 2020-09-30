// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_servicemethod is a model for public.metadata_servicemethod table. This table store metadata_servicemethod information.
package metadata_servicemethod

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/thaiwater30/model/lt_servicemethod"
	"haii.or.th/api/util/pqx"
)

//	get metadata servicemethod
//	Parameters:
//		m_id
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_LtServicemethod
func GetMetadataServicemethod(m_id int64) ([]*lt_servicemethod.Struct_LtServicemethod, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	var (
		obj *lt_servicemethod.Struct_LtServicemethod

		_id                 int64
		_servicemethod_name sql.NullString
	)
	row, err := db.Query(SQL_GetServicemethodFromMetadataId, m_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*lt_servicemethod.Struct_LtServicemethod, 0)
	for row.Next() {
		err = row.Scan(&_id, &_servicemethod_name)
		if err != nil {
			return nil, err
		}
		if !_servicemethod_name.Valid || _servicemethod_name.String == "" {
			_servicemethod_name.String = "{}"
		}

		obj = &lt_servicemethod.Struct_LtServicemethod{}
		obj.ID = _id
		obj.Name = json.RawMessage(_servicemethod_name.String)

		data = append(data, obj)
	}
	return data, nil
}
