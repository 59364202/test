// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

// package for model สำหรับอบรม
package a

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	//	"log"
	//	"fmt"
	"strconv"
	//สำหรับรับค่าที่เป็น null จากฐานข้อมูล
	"database/sql"
	"encoding/json"
)

//ตัวแปร ขึ้นต้นด้วยตัวใหญ่คือ public
//ขึ้นต้นด้วยตัวเล็กคือ private

//struc for return json
//`json:"-"` หมายถึงไม่ return json นั้น

//struct ที่จะ return ถ้าต้องการ return null ต้องเป็น inteface เท่านนั้น
//interface สามารถเก็ยข้อมูลได้ทุกประเภท เก็บ array ได้
//	Region_id   int64 `json:"region_id"`
//	Province_name string `json:"-"`
//ถ้าไม่ใส่ `json:"region_id"` จะ return struct ตามชื่อตัวแปร Region_id
type thitipornProvinceStruc struct {
	Region_id          interface{}     `json:"region_id"`          // example:`81` รหัสจงหวัด
	Province_id        int64           `json:"province_id"`        // example:`5` รหัสภาค
	Province_name      string          `json:"province_name"`      // example:`กระบี่` ชื่อจังหวัด
	Province_name_json json.RawMessage `json:"province_name_json"` // example:`{"th" : "กระบี่ "}` ชื่อจังหวัดแบบ json มีได้หลายภาษา
}

//for swagger api document model description
// example:`81` รหัสจงหวัด

//sturc for parameter input from model and service only query string parameter
//api2.thaiwater.net/test/thitiporn?region_id=5&province_id=81
//province_id ชื่อ json ต้องตรงกับ query string ที่ส่งมาจาก service
type Param_handlerGetProvince struct {
	Province_id string `json:"province_id"` // example:`5` required:true รหัสภาค
	Region_id   string `json:"region_id"`   // example:81   required:false รหัสจังหวัด
}

// no parameter return struct,error
//func Thitiporn_Province() ([]*thitipornProvinceStruc, error) {

//รับ parameter เป็น string
//func Thitiporn_Province(region_code string,prov_code string) ([]*thitipornProvinceStruc, error) {

//function for query data
//return []*provinceStruct -> array pointer of province struct
func Thitiporn_Province(param *Param_handlerGetProvince) ([]*thitipornProvinceStruc, error) {
	db, err := pqx.Open()
	if err != nil {
		//	บันทึก error ใน log และ return จาก function
		return nil, errors.Repack(err)
	}

	//	ถ้าต้องการให้ใส่ทั้ง 2 parameter ถ้าไม่ใส่ต้องการให้ return error
	//	if param.Province_id == "" || param.Region_id == "" {
	//		return nil, errors.New("invalid param")
	//	}

	//parameter interface
	paramInterface := []interface{}{}

	//  ถ้า query return null ถ้าตัวปรกที่มารับค่าเป็น string or int จะเกิด error
	//  ถ้าต้องการให้แสดงค่าได้ต้องเปลี่ยน struct ตัวแปร เป็น string
	//  หรือ เปลี่ยนประเภท ตัวแปรใน struct เป็น interface

	// query
	//	SELECT province_code, area_code, province_name->>'th'
	var queryString string = `
	SELECT null, null, province_name->>'th',province_name
	FROM public.lt_geocode
	WHERE 
	(
		(amphoe_code = '  ' AND tumbon_code = '  ') 
		OR (amphoe_name->>'th' = '' AND tumbon_name->>'th' = '')
	)`

	//sql injection $1 = parameter 1
	//	if region_ocde != "" {
	//		q += " AND region_code = $1 "
	//	}

	//sql injection $2 = parameter 2
	//	if prov_code != "" {
	//		q += " AND province_code = $2 "
	//	}

	//string convert -> interger to string range of paramInterface
	//strconv.Itoa(len(paramInterface))
	//ตรวจสอบลำดับของ parameter array ถ้าไม่ใส่ region id มา province_id จะเป็นลำดับ array ที่ 1
	if param.Region_id != "" {
		paramInterface = append(paramInterface, param.Region_id)
		queryString += " AND area_code = $" + strconv.Itoa(len(paramInterface))

	}
	if param.Province_id != "" {
		paramInterface = append(paramInterface, param.Province_id)
		queryString += " AND province_code = $" + strconv.Itoa(len(paramInterface))
	}

	queryString += " ORDER BY province_name->>'th' ASC "

	rows, err := db.Query(queryString, paramInterface...)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//	กำหนดตัวแปร array เพื่อจะได้ใส่ค่าใน struct ได้แบบไม่กำหนด range
	returnStruc := make([]*thitipornProvinceStruc, 0)

	for rows.Next() {
		//		"database/sql"
		//	"encoding/json"

		//		กำหนดตัวแปรเพื่อมารับค่าจากฐานข้อมูล
		//		ถ้ากำหนดแบบนี้ แล้วค่าที่ ได้จากฐานข้อมูล เป็น null จะมี error
		//		var (
		//			_province_code int64
		//			_area_code     int64
		//			_province_name string
		//		)

		var (
			_province_code      sql.NullInt64
			_area_code          sql.NullInt64
			_province_name      sql.NullString
			_province_name_json pqx.JSONRaw
		)

		//		rows.Scan เอาค่าจากฐานข้อมูลมาใส่ตัวแปร
		err = rows.Scan(&_province_code, &_area_code, &_province_name, &_province_name_json)
		if err != nil {
			return nil, errors.Repack(err)
		}

		// append data object to array
		p := &thitipornProvinceStruc{
			Province_id: _province_code.Int64,
			//			Region_id:     _area_code.Int64,
			Province_name:      _province_name.String,
			Province_name_json: _province_name_json.JSON(),
		}

		//		p.Province_name_json = _province_name_json.JSON()

		//		ตรวจสอบว่าข้อมูลเป็น null หรือไม่
		if !_province_name.Valid {
			_province_name.String = "{}"
		}
		//		p.Province_name_json = json.RawMessage(_province_name.String)

		//	การกำหนดค่าให้ interface ทำได้ 2 วิธี
		//	1. ใช้ Value()
		//		assign value to interface
		p.Region_id, _ = _area_code.Value()

		// 2. เช็คว่า ถ้าไม่ใส่ null ให้ใส่ค่า
		if _area_code.Valid {
			p.Region_id = _area_code.Int64
		}

		returnStruc = append(returnStruc, p)
	}
	return returnStruc, err
}
