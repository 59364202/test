// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata_method is a model for public.metadata_method table. This table store metadata_method information.
package metadata_method

import (
	"database/sql"
	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// get metadata method list
//  Parameters:
//		None
//  Return:
//		Array MetadataMethodParams
func GetMetadataMethod() ([]*MetadataMethodParams, error) {
	// open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql get metadata method
	q := getMetadataMethod
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := make([]*MetadataMethodParams, 0)
	for rows.Next() {
		dataRow := &MetadataMethodParams{}
		var (
			id   sql.NullInt64
			name sql.NullString
		)
		// scan data
		rows.Scan(&id, &name)
		sid := strconv.FormatInt(id.Int64, 10)
		dataRow.ID = id.Int64
		dataRow.MethodID, _ = model.GetCipher().EncryptText(sid)
		dataRow.MethodName = name.String
		data = append(data, dataRow)
	}

	return data, nil
}
