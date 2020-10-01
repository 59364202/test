// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agency is a model for public.agency table. This table store agency.
package agency

import (
	"bufio"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"haii.or.th/api/thaiwater30/util/file"
	"haii.or.th/api/util/pqx"
	//	"log"
)

// 	get agency หาหน่วยงานทั้งหมดตาม id, department_id, ministry_id, sqlCmdWhere อย่างใดอย่างหนึ่ง
//	Parameters:
//		id
//			agency id
//		department_id
//			department id
//		ministry_id
//			ministry id
//		sqlCmdWhere
//			คำสั่ง sql ในส่วนของ where
// 	Return:
//		[]Struct_Agency
func getAgency(id, department_id, ministry_id int64, sqlCmdWhere string) ([]*Struct_Agency, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	strSql := SQL_SelectAgency
	var (
		param int64 = 0
		row   *sql.Rows

		data   []*Struct_Agency
		agency *Struct_Agency

		_id               int64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString
		_department_id    sql.NullInt64
		_deparmtnet_name  sql.NullString
		_ministry_id      sql.NullInt64
		_ministry_name    sql.NullString
		_logo             sql.NullString
		_aspects          sql.NullString
	)

	var sqlOrderBy string = " ORDER BY agency_name->>'th' "

	if id == 0 && department_id == 0 && ministry_id == 0 {
		//		log.Printf(strSql + sqlCmdWhere + sqlOrderBy)
		// query
		row, err = db.Query(strSql + sqlCmdWhere + sqlOrderBy)
	} else {
		if id != 0 {
			strSql += " AND a.id = $1 AND a.deleted_at = '1970-01-01 07:00:00+07' "
			param = id
		} else if department_id != 0 {
			strSql += " AND d.id = $1 AND d.deleted_at = '1970-01-01 07:00:00+07' "
			param = department_id
		} else if ministry_id != 0 {
			strSql += " AND m.id = $1 AND m.deleted_at = '1970-01-01 07:00:00+07' "
			param = ministry_id
		}
		//		log.Printf(strSql+sqlOrderBy, param)
		// query
		row, err = db.Query(strSql+sqlOrderBy, param)
	}
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for row.Next() {
		err = row.Scan(&_id, &_agency_name, &_agency_shortname, &_department_id, &_deparmtnet_name, &_ministry_id, &_ministry_name, &_logo, &_aspects)
		if err != nil {
			return nil, err
		}

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if _deparmtnet_name.String == "" {
			_deparmtnet_name.String = "{}"
		}
		if _aspects.String == "" {
			_aspects.String = "{}"
		}
		if _logo.String != "" {
			fn, _ := os.Open(filepath.Join(file.UploadPath, _logo.String))
			content, _ := ioutil.ReadAll(bufio.NewReader(fn))
			_logo.String = base64.StdEncoding.EncodeToString(content)
		} else {
			//defualt logo
			_logo.String = "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEABAMAAACuXLVVAAAAMFBMVEXDw8P///8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC8xHK4AAAACXBIWXMAAAsTAAALEwEAmpwYAAACKklEQVR42u2bUXKEMAxD8Q2s+1+2nelfZ1rCxuYlIF3AbyWTDYk5DsuyLMuyLMuyLMuyLMuyLMuydpZ+lGx5iCH0SwnX/xblPkGgP5Rw/bsIJJZA/4rp/zsJzuq3h6BT0fWbCUYAEq7faUEIJhAMMGpAG4FggHEDmh4ECbbgCkCyCfRYIBjgmgENGVwEEJxAPcBVA8ozuAwgOIEFABJOgAcQnMACAAknYIDSJtCeAPkcgDCA4C7EAWQAwV1IA4QdMMC2ALIDdsAOeCEygHdErwd4zrbcr2Z+Pd/1iOZJp2SxJUDpeb0B8PuC2BCg+OJww2u7YgD85vSAW4C/PefnBw44AX6G5FoGLZNM+CBT0AAHnMAC43wBG7DASGfABowSHI2CDRiyoHm4PNAARkJo/8Qh4PonBLd8ZBJgA5x2Afz7F+jBBZYBfiXm/4v47QC/I+shwN8L9no1fN5sOX5IRZ8T0tcF+HE9fW2HX15/PEAh2IAqCyTWgpgBSLh+hQWTAIl2QIUFMQuQsAGzFsQ8QMIGTFpQAZBwfR5AaAvOWSCxFkQVQO4KILEZRB1AwgbwAIIT+MwCHECCM6gFSDaBTwAkOAMaIKoBcjcACc6ABoh6gNwLQIIzoAGiAyB3ApDgDF4PED0AuQ+ABGfweoDoAkgDwC0w3ASvB4g+gDQA3AKDTWCAToBke3DMgXi0AwZI+CEY6kIDGIAGiF6ANIABDAAvAwMLgR14HcAXzNBOp6npkCsAAAAASUVORK5CYII="
		}

		agency = &Struct_Agency{}
		agency.Id = _id
		agency.Agency_name = json.RawMessage(_agency_name.String)
		agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		agency.Department_id = _department_id.Int64
		agency.Department_name = json.RawMessage(_deparmtnet_name.String)
		agency.Ministry_id = _ministry_id.Int64
		agency.Ministry_name = json.RawMessage(_ministry_name.String)
		agency.Aspects = json.RawMessage(_aspects.String)
		agency.Logo = _logo.String

		data = append(data, agency)
	}

	return data, nil
}

// โลโก้ของหน่วยงาน
//	Parameters:
// 	Return:
//		[]Logo_Agency
func GetAgencyLogo() ([]*Logo_Agency, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	strSql := SQL_AgencyLogo
	var (
		row  *sql.Rows
		data []*Logo_Agency

		_id   int64
		_logo sql.NullString
	)
	row, err = db.Query(strSql)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_logo)
		if err != nil {
			return nil, err
		}

		if _logo.String != "" {
			fn, _ := os.Open(filepath.Join(file.UploadPath, _logo.String))
			content, _ := ioutil.ReadAll(bufio.NewReader(fn))
			_logo.String = base64.StdEncoding.EncodeToString(content)
		}

		agency_logo := &Logo_Agency{}
		agency_logo.Id = _id
		agency_logo.Logo = _logo.String

		data = append(data, agency_logo)
	}

	return data, nil
}

//	 get all agency
// 	Return:
//		[]Struct_Agency
func GetAllAgency() ([]*Struct_Agency, error) {
	return getAgency(0, 0, 0, " AND a.deleted_at = '1970-01-01 07:00:00+07' ")
}

//	get agency ตาม agency id
//	Parameters:
//		id
//			agency id
// 	Return:
//		[]Struct_Agency
func GetAgencyById(id int64) ([]*Struct_Agency, error) {
	return getAgency(id, 0, 0, "")
}

//	get agency ตาม department id
//	Parameters:
//		department_id
//			department id
// 	Return:
//		[]Struct_Agency
func GetAgencyByDepartmentId(department_id int64) ([]*Struct_Agency, error) {
	return getAgency(0, department_id, 0, "")
}

//	get agency ตาม ministry id
//	Parameters:
//		ministry_id
//			ministry id
// 	Return:
//		[]Struct_Agency
func GetAgencyByMinistryId(ministry_id int64) ([]*Struct_Agency, error) {
	return getAgency(0, 0, ministry_id, "")
}

//	get agency ที่มีอยู่ใน table
//	Parameters:
//		tableName
//			ชื่อตาราง ที่ต้องการหาหน่วยงาน
// 	Return:
//		[]Struct_Agency
func GetAgencyInTable(tableName string) ([]*Struct_Agency, error) {
	var sqlCmdWhere = " AND EXISTS (SELECT agency_id FROM " + tableName + " WHERE " + tableName + ".agency_id = a.id) "
	return getAgency(0, 0, 0, sqlCmdWhere)
}

//	get agency ตามขนาดของเขื่อน
//	Parameters:
//		damSize
//			ขนาดของเขื่อน (0, 1, 2)
// 	Return:
//		[]Struct_Agency
func GetAgencyByDam(damSize int64) ([]*Struct_Agency, error) {

	sqlCmd := ""
	switch damSize {
	case 0:
		sqlCmd = " SELECT dam_id FROM dam_daily UNION ALL SELECT mediumdam_id FROM medium_dam "
	case 1:
		sqlCmd = " SELECT dam_id FROM dam_daily "
	case 2:
		sqlCmd = " SELECT md.agency_id FROM (" + sqlCmd + ") dam LEFT JOIN m_dam md ON dam.dam_id = md.id WHERE md.agency_id = a.id "
	default:
		return nil, nil
	}

	var sqlCmdWhere = " AND EXISTS (SELECT md.agency_id FROM (" + sqlCmd + ") dam LEFT JOIN m_dam md ON dam.dam_id = md.id WHERE md.agency_id = a.id ) "
	return getAgency(0, 0, 0, sqlCmdWhere)
}

//	get agency ที่มีภาพอยู่ใน ตาราง media
// 	Return:
//		[]Struct_Agency
func GetAgencyInMediaTable() ([]*Struct_Agency, error) {
	var sqlCmdWhere = " AND EXISTS( 	SELECT agency_id " +
		"					FROM media " +
		"					LEFT JOIN lt_media_type ON media.media_type_id = lt_media_type.id " +
		"					WHERE lt_media_type.media_category = 'image' " +
		"  				  AND media.agency_id = a.id " +
		"					  AND lt_media_type.deleted_by IS NULL " +
		"  				  AND media.deleted_by IS NULL) " +
		"   AND a.deleted_by IS NULL "
	return getAgency(0, 0, 0, sqlCmdWhere)
}

//	get agency	ที่มีอยู่ในตารางตามเงื่อนไข
//	Parameters:
//		tableName
//			ชื่อตาราง master
//		stationTableName
//			ชื่อตาราง transection
//		stationColumnName
//			ชื่อ field ที่เป็น foreign key ของ tableName
// 	Return:
//		[]Struct_Agency
func GetAgencyInStationTable(tableName string, stationTableName string, stationColumnName string) ([]*Struct_Agency, error) {

	var sqlCmdWhere = " AND a.id IN ( "

	switch tableName {
	case "rainfall":
		sqlCmdWhere += " SELECT distinct(agency_id) FROM m_tele_station m WHERE m.tele_station_type IN ('R','A') "
	case "tele_waterlevel":
	case "waterlevel":
		sqlCmdWhere += " SELECT distinct(agency_id) FROM m_tele_station m WHERE m.tele_station_type IN ('W','A') "
	case "dam_daily":
		sqlCmdWhere += " SELECT distinct(agency_id) FROM m_dam d WHERE EXISTS (SELECT * FROM dam_daily dd WHERE d.id = dd.dam_id) "
	case "dam_hourly":
		sqlCmdWhere += " SELECT distinct(agency_id) FROM m_dam d WHERE EXISTS (SELECT * FROM dam_hourly dd WHERE d.id = dd.dam_id) "
	}

	sqlCmdWhere += " )"

	//	var sqlCmdWhere = " AND EXISTS (SELECT agency_id FROM " + stationTableName + " LEFT JOIN " + tableName + " ON " + stationTableName + ".id = " + tableName + "." + stationColumnName + " WHERE " + tableName + "." + stationColumnName + " IS NOT NULL AND " + stationTableName + ".agency_id = a.id) "
	return getAgency(0, 0, 0, sqlCmdWhere)
}

//	get agency ที่มีอยู่ใน ตาราง m_tele_station, canal_waterlevel
// 	Return:
//		[]Struct_Agency
func GetAgencyInCanalWaterlevelTable() ([]*Struct_Agency, error) {
	var sqlCmdWhere = " AND EXISTS (  " +
		"			SELECT agt.agency_id  " +
		"			FROM (  " +
		"					SELECT cs.agency_id FROM m_canal_station cs WHERE cs.is_ignore = false AND cs.deleted_at = '1970-01-01 07:00:00+07'  " +
		"					UNION   " +
		"					SELECT ts.agency_id FROM m_tele_station ts LEFT JOIN ignore ig ON m.id = ig.station_id ::int AND ig.data_category = 'tele_waterlevel' AND (ig.is_ignore = false OR ig.is_ignore IS NULL) WHERE ts.deleted_at = '1970-01-01 07:00:00+07' AND (upper(ts.tele_station_type) = 'A' OR upper(ts.tele_station_type) = 'W' OR ts.tele_station_type IS NULL) " +
		"				 ) agt  " +
		"			WHERE a.id = agt.agency_id  " +
		"		   ) "
	return getAgency(0, 0, 0, sqlCmdWhere)
}

//	get ข้อมูล ที่ใช้ในหน้า agency_shopping
//	สถานะต่างๆของบัญชีข้อมูล แยกตามหน่วยงาน
//	Parameters:
//		user_id
//			id ของผู้ใช้
// 	Return:
//		[]Struct_Agency
func GetAgencyShopping(user_id int64) ([]*Struct_AgencyShopping, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		_id           int64
		_agency_name  sql.NullString
		_metadata     sql.NullInt64
		_connect      sql.NullInt64
		_wait_update  sql.NullInt64
		_wait_connect sql.NullInt64
		_dataservice  sql.NullInt64
	)
	// query
	row, err := db.Query(SQL_SelectAgencyShopping, user_id)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := make([]*Struct_AgencyShopping, 0)
	for row.Next() {
		err = row.Scan(&_id, &_agency_name, &_metadata, &_connect, &_wait_update, &_wait_connect, &_dataservice)
		if err != nil {
			return nil, err
		}

		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		agency := new(Struct_AgencyShopping)
		agency.AgencyId = _id
		agency.AgencyName = json.RawMessage(_agency_name.String)
		agency.Metadata = _metadata.Int64
		agency.MetadataStatus_Connect = _connect.Int64
		agency.MetadataStatus_WaitUpdate = _wait_update.Int64
		agency.MetadataStatus_WaitConnect = _wait_connect.Int64
		agency.DataService = _dataservice.Int64
		data = append(data, agency)
	}

	return data, nil
}

////	get ข้อมูล ที่ใชี้ในหน้า agency_shopping
//func GetAgencyMetadataSummary(year, month int64) ([]*Struct_AgencyMetadataSummary, error) {
//	db, err := pqx.Open()
//	if err != nil {
//		return nil, err
//	}
//	if year == 0 || month == 0 || month > 12 || month < 1 {
//		return nil, rest.NewError(500, " wrong param ", nil)
//	}
//	row, err := db.Query(SQL_SelectAgencyMetadataSummary, year, month)
//	if err != nil {
//		return nil, pqx.GetRESTError(err)
//	}
//	data := make([]*Struct_AgencyMetadataSummary, 0)
//	var (
//		object *Struct_AgencyMetadataSummary
//
//		_id    int64
//		_name  sql.NullString
//		_total sql.NullInt64
//		_count sql.NullInt64
//	)
//	for row.Next() {
//		err = row.Scan(&_id, &_name, &_total, &_count)
//		if err != nil {
//			return nil, err
//		}
//
//		object = &Struct_AgencyMetadataSummary{AgencyId: _id}
//		object.AgencyName = _name.String
//		object.Total = _total.Int64
//		object.Count = _count.Int64
//
//		data = append(data, object)
//	}
//
//	return data, nil
//}
