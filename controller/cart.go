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
	str, err := cartFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str, "加入成功", 1)
}

// 批量移除购物车商品
func (m *Cart) Delete() {
	str, err := cartFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str, "删除成功", 1)
}

// 获取购物车列表
func (m *Cart) Index() {
	str, err := cartFilter.Index()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}

// 获取购物车数量
func (m *Cart) Count() {
	total := cartFilter.Count()
	m.SetResponse(total)
}

// 批量选择或者取消购物车商品
func (m *Cart) Checked() {
	str, err := cartFilter.Checked()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str, "成功", 1)
}

func (m *Cart) Buy() {
	str, err := cartFilter.Buy()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}
