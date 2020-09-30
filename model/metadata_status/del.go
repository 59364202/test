// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_status is a model for public.lt_metadata_status table. This table store metadata status information.
package metadata_status

import (
	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// soft deleted metadata status by id
//  Parameters:
//		uid
//			user id
//		method_id
//			metadata status id
//  Return:
//		"Delete Successful"
func DelMetadataStatus(uid int64, method_id string) (string, error) {

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
	q := delMetadataMethod
	p := []interface{}{uid, mid}

	stmt, err := db.Prepare(q)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
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
