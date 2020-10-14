package filter

import (
	"encoding/json"
	"fmt"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"goshop/front-api/model/cart"
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

func (m *Cart) Add() (carts *cart.Carts, err error) {
	productId := m.PostForm("product_id")
	productSpecId := m.PostForm("product_spec_id")
	nums := m.PostForm("nums")
	isSelect := m.PostForm("is_select")
	
	m.validation.Numeric(productId).Message("请选择商品！")
	m.validation.Numeric(productSpecId).Message("请选择商品！")
	m.validation.Integer(nums).Message("数量是整形！")
	m.validation.Bool(isSelect).Message("是否选中？")
	
	if m.validation.HasError() {
		err = m.validation.GetError()
		return
	}
	
	err = service.NewCart(m.Context).Add()
	if err != nil {
		return
	}
	
	return service.NewCart(m.Context).Index()
}

func (m *Cart) Delete() (carts *cart.Carts, err error) {
	isAll := m.DefaultPostForm("is_all", "0")
	cartIds := m.PostForm("cart_ids")
	
	m.validation.Switch(isAll).Message("是否删除全部")
	m.validation.Required(cartIds).Message("请选择购物车商品！")
	
	if m.validation.HasError() {
		err = m.validation.GetError()
		return
	}
	
	isAllNum, _ := strconv.ParseInt(m.DefaultPostForm("is_all", "0"), 10, 8)
	
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	cartIdList := make([]uint64, 0, 32)
	
	if isAllNum == 0 { //根据id删除
		cartIds := m.PostForm("cart_ids")
		
		if err := json.Unmarshal([]byte(cartIds), &cartIdList); err != nil {
			return nil, fmt.Errorf("请选择购物车商品, err: %v", err)
		}
	}
	
	err = service.NewCart(m.Context).Delete(int8(isAllNum), memberId, cartIdList)
	
	return service.NewCart(m.Context).Index()
}

func (m *Cart) Index() (*cart.Carts, error) {
	return service.NewCart(m.Context).Index()
}

func (m *Cart) Count() uint64 {
	return service.NewCart(m.Context).Count()
}

func (m *Cart) Checked() (carts *cart.Carts, err error) {
	cartChecked := m.PostForm("cart_checked")
	
	m.validation.Required(cartChecked).Message("请选择购物车商品！")
	
	if m.validation.HasError() {
		err = m.validation.GetError()
		return
	}
	
	err = service.NewCart(m.Context).Checked()
	
	return service.NewCart(m.Context).Index()
}

func (m *Cart) Buy() (*cart.BuyRes, error) {
	products := m.PostForm("products")
	
	m.validation.Required(products).Message("请选择商品！")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	
	req := make([]*cart.BuyReq, 0, 32)
	if err := json.Unmarshal([]byte(products), &req); err != nil {
		return nil, fmt.Errorf("参数错误, err: %v", err)
	}
	
	return service.NewCart(m.Context).Buy(req)
}
