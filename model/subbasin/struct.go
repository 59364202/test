package subbasin

import (
	//model_basin "haii.or.th/api/thaiwater30/model/basin"
	"encoding/json"
)

type Struct_Subbasin struct {
	//Basin		*model_basin.Struct_Basin `json:"basin"`
	Id         int64           `json:"id"`                   // example:`1` รหัสลุ่มน้ำ
	Basin_name json.RawMessage `json:"basin_name,omitempty"` // example:`{"th":"ลุ่มน้ำสาละวิน","en":"MAE NAM SALAWIN"}` ชื่อลุ่มน้ำ
}

type Basin struct {
	BasinCode string          `json:"basin_code"` // example:`1` รหัสลุ่มน้ำ
	BasinName json.RawMessage `json:"basin_name"` // example:`{"th":"ลุ่มน้ำสาละวิน","en":"MAE NAM SALAWIN"}` ชื่อลุ่มน้ำ
	Subbasin  []*Subbasin     `json:"subbasin"`   // ลุ่มน้ำสาขา
}

type Subbasin struct {
	ID           int64           `json:"subbasin_id"`    // example:`1` รหัสลุ่มน้ำสาขา
	SubbasinCode string          `json:"subbasin_code"`  // example:`102` รหัสลุ่มน้ำสาขา
	SubbasinName json.RawMessage `json:"sbubbasin_name"` // example:`{"th":"น้ำแม่ปายตอนบน","en":"UPPER PART OF NAM MAE PAI"}` ชื่อรหัสลุ่มน้ำสาขา
}
