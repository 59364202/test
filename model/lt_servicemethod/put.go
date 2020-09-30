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

//	update data
//	Parameters:
//		p
//			Struct_LtServicemethod
//	Return:
//		Struct_LtServicemethod
func PutServiceMethod(p *Struct_LtServicemethod) (*Struct_LtServicemethod, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Convert name to db-json type
	jsonName, err := p.Name.MarshalJSON()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sqlUpdateServiceMethod, p.ID, p.UserId, jsonName)
	if err != nil {
		return nil, err
	}

	return p, nil
}
