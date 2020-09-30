package metadata

import (
	"database/sql"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
	"time"
)

var sqlGetMetadata = "SELECT m.id, m.subcategory_id, m.agency_id " +
	" , m.dataunit_id, m.dataformat_id , m.servicemethod_id " +
	" , m.connection_format, m.metadata_method, m.metadata_channel " +
	" , f.datafrequency AS metadata_datafrequency, m.metadata_convertfrequency, m.metadata_contact " +
	" , m.metadata_agencystoredate, m.metadata_startdatadate, m.scriptname " +
	" , m.metadata_update_plan, m.metadata_laws, m.metadata_remark " +
	" , m.metadata_status , l.code , mt.metadataagency_name " +
	" , mt.metadataservice_name , mt.metadata_tag , mt.metadata_description " +
	" FROM metadata m " +
	" LEFT JOIN metadata_translate mt ON m.id = mt.metadata_id " +
	" LEFT JOIN language l ON l.id = mt.language_id " +
	" LEFT JOIN metadata_frequency f ON m.id = f.metadata_id " +
	" WHERE m.deleted_by IS NULL AND mt.deleted_by IS NULL AND f.deleted_by IS NULL AND l.deleted_by IS NULL "
var sqlGetMetadataWhereId = " AND m.id = $1 "
var sqlGetMetadataOrderBy = " ORDER BY m.id, f.datafrequency, l.id "

var sqlGetHydroMetadata = " SELECT h.hydroinfo_id " +
	" FROM metadata_hydroinfo h " +
	" WHERE h.metadata_id = $1 AND h.deleted_by IS NULL "

var sqlGetHistoryMetadata = " SELECT h.metadata_datetime AS history_datetime , h.history_description , u.full_name " +
	" FROM metadata_history h " +
	" LEFT JOIN api.user u ON h.created_by = u.id " +
	" WHERE h.deleted_by IS NULL AND h.metadata_id = $1 "

var sqlGetMetadataFrequency = " SELECT datafrequency FROM metadata_frequency WHERE metadata_id = $1 ORDER BY metadata_id, datafrequency "

func GetMetadata(metadataId string) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	_data, err := getData(db, metadataId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	_hydro, err := getHydro(db, metadataId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	_history, err := getHistory(db, metadataId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	data := &Metadata_struct{Data: _data, Hydro: _hydro, History: _history}
	return result.Result1(data), nil
}

// GET DATA
func getData(db *pqx.DB, metadataId string) (interface{}, error) {

	_result, err := getDataResult(db, metadataId)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := getDataResultScan(_result)

	return data, nil
}

func getDataResult(db *pqx.DB, metadataId string) (*sql.Rows, error) {
	var _result *sql.Rows
	var err error
	if metadataId == "" {
		_result, err = db.Query(sqlGetMetadata + sqlGetMetadataOrderBy)
	} else {
		countParam := 1
		_param := make([]interface{}, 0)
		arrayMetadataId := strings.Split(metadataId, ",")
		_sql := sqlGetMetadata + " AND m.id IN ("

		for i, v := range arrayMetadataId {
			_sql += "$" + strconv.Itoa(countParam)
			if i < len(arrayMetadataId)-1 {
				_sql += ","
			} else {
				_sql += ") "
			}
			countParam++
			_param = append(_param, v)
		}
		_result, err = db.Query(_sql+sqlGetMetadataOrderBy, _param...)
	}
	return _result, err
}

func getDataResultScan(_result *sql.Rows) interface{} {
	columns, _ := _result.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	data := make([]interface{}, 0)

	var (
		vals                 map[string]interface{}
		metadataagency_name  map[string]string
		metadataservice_name map[string]string
		metadata_tag         map[string]string
		metadata_description map[string]string
		metadata_frequency   []string

		id int64 = 0
	)

	for _result.Next() {
		vals = make(map[string]interface{})
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}

		_result.Scan(valuePtrs...)

		//Loop by Column
		for i, col := range columns {
			switch v := values[i].(type) {
			case time.Time:
				vals[col] = v.Format("2006-01-02")
			case nil:
				vals[col] = ""
			default:
				vals[col] = v
			}
		}

		if id != vals["id"].(int64) {
			if id != 0 {
				row := make(map[string]interface{})
				for k, v := range vals {
					row[k] = v
				}
				row["metadataagency_name"] = metadataagency_name
				row["metadataservice_name"] = metadataservice_name
				row["metadata_tag"] = metadata_tag
				row["metadata_description"] = metadata_description
				row["metadata_datafrequency"] = getMetadataFrequencyString(metadata_frequency)
				delete(row, "code")

				data = append(data, row)
			}
			id = vals["id"].(int64)
			metadataagency_name = make(map[string]string)
			metadataservice_name = make(map[string]string)
			metadata_tag = make(map[string]string)
			metadata_description = make(map[string]string)
			metadata_frequency = make([]string, 0)
		}
		metadataagency_name[vals["code"].(string)] = vals["metadataagency_name"].(string)
		metadataservice_name[vals["code"].(string)] = vals["metadataservice_name"].(string)
		metadata_tag[vals["code"].(string)] = vals["metadata_tag"].(string)
		metadata_description[vals["code"].(string)] = vals["metadata_description"].(string)
		metadata_frequency = append(metadata_frequency, vals["metadata_datafrequency"].(string))
	}
	if id != 0 {
		row := make(map[string]interface{})
		for k, v := range vals {
			row[k] = v
		}
		row["metadataagency_name"] = metadataagency_name
		row["metadataservice_name"] = metadataservice_name
		row["metadata_tag"] = metadata_tag
		row["metadata_description"] = metadata_description
		row["metadata_datafrequency"] = getMetadataFrequencyString(metadata_frequency)
		delete(row, "code")

		data = append(data, row)
	}
	return data
}

// GET HYDRO
func getHydro(db *pqx.DB, metadataId string) ([]int64, error) {
	_data := make([]int64, 0)

	_result, err := db.Query(sqlGetHydroMetadata, metadataId)
	if err != nil {
		return nil, err
	}

	for _result.Next() {
		var _id int64
		_result.Scan(&_id)

		_data = append(_data, _id)
	}

	return _data, nil
}

// GET HISTORY
func getHistory(db *pqx.DB, metadataId string) ([]*History_struct, error) {
	_data := make([]*History_struct, 0)

	_result, err := db.Query(sqlGetHistoryMetadata, metadataId)
	if err != nil {
		return nil, err
	}

	for _result.Next() {
		var (
			_datetime    sql.NullString
			_description sql.NullString
			_createdBy   sql.NullString
		)
		err = _result.Scan(&_datetime, &_description, &_createdBy)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}

		history := &History_struct{}
		history.HistoryDatetime = _datetime.String
		history.HistoryDescription = _description.String
		history.CreatedBy = _createdBy.String
		_data = append(_data, history)
	}

	return _data, nil
}

// GET METADATA FREQUENCY STRING
func getMetadataFrequencyString(arrFrequency []string) string {
	strFreq := ""
	arrResult := []string{}
	for _, freq := range arrFrequency {
		if strFreq != freq {
			arrResult = append(arrResult, freq)
			strFreq = freq
		}
	}
	return strings.Join(arrResult, ", ")
}
