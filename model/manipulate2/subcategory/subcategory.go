package subcategory

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

//var sqlGetSubCategory = " SELECT  t.subcategory_id,t.subcategory_name ,tc.category_id ,tc.category_name ,l.code  " +
//	" FROM language l  " +
//	" LEFT JOIN lt_category_translate tc ON tc.language_id = l.id " +
//	" LEFT JOIN lt_subcategory lt ON lt.category_id = tc.category_id " +
//	" LEFT JOIN lt_subcategory_translate t ON t.subcategory_id = lt.id AND t.language_id = l.id "

var sqlGetSubCategory = "SELECT s.id , s.subcategory_name, c.id , c.category_name " +
	" FROM lt_subcategory s " +
	" INNER JOIN lt_category c ON s.category_id = c.id and c.deleted_at = '1970-01-01 07:00:00+07' "

//var sqlGetSubCategoryWhereCategoryId = " WHERE s.category_id IN ($1) AND s.deleted_by IS NULL "
var sqlGetSubCategoryWhereDeletedIsNull = " WHERE s.deleted_at = '1970-01-01 07:00:00+07' "
var sqlGetSubCategoryOrderBy = " ORDER BY s.id , c.id "

var sqlInsertSubCategory = " INSERT INTO lt_subcategory (category_id , created_by , subcategory_name) VALUES($1 , $2 , $3) RETURNING id "

var sqlUpdateSubCategory = " UPDATE lt_subcategory SET category_id = $1 , updated_by = $2 , subcategory_name = $3 , updated_at = NOW() WHERE id = $4 "

var sqlCheckChild = " SELECT id FROM metadata WHERE subcategory_id = $1 AND deleted_by IS NULL LIMIT 1 "
var sqlUpdateToDeleteSubCategory = " UPDATE lt_subcategory SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "

type SubCategory_struct struct {
	Id              int64           `json:"id"`
	SubCategoryName json.RawMessage `json:"subcategory_name,omitempty"` // example:`{"th": "ภาษาไทย", "en": "english"}`
	CategoryId      interface{}     `json:"category_id"`
	CategoryName    json.RawMessage `json:"category_name,omitempty"`
}

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

	//	if categoryId == "" {
	//		_result, err = db.Query(sqlGetSubCategory + sqlGetSubCategoryWhereDeletedIsNull + sqlGetSubCategoryOrderBy)
	//	} else {
	//		var _sql = sqlGetSubCategory + sqlGetSubCategoryWhereDeletedIsNull + " AND s.category_id IN ("
	//
	//		var _param = make([]interface{}, 0)
	//		var countParam = 1
	//		var arrayCategoryId = strings.Split(categoryId, ",")
	//		for i, v := range arrayCategoryId {
	//			_sql += "$" + strconv.Itoa(countParam)
	//			if i < len(arrayCategoryId)-1 {
	//				_sql += ","
	//			} else {
	//				_sql += ") "
	//			}
	//
	//			_param = append(_param, v)
	//			countParam++
	//		}
	//
	//		_result, err = db.Query(_sql, _param...)
	//	}
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

func PostSubCategory(userId int64, categoryId string, mapTxt json.RawMessage) (*SubCategory_struct, error) {
	db, err := pqx.Open()

	if err != nil {
		return nil, errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return nil, errors.Repack(err)
	}
	var (
		_id int64
	)
	err = db.QueryRow(sqlInsertSubCategory, categoryId, userId, jsonText).Scan(&_id)
	if err != nil {
		return nil, err
	}

	data := &SubCategory_struct{Id: _id, SubCategoryName: mapTxt}
	return data, nil
}

func PutSubCategory(subCategoryId string, userId int64, categoryId string, mapTxt json.RawMessage) (string, error) {
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	jsonText, err := mapTxt.MarshalJSON()
	if err != nil {
		return "", errors.Repack(err)
	}

	_, err = db.Exec(sqlUpdateSubCategory, categoryId, userId, jsonText, subCategoryId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "update success", nil
}

func DeleteSubCategory(subcategoryId string, userId int64) (string, error) {
	hasChild := false
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	row, err := db.Query(sqlCheckChild, subcategoryId)
	if err != nil {
		return "", pqx.GetRESTError(err)
	}
	for row.Next() {
		hasChild = true
	}

	if hasChild {
		return "", rest.NewError(422, "หมวดหมู่ย่อยได้ถูกใช้โดย บัญชีข้อมูล", nil)
	}

	_, err = db.Exec(sqlUpdateToDeleteSubCategory, subcategoryId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	return "delete success", nil
}
