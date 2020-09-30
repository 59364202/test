// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agent is a model for api.agent table. This table store agent information.
package agent

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

// Set secret_key to null table api.agent
//  Parameters:
//		uid								
//					user id for tag update
//		id								
//					agent id for deleted key
//  Return:
//		return result and error
func DeleteKey(uid int64, id int64) (*result.Result, error) {

	// connect database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.Repack(err)
	}
	// sql deleted secret_key to table api.agent
	q := sqlDeleteKey
	p := []interface{}{uid, id}

	// process update secret_key
	_, err = db.Query(q, p...)
	if err != nil {
		// return error
		return nil, err
	}

	// return error = null
	return result.Result1(err), nil
}

// soft deleted table api.agent params agent.id
//  Parameters:
//		id			agent id for soft deleted agent user
//  Return:
//		return result and error
func DeleteKeyAccess(id int64) (*result.Result, error) {

	// connect database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.Repack(err)
	}
	
	p := []interface{}{id}
	// check user
	q := sqlDeleteLookupUser
	var user_id int64
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user_id)
	}

	p = []interface{}{id}
	// soft deleted agent
	q = sqlDeleteAgent
	stmt, _ := db.Prepare(q)
	res, err := stmt.Exec(p...)
	p = []interface{}{user_id}

	// soft deleted user
	q = sqlDeleteUser
	stmt, _ = db.Prepare(q)
	res, err = stmt.Exec(p...)
	res.RowsAffected()
	return result.Result1(nil), nil
}
