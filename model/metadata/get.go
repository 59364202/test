// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for public.metadata table. This table store metadata.
package metadata

import (
	"database/sql"
	"encoding/json"
	//	logx "log"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"

	model "haii.or.th/api/server/model"
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_hydroinfo "haii.or.th/api/thaiwater30/model/hydroinfo"
	model_category "haii.or.th/api/thaiwater30/model/lt_category"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_subcategory "haii.or.th/api/thaiwater30/model/lt_subcategory"
	model_metadata_frequency "haii.or.th/api/thaiwater30/model/metadata_frequency"
	model_metadata_history "haii.or.th/api/thaiwater30/model/metadata_history"
	model_metadata_hydroinfo "haii.or.th/api/thaiwater30/model/metadata_hydroinfo"
	model_metadata_servicemethod "haii.or.th/api/thaiwater30/model/metadata_servicemethod"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/b64"
	"haii.or.th/api/thaiwater30/util/sqltime"
	"haii.or.th/api/thaiwater30/util/validdata"
)

//	get shopping table
//	Parameters:
//		p
//			Param_Metadata
//	Return:
//		Array Struct_Metadata
func GetMetadataShoppingTable(p *Param_Metadata) ([]*Struct_Metadata, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		data     []*Struct_Metadata
		metadata *Struct_Metadata

		_id                        int64
		_metadataservice_name      sql.NullString
		_metadata_description      sql.NullString
		_metadata_convertfrequency sql.NullString
		_sucategory_id             int64
		_sucategory_name           sql.NullString
		_agency_id                 int64
		_agency_name               sql.NullString
		_jsontable                 sql.NullString
		_table                     sql.NullString
		_column                    sql.NullString

		//		_fromdate time.Time
		//		_todate   time.Time
	)
	strSQL := SQL_selectMetadata
	strSQL_WHERE := ""
	if p.Data_name != "" {
		strSQL_WHERE += " AND m.metadataservice_name->>'th' LIKE '%" + p.Data_name + "%' "
	}
	if p.Category != 0 {
		strSQL_WHERE += " AND lc.id = " + strconv.FormatInt(p.Category, 10) + " "
	}
	if p.Subcategory != 0 {
		strSQL_WHERE += " AND ls.id = " + strconv.FormatInt(p.Subcategory, 10) + " "
	}
	if p.Ministry != 0 {
		strSQL_WHERE += " AND lm.id = " + strconv.FormatInt(p.Ministry, 10) + " "
	}
	if p.Department != 0 {
		strSQL_WHERE += " AND ld.id = " + strconv.FormatInt(p.Department, 10) + " "
	}
	if p.Agency_id != 0 {
		strSQL_WHERE += " AND a.id = " + strconv.FormatInt(p.Agency_id, 10) + " "
	}
	strSQL += strSQL_WHERE
	//	if p.Data_name != "" || p.Category != 0 || p.Subcategory != 0 || p.Ministry != 0 || p.Department != 0 {
	//		strSQL += "  " + strSQL_WHERE
	//	}

	fmt.Println("-- hello --",SQL_selectMetadata_Head + strSQL + SQL_selectMetadata_Orderby + SQL_selectMetadata_Foot)

	row, err := db.Query(SQL_selectMetadata_Head + strSQL + SQL_selectMetadata_Orderby + SQL_selectMetadata_Foot)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_metadataservice_name, &_metadata_description, &_metadata_convertfrequency, &_sucategory_id, &_sucategory_name, &_agency_id, &_agency_name,
			&_jsontable, &_table, &_column)
		if err != nil {
			return nil, err
		}

		if !_metadataservice_name.Valid || _metadataservice_name.String == "" {
			_metadataservice_name.String = "{}"
		}
		if !_metadata_description.Valid || _metadata_description.String == "" {
			_metadata_description.String = "{}"
		}
		if !_sucategory_name.Valid || _sucategory_name.String == "" {
			_sucategory_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		metadata = &Struct_Metadata{}
		data = append(data, metadata)

		metadata.Id = _id
		metadata.Table = _table.String
		metadata.Metadataservice_Name = json.RawMessage(_metadataservice_name.String)
		metadata.Metadata_Description = json.RawMessage(_metadata_description.String)
		metadata.Metadata_Convertfrequency = _metadata_convertfrequency.String
		metadata.Subcategory = &model_subcategory.Struct_subcategory{Id: _sucategory_id, Subcategory_name: json.RawMessage(_sucategory_name.String)}
		metadata.MetadataId, _ = model.GetCipher().EncryptText(strconv.FormatInt(_id, 10))

		metadata.Agency = &model_agency.Struct_Agency{}
		metadata.Agency.Id = _agency_id
		metadata.Agency.Agency_name = json.RawMessage(_agency_name.String)

		//		if !_table.Valid && _table.String == "{}" {
		//			continue
		//		}
		//
		//		if _column.Valid {
		//			err = db.QueryRow("SELECT min("+_column.String+"), max("+_column.String+") FROM "+_table.String).Scan(&_fromdate, &_todate)
		//		} else {
		//			err = db.QueryRow("SELECT min(updated_at), max(updated_at) FROM "+_table.String).Scan(&_fromdate, &_todate)
		//		}
		//		if err != nil {
		//			continue
		//		}
		//		metadata.FormDate = _fromdate.Format(setting.GetSystemSetting("setting.Default.DateFormat"))
		//		metadata.ToDate = _todate.Format(setting.GetSystemSetting("setting.Default.DateFormat"))

		//		_metadata_service, err := model_metadata_servicemethod.GetMetadataServicemethod(metadata.Id)
		//		if err == nil {
		//			metadata.Servicemethod = _metadata_service
		//		}
		_metadata_frequency, err := model_metadata_frequency.GetMetadataFrequency(metadata.Id)
		if err == nil {
			metadata.Frequency = _metadata_frequency
		}

		//		column := GetColumnDateTime(_table.String)
		//		if column == "" {
		//			continue
		//		}
		//
		//		err = db.QueryRow("SELECT min("+column+"), max("+column+") FROM "+_table.String).Scan(&_fromdate, &_todate)
		//		if err != nil {
		//			continue
		//		}
		//		metadata.FormDate = _fromdate.Format("2006-01-02")
		//		metadata.ToDate = _todate.Format("2006-01-02")
	}

	return data, nil
}

//	get shopping metadata detail
//	Parameters:
//		id
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_Metadata
func GetMetadataShoppingDetail(id int64) ([]*Struct_Metadata, error) {

	meta := &Struct_Metadata{}
	data := make([]*Struct_Metadata, 0)
	data = append(data, meta)

	meta.Id = id
	p, err := getAllProvinceFromMetaId(id)
	if err != nil || p == nil {
		meta.Province = nil
	} else {
		meta.Province = p
	}

	b, err := model_basin.GetAllBasinFromMetaId(id)
	if err != nil || b == nil {
		meta.Basin = nil
	} else {
		meta.Basin = b
	}

	_metadata_service, err := model_metadata_servicemethod.GetMetadataServicemethod(id)
	if err != nil {
		meta.Servicemethod = nil
	} else {
		meta.Servicemethod = _metadata_service
	}
	_metadata_frequency, err := model_metadata_frequency.GetMetadataFrequency(id)
	if err != nil {
		meta.Frequency = nil
	} else {
		meta.Frequency = _metadata_frequency
	}
	getFromDateToDate(meta)

	return data, nil
}

//	หา min date, max date ของ metadata
//	Parameters:
//		meta
//			Struct_Metadata
func getFromDateToDate(meta *Struct_Metadata) {
	db, err := pqx.Open()
	if err != nil {
		return
	}
	var (
		metadata_fromdate   sql.NullString
		metadata_todate     sql.NullString
		metadata_last_check sql.NullString
	)
	// หา min, max จาก metadata
	itf := make([]interface{}, 0)
	itf = append(itf, meta.Id)
	strSql := "SELECT data_min_date, data_max_date, data_last_check FROM metadata WHERE id = $1"
	err = db.QueryRow(strSql, itf...).Scan(&metadata_fromdate, &metadata_todate, &metadata_last_check)
	if err != nil {
		log.Locationf("metadata id : %v , SQL_genSqlSelectMaxMinDateFromTable err : %s , %s , %v", meta.Id, err, strSql, itf)
		return
	}

	if metadata_fromdate.Valid {
		meta.FormDate = pqx.NullStringToTime(metadata_fromdate).Format(setting.GetSystemSetting("setting.Default.DateFormat"))

	}
	metadata_todate_date := pqx.NullStringToTime(metadata_todate)
	if metadata_todate.Valid {
		meta.ToDate = metadata_todate_date.Format(setting.GetSystemSetting("setting.Default.DateFormat"))
	}

	//	// วันที่ล่าสุดของข้อมูล ไม่ใช่วันปัจจุบัน ให้ไปเช็ค วันที่ล่าสุด จากข้อมูลใหม่
	//	if metadata_todate_date.Format("20060102") != time.Now().Format("20160102") {
	//		var (
	//			_table           sql.NullString
	//			_column          sql.NullString
	//			_agency_id       int64
	//			_convert_setting sql.NullString
	//
	//			_fromdate sql.NullString
	//			_todate   sql.NullString
	//		)
	//
	//		// เอา max ไปเช็คเพื่อดูว่า max เปลี่ยนรึป่าว
	//		err = db.QueryRow(SQL_selectPartitionFieldFromMetadata, meta.Id).Scan(&_table, &_column, &_agency_id, &_convert_setting)
	//		if !_table.Valid && _table.String == "{}" || err != nil {
	//			return
	//		}
	//		strSql, itf = SQL_genSqlSelectMaxMinDateFromTable(_table.String, _agency_id, _convert_setting.String, metadata_last_check.String)
	//		if strSql == "" {
	//			return
	//		}
	//
	//		err = db.QueryRow(strSql, itf...).Scan(&_fromdate, &_todate)
	//		if err != nil {
	//			log.Log(err.Error())
	//			return
	//		}
	//		log.Log(strSql, itf)
	//		if _todate.Valid {
	//			// มีวันที่ข้อมูลใหม่กว่ามา ต้องไปอัพเดทใน metadata ด้วย
	//			meta.ToDate = pqx.NullStringToTime(_todate).Format(setting.GetSystemSetting("setting.Default.DateFormat"))
	//			strSql = "UPDATE metadata SET data_last_check = NOW(), data_max_date = $1 WHERE id = $2"
	//			db.Exec(strSql, _todate.String, meta.Id)
	//		}
	//	}
}

//	หา min date, max date ของ metadata จาก config
//	Parameters:
//		meta
//			Struct_Metadata
//		metadata_last_check
//			วันที่เช็ค ล่าสุด
//		datasetId
//			dataset id
//		additionalDataset
//			dataset id ที่เกี่ยวข้อง ex. 1,2,3
//		find
//			ค่าที่ต้องการหา {min,max,} ไม่ใส่ถือว่าหา min, max
func findFromDateToDate(meta *Struct_Metadata, metadata_last_check string, datasetId int64, additionalDataset string, find string) error {
	db, err := pqx.Open()
	if err != nil {
		return err
	}
	var (
		_table           sql.NullString
		_column          sql.NullString
		_agency_id       int64
		_convert_setting sql.NullString

		_fromdate sql.NullString
		_todate   sql.NullString
	)
	sDatasetId := datatype.MakeString(datasetId)

	err = db.QueryRow(SQL_selectPartitionFieldFromMetadata, meta.Id).Scan(&_table, &_column, &_agency_id, &_convert_setting)
	if !_table.Valid && _table.String == "{}" || err != nil {
		return nil
	}

	switch find {
	case "max":
		if _todate.Valid {
			metadata_last_check = _todate.String[0:10]
		}
	case "min":
		if _fromdate.Valid {
			metadata_last_check = _fromdate.String[0:10]
		}
	}

	strSql, itf := SQL_genSqlSelectMaxMinDateFromTable(_table.String, _agency_id, _convert_setting.String, metadata_last_check, sDatasetId, additionalDataset, find, _todate, _fromdate)
	if strSql == "" {
		return nil
	}

	//	return nil
	switch find {
	case "max":
		err = db.QueryRow(strSql, itf...).Scan(&_todate)
	case "min":
		err = db.QueryRow(strSql, itf...).Scan(&_fromdate)
	default:
		err = db.QueryRow(strSql, itf...).Scan(&_fromdate, &_todate)
	}

	if err != nil {
		log.Locationf("metadata id : %v , SQL_genSqlSelectMaxMinDateFromTable err : %s , %s , %v", meta.Id, err, strSql, itf)
		return err
	}
	if _fromdate.Valid {
		meta.FormDate = pqx.NullStringToTime(_fromdate).Format(setting.GetSystemSetting("setting.Default.DateFormat"))
		//		meta.FormDate = pqx.NullStringToTime(_fromdate).Format("2006-01-02")
		//		meta.ToDate = pqx.NullStringToTime(_todate).Format("2006-01-02")
	}
	if _todate.Valid {
		meta.ToDate = pqx.NullStringToTime(_todate).Format(setting.GetSystemSetting("setting.Default.DateFormat"))
	}
	return nil
}

//	get province ทั้งหมดตาม metadata id
//	Parameters:
//		meta_id
//			รหัสบัญชีข้อมูล
//	Return:
//		Array Struct_Province
func getAllProvinceFromMetaId(meta_id int64) ([]*model_lt_geocode.Struct_Province, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	//	หา lookup table จาก dataset
	lu, err := getTableFromLookupTable(meta_id)
	if err != nil {
		return nil, err
	}
	if lu.Table_name.String == "" {
		return nil, rest.NewError(500, "err table_name", err)
	}

	var (
		data     []*model_lt_geocode.Struct_Province
		province *model_lt_geocode.Struct_Province

		_province_code string
		_province_name sql.NullString
	)
	strSql := ""
	if lu.Table_name.String != "" && lu.Table_name.String != "lt_geocode" {
		strSql = lu.Table_name.String + " WHERE agency_id = " + lu.Agency_id.String
	}
	row, err := db.Query(SQL_selectAllProvinceFromTable + strSql + SQL_selectAllProvinceFromTable_end)
	if err != nil {
		return nil, rest.NewError(500, "err query", err)
	}
	for row.Next() {
		err = row.Scan(&_province_code, &_province_name)
		if err != nil {
			return nil, rest.NewError(500, "err scan", err)
		}

		if !_province_name.Valid || _province_name.String == "" {
			_province_name.String = "{}"
		}

		province = &model_lt_geocode.Struct_Province{}
		province.Province_code = _province_code
		province.Province_name = json.RawMessage(_province_name.String)

		data = append(data, province)
	}
	return data, nil
}

//	get table from LookupTable field by metadata id
//	Parameters:
//		meta_id
//			รหัสบัญชีข้อมูล
//	Return:
//		Strct_Data
func getTableFromLookupTable(meta_id int64) (*Strct_Data, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		_table     sql.NullString
		_agency_id sql.NullString
	)

	row, err := db.Query(SQL_selectTableFromDatainportDataset, meta_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err := row.Scan(&_table, &_agency_id)
		if err != nil {
			return nil, err
		}
	}
	data := &Strct_Data{}
	data.Table_name = _table
	data.Agency_id = _agency_id

	return data, nil
}

//	get สรุปบัญชีข้อมูล กรุ๊ปตาม หน่วยงาน
//	Return:
//		Array Struct_MetadataSummary
func GetSummaryMetadataGroupByAgency() ([]*Struct_MetadataSummary, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Setting Variables
	var (
		data        []*Struct_MetadataSummary
		objMetadata *Struct_MetadataSummary

		_agency_id           sql.NullInt64
		_agency_name         sql.NullString
		_agency_shortname    sql.NullString
		_metadata_sum_total  sql.NullInt64
		_metadata_sum_import sql.NullInt64

		_result *sql.Rows
	)

	//Query
	_result, err = db.Query(SQL_selectSummaryMetadataByAgency)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for _result.Next() {
		err = _result.Scan(&_agency_id, &_agency_shortname, &_agency_name, &_metadata_sum_import, &_metadata_sum_total)
		if err != nil {
			return nil, err
		}

		if _agency_shortname.String == "" || !_agency_shortname.Valid {
			_agency_shortname.String = "{}"
		}
		if _agency_name.String == "" || !_agency_name.Valid {
			_agency_name.String = "{}"
		}

		objMetadata = &Struct_MetadataSummary{}
		objMetadata.Summary_Import = _metadata_sum_import.Int64
		objMetadata.Summary_Total = _metadata_sum_total.Int64

		objMetadata.Agency = &model_agency.Struct_Agency{}
		objMetadata.Agency.Id = _agency_id.Int64
		objMetadata.Agency.Agency_name = json.RawMessage(_agency_name.String)
		objMetadata.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)

		data = append(data, objMetadata)
	}

	return data, nil
}

//	get สรุปนำเข้าบัญชีข้อมูล กรุ๊ปตาม หน่วยงาน
//	Parameters:
//		agencyID
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_MetadataSummary
func getSummaryMetadataImportedByAgency(agencyID int64) ([]*Struct_MetadataSummary, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Setting Variables
	var (
		data        []*Struct_MetadataSummary
		objMetadata *Struct_MetadataSummary

		_category_id    sql.NullInt64
		_category_name  sql.NullString
		_metadata_count sql.NullInt64

		_result *sql.Rows
	)

	//Query
	//log.Printf(SQL_selectSummaryMetadataImportedByAgency + SQL_selectSummaryMetadataImportedByAgency_OrderBy, agencyID)
	_result, err = db.Query(SQL_selectSummaryMetadataImportedByAgency+SQL_selectSummaryMetadataImportedByAgency_OrderBy, agencyID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for _result.Next() {
		err = _result.Scan(&_category_id, &_category_name, &_metadata_count)
		if err != nil {
			return nil, err
		}

		if _category_name.String == "" || !_category_name.Valid {
			_category_name.String = "{}"
		}

		objMetadata = &Struct_MetadataSummary{}
		objMetadata.Count_Metadata = _metadata_count.Int64

		objMetadata.Subcategory = &model_subcategory.Struct_subcategory{}
		objMetadata.Subcategory.Category = &model_category.Struct_category{}
		objMetadata.Subcategory.Category.Id = _category_id.Int64
		objMetadata.Subcategory.Category.Category_name = json.RawMessage(_category_name.String)

		data = append(data, objMetadata)
	}

	return data, nil
}

//	get สรุปนำเข้าบัญชีข้อมูล กรุ๊ปตาม หมวดหมู่หลัก
//	Parameters:
//		categoryID
//			รหัสหมวดหมู่หลัก
//	Return:
//		Array Struct_MetadataSummary
func getSummaryMetadataImportedByCategory(categoryID int64) ([]*Struct_MetadataSummary, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Setting Variables
	var (
		data        []*Struct_MetadataSummary
		objMetadata *Struct_MetadataSummary

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString
		_metadata_count   sql.NullInt64

		_result *sql.Rows
	)

	//Query
	_result, err = db.Query(SQL_selectSummaryMetadataImportedByCategory+SQL_selectSummaryMetadataImportedByCategory_OrderBy, categoryID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for _result.Next() {
		err = _result.Scan(&_agency_id, &_agency_shortname, &_agency_name, &_metadata_count)
		if err != nil {
			return nil, err
		}

		if _agency_shortname.String == "" || !_agency_shortname.Valid {
			_agency_shortname.String = "{}"
		}
		if _agency_name.String == "" || !_agency_name.Valid {
			_agency_name.String = "{}"
		}

		objMetadata = &Struct_MetadataSummary{}
		objMetadata.Count_Metadata = _metadata_count.Int64

		objMetadata.Agency = &model_agency.Struct_Agency{}
		objMetadata.Agency.Id = _agency_id.Int64
		objMetadata.Agency.Agency_name = json.RawMessage(_agency_name.String)
		objMetadata.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)

		data = append(data, objMetadata)
	}

	return data, nil
}

//	get สรุปนำเข้าบัญชีข้อมูล กรุ๊ปตาม หมวดหมู่หลัก
//	Parameters:
//		categoryID
//			รหัสหมวดหมู่หลัก
//	Return:
//		Array Struct_MetadataSummary
func getMetadataImportedByAgencyCategory(agencyID int64, categoryID int64) ([]*Struct_Metadata, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Setting Variables
	var (
		data        []*Struct_Metadata
		objMetadata *Struct_Metadata

		_id                   sql.NullInt64
		_metadata_servicename sql.NullString
		_metadata_agencyname  sql.NullString
		_total_record         sql.NullInt64
		_last_import_date     sql.NullString

		_result *sql.Rows
	)

	//Query
	//log.Println(SQL_selectMetadataImportedByAgencyCategory + SQL_selectMetadataImportedByAgencyCategory_OrderBy)
	_result, err = db.Query(SQL_selectMetadataImportedByAgencyCategory+SQL_selectMetadataImportedByAgencyCategory_OrderBy, agencyID, categoryID)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	for _result.Next() {
		err = _result.Scan(&_id, &_metadata_servicename, &_metadata_agencyname, &_total_record, &_last_import_date)
		if err != nil {
			return nil, err
		}

		if _metadata_servicename.String == "" || !_metadata_servicename.Valid {
			_metadata_servicename.String = "{}"
		}
		if _metadata_agencyname.String == "" || !_metadata_agencyname.Valid {
			_metadata_agencyname.String = "{}"
		}

		var strLastImportDate string = ""
		if _last_import_date.String != "" {
			dtLastImportDate, err := time.Parse("2006-01-02T15:04:05Z07:00", _last_import_date.String)
			if err != nil {
				return nil, errors.Repack(err)
			}
			strLastImportDate = dtLastImportDate.Format(strDatetimeFormat)
		}

		objMetadata = &Struct_Metadata{}
		objMetadata.Id = _id.Int64
		objMetadata.Metadataservice_Name = json.RawMessage(_metadata_servicename.String)
		objMetadata.Metadataagency_Name = json.RawMessage(_metadata_agencyname.String)
		objMetadata.Total_Import_Record = _total_record.Int64
		objMetadata.Last_Import_Date = strLastImportDate

		data = append(data, objMetadata)
	}

	return data, nil
}

//	get สรุปนำเข้าบัญชีข้อมูล
//	Parameters:
//		param
//			Struct_Metadata_InputParam
//	Return:
//		Array Struct_MetadataSummary
func GetSummaryMetadataImported(param *Struct_Metadata_InputParam) ([]*Struct_MetadataSummary, error) {
	var intAgencyID int64 = 0
	var intCategoryID int64 = 0
	var err error

	if param.Agency_Id != "" {
		//Convert AgencyID type from string to int64
		intAgencyID, err = strconv.ParseInt(param.Agency_Id, 10, 64)
		if err != nil {
			return nil, errors.Repack(err)
		}
	}

	if param.Category_Id != "" {
		//Convert CategoryID type from string to int64
		intCategoryID, err = strconv.ParseInt(param.Category_Id, 10, 64)
		if err != nil {
			return nil, errors.Repack(err)
		}
	}

	if (intAgencyID == 0) && (intCategoryID == 0) {
		return nil, errors.New("'agency_id' and 'category_id' most not be null both.")
	}

	if intAgencyID != 0 {
		return getSummaryMetadataImportedByAgency(intAgencyID)
	} else {
		return getSummaryMetadataImportedByCategory(intCategoryID)
	}
}

//	get รายละเอียดบัญชีข้อมูลที่เชื่อมโยงของหน่วยงาน
//	Parameters:
//		param
//			Struct_Metadata_InputParam
//	Return:
//		Array Struct_Metadata
func GetMetadataImported(param *Struct_Metadata_InputParam) ([]*Struct_Metadata, error) {
	var intAgencyID int64 = 0
	var intCategoryID int64 = 0
	var err error

	if param.Agency_Id != "" {
		//Convert AgencyID type from string to int64
		intAgencyID, err = strconv.ParseInt(param.Agency_Id, 10, 64)
		if err != nil {
			return nil, errors.Repack(err)
		}
	}

	if param.Category_Id != "" {
		//Convert CategoryID type from string to int64
		intCategoryID, err = strconv.ParseInt(param.Category_Id, 10, 64)
		if err != nil {
			return nil, errors.Repack(err)
		}
	}

	return getMetadataImportedByAgencyCategory(intAgencyID, intCategoryID)
}

//	get metadata ตาม agency id, status
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//		metadata_status
//			รหัสสถานะ
//	Return:
//		Array Struct_M
func getMetadataByAgencyAndStatus(agency_id int64, metadata_status string) ([]*Struct_M, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	if agency_id < 1 || metadata_status == "" {
		return nil, errors.New("no param")
	}

	row, err := db.Query(SQL_selectMetadataByAgencyAndStatus, agency_id, metadata_status)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data := make([]*Struct_M, 0)
	var (
		object *Struct_M

		_id                     int64
		_metadata_name          sql.NullString
		_metadata_frequency     sql.NullString
		_connection_format      sql.NullString
		_dataimport_download_id sql.NullInt64
		_dataimport_dataset_id  sql.NullInt64
	)
	for row.Next() {
		err = row.Scan(&_id, &_metadata_name, &_metadata_frequency, &_connection_format, &_dataimport_download_id, &_dataimport_dataset_id)
		if err != nil {
			return nil, err
		}

		if !_metadata_name.Valid || _metadata_name.String == "" {
			_metadata_name.String = "{}"
		}

		object = &Struct_M{Id: _id}
		object.Metadata_Convertfrequency = _metadata_frequency.String
		object.Metadataservice_Name = json.RawMessage(_metadata_name.String)
		object.Connection_Format = _connection_format.String
		object.Dataimport_Download_Id = validdata.ValidData(_dataimport_download_id.Valid, _dataimport_download_id.Int64)
		object.Dataimport_Dataset_Id = validdata.ValidData(_dataimport_dataset_id.Valid, _dataimport_dataset_id.Int64)

		data = append(data, object)
	}

	return data, nil
}

//	get metadata ตาม agency id ที่สถานะเป็น connecct
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_M
func GetMetadataStatusConnect(agency_id int64) ([]*Struct_M, error) {
	return getMetadataByAgencyAndStatus(agency_id, MetadataStatus_Connect)
}

//	get metadata ตาม agency id ที่สถานะเป็น wait for update
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_M
func GetMetadataStatusWaitUpdate(agency_id int64) ([]*Struct_M, error) {
	return getMetadataByAgencyAndStatus(agency_id, MetadataStatus_WaitUpdate)
}

//	get metadata ตาม agency id ที่สถานะเป็น wait for connect
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_M
func GetMetadataStatusWaitConnect(agency_id int64) ([]*Struct_M, error) {
	return getMetadataByAgencyAndStatus(agency_id, MetadataStatus_WaitConnect)
}

//	get metadata ตาม agency id ที่สถานะเป็น cancel
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//	Return:
//		Array Struct_M
func GetMetadataStatusCancel(agency_id int64) ([]*Struct_M, error) {
	return getMetadataByAgencyAndStatus(agency_id, MetadataStatus_Cancel)
}

//	get ข้อมูลที่นำเข้าล่าสุดของ metadata
//	Parameters:
//		metadata_id
//			รหัสบัญชีข้อมูล
//		media_url
//			ลิงค์ media
//	Return:
//		Struct_MetadataImportByAgency
func GetLastMetadataImportByAgency(metadata_id int64, media_url string) (*Struct_MetadataImportByAgency, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	if metadata_id < 1 {
		return nil, errors.New("no param")
	}
	var (
		_metadata_id       int64
		_agency_id         sql.NullInt64
		_connection_format sql.NullString
		_table_name        sql.NullString
		_lookup_table      sql.NullString
		_fields            sql.NullString
		_itfFields         []map[string]interface{}
		_mapField          map[string]interface{}
		isMediaTypeId      bool = false

		_media_type_id interface{}
		media_type_id  string
	)
	row, err := db.Query(SQL_SelectMetadataImportTableByAgency, metadata_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for row.Next() {
		err = row.Scan(&_metadata_id, &_agency_id, &_connection_format, &_table_name, &_lookup_table, &_fields)
		if err != nil {
			return nil, err
		}
	}

	if !_agency_id.Valid || !_connection_format.Valid || !_table_name.Valid {
		return nil, errors.New(" this metadata no dataset_id ")
	}

	if IsMedia(_table_name.String) && _fields.Valid && _fields.String != "" && _fields.String != "{}" {
		// เป็น media หา media_type_id
		err = json.Unmarshal([]byte(_fields.String), &_itfFields)
		if err != nil {
			return nil, err
		}
	L:
		for _, iv := range _itfFields {
			_mapField = iv
			if _mapField["name"].(string) == "media_type_id" {
				isMediaTypeId = true
			}
			if isMediaTypeId {
				break L
			}
		}
		_media_type_id = _mapField["transform_params"]
		if str, ok := _media_type_id.(string); ok {
			media_type_id = str
		} else {
			media_type_id = strconv.Itoa(_media_type_id.(int))
		}
	}
	//	log.Println(_mapField)
	str_agency_id := strconv.FormatInt(_agency_id.Int64, 10)
	// gen sql
	_connection_format.String = strings.ToUpper(strings.TrimSpace(_connection_format.String)) // เปลี่ยน ให้เป็นตัวใหญ่

	q := SQL_GenSQLSelectLatestImport(_connection_format.String, _table_name.String, str_agency_id, media_type_id)

	row, err = db.Query(q)
	if err != nil {
		return nil, err
	}

	data := &Struct_MetadataImportByAgency{}
	if IsMedia(_table_name.String) {
		rs, err := ScanData_Media(row, media_url)
		if err != nil {
			return nil, err
		}
		data.Img = rs
	} else {
		rs, err := Scan_GetLastMetadataImportByAgency(row)
		if err != nil {
			return nil, err
		}

		if _connection_format.String == "ONLINE" {
			data.Weather = rs
		} else {
			data.Table = rs
		}
	}
	data.TableName = _table_name.String

	return data, nil
}

//	scan sql result ลงใน Struct_Data_Media
//	Parameters:
//		row
//			sql.Rows
//		media_url
//			ลิงค์ media
//	Return:
//		Array Struct_Data_Media
func ScanData_Media(row *sql.Rows, media_url string) ([]*Struct_Data_Media, error) {
	columns, _ := row.Columns()
	count := len(columns)
	values := make([]interface{}, count) // สำหรับ scan dynamic column
	valuePtrs := make([]interface{}, count)
	//	final_result := make([]map[string]interface{}, 0)

	data := make([]*Struct_Data_Media, 0)
	var obj *Struct_Data_Media

	for row.Next() {
		// scan data ลง array
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		row.Scan(valuePtrs...)
		tmp_struct := map[string]interface{}{}

		var (
			_temp_media_path string
			_temp_filename   string
		)

		for i, col := range columns {
			val := values[i]

			if col == "media_path" {
				_temp_media_path = val.(string)
			}
			if col == "filename" {
				_temp_filename = val.(string)
			}

			if b, ok := val.(string); ok { // type string
				tmp_struct[col] = b
			} else if float, ok := val.(float64); ok { // type float64
				tmp_struct[col] = float
			} else {
				tmp_struct[col] = val
			}
			if _temp_media_path != "" && _temp_filename != "" { // ได้ media_path, filename แล้วทำ EncryptText url
				tmp_struct["path"] = tmp_struct["media_path"] // เก็บ path ไว้สำหรับไปดึงไฟล์ แต่ไม่ส่งออกเป็น json
				eid, _ := model.GetCipher().EncryptText(filepath.Join(_temp_media_path, _temp_filename))
				tmp_struct["media_path"] = media_url + "?file=" + eid
				_temp_media_path = "" // เคลีย _temp_media_path, _temp_filename เพื่อที่จะได้ไม่ต้อง EncryptText url ซ้ำ
				_temp_filename = ""
			}
		}
		//		dt, _ := time.Parse(time.RFC3339, tmp_struct["media_datetime"].(string))
		// ใส่ลง struct
		obj = &Struct_Data_Media{
			Agency_id: tmp_struct["agency_id"],
			Filename:  tmp_struct["filename"],
			//			Media_Datetime: dt.Format("2006-01-02 15:04"),
			Media_Datetime: tmp_struct["media_datetime"],
			Media_Desc:     tmp_struct["media_desc"],
			Media_Path:     tmp_struct["media_path"],
			Media_Type_id:  tmp_struct["media_type_id"],
			Refer_Source:   tmp_struct["refer_source"],
			Path:           tmp_struct["path"].(string),
		}

		data = append(data, obj)
	}
	return data, nil
}

//	get ข้อมูล media ที่นำเข้าล่าสุด ตาม table name, agency id, media type id
//	Parameters:
//		_table_name
//			ชื่อตาราง
//		_agency_id
//			รหัสหน่วยงาน
//		media_type_id
//			รหัสประเภทข้อมูลสื่อ
//	Return:
//		Array Struct_MetadataImportByAgency_Img
func GetLastMetadataImportByAgency_Media(_table_name, _agency_id string, media_type_id string) ([]*Struct_MetadataImportByAgency_Img, error) {
	data := []*Struct_MetadataImportByAgency_Img{}
	db, _ := pqx.Open()
	var (
		obj       *Struct_MetadataImportByAgency_Img
		_path     sql.NullString
		_filename sql.NullString
	)

	strSql := SQL_GenSelectLastMetadataImportMedia(_table_name, _agency_id, media_type_id)

	row, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		err = row.Scan(&_path, &_filename)
		if err != nil {
			return nil, err
		}
		encryptText, _ := b64.EncryptText(_path.String + "," + _filename.String)
		obj = &Struct_MetadataImportByAgency_Img{Name: _filename.String, Path: encryptText, FilePath: _path.String}
		data = append(data, obj)
	}

	return data, nil
}

//func GetLastMetadataImportByAgency_Online(_agency_id, _table_name, _lookup_table string) (*Struct_MetadataImportByAgency, error) {
//
//	if _lookup_table == "" || _lookup_table == "lt_geocode" {
//		return GetLastMetadataImportByAgency_Offline(_agency_id, _table_name, _lookup_table)
//	}
//
//	row, err := query(SQL_GenSQLSelectFromMetadata_Online(_table_name, _lookup_table, _agency_id))
//	if err != nil {
//		return nil, err
//	}
//	rs, err := Scan_GetLastMetadataImportByAgency(row)
//	data := &Struct_MetadataImportByAgency{Weather: rs, TableName: _table_name}
//	return data, err
//}
//
//func GetLastMetadataImportByAgency_Offline(_agency_id string, _table_name, _lookup_table string) (*Struct_MetadataImportByAgency, error) {
//	row, err := query(SQL_GenSQLSelectFromMetadata_Offline(_table_name, _lookup_table, _agency_id))
//	if err != nil {
//		return nil, err
//	}
//	rs, err := Scan_GetLastMetadataImportByAgency(row)
//	data := &Struct_MetadataImportByAgency{Table: rs}
//	return data, err
//}
//
//func query(strSql string) (*sql.Rows, error) {
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, err
//	}
//
//	row, err := db.Query(strSql)
//
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//	return row, nil
//}

//	scan sql result ที่ไม่ใช่ media ลงใน Struct_MetadataImportByAgency_Table
//	Parameters:
//		row
//			sql.Rows
//	Return:
//		Struct_MetadataImportByAgency_Table
func Scan_GetLastMetadataImportByAgency(row *sql.Rows) (*Struct_MetadataImportByAgency_Table, error) {
	rs := &Struct_MetadataImportByAgency_Table{}

	columns, _ := row.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	final_result := make([]map[string]interface{}, 0)
	final_array := make([][]interface{}, 0)

	for row.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		row.Scan(valuePtrs...)
		tmp_struct := map[string]interface{}{}
		tmp_array := make([]interface{}, 0)

		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				var js map[string]interface{}
				// check json
				if json.Unmarshal(b, &js) == nil {
					tmp_struct[col] = js
					tmp_array = append(tmp_array, js)
				} else {
					// string
					tmp_struct[col] = string(b)
					tmp_array = append(tmp_array, string(b))
				}
			} else if float, ok := val.(float64); ok {
				// float64
				tmp_struct[col] = float
				tmp_array = append(tmp_array, float)
			} else {
				tmp_struct[col] = val
				tmp_array = append(tmp_array, val)
			}
		}
		final_result = append(final_result, tmp_struct)
		final_array = append(final_array, tmp_array)
	}
	rs.Columns = columns
	rs.Data = final_result
	rs.DataArray = final_array
	return rs, nil
}

//	get log การที่ขอใช้บริการข้อมูล จำแนกรายหน่วยงาน
//	Parameters:
//		agency_id
//			รหัสหน่วยงานของผู้ใช้
//		m_agency
//			รหัสหน่วยงานของบัญชีข้อมูล
//	Return:
//		Array Struct_MetadataStatus
func GetOrderResult(agency_id, m_agency int64) ([]*Struct_MetadataStatus, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var obj *Struct_MetadataStatus
	//	query
	row, err := db.Query(SQL_SelectMetadataStatusByAgency, agency_id, m_agency)
	if err != nil {
		return nil, err
	}
	data := make([]*Struct_MetadataStatus, 0)
	var format string = setting.GetSystemSetting("setting.Default.DatetimeFormat")
	for row.Next() {
		var (
			_id                 int64
			_user_name          sql.NullString
			_metadata_name      sql.NullString
			_servicemethod_name sql.NullString
			_create_date        sqltime.NullTime
			_result_date        sqltime.NullTime
			_result             sql.NullString
		)
		if err = row.Scan(&_id, &_user_name, &_metadata_name, &_servicemethod_name, &_create_date, &_result_date, &_result); err != nil {
			return nil, err
		}

		if !_metadata_name.Valid || _metadata_name.String == "" {
			_metadata_name.String = "{}"
		}
		if !_servicemethod_name.Valid || _servicemethod_name.String == "" {
			_servicemethod_name.String = "{}"
		}

		obj = &Struct_MetadataStatus{Id: _id, UserName: _user_name.String}
		obj.MetadataName = json.RawMessage(_metadata_name.String)
		obj.ServicemethodName = json.RawMessage(_servicemethod_name.String)
		obj.CreateDate = _create_date.Time.Format(format)
		if _result_date.Valid {
			obj.ResultDate = _result_date.Time.Format(format)
		}
		obj.Result = _result.String

		data = append(data, obj)
	}
	return data, nil
}

//	get metadata
//	Parameters:
//		arrMetadataID
//			array รหัสหน่วยบัญชีข้อมูล
//	Return:
//		Array Struct_Metadata_Data
func GetMetadata(arrMetadataID []string) ([]*Struct_Metadata_Data, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_Metadata_Data
		objMetadata *Struct_Metadata_Data

		_id                        sql.NullInt64
		_subcategory_id            sql.NullInt64
		_agency_id                 sql.NullInt64
		_dataunit_id               sql.NullInt64
		_dataformat_id             sql.NullInt64
		_connection_format         sql.NullString
		_metadata_channel          sql.NullString
		_metadata_convertfrequency sql.NullString
		_metadata_contact          sql.NullString
		_metadata_agencystoredate  sqltime.NullTime
		_metadata_startdatadate    sqltime.NullTime
		_metadata_update_plan      sql.NullInt64
		_metadata_laws             sql.NullString
		_metadata_remark           sql.NullString
		_metadataagency_name       sql.NullString
		_metadataservice_name      sql.NullString
		_metadata_tag              sql.NullString
		_metadata_description      sql.NullString
		_metadata_status_id        sql.NullInt64
		_metadata_method_id        sql.NullInt64
		_category_id               sql.NullInt64
		_metadata_receive_date     sqltime.NullTime

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""

	//Check Filter agency_id
	if len(arrMetadataID) > 0 {
		if len(arrMetadataID) == 1 {
			strMetadataID, err := model.GetCipher().DecryptText(arrMetadataID[0])
			if err != nil {
				return nil, errors.New("Invalid Key")
			}
			arrParam = append(arrParam, strMetadataID)
			sqlCmdWhere += " AND m.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, strId := range arrMetadataID {
				strMetadataID, err := model.GetCipher().DecryptText(strId)
				if err != nil {
					return nil, errors.New("Invalid Key")
				}
				arrParam = append(arrParam, strMetadataID)
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND m.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DateFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	}

	//	log.Printf(sqlGetMetadata+sqlCmdWhere+sqlGetMetadataGroupby, arrParam...)
	_result, err = db.Query(sqlGetMetadata+sqlCmdWhere+sqlGetMetadataGroupby, arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_id, &_subcategory_id, &_agency_id, &_dataunit_id, &_dataformat_id,
			&_connection_format, &_metadata_channel, &_metadata_convertfrequency, &_metadata_contact,
			&_metadata_agencystoredate, &_metadata_startdatadate,
			&_metadata_update_plan, &_metadata_laws, &_metadata_remark,
			&_metadataagency_name, &_metadataservice_name, &_metadata_tag, &_metadata_description,
			&_metadata_status_id, &_metadata_method_id, &_category_id, &_metadata_receive_date)
		if err != nil {
			return nil, err
		}

		objServiceMethod, err := model_metadata_servicemethod.GetMetadataServicemethod(_id.Int64)
		if err != nil {
			return nil, err
		}

		objHydroInfo, err := model_metadata_hydroinfo.GetMetadataHydroinfo(_id.Int64)
		if err != nil {
			return nil, err
		}

		objFrequency, err := model_metadata_frequency.GetMetadataFrequency(_id.Int64)
		if err != nil {
			return nil, err
		}

		objHistory, err := model_metadata_history.GetMetadataHistory(_id.Int64)
		if err != nil {
			return nil, err
		}

		if !_metadataagency_name.Valid || _metadataagency_name.String == "" {
			_metadataagency_name.String = "{}"
		}
		if !_metadataservice_name.Valid || _metadataservice_name.String == "" {
			_metadataservice_name.String = "{}"
		}
		if !_metadata_tag.Valid || _metadata_tag.String == "" {
			_metadata_tag.String = "{}"
		}
		if !_metadata_description.Valid || _metadata_description.String == "" {
			_metadata_description.String = "{}"
		}

		objMetadata = &Struct_Metadata_Data{}
		objMetadata.Id = _id.Int64
		objMetadata.MetadataId, _ = model.GetCipher().EncryptText(strconv.FormatInt(_id.Int64, 10))
		objMetadata.SubcategoryId = _subcategory_id.Int64
		objMetadata.AgencyId = _agency_id.Int64
		objMetadata.DataunitId = _dataunit_id.Int64
		objMetadata.DataformatId = _dataformat_id.Int64
		objMetadata.ConnectionFormat = _connection_format.String
		objMetadata.MetadataChannel = _metadata_channel.String
		objMetadata.MetadataConvertfrequency = _metadata_convertfrequency.String
		objMetadata.MetadataContact = _metadata_contact.String

		if _metadata_agencystoredate.Valid {
			objMetadata.MetadataAgencystoredate = _metadata_agencystoredate.Time.Format(strDatetimeFormat)
		}
		if _metadata_startdatadate.Valid {
			objMetadata.MetadataStartdatadate = _metadata_startdatadate.Time.Format(strDatetimeFormat)
		}

		if _metadata_receive_date.Valid {
			objMetadata.MetadataReceiveDate = _metadata_receive_date.Time.Format(strDatetimeFormat)
		}

		objMetadata.MetadataUpdatePlan = _metadata_update_plan.Int64
		objMetadata.MetadataLaws = _metadata_laws.String
		objMetadata.MetadataRemark = _metadata_remark.String
		objMetadata.MetadataAgencyName = json.RawMessage(_metadataagency_name.String)
		objMetadata.MetadataServiceName = json.RawMessage(_metadataservice_name.String)
		objMetadata.MetadataTag = json.RawMessage(_metadata_tag.String)
		objMetadata.MetadataDescription = json.RawMessage(_metadata_description.String)
		objMetadata.MetadataStatusID = _metadata_status_id.Int64

		objMetadata.Servicemethod = objServiceMethod
		objMetadata.Hydroinfo = objHydroInfo
		objMetadata.Frequency = objFrequency
		objMetadata.History = objHistory

		objMetadata.CategoryID = _category_id.Int64
		objMetadata.MethodID = _metadata_method_id.Int64

		data = append(data, objMetadata)
	}

	return data, err
}

//	get ตารางบัญชีข้อมูล
//	Parameters:
//		param
//			Struct_Metadata_Table_InputParam
//	Return:
//		Array Struct_Metadata
func GetMetadataTable(param *Struct_Metadata_Table_InputParam) ([]*Struct_Metadata, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data        []*Struct_Metadata
		objMetadata *Struct_Metadata

		_id                   sql.NullInt64
		_metadataagency_name  sql.NullString
		_metadataservice_name sql.NullString

		_agency_id        sql.NullInt64
		_agency_shortname sql.NullString
		_agency_name      sql.NullString

		_subcat_id   sql.NullInt64
		_subcat_name sql.NullString

		_cat_id   sql.NullInt64
		_cat_name sql.NullString

		_hydro sql.NullString

		_result *sql.Rows
	)

	//-- Check Filter by parameters --//
	var arrParam = make([]interface{}, 0)
	var sqlCmdWhere string = ""

	//Check Filter agency_id
	if len(param.AgencyID) > 0 {
		if len(param.AgencyID) == 1 {
			arrParam = append(arrParam, param.AgencyID[0])
			sqlCmdWhere += " AND m.agency_id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, intId := range param.AgencyID {
				arrParam = append(arrParam, intId)
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND m.agency_id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter subcategory_id
	if len(param.SubcategoryID) > 0 {
		if len(param.SubcategoryID) == 1 {
			arrParam = append(arrParam, param.SubcategoryID[0])
			sqlCmdWhere += " AND sc.id = $" + strconv.Itoa(len(arrParam))
		} else {
			arrSqlCmd := []string{}
			for _, intId := range param.SubcategoryID {
				arrParam = append(arrParam, intId)
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND sc.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
		}
	}

	//Check Filter hydroinfo_id
	if len(param.Hydroinfo) > 0 {
		if len(param.Hydroinfo) == 1 {
			arrParam = append(arrParam, param.Hydroinfo[0])
			sqlCmdWhere += " AND EXISTS (SELECT hf.metadata_id FROM metadata_hydroinfo hf WHERE hf.hydroinfo_id = $" + strconv.Itoa(len(arrParam)) + " AND hf.deleted_by IS NULL AND hf.metadata_id = m.id) "
		} else {
			arrSqlCmd := []string{}
			for _, intId := range param.Hydroinfo {
				arrParam = append(arrParam, intId)
				arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
			}
			sqlCmdWhere += " AND EXISTS (SELECT hf.metadata_id FROM metadata_hydroinfo hf WHERE hf.hydroinfo_id IN (" + strings.Join(arrSqlCmd, ",") + ") AND hf.deleted_by IS NULL AND hf.metadata_id = m.id) "
		}
	}

	if param.CategoryID != 0 {
		arrParam = append(arrParam, param.CategoryID)
		sqlCmdWhere += " AND c.id = $" + strconv.Itoa(len(arrParam))
	}

	//	log.Printf(sqlGetMetadataTable+sqlCmdWhere+sqlGetMetadataTableGroupBy+" ORDER BY m.metadataservice_name->>'th' ", arrParam...)
	_result, err = db.Query(sqlGetMetadataTable+sqlCmdWhere+sqlGetMetadataTableGroupBy+" ORDER BY m.metadataservice_name->>'th' ", arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_Metadata, 0)

	// Loop data result
	for _result.Next() {

		//Scan to execute query with variables
		err := _result.Scan(&_id, &_metadataservice_name, &_metadataagency_name,
			&_agency_id, &_agency_shortname, &_agency_name,
			&_subcat_id, &_subcat_name,
			&_cat_id, &_cat_name,
			&_hydro)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}

		if !_metadataagency_name.Valid || _metadataagency_name.String == "" {
			_metadataagency_name.String = "{}"
		}
		if !_metadataservice_name.Valid || _metadataservice_name.String == "" {
			_metadataservice_name.String = "{}"
		}
		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_subcat_name.Valid || _subcat_name.String == "" {
			_subcat_name.String = "{}"
		}
		if !_cat_name.Valid || _cat_name.String == "" {
			_cat_name.String = "{}"
		}

		objMetadata = &Struct_Metadata{}

		objMetadata.Id = _id.Int64
		objMetadata.MetadataId, _ = model.GetCipher().EncryptText(strconv.FormatInt(_id.Int64, 10))
		objMetadata.Metadataagency_Name = json.RawMessage(_metadataagency_name.String)
		objMetadata.Metadataservice_Name = json.RawMessage(_metadataservice_name.String)

		objMetadata.Agency = &model_agency.Struct_Agency{}
		objMetadata.Agency.Id = _agency_id.Int64
		objMetadata.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		objMetadata.Agency.Agency_name = json.RawMessage(_agency_name.String)

		objMetadata.Subcategory = &model_subcategory.Struct_subcategory{}
		objMetadata.Subcategory.Id = _subcat_id.Int64
		objMetadata.Subcategory.Subcategory_name = json.RawMessage(_subcat_name.String)

		objMetadata.Subcategory.Category = &model_category.Struct_category{}
		objMetadata.Subcategory.Category.Id = _cat_id.Int64
		objMetadata.Subcategory.Category.Category_name = json.RawMessage(_cat_name.String)

		if !_hydro.Valid {
			objMetadata.Hydroinfo = nil
		} else {
			arrHydro := strings.Split(_hydro.String, "|")
			for _, strHydro := range arrHydro {
				arrStrHydro := strings.Split(strHydro, "##")
				intHydroInfo, err := strconv.ParseInt(arrStrHydro[0], 10, 64)
				if err != nil {
					return nil, err
				}
				objHydro := &model_hydroinfo.Struct_Hydroinfo{}
				objHydro.ID = intHydroInfo
				objHydro.HydroinfoName = json.RawMessage(arrStrHydro[1])
				objMetadata.Hydroinfo = append(objMetadata.Hydroinfo, objHydro)
			}
		}

		data = append(data, objMetadata)
	}

	return data, err
}
