package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
	"goshop/front-api/pkg/grpc/gclient"
)

type Member struct {
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c}
}

func (m *Member) Info() (*memberpb.LoginRes, error) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err := gclient.MemberClient.GetMemberForLogin(ctx, &memberpb.MemberIdReq{MemberId: memberId})
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取失败， err：%v", err)
	}
	return row, nil
}
