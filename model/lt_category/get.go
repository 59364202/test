package lt_category

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"haii.or.th/api/util/pqx"
)

//	select data
//	Parameters:
//		id
//			รหัสประเภทข้อมูล
//	Return:
//		[]Struct_category
func getCategory(id int64) ([]*Struct_category, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		strSql string

		data     []*Struct_category
		category *Struct_category

		_id            int64
		_category_name sql.NullString
	)
	strSql = SQL_selectCategory
	if id != 0 {
		strSql += " WHERE id = " + strconv.FormatInt(id, 64)
	}

	row, err := db.Query(strSql)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	for row.Next() {
		err = row.Scan(&_id, &_category_name)
		if err != nil {
			return nil, err
		}

		if !_category_name.Valid || _category_name.String == "" {
			_category_name.String = "{}"
		}

		category = &Struct_category{}
		category.Id = _id
		category.Category_name = json.RawMessage(_category_name.String)

		data = append(data, category)
	}

	return data, nil
}

//	select all data
//	Return:
//		[]Struct_category
func GetAllCategory() ([]*Struct_category, error) {
	return getCategory(0)
}

//	select data from id
//	Parameters:
//		id
//			รหัสประเภทข้อมูล
//	Return:
//		[]Struct_category
func GetCategoryFromId(id int64) ([]*Struct_category, error) {
	return getCategory(id)
}
