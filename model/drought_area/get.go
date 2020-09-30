// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package drought_area is a model for public.drought_area table. This table store drought_area.
package drought_area

import (
	"database/sql"
	"haii.or.th/api/util/pqx"
)

//	get drought area วันที่ล่าสุด
//	Return:
//		[]Struct_DroughtArea
func GetLatestDrought() ([]*Struct_DroughtArea, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		data []*Struct_DroughtArea
		obj  *Struct_DroughtArea
	)

	data = make([]*Struct_DroughtArea, 0)
	row, err := db.Query(SQL_GetLatestDrought)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var (
			_provinceCode sql.NullString
			_provinceName pqx.JSONRaw
		)
		err = row.Scan(&_provinceCode, &_provinceName)
		if err != nil {
			return nil, err
		}
		obj = &Struct_DroughtArea{}
		obj.ProvinceCode = _provinceCode.String
		obj.ProvinceName = _provinceName.JSONPtr()

		data = append(data, obj)
	}

	return data, nil
}
