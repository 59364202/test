package shared

import (
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"sort"
	"strings"
	//"log"
)

//	Get list of table that used by metadata
//	Return:
//		Array TableStruct
func GetTable() ([]*TableStruct, error) {
	//Open DB
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Variables
	var (
		listOfTableName []*TableStruct
		objTableName    *TableStruct

		_result     *sql.Rows
		_table_name sql.NullString
	)

	//Query
	_result, err = db.Query(sqlGetTableName)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	listOfTableName = make([]*TableStruct, 0)

	//Loop result
	for _result.Next() {
		_result.Scan(&_table_name)

		objTableName = &TableStruct{}
		objTableName.TableName = _table_name.String

		listOfTableName = append(listOfTableName, objTableName)
	}

	//Return result
	return listOfTableName, nil
}

//	รายชื่อตาราง
//	Return:
//		Array Struct_MetadataTable
func GetMatadataTable() ([]*Struct_MetadataTable, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		resultData       []*Struct_MetadataTable
		objMetadataTable *Struct_MetadataTable

		_table_name  sql.NullString
		_table_desc  sql.NullString
		_column_list sql.NullString

		_result *sql.Rows
	)

	//Query
	//log.Printf(sqlGetMetadataTable)
	_result, err = db.Query(sqlGetMetadataTable)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Loop data result
	resultData = make([]*Struct_MetadataTable, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_table_name, &_table_desc, &_column_list)
		if err != nil {
			return nil, err
		}

		objMetadataTable = &Struct_MetadataTable{}
		objMetadataTable.ID = _table_name.String
		objMetadataTable.TableName = _table_name.String
		objMetadataTable.TableDescription = _table_desc.String

		arrColumns := strings.Split(_column_list.String, ",")
		sort.Strings(arrColumns)
		objMetadataTable.ColumnList = arrColumns

		resultData = append(resultData, objMetadataTable)
	}

	return resultData, nil
}
