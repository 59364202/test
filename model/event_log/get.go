package event_log

import (
	"database/sql"
	"encoding/json"
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_event_code "haii.or.th/api/thaiwater30/model/event_code"
	model_event_category "haii.or.th/api/thaiwater30/model/event_log_category"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//	"log"
	"strconv"
	"strings"
	"time"
)

//Description: Get EventLogSummary (Group by Agency and EventCategory)
//Input Parameter: start_date*, end_date*, agency_id(multi)
//  Parameters:
//		param
//			Struct_EventLog_InputParam
//  Return:
//		EventLogSummary (Group by Agency and EventCategory)
func GetEventLogSummaryGroupByAgencyCategory(param *Struct_EventLog_InputParam) ([]*Struct_EventLogSummary_GroupByAgencyCategory, error) {
	//Check Input Parameters
	strStartDate, strEndDate, err := checkInputParam(param)
	if err != nil {
		return nil, err
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data               []*Struct_EventLogSummary_GroupByAgencyCategory
		objEventLogSummary *Struct_EventLogSummary_GroupByAgencyCategory

		intAgencyID int64 = 0
		//arrEventLogColumn 	map[string]int64
		dataEventLogCategory []*model_event_category.Struct_EventLogCategory
		objEventLogCategory  *model_event_category.Struct_EventLogCategory

		_agency_id        sql.NullInt64
		_agency_shortname sql.NullString
		_agency_name      sql.NullString

		_event_type_id     sql.NullInt64
		_event_type_code   sql.NullString
		_event_type_desc   sql.NullString
		_event_log_summary sql.NullInt64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = sqlSelectEventLogWhere
	arrAgencyId := []string{}

	arrParam = append(arrParam, strStartDate)
	arrParam = append(arrParam, strEndDate)

	if param.AgencyID != "" {
		arrAgencyId = strings.Split(param.AgencyID, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyId) > 0 {
		if len(arrAgencyId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.AgencyID, " "))
			sqlCmdWhere += " AND mt.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrAgencyId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND mt.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	sqlCmdQuery := " SELECT agency_id, agency_shortname::jsonb, agency_name::jsonb, event_type_id, event_type_code, event_type_desc::jsonb, COUNT(*) AS summary_event_log " +
		" FROM (" + sqlSelectEventLog + sqlCmdWhere + ") aa " +
		" GROUP BY agency_id, agency_shortname::jsonb, agency_name::jsonb, event_type_id, event_type_code, event_type_desc::jsonb " +
		" ORDER BY agency_name::jsonb->>'th' "

	//Query
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err = db.Query(sqlCmdQuery, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_EventLogSummary_GroupByAgencyCategory, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_agency_id, &_agency_shortname, &_agency_name, &_event_type_id, &_event_type_code, &_event_type_desc, &_event_log_summary)
		if err != nil {
			return nil, err
		}

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_event_type_desc.Valid || _event_type_desc.String == "" {
			_event_type_desc.String = "{}"
		}

		if intAgencyID != _agency_id.Int64 {
			if intAgencyID != 0 {
				objEventLogSummary.ListOfEventCategory = dataEventLogCategory
				data = append(data, objEventLogSummary)
			}

			objEventLogSummary = &Struct_EventLogSummary_GroupByAgencyCategory{}
			objEventLogSummary.Agency = &model_agency.Struct_Agency{}
			objEventLogSummary.Agency.Id = _agency_id.Int64
			objEventLogSummary.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
			objEventLogSummary.Agency.Agency_name = json.RawMessage(_agency_name.String)

			dataEventLogCategory = make([]*model_event_category.Struct_EventLogCategory, 0)

			intAgencyID = _agency_id.Int64
		}

		objEventLogCategory = &model_event_category.Struct_EventLogCategory{}
		objEventLogCategory.ID = _event_type_id.Int64
		objEventLogCategory.Code = _event_type_code.String
		objEventLogCategory.Description = json.RawMessage(_event_type_desc.String)
		objEventLogCategory.SummaryEvent = _event_log_summary.Int64
		dataEventLogCategory = append(dataEventLogCategory, objEventLogCategory)
	}

	if intAgencyID != 0 {
		objEventLogSummary.ListOfEventCategory = dataEventLogCategory
		data = append(data, objEventLogSummary)
	}

	//Return Data
	return data, nil
}

//Description: Get EventLogSummary (Group by EventCategory)
//Input Parameter: start_date*, end_date*, agency_id(multi)
//  Parameters:
//		param
//			Struct_EventLog_InputParam
//  Return:
//		EventLogSummary (Group by EventCategory)
func GetEventLogSummaryGroupByCategory(param *Struct_EventLog_InputParam) ([]*model_event_category.Struct_EventLogCategory, error) {
	//Check Input Parameters
	strStartDate, strEndDate, err := checkInputParam(param)
	if err != nil {
		return nil, err
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data                []*model_event_category.Struct_EventLogCategory
		objEventLogCategory *model_event_category.Struct_EventLogCategory

		_event_type_id     sql.NullInt64
		_event_type_code   sql.NullString
		_event_type_desc   sql.NullString
		_event_log_summary sql.NullInt64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = sqlSelectEventLogWhere
	arrAgencyId := []string{}

	arrParam = append(arrParam, strStartDate)
	arrParam = append(arrParam, strEndDate)

	if param.AgencyID != "" {
		arrAgencyId = strings.Split(param.AgencyID, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyId) > 0 {
		if len(arrAgencyId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.AgencyID, " "))
			sqlCmdWhere += " AND mt.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrAgencyId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND mt.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	sqlCmdQuery := " SELECT event_type_id, event_type_code, event_type_desc::jsonb, COUNT(*) AS summary_event_log " +
		" FROM (" + sqlSelectEventLog + sqlCmdWhere + ") aa " +
		" GROUP BY event_type_id, event_type_code, event_type_desc::jsonb " +
		" ORDER BY event_type_code "

	//Query
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err = db.Query(sqlCmdQuery, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*model_event_category.Struct_EventLogCategory, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_event_type_id, &_event_type_code, &_event_type_desc, &_event_log_summary)
		if err != nil {
			return nil, err
		}

		if !_event_type_desc.Valid || _event_type_desc.String == "" {
			_event_type_desc.String = "{}"
		}

		objEventLogCategory = &model_event_category.Struct_EventLogCategory{}
		objEventLogCategory.ID = _event_type_id.Int64
		objEventLogCategory.Code = _event_type_code.String
		objEventLogCategory.Description = json.RawMessage(_event_type_desc.String)
		objEventLogCategory.SummaryEvent = _event_log_summary.Int64

		data = append(data, objEventLogCategory)
	}

	//Return Data
	return data, nil
}

//Description: Get EventLogSummary (Group by EventCode)
//Input Parameter: start_date*, end_date*, agency_id(multi), event_log_category_id(multi)
//  Parameters:
//		param
//			Struct_EventLog_InputParam
//  Return:
//		EventLogSummary (Group by EventCode)
func GetEventLogSummaryGroupByCode(param *Struct_EventLog_InputParam) ([]*model_event_code.Struct_EventCode_SummaryEvent, error) {
	//Check Input Parameters
	strStartDate, strEndDate, err := checkInputParam(param)
	if err != nil {
		return nil, err
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data         []*model_event_code.Struct_EventCode_SummaryEvent
		objEventCode *model_event_code.Struct_EventCode_SummaryEvent

		_event_type_id   sql.NullInt64
		_event_type_code sql.NullString
		_event_type_desc sql.NullString

		_event_subtype_id   sql.NullInt64
		_event_subtype_code sql.NullString
		_event_subtype_desc sql.NullString
		_event_log_summary  sql.NullInt64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = sqlSelectEventLogWhere
	arrAgencyId := []string{}
	arrEventLogCategoryId := []string{}

	arrParam = append(arrParam, strStartDate)
	arrParam = append(arrParam, strEndDate)

	if param.AgencyID != "" {
		arrAgencyId = strings.Split(param.AgencyID, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyId) > 0 {
		if len(arrAgencyId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.AgencyID, " "))
			sqlCmdWhere += " AND mt.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrAgencyId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND mt.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter event_log_category_id
	if len(arrEventLogCategoryId) > 0 {
		if len(arrEventLogCategoryId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.EventCategoryID, " "))
			sqlCmdWhere += " AND b.event_log_category_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrEventLogCategoryId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND b.event_log_category_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	sqlCmdQuery := " SELECT event_type_id, event_type_code, event_type_desc::jsonb, event_subtype_id, event_subtype_code, event_subtype_desc::jsonb, COUNT(*) AS summary_event_log " +
		" FROM (" + sqlSelectEventLog + sqlCmdWhere + ") aa " +
		" GROUP BY event_type_id, event_type_code, event_type_desc::jsonb, event_subtype_id, event_subtype_code, event_subtype_desc::jsonb " +
		" ORDER BY event_subtype_code "

	//Query
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err = db.Query(sqlCmdQuery, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*model_event_code.Struct_EventCode_SummaryEvent, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_event_type_id, &_event_type_code, &_event_type_desc, &_event_subtype_id, &_event_subtype_code, &_event_subtype_desc, &_event_log_summary)
		if err != nil {
			return nil, err
		}

		if !_event_type_desc.Valid || _event_type_desc.String == "" {
			_event_type_desc.String = "{}"
		}
		if !_event_subtype_desc.Valid || _event_subtype_desc.String == "" {
			_event_subtype_desc.String = "{}"
		}

		objEventCode = &model_event_code.Struct_EventCode_SummaryEvent{}
		objEventCode.ID = _event_subtype_id.Int64
		objEventCode.Code = _event_subtype_code.String
		objEventCode.Description = json.RawMessage(_event_subtype_desc.String)
		objEventCode.SummaryEvent = _event_log_summary.Int64

		objEventCode.EventCategory = &model_event_category.Struct_EventLogCategory{}
		objEventCode.EventCategory.ID = _event_type_id.Int64
		objEventCode.EventCategory.Code = _event_type_code.String
		objEventCode.EventCategory.Description = json.RawMessage(_event_type_desc.String)

		data = append(data, objEventCode)
	}

	//Return Data
	return data, nil
}

//Description: Get EventLogDetail
//Input Parameter: start_date*, end_date*, agency_id(multi), event_log_category_id(multi), event_code(multi)
//  Parameters:
//		param
//			Struct_EventLog_InputParam
//  Return:
//		[]Struct_EventLog
func GetEventLogDetail(param *Struct_EventLog_InputParam) ([]*Struct_EventLog, error) {
	//Check Input Parameters
	strStartDate, strEndDate, err := checkInputParam(param)
	if err != nil {
		return nil, err
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_EventLog
		objEventLog *Struct_EventLog

		_id                 sql.NullInt64
		_event_log_date     time.Time
		_event_log_data     sql.NullString
		_event_log_message  sql.NullString
		_event_log_duration sql.NullInt64

		_metadata_id           sql.NullInt64
		_metadata_service_name sql.NullString
		_metadata_agency_name  sql.NullString

		_agency_id        sql.NullInt64
		_agency_shortname sql.NullString
		_agency_name      sql.NullString

		_event_type_id   sql.NullInt64
		_event_type_code sql.NullString
		_event_type_desc sql.NullString

		_subevent_id   sql.NullInt64
		_subevent_code sql.NullString
		_subevent_desc sql.NullString

		_result *sql.Rows
	)

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = sqlSelectEventLogWhere
	arrAgencyId := []string{}
	arrEventCodeId := []string{}
	arrEventLogCategoryId := []string{}

	arrParam = append(arrParam, strStartDate)
	arrParam = append(arrParam, strEndDate)

	if param.AgencyID != "" {
		arrAgencyId = strings.Split(param.AgencyID, ",")
	}
	if param.EventCategoryID != "" {
		arrEventLogCategoryId = strings.Split(param.EventCategoryID, ",")
	}
	if param.EventCodeID != "" {
		arrEventCodeId = strings.Split(param.EventCodeID, ",")
	}

	//Check Filter agency_id
	if len(arrAgencyId) > 0 {
		if len(arrAgencyId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.AgencyID, " "))
			sqlCmdWhere += " AND mt.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrAgencyId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND mt.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter event_log_category_id
	if len(arrEventLogCategoryId) > 0 {
		if len(arrEventLogCategoryId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.EventCategoryID, " "))
			sqlCmdWhere += " AND b.event_log_category_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrEventLogCategoryId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND b.event_log_category_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter agency_id
	if len(arrEventCodeId) > 0 {
		if len(arrEventCodeId) == 1 {
			arrParam = append(arrParam, strings.Trim(param.EventCodeID, " "))
			sqlCmdWhere += " AND b.event_code_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrEventCodeId {
				arrParam = append(arrParam, strings.Trim(strId, " "))
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND b.event_code_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Query
	//	log.Printf(sqlSelectEventLog+sqlCmdWhere+sqlSelectEventLogOrderBy, arrParam...)
	_result, err = db.Query(sqlSelectEventLog+sqlCmdWhere+sqlSelectEventLogOrderBy, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_EventLog, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_event_log_date, &_event_log_data, &_event_log_message, &_event_log_duration, &_metadata_id, &_metadata_service_name, &_metadata_agency_name, &_agency_id, &_agency_shortname, &_agency_name, &_event_type_id, &_event_type_code, &_event_type_desc, &_subevent_id, &_subevent_code, &_subevent_desc)
		if err != nil {
			return nil, err
		}

		if !_event_log_data.Valid || _event_log_data.String == "" {
			_event_log_data.String = "{}"
		}
		if !_metadata_service_name.Valid || _metadata_service_name.String == "" {
			_metadata_service_name.String = "{}"
		}
		if !_metadata_agency_name.Valid || _metadata_agency_name.String == "" {
			_metadata_agency_name.String = "{}"
		}
		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_event_type_desc.Valid || _event_type_desc.String == "" {
			_event_type_desc.String = "{}"
		}
		if !_subevent_desc.Valid || _subevent_desc.String == "" {
			_subevent_desc.String = "{}"
		}

		//Generate EventLog object
		objEventLog = &Struct_EventLog{}
		objEventLog.Id = _id.Int64
		objEventLog.EventLogDate = _event_log_date.Format(strDatetimeFormat)
		objEventLog.EventLogMessage = _event_log_message.String
		objEventLog.EventLogData = json.RawMessage(_event_log_data.String)
		objEventLog.EventLogDuration = _event_log_duration.Int64
		objEventLog.EventLogDurationTime = _event_log_duration.Int64

		objEventLog.Metadata = &model_metadata.Struct_Metadata{}
		objEventLog.Metadata.Id = _metadata_id.Int64
		objEventLog.Metadata.Metadataservice_Name = json.RawMessage(_metadata_service_name.String)
		objEventLog.Metadata.Metadataagency_Name = json.RawMessage(_metadata_agency_name.String)

		objEventLog.Agency = &model_agency.Struct_Agency{}
		objEventLog.Agency.Id = _agency_id.Int64
		objEventLog.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		objEventLog.Agency.Agency_name = json.RawMessage(_agency_name.String)

		objEventLog.EventCode = &model_event_code.Struct_EventCode{}
		objEventLog.EventCode.ID = _subevent_id.Int64
		objEventLog.EventCode.Code = _subevent_code.String
		objEventLog.EventCode.Description = json.RawMessage(_subevent_desc.String)

		objEventLog.EventCode.EventCategory = &model_event_category.Struct_EventLogCategory{}
		objEventLog.EventCode.EventCategory.ID = _event_type_id.Int64
		objEventLog.EventCode.EventCategory.Code = _event_type_code.String
		objEventLog.EventCode.EventCategory.Description = json.RawMessage(_event_type_desc.String)

		data = append(data, objEventLog)
	}

	//Return Data
	return data, nil
}

// check input param
//  Parameters:
//		param
//				Struct_EventLog_InputParam
//  Return:
//		start date, enddate, error
func checkInputParam(param *Struct_EventLog_InputParam) (string, string, error) {

	var strStartDate string
	var strEndDate string

	//Check Parameters
	if param.StartDate == "" {
		return strStartDate, strEndDate, errors.New("'start_date' is not null.")
	} else {
		if len(strings.Split(strings.TrimSpace(param.StartDate), " ")) == 1 {
			strStartDate = strings.TrimSpace(param.StartDate) + " 00:00:00"
		} else {
			strStartDate = strings.TrimSpace(param.StartDate) + ":00"
		}
	}

	if param.EndDate == "" {
		return strStartDate, strEndDate, errors.New("'end_date' is not null.")
	} else {
		if len(strings.Split(strings.TrimSpace(param.EndDate), " ")) == 1 {
			strEndDate = strings.TrimSpace(param.EndDate) + " 23:59:59"
		} else {
			strEndDate = strings.TrimSpace(param.EndDate) + ":59"
		}
	}

	return strStartDate, strEndDate, nil
}

//Description: Get EventLogSummary for report (Group by EventCategory)
//Input Parameter: start_date*, end_date*
//  Parameters:
//		param
//			Struct_EventLog_InputParam
//  Return:
//		[]Struct_EventLogSummaryReport
func GetEventLogSummaryReport(param *Struct_EventLog_InputParam) ([]*Struct_EventLogSummaryReport, error) {
	//Check Input Parameters
	strStartDate, strEndDate, err := checkInputParam(param)
	if err != nil {
		return nil, err
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data []*Struct_EventLogSummaryReport

		objEventCatReport *Struct_EventLogSummaryReport
		objAgency         *model_agency.Struct_Agency
		objEventCode      *Struct_EventCodeSummaryReport

		_event_type_id     sql.NullInt64
		_event_type_code   sql.NullString
		_event_type_desc   sql.NullString
		_event_log_count   sql.NullInt64
		_event_log_percent sql.NullFloat64

		_event_subtype_id   sql.NullInt64
		_event_subtype_code sql.NullString
		_event_subtype_desc sql.NullString

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_result *sql.Rows

		intEventCateID int64 = -1
		intEventCodeID int64 = -1
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	arrParam = append(arrParam, strStartDate)
	arrParam = append(arrParam, strEndDate)

	sqlCmdQuery := " SELECT ec.id AS event_type_id, ec.code AS event_type_code, ec.description AS event_type_desc " +
		" 	  , ab.event_subtype_id, ab.event_subtype_code, ab.event_subtype_desc " +
		"	  , ab.agency_id, ab.agency_shortname, ab.agency_name " +
		"	  , ab.event_cnt, (((CASE WHEN ab.event_type_id IS NULL THEN 0 ELSE SUM(ab.event_cnt) OVER (PARTITION BY ec.id) END) / SUM(ab.event_cnt) OVER ())*100) AS event_percent " +
		" FROM api.lt_event_log_category ec " +
		" LEFT JOIN ( SELECT event_type_id, event_subtype_id, event_subtype_code, event_subtype_desc::jsonb, aa.agency_id, aa.agency_shortname::jsonb, aa.agency_name::jsonb, COUNT(*) AS event_cnt " +
		"  			 FROM (" + sqlSelectEventLog + ") aa " +
		"  			 GROUP BY event_type_id, event_subtype_id, event_subtype_code, event_subtype_desc::jsonb, agency_id, agency_shortname::jsonb, agency_name::jsonb " +
		"			) ab ON ec.id = ab.event_type_id " +
		" WHERE ec.deleted_by IS NULL AND (ec.deleted_at IS NULL OR ec.deleted_at = '1970-01-01 07:00:00+07') " +
		" ORDER BY event_type_id, event_subtype_id, agency_id "

	//Query
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err = db.Query(sqlCmdQuery, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_EventLogSummaryReport, 0)

	//objAgency = make(*model_agency.Struct_Agency, 0)
	//objEventCode = make(*Struct_EventCodeSummaryReport, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_event_type_id, &_event_type_code, &_event_type_desc,
			&_event_subtype_id, &_event_subtype_code, &_event_subtype_desc,
			&_agency_id, &_agency_shortname, &_agency_name,
			&_event_log_count, &_event_log_percent)
		if err != nil {
			return nil, err
		}

		if intEventCateID != _event_type_id.Int64 {
			if intEventCateID != -1 {
				if objEventCode != nil {
					objEventCatReport.ListEventCode = append(objEventCatReport.ListEventCode, objEventCode)
				}
				data = append(data, objEventCatReport)
			}

			if !_event_type_desc.Valid || _event_type_desc.String == "" {
				_event_type_desc.String = "{}"
			}
			objEventCatReport = &Struct_EventLogSummaryReport{}
			objEventCatReport.ID = _event_type_id.Int64
			objEventCatReport.Code = _event_type_code.String
			objEventCatReport.Description = json.RawMessage(_event_type_desc.String)
			objEventCatReport.PercentEvent = _event_log_percent.Float64

			//objEventCatReport.ListEventCode = []*Struct_EventCodeSummaryReport{}
			//objEventCatReport.ListEventCode = nil
			objEventCode = nil
			objAgency = nil

			intEventCateID = _event_type_id.Int64
		}

		if _event_subtype_id.Int64 != 0 {
			if intEventCodeID != _event_subtype_id.Int64 {
				if objEventCode != nil {
					objEventCatReport.ListEventCode = append(objEventCatReport.ListEventCode, objEventCode)
				}

				if !_event_subtype_desc.Valid || _event_subtype_desc.String == "" {
					_event_subtype_desc.String = "{}"
				}
				objEventCode = &Struct_EventCodeSummaryReport{}
				objEventCode.ID = _event_subtype_id.Int64
				objEventCode.Code = _event_subtype_code.String
				objEventCode.Description = json.RawMessage(_event_subtype_desc.String)
				//objEventCode.ListAgency = []*model_agency.Struct_Agency{}
				//objAgency = make(*model_agency.Struct_Agency, 0)
				intEventCodeID = _event_subtype_id.Int64
			}
		}

		if _agency_id.Int64 != 0 {
			if !_agency_shortname.Valid || _agency_shortname.String == "" {
				_agency_shortname.String = "{}"
			}
			if !_agency_name.Valid || _agency_name.String == "" {
				_agency_name.String = "{}"
			}
			objAgency = &model_agency.Struct_Agency{}
			objAgency.Id = _agency_id.Int64
			objAgency.Agency_shortname = json.RawMessage(_agency_shortname.String)
			objAgency.Agency_name = json.RawMessage(_agency_name.String)
			objEventCode.ListAgency = append(objEventCode.ListAgency, objAgency)
		}
	}

	if intEventCateID != -1 {
		if objEventCode != nil {
			objEventCatReport.ListEventCode = append(objEventCatReport.ListEventCode, objEventCode)
		}
		data = append(data, objEventCatReport)
	}

	//Return Data
	return data, nil
}

//	Description: Get Eventreport
//	Input Parameter: start_date*, end_date*
//  Parameters:
//		param
//			Struct_EventReport_Input
//  Return:
//		[]Struct_EventReport
func GetEventReport(param *Struct_EventReport_Input) ([]*Struct_EventReport, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(SQL_EventReport, param.Date, param.Agent, param.Event)
	if err != nil {
		return nil, err
	}

	rs := make([]*Struct_EventReport, 0)
	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	for rows.Next() {
		var _datetime time.Time
		s := &Struct_EventReport{}
		rows.Scan(&s.Type, &s.Id, &s.LogId, &_datetime, &s.Name, &s.Detail)

		s.DateTime = _datetime.Format(strDatetimeFormat)
		rs = append(rs, s)
	}

	return rs, nil
}
