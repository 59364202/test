package metadata

import (
	"database/sql"
	//	model_language "haii.or.th/api/server/model/language"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

var sql_ = `SELECT m.id, m.subcategory_id, m.agency_id, m.dataunit_id, m.dataformat_id
			     , m.servicemethod_id, m.connection_format, m.metadata_method, m.metadata_chanel, m.metadata_datafrequency
			     , m.metadata_convertfrequency, m.metadata_contact, m.metadata_agenystoredate, m.metadata_startdatadate, m.scriptname
			FROM metadata m
			LEFT JOIN metadata_translate mt ON m.id = mt.metadata_id, `

var sqlGetMatedataTable = " SELECT m.id , m.subcategory_id , m.agency_id , mt.metadataservice_name , l.code , mt.metadataagency_name " +
	" FROM metadata m " +
	" LEFT JOIN metadata_translate mt ON m.id = mt.metadata_id AND mt.deleted_by IS NULL " +
	" LEFT JOIN language l ON l.id = mt.language_id AND l.deleted_by IS NULL "

//var sqlGetMatedataTableWhere = " WHERE m.deleted_by IS NULL m.subcategory_id IN ($1) AND m.agency_id IN ($2) AND "
var sqlGetMatedataTableWhere = " WHERE m.deleted_by IS NULL AND mt.deleted_by IS NULL AND l.deleted_by IS NULL"
var sqlGetMatedataTableOrderBy = " ORDER BY m.id , l.id "

func GetMatadataTable(subcategoryId, agencyId string) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		id         int64
		text       map[string]string
		agencyName map[string]string

		data     []*MetadataTable_struct
		metadata *MetadataTable_struct

		_id            int64
		_subcategoryId int64
		_agencyId      int64
		_text          sql.NullString
		_agencyName    sql.NullString
		_code          sql.NullString
	)

	var _sql = sqlGetMatedataTable + sqlGetMatedataTableWhere

	var arraySubCategoryId = strings.Split(subcategoryId, ",")
	var arrayAgencyId = strings.Split(agencyId, ",")
	var countParam int = 1
	var _param = make([]interface{}, 0)

	if arraySubCategoryId[0] != "" {
		_sql += " AND m.subcategory_id IN ( "
		for i, v := range arraySubCategoryId {
			_sql += "$" + strconv.Itoa(countParam)
			if i != len(arraySubCategoryId)-1 {
				_sql += ","
			} else {
				_sql += ") "
			}

			_param = append(_param, v)
			countParam++
		}
	}

	if arrayAgencyId[0] != "" {
		_sql += " AND m.agency_id IN ("
		for i, v := range arrayAgencyId {
			_sql += "$" + strconv.Itoa(countParam)
			if i < len(arrayAgencyId)-1 {
				_sql += ","
			} else {
				_sql += ") "
			}

			_param = append(_param, v)
			countParam++
		}
	}

	_result, err := db.Query(_sql+sqlGetMatedataTableOrderBy, _param...)
	if err != nil {
		return nil, errors.Repack(err)
	}

	for _result.Next() {
		err = _result.Scan(&_id, &_subcategoryId, &_agencyId, &_text, &_code, &_agencyName)
		if err != nil {
			return nil, errors.Repack(err)
		}

		if id != _id {
			if id != 0 {
				metadata.MetadataServiceName = text
				metadata.MetadataAgencyName = agencyName
				data = append(data, metadata)
			}
			id = _id

			metadata = &MetadataTable_struct{}
			metadata.Id = id
			metadata.AgencyId = _agencyId
			metadata.SubcategoryId = _subcategoryId

			text = make(map[string]string)
			agencyName = make(map[string]string)
		}
		if _text.Valid {
			text[_code.String] = _text.String
		}
		if _agencyName.Valid {
			agencyName[_code.String] = _agencyName.String
		}
	}

	if id != 0 {
		metadata.MetadataServiceName = text
		metadata.MetadataAgencyName = agencyName
		data = append(data, metadata)
	}

	return result.Result1(data), nil
}
