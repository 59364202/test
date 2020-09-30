package agency

import (
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

var sqlInsertAgency = "INSERT INTO agency (agency_name, agency_shortname, department_id, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $4, NOW(), NOW()) RETURNING id "

func PostAgency(userId int64, agencyName json.RawMessage, agencyshortName json.RawMessage, departmentId string) (*Agency_struct, error) {
	//Convert departmentId type from string to int64
	intDepartmentId, err := strconv.ParseInt(departmentId, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Try to open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Insert agency table
	newId, err := insertAgency(tx, agencyName, agencyshortName, departmentId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	data := &Agency_struct{Id: newId, AgencyName: agencyName, AgencyShortName: agencyshortName, DepartmentId: intDepartmentId}
	return data, nil
}

//Insert to agency table
func insertAgency(tx *pqx.Tx, agencyName json.RawMessage, agencyshortName json.RawMessage, departmentId string, userId int64) (int64, error) {
	var _id int64

	//Convert agencyName to db-json type
	jsonAgencyName, err := agencyName.MarshalJSON()
	if err != nil {
		return 0, err
	}

	//Convert agencyshortName to db-json type
	jsonAgencyShortName, err := agencyshortName.MarshalJSON()
	if err != nil {
		return 0, err
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertAgency)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(string(jsonAgencyName[:]), string(jsonAgencyShortName[:]), departmentId, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}
