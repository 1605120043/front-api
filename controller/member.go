package controller

import "goshop/front-api/filter"

var memberFilter *filter.Member

type Member struct {
	Base
}

func (m *Member) Initialise() {
	memberFilter = filter.NewMember(m.Context)
}

func (m *Member) Info() {
	str, err := memberFilter.Info()
	if err != nil {
		m.SetResponse(str, err)
		return
	}

	m.SetResponse(str)
}

func (m *Member) Update() {
	res, err := memberFilter.Update()
	if err != nil {
		m.SetResponse(res, err)
		return
	}

	m.SetResponse(res)
}
