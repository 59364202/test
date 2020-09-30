package selectoption

import ()

type SelectOption struct {
	Option []*Option `json:"option"`
}

type Option struct {
	Value interface{} `json:"value"`
	Text  interface{} `json:"text"`
}

func NewSelect() *SelectOption {

	s := &SelectOption{}

	o := make([]*Option, 0)
	s.Option = o

	return s
}

func (selectOption *SelectOption) Add(v, t interface{}) {
	selectOption.Option = append(selectOption.Option, &Option{Text: t, Value: v})
}
