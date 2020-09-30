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

//	update ministry
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		ministryId
//			ลำดับกระทรวง
//		ministryCode
//			รหัสกระทรวง
//		mapShortTxt
//			ชื่อย่อกระทรวง
//		mapTxt
//			ชื่อกระทรวง
//	Return:
//		Update Successful
func PutMinistry(userId int64, ministryId string, ministryCode string, mapShortTxt, mapTxt json.RawMessage) (string, error) {
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}
	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}
	jsonShortText, err := mapShortTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}
	_, err = db.Exec(sqlUpdateMinistry, ministryCode, jsonShortText, jsonText, userId, ministryId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Update Successful", nil
}
