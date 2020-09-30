package agency

import (
	"encoding/json"
)

type Agency_struct struct {
	Id              int64           `json:"id"`
	AgencyName      json.RawMessage `json:"agency_name,omitempty"`
	AgencyShortName json.RawMessage `json:"agency_shortname,omitempty"`
	DepartmentId    int64           `json:"department_id"`
	DepartmentName  json.RawMessage `json:"department_name,omitempty"`
	MinistryId      int64           `json:"ministry_id"`
	MinistryName    json.RawMessage `json:"ministry_name,omitempty"`
}
