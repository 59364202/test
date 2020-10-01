// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package order_header is a model for dataservice.order_header table. This table store order_header.
package order_header

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"haii.or.th/api/util/pqx"

	model_order_status "haii.or.th/api/thaiwater30/model/order_status"
)

var StringDateLayout string = "2006-01-02T15:04"

//	get order header by user id
//	Parameters:
//		user_id
//			รหัสผู้ใช้
//	Return:
//		Array Struct_OH
func GetOrderHeaderByUserId(user_id int64) ([]*Struct_OH, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		data        []*Struct_OH
		orderHeader *Struct_OH

		_id        int64
		_datetime  time.Time
		_quanlity  int64
		_status_id int64
		_status    string
	)

	row, err := db.Query(SQL_SelectOrderHeaderByUserId, user_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_datetime, &_quanlity, &_status, &_status)
		if err != nil {
			return nil, err
		}

		orderHeader = &Struct_OH{}
		orderHeader.Id = _id
		orderHeader.Order_Datetime = _datetime.Format("2006-01-02 15:04:05")
		orderHeader.Order_Quality = _quanlity
		orderHeader.OrderStatus = &model_order_status.Struct_OrderStatus{Id: _status_id, Order_Status: _status}

		data = append(data, orderHeader)
	}

	return data, nil
}

//	get order header by order header id
//	Paremeters:
//		orderHeader
//			Struct_OrderHeader
//	Return:
//		Struct_OrderHeader
//func GetOrderByOrderHeaderId(orderHeader *Struct_OrderHeader) (*Struct_OrderHeader, error) {
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, err
//	}
//	var (
//		_status_id   int64
//		_quality     int64
//		_purpose     string
//		_forexternal sql.NullBool
//	)
//	row, err := db.Query(SQL_SelectOrderById, orderHeader.Id)
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//	for row.Next() {
//		err = row.Scan(&_status_id, &_quality, &_purpose, &_forexternal)
//		if err != nil {
//			return nil, err
//		}
//
//		if !_forexternal.Valid {
//			_forexternal.Bool = false
//		}
//
//		orderHeader.Order_Status_id = _status_id
//		orderHeader.Order_Quality = _quality
//		orderHeader.Order_Purpose = _purpose
//		orderHeader.Order_Forexternal = _forexternal.Bool
//	}
//
//	return nil, nil
//}

//	get order_header สำหรับบุคคลภายนอก
//	Parameter:
//		param
//			Param
//	Return:
//		Struct_OrderHeader
func GetOrderForExternal(param *Param) ([]*Struct_OrderHeader, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows

		data  []*Struct_OrderHeader
		order *Struct_OrderHeader

		_id   int64
		_date time.Time
	)

	if param.Date != "" {
		strSql := " AND date(order_datetime) = $1"
		row, err = db.Query(SQL_SelectOrderForExternal+strSql, param.Date)
	} else {
		row, err = db.Query(SQL_SelectOrderForExternal)
	}
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for row.Next() {
		err = row.Scan(&_id, &_date)
		if err != nil {
			return nil, err
		}
		order = &Struct_OrderHeader{Id: _id, Order_Datetime: _date.Format("2006-01-02")}
		data = append(data, order)
	}

	return data, nil
}

//	get order header ของหน้า data_service_management
//	Parameters:
//		param
//			Param
//	Return:
//		Struct_OrderHeader
func GetOrderHeaderByParamManagement(param *Param) ([]*Struct_OrderHeader, error) {
	return getOrderHeader(param.Date_Start, param.Date_End, param.Agency_Id, param.Statud_Id)
}

//  ดึงข้อมูล order_purpose ที่ใช้งานบ่อยที่สุด 10 อันดับจากตาราง order_header
//	Parameters:
//		None
//	Return:
//		Struct_Order_Purpose
func GetPopularOrderPurpose() ([]*Struct_Order_Purpose, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSQL string = SQL_SelectOrderHeaderForOrderPurposePopular
		row    *sql.Rows
		data   []*Struct_Order_Purpose

		_order_purpose           string //วัตุประสงค์
		_order_purpose_frequency int64  //จำนวนวัตถุประสงค์ที่ซ้ำ
	)

	row, err = db.Query(strSQL)

	for row.Next() {
		err = row.Scan(&_order_purpose, &_order_purpose_frequency)
		if err != nil {
			return nil, err
		}

		OderPurpose := &Struct_Order_Purpose{}
		OderPurpose.Order_Purpose = _order_purpose
		OderPurpose.Id = _order_purpose_frequency //ไม่ใช่ id ของ order_purpose แต่เป็นจำนวนของ order_purpose ที่ซ้ำกัน
		data = append(data, OderPurpose)
	}

	return data, nil
}

//	get order header ตาม datestart, dateend, agency_id, status_id
//	Parameters:
//		datestart
//			วันที่เริ่มต้น
//		dateend
//			วันที่สิ้นสุด
//		agency_id
//			รหัสหน่วยงาน
//		status_id
//			รหัสของสถานะการขอข้อมูล
//	Return:
//		Array Struct_OrderHeader
func getOrderHeader(datestart, dateend string, agency_id, status_id int64) ([]*Struct_OrderHeader, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSQL       string = SQL_SelectOrderHeader
		strSQL_Where string = ""
		param        []interface{}
		countParam   int = 0

		row         *sql.Rows
		data        []*Struct_OrderHeader
		orderHeader *Struct_OrderHeader

		_id               int64
		_datetime         time.Time
		_user_id          int64
		_user_name        string
		_status_id        int64
		_order_status     string
		_user_agency_id   sql.NullInt64
		_user_agency_name sql.NullString
	)

	if datestart != "" {
		countParam++
		strSQL_Where += " date(oh.order_datetime) >= $" + strconv.Itoa(countParam) + " AND"
		param = append(param, datestart)
	}
	if dateend != "" {
		countParam++
		strSQL_Where += " date(oh.order_datetime) <= $" + strconv.Itoa(countParam) + " AND"
		param = append(param, dateend)
	}
	if agency_id != 0 {
		countParam++
		strSQL_Where += " u.agency_id = $" + strconv.Itoa(countParam) + " AND"
		param = append(param, agency_id)
	}
	if status_id != 0 {
		countParam++
		strSQL_Where += " ohs.id = $" + strconv.Itoa(countParam) + " AND"
		param = append(param, status_id)
	}

	if countParam != 0 {
		strSQL_Where = " WHERE " + strSQL_Where[0:len(strSQL_Where)-3]
		row, err = db.Query(strSQL+strSQL_Where+SQL_SelectOrderHeader_OrderById, param...)
	} else {
		row, err = db.Query(strSQL + SQL_SelectOrderHeader_OrderById)
	}
	if err != nil {
		return nil, err
	}
	for row.Next() {
		err = row.Scan(&_id, &_datetime, &_user_id, &_user_name, &_status_id, &_order_status, &_user_agency_id, &_user_agency_name)
		if err != nil {
			return nil, err
		}
		if !_user_agency_name.Valid || _user_agency_name.String == "" {
			_user_agency_name.String = "{}"
		}

		orderHeader = &Struct_OrderHeader{}
		orderHeader.Id = _id
		orderHeader.Order_Datetime = _datetime.Format("2006-01-02 15:04")
		orderHeader.User_Id = _user_id
		orderHeader.User_Fullname = _user_name
		orderHeader.User_Agency_Id = _user_agency_id.Int64
		orderHeader.User_Agency_Name = json.RawMessage(_user_agency_name.String)
		orderHeader.Order_Status_id = _status_id
		orderHeader.OrderStatus = &model_order_status.Struct_OrderStatus{Id: _status_id, Order_Status: _order_status}

		data = append(data, orderHeader)
	}

	return data, nil
}
