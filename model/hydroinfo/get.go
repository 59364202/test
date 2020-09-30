package hydroinfo

import (
	"database/sql"
	"encoding/json"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

//	Get all hydroinfo
//	Return:
//		[]Struct_Hydroinfo
func GetAllHydroinfo() ([]*Struct_Hydroinfo, error) {
	return getHydroinfo("", 0)
}

//	Get hydroinfo ของบัญชีข้อมูล
//	Parameters:
//		metadataID
//			รหัสบัญชีข้อมูล
//	Return:
//		[]Struct_Hydroinfo
func GetHydroinfoByMetadata(metadataID int64) ([]*Struct_Hydroinfo, error) {
	return getHydroinfo("", metadataID)
}

//	get hydroinfo ตาม hydroinfoID, metadataID
//	Parameters:
//		hydroinfoID
//			รหัสกรมทรัพยากรน้ำ
//		metadataID
//			รหัสบัญชีข้อมูล
//	Return:
//		[]Struct_Hydroinfo
func getHydroinfo(hydroinfoID string, metadataID int64) ([]*Struct_Hydroinfo, error) {

	//Open database
	db, err := pqx.Open()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	//Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer tx.Rollback()

	//Variables
	var (
		data         []*Struct_Hydroinfo
		objHydroinfo *Struct_Hydroinfo

		arrObjAgency []*model_agency.Struct_Agency
		objAgency    *model_agency.Struct_Agency

		intHydroInfoID int64 = 0

		_id               sql.NullInt64
		_name             sql.NullString
		_hydroinfo_number sql.NullInt64
		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_result *sql.Rows
	)

	var arrParam = make([]interface{}, 0)
	var sqlQuery string = ""

	if metadataID == 0 {
		//-- Check Filter by parameters --//
		var sqlCmdWhere string = ""
		arrHydroinfoID := []string{}

		if hydroinfoID != "" {
			arrHydroinfoID = strings.Split(hydroinfoID, ",")
		}

		//Check Filter hydroinfoID
		if len(arrHydroinfoID) > 0 {
			if len(arrHydroinfoID) == 1 {
				arrParam = append(arrParam, strings.Trim(hydroinfoID, " "))
				sqlCmdWhere += " AND h.id = $" + strconv.Itoa(len(arrParam))
			} else {
				arrSqlCmd := []string{}
				for _, strId := range arrHydroinfoID {
					arrParam = append(arrParam, strings.Trim(strId, " "))
					arrSqlCmd = append(arrSqlCmd, "$"+strconv.Itoa(len(arrParam)))
				}
				sqlCmdWhere += " AND h.id IN (" + strings.Join(arrSqlCmd, ",") + ")"
			}
		}

		sqlQuery = sqlGetHydroinfo + sqlCmdWhere
	} else {
		sqlQuery = sqlGetHydroinfoByMetadata
		arrParam = append(arrParam, metadataID)
	}

	//Query
	//log.Printf(sqlQuery+" ORDER BY h.hydroinfo_name->>'th', ag.agency_name->>'th' ", arrParam...)
	//_result, err = db.Query(sqlQuery + " ORDER BY h.hydroinfo_name->>'th', ag.agency_name->>'th' ", arrParam...)
	_result, err = db.Query(sqlQuery+" ORDER BY h.hydroinfo_number ", arrParam...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	data = make([]*Struct_Hydroinfo, 0)

	// Loop data result
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_name, &_hydroinfo_number, &_agency_id, &_agency_shortname, &_agency_name)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}

		if intHydroInfoID != _id.Int64 {
			if intHydroInfoID != 0 {
				objHydroinfo.Agency = arrObjAgency
				data = append(data, objHydroinfo)
			}

			if !_name.Valid || _name.String == "" {
				_name.String = "{}"
			}
			objHydroinfo = &Struct_Hydroinfo{}
			objHydroinfo.ID = _id.Int64
			objHydroinfo.HydroinfoName = json.RawMessage(_name.String)
			objHydroinfo.HydroinfoNumber = ValidData(_hydroinfo_number.Valid, _hydroinfo_number.Int64)

			arrObjAgency = make([]*model_agency.Struct_Agency, 0)
			intHydroInfoID = _id.Int64
		}

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		objAgency = &model_agency.Struct_Agency{}
		objAgency.Id = _agency_id.Int64
		objAgency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		objAgency.Agency_name = json.RawMessage(_agency_name.String)

		arrObjAgency = append(arrObjAgency, objAgency)
	}

	if intHydroInfoID != 0 {
		objHydroinfo.Agency = arrObjAgency
		data = append(data, objHydroinfo)
	}

	//Return data
	return data, nil
}

// check valid data return valid return data, invalid return null
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
