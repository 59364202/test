package dataunit

import (
	"encoding/json"
)

type Dataunit_struct struct {
	Id           int64           `json:"id"`
	DataunitName json.RawMessage `json:"dataunit_name,omitempty"`
}
