// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package event_log_sink_method is a model for api.event_log_sink_method table. This table store method config e-mail infomation.
package event_log_sink_method

import (
	model_method_type "haii.or.th/api/thaiwater30/model/event_log_sink_method_type"
	//result "haii.or.th/api/thaiwater30/util/result"
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

// get event log sink method information
//  Parameters:
//		param
//			Struct_EventLogSinkMethod_InputParam
//  Return:
//		Struct_EventLogSinkMethod
func GetEventLogSinkMethod(param *Struct_EventLogSinkMethod_InputParam) ([]*Struct_EventLogSinkMethod, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data                  []*Struct_EventLogSinkMethod
		objEventLogSinkMethod *Struct_EventLogSinkMethod

		_id                      sql.NullInt64
		_description             sql.NullString
		_sink_params             sql.NullString
		_method_type_id          sql.NullInt64
		_method_type_description sql.NullString

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	sqlCmdWhere := ""
	arrParam := make([]interface{}, 0)
	arrId := []string{}
	arrTypeId := []string{}

	if param.Id != "" {
		arrId = strings.Split(param.Id, ",")
	}

	if param.Event_log_sink_method_type_id != "" {
		arrTypeId = strings.Split(param.Event_log_sink_method_type_id, ",")
	}

	//Check Filter id
	if len(arrId) > 0 {
		if len(arrId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Id, " "))
			sqlCmdWhere += " AND sm.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND sm.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter event_log_sink_method_type_id
	if len(arrTypeId) > 0 {
		if len(arrTypeId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.Event_log_sink_method_type_id, " "))
			sqlCmdWhere += " AND sm.event_log_sink_method_type_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrTypeId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND sm.event_log_sink_method_type_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Query
	//log.Printf(sqlGetEventLogSinkMethod + sqlCmdWhere + sqlGetEventLogSinkMethodOrderby, arrParam...)
	_result, err = db.Query(sqlGetEventLogSinkMethod+sqlCmdWhere+sqlGetEventLogSinkMethodOrderby, arrParam...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_EventLogSinkMethod, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_description, &_sink_params, &_method_type_id, &_method_type_description)
		if err != nil {
			return nil, err
		}

		if _sink_params.String == "" {
			_sink_params.String = "{}"
		}

		nSetting := map[string]interface{}{}
		dataSetting := []byte(_sink_params.String)
		err = json.Unmarshal(dataSetting, &nSetting)
		objEventLogSinkMethod = &Struct_EventLogSinkMethod{}
		if err != nil {
		} else {
			if nSetting["system_setting_name"] != nil {
				b := nSetting["system_setting_name"].(string)
				cfname := strings.Split(b, ".")
				objEventLogSinkMethod.ConfigName = cfname[len(cfname)-1]
			}
		}
		//Generate objEventLogSinkMethod object
		objEventLogSinkMethod.Id = _id.Int64
		objEventLogSinkMethod.Description = _description.String
		objEventLogSinkMethod.Sink_params = json.RawMessage(_sink_params.String)

		objEventLogSinkMethod.Event_log_sink_method_type = &model_method_type.Struct_EventLogSinkMethodType{}
		objEventLogSinkMethod.Event_log_sink_method_type.Id = _method_type_id.Int64
		objEventLogSinkMethod.Event_log_sink_method_type.Description = _method_type_description.String

		data = append(data, objEventLogSinkMethod)
	}

	//Return result
	return data, nil
}
