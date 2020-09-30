package hydroinfo

import (
	"encoding/json"
)

type Hydroinfo_struct struct {
	Id              int64             `json:"id"`
	HydroinfoName   json.RawMessage   `json:"hydroinfo_name,omitempty"`
	AgencyId        []int             `json:"agency_id"`
	AgencyName      []json.RawMessage `json:"agency_name,omitempty"`
	AgencyShortName []json.RawMessage `json:"agency_shortname,omitempty"`
}

type Hydroinfo_struct1 struct {
	Id              int64   `json:"id"`
	HydroinfoName   *Lang   `json:"hydroinfo_name,omitempty"`
	AgencyId        []int   `json:"agency_id"`
	AgencyName      []*Lang `json:"agency_name,omitempty"`
	AgencyShortName []*Lang `json:"agency_shortname,omitempty"`
}

type Lang struct {
	TH string `json:"th"`
	EN string `json:"en"`
	JP string `json:"jp"`
}
