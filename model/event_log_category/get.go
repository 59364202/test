package event_log_category

import (
	//result "haii.or.th/api/thaiwater30/util/result"
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strings"
)

//	Get Event log Category list
//  Parameters:
//		param
//			Struct_EventLogCategory_InputParam
//  Return:
//		[]Struct_ELC
func GetEventLogCategory(param *Struct_EventLogCategory_InputParam) ([]*Struct_ELC, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data                []*Struct_ELC
		objEventLogCategory *Struct_ELC

		_id          sql.NullInt64
		_code        sql.NullString
		_color       sql.NullString
		_description sql.NullString

		_result *sql.Rows
	)

	sqlCmd := ""
	if param.ID != "" {
		if len(strings.Split(param.ID, ",")) == 0 {
			sqlCmd = " AND elc.id = " + param.ID
		} else {
			sqlCmd = " AND elc.id IN (" + param.ID + ")"
		}
	}

	//Query
	//log.Printf(sqlGetEventLogCategory + sqlCmd + sqlGetEventLogCategoryOrderby)
	_result, err = db.Query(sqlGetEventLogCategory + sqlCmd + sqlGetEventLogCategoryOrderby)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_ELC, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_code, &_description, &_color)

		if err != nil {
			return nil, err
		}

		//Generate EventLogCategory object
		objEventLogCategory = &Struct_ELC{}
		objEventLogCategory.ID = _id.Int64
		objEventLogCategory.Code = _code.String
		objEventLogCategory.Color = _color.String

		if _description.String == "" {
			_description.String = "{}"
		}
		objEventLogCategory.Description = json.RawMessage(_description.String)

		data = append(data, objEventLogCategory)
	}

	//Return result
	//return result.Result1(data), nil
	return data, nil
}
