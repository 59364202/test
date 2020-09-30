package tmd_report

import ()

type Struct_Temperature struct {
	Temperature  string `json:"temperature"`  // example:`28.2` อุณหภูมิ
	Current_time string `json:"current_time"` // example:`22:00 น.` เวลาของข้อมูล
}

//	inittial default value
func (s *Struct_Temperature) Init() {
	s.Temperature = "-"
	s.Current_time = "-"
}
