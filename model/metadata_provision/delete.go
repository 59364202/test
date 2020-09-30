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

// soft deleted metadata provision
//  Parameters:
//		uid
//			user id
//		methodProvisionID
//			metadata provision id
//  Return:
//		metadata provision id
func DelMetadataProvision(uid int64, methodProvisionID int64) (interface{}, error) {

	// open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql deleted metadata method
	q := sqlDelMetadataMethod
	p := []interface{}{uid, methodProvisionID}
	
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

	return methodProvisionID, nil
}
