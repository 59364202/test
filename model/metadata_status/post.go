// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_status is a model for public.lt_metadata_status table. This table store metadata status information.
package metadata_status

import (
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// new metadata status
//  Parameters:
//		uid
//			user id
//		status_name
//			metadata status name
//  Return:
//		metadata status id
func PostMetadataStatus(uid int64, status_name json.RawMessage) (interface{}, error) {

	// open dabase
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// convert map[string]interface{} to json.rawmessage
	b, err := status_name.MarshalJSON()
	if err != nil {
		return nil, err
	}

	// sql insert data
	q := insMetadataMethod
	p := []interface{}{b, uid}

	var methodID int64
	err = db.QueryRow(q, p...).Scan(&methodID)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	// return metadata status id
	return methodID, nil
}
