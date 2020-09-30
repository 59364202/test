// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agent is a model for api.agent table. This table store agent information.
package agent

import (
	"database/sql"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

// update table api.agent
//  Parameters:
//		uid			
//				user id
//		id				
//				agent id 
//  Return:
//		Data for service create folder dataimport
func UpdateKey(uid, id int64) (*result.Result, error) {
	// random secret_key
	ranKey := GenarateKey()

	// sql for update secret_key
	q := sqlPutUpdateKey
	p := []interface{}{ranKey, uid, id}

	// connect database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.Repack(err)
	}

	// prepare statement
	statement, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	// execute sql
	_, err = statement.Exec(p...)
	if err != nil {
		return nil, err
	}

	// sql for get agent user info by id
	q = sqlGetAgentUserInfo
	p = []interface{}{id}
	// process query get user
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	var user_account string
	// get user account
	for rows.Next() {
		rows.Scan(&user_account)
	}
	// create json for change secret_key in machine converter
	data := NewCronAgent(user_account, ranKey)
	// return data
	return result.Result1(data), nil
}

// update api.agent
// params user_id , agent.id, callback url, reqyest_origin, agent_type, key_access, permission realm , user account
func EditKeyAccess(uid int64, id int64, callbackurl string, request_origin string, agent_type_id int64, key_access string, permission_realm_id int64, user_account string) (*result.Result, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	// get sql for update table agent and table user
	// add params for update sql
	qAgent := sqlPutAgent
	pAgent := []interface{}{agent_type_id, callbackurl, permission_realm_id, uid, id}

	var user_id int64

	// prepare statment agent
	statement, err := db.Prepare(qAgent)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	// update agent
	err = statement.QueryRow(pAgent...).Scan(&user_id)
	if err != nil {
		return nil, err
	}

	// create json for change agent account in machine converter
	data := NewCronAgent(user_account, key_access)

	// return data
	return result.Result1(data), nil
}

// lookup table for dropdown box when add or edit agent
func LookupKeyAccess(db *pqx.DB, agent_type_id, permission_realm_id int64) (*SelectOption, *SelectOption) {

	// lookup agent type
	q := sqlPutLookupAgentType
	p := []interface{}{agent_type_id}
	lt_agent_type, _ := lookupData(db, q, p)

	// lookup permission realm
	q = sqlPutLookupPermissionRealm
	p = []interface{}{permission_realm_id}
	lt_permission_realm, _ := lookupData(db, q, p)
	return lt_agent_type, lt_permission_realm
}

// func for lookup data and add data to array dropdown box
func lookupData(db *pqx.DB, q string, p []interface{}) (*SelectOption, error) {

	// query lookuptable
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var (
		name sql.NullString
		id   sql.NullInt64
	)
	// loop lookup table
	for rows.Next() {
		rows.Scan(&id, &name)
	}
	// add data to select option
	sOption := new(SelectOption)
	sOption.Value = id.Int64
	sOption.Text = name.String

	// return select option
	return sOption, nil
}
