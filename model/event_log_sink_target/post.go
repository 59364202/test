// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_target is a model for api.event_log_sink_target table. This table store target send email information.
package event_log_sink_target

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// new event log sink target
//  Parameters:
//		uid
//			user id for updated_by
//		condition
//			sink condition id
//		method
//			sink method id 
//		group
//			array permission group id
//		lang
//			lang for send email template
//		color
//			color of group
//  Return:
//		sink target id
func AddTarget(uid, condition, method int64, group []int64, lang, color string) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := insTarget
	var targetID int64

	for _, v := range group {
		p := []interface{}{condition, method, lang, v, uid}
		err = db.QueryRow(q, p...).Scan(&targetID)
		if err != nil {
			return 0, pqx.GetRESTError(err)
		}
		makeDetailsSystemSetting(color, uid, v)
	}

	return targetID, nil
}

// make color target
//  Parameters:
//		color
//			color code
//		uid
//			user id update color group
//		gid
//			permission group id
//  Return:
//		None
func makeDetailsSystemSetting(color string, uid, gid int64) {
	name := "bof.EventMgt.PermissionGroup.Color_" + strconv.FormatInt(gid, 10)
	desc := "Color Setting PermissionGroup ID " + strconv.FormatInt(gid, 10)
	addColor(name, color, desc, uid)
}

// up color of sink target to system_setting
//  Parameters:
//		name
//			name system_setting for set color
//		color
//			color code
//		uid
//			user id update color group
//		gid
//			permission group id
//  Return:
//		None
func addColor(name, color, desc string, uid int64) (int64, error) {

	db, err := pqx.Open()
	if err != nil {
		return 0, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	q := upsertColor
	p := []interface{}{1, name, true, color, desc, uid}

	var settingID int64
	err = db.QueryRow(q, p...).Scan(&settingID)
	if err != nil {
		return 0, pqx.GetRESTError(err)
	}
	return settingID, nil
}
