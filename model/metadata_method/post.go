// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_method is a model for public.metadata_method table. This table store metadata_method information.
package metadata_method

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// new metadata method
//  Parameters:
//		uid
//			user id
//		metadata_name
//			metadata method name
//  Return:
//		metadata method id
func PostMetadataMethod(uid int64, metadata_name string) (interface{}, error) {
	
	// open database
	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql insert metadata method
	q := insMetadataMethod
	p := []interface{}{metadata_name, uid}

	var methodID int64
	// insert data and get id
	err = db.QueryRow(q, p...).Scan(&methodID)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	return methodID, nil
}
