package result

import (
	"encoding/json"
)

type Result struct {
	Result string      `json:"result"` // example:`OK` "OK", "NO" 
	Data   interface{} `json:"data,omitempty"`
}

type ResultJson struct {
	Result string          `json:"result"` // example:`OK` "OK", "NO"
	Data   json.RawMessage `json:"data,omitempty"`
}

func rs(s string, data interface{}) *Result {
	return &Result{Result: s, Data: data}
}

func Result0(data interface{}) *Result {
	return rs("NO", data)
}
func Result1(data interface{}) *Result {
	return rs("OK", data)
}

func ResultJson1(data json.RawMessage) *ResultJson {
	return &ResultJson{Result: "OK", Data: data}
}
