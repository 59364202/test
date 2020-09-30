// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_status is a model for public.lt_metadata_status table. This table store metadata status information.
package metadata_status

import (
	"encoding/json"
	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// update metadata status by id
//  Parameters:
//		uid
//			user id
//		method_id
//			metadata status id
//		status_name
//			metadata status name
//  Return:
//		metadata status id
func PutMetadataStatus(uid int64, method_id string, status_name json.RawMessage) (interface{}, error) {

	// opend database
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

	// convert map[string]interface{} to json.rawmessage
	b, err := status_name.MarshalJSON()
	if err != nil {
		return nil, err
	}

	// sql update data
	q := updateMetadataMethod
	p := []interface{}{uid, mid, b}

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
