// Copyright 2017 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package validdata provides a function to check valid data and return data.
package validdata

import (
	"database/sql/driver"
	"haii.or.th/api/util/datatype"
)

// ValidData return valid value.
//
//	Parameters:
//		valid value type boolean
//		value type interface
// 	Return:
//		format datetime
//
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

//	convert sql.NullInt64, sql.NullString, sql.NullFloat64 to string
//	Parameters:
//		v, err from sql.NullInt64.Value()
//	Return:
//		string v
func DataString(v driver.Value, err error) string {
	return datatype.MakeString(v)
}
