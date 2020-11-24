package filter

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"unicode/utf8"
	
	"goshop/front-api/model/wxapp"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
)

type Member struct {
	validation validation.Validation
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c, validation: validation.Validation{}}
}

func (m *Member) Info() (*memberpb.LoginRes, error) {
	return service.NewMember(m.Context).Info()
}

func (m *Member) Update() (bool, error) {
	avatar := m.PostForm("avatar")
	nickname := m.PostForm("nickname")
	gender := m.PostForm("gender")
	birthday := m.PostForm("birthday")
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	valid := validation.Validation{}
	//valid.Required(avatar).Message("请上传头像信息！")
	//valid.Match(avatar, regexp.MustCompile(`^[a-zA-z0-9,\-\.]+$`)).Message("头像格式错误")
	//valid.Required(nickname).Message("请填写昵称信息！")
	if len(nickname) > 0 {
		valid.Match(nickname, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9%|！]+$`)).Message("昵称格式错误")
	}
	valid.Required(gender).Message("请选择性别！")
	valid.Match(gender, regexp.MustCompile(`^0|1|2$`)).Message("性别格式错误")
	if len(birthday) > 0 {
		valid.Match(birthday, regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)).Message("生日信息格式错误")
	}
	if valid.HasError() {
		return false, valid.GetError()
	}
	if utf8.RuneCountInString(nickname) > 30 {
		return false, errors.New("昵称长度超过30个字符！")
	}
	genderNum, _ := strconv.ParseUint(gender, 10, 64)
	genderType := memberpb.MemberGender(genderNum)
	
	req := &memberpb.Member{
		MemberId: memberId,
		Nickname: nickname,
		Gender:   genderType,
		Birthday: birthday,
		Avatar:   avatar,
	}
	return service.NewMember(m.Context).Update(req)
}

func (m *Member) Pay() (map[string]string, error) {
	orderId := m.PostForm("order_id")
	paymentCode := m.PostForm("payment_code")
	tradeType := m.PostForm("trade_type")
	
	m.validation.Required(orderId).Message("订单编号错误！")
	m.validation.Required(paymentCode).Message("请选择支付方式！")
	m.validation.Required(tradeType).Message("参数出错！")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	orderIdNum, _ := strconv.ParseUint(orderId, 10, 64)
	
	return service.NewMember(m.Context).Pay(memberId, orderIdNum, paymentCode, tradeType)
}

func (m *Member) WxNotify() (wxn wxapp.WXPayNotify, err error) {
	var body []byte
	body, err = ioutil.ReadAll(m.Request.Body)
	if err != nil {
		return
	}
	
	log.Println(string(body))
	err = xml.Unmarshal(body, &wxn)
	if err != nil {
		return
	}
	return
}
