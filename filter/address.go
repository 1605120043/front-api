package filter

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/model/address"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Address struct {
	validation validation.Validation
	*gin.Context
}

func NewAddress(c *gin.Context) *Address {
	return &Address{Context: c, validation: validation.Validation{}}
}

func (m *Address) Index() ([]*address.AddressList, error) {
	return service.NewAddress(m.Context).Index()
}

func (m *Address) Detail() (*address.AddressDetail, error) {
	return service.NewAddress(m.Context).Detail()
}

// 添加收货地址
func (m *Address) Add() error {
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	longitude := m.PostForm("longitude")
	latitude := m.PostForm("latitude")
	
	m.validation.Required(name).Message("收货人姓名不能为空！")
	m.validation.Mobile(mobile).Message("手机号格式不正确！")
	m.validation.Required(addressDetail).Message("收货地址不能为空！")
	m.validation.Required(roomNumber).Message("门牌号不能为空！")
	m.validation.Bool(isDefault).Message("是否默认！")
	m.validation.Required(longitude).Message("经度不能为空！")
	m.validation.Required(latitude).Message("纬度不能为空！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	if err := service.NewAddress(m.Context).Add(); err != nil {
		return err
	}
	
	return nil
}

// 编辑收货地址
func (m *Address) Edit() error {
	addressId := m.PostForm("id")
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	longitude := m.PostForm("longitude")
	latitude := m.PostForm("latitude")
	
	m.validation.Numeric(addressId).Message("参数错误！")
	m.validation.Required(name).Message("收货人姓名不能为空！")
	m.validation.Mobile(mobile).Message("手机号格式不正确！")
	m.validation.Required(addressDetail).Message("收货地址不能为空！")
	m.validation.Required(roomNumber).Message("门牌号不能为空！")
	m.validation.Bool(isDefault).Message("是否默认！")
	m.validation.Required(longitude).Message("经度不能为空！")
	m.validation.Required(latitude).Message("纬度不能为空！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	if err := service.NewAddress(m.Context).Edit(); err != nil {
		return err
	}
	
	return nil
}
