// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_ministry is a model for public.lt_ministry table. This table store lt_ministry information.
package lt_ministry

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//	delete data
//	Parameters:
//		userId
//			รหัสผู้ใช้งาน
//		ministryId
//			ลำดับกระทรวง
//	Return:
//		Delete Successful
func DeleteMinistry(userId int64, ministryId string) (string, error) {
	hasChild := false
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	row, err := db.Query(sqlCheckChild, ministryId)
	if err != nil {
		return "", errors.Repack(err)
	}
	for row.Next() {
		hasChild = true
	}

	if hasChild {
		return "", rest.NewError(422, "กระทรวงได้ถูกใช้โดย บัญชีข้อมูล", nil)
	}

	_, err = db.Exec(sqlUpdateToDeleteMinistry, ministryId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "Delete Successful", nil
}
