// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_servicemethod is a model for public.lt_servicemethod table. This table store lt_servicemethod information.
package lt_servicemethod

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlUpdateToDeleteServiceMethod = ` UPDATE lt_servicemethod
										  SET deleted_by = $2,
											  deleted_at = NOW(),
											  updated_by = $2 ,
											  updated_at = NOW()
										  WHERE id = $1 `

//	delete data
//	Parameters:
//		p
//			Struct_LtServicemethod
//	Return:
//		Delete Successful
func DeleteServiceMethod(p *Struct_LtServicemethod) (string, error) {
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	_, err = db.Exec(sqlUpdateToDeleteServiceMethod, p.ID, p.UserId)
	if err != nil {
		return "", err
	}

	return "Delete Successful", nil
}
