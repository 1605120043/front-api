package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"goshop/front-api/pkg/grpc/gclient"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
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

func (m *Member) Update(req *memberpb.Member) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	res, err := gclient.MemberClient.EditMember(ctx, req)
	cancel()
	if err != nil {
		return false, fmt.Errorf("修改信息失败， err：%v", err)
	}
	if res.State == 0 {
		return false, fmt.Errorf("修改信息失败， err：%v", err)
	}

	return true, nil
}
