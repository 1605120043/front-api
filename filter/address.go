package filter

import (
	"regexp"
	"strconv"
	
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
	pageSize := m.DefaultQuery("page_size", "10")
	
	valid := validation.Validation{}
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("拉取格式不正确")
	if valid.HasError() {
		return nil, valid.GetError()
	}
	
	return service.NewAddress(m.Context).Index()
}

func (m *Address) Detail() (*address.AddressDetail, error) {
	addressId := m.Query("address_id")
	m.validation.Numeric(addressId).Message("参数错误！")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	return service.NewAddress(m.Context).Detail()
}

// 添加收货地址
func (m *Address) Add() error {
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	codeProv := m.PostForm("code_prov")
	codeCity := m.PostForm("code_city")
	codeCoun := m.PostForm("code_coun")
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	//longitude := m.PostForm("longitude")
	//latitude := m.PostForm("latitude")
	
	m.validation.Required(name).Message("收货人姓名不能为空！")
	m.validation.Mobile(mobile).Message("手机号格式不正确！")
	m.validation.Numeric(codeProv).Message("请选择省市区")
	m.validation.Numeric(codeCity).Message("请选择省市区")
	m.validation.Numeric(codeCoun).Message("请选择省市区")
	m.validation.Required(addressDetail).Message("收货地址不能为空！")
	m.validation.Required(roomNumber).Message("门牌号不能为空！")
	m.validation.Bool(isDefault).Message("是否默认！")
	//m.validation.Required(longitude).Message("经度不能为空！")
	//m.validation.Required(latitude).Message("纬度不能为空！")
	
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
	addressId := m.PostForm("address_id")
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	codeProv := m.PostForm("code_prov")
	codeCity := m.PostForm("code_city")
	codeCoun := m.PostForm("code_coun")
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	//longitude := m.PostForm("longitude")
	//latitude := m.PostForm("latitude")
	
	m.validation.Numeric(addressId).Message("参数错误！")
	m.validation.Required(name).Message("收货人姓名不能为空！")
	m.validation.Mobile(mobile).Message("手机号格式不正确！")
	m.validation.Numeric(codeProv).Message("请选择省市区")
	m.validation.Numeric(codeCity).Message("请选择省市区")
	m.validation.Numeric(codeCoun).Message("请选择省市区")
	m.validation.Required(addressDetail).Message("收货地址不能为空！")
	m.validation.Required(roomNumber).Message("门牌号不能为空！")
	m.validation.Bool(isDefault).Message("是否默认！")
	//m.validation.Required(longitude).Message("经度不能为空！")
	//m.validation.Required(latitude).Message("纬度不能为空！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	if err := service.NewAddress(m.Context).Edit(); err != nil {
		return err
	}
	
	return nil
}

// 删除收货地址
func (m *Address) Delete() error {
	addressId := m.PostForm("address_id")
	
	m.validation.Numeric(addressId).Message("参数错误！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	addressIdNumber, _ := strconv.ParseUint(addressId, 10, 64)
	if err := service.NewAddress(m.Context).Delete(addressIdNumber); err != nil {
		return err
	}
	
	return nil
}
