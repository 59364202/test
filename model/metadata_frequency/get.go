// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_frequency is a model for public.metadata_frequency table. This table store metadata_frequency information.
package metadata_frequency

import (
	"database/sql"
	"haii.or.th/api/util/pqx"
)

//	Function for delete data in 'metadata_frequency' table with metadata_id
//	Parameters:
//		m_id
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_MetadataFrequency
func GetMetadataFrequency(m_id int64) ([]*Struct_MetadataFrequency, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	var (
		obj *Struct_MetadataFrequency

		_id            int64
		_datafrequency sql.NullString
	)

	row, err := db.Query(SQL_GetFrequencyFromMetadataId, m_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*Struct_MetadataFrequency, 0)
	for row.Next() {
		err = row.Scan(&_id, &_datafrequency)
		if err != nil {
			return nil, err
		}

		obj = &Struct_MetadataFrequency{}
		obj.Id = _id
		obj.Datafrequency = _datafrequency.String

		data = append(data, obj)
	}

	return data, nil
}
