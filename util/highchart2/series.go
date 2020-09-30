package highchart

import (
	"time"
)

type Series struct {
	Line
	Name                string        `json:"name"`
	Data                []interface{} `json:"data"`
	LineWidth           int           `json:"lineWidth"`
	Type                string        `json:"type"`
	YAxis               int           `json:"yAxis"`
	ShowInLegend        bool          `json:"showInLegend"`
	EnableMouseTracking bool          `json:"enableMouseTracking"`
	Marker              *Marker       `json:"marker"`
}
type Marker struct {
	Enabled bool `json:"enabled"`
}

func (s *Series) _constructor(stype, name string) {
	s.YLeft()
	s.LineWidth = 1
	s.Name = name
	s.Type = stype
	s.ShowInLegend = true
	s.EnableMouseTracking = true

	s.Marker = &Marker{Enabled: false}
}

//add data
func (s *Series) AddData(f interface{}) {
	if f != nil {
		s.Data = append(s.Data, f)
	} else {
		s.Data = append(s.Data, nil)
	}
}

//add date and data
func (s *Series) AddDateData(date, f interface{}) {
	itf := make([]interface{}, 0)

	if date != nil {
		d := date.(time.Time)
		dd := time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), time.UTC)
		itf = append(itf, dd.UnixNano()/int64(time.Millisecond))
		//				itf = append(itf, date)
	} else {
		itf = append(itf, nil)
	}

	if f != nil {
		itf = append(itf, f)
	} else {
		itf = append(itf, nil)
	}

	s.Data = append(s.Data, itf)
}

func (s *Series) YLeft() {
	s.YAxis = 0
}
func (s *Series) YRight() {
	s.YAxis = 1
}
func (s *Series) CurrentYear() {
	s.Color = "red"
}
func (s *Series) PlotLine() {
	s.LineWidth = 0
	s.ShowInLegend = false
	s.EnableMouseTracking = false
	s.Marker.Enabled = false
}
