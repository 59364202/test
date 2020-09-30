// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package subcategory is a model for public.subcategory table. This table store subcategory information.
package subcategory

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	get subcategory
//	Parameters:
//		categoryId
//			รหัสหมวดหมู่หลัก
//	Return:
//		Array SubCategory_struct
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
