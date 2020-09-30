// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package partition is a model for manage partition table
package partition

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	get list partition table
//	Return:
//		Array string
func ListPartitionTable() ([]string, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// Get all table from public schema with pt_table_name_field_name and without _yddddmdd
	r, err := pqx.ListPartionableTable(db, "public")
	if err == nil {
		return r, nil
	}
	return nil, errors.Repack(err)
}
