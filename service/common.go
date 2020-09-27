package service

import (
	"context"
	"fmt"
	"time"
	
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

// 根据手机号和密码登录
func (m *Common) MobileLoginByPassword() (*MemberLoginRes, error) {
	mobile := m.PostForm("mobile")
	password := m.PostForm("password")
	
	req := memberpb.MobilePasswdReq{
		Mobile:   mobile,
		Password: password,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err := gclient.MemberClient.LoginByMobile(ctx, &req)
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("登录失败， err：%v", err)
	}
	
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

// 根据手机号和验证码登录
func (m *Common) MobileLoginByCode() (*MemberLoginRes, error) {
	return nil, nil
}

func (m *Common) MobileRegisterByPassword() (*MemberLoginRes, error) {
	mobile := m.PostForm("mobile")
	code := m.PostForm("code")
	password := m.PostForm("password")
	
	if code != "0000" {
		return nil, fmt.Errorf("验证码错误!")
	}
	
	req := memberpb.MobilePasswdReq{
		Mobile:   mobile,
		Password: password,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err := gclient.MemberClient.RegisterByMobile(ctx, &req)
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("注册失败， err：%v", err)
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
