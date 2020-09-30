// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_provision is a model for public.metadata_provision table. This table store metadata provision information.
package metadata_provision

import (
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// get metadata provision by id
//  Parameters:
//		metadataID
//			metadata provision id
//  Return:
//		Array MetadataProvisionOutput
func GetMetadataProvision(metadataID string) (interface{}, error) {
	// connect database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlSelectMetadataProvision
	p := []interface{}{}
	if metadataID != "" {
		q += " AND md.id=$1"
		p = append(p, metadataID)
	}

	// query data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := make([]*MetadataProvisionOutput, 0)
	// loop add data to array
	for rows.Next() {
		dataRow := &MetadataProvisionOutput{}
		var (
			id           sql.NullInt64
			name         pqx.JSONRaw
			downloadName sql.NullString
			datasetName  sql.NullString
		)
		rows.Scan(&id, &name, &downloadName, &datasetName)

		dataRow.ID = id.Int64
		dataRow.Name = name.JSON()
		dataRow.DownloadName = downloadName.String
		dataRow.DatasetName = datasetName.String
		data = append(data, dataRow)
	}

	return data, nil
}
