package controller

import "goshop/front-api/filter"

var cartFilter *filter.Cart

type Cart struct {
	Base
}

func (m *Cart) Initialise() {
	cartFilter = filter.NewCart(m.Context)
}

// 单个商品加入购物车
func (m *Cart) Add() {
	err := cartFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}

// 移除购物车
func (m *Cart) Delete() {
	err := cartFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse()
}

// 获取购物车列表
func (m *Cart) Index() {
	str, err := cartFilter.Index()
	if err != nil {
		m.SetResponse(str, err)
		return
	}
	
	m.SetResponse(str)
}
