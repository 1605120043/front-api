package filter

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Cart struct {
	validation validation.Validation
	*gin.Context
}

func NewCart(c *gin.Context) *Cart {
	return &Cart{Context: c, validation: validation.Validation{}}
}

func (m *Cart) Add() error {
	productId := m.PostForm("product_id")
	productSpecId := m.PostForm("product_spec_id")
	nums := m.PostForm("nums")
	isSelect := m.PostForm("is_select")
	
	m.validation.Numeric(productId).Message("请选择商品！")
	m.validation.Numeric(productSpecId).Message("请选择商品！")
	m.validation.Integer(nums).Message("数量是整形！")
	m.validation.Bool(isSelect).Message("是否选中？")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	return service.NewCart(m.Context).Add()
}

func (m *Cart) Delete() error {
	isAll := m.DefaultPostForm("is_all", "0")
	cartIds := m.PostForm("cart_ids")
	
	m.validation.Switch(isAll).Message("是否删除全部")
	m.validation.Required(cartIds).Message("请选择购物车商品！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	return service.NewCart(m.Context).Delete()
}

func (m *Cart) Index() (map[string]interface{}, error) {
	return service.NewCart(m.Context).Index()
}

func (m *Cart) Checked() error {
	cartChecked := m.PostForm("cart_checked")
	
	m.validation.Required(cartChecked).Message("请选择购物车商品！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	return service.NewCart(m.Context).Checked()
}