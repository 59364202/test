package hydroinfo

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlDeleteHydroinfoAgency = ` DELETE FROM lt_hydroinfo_agency WHERE hydroinfo_id = $1 `

func DeleteHydroinfo(userId int64, hydroinfoId string) (*result.Result, error) {
	//Try to open database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Repack(err)
	}
	defer tx.Rollback()

	//Check child table of lt_hydroinfo table
	isHasChild, err := checkHydroinfoChild(db, hydroinfoId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return result.Result0(""), nil
	}

	//Update to Delete lt_hydroinfo table
	err = updateToDeleteHydroinfo(tx, hydroinfoId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Update to Delete lt_hydroinfo_agency table
	err = updateToDeleteHydroinfoAgency(tx, hydroinfoId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return result.Result1("Delete Successful."), nil
}

func deleteHydroinfoAgency(tx *pqx.Tx, hydroinfoId string) error {
	//Prepare statement
	statement, err := tx.Prepare(sqlDeleteHydroinfoAgency)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Execute statement with parameters
	_, err = statement.Exec(hydroinfoId)
	if err != nil {
		return err
	}

	//Return result
	return nil
}
