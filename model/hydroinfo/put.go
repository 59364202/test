package hydroinfo

import (
	"encoding/json"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	update lt_hydroinfo table
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		param
//			parameter ที่หน้าจอส่งมา
//	Return:
//		Updated Successful
func PutHydroinfo(userId int64, param *Struct_Hydroinfo_InputParam) (string, error) {

	//Try to open database
	db, err := pqx.Open()
	if err != nil {
		return "", errors.Repack(err)
	}

	//Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Repack(err)
	}
	defer tx.Rollback()

	// Update lt_hydroinfo table
	err = updateHydroinfo(tx, param.ID, param.HydroInfoName, param.HydroinfoNumber, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Delete lt_hydroinfo_agency table
	err = deleteHydroinfoAgency(tx, param.ID)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Insert lt_hydroinfo_agency table
	err = insertHydroinfoAgency(tx, param.ID, param.AgencyID, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit transaction
	tx.Commit()

	//Set object fot return result
	//data := &Struct_Hydroinfo{ID: hydroinfoId, HydroinfoName: hydroninfoName, HydroinfoNumber: hydroinfoNumber}

	//Return result
	return "Updated Successful", nil
}

//	Update lt_hydroinfo table
//	Parameters:
//		tx
//			Transaction
//		hydroinfoId
//			รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
//		hydroninfoName
//			ชื่อด้าน
//		hydroinfoNumber
//			ลำดับของชือ เพื่อใช้ในการนำเสนอ
//		userId
//			รหัสผู้ใช้
//	Return:
//		nil, error
func updateHydroinfo(tx *pqx.Tx, hydroinfoId int64, hydroninfoName json.RawMessage, hydroinfoNumber int64, userId int64) error {
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
	_, err = statement.Exec(hydroinfoId, jsonHydroinfoName, hydroinfoNumber, userId)
	if err != nil {
		return err
	}

	return nil
}

//  Update lt_hydroinfo table to set 'Delete'
//	Parameters:
//		tx
//			Transaction
//		hydroinfoId
//			รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
//		hydroninfoName
//			ชื่อด้าน
//		hydroinfoNumber
//			ลำดับของชือ เพื่อใช้ในการนำเสนอ
//		userId
//			รหัสผู้ใช้
//	Return:
//		nil, error
func updateToDeleteHydroinfo(tx *pqx.Tx, hydroinfoId int64, userId int64) error {
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

//	Update lt_hydroinfo_agency to set 'Delete'
//	Parameters:
//		tx
//			Transaction
//		hydroinfoId
//			รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
//		userId
//			รหัสผู้ใช้
//	Return:
//		nil, error
func updateToDeleteHydroinfoAgency(tx *pqx.Tx, hydroinfoId int64, userId int64) error {
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
