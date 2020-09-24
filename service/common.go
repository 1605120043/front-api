package service

import (
	"context"
	"fmt"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
	"goshop/front-api/pkg/grpc/gclient"
)

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
