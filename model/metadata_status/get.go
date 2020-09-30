// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_status is a model for public.lt_metadata_status table. This table store metadata status information.
package metadata_status

import (
	"database/sql"
	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// get metadata status list
//  Parameters:
//		None
//  Return:
//		Array information cctv
func GetMetadataStatus() ([]*MetadataStatusParams, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	
	// get sql
	q := getMetadataMethod
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	// make array data
	data := make([]*MetadataStatusParams, 0)
	for rows.Next() {
		dataRow := &MetadataStatusParams{}
		var (
			id   sql.NullInt64
			name pqx.JSONRaw
		)
		rows.Scan(&id, &name)
		sid := strconv.FormatInt(id.Int64, 10)
		dataRow.ID = id.Int64
		dataRow.StatusID, _ = model.GetCipher().EncryptText(sid)
		dataRow.StatusName = name.JSON()
		if id.Int64 != 1 {
			dataRow.IsDeleted = true
		}
		// add data to array
		data = append(data, dataRow)
	}
	
	// return data
	return data, nil
}
