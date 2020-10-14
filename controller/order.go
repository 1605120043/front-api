package controller

import "goshop/front-api/filter"

var orderFilter *filter.Order

type Order struct {
	Base
}

func (m *Order) Initialise() {
	orderFilter = filter.NewOrder(m.Context)
}

func (m *Order) Index() {
	str, err := orderFilter.Index()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Order) Info() {
	info, err := orderFilter.Info()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(info)
}

func (m *Order) GetUserOrderStatusCount() {
	list, err := orderFilter.GetUserOrderStatusCount()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(list)
}

func (m *Order) CreateOrder() {
	str, err := orderFilter.CreateOrder()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

func (m *Order) CancelOrder() {
	err := orderFilter.CancelOrder()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}

func (m *Order) DeleteOrder() {
	err := orderFilter.DeleteOrder()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}
