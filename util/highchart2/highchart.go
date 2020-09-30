package highchart

import ()

var LeapYear = 2016

type HighChart struct {
	Series    []*Series   `json:"series"`
	PlotLines []*Plotline `json:"plotLines"`
}

func NewChart() *HighChart {
	m := new(HighChart)
	return m
}

func (hc *HighChart) NewSerie(name string) *Series {
	return hc.newSeries("", name)
}
func (hc *HighChart) NewLineSerie(name string) *Series {
	return hc.newSeries("line", name)
}
func (hc *HighChart) NewColumnSerie(name string) *Series {
	return hc.newSeries("column", name)
}

func (hc *HighChart) newSeries(stype, name string) *Series {
	s := new(Series)

	s._constructor(stype, name)

	hc.Series = append(hc.Series, s)
	return s
}

func (hc *HighChart) NewPlotLines(name string) (*Plotline, *Series) {
	p := &Plotline{}

	p._constructor(name)
	hc.PlotLines = append(hc.PlotLines, p)

	s := hc.NewLineSerie(name)
	s._constructor("", name)
	s.PlotLine()

	return p, s
}
