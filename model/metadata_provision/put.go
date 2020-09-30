// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_provision is a model for public.metadata_provision table. This table store metadata provision information.
package metadata_provision

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// update metadata provision
//  Parameters:
//		uid
//			user id
//		metadataserviceID
//			metadata service name
//		metadataserviceName
//			metadata provision service name
//  Return:
//		metadata provision id
func PutMetadataPrivision(uid, metadataserviceID int64, metadataserviceName string) (interface{}, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	
	q := sqlUpdateMetadataMethod
	p := []interface{}{uid, metadataserviceID, metadataserviceName}

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	return metadataserviceName, nil
}