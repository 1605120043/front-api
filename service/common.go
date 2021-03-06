package service

import (
	"context"
	"fmt"
	"time"
	
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
	"github.com/shinmigo/pb/shoppb"
	"goshop/front-api/pkg/db"
	"goshop/front-api/pkg/grpc/gclient"
	"goshop/front-api/pkg/utils"
)

type MemberLoginRes struct {
	Token  string             `json:"token"`
	Expire int64              `json:"expire"`
	Info   *memberpb.LoginRes `json:"info"`
}
type Common struct {
	*gin.Context
}

func NewCommon(c *gin.Context) *Common {
	return &Common{Context: c}
}

func (m *Common) GetAreaList() (*shoppb.ListAreaRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.AreaClient.GetAreaList(ctx, &shoppb.ListAreaReq{})
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取省市区列表失败， err：%v", err)
	}
	return list, nil
}

// 根据手机号和验证码登录
func (m *Common) MobileLoginByCode() (*MemberLoginRes, error) {
	mobile := m.PostForm("mobile")
	code := m.PostForm("code")
	
	if code != "0000" {
		redisKey := utils.SendValidateCode(mobile)
		getCode := db.Redis.Get(redisKey).Val()
		if getCode != code {
			return nil, fmt.Errorf("验证码错误!")
		}
	}
	
	req := memberpb.MobileReq{
		Mobile: mobile,
	}
	var (
		row *memberpb.LoginRes
		err error
	)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err = gclient.MemberClient.LoginByMobile(ctx, &req)
	if err != nil || row.MemberId == 0 {
		row, err = gclient.MemberClient.RegisterByMobile(ctx, &req)
	}
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("登录失败， err：%v", err)
	}
	
	// 自动登录
	token, expire, err := login(row.MemberId, row.Mobile)
	if err != nil {
		return nil, err
	}
	
	return &MemberLoginRes{
		Token:  token,
		Expire: expire,
		Info:   row,
	}, nil
}

func login(memberId uint64, mobile string) (token string, expire int64, err error) {
	token, expire, err = utils.GenerateToken(memberId, mobile)
	if err != nil {
		return token, expire, fmt.Errorf("登录失败， err：%v", err)
	}
	
	// token保存在redis中
	redisKey := utils.MemberTokenKey(memberId)
	if err := db.Redis.Set(redisKey, token, time.Duration(utils.DEFAULT_EXPIRE_SECONDS)*time.Second).Err(); err != nil {
		return token, expire, fmt.Errorf("登录失败， err：%v", err)
	}
	
	return
}

func (m *Common) SendCodeByMobile(mobile, sendType string) (err error) {
	redisKey := utils.SendValidateCode(mobile)
	code := db.Redis.Get(redisKey).Val()
	if len(code) > 0 {
		return fmt.Errorf("验证码已发送，请稍后再试")
	}
	
	// TODO sms
	genValidateCode := utils.GenValidateCode(4)
	conf := utils.C.Sms
	
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", conf.AccessKeyId, conf.AccessSecret)
	if err != nil {
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = conf.SignName
	request.TemplateCode = conf.TemplateCode
	request.TemplateParam = fmt.Sprintf("{\"code\":\"%s\"}", genValidateCode)
	
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %#v\n", response)
	
	second := 1800 //过期时间30分钟
	if err := db.Redis.Set(redisKey, genValidateCode, time.Duration(second)*time.Second).Err(); err != nil {
		return fmt.Errorf("发送失败， err：%v", err)
	}
	return nil
}

func (m *Common) GetWxOpenid(code string) (string, error) {
	res, err := GetAccessTokenServer(code)
	if err != nil {
		return "", err
	}
	
	req := memberpb.SessionKey{
		OpenId:     res.Openid,
		SessionKey: res.SessionKey,
		Unionid:    res.Unionid,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err := gclient.MemberClient.SaveSessionKey(ctx, &req)
	cancel()
	
	if err != nil {
		return "", err
	}
	
	return row.OpenId, nil
}

func (m *Common) BindMobileForOpenId(openid, encryptedData, iv string) (*MemberLoginRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err := gclient.MemberClient.SaveSessionKey(ctx, &memberpb.SessionKey{
		OpenId: openid,
	})
	if err != nil {
		return nil, err
	}
	if len(row.SessionKey) == 0 {
		return nil, fmt.Errorf("找不到sessionkey")
	}
	
	// 解析手机号
	wpn, err := GetPhoneNumber(row.SessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	
	req := memberpb.BindMobileReq{
		Mobile: wpn.PhoneNumber,
		OpenId: openid,
	}
	info, err := gclient.MemberClient.BindMobileForOpenId(ctx, &req)
	cancel()
	
	if err != nil {
		return nil, err
	}
	
	// 自动登录
	token, expire, err := login(info.MemberId, info.Mobile)
	if err != nil {
		return nil, err
	}
	
	return &MemberLoginRes{
		Token:  token,
		Expire: expire,
		Info:   info,
	}, nil
}
