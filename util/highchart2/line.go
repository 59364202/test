package highchart

import ()

type Line struct {
	Color     string `json:"color"`
	DashStyle string `json:"dashStyle"`
}

func (s *Line) SetDashStlye_Solid() {
	s.DashStyle = "Solid"
}
func (s *Line) SetDashStlye_ShortDash() {
	s.DashStyle = "ShortDash"
}
func (s *Line) SetDashStlye_ShortDot() {
	s.DashStyle = "ShortDot"
}
func (s *Line) SetDashStlye_ShortDashDot() {
	s.DashStyle = "Solid"
}
func (s *Line) SetDashStlye_Dot() {
	s.DashStyle = "Dot"
}
func (s *Line) SetDashStlye_Dash() {
	s.DashStyle = "Dash"
}
func (s *Line) SetDashStlye_LongDash() {
	s.DashStyle = "LongDash"
}
func (s *Line) SetDashStlye_DashDot() {
	s.DashStyle = "DashDot"
}
func (s *Line) SetDashStlye_LongDashDot() {
	s.DashStyle = "LongDashDot"
}
func (s *Line) SetDashStlye_LongDashDotDot() {
	s.DashStyle = "LongDash"
}
