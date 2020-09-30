// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package lt_subcategory is a model for public.lt_subcategory table. This table store lt_subcategory information.
package lt_subcategory

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	get sub category
//	Parameters:
//		categoryId
//			[]รหัสหมวดข้อมูลหลัก
//	Return:
//		[]SubCategory_struct
func GetSubCategory(categoryId []int) ([]*SubCategory_struct, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		data        []*SubCategory_struct
		subcategory *SubCategory_struct

		_id            sql.NullInt64
		_text          sql.NullString
		_category_id   sql.NullInt64
		_category_text sql.NullString
	)
	var _result *sql.Rows
	var _sql = sqlGetSubCategory + sqlGetSubCategoryWhereDeletedIsNull
	for i, v := range categoryId {
		if i == 0 {
			_sql += " AND s.category_id IN ("
		} else {
			_sql += ","
		}
		_sql += strconv.Itoa(v)
	}

	if len(categoryId) == 0 {
		_result, err = db.Query(_sql + sqlGetSubCategoryOrderBy)
	} else {
		_result, err = db.Query(_sql + ") " + sqlGetSubCategoryOrderBy)
	}

	if err != nil {
		return nil, err
	}
	defer _result.Close()

	data = make([]*SubCategory_struct, 0)

	for _result.Next() {
		err := _result.Scan(&_id, &_text, &_category_id, &_category_text)
		if err != nil {
			return nil, err
		}
		if !_text.Valid || _text.String == "" {
			_text.String = "{}"
		}
		if !_category_text.Valid || _category_text.String == "" {
			_category_text.String = "{}"
		}

		subcategory = &SubCategory_struct{Id: _id.Int64, CategoryId: _category_id.Int64}
		subcategory.SubCategoryName = json.RawMessage(_text.String)
		subcategory.CategoryName = json.RawMessage(_category_text.String)

		data = append(data, subcategory)
	}

	return data, nil
}

//	get data
//	Parameters:
//		id
//			รหัสหมวดข้อมูลย่อย
//		category_id
//			รหัสหมวดข้อมูลหลัก
//	Return:
//		[]Struct_subcategory
func getSubcategory(id, category_id int64) ([]*Struct_subcategory, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var (
		strSql string

		data        []*Struct_subcategory
		subcategory *Struct_subcategory

		_id               int64
		_subcategory_name sql.NullString
	)
	strSql = SQL_selectSubcategory
	if id != 0 {
		strSql += " WHERE id = " + strconv.FormatInt(id, 10)
	} else if category_id != 0 {
		strSql += " WHERE category_id = " + strconv.FormatInt(category_id, 10)
	}

	row, err := db.Query(strSql)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_subcategory_name)
		if err != nil {
			return nil, err
		}

		if !_subcategory_name.Valid || _subcategory_name.String == "" {
			_subcategory_name.String = "{}"
		}

		subcategory = &Struct_subcategory{}
		subcategory.Id = _id
		subcategory.Subcategory_name = json.RawMessage(_subcategory_name.String)

		data = append(data, subcategory)
	}

	return data, nil
}

//	get all subcategory
//	Return:
//		[]Struct_subcategory
func GetAllSubcategory() ([]*Struct_subcategory, error) {
	return getSubcategory(0, 0)
}

//	get subcategory from category id
//	Parameters:
//		category_id
//			รหัสหมวดข้อมูลหลัก
//	Return:
//		[]Struct_subcategory
func GetSubcategoryFromCategoryId(category_id int64) ([]*Struct_subcategory, error) {
	return getSubcategory(0, category_id)
}
