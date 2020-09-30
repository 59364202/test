package hydroinfo

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	Delete hydroinfo (soft delete)
//	Parameters:
//		userId
//			รหัสผู้ใช้ที่เป็นคน delete
//		hydroinfoId
//			รหัสกรมทรัพยากรน้ำ ที่ต้องการ delete
//	Return:
//		Delete Successful
func DeleteHydroinfo(userId int64, hydroinfoId int64) (string, error) {
	//Try to open database
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

	//Check child table of lt_hydroinfo table
	//isHasChild, err := checkHydroinfoChild(db, hydroinfoId)
	//if err != nil {
	//	return nil, errors.Repack(err)
	//}

	//Can't delete this data. It's has been used.
	//if isHasChild {
	//	return result.Result0("Can't Delete."), nil
	//}

	//Update to Delete lt_hydroinfo table
	err = updateToDeleteHydroinfo(tx, hydroinfoId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Update to Delete lt_hydroinfo_agency table
	err = updateToDeleteHydroinfoAgency(tx, hydroinfoId, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Return result
	return "Delete Successful", nil
}

//	Delete lt_hydroinfo_agency table
//	Parameters:
//		tx
//			Transaction
//		hydroinfoId
//			รหัสกรมทรัพยากรน้ำ ที่ต้องการ delete
//	Return:
//		nil, error
func deleteHydroinfoAgency(tx *pqx.Tx, hydroinfoId int64) error {
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
