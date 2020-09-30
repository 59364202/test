package agency

import (
	"database/sql"
	"encoding/json"
	"strings"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

var sqlSelectAgency = ` SELECT ag.id, ag.agency_name, ag.agency_shortname, d.id, d.department_name, m.id, m.ministry_name 
	FROM agency ag 
	INNER JOIN lt_department d ON ag.department_id = d.id AND d.deleted_at = '1970-01-01 07:00:00+07'
	INNER JOIN lt_ministry m ON d.ministry_id = m.id AND m.deleted_at = '1970-01-01 07:00:00+07'
`
var sqlSelectAgency_Where = ` WHERE ag.deleted_by IS NULL AND ag.deleted_at = '1970-01-01 07:00:00+07' `
var sqlSelectAgency_OrderBy = ` ORDER BY ag.agency_name->>'th' `

func GetAgency(agencyId string, departmentId string) ([]*Agency_struct, error) {
	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Varaibles
	var (
		data   []*Agency_struct
		agency *Agency_struct

		_id               sql.NullInt64
		_agency_name      sql.NullString
		_agency_shortname sql.NullString
		_department_id    sql.NullInt64
		_deparmtnet_name  sql.NullString
		_ministry_id      sql.NullInt64
		_ministry_name    sql.NullString

		_result *sql.Rows
	)

	//Set 'sqlCmdWhere' variables
	sqlCmdWhere := sqlSelectAgency_Where

	//Check agency_id
	if agencyId != "" {
		sqlCmdWhere += " AND ag.id = " + agencyId
	}

	//Check department_id
	if departmentId != "" {
		arrDepartmentId := strings.Split(departmentId, ",")
		if len(arrDepartmentId) == 1 {
			sqlCmdWhere += " AND ag.department_id = " + departmentId
		} else {
			sqlCmdWhere += " AND ag.department_id IN (" + departmentId + ") "
		}
	}

	_result, err = db.Query(sqlSelectAgency + sqlCmdWhere + sqlSelectAgency_OrderBy)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	// Loop data result
	data = make([]*Agency_struct, 0)

	for _result.Next() {
		//Scan to execute query with variable
		err := _result.Scan(&_id, &_agency_name, &_agency_shortname, &_department_id, &_deparmtnet_name, &_ministry_id, &_ministry_name)
		if err != nil {
			return nil, err
		}

		//Generate Agency object
		agency = &Agency_struct{}
		agency.Id = _id.Int64
		agency.DepartmentId = _department_id.Int64
		agency.MinistryId = _ministry_id.Int64

		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}
		agency.AgencyName = json.RawMessage(_agency_name.String)

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}
		agency.AgencyShortName = json.RawMessage(_agency_shortname.String)

		if _deparmtnet_name.String == "" {
			_deparmtnet_name.String = "{}"
		}
		agency.DepartmentName = json.RawMessage(_deparmtnet_name.String)

		if _ministry_name.String == "" {
			_ministry_name.String = "{}"
		}
		agency.MinistryName = json.RawMessage(_ministry_name.String)

		data = append(data, agency)
	}

	return data, nil
}
