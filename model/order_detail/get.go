// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"database/sql"
	"encoding/json"
	"strings"

	//	model_order_header "haii.or.th/api/thaiwater30/model/order_header"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"

	//	"haii.or.th/api/thaiwater30/util/b64"
	"haii.or.th/api/thaiwater30/util/file"
	"haii.or.th/api/thaiwater30/util/sqltime"

	model "haii.or.th/api/server/model"
	"haii.or.th/api/server/model/setting"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_category "haii.or.th/api/thaiwater30/model/lt_category"
	model_lt_department "haii.or.th/api/thaiwater30/model/lt_department"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_lt_ministry "haii.or.th/api/thaiwater30/model/lt_ministry"
	model_lt_servicemethod "haii.or.th/api/thaiwater30/model/lt_servicemethod"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_order_detail_status "haii.or.th/api/thaiwater30/model/order_detail_status"
)

var StringDateLayout string = time.RFC3339

type URLBuilder interface {
	BuildURL(version int64, service string, buildabs bool) string
}

// scan order_detail ใส่ลง []Struct_OrderDetail
func scanOrderDetail(row *sql.Rows, bx URLBuilder) ([]*Struct_OrderDetail, error) {
	var (
		data        []*Struct_OrderDetail
		orderDetail *Struct_OrderDetail

		_metadata_id               int64
		_metadataservice_name      sql.NullString
		_status_id                 int64
		_detail_status             sql.NullString
		_service_id                int64
		_servicemethod_name        sql.NullString
		_detail_province           sql.NullString
		_detail_basin              sql.NullString
		_category_id               int64
		_category_name             sql.NullString
		_agency_id                 int64
		_agency_name               sql.NullString
		_from_date                 sql.NullString
		_to_date                   sql.NullString
		_detail_letterdate         sql.NullString
		_detail_letterno           sql.NullString
		_order_header_id           int64
		_department_id             int64
		_department_name           sql.NullString
		_ministry_id               int64
		_ministry_name             sql.NullString
		_id                        int64
		_detail_source_result      sql.NullString
		_detail_source_result_date sql.NullString
		_user_agency               sql.NullString
		_user_id                   int64
		_user_office_name          sql.NullString
		_user_fullname             sql.NullString
		_user_ministry_name        sql.NullString
		_user_contact_phone        sql.NullString
		_order_purpose             sql.NullString
		_e_id                      sql.NullString
		_is_enabled                sql.NullBool
		_count_log                 sql.NullInt64
		_latest_access_time        sqltime.NullTime
		_detail_frequency          sql.NullString
	)
	for row.Next() {
		err := row.Scan(&_metadata_id, &_metadataservice_name, &_status_id, &_detail_status, &_service_id, &_servicemethod_name, &_detail_province,
			&_detail_basin, &_category_id, &_category_name, &_agency_id, &_agency_name, &_from_date, &_to_date, &_detail_letterdate, &_detail_letterno, &_order_header_id,
			&_department_id, &_department_name, &_ministry_id, &_ministry_name, &_id, &_detail_source_result, &_detail_source_result_date, &_user_agency, &_user_id,
			&_user_office_name, &_user_fullname, &_user_ministry_name, &_user_contact_phone, &_order_purpose, &_e_id, &_is_enabled, &_count_log, &_latest_access_time,
			&_detail_frequency)
		if err != nil {
			return nil, err
		}
		if !_metadataservice_name.Valid || _metadataservice_name.String == "" {
			_metadataservice_name.String = "{}"
		}
		if !_servicemethod_name.Valid || _servicemethod_name.String == "" {
			_servicemethod_name.String = "{}"
		}
		if !_category_name.Valid || _category_name.String == "" {
			_category_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_department_name.Valid || _department_name.String == "" {
			_department_name.String = "{}"
		}
		if !_ministry_name.Valid || _ministry_name.String == "" {
			_ministry_name.String = "{}"
		}
		if !_user_agency.Valid || _user_agency.String == "" {
			_user_agency.String = "{}"
		}
		if !_user_ministry_name.Valid || _user_ministry_name.String == "" {
			_user_ministry_name.String = "{}"
		}

		orderDetail = &Struct_OrderDetail{}
		orderDetail.Id = _id
		orderDetail.Order_Header_Id = _order_header_id
		orderDetail.Order_Purpose = _order_purpose.String
		orderDetail.Detail_Frequency = _detail_frequency.String
		orderDetail.Detail_Letterdate = _detail_letterdate.String
		orderDetail.Detail_Letterno = _detail_letterno.String
		orderDetail.Detail_Source_Result = _detail_source_result.String
		orderDetail.IsEnabled = _is_enabled.Bool
		orderDetail.CountLog = _count_log.Int64

		orderDetail.User_Id = _user_id
		orderDetail.User_Fullname = _user_fullname.String
		orderDetail.User_ContractPhone = _user_contact_phone.String
		orderDetail.User_OfficeName = _user_office_name.String
		orderDetail.E_Id = _e_id.String
		orderDetail.User_MinistryName = json.RawMessage(_user_ministry_name.String)
		orderDetail.User_AgencyName = json.RawMessage(_user_agency.String)

		if _detail_source_result_date.Valid {
			t, err := time.Parse(StringDateLayout, _detail_source_result_date.String)
			if err == nil {
				orderDetail.Detail_Source_Result_Date = t.Format("2006-01-02")
			}
		}

		orderDetail.Service_Id = _service_id
		orderDetail.Metadata = &model_metadata.Struct_Metadata{}
		orderDetail.Metadata.Id = _metadata_id
		orderDetail.Metadata.Metadataservice_Name = json.RawMessage(_metadataservice_name.String)

		orderDetail.Service = &model_lt_servicemethod.Struct_LtServicemethod{ID: _service_id, Name: json.RawMessage(_servicemethod_name.String)}
		orderDetail.Order_Detail_Status = &model_order_detail_status.Struct_OrderDetailStatus{Id: _status_id, Detail_Status: _detail_status.String}
		orderDetail.Category = &model_lt_category.Struct_category{Id: _category_id, Category_name: json.RawMessage(_category_name.String)}
		orderDetail.Agency = &model_agency.Struct_A{Id: _agency_id, Agency_name: json.RawMessage(_agency_name.String)}
		orderDetail.Department = &model_lt_department.Struct_Department{Id: _department_id, Department_Name: json.RawMessage(_department_name.String)}
		orderDetail.Ministry = &model_lt_ministry.Struct_Ministry{Id: _ministry_id, Ministry_Name: json.RawMessage(_ministry_name.String)}

		if _latest_access_time.Valid {
			orderDetail.LatestAccessTime = _latest_access_time.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		}

		if _from_date.Valid {
			t, err := time.Parse(StringDateLayout, _from_date.String)
			if err == nil {
				orderDetail.Detail_Fromdate = t.Format("2006-01-02")
			}
		}

		if _to_date.Valid {
			t, err := time.Parse(StringDateLayout, _to_date.String)
			if err == nil {
				orderDetail.Detail_Todate = t.Format("2006-01-02")
			}
		}

		if _detail_province.Valid && _detail_province.String != "" {
			rs, err := model_lt_geocode.GetProvinceFromGeocode(_detail_province.String)
			if err == nil {
				orderDetail.Province = rs
			}
		}
		if _detail_basin.Valid && _detail_basin.String != "" {
			rs, err := model_basin.GetBasinFromCode(_detail_basin.String)
			if err == nil {
				orderDetail.Basin = rs
			}
		}

		if _e_id.String != "" {
			// มี e_id สร้าง service url
			mid := strconv.FormatInt(_metadata_id, 10)
			orderDetail.Service_Url = bx.BuildURL(0, "thaiwater30/api_service", true) + "?mid=" + mid + "&eid=" + _e_id.String
		}

		data = append(data, orderDetail)
	}
	return data, nil
}

// get order_detail จาก order_header_id
func GetOrderDetailByOrderHeaderId(id int64, bx URLBuilder) ([]*Struct_OrderDetail, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSQL string
		row    *sql.Rows
	)

	strSQL = SQL_SelectOrderDetail

	if id != 0 {
		strSQL += "WHERE od.order_header_id = $1"
		row, err = db.Query(strSQL+SQL_SelectOrderDetail_OrderBy, id)
	} else {
		row, err = db.Query(strSQL + SQL_SelectOrderDetail_OrderBy)
	}

	if err != nil {
		return nil, err
	}

	return scanOrderDetail(row, bx)
}

// get order_detail จาก order_header_id โดยไม่เอา service cd/dvd
func GetOrderDetailByOrderHeaderIdWithoutCDDVD(id int64, bx URLBuilder) ([]*Struct_OrderDetail, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSQL string
		row    *sql.Rows
	)

	strSQL = SQL_SelectOrderDetail + " WHERE od.service_id <> 3"

	if id != 0 {
		strSQL += " AND od.order_header_id = $1 "
		row, err = db.Query(strSQL+SQL_SelectOrderDetail_OrderBy, id)
	} else {
		row, err = db.Query(strSQL + SQL_SelectOrderDetail_OrderBy)
	}

	if err != nil {
		return nil, err
	}

	return scanOrderDetail(row, bx)
}

// get order_detail ตาม parameter
func GetOrderDetail(param *Param, bx URLBuilder) ([]*Struct_OrderDetail, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	strSQL := SQL_SelectOrderDetail
	strSQLWhere := " WHERE od.service_id <> 3 AND oh.order_status_id <> 3 "
	strAalWhere := " WHERE aal.service_id = 107  "
	itf := make([]interface{}, 0)

	if param.User_Id > 0 {
		itf = append(itf, param.User_Id)
		strSQLWhere += " AND au.id = $" + strconv.Itoa(len(itf))
	}
	if param.Agency_Id > 0 {
		itf = append(itf, param.Agency_Id)
		strSQLWhere += " AND a.id = $" + strconv.Itoa(len(itf))
	}
	if param.UserAgency_Id > 0 {
		itf = append(itf, param.UserAgency_Id)
		strSQLWhere += " AND ua.id = $" + strconv.Itoa(len(itf))
	}
	if param.Date_Start != "" {
		itf = append(itf, param.Date_Start)
		// strSQLWhere += " AND od.created_at >= $" + strconv.Itoa(len(itf))
		strAalWhere += " AND aal.created_at >= $" + strconv.Itoa(len(itf))
	}
	if param.Date_End != "" {
		// itf = append(itf, param.Date_End)
		itf = append(itf, param.Date_End+" 23:59:59")
		// strSQLWhere += " AND od.created_at < $" + strconv.Itoa(len(itf)) + "::DATE + interval '1 day' "
		// strSQLWhere += " AND od.created_at <= $" + strconv.Itoa(len(itf))
		strAalWhere += " AND aal.created_at <= $" + strconv.Itoa(len(itf))
	}
	if param.Eid != "" {
		itf = append(itf, param.Eid)
		strSQLWhere += " AND od.e_id = $" + strconv.Itoa(len(itf))
	}
	strSQL = strings.Replace(strSQL, "--WHERE", strAalWhere, 1)
	log.Println(strSQL, itf)
	row, err := db.Query(strSQL+strSQLWhere, itf...)
	if err != nil {
		return nil, err
	}

	return scanOrderDetail(row, bx)
}

// get order_detail ที่ status ตาม parameter
func GetOrderDetailByStatus(p *Param_OrderApprove, bx URLBuilder) ([]*Struct_OrderDetail, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSQL string
		row    *sql.Rows

		countP int = 2
		param  []interface{}
	)

	strSQL = SQL_SelectOrderDetail + " WHERE od.service_id = $1 AND od.detail_letterpath IS NOT NULL AND " +
		" od.detail_source_result IS NULL AND od.detail_source_result_date IS NULL "
	param = append(param, p.Service_Id)

	if p.Detail_Letterno != "0" && p.Detail_Letterno != "" {
		strSQL += " AND od.detail_letterno = $" + strconv.Itoa(countP)
		param = append(param, p.Detail_Letterno)
		countP++
	} else {
		strSQL += " AND od.detail_letterno IS NOT NULL"
	}

	if p.Agency_Id != 0 {
		strSQL += " AND a.id = $" + strconv.Itoa(countP)
		param = append(param, p.Agency_Id)
	}
	row, err = db.Query(strSQL, param...)
	if err != nil {
		return nil, err
	}
	return scanOrderDetail(row, bx)
}

//	get order_detail ที่ service_id = 3  และได้มีการ print เอกสารคำขอ ส่งหน่วยงานแล้ว
//	โดยจะกรุ๊ปข้อมูล ตาม order_header_id และ agency_id
func GetOrderDetailGroupByAgency() ([]*Struct_OrderDetail, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		data        []*Struct_OrderDetail
		orderDetail *Struct_OrderDetail

		_order_header_id   int64
		_agency_id         int64
		_agency_name       sql.NullString
		_detail_letterno   sql.NullString
		_detail_letterpath sql.NullString
		_department_id     int64
		_department_name   sql.NullString
		_ministry_id       int64
		_ministry_name     sql.NullString
		_detail_fromdate   sql.NullString
		_detail_todate     sql.NullString
		_metadata_name     sql.NullString
	)

	row, err := db.Query(SQL_SelectOrderDetailGroupByAgency)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_order_header_id, &_agency_id, &_agency_name, &_detail_letterno, &_detail_letterpath,
			&_department_id, &_department_name, &_ministry_id, &_ministry_name, &_detail_fromdate, &_detail_todate, &_metadata_name)
		if err != nil {
			return nil, err
		}

		if !_department_name.Valid || _department_name.String == "" {
			_department_name.String = "{}"
		}
		if !_ministry_name.Valid || _ministry_name.String == "" {
			_ministry_name.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_metadata_name.Valid || _metadata_name.String == "" {
			_metadata_name.String = "{}"
		}

		orderDetail = &Struct_OrderDetail{}
		orderDetail.Order_Header_Id = _order_header_id
		orderDetail.Detail_Letterno = _detail_letterno.String
		if _detail_letterpath.String != "" {
			//			orderDetail.Detail_Letterpath, _ = b64.EncryptText(filepath.Join(file.UploadPath, _detail_letterpath.String))
			orderDetail.Detail_Letterpath, _ = model.GetCipher().EncryptText(filepath.Join(file.UploadPath, _detail_letterpath.String))
		}

		orderDetail.Department = &model_lt_department.Struct_Department{Id: _department_id, Department_Name: json.RawMessage(_department_name.String)}
		orderDetail.Ministry = &model_lt_ministry.Struct_Ministry{Id: _ministry_id, Ministry_Name: json.RawMessage(_ministry_name.String)}
		orderDetail.Agency = &model_agency.Struct_A{Id: _agency_id, Agency_name: json.RawMessage(_agency_name.String)}
		orderDetail.Metadata = &model_metadata.Struct_Metadata{}
		orderDetail.Metadata.Metadataservice_Name = json.RawMessage(_metadata_name.String)

		if _detail_fromdate.Valid {
			t, err := time.Parse(StringDateLayout, _detail_fromdate.String)
			if err == nil {
				orderDetail.Detail_Fromdate = t.Format("2006-01-02")
			}
		}

		if _detail_todate.Valid {
			t, err := time.Parse(StringDateLayout, _detail_todate.String)
			if err == nil {
				orderDetail.Detail_Todate = t.Format("2006-01-02")
			}
		}

		data = append(data, orderDetail)
	}

	return data, nil
}

//	get order_detail for page data_service_summary
func GetOrderDetailSummary(p *Param) ([]*Struct_OrderDetail, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSQL       string
		strSQL_Where string = ""
		row          *sql.Rows
		param        []interface{}

		data        []*Struct_OrderDetail
		orderDetail *Struct_OrderDetail

		_id                        int64
		_order_header_id           int64
		_metadataservice_name      sql.NullString
		_user_full_name            string
		_servicemethod_name        sql.NullString
		_order_datetime            time.Time
		_detail_source_result_date sql.NullString
		_detail_source_result      sql.NullString
		_user_id                   int64
		_user_agency_name          sql.NullString
	)

	strSQL = SQL_SelectOrderDetailSummary

	if p.Date_Start != "" {
		param = append(param, p.Date_Start)
		strSQL_Where += " date(oh.order_datetime) >= $" + strconv.Itoa(len(param)) + " AND"

	}
	if p.Date_End != "" {
		param = append(param, p.Date_End)
		strSQL_Where += " date(oh.order_datetime) <= $" + strconv.Itoa(len(param)) + " AND"
	}
	if p.User_Id != 0 {
		param = append(param, p.User_Id)
		strSQL_Where += " u.id = $" + strconv.Itoa(len(param)) + " AND"
	}
	if p.Agency_Id != 0 {
		param = append(param, p.Agency_Id)
		strSQL_Where += " m.agency_id = $" + strconv.Itoa(len(param)) + " AND"
	}

	if len(param) != 0 {
		strSQL += " WHERE " + strSQL_Where[0:len(strSQL_Where)-3]
		row, err = db.Query(strSQL+SQL_SelectOrderDetailSummary_OrderBy, param...)
	} else {
		row, err = db.Query(strSQL + SQL_SelectOrderDetailSummary_OrderBy)
	}
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_order_header_id, &_metadataservice_name, &_user_full_name, &_servicemethod_name, &_order_datetime,
			&_detail_source_result_date, &_detail_source_result, &_user_id, &_user_agency_name)
		if err != nil {
			return nil, err
		}

		if !_metadataservice_name.Valid || _metadataservice_name.String == "" {
			_metadataservice_name.String = "{}"
		}
		if !_servicemethod_name.Valid || _servicemethod_name.String == "" {
			_servicemethod_name.String = "{}"
		}
		if !_user_agency_name.Valid || _user_agency_name.String == "" {
			_user_agency_name.String = "{}"
		}

		if _detail_source_result_date.Valid {
			t, err := time.Parse(StringDateLayout, _detail_source_result_date.String)
			if err == nil {
				_detail_source_result_date.String = t.Format("2006-01-02")
			}
		}

		orderDetail = &Struct_OrderDetail{}
		orderDetail.Id = _id
		orderDetail.User_Id = _user_id
		orderDetail.Order_Header_Id = _order_header_id
		orderDetail.User_Fullname = _user_full_name
		orderDetail.Detail_Source_Result = _detail_source_result.String
		orderDetail.Detail_Source_Result_Date = _detail_source_result_date.String
		orderDetail.Order_Header_Order_Datetime = _order_datetime.Format("2006-01-02")
		orderDetail.Service = &model_lt_servicemethod.Struct_LtServicemethod{Name: json.RawMessage(_servicemethod_name.String)}
		orderDetail.Metadata = &model_metadata.Struct_Metadata{}
		orderDetail.Metadata.Metadataservice_Name = json.RawMessage(_metadataservice_name.String)
		orderDetail.User_AgencyName = json.RawMessage(_user_agency_name.String)

		data = append(data, orderDetail)
	}

	return data, nil
}

// get จำนวนคำขอข้อมูลของหน่วยงาน(agency_id) แยกตามหน่วยงาน
func GetCountOrderDetailByAgencyId(agency_id int64) ([]*Struct_CountOrderDetailByAgencyId, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	if agency_id <= 0 {
		return nil, errors.New("no param")
	}
	row, err := db.Query(SQL_SelectCountOrderDetailByAgencyId, agency_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*Struct_CountOrderDetailByAgencyId, 0)
	var (
		object *Struct_CountOrderDetailByAgencyId

		_id    int64
		_name  sql.NullString
		_count int64
	)

	for row.Next() {
		err = row.Scan(&_id, &_name, &_count)
		if err != nil {
			return nil, err
		}

		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}

		object = &Struct_CountOrderDetailByAgencyId{}
		object.AgencyId = _id
		object.AgencyName = json.RawMessage(_name.String)
		object.Count = _count

		data = append(data, object)
	}

	return data, nil
}

// scan count from query select count(x) from xxx
func scanCount(q string, p []interface{}) (int64, error) {
	db, err := pqx.Open()
	if err != nil {
		return 0, err
	}

	row, err := db.Query(q, p...)
	if err != nil {
		return 0, err
	}
	var _count sql.NullInt64

	for row.Next() {
		err = row.Scan(&_count)
		if err != nil {
			return 0, err
		}
	}
	return _count.Int64, nil
}

// get จำนวน order detail ที่ยังไม่ได้บันทึกผลการอนุมัติ และ order header ยังไม่ถูกยกเลิก
func GetCountOrderDetailNoResult() (int64, error) {
	return scanCount(SQL_CountNoResult, []interface{}{})
}

// get จำนวน order detail ชนิด download ที่เปิดใช้งานอยู่ และ ยังไม่หมดอายุ
func GetCountOrderDetailDownloadOnlyEnable() (int64, error) {
	date := setting.GetSystemSetting("service.api-service.download.ExpireDate")
	return scanCount(Gen_SQL_CountOrderDetailDownloadOnlyEnable(date), []interface{}{})
}

// get จำนวน order detail ชนิดอื่นที่ไม่ใช่ download ที่เปิดใช้งานอยู่
func GetCountOrderDetailEnableWithoutDownload() (int64, error) {
	return scanCount(SQL_CountOrderDetailEnableWithoutDownload, []interface{}{})
}
