// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package flood_situation is a model for public.flood_situation table. This table store flood_situation.
package flood_situation

import (
	"database/sql"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/pqx"
)

//	get flood situation วันล่าสุด
//	Return:
//		[]Struct_FloodSituation
func GetLatestFloodSituation() ([]*Struct_FloodSituation, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		data []*Struct_FloodSituation
		obj  *Struct_FloodSituation
	)

	data = make([]*Struct_FloodSituation, 0)
	row, err := db.Query(SQL_GetLatestFloodSituation)
	if err != nil {
		return nil, err
	}
	//	datetime format
	layout := setting.GetSystemSetting("setting.Default.DatetimeFormat")
	for row.Next() {
		var (
			_provinceCode     sql.NullString
			_provinceName     pqx.JSONRaw
			_floodDatetime    sql.NullString
			_floodName        sql.NullString
			_floodLink        sql.NullString
			_floodDescription sql.NullString
			_floodAuthor      sql.NullString
			_floodColorlevel  sql.NullString
		)
		err = row.Scan(&_provinceCode, &_provinceName, &_floodDatetime, &_floodName, &_floodLink, &_floodDescription, &_floodAuthor, &_floodColorlevel)
		if err != nil {
			return nil, err
		}

		obj = &Struct_FloodSituation{}
		obj.ProvinceCode = _provinceCode.String
		obj.ProvinceName = _provinceName.JSONPtr()
		obj.FloodName = _floodName.String
		obj.FloodLink = _floodLink.String
		obj.FloodDescription = _floodDescription.String
		obj.FloodAuthor = _floodAuthor.String
		obj.FloodColorlevel = _floodColorlevel.String

		if _floodDatetime.Valid {
			obj.FloodDatetime = pqx.NullStringToTime(_floodDatetime).Format(layout)
		}

		data = append(data, obj)
	}

	return data, nil
}
