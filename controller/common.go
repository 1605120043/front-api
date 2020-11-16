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
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Common) MobileLogin() {
	str, err := commonFilter.MobileLoginByCode()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Common) SendCode() {
	err := commonFilter.SendCodeByMobile()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(nil)
}

func (m *Common) WxLogin() {
	str, err := commonFilter.MemberLoginByWXApp()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}
