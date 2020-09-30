// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_ministry is a model for public.lt_ministry table. This table store lt_ministry information.
package lt_ministry

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	insert ministry
//	Parameters:
//		mapTxt
//			ชื่อกระทรวง
//		mapShortTxt
//			ชื่อย่อกระทรวง
//		code
//			รหัสกระทรวง
//		userId
//			รหัสผู้ใช้งาน
//	Return:
//		Ministry_struct
func PostMinistry(mapTxt, mapShortTxt json.RawMessage, code string, userId int64) (*Ministry_struct, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}
	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return nil, errors.Repack(err)
	}
	jsonShortText, err := mapShortTxt.MarshalJSON()
	if err != nil {
		return nil, errors.Repack(err)
	}

	var newId int64
	err = db.QueryRow(sqlInsertMinistry, code, jsonShortText, jsonText, userId).Scan(&newId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	data := &Ministry_struct{Id: newId, MinistryName: mapTxt, MinistryShortName: mapShortTxt}
	return data, nil
}
