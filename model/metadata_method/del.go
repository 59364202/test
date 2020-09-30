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

// soft deleted metadata method
//  Parameters:
//		uid
//			user id for update data
//		method_id
//			metadata method id for soft deleted data
//  Return:
//		"Delete Successful"
func DelMetadataMethod(uid int64, method_id string) (string, error) {

	// open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	inf, err := model.GetCipher().DecryptText(method_id)
	if err != nil {
		return "", errors.New("Can not deleted data invalid key")
	}
	mid, err := strconv.Atoi(inf)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	// get sql deleted metadata method
	q := delMetadataMethod
	p := []interface{}{uid, mid}

	// prepare statement
	stmt, err := db.Prepare(q)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	// execute
	res, err := stmt.Exec(p...)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	return "Delete Successful", nil
}
