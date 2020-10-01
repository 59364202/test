package ignore

import (
	model_history "haii.or.th/api/thaiwater30/model/ignore_history"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	//	"log"
	"strings"
)

//	update สถานะ ignore
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		param
//			*Struct_IgnoreStation_InputParam
//	Return:
//		Ignore Data Successful
func PatchIgnoreStation(userId int64, param *Struct_IgnoreStation_InputParam) (string, error) {

	//Check input params
	if param.TableName == "" {
		return "", rest.NewError(422, "'table_name' is not null.", errors.New("parameter 'table_name' is not null."))
	}
	if param.ID == "" {
		return "", rest.NewError(422, "'data_id' is not null.", errors.New("parameter 'data_id' is not null."))
	}
	if param.StationID == "" {
		return "", rest.NewError(422, "'station_id' is not null.", errors.New("parameter 'station_id' is not null."))
	}

	//arrDataID := strings.Split(param.ID, ",")

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", pqx.GetRESTError(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Check Exist Partition Table of param.Table in param.Year
	_, sqlCmdIgnoreInfo, err := getIgnoreStationInfo(db, param.TableName, param.StationID, param.ID, param.IsIgnore, userId)
	if err != nil {
		return "", err
	}
	if sqlCmdIgnoreInfo == "" {
		return "", rest.NewError(422, "ไม่พบข้อมูลที่ต้องการดำเนินการ", errors.New("Can't find data from filter variables."))
	}

	//Update ignore station
	err = updateIgnoreStation(tx, param.TableName, param.StationID, param.ID, param.IsIgnore, userId, sqlCmdIgnoreInfo)
	if err != nil {
		return "", err
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	return "Ignore Data Successful", nil
}

//	Insert to dbamodule_history table
//	Parameters:
//		tx
//			Transaction
//		tableName
//			ชื่อตาราง
//		stationID
//			รหัสสถานี
//		dataID
//			รหัสข้อมูล
//		isIgnore
//			สถานะ ignore
//		userId
//			รหัสผู้ใช้
//		sqlCmdCheckIgnoreInfo
//			sql query ที่ต้องใช้
//	Return:
//		nil, error
func updateIgnoreStation(tx *pqx.Tx, tableName string, stationID string, dataID string, isIgnore bool, userId int64, sqlCmdCheckIgnoreInfo string) error {

	// var _id int64
	var metadataTableName string = ""

	switch tableName {
	case "rainfall_24h":
		metadataTableName = "m_tele_station"
	case "dam_daily":
		metadataTableName = "m_dam"
	case "dam_hourly":
		metadataTableName = "m_dam"
	case "tele_waterlevel":
		metadataTableName = "m_tele_station"
	case "waterquality":
		metadataTableName = "m_waterquality_station"
	default:
		return rest.NewError(404, "Unknown table_name", nil)
	}

	//======= update ignore station =======//
	// Prepare Statement
	//	log.Printf(" UPDATE public."+metadataTableName+" SET is_ignore = $2 WHERE id = $1 ; ", stationID, isIgnore)
	statement, err := tx.Prepare(" UPDATE public." + metadataTableName + " SET is_ignore = $2 WHERE id = $1 ; ")
	if err != nil {
		return pqx.GetRESTError(err)
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(stationID, isIgnore)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	//======= Insert ignore_history table =======//
	var remarks string
	if isIgnore {
		remarks = "Ignore station"
	} else {
		remarks = "Unlock Ignore"
	}

	//paramHistory := &model_history.Struct_IgnoreHistory_InputParam{}
	//paramHistory.SQLCmdIgnoreInfo = strings.Replace(sqlCmdCheckIgnoreInfo, "#{Remarks}", remarks, -1)

	//Execute insert statement with parameters and returning id
	var arrParam = make([]interface{}, 0)
	//arrParam = append(arrParam, isIgnore)
	arrParam = append(arrParam, stationID)
	arrParam = append(arrParam, dataID)
	newID, err := model_history.PostIgnoreHistoryWithSqlCmd(strings.Replace(sqlCmdCheckIgnoreInfo, "#{Remarks}", remarks, -1), arrParam)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	upsertSQL := `
	INSERT INTO
		public.ignore(
			ignore_datetime,
			data_category,
			station_id,
			station_oldcode,
			station_name,
			station_province,
			agency_shortname,
			data_datetime,
			data_value,
			data_id,
			remark,
			is_ignore,
			created_by
		)
	SELECT
		ignore_datetime,
		data_category,
		station_id :: bigint,
		station_oldcode,
		station_name,
		station_province,
		agency_shortname,
		data_datetime,
		data_value,
		data_id,
		remark,
		$2,
		created_by
	FROM
		public.ignore_history
	WHERE
		id = $1 ON CONFLICT (data_category, station_id) 
	DO UPDATE SET
		ignore_datetime = excluded.ignore_datetime,
		data_category = excluded.data_category,
		station_id = excluded.station_id,
		station_oldcode = excluded.station_oldcode,
		station_name = excluded.station_name,
		station_province = excluded.station_province,
		agency_shortname = excluded.agency_shortname,
		data_datetime = excluded.data_datetime,
		data_value = excluded.data_value,
		data_id = excluded.data_id,
		remark = excluded.remark,
		is_ignore = excluded.is_ignore,
		updated_by = excluded.created_by
	`
	upsertParam := []interface{}{newID, isIgnore}
	db, err := pqx.Open()
	if err != nil {
		return pqx.GetRESTError(err)
	}
	_, err = db.Exec(upsertSQL, upsertParam...) // upsert to ignore
	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}
