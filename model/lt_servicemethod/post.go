// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_servicemethod is a model for public.lt_servicemethod table. This table store lt_servicemethod information.
package lt_servicemethod

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	"database/sql"
	"strconv"

	model "haii.or.th/api/server/model"
)

var sqlInsertServiceMethod = ` INSERT INTO lt_servicemethod (created_by, updated_by, servicemethod_name) VALUES ($1, $1, $2) RETURNING id `

//	insert data
//	Parameters:
//		p
//			Struct_LtServicemethod
//	Return:
//		Struct_LtServicemethod
func PostServiceMethod(p *Struct_LtServicemethod) (*Struct_LtServicemethod, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}
	//Convert name to db-json type
	jsonName, err := p.Name.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var _id sql.NullInt64
	err = db.QueryRow(sqlInsertServiceMethod, p.UserId, jsonName).Scan(&_id)
	if err != nil {
		return nil, err
	}
	p.ID = _id.Int64
	p.Name = p.Name
	p.UserId = p.UserId
	p.ServiceMethodID, _ = model.GetCipher().EncryptText(strconv.FormatInt(_id.Int64, 10))

	return p, nil
}
