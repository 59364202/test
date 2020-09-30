// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_method is a model for public.metadata_method table. This table store metadata_method information.
package metadata_method

import (
	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// update metadata method
//  Parameters:
//		uid
//			user id
//		method_id
//			metadata method id
//		method_name
//			metadata method name
//  Return:
//		metadata method id
func PutMetadataMethod(uid int64, method_id, method_name string) (interface{}, error) {
	
	// open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	inf, err := model.GetCipher().DecryptText(method_id)
	if err != nil {
		return nil, errors.New("Can not Delted Invalid Key")
	}
	mid, err := strconv.Atoi(inf)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// sql update metadata method
	q := updateMetadataMethod
	p := []interface{}{uid, mid, method_name}
	// prepare
	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	// execute
	res, err := stmt.Exec(p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	return mid, nil
}
