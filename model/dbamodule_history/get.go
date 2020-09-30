package dbamodule_history

import (
	model_setting "haii.or.th/api/server/model/setting"
	//result "haii.or.th/api/thaiwater30/util/result"
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//"strconv"
	//	"log"
	"time"
)

//	get create partition table history
//	Parameters:
//		param
//			ใช้ในส่วน TableName เพื่อดูประวัติของตารางนั้น
//			ใช้ในส่วน Year เพื่อดูประวัติของปีนั้น
//	Return:
//		[]Struct_DBAModuleHistory
func GetDBAModuleHistory(param *Struct_DBAModuleHistory_InputParam) ([]*Struct_DBAModuleHistory, error) {

	//Check input params
	if param.TableName == "" {
		return nil, errors.Repack(errors.New("'table_name' is not null."))
	}
	if param.Year == "" {
		return nil, errors.Repack(errors.New("'year' is not null."))
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	var (
		data                []*Struct_DBAModuleHistory
		objDBAModuleHistory *Struct_DBAModuleHistory

		_id            sql.NullInt64
		_datetime      time.Time
		_month         sql.NullString
		_remark        sql.NullString
		_user_id       sql.NullInt64
		_user_fullname sql.NullString

		_result *sql.Rows
	)

	//Open DB
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Query SQL command
	//	log.Println(sqlSelectHistory, param.TableName, param.Year)
	_result, err = db.Query(sqlSelectHistory, param.TableName, param.Year)

	data = make([]*Struct_DBAModuleHistory, 0)

	//Loop result
	for _result.Next() {
		_result.Scan(&_id, &_datetime, &_month, &_remark, &_user_id, &_user_fullname)

		objDBAModuleHistory = &Struct_DBAModuleHistory{}
		objDBAModuleHistory.ID = _id.Int64
		objDBAModuleHistory.Datetime = _datetime.Format(strDatetimeFormat)
		objDBAModuleHistory.Month = _month.String
		objDBAModuleHistory.Remark = _remark.String
		objDBAModuleHistory.TableName = param.TableName
		objDBAModuleHistory.Year = param.Year

		objDBAModuleHistory.User = &Struct_User{}
		objDBAModuleHistory.User.ID = _user_id.Int64
		objDBAModuleHistory.User.FullName = _user_fullname.String

		data = append(data, objDBAModuleHistory)
	}

	//Return Data
	return data, nil
}
