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

// new metadata provision
//  Parameters:
//		uid
//			user id
//		metadataserviceName
//			metadata provision service name
//  Return:
//		metadata provision id
func PostMetadataProvision(uid int64, metadataserviceName string) (interface{}, error) {
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := sqlInsMetadataProvision
	p := []interface{}{metadataserviceName, uid}

	var methodID int64
	err = db.QueryRow(q, p...).Scan(&methodID)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	return methodID, nil
}
