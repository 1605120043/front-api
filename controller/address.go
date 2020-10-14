package controller

import "goshop/front-api/filter"

var addressFilter *filter.Address

type Address struct {
	Base
}

func (m *Address) Initialise() {
	addressFilter = filter.NewAddress(m.Context)
}

func (m *Address) Index() {
	str, err := addressFilter.Index()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Address) Detail() {
	str, err := addressFilter.Detail()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Address) Add() {
	err := addressFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}

func (m *Address) Edit() {
	err := addressFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}

func (m *Address) Delete() {
	err := addressFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}
