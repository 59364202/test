package ignore_history

import (
	"database/sql"
	"encoding/json"
	//	"log"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	"haii.or.th/api/thaiwater30/util/validdata"

	model_setting "haii.or.th/api/server/model/setting"
)

//	get ignore history
//	Parameters:
//		param
//			ใช้ในส่วน table_name
//	Return:
//		[]Struct_IgnoreHistory
func GetIgnoreHistory(param *Struct_IgnoreHistory_InputParam) ([]*Struct_IgnoreHistory, error) {
	//Check input params
	if param.TableName == "" {
		return nil, errors.Repack(errors.New("'table_name' is not null."))
	}

	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere = ""

	if param.TableName == "waterlevel" {
		arrSqlCmd := []string{}

		arrParam = append(arrParam, "tele_waterlevel")
		arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))

		arrParam = append(arrParam, "canal_waterlevel")
		arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))

		sqlCmdWhere += " WHERE h.data_category IN (" + strings.Join(arrSqlCmd, ",") + ")"
	} else {
		arrParam = append(arrParam, param.TableName)
		sqlCmdWhere += " WHERE h.data_category = $" + strconv.Itoa(len(arrParam))
	}

	return getIgnoreHistory(param, sqlSelectHistory+sqlCmdWhere+sqlSelectHistory_OrderBy, arrParam)
}

//	get ignore data
//	Parameters:
//		param
//			ใช้ในส่วน table_name
//	Return:
//		[]Struct_IgnoreHistory
func GetIgnoreData(param *Struct_IgnoreHistory_InputParam) ([]*Struct_IgnoreHistory, error) {
	//Check input params
	if param.TableName == "" {
		return nil, errors.Repack(errors.New("'table_name' is not null."))
	}

	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere = ""

	if param.TableName == "waterlevel" {
		arrSqlCmd := []string{}

		arrParam = append(arrParam, "tele_waterlevel")
		arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))

		arrParam = append(arrParam, "canal_waterlevel")
		arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))

		sqlCmdWhere += " AND ih.data_category IN (" + strings.Join(arrSqlCmd, ",") + ")"
	} else {
		arrParam = append(arrParam, param.TableName)
		sqlCmdWhere += " AND ih.data_category = $" + strconv.Itoa(len(arrParam))
	}

	return getIgnoreHistory(param, sqlSelectIgnoreData+sqlCmdWhere, arrParam)
}

//	get ignore history
//	Parameters:
//		param
//			ใช้ในส่วน table_name
//		sqlCmdQuery
//			sql query
//		arrParam
//			parameter ที่ใช้ร่วมกับ query
//	Return:
//		[]Struct_IgnoreHistory
func getIgnoreHistory(param *Struct_IgnoreHistory_InputParam, sqlCmdQuery string, arrParam []interface{}) ([]*Struct_IgnoreHistory, error) {

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Find date default format
	strDateFormat := model_setting.GetSystemSetting("bof.Default.DateFormat")
	if strDateFormat == "" {
		strDateFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}

	var (
		data       []*Struct_IgnoreHistory
		objHistory *Struct_IgnoreHistory

		_id               sql.NullInt64
		_ignore_datetime  time.Time
		_data_category    sql.NullString
		_station_id       sql.NullInt64
		_station_oldcode  sql.NullString
		_station_name     sql.NullString
		_station_province sql.NullString
		_agency_name      sql.NullString
		_agency_shortname sql.NullString
		_data_id          sql.NullInt64
		_data_datetime    time.Time
		_remark           sql.NullString
		_user_id          sql.NullInt64
		_user_fullname    sql.NullString
		_data_value       sql.NullFloat64

		_result *sql.Rows
	)

	//Open DB
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Query SQL command
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err = db.Query(sqlCmdQuery, arrParam...)

	//Loop result
	data = make([]*Struct_IgnoreHistory, 0)
	for _result.Next() {
		_result.Scan(&_id, &_ignore_datetime, &_data_category,
			&_station_id, &_station_oldcode, &_station_name, &_station_province,
			&_agency_shortname, &_agency_name,
			&_data_id, &_data_datetime, &_remark,
			&_user_id, &_user_fullname, &_data_value)

		if !_station_name.Valid || _station_name.String == "" {
			_station_name.String = "{}"
		}
		if !_station_province.Valid || _station_province.String == "" {
			_station_province.String = "{}"
		}
		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		//if !_agency_name.Valid || _agency_name.String == "" {
		//	_agency_name.String = "{}"
		//}

		objHistory = &Struct_IgnoreHistory{}
		objHistory.ID = _id.Int64
		objHistory.IgnoreDatetime = _ignore_datetime.Format(strDatetimeFormat)
		objHistory.DataCategory = _data_category.String
		objHistory.StationID = _station_id.Int64
		objHistory.StationOldCode = _station_oldcode.String
		objHistory.StationName = json.RawMessage(_station_name.String)
		objHistory.StationProvince = json.RawMessage(_station_province.String)
		//objHistory.AgencyName = json.RawMessage(_agency_name.String)
		objHistory.AgencyShortname = json.RawMessage(_agency_shortname.String)
		objHistory.DataID = _data_id.Int64

		//log.Println(_data_category.String)

		if _data_category.String == "dam_daily" {
			objHistory.DataDate = _data_datetime.Format(strDateFormat)
		} else {
			objHistory.DataDate = _data_datetime.Format(strDatetimeFormat)
		}

		//log.Println(strDateFormat)
		//log.Println(strDatetimeFormat)
		//log.Println(_data_datetime)
		//log.Println(objHistory.DataDate)

		objHistory.Remark = _remark.String
		objHistory.DataValue = validdata.ValidData(_data_value.Valid, _data_value.Float64)

		objHistory.User = &Struct_User{}
		objHistory.User.ID = _user_id.Int64
		objHistory.User.FullName = _user_fullname.String

		data = append(data, objHistory)
	}

	//Return Data
	return data, nil
}
