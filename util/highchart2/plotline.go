package highchart

import ()

type Plotline struct {
	Line
	Value float64           `json:"value"`
	Width int               `json:"width"`
	Y     float64           `json:"y"`
	Label map[string]string `json:"label"`
}

func (p *Plotline) _constructor(name string) {
	p.Label = map[string]string{
		"text": name,
	}
	p.Width = 1
	p.Color = "yellow"
	p.SetDashStlye_Dash()
}
