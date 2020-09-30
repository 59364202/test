// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dashboard is a model for api.dataimport_download,api.dataimport_dataset,login_session table. This table store dataimport logs information.
package dashboard

import (
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// update metadata offline date
//  Parameters:
//		metadataID
//				metadata id for update
//  Return:
//		metadata id
func UpdateMetadataOfflineDate(metadataID int64) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := sqlUpdateMetadataOfflineDate
	p := []interface{}{time.Now().Format(time.RFC3339), metadataID}

	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	res, err := stmt.Exec(p...)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return 0, pqx.GetRESTError(err)
	}

	return metadataID, nil
}
