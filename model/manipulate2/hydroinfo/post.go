package hydroinfo

import (
	"encoding/json"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

var sqlInsertHydroinfo = "INSERT INTO lt_hydroinfo (hydroinfo_name, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $2, NOW(), NOW()) RETURNING id "
var sqlInsertHydroinfoAgency = "INSERT INTO lt_hydroinfo_agency (hydroinfo_id, agency_id, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $3, NOW(), NOW()) "

func PostHydroInfo(userId int64, hydroninfoName json.RawMessage, agencyId string) (*result.Result, error) {

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

	//Insert lt_hydroinfo table
	newId, err := insertHydroinfo(tx, hydroninfoName, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Insert lt_hydroinfo_agency table
	arrAgency, err := insertHydroinfoAgency(tx, strconv.FormatInt(newId, 10), agencyId, userId)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Commit Transaction
	tx.Commit()

	//Return data
	data := &Hydroinfo_struct{Id: newId, HydroinfoName: hydroninfoName, AgencyId: arrAgency}
	return result.Result1(data), nil
}

//Insert to lt_hydroinfo table
func insertHydroinfo(tx *pqx.Tx, hydroninfoName json.RawMessage, userId int64) (int64, error) {
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
	err = statement.QueryRow(string(jsonHydroinfoName[:]), userId).Scan(&_id)
	if err != nil {
		return 0, err
	}

	return _id, nil
}

//Insert to lt_hydroinfo_agency table
func insertHydroinfoAgency(tx *pqx.Tx, hydroinfoId string, agencyId string, userId int64) ([]int, error) {

	//Set arrAgencyId for return array of agency_id
	arrAgencyId := []int{}

	//Prepare Statement
	statement, err := tx.Prepare(sqlInsertHydroinfoAgency)
	if err != nil {
		return []int{}, err
	}
	defer statement.Close()

	//Loop by number of agencyId
	for _, strAgencyId := range strings.Split(agencyId, ",") {

		//Convert strAgencyId's type that string to integer
		intAgencyId, err := strconv.Atoi(strings.TrimSpace(strAgencyId))
		if err != nil {
			return []int{}, err
		}

		//Execute insert statement with parameters
		_, err = statement.Exec(hydroinfoId, intAgencyId, userId)
		if err != nil {
			return []int{}, err
		}

		//Append agency_id to arrAgencyId
		arrAgencyId = append(arrAgencyId, intAgencyId)
	}

	return arrAgencyId, nil
}
