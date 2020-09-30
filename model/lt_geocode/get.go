// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_geocode is a model for public.lt_geocode table. This table store lt_geocode information.
package lt_geocode

import (
	"database/sql"
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	//	"log"
	"strconv"
	"strings"
	"time"
)

//	scan data
//	Parameters:
//		row
//			result of a query
//		hasArea
//			true ถ้าต้องการข้อมูล area
//		hasProvince
//			true ถ้าต้องการข้อมูล province
//		hasAmphoe
//			true ถ้าต้องการข้อมูล amphoe
//		hasTumbon
//			true ถ้าต้องการข้อมูล tumbon
//	Return:
//		[]Struct_Geocode
func scanGeocode(row *sql.Rows, hasArea, hasProvince, hasAmphoe, hasTumbon bool) ([]*Struct_Geocode, error) {
	var (
		data      []*Struct_Geocode
		s_geocode *Struct_Geocode
		err       error

		_id            int64
		_geocode       string
		_area_code     sql.NullString
		_province_code sql.NullString
		_amphoe_code   sql.NullString
		_tumbon_code   sql.NullString
		_area_name     sql.NullString
		_province_name sql.NullString
		_amphoe_name   sql.NullString
		_tumbon_name   sql.NullString
	)
	for row.Next() {
		err = row.Scan(&_id, &_geocode, &_area_code, &_province_code, &_amphoe_code, &_tumbon_code, &_area_name, &_province_name, &_amphoe_name, &_tumbon_name)
		if err != nil {
			//			log.Println(err)
			return nil, err
		}

		s_geocode = &Struct_Geocode{}
		s_geocode.Id = _id
		s_geocode.Geocode = _geocode

		if hasArea {
			if !_area_name.Valid || _area_name.String == "" {
				_area_name.String = "{}"
			}

			s_geocode.Area_code = _area_code.String
			s_geocode.Area_name = json.RawMessage(_area_name.String)
		}
		if hasProvince {
			if !_province_name.Valid || _province_name.String == "" {
				_province_name.String = "{}"
			}

			s_geocode.Province_code = _province_code.String
			s_geocode.Province_name = json.RawMessage(_province_name.String)
		}
		if hasAmphoe {
			if !_amphoe_name.Valid || _amphoe_name.String == "" {
				_amphoe_name.String = "{}"
			}

			s_geocode.Amphoe_code = _amphoe_code.String
			s_geocode.Amphoe_name = json.RawMessage(_amphoe_name.String)
		}
		if hasTumbon {
			if !_tumbon_name.Valid || _tumbon_name.String == "" {
				_tumbon_name.String = "{}"
			}

			s_geocode.Tumbon_code = _tumbon_code.String
			s_geocode.Tumbon_name = json.RawMessage(_tumbon_name.String)
		}
		data = append(data, s_geocode)
	}
	return data, nil
}

//	get all province
//	Return:
//		[]Struct_Geocode
func GetAllProvince() ([]*Struct_Geocode, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows
	)
	row, err = db.Query(SQL_selectGeocodeFromGeocode + SQL_selectGeocodeFromGeocode_AllProvince + SQL_selectGeocodeFromGeocode_OrderBy)
	if err != nil {
		return nil, err
	}
	return scanGeocode(row, false, true, false, false)
}

//	get province from prov_id
//	Parameters:
//		prov_id
//			รหัส จังหวัด
//	Return:
//		Struct_Geocode
func GetProvince(prov_id string) (*Struct_Geocode, error) {
	if prov_id == "" {
		return nil, errors.New("invalid prov_id")
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	q := SQL_selectGeocodeFromGeocode + SQL_selectGeocodeFromGeocode_AllProvince + " AND province_code = $1 " + SQL_selectGeocodeFromGeocode_OrderBy
	var s *Struct_Geocode
	row, err := db.Query(q, prov_id)
	if err != nil {
		return nil, err
	}
	sa, err := scanGeocode(row, false, true, false, false)
	if err != nil {
		return nil, err
	}
	if len(sa) == 1 {
		s = sa[0]
	}
	return s, nil
}

//	get geocode
//	Parameters:
//		geocode
//			geocode
//		hasArea
//			true ถ้าต้องการข้อมูล area
//		hasProvince
//			true ถ้าต้องการข้อมูล province
//		hasAmphoe
//			true ถ้าต้องการข้อมูล amphoe
//		hasTumbon
//			true ถ้าต้องการข้อมูล tumbon
//	Return:
//		[]Struct_Geocode
func getGeocodeFromGeocode(geocode string, hasArea, hasProvince, hasAmphoe, hasTumbon bool) ([]*Struct_Geocode, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows
	)
	strSql := SQL_selectGeocodeFromGeocode
	if geocode != "" {
		param := []interface{}{}
		strSql += " WHERE "
		gcs := strings.Split(geocode, ",")
		lastIndex := len(gcs) - 1

		for i, s := range gcs {
			param = append(param, s)
			strSql += " geocode = $" + strconv.Itoa(len(param))
			if i != lastIndex {
				strSql += " OR "
			}
		}
		row, err = db.Query(strSql+SQL_selectGeocodeFromGeocode_OrderBy, param...)
	} else {
		row, err = db.Query(strSql + SQL_selectGeocodeFromGeocode_OrderBy)
	}

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	return scanGeocode(row, hasArea, hasProvince, hasAmphoe, hasTumbon)
}

//	get province from geocode
//	Parameters:
//		geocode
//			geocode
//	Return:
//		[]Struct_Geocode
func GetProvinceFromGeocode(geocode string) ([]*Struct_Geocode, error) {
	return getGeocodeFromGeocode(geocode, false, true, false, false)
}

//	get all area
//	Return:
//		[]Struct_Region
func GetAllArea() ([]*Struct_Region, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlGetAllArea
	p := []interface{}{}

	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	data := make([]*Struct_Region, 0)
	for rows.Next() {
		var (
			areacode sql.NullString
			areaname pqx.JSONRaw
		)
		rows.Scan(&areacode, &areaname)
		dd := &Struct_Region{}
		dd.Area_code = areacode.String
		dd.Area_name = areaname.JSON()

		data = append(data, dd)
	}
	dd := &Struct_Region{}
	dd.Area_code = ""
	dd.Area_name = json.RawMessage("{\"th\":\"แสดงทั้งประเทศ\"}")
	data = append(data, dd)
	return data, nil
}

//	get station group by province
//	Return:
//		[]Province
func GetProvinceByStation() ([]*Province, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlGetAllProvince
	p := []interface{}{}

	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	data := make([]*Province, 0)
	aStation := make([]*Station, 0)
	province := &Province{}
	sProvinceCode := ""
	for rows.Next() {
		var (
			basincode    sql.NullString
			basinname    pqx.JSONRaw
			subbasincode sql.NullString
			subbasinname pqx.JSONRaw
			mid          sql.NullInt64
			maxDate      time.Time
			minDate      time.Time
		)
		rows.Scan(&basincode, &basinname, &subbasincode, &subbasinname, &mid, &maxDate, &minDate)

		if sProvinceCode != "" {
			if sProvinceCode != basincode.String {
				province.TeleStation = aStation
				data = append(data, province)

				aStation = make([]*Station, 0)
				province = &Province{}
				sProvinceCode = basincode.String
				province.ProvinceCode = basincode.String
				province.ProvinceName = basinname.JSON()
				sb := &Station{}
				sb.StationCode = subbasincode.String
				sb.StationName = subbasinname.JSON()
				sb.StationID = mid.Int64
				if maxDate.Year() > 1 {
					sb.MaxYear = maxDate.Year()
				}
				if minDate.Year() > 1 {
					sb.MinYear = minDate.Year()
				}
				if subbasincode.Valid {
					aStation = append(aStation, sb)
				}
				province.TeleStation = aStation
			} else {
				sb := &Station{}
				sb.StationCode = subbasincode.String
				sb.StationName = subbasinname.JSON()
				sb.StationID = mid.Int64
				if maxDate.Year() > 1 {
					sb.MaxYear = maxDate.Year()
				}
				if minDate.Year() > 1 {
					sb.MinYear = minDate.Year()
				}
				if subbasincode.Valid {
					aStation = append(aStation, sb)
				}
				province.TeleStation = aStation
			}
		} else {
			sProvinceCode = basincode.String
			province.ProvinceCode = basincode.String
			province.ProvinceName = basinname.JSON()
			sb := &Station{}
			sb.StationCode = subbasincode.String
			sb.StationName = subbasinname.JSON()
			sb.StationID = mid.Int64
			if maxDate.Year() > 1 {
				sb.MaxYear = maxDate.Year()
			}
			if minDate.Year() > 1 {
				sb.MinYear = minDate.Year()
			}
			if subbasincode.Valid {
				aStation = append(aStation, sb)
			}
			province.TeleStation = aStation
		}
	}

	if len(aStation) > 0 {
		province.TeleStation = aStation
		data = append(data, province)
	}

	return data, nil
}

type Struct_ProvinceRegion struct {
	ProvinceId   int64  `json:"province_id"`   // example:`81` รหัสจังหวัด
	RegionId     int64  `json:"region_id"`     // example:`5` รหัสลุ่มแม่น้ำ
	ProvinceName string `json:"province_name"` // exmaple:`กระบี่` ชื่อจังหวัด
}

//	get province with region id
//	Returns:
//		array province region
func GetProvinceRegion() ([]*Struct_ProvinceRegion, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// query
	var q string = `
	SELECT province_code, area_code, province_name->>'th'
	FROM public.lt_geocode
	WHERE amphoe_code = '  ' AND tumbon_code = '  ' OR amphoe_name->>'th' = '' AND tumbon_name->>'th' = ''
	ORDER BY province_name->>'th' ASC
	`
	//	query result
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	//	result function
	var rs []*Struct_ProvinceRegion = make([]*Struct_ProvinceRegion, 0)

	// loop
	for rows.Next() {
		var (
			_province_code sql.NullInt64
			_area_code     sql.NullInt64
			_province_name sql.NullString
		)
		err = rows.Scan(&_province_code, &_area_code, &_province_name)
		if err != nil {
			return nil, err
		}
		// create struct and add data into struct
		s := &Struct_ProvinceRegion{
			ProvinceId:   _province_code.Int64,
			RegionId:     _area_code.Int64,
			ProvinceName: _province_name.String,
		}
		// append struct into array
		rs = append(rs, s)
	}

	return rs, nil
}

// Get geocode id from geocode (string)
//	Parameters:
//		geocode geocode
//	Returns:
//		int geocode id
func GetGeocodeIdFromGeocode(geocode string) (*Struct_Geocode_Id, error) {
	
	// connection
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}
	
	// main query
	q := `
	SELECT id, geocode, province_name->>'th' AS province_name, amphoe_name->>'th' AS amphoe_name, tumbon_name->>'th' AS tumbon_name
	FROM lt_geocode
	WHERE geocode = $1
	`
	
	// parameter
	p := geocode
	
	lt_geocode := &Struct_Geocode_Id{}
	
	err = db.QueryRow(q, p).Scan(&lt_geocode.Id)
	
	if err != nil {
		return nil, errors.Repack(err)
	}

	return lt_geocode, nil
}
