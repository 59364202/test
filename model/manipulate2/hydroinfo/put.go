package hydroinfo

import (
	"encoding/json"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
)

var sqlUpdateHydroinfo = ` UPDATE lt_hydroinfo
						   SET hydroinfo_name = $2
						     , updated_by = $3
						     , updated_at = NOW()
						   WHERE id = $1 `

var sqlUpdateToDeleteHydroinfo = ` UPDATE lt_hydroinfo
								   SET deleted_by = $2
									 , deleted_at = NOW()
									 , updated_by = $2
									 , updated_at = NOW()
								   WHERE id = $1 `

var sqlUpdateToDeleteHydroinfoAgency = ` UPDATE lt_hydroinfo_agency
										 SET deleted_by = $2
										   , deleted_at = NOW()
										   , updated_by = $2
										   , updated_at = NOW()
										   WHERE hydroinfo_id = $1 `

func PutHydroinfo(userId int64, hydroinfoId string, hydroninfoName json.RawMessage, agencyId string) (*result.Result, error) {
	//Convert hydroinfoId's type from string to int64
	intHydroinfoId, err := strconv.ParseInt(hydroinfoId, 10, 64)
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

	//Update lt_hydroinfo table
	err = updateHydroinfo(tx, hydroinfoId, hydroninfoName, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Delete lt_hydroinfo_agency table
	err = deleteHydroinfoAgency(tx, hydroinfoId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Insert lt_hydroinfo_agency table
	arrAgency, err := insertHydroinfoAgency(tx, hydroinfoId, agencyId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	data := &Hydroinfo_struct{Id: intHydroinfoId, HydroinfoName: hydroninfoName, AgencyId: arrAgency}

	//Return result
	return result.Result1(data), nil
}

//Update lt_hydroinfo table
func updateHydroinfo(tx *pqx.Tx, hydroinfoId string, hydroninfoName json.RawMessage, userId int64) error {
	//Convert hydroninfoName to db-json type
	jsonHydroinfoName, err := hydroninfoName.MarshalJSON()
	if err != nil {
		return err
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlUpdateHydroinfo)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute update statement with parameters
	_, err = statement.Exec(hydroinfoId, jsonHydroinfoName, userId)
	if err != nil {
		return err
	}

	return nil
}

//Update lt_hydroinfo table to set 'Delete'
func updateToDeleteHydroinfo(tx *pqx.Tx, hydroinfoId string, userId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateToDeleteHydroinfo)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(hydroinfoId, userId)
	if err != nil {
		return err
	}

	//Return result
	return nil
}

//Update lt_hydroinfo_agency to set 'Delete'
func updateToDeleteHydroinfoAgency(tx *pqx.Tx, hydroinfoId string, userId int64) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlUpdateToDeleteHydroinfoAgency)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(hydroinfoId, userId)
	if err != nil {
		return err
	}

	//Return result
	return nil
}
