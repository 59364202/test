package hydroinfo

import (
	"database/sql"
	"encoding/json"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

var sqlGetHydroinfo = ` SELECT h.id
					 , h.hydroinfo_name
					 , array_to_string(array_agg(ha.agency_id), ', ') AS agency_id
					 , array_to_string(array_agg(ag.agency_name::text), ', ') AS agency_name
					 , array_to_string(array_agg(ag.agency_shortname::text), ', ') AS agency_shortname
					 , hydroinfo_number
				FROM lt_hydroinfo h 
				LEFT JOIN lt_hydroinfo_agency ha ON h.id = ha.hydroinfo_id
				LEFT JOIN agency ag ON ha.agency_id = ag.id `
var sqlGetHydroinfo_Where = ` WHERE h.deleted_at = to_timestamp(0) AND ha.deleted_at = to_timestamp(0) AND ag.deleted_at = to_timestamp(0) `
var sqlGetHydroinfo_GroupBy = ` GROUP BY h.id `
var sqlGetHydroinfo_OrderBy = ` ORDER BY h.id `

func GetHydroinfo(hydroinfoId string, agencyId string) (*result.Result, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		data      []*Hydroinfo_struct1
		hydroinfo *Hydroinfo_struct1

		_id              sql.NullInt64
		_hydroinfoName   sql.NullString
		_agencyId        sql.NullString
		_agencyName      sql.NullString
		_agencyShortName sql.NullString

		_hydroinfo_no sql.NullInt64
	)
	var _result *sql.Rows

	//Set 'sqlCmdWhere' variables
	sqlCmdWhere := sqlGetHydroinfo_Where

	//Check hydroinfo_id
	if hydroinfoId != "" {
		sqlCmdWhere += " AND h.id = " + hydroinfoId
	}

	//Check agency_id
	if agencyId != "" {
		arrAgencyId := strings.Split(agencyId, ",")
		if len(arrAgencyId) == 1 {
			sqlCmdWhere += " AND EXISTS (SELECT agency_id FROM lt_hydroinfo_agency ha2 WHERE ha2.agency_id = " + arrAgencyId[0] + " AND h.id = ha2.hydroinfo_id) "
		} else {
			sqlCmdWhere += " AND EXISTS (SELECT agency_id FROM lt_hydroinfo_agency ha2 WHERE ha2.agency_id IN (" + agencyId + ") AND h.id = ha2.hydroinfo_id) "
		}
	}

	_result, err = db.Query(sqlGetHydroinfo + sqlCmdWhere + sqlGetHydroinfo_GroupBy + sqlGetHydroinfo_OrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Hydroinfo_struct1, 0)

	for _result.Next() {
		//Scan to execute query with variable
		err := _result.Scan(&_id, &_hydroinfoName, &_agencyId, &_agencyName, &_agencyShortName, &_hydroinfo_no)
		if err != nil {
			return nil, err
		}

		//Generate hydroinfo object
		hydroinfo = &Hydroinfo_struct1{}
		hydroinfo.Id = _id.Int64

		if _hydroinfoName.Valid {
			hyn := &Lang{}
			b := []byte(_hydroinfoName.String)
			json.Unmarshal(b, hyn)
			hydroinfo.HydroinfoName = hyn
		}

		arrAgencyId := []int{}
		for _, strAgencyId := range strings.Split(_agencyId.String, ", ") {
			intAgencyId, _ := strconv.Atoi(strAgencyId)
			arrAgencyId = append(arrAgencyId, intAgencyId)
		}
		hydroinfo.AgencyId = arrAgencyId

		if !_agencyName.Valid || _agencyName.String == "" {
			_agencyName.String = "{}"
		}
		arrAgencyName := make([]*Lang, 0)
		for _, strAgencyName := range strings.Split(_agencyName.String, ", ") {
			//if strAgencyName == "" {
			//	strAgencyName = "{}"
			//}
			an := &Lang{}
			b := []byte(strAgencyName)
			json.Unmarshal(b, an)
			arrAgencyName = append(arrAgencyName, an)
		}
		hydroinfo.AgencyName = arrAgencyName

		if !_agencyShortName.Valid || _agencyShortName.String == "" {
			_agencyShortName.String = "{}"
		}
		arrAgencyShortName := make([]*Lang, 0)
		for _, strAgencyShortName := range strings.Split(_agencyShortName.String, ", ") {
			//if strAgencyShortName == "" {
			//	strAgencyShortName = "{}"
			//}
			asn := &Lang{}
			b := []byte(strAgencyShortName)
			json.Unmarshal(b, asn)
			arrAgencyShortName = append(arrAgencyShortName, asn)
		}
		hydroinfo.AgencyShortName = arrAgencyShortName

		data = append(data, hydroinfo)
	}

	return result.Result1(data), nil
}

//Check child table
func checkHydroinfoChild(db *pqx.DB, hydroinfoId string) (bool, error) {
	//Set SQL for check child
	var sqlCheckHydroinfoChild string = ` SELECT id FROM metadata_hydroinfo WHERE hydroinfo_id = $1 LIMIT 1 `

	//Set default of return value
	var isHasChild bool = false

	//Query statement with parameters
	row, err := db.Query(sqlCheckHydroinfoChild, hydroinfoId)
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
