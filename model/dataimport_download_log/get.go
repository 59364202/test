package dataimport_download_log

import (
	"database/sql"
	"encoding/json"
	"fmt"

	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//	"log"
	"strconv"
	"strings"
	"time"
)

func GetOverAllPercentDownload(param *Struct_DownloadLog_Inputparam) ([]*Struct_DownloadLog_Summary_Agency, error) {

	//Check Input Parameters
	if param.Year == "" {
		return nil, errors.Repack(errors.New("'year' is not null."))
	} else {
		arrYear := []string{}
		arrYear = strings.Split(param.Year, ",")
		if len(arrYear) > 1 {
			return nil, errors.Repack(errors.New("'year' is not multiple."))
		}
	}

	if param.Month != "" {
		arrMonth := []string{}
		arrMonth = strings.Split(param.Month, ",")
		if len(arrMonth) > 1 {
			return nil, errors.Repack(errors.New("'month' is not multiple."))
		}
	}

	return getPercentDownloadInfo(param.Year, param.Month, "", "", param.ConnectionFormat)
}

func GetPercentDownload(param *Struct_DownloadLog_Inputparam) ([]*Struct_DownloadLog_Summary_Agency, error) {

	//Check Input Parameters
	strStartDate, strEndDate, err := checkInputParam(param)
	if err != nil {
		return nil, err
	}
	/*
		strStartDate := param.StartDate
		strEndDate := param.EndDate

		//Check Input Parameters
		if (strStartDate == ""){
			return nil, errors.Repack(errors.New("'start_date' is not null."))
		}else{
			arrStartDate := []string{}
			arrStartDate = strings.Split(strStartDate, " ")
			if (len(arrStartDate) == 1){
				//return nil, errors.Repack(errors.New("'start_date' is not multiple."))
				strStartDate += " 00:00:00"
			}
		}

		if (param.EndDate == ""){
			return nil, errors.Repack(errors.New("'end_date' is not null."))
		}else{
			arrEndDate := []string{}
			arrEndDate = strings.Split(param.EndDate, ",")
			if (len(arrEndDate) == 1){
				//return nil, errors.Repack(errors.New("'end_date' is not multiple."))
				strEndDate += " 23:59:59"
			}
		}*/

	return getPercentDownloadInfo("", "", strStartDate, strEndDate, param.ConnectionFormat)
}

//Description: Get PercentDownload (Group by Agency)
//Input Parameter: year, month, start_date, end_date
func getPercentDownloadInfo(year string, month string, start_date string, end_date string, connection_format string) ([]*Struct_DownloadLog_Summary_Agency, error) {

	//-- Check Filter by parameters --//
	var sqlCmdWhereOnline string
	var sqlCmdWhereOffline string
	var sqlCmdDateRange string
	var sqlSelectOverallDownloadCount string
	var sqlCmdQuery string

	//Check Input Parameters
	if year == "" {
		sqlCmdWhereOffline = " AND dll.download_begin_at BETWEEN '" + start_date + "' AND '" + end_date + "' "
		sqlCmdWhereOnline = " AND dll.download_begin_at BETWEEN '" + start_date + "' AND '" + end_date + "' "
		sqlCmdDateRange = " * DATE_PART('day', '" + end_date + "'::timestamp - '" + start_date + "'::timestamp + '1 DAY'::INTERVAL) "
	} else {
		var (
			dateEnd   time.Time
			dateStart time.Time
			err       error
		)
		if month == "" {
			dateStart, err = time.Parse("2006-1-02", year+"-01-01")
			if err != nil {
				return nil, err
			}
			dateEnd = dateStart.AddDate(1, 0, 0)
			// ds := dateStart.Format(time.RFC3339Nano)
			// de := dateEnd.Format(time.RFC3339Nano)
			// sqlCmdWhereOffline = " AND dll.download_begin_at >= '" + ds + "' AND download_begin_at < '" + de + "' "
			// sqlCmdWhereOnline += " AND dll.download_begin_at >= '" + ds + "' AND download_begin_at < '" + de + "' "
			sqlCmdDateRange = " * DATE_PART('DOY', (CASE WHEN '" + year + "-12-31'::DATE > NOW() THEN NOW() ELSE '" + year + "-12-31'::DATE END)) "
		} else {
			dateStart, err = time.Parse("2006-1-02", year+"-"+month+"-01")
			if err != nil {
				return nil, err
			}
			dateEnd = dateStart.AddDate(0, 1, 0) //

			sqlCmdDateRange = " * (CASE WHEN EXTRACT(YEAR FROM NOW()) = '" + month + "' AND EXTRACT(MONTH FROM NOW()) = '" + month + "' THEN DATE_PART('DAY', NOW()) ELSE DATE_PART('DAY', '" + year + "-" + month + "-01'::DATE + '1 MONTH'::INTERVAL - '1 DAY'::INTERVAL) END) "
		}
		ds := dateStart.Format("2006-01-02 15:04")
		de := dateEnd.Format("2006-01-02 15:04")

		sqlCmdWhereOffline = " AND dll.download_begin_at >= '" + ds + "' AND download_begin_at < '" + de + "' "
		sqlCmdWhereOnline += " AND dll.download_begin_at >= '" + ds + "' AND download_begin_at < '" + de + "' "
	}
	switch connection_format {
	case "online":
		sqlSelectOverallDownloadCount = strings.Replace(strings.Replace(sqlSelectOverAllDownloadCount_Online, "$sqlCmdDateRange", sqlCmdDateRange, 1), "$sqlCmdWhereOnline", sqlCmdWhereOnline, 1)
		sqlCmdQuery = `
		SELECT tmt.agency_id 
				, agt.agency_shortname::jsonb 
				, agt.agency_name::jsonb 
				, SUM(expected_download_count) AS expected_download 
				, SUM(actual_download_count) AS actual_download_count 
				, SUM(download_count) AS download_count 
				, SUM(download_files_count) AS download_files_count 
				, SUM(import_record_count) AS import_record_count 
				, CASE WHEN SUM(expected_download_count) = 0 THEN 0 ELSE (SUM(download_count)/SUM(expected_download_count))*100 END AS percent_download 
		FROM ( ` + sqlSelectOverallDownloadCount + `) tmt 
		LEFT JOIN agency agt ON agt.id = tmt.agency_id 
		WHERE (agt.deleted_by IS NULL AND (agt.deleted_at IS NULL OR agt.deleted_at = '1970-01-01 07:00:00')) 
		GROUP BY tmt.agency_id, agt.agency_shortname::jsonb, agt.agency_name::jsonb 
		ORDER BY agt.agency_name::jsonb->>'th' 
		`

		break
	case "offline":
		sqlSelectOverallDownloadCount = strings.Replace(sqlSelectOverAllDownloadCount_Offline, "$sqlCmdWhereOffline", sqlCmdWhereOffline, 1)
		sqlCmdQuery = `
		SELECT tmt.agency_id 
				, agt.agency_shortname::jsonb 
				, agt.agency_name::jsonb 
				, SUM(expected_download_count) AS expected_download 
				, SUM(actual_download_count) AS actual_download_count 
				, SUM(download_count) AS download_count 
				, SUM(download_files_count) AS download_files_count 
				, SUM(import_record_count) AS import_record_count 
				, 0 AS percent_download 
		FROM ( ` + sqlSelectOverallDownloadCount + `) tmt 
		LEFT JOIN agency agt ON agt.id = tmt.agency_id 
		WHERE (agt.deleted_by IS NULL AND (agt.deleted_at IS NULL OR agt.deleted_at = '1970-01-01 07:00:00')) 
		GROUP BY tmt.agency_id, agt.agency_shortname::jsonb, agt.agency_name::jsonb 
		ORDER BY agt.agency_name::jsonb->>'th' 
		`
		break
		//	default:
		//		sqlSelectOverallDownloadCount = strings.Replace(strings.Replace(sqlSelectOverAllDownloadCount_Online, "$sqlCmdDateRange", sqlCmdDateRange, 1), "$sqlCmdWhereOnline", sqlCmdWhereOnline, 1) +
		//			`
		//			UNION
		//			` +
		//			strings.Replace(sqlSelectOverAllDownloadCount_Offline, "$sqlCmdWhereOffline", sqlCmdWhereOffline, 1)
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data               []*Struct_DownloadLog_Summary_Agency
		objPercentDownload *Struct_DownloadLog_Summary_Agency

		_agency_id        sql.NullInt64
		_agency_shortname sql.NullString
		_agency_name      sql.NullString

		_expected_download_count sql.NullFloat64
		_actual_download_count   sql.NullInt64
		_download_count          sql.NullInt64
		_number_of_file          sql.NullInt64
		_number_of_record        sql.NullInt64
		_percent_download_count  sql.NullFloat64

		_result *sql.Rows
	)

	//Query
	//	log.Printf(sqlCmdQuery)
	fmt.Println(sqlCmdQuery)
	_result, err = db.Query(sqlCmdQuery)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_DownloadLog_Summary_Agency, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_agency_id, &_agency_shortname, &_agency_name, &_expected_download_count, &_actual_download_count, &_download_count, &_number_of_file, &_number_of_record, &_percent_download_count)
		if err != nil {
			return nil, err
		}

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		objPercentDownload = &Struct_DownloadLog_Summary_Agency{}
		objPercentDownload.NumberOfExpectedDownload = _expected_download_count.Float64
		objPercentDownload.NumberOfDownloadCountActual = _actual_download_count.Int64
		objPercentDownload.NumberOfDownloadCount = _download_count.Int64
		objPercentDownload.PercentDownloadCount = _percent_download_count.Float64
		objPercentDownload.NumberOfFileDownload = _number_of_file.Int64
		objPercentDownload.NumberOfDataRecord = _number_of_record.Int64

		objPercentDownload.Agency = &model_agency.Struct_Agency{}
		objPercentDownload.Agency.Id = _agency_id.Int64
		objPercentDownload.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		objPercentDownload.Agency.Agency_name = json.RawMessage(_agency_name.String)

		data = append(data, objPercentDownload)
	}

	//Return Data
	return data, nil
}

//Description: Get Download Size (Group by Each Date in a Month)
//Input Parameter: year*, month*, agency_id*
func GetDownloadSizeByAgency(param *Struct_DownloadLog_Inputparam) ([]*Struct_DownloadLog, error) {

	//Check Input Parameters
	if param.Year == "" {
		return nil, errors.Repack(errors.New("'year' is not null."))
	} else {
		arrYear := []string{}
		arrYear = strings.Split(param.Year, ",")
		if len(arrYear) > 1 {
			return nil, errors.Repack(errors.New("'year' is not multiple."))
		}
	}

	if param.Month == "" {
		return nil, errors.Repack(errors.New("'month' is not null."))
	} else {
		arrMonth := []string{}
		arrMonth = strings.Split(param.Month, ",")
		if len(arrMonth) > 1 {
			return nil, errors.Repack(errors.New("'month' is not multiple."))
		}
	}

	if param.AgencyID == "" {
		return nil, errors.Repack(errors.New("'agency_id' is not null."))
	} else {
		arrAgencyID := []string{}
		arrAgencyID = strings.Split(param.AgencyID, ",")
		if len(arrAgencyID) > 1 {
			return nil, errors.Repack(errors.New("'agency_id' is not multiple."))
		}
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data            []*Struct_DownloadLog
		objDownloadSize *Struct_DownloadLog

		_download_date    time.Time
		_download_size    sql.NullFloat64
		_number_of_file   sql.NullInt64
		_number_of_record sql.NullInt64

		_result *sql.Rows
	)

	//Query
	//	log.Printf(sqlSelectMonthlyDownloadSizeByAgency, param.Month, param.Year, param.AgencyID, param.Year+"-"+param.Month+"-01")
	_result, err = db.Query(sqlSelectMonthlyDownloadSizeByAgency, param.Month, param.Year, param.AgencyID, param.Year+"-"+param.Month+"-01")
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_DownloadLog, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_download_date, &_download_size, &_number_of_file, &_number_of_record)
		if err != nil {
			return nil, err
		}

		objDownloadSize = &Struct_DownloadLog{}
		objDownloadSize.DownloadDate = _download_date.Format(strDatetimeFormat)
		objDownloadSize.NumberOfDownloadSize = _download_size.Float64
		objDownloadSize.NumberOfFileDownload = _number_of_file.Int64
		objDownloadSize.NumberOfDataRecord = _number_of_record.Int64

		data = append(data, objDownloadSize)
	}

	//Return Data
	return data, nil
}

//Description: Get PercentDownloadDetail (Group by Metadata)
//Input Parameter: start_date, end_date, agency_id
func GetPercentDownloadDetail(param *Struct_DownloadLog_Inputparam) ([]*Struct_DownloadLog_Summary_Metadata, error) {

	//Check Input Parameters
	if param.StartDate == "" {
		return nil, errors.Repack(errors.New("'start_date' is not null."))
	} else {
		arrStartDate := []string{}
		arrStartDate = strings.Split(param.StartDate, ",")
		if len(arrStartDate) > 1 {
			return nil, errors.Repack(errors.New("'start_date' is not multiple."))
		}
	}

	if param.EndDate == "" {
		return nil, errors.Repack(errors.New("'end_date' is not null."))
	} else {
		arrEndDate := []string{}
		arrEndDate = strings.Split(param.EndDate, ",")
		if len(arrEndDate) > 1 {
			return nil, errors.Repack(errors.New("'end_date' is not multiple."))
		}
	}

	if param.AgencyID == "" {
		return nil, errors.Repack(errors.New("'agency_id' is not null."))
	} else {
		arrAgencyID := []string{}
		arrAgencyID = strings.Split(param.AgencyID, ",")
		if len(arrAgencyID) > 1 {
			return nil, errors.Repack(errors.New("'agency_id' is not multiple."))
		}
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data               []*Struct_DownloadLog_Summary_Metadata
		objPercentDownload *Struct_DownloadLog_Summary_Metadata

		_metadata_id           sql.NullInt64
		_metadata_agency_name  sql.NullString
		_metadata_service_name sql.NullString

		_expected_download_count sql.NullFloat64
		_actual_download_count   sql.NullInt64
		_download_count          sql.NullInt64
		_number_of_file          sql.NullInt64
		_number_of_record        sql.NullInt64
		_percent_download_count  sql.NullFloat64
		_lastest_import_date     sql.NullString

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var sqlCmdWhereOnline string = " AND dll.download_begin_at BETWEEN '" + param.StartDate + " 00:00:00' AND '" + param.EndDate + " 23:59:59' "
	var sqlCmdWhereOffline string = " AND dll.download_begin_at BETWEEN '" + param.StartDate + " 00:00:00' AND '" + param.EndDate + " 23:59:59' "
	var sqlCmdDateRange string = " * DATE_PART('day', '" + param.EndDate + " 23:59:59'::timestamp - '" + param.StartDate + " 00:00:00'::timestamp + '1 DAY'::INTERVAL) "
	var sqlSelectOverallDownloadCount string = ""

	switch param.ConnectionFormat {
	case "online":
		sqlSelectOverallDownloadCount = strings.Replace(strings.Replace(sqlSelectOverAllDownloadCount_Online, "$sqlCmdDateRange", sqlCmdDateRange, 1), "$sqlCmdWhereOnline", sqlCmdWhereOnline, 1)
		break
	case "offline":
		sqlSelectOverallDownloadCount = strings.Replace(sqlSelectOverAllDownloadCount_Offline, "$sqlCmdWhereOffline", sqlCmdWhereOffline, 1)
		break
	default:
		sqlSelectOverallDownloadCount = strings.Replace(strings.Replace(sqlSelectOverAllDownloadCount_Online, "$sqlCmdDateRange", sqlCmdDateRange, 1), "$sqlCmdWhereOnline", sqlCmdWhereOnline, 1) +
			`
			UNION 
			` +
			strings.Replace(sqlSelectOverAllDownloadCount_Offline, "$sqlCmdWhereOffline", sqlCmdWhereOffline, 1)
	}

	sqlCmdQuery := " SELECT tmt.metadata_id " +
		"	  , tmt.metadataagency_name::jsonb " +
		"	  , tmt.metadataservice_name::jsonb " +
		"	  , SUM(expected_download_count) AS expected_download " +
		"	  , SUM(actual_download_count) AS actual_download_count " +
		"	  , SUM(download_count) AS download_count " +
		"	  , SUM(download_files_count) AS download_files_count " +
		"	  , SUM(import_record_count) AS import_record_count " +
		"	  , CASE WHEN SUM(expected_download_count) = 0 THEN 0 ELSE (SUM(download_count)/SUM(expected_download_count))*100 END AS percent_download " +
		" 	  , MAX(lastest_import) AS lastest_import " +
		" FROM ( " + sqlSelectOverallDownloadCount + ") tmt " +
		" LEFT JOIN agency agt ON agt.id = tmt.agency_id " +
		" WHERE tmt.agency_id = $1 " +
		" GROUP BY tmt.metadata_id, tmt.metadataagency_name::jsonb, tmt.metadataservice_name::jsonb " +
		" ORDER BY tmt.metadataservice_name::jsonb->>'th' "

	//Query
	//	log.Printf(sqlCmdQuery, param.AgencyID)
	_result, err = db.Query(sqlCmdQuery, param.AgencyID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_DownloadLog_Summary_Metadata, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_metadata_id, &_metadata_agency_name, &_metadata_service_name, &_expected_download_count, &_actual_download_count, &_download_count, &_number_of_file, &_number_of_record, &_percent_download_count, &_lastest_import_date)
		if err != nil {
			return nil, err
		}

		if !_metadata_agency_name.Valid || _metadata_agency_name.String == "" {
			_metadata_agency_name.String = "{}"
		}
		if !_metadata_service_name.Valid || _metadata_service_name.String == "" {
			_metadata_service_name.String = "{}"
		}

		objPercentDownload = &Struct_DownloadLog_Summary_Metadata{}
		objPercentDownload.NumberOfExpectedDownload = _expected_download_count.Float64
		objPercentDownload.NumberOfDownloadCountActual = _actual_download_count.Int64
		objPercentDownload.NumberOfDownloadCount = _download_count.Int64
		objPercentDownload.PercentDownloadCount = _percent_download_count.Float64
		objPercentDownload.NumberOfFileDownload = _number_of_file.Int64
		objPercentDownload.NumberOfDataRecord = _number_of_record.Int64

		if !_lastest_import_date.Valid || _lastest_import_date.String == "" {
			objPercentDownload.DownloadLastestDate = ""
		} else {
			t, err := time.Parse(time.RFC3339, _lastest_import_date.String)
			if err != nil {
				objPercentDownload.DownloadLastestDate = ""
			} else {
				objPercentDownload.DownloadLastestDate = t.Format(strDatetimeFormat)
			}
		}

		objPercentDownload.Metadata = &model_metadata.Struct_Metadata{}
		objPercentDownload.Metadata.Id = _metadata_id.Int64
		objPercentDownload.Metadata.Metadataagency_Name = json.RawMessage(_metadata_agency_name.String)
		objPercentDownload.Metadata.Metadataservice_Name = json.RawMessage(_metadata_service_name.String)

		data = append(data, objPercentDownload)
	}

	//Return Data
	return data, nil
}

//Description: Get PercentDownloadYearlyCompare (Group by Month)
//Input Parameter: agency_id, year*
func PercentDownloadYearlyCompare(param *Struct_DownloadLog_Inputparam) ([]*Struct_DownloadLog_YearlyCompare, error) {

	//Check Input Parameters
	if param.Year == "" {
		return nil, errors.Repack(errors.New("'year' is not null."))
	}

	if param.AgencyID == "" {
		return nil, errors.Repack(errors.New("'agency_id' is not null."))
	} else {
		arrAgencyID := []string{}
		arrAgencyID = strings.Split(param.AgencyID, ",")
		if len(arrAgencyID) > 1 {
			return nil, errors.Repack(errors.New("'agency_id' is not multiple."))
		}
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data               []*Struct_DownloadLog_YearlyCompare
		objPercentDownload *Struct_DownloadLog_YearlyCompare

		//objDownloadFileCount	map[int64]int64
		//objImportRecord			map[int64]int64
		//objPercentDownload		map[int64]float64

		intYear int64 = 0

		_year                    sql.NullInt64
		_month                   sql.NullInt64
		_expected_download_count sql.NullInt64
		_actual_download_count   sql.NullInt64
		_download_count          sql.NullInt64
		_number_of_file          sql.NullInt64
		_number_of_record        sql.NullInt64
		_percent_download_count  sql.NullFloat64

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdSelectTempDate string = ""
	var sqlCmdWhere string = ""

	//	var sqlCmdWhereOffline string = ""
	var sqlCmdSelectTempDateOffline string = ""
	var sqlSelectCompareDownloadCount string = ""

	arrParam = append(arrParam, param.AgencyID)

	arrYear := []string{}
	arrYear = strings.Split(param.Year, ",")
	if len(arrYear) == 1 {
		arrParam = append(arrParam, strings.Trim(param.Year, " "))
		sqlCmdWhere = " AND EXTRACT(YEAR FROM dll.download_begin_at) = $" + strconv.Itoa(len(arrParam))
		sqlCmdSelectTempDate = "SELECT GENERATE_SERIES('" + param.Year + "-01-01'::DATE, '" + param.Year + "-12-31'::DATE, '1 day')::DATE AS dt_date"
		sqlCmdSelectTempDateOffline = " SELECT " + param.Year + " as year "
	} else {
		arrSqlCmd := []string{}
		arrSqlCmdSelectTempDate := []string{}
		arrSqlCmdSelectTempDateOffline := []string{}
		for _, strYear := range arrYear {
			arrParam = append(arrParam, strings.Trim(strYear, " "))
			arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			arrSqlCmdSelectTempDate = append(arrSqlCmdSelectTempDate, "SELECT GENERATE_SERIES('"+strYear+"-01-01'::DATE, '"+strYear+"-12-31'::DATE, '1 day')::DATE AS dt_date")
			arrSqlCmdSelectTempDateOffline = append(arrSqlCmdSelectTempDateOffline, " SELECT "+strYear+" as year ")
		}
		sqlCmdWhere = " AND EXTRACT(YEAR FROM dll.download_begin_at) IN (" + strings.Join(arrSqlCmd, ",") + ")"
		sqlCmdSelectTempDate = strings.Join(arrSqlCmdSelectTempDate, " UNION ")
		sqlCmdSelectTempDateOffline = strings.Join(arrSqlCmdSelectTempDateOffline, " UNION ")
	}

	switch param.ConnectionFormat {
	case "online":
		sqlSelectCompareDownloadCount = strings.Replace(strings.Replace(sqlSelectCompareDownloadCount_Online, "$sqlCmdSelectTempDate", sqlCmdSelectTempDate, 1), "$sqlCmdWhere", sqlCmdWhere, 1)
		break
	case "offline":
		sqlSelectCompareDownloadCount = strings.Replace(strings.Replace(sqlSelectCompareDownloadCount_Offline, "$sqlCmdSelectTempDateOffline", sqlCmdSelectTempDateOffline, 1), "$sqlCmdWhereOffline", sqlCmdWhere, 1)
		break
	default:
		sqlSelectCompareDownloadCount = strings.Replace(strings.Replace(sqlSelectCompareDownloadCount_Online, "$sqlCmdSelectTempDate", sqlCmdSelectTempDate, 1), "$sqlCmdWhere", sqlCmdWhere, 1) +
			`
			UNION 
			` +
			strings.Replace(strings.Replace(sqlSelectCompareDownloadCount_Offline, "$sqlCmdSelectTempDateOffline", sqlCmdSelectTempDateOffline, 1), "$sqlCmdWhereOffline", sqlCmdWhere, 1)
	}
	sqlCmdQuery := `
	SELECT download_year
			, download_month
			, SUM(expected_download_count) AS expected_download_count
			, SUM(actual_download_count) AS actual_download_count
			, SUM(download_count) AS download_count
			, SUM(download_files_count) AS download_files_count
			, SUM(import_record_count) AS import_record_count
			, CASE WHEN SUM(expected_download_count) = 0 THEN 0 ELSE (SUM(download_count)/SUM(expected_download_count))*100 END AS percent_download
	FROM ( ` + sqlSelectCompareDownloadCount + ` ) aa
	GROUP BY download_year, download_month
	ORDER BY download_year, download_month 
	`
	//	sqlCmdQuery := strings.Replace(strings.Replace(sqlSelectCompareDownloadCount, "$sqlCmdSelectTempDate", sqlCmdSelectTempDate, 1), "$sqlCmdWhere", sqlCmdWhere, 1)

	//Query
	//	log.Printf(sqlCmdQuery, arrParam...)
	_result, err = db.Query(sqlCmdQuery, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_DownloadLog_YearlyCompare, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_year, &_month, &_expected_download_count, &_actual_download_count, &_download_count, &_number_of_file, &_number_of_record, &_percent_download_count)
		if err != nil {
			return nil, err
		}
		//		log.Println(intYear)
		if intYear != _year.Int64 {
			if intYear != 0 {
				objPercentDownload.Year = intYear
				//objPercentDownloadCompare.NumberOfFileDownload = objDownloadFileCount
				//objPercentDownloadCompare.NumberOfDataRecord = objImportRecord
				//objPercentDownloadCompare.PercentDownloadCount = objPercentDownload

				//				log.Println(objPercentDownload.Year)
				//				log.Println("--------")
				data = append(data, objPercentDownload)
			}
			objPercentDownload = &Struct_DownloadLog_YearlyCompare{}
			objPercentDownload.NumberOfFileDownload = []int64{}
			objPercentDownload.NumberOfDataRecord = []int64{}
			objPercentDownload.PercentDownloadCount = []float64{}

			intYear = _year.Int64

			//objDownloadFileCount = map[int64]int64{}
			//objImportRecord = map[int64]int64{}
			//objPercentDownload = map[int64]float64{}

			objPercentDownload.Year = intYear
		}

		objPercentDownload.NumberOfFileDownload = append(objPercentDownload.NumberOfFileDownload, _number_of_file.Int64)
		objPercentDownload.NumberOfDataRecord = append(objPercentDownload.NumberOfDataRecord, _number_of_record.Int64)
		objPercentDownload.PercentDownloadCount = append(objPercentDownload.PercentDownloadCount, _percent_download_count.Float64)
		//objDownloadFileCount[_month.Int64] = _number_of_file.Int64
		//objImportRecord[_month.Int64] = _number_of_record.Int64
		//objPercentDownload[_month.Int64] = _percent_download_count.Float64
	}

	if intYear != 0 {
		objPercentDownload.Year = intYear
		//objPercentDownloadCompare.NumberOfFileDownload = objDownloadFileCount
		//objPercentDownloadCompare.NumberOfDataRecord = objImportRecord
		//objPercentDownloadCompare.PercentDownloadCount = objPercentDownload
		//		log.Println(objPercentDownload.Year)
		//		log.Println("++++++++")
		data = append(data, objPercentDownload)
	}

	//Return Data
	return data, nil
}

func checkInputParam(param *Struct_DownloadLog_Inputparam) (string, string, error) {

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
