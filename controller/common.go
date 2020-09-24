package controller

import "goshop/front-api/filter"

var commonFilter *filter.Common

type Common struct {
	Base
}

func (m *Common) Initialise() {
	commonFilter = filter.NewCommon(m.Context)
}

func (m *Common) GetAreaList() {
	str, err := commonFilter.GetAreaList()
	if err != nil {
		m.SetResponse(str, err)
		return
	}
	
	m.SetResponse(str)
}
