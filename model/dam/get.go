package dam

import (
	"database/sql"
	"encoding/json"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"strconv"
	"strings"
)

//	get dam
//	Parameters:
//		dam_id
//			dam id
//		agency_id
//			agency id
//	Return:
//		[]Struct_GetDam
func GetDam(dam_id, agency_id string) ([]*Struct_GetDam, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	if agency_id == "" {
		agency_id = "12"
	}
	var (
		data []*Struct_GetDam
		dam  *Struct_GetDam

		_result *sql.Rows

		_id          int64
		_geocode_id  int64
		_agency_id   int64
		_dam_lat     sql.NullFloat64
		_dam_long    sql.NullFloat64
		_dam_name    sql.NullString
		_max_storage sql.NullFloat64
		_min_storage sql.NullFloat64
		_agency_name sql.NullString
	)
	//	query
	if dam_id == "" {
		_result, err = db.Query(sqlGetDam+" m_dam.agency_id = $1 "+sqlGetDam_Order, agency_id)
	} else {
		_result, err = db.Query(sqlGetDam+" m_dam.id = $1 "+sqlGetDam_Order, dam_id)
	}

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	data = make([]*Struct_GetDam, 0)
	for _result.Next() {
		_result.Scan(&_id, &_geocode_id, &_agency_id, &_dam_lat, &_dam_long, &_dam_name, &_max_storage, &_min_storage, &_agency_name)

		if _dam_name.String == "" || !_dam_name.Valid {
			_dam_name.String = "{}"
		}

		if _agency_name.String == "" || !_agency_name.Valid {
			_agency_name.String = "{}"
		}

		dam = &Struct_GetDam{}
		dam.Id = _id

		objAgency := &model_agency.Struct_A{}
		objAgency.Id = _agency_id
		objAgency.Agency_name = json.RawMessage(_agency_name.String)
		dam.Agency = objAgency

		dam.Dam_lat, _ = _dam_lat.Value()
		dam.Dam_long, _ = _dam_long.Value()
		dam.Dam_name = json.RawMessage(_dam_name.String)
		dam.Max_storage = _max_storage.Float64
		dam.Min_storage = _min_storage.Float64

		data = append(data, dam)
	}

	return data, nil
}

//	get dam จาก dam id
//	Parameters:
//		dam_id
//			dam id
//	Return:
//		[]Struct_GetDam
func GetDamFromId(dam_id string) ([]*Struct_GetDam, error) {
	return GetDam(dam_id, "")
}

//	get dam จาก agency id
//	Parameters:
//		agency_id
//			agency id
//	Return:
//		[]Struct_GetDam
func GetDamFromAgency(agency_id string) ([]*Struct_GetDam, error) {
	return GetDam("", agency_id)
}

//
//func GetDamDatatype() ([]map[string]string, error) {
//	data := make([]map[string]string, 0)
//
//	m := map[string]string{
//		"text":  "น้ำไหลลงอ่าง",
//		"value": "dam_inflow",
//	}
//	data = append(data, m)
//
//	m = map[string]string{
//		"text":  "ปริมาตรน้ำในอ่างฯ",
//		"value": "dam_storage",
//	}
//	data = append(data, m)
//
//	m = map[string]string{
//		"text":  "ใช้การได้จริง",
//		"value": "dam_uses_water",
//	}
//	data = append(data, m)
//
//	m = map[string]string{
//		"text":  "น้ำระบาย",
//		"value": "dam_released",
//	}
//
//	data = append(data, m)
//	return data, nil
//}

func GetDamStationByDataType(param *Struct_GetDam_InputParam) (*result.Result, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		//data      		[]*DamStruct
		//objDamStation 	*DamStruct

		data          []*Struct_GetDam_OutputParam
		objDamStation *Struct_GetDam_OutputParam

		_id              sql.NullInt64
		_geocode_id      sql.NullInt64
		_agency_id       sql.NullInt64
		_oldcode         sql.NullString
		_lat             sql.NullFloat64
		_long            sql.NullFloat64
		_name            sql.NullString
		_max_water_level sql.NullFloat64
		_min_water_level sql.NullFloat64
		_max_storage     sql.NullFloat64
		_min_storage     sql.NullFloat64

		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_result *sql.Rows
	)

	//Set 'sqlCmdWhere' variables
	arrSqlCmdWhere := []string{}
	strSqlCmdWhere := ""

	switch param.DataType {
	case "dam_daily":
		arrSqlCmdWhere = append(arrSqlCmdWhere, sqlConditionDamStationByDaily)
	case "dam_hourly":
		arrSqlCmdWhere = append(arrSqlCmdWhere, sqlConditionDamStationByHourly)
	default:

	}

	if len(arrSqlCmdWhere) > 0 {
		strSqlCmdWhere = " WHERE " + strings.Join(arrSqlCmdWhere, " AND ")
	}

	_result, err = db.Query(sqlGetDamStationByDataType + strSqlCmdWhere + sqlGetDamStationByDataTypeOrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	//data = make([]*DamStruct, 0)
	data = make([]*Struct_GetDam_OutputParam, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_geocode_id, &_agency_id, &_lat, &_long, &_oldcode, &_name, &_max_water_level, &_min_water_level, &_max_storage, &_min_storage, &_agency_shortname, &_agency_name)
		if err != nil {
			return nil, err
		}

		//Generate DamStation object
		/*
			objDamStation = &DamStruct{}
			objDamStation.Dam_id = _id.Int64
			objDamStation.District_id = _district_id.Int64
			objDamStation.Agency_id = _agency_id.Int64
			objDamStation.Dam_lat = _lat.String
			objDamStation.Dam_long = _long.String
			objDamStation.Dam_name = json.RawMessage(_name.String)
			objDamStation.Dam_oldcode = _oldcode.String
			objDamStation.Max_water_level = _max_water_level.Float64
			objDamStation.Min_water_level = _max_water_level.Float64
			objDamStation.Max_storage = _max_storage.Float64
			objDamStation.Min_storage = _max_storage.Float64

			if (_name.String == ""){_name.String = "{}"}
			objDamStation.Dam_name = json.RawMessage(_name.String)

			data = append(data, objDamStation)
		*/

		objDamStation = &Struct_GetDam_OutputParam{}
		objDamStation.Id = _id.Int64

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		if !_name.Valid || _name.String == "" {
			_name.String = "{}"
		}
		objDamStation.Station_name = json.RawMessage(_name.String)
		objDamStation.Station_oldcode = _oldcode.String
		objDamStation.Agency_name = json.RawMessage(_agency_name.String)
		objDamStation.Agency_shortname = json.RawMessage(_agency_shortname.String)

		data = append(data, objDamStation)
	}

	return result.Result1(data), nil
}

func GetDamGroupByAgency() ([]*Struct_DamGroupByAgency, error) {

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		data   []*Struct_DamGroupByAgency
		objDam *Struct_DamGroupByAgency

		_dam sql.NullString

		_agency_id        sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString

		_result *sql.Rows
	)

	_result, err = db.Query(sqlGetDamGroupbyAgency)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Struct_DamGroupByAgency, 0)
	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_agency_id, &_agency_shortname, &_agency_name, &_dam)
		if err != nil {
			return nil, err
		}

		//Generate DamStation object
		objDam = &Struct_DamGroupByAgency{}

		if !_agency_shortname.Valid || _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		if !_agency_name.Valid || _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		objDam.Agency = &model_agency.Struct_Agency{}
		objDam.Agency.Id = _agency_id.Int64
		objDam.Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		objDam.Agency.Agency_name = json.RawMessage(_agency_name.String)

		if _dam.String == "" {
			objDam.Dam = nil
		} else {
			arrDamStationStruct := []*Struct_Dam{}
			arrDamStation := strings.Split(_dam.String, "|")
			for _, dam := range arrDamStation {
				arrDam := strings.Split(dam, "##")
				intID, err := strconv.ParseInt(arrDam[0], 10, 64)
				if err != nil {
					return nil, err
				}
				strName := arrDam[1]
				if strName == "" {
					strName = "{}"
				}
				objDamStation := &Struct_Dam{}
				objDamStation.Id = intID
				objDamStation.Dam_name = json.RawMessage(strName)
				arrDamStationStruct = append(arrDamStationStruct, objDamStation)
			}
			objDam.Dam = arrDamStationStruct
		}

		data = append(data, objDam)
	}

	return data, nil
}
