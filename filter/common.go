package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Common struct {
	validation validation.Validation
	*gin.Context
}

func NewCommon(c *gin.Context) *Common {
	return &Common{Context: c, validation: validation.Validation{}}
}

func (m *Common) GetAreaList() (*shoppb.ListAreaRes, error) {
	return service.NewCommon(m.Context).GetAreaList()
}

// 根据手机号和密码登录
func (m *Common) MobileLoginByPassword() (*service.MemberLoginRes, error) {
	mobile := m.PostForm("mobile")
	password := m.PostForm("password")
	
	m.validation.Required(mobile).Message("请输入手机号码！")
	m.validation.Required(password).Message("请输入登录密码！")
	m.validation.Mobile(mobile).Message("请输入正确的手机号码！")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	return service.NewCommon(m.Context).MobileLoginByPassword()
}

// 根据手机号和验证码登录
func (m *Common) MobileLoginByCode() (*service.MemberLoginRes, error) {
	mobile := m.PostForm("mobile")
	code := m.PostForm("code")
	
	m.validation.Required(mobile).Message("请输入手机号码！")
	m.validation.Required(code).Message("请输入验证码！")
	m.validation.Mobile(mobile).Message("请输入正确的手机号码！")
	m.validation.Numeric(code).Message("请输入正确的验证码！")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	return service.NewCommon(m.Context).MobileLoginByCode()
}

// 根据手机号和密码注册
func (m *Common) MobileRegisterByPassword() (*service.MemberLoginRes, error) {
	mobile := m.PostForm("mobile")
	code := m.PostForm("code")
	password := m.PostForm("password")
	
	m.validation.Required(mobile).Message("请输入手机号码！")
	m.validation.Required(code).Message("请输入验证码！")
	m.validation.Required(password).Message("请输入登录密码！")
	m.validation.Numeric(code).Message("请输入正确的验证码！")
	m.validation.Mobile(mobile).Message("请输入正确的手机号码！")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	return service.NewCommon(m.Context).MobileRegisterByPassword()
}