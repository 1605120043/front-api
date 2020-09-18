package service

import (
	"context"
	"fmt"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/pkg/grpc/gclient"
)

type Category struct {
	*gin.Context
}

func NewCategory(c *gin.Context) *Category {
	return &Category{Context: c}
}

func (m *Category) Index(param *productpb.ListCategoryReq) (*productpb.ListCategoryRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.ProductCategoryClient.GetCategoryList(ctx, param)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取分类列表失败， err：%v", err)
	}
	return list, nil
}
