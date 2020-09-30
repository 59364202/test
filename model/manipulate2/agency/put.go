package agency

import (
	"encoding/json"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

var sqlUpdateAgency = ` UPDATE agency
					    SET agency_name = $2
					      , agency_shortname = $3
					      , department_id = $4
					      , updated_by = $5
					      , updated_at = NOW()
					    WHERE id = $1 `

var sqlUpdateToDeleteAgency = ` UPDATE agency
								SET deleted_by = $2
								  , deleted_at = NOW()
								  , updated_by = $2
								  , updated_at = NOW()
								WHERE id = $1 `

func PutAgency(userId int64, agencyId string, agencyName json.RawMessage, agencyShortName json.RawMessage, departmentId string) (*Agency_struct, error) {
	//Convert agencyId type from string to int64
	intAgencyId, err := strconv.ParseInt(agencyId, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Convert departmentId type from string to int64
	intDepartmentId, err := strconv.ParseInt(departmentId, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Open database
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

	//Update agency table
	err = updateAgency(tx, agencyId, agencyName, agencyShortName, departmentId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	data := &Agency_struct{Id: intAgencyId, AgencyName: agencyName, AgencyShortName: agencyShortName, DepartmentId: intDepartmentId}

	//Return result
	return data, nil
}

//Update agency table
func updateAgency(tx *pqx.Tx, agencyId string, agencyName json.RawMessage, agencyShortName json.RawMessage, departmentId string, userId int64) error {
	//Convert agencyName to db-json type
	jsonAgencyName, err := agencyName.MarshalJSON()
	if err != nil {
		return err
	}

	//Convert agencyShortName to db-json type
	jsonAgencyShortName, err := agencyShortName.MarshalJSON()
	if err != nil {
		return err
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateAgency)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(agencyId, jsonAgencyName, jsonAgencyShortName, departmentId, userId)
	if err != nil {
		return err
	}

	return nil
}

//Update agency table to set 'Delete'
func updateToDeleteAgency(tx *pqx.Tx, agencyId string, userId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateToDeleteAgency)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(agencyId, userId)
	if err != nil {
		return err
	}

	//Return result
	return nil
}
