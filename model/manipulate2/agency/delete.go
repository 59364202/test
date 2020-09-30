package agency

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
)

func DeleteAgency(userId int64, agencyId string) (string, error) {
	//Open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	//Check child table of agency
	isHasChild, err := checkAgencyChild(db, agencyId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Can't delete this data. It's has been used.
	if isHasChild {
		return "", rest.NewError(422, "ไม่สามารถลบหน่วยงานนี้ได้ เนื่องจากถูกใช้งานอยู่", nil)
	}

	//Update to Delete agency table
	err = updateToDeleteAgency(tx, agencyId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful.", nil
}

//Check child table of agency
func checkAgencyChild(db *pqx.DB, agencyId string) (bool, error) {
	//Set SQL for check child
	var sqlCheckChild string = ` SELECT agency_id FROM metadata WHERE agency_id = $1
	  UNION
	  SELECT agency_id FROM lt_hydroinfo_agency WHERE agency_id = $1 `

	//Set default of return value
	var isHasChild bool = false

	//Query statement with parameters
	row, err := db.Query(sqlCheckChild, agencyId)
	if err != nil {
		return isHasChild, err
	}

	//Check child
	for row.Next() {
		isHasChild = true
	}

	//Return result
	return isHasChild, nil
}
