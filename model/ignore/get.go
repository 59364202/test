package ignore

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_dam_hourly "haii.or.th/api/thaiwater30/model/dam_hourly"
	model_rainfall24h "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"

	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"strconv"
	"strings"
	"time"
	//	"log"
)

//	get Exist Data
//	Parameters:
//		db
//			DB connection
//		tableName
//			ชื่อตาราง
//		stationID
//			รหัสสถานี
//		dataID
//			รหัสข้อมูล
//		isIgnore
//			สถานะ ignore
//		userID
//			รหัสผู้ใช้
//	Return:
//		sql.Rows
//		sql query
func getIgnoreStationInfo(db *pqx.DB, tableName string, stationID string, dataID string, isIgnore bool, userID int64) (*sql.Rows, string, error) {
	var (
		sqlCmdQuery string = ""
		//remarks		string	= ""
		arrParam    []interface{}
		intRowCount int64 = 0
	)

	switch tableName {
	case "rainfall_24h":
		sqlCmdQuery = model_rainfall24h.SQLSelectIgnoreData
	case "tele_waterlevel":
		sqlCmdQuery = model_waterlevel.SQLSelectIgnoreData
	case "waterquality":
		sqlCmdQuery = model_waterquality.SQLSelectIgnoreData
	case "dam_daily":
		sqlCmdQuery = model_dam_daily.SQLSelectIgnoreData
	case "dam_hourly":
		sqlCmdQuery = model_dam_hourly.SQLSelectIgnoreData
	default:
		return nil, sqlCmdQuery, rest.NewError(422, "unknown table_name", errors.New("unknown table_name"))
	}

	arrParam = make([]interface{}, 0)
	// arrParam = append(arrParam, isIgnore)
	//arrParam = append(arrParam, userID)

	if stationID != "" {
		arrParam = append(arrParam, strings.Trim(stationID, " "))
		sqlCmdQuery += " AND d.id = $" + strconv.Itoa(len(arrParam))
	}

	if dataID != "" {
		arrParam = append(arrParam, strings.Trim(dataID, " "))
		sqlCmdQuery += " AND dd.id = $" + strconv.Itoa(len(arrParam))
	}

	/*if len(arrDataID) > 0 {
		if len(arrDataID) == 1 {
			arrParam = append(arrParam, strings.Trim(arrDataID[0], " "))
			sqlCmdQuery += " AND dd.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrDataID {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdQuery += " AND dd.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}*/

	sqlCmdQuery = strings.Replace(sqlCmdQuery, "#{UserID}", strconv.FormatInt(userID, 10), -1)
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err := db.Query(sqlCmdQuery, arrParam...)
	if err != nil {
		return nil, "", pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	for _result.Next() {
		intRowCount += 1
	}

	if intRowCount == 0 {
		return _result, "", nil
	}

	return _result, sqlCmdQuery, nil
}

//	get วันเวลาล่าสุดที่ ignore
//	Parameters:
//		jsonStationList
//			system setting ที่กำหนดว่าจะใช้ ตารางไหน
//	Return:
//		jsonStationList พร้อมวันเวลาล่าสุดที่ ignore
func GetLastestIgnoreStation(jsonStationList json.RawMessage) ([]map[string]string, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Variables
	var (
		data             []map[string]string
		objIgnoreStation map[string]string

		_data_category       sql.NullString
		_max_ignore_datetime time.Time

		_result *sql.Rows
	)

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Query
	//	log.Printf("SELECT data_category, MAX(ignore_datetime) AS max_ignore_datetime FROM ignore_history GROUP BY data_category")
	_result, err = db.Query("SELECT data_category, MAX(ignore_datetime) AS max_ignore_datetime FROM ignore_history GROUP BY data_category")
	if err != nil {
		//		log.Println(err)
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	objIgnoreStation = make(map[string]string, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_data_category, &_max_ignore_datetime)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}

		objIgnoreStation[_data_category.String] = _max_ignore_datetime.Format(strDatetimeFormat)
	}

	// a map container to decode the JSON structure into
	objStationList := make([]map[string]string, 0)

	// unmarschal JSON
	err = json.Unmarshal(jsonStationList, &objStationList)
	if err != nil {
		return nil, rest.NewError(500, err.Error(), err)
	}

	data = make([]map[string]string, 0)
	for _, mapValue := range objStationList {
		objData := make(map[string]string, 0)
		objData = mapValue
		objData["lastest_ignore_datetime"] = objIgnoreStation[mapValue["data_type"]]
		data = append(data, objData)
	}

	//Return data
	return data, nil
}
