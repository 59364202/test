package dam_yearly

import ()

type GraphDamYearlyInput struct {
	DataType string  `json:"data_type"` // field ที่ต้องการ เช่น dam_inflow
	DamID    []int64 `json:"dam_id"`    // รหัสเขื่อน เช่น [1,2,3,4]
	Year     []int64 `json:"year"`	// ปี  เช่น [2015,2016]
	Month    string  `json:"month"` // เดือน เช่น 10
	Day      string  `json:"day"` // วัน เช่น 23
}

type GraphDamMediumInput struct {
	DataType string  `json:"data_type"` // field ที่ต้องการ เช่น  mediumdam_inflow
	DamID    []int64 `json:"medium_id"` // รหัสเขื่อน เช่น [1,2,3,4]
	Year     []int64 `json:"year"` // ปี  เช่น [2015,2016]
}

type GraphDamOutput struct {
	GraphData      []*DataOutput `json:"graph_data"`       // ข้อมูลกราฟ
	UpperRuleCurve []*RuleCurve  `json:"upper_rule_curve"` // ข้อมูล upper rule curve
	LowerRuleCurve []*RuleCurve  `json:"lower_rule_curve"` // ข้อมูล lower rule curve
	LowerBound     interface{}   `json:"lower_bound"`      // example:`213` กักเก็บต่ำสุด
	UpperBound     interface{}   `json:"upper_bound"`      // example:`260` กักเก็บสูงสุด
	NormalBound    interface{}   `json:"normal_bound"`     // example:`260` กักเก็บปกติ
}

type DataOutput struct {
	Year int64       `json:"year"` // example:`2006` ปี
	Data []*DataYear `json:"data"` // ข้อมูล
}

type DataYear struct {
	DamDate string      `json:"date"`  // example:`2006-01-02` วันที่เก็บข้อมูล
	Value   interface{} `json:"value"` // example:`12` ค่า
}

type RuleCurve struct {
	Date string    `json:"date"`  // example:`2006-01-02` วันที่
	Value interface{} `json:"value"` // example:`32` ค่า
}
