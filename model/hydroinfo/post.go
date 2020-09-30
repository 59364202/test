package hydroinfo

import (
	//model_agency "haii.or.th/api/thaiwater30/model/agency"
	"encoding/json"
	"strconv"
	"strings"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

//	insert data
//	Parameters:
//		userId
//			รหัสผู้ใช้
//		param
//			parameter ที่หน้าจอส่งมา
//	Return:
//		Inserted Successful
func PostHydroInfo(userId int64, param *Struct_Hydroinfo_InputParam) (string, error) {

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

	//Insert lt_hydroinfo table
	newId, err := insertHydroinfo(tx, param.HydroInfoName, param.HydroinfoNumber, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Insert lt_hydroinfo_agency table
	err = insertHydroinfoAgency(tx, newId, param.AgencyID, userId)
	if err != nil {
		return "", errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	//data := &Struct_Hydroinfo{ID: newId, HydroinfoName: hydroninfoName, HydroinfoNumber: hydroinfoNumber}
	return "Inserted Successful", nil
}

//	Insert data to lt_hydroinfo table
//	Parameters:
//		tx
//			Transaction
//		hydroninfoName
//			ชื่อด้าน
//		hydroinfoNumber
//			ลำดับของชือ เพื่อใช้ในการนำเสนอ
//		userId
//			รหัสผู้ใช้
//	Return:
//		รหัสข้อมูลที่เพิ่ม
func insertHydroinfo(tx *pqx.Tx, hydroninfoName json.RawMessage, hydroninfoNumber int64, userId int64) (int64, error) {
	var _id int64

	//Convert hydroninfoName to db-json type
	jsonHydroinfoName, err := hydroninfoName.MarshalJSON()
	if err != nil {
		return 0, err
	}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertHydroinfo)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//Execute insert statement with parameters and returning id
	err = statement.QueryRow(string(jsonHydroinfoName[:]), hydroninfoNumber, userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}

//	Insert to lt_hydroinfo_agency table
//	Parameters:
//		tx
//			Transaction
//		hydroinfoId
//			รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
//		agencyId
//			รหัสหน่วยงาน
//		userId
//			รหัสผู้ใช้
//	Return:
//		nil, error
func insertHydroinfoAgency(tx *pqx.Tx, hydroinfoId int64, agencyId string, userId int64) error {

	//Set arrAgencyId for return array of agency_id
	arrAgencyId := []int{}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertHydroinfoAgency)
	if err != nil {
		return err
	}
	defer statement.Close()

	//Loop by number of agencyId
	for _, strAgencyId := range strings.Split(agencyId, ",") {

		//Convert strAgencyId's type that string to integer
		intAgencyId, err := strconv.Atoi(strings.TrimSpace(strAgencyId))
		if err != nil {
			return err
		}

		//Execute insert statement with parameters
		_, err = statement.Exec(hydroinfoId, intAgencyId, userId)
		if err != nil {
			return err
		}

		//Append agency_id to arrAgencyId
		arrAgencyId = append(arrAgencyId, intAgencyId)
	}

	return nil
}
