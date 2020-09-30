// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agent is a model for api.agent table. This table store agent information.
package agent

import (
	"crypto/rand"
	"database/sql"
	result "haii.or.th/api/thaiwater30/util/result"
	selectOption "haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

// get data table api.agent
//  Parameters:
//		None
//  Return:
//		Agent Information
func GetKeyAccessTable() (interface{}, error) {
	// connect  database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.Repack(err)
	}

	// sql get data from api.agent
	q := sqlGetKeyAccessTable
	p := []interface{}{}

	// process get data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*KeyAccess, 0)
	// loop data for created json
	for rows.Next() {
		//declare value for get data from table
		var (
			id                  int64
			callback_url        sql.NullString
			agent_type          string
			secret_key          sql.NullString
			permission_realm    sql.NullString
			agent_type_id       int64
			permission_realm_id sql.NullInt64
			account             string
			full_name           sql.NullString
		)
		d := new(KeyAccess)
		rows.Scan(&id, &callback_url, &agent_type, &secret_key, &permission_realm, &agent_type_id, &permission_realm_id, &account, &full_name)
		d.ID = id
		d.UserAccount = account
		d.FullName = full_name.String
		sOptionAgent := new(SelectOption)
		sOptionAgent.Text = agent_type
		sOptionAgent.Value = agent_type_id
		d.AgentType = sOptionAgent
		d.CallbackURL = callback_url.String
		d.KeyAccess = secret_key.String
		sOptionPermissionRealm := new(SelectOption)
		sOptionPermissionRealm.Text = permission_realm.String
		sOptionPermissionRealm.Value = permission_realm_id.Int64
		d.PermissionRealm = sOptionPermissionRealm
		data = append(data, d)
	}

	// return  data
	return result.Result1(data), nil
}

// func random text
//  Parameters:
//		None
//  Return:
//		Secret Key Genarate
func GenarateKey() string {
	dictionary := "0123456789abcdef"
	var bytes = make([]byte, 64)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	ranKey := string(bytes)
	return ranKey
}

// get agent type
//  Parameters:
//		None
//  Return:
//		All Type of agent
func GetAgentType() (interface{}, error) {

	// connect database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.Repack(err)
	}

	// sql get agent type
	q := sqlGetAgentType
	p := []interface{}{}

	// process agent type
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var data = selectOption.NewSelect()
	// loop data
	for rows.Next() {
		//declare value for get data from table
		var (
			id   sql.NullInt64
			text sql.NullString
		)
		rows.Scan(&id, &text)
		data.Add(id.Int64, text.String)
	}
	// return data
	return result.Result1(data.Option), nil
}

// get permission realm
//  Parameters:
//		None
//  Return:
//		All Permission Realm Name and ID
func GetPermissionRealm() (*result.Result, error) {

	// connect database server
	db, err := pqx.Open()
	// check error
	if err != nil {
		return nil, errors.Repack(err)
	}

	// get sql permission realm
	q := sqlGetPermissionRealm
	p := []interface{}{}

	// process get data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	var data = selectOption.NewSelect()
	// loop for get data
	for rows.Next() {
		//declare value for get data from table
		var (
			id   sql.NullInt64
			text sql.NullString
		)
		rows.Scan(&id, &text)
		data.Add(id.Int64, text.String)
	}
	// return data
	return result.Result1(data.Option), nil
}
