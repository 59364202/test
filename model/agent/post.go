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
	"strings"
)

// New agency to table api.agent
//  Parameters:
//		user_id							
//					user id for tag update.
//		user_account								
//					agent account.
//		callbackurl									
//					callback url for type www.
//		request_origin								
//					origin header for type www.
//		agent_type_id								
//					agent type id.
//		permission_realm							
//					permission realm agent.
//  Return:
//		return result and error
func AddKeyAccess(uid int64, user_account string, callbackurl string, request_origin string, agent_type_id int64, key_access string, permission_realm_id int64) (*result.Result, error) {

	// connect database server
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}
	au := strings.Split(user_account, " ")
	user_account = strings.Join(au, "")

	// sql add user
	qUser := sqlPostUser
	pUser := []interface{}{user_account}

	// sql add agent user
	qAgent := sqlPostAgent
	keyAgent := GenarateKey()
	pAgent := []interface{}{agent_type_id, callbackurl, keyAgent, permission_realm_id, uid}

	// func inser user agent
	_, err = insertUserAgent(db, qUser, pUser, qAgent, pAgent, agent_type_id)

	if err != nil {
		return nil, err
	}

	data := NewCronAgent(user_account, keyAgent)

	//	agentType, permissionRealm := LookupKeyAccess(db, agent_type_id, permission_realm_id)
	//	data := &KeyAccess{ID: id, UserAccount: user_account, AgentType: agentType, CallbackURL: callbackurl, RequestOrigin: request_origin, KeyAccess: key_access, PermissionRealm: permissionRealm}
	return result.Result1(data), nil
}

// inser user agent
//  Parameters:
//		db							
//					db.open()
//		qUser								
//					sql regist user agent
//		pUser								
//					information regist user
//		qAgent								
//					sql regist agent
//		pAgent								
//					information regist agent
//		agent_type_id   					
//					type check for get permission
//  Return:
//		agent id
func insertUserAgent(db *pqx.DB, qUser string, pUser []interface{}, qAgent string, pAgent []interface{}, agent_type_id int64) (int64, error) {

	var id int64
	var user_id int64
	var permission_id int64

	// prepare statement agent and user
	statement, err := db.Prepare(qAgent)
	statement1, err := db.Prepare(qUser)
	qGroup := sqlPostPermissionGroup
	statement2, err := db.Prepare(qGroup)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	// insert user for get new user id
	err = statement1.QueryRow(pUser...).Scan(&user_id)
	if err != nil {
		return 0, err
	}

	// insert new agent by new user id
	pAgent = append(pAgent, user_id)
	err = statement.QueryRow(pAgent...).Scan(&id)
	if err != nil {
		return 0, err
	}
	// add permission new agent converter
	pPermission := []interface{}{user_id}
	err = statement2.QueryRow(pPermission...).Scan(&permission_id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// created json for service created dataimport
//  Parameters:
//		keyAgent								
//					Secret key agent
//		user_account							
//					Agent account 
//  Return:
//		Data for service create folder dataimport
func NewCronAgent(user_account, keyAgent string) *NewAgent {
	data := &NewAgent{}
	agency := strings.Split(user_account, "-")
	if len(agency) < 2 {
		return nil
	}
	data.Agency = agency[len(agency)-1]
	data.AgentName = user_account
	data.AgentKey = keyAgent

	return data
}
