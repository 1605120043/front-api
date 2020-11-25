package controller

import (
	"goshop/front-api/filter"
	"goshop/front-api/pkg/utils"
)

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
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Member) Update() {
	res, err := memberFilter.Update()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(res)
}

func (m *Member) Payment() {
	m.SetResponse(utils.Payments.Payment)
}

func (m *Member) Pay() {
	str, err := memberFilter.Pay()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	m.SetResponse(str)
}

func (m *Member) Notify() {
	if err := memberFilter.WxNotify(); err != nil {
		m.SetResponse(nil, err)
		return
	}
	m.SetResponse()
}
