package filter

import (
	"fmt"
	
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

func (m *Common) SendCodeByMobile() error {
	mobile := m.PostForm("mobile")
	sendType := m.DefaultPostForm("send_type", "login")
	
	m.validation.Required(mobile).Message("请输入手机号码！")
	
	if m.validation.HasError() {
		return m.validation.GetError()
	}
	
	return service.NewCommon(m.Context).SendCodeByMobile(mobile, sendType)
}

func (m *Common) MemberLoginByWXApp() (*service.MemberLoginRes, error) {
	code := m.GetString("code")
	encryptedData := m.GetString("encryptedData")
	iv := m.GetString("iv")
	
	m.validation.Required(code).Message("检查参数code")
	m.validation.Required(encryptedData).Message("检查参数encryptedData")
	m.validation.Required(iv).Message("检查参数iv")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	row, err := service.NewCommon(m.Context).MemberLoginByWXApp(code, encryptedData, iv)
	if len(err) > 0 {
		return row, fmt.Errorf(err)
	}
	
	return row, nil
}
