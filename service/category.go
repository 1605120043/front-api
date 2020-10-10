package service

import (
	"context"
	"fmt"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/model/product"
	"goshop/front-api/pkg/grpc/gclient"
)

type Category struct {
	*gin.Context
}

func NewCategory(c *gin.Context) *Category {
	return &Category{Context: c}
}

func (m *Category) Index(param *productpb.ListCategoryReq) ([]*product.CategoryList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.ProductCategoryClient.GetCategoryList(ctx, param)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取分类列表失败， err：%v", err)
	}
	
	list := make([]*product.CategoryList, 0, res.Total)
	if res.Total > 0 {
		for k := range res.Categories {
			list = append(list, &product.CategoryList{
				Id:   res.Categories[k].CategoryId,
				Pid:  res.Categories[k].ParentId,
				Name: res.Categories[k].Name,
				Icon: "https://img01.yimishiji.com/v1/img/" + res.Categories[k].Icon,
				Sort: res.Categories[k].Sort,
			})
		}
	}
	return list, nil
}
