package event_code

import (
	model_event_log_category "haii.or.th/api/thaiwater30/model/event_log_category"
	//result "haii.or.th/api/thaiwater30/util/result"
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
	//"log"
)

//	Get Event code list
//  Parameters:
//		param
//			Struct_EventCode_InputParam
//  Return:
//		[]Struct_EventCode
func GetEventCode(param *Struct_EventCode_InputParam) ([]*Struct_EventCode, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data         []*Struct_EventCode
		objEventCode *Struct_EventCode

		_id                sql.NullInt64
		_code              sql.NullString
		_description       sql.NullString
		_is_autoclose      sql.NullBool
		_troubleshoot      sql.NullString
		_subtype_category  sql.NullString
		_event_id          sql.NullInt64
		_event_code        sql.NullString
		_event_description sql.NullString
		//_channel_id			sql.NullInt64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	sqlCmdWhere := ""
	var arrParam = make([]interface{}, 0)
	arrEventCodeId := []string{}
	arrEventLogCategoryId := []string{}

	if param.ID != "" {
		arrEventCodeId = strings.Split(param.ID, ",")
	}
	if param.EventCategoryID != "" {
		arrEventLogCategoryId = strings.Split(param.EventCategoryID, ",")
	}

	//Check Filter id
	if len(arrEventCodeId) > 0 {
		if len(arrEventCodeId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.ID, " "))
			sqlCmdWhere += " AND subevent.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrEventCodeId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND subevent.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter event_code_category_id
	if len(arrEventLogCategoryId) > 0 {
		if len(arrEventLogCategoryId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.EventCategoryID, " "))
			sqlCmdWhere += " AND subevent.event_log_category_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrEventLogCategoryId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND subevent.event_log_category_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Query
	//log.Printf(sqlGetEventCode + sqlCmdWhere + sqlGetEventCodeOrderby, arrParam...)
	_result, err = db.Query(sqlGetEventCode+sqlCmdWhere+sqlGetEventCodeOrderby, arrParam...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_EventCode, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_code, &_description, &_is_autoclose, &_troubleshoot, &_subtype_category, &_event_id, &_event_code, &_event_description)

		if err != nil {
			return nil, err
		}

		//Generate EventCode object
		if _description.String == "" {
			_description.String = "{}"
		}

		if _event_description.String == "" {
			_event_description.String = "{}"
		}

		objEventCode = &Struct_EventCode{}
		objEventCode.ID = _id.Int64
		objEventCode.Code = _code.String
		objEventCode.IsAutoclose = _is_autoclose.Bool
		objEventCode.Troubleshoot = _troubleshoot.String
		objEventCode.SubtypeCategory = _subtype_category.String
		objEventCode.Description = json.RawMessage(_description.String)

		objEventCode.EventCategory = &model_event_log_category.Struct_EventLogCategory{}
		objEventCode.EventCategory.ID = _event_id.Int64
		objEventCode.EventCategory.Code = _event_code.String
		objEventCode.EventCategory.Description = json.RawMessage(_event_description.String)

		data = append(data, objEventCode)
	}

	//return result.Result1(data), nil
	return data, nil
}

//	Get Event category list wtih event code
//  Return:
//		[]Struct_Event
func GetEventCategoryEventCode() ([]*Struct_Event, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	rs := make([]*Struct_Event, 0)
	mapCategory := map[int64]*Struct_Event{}

	rows, err := db.Query(sqlEventCategoryEventCode)
	if err != nil {
		return nil, err
	}

	// นำข้อมูลที่ได้จาก db เก็บลงใน map
	for rows.Next() {
		var (
			_category_id       sql.NullInt64
			_category_code     sql.NullString
			_event_id          sql.NullInt64
			_event_description pqx.JSONRaw

			c  *Struct_Event
			ok bool
		)
		err = rows.Scan(&_category_id, &_category_code, &_event_id, &_event_description)
		if err != nil {
			return nil, err
		}

		if c, ok = mapCategory[_category_id.Int64]; !ok {
			c = &Struct_Event{
				Id:       _category_id.Int64,
				Code:     _category_code.String,
				Subevent: make([]*Struct_Subevent, 0),
			}
		}
		c.Subevent = append(c.Subevent, &Struct_Subevent{Id: _event_id.Int64, Description: _event_description.JSON()})
		mapCategory[_category_id.Int64] = c
	}

	// นำข้อมูลจาก map เก็บลง array เพื่อที่จะ return เป็น json
	for _, v := range mapCategory {
		rs = append(rs, v)
	}

	return rs, nil
}
