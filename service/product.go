package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/pkg/grpc/gclient"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
}

func (m *Product) Index() (*productpb.ListProductRes, error) {
	categoryId, _ := strconv.ParseUint(m.DefaultQuery("category_id", "0"), 10, 64)
	page, _ := strconv.ParseUint(m.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseUint(m.DefaultQuery("page_size", "20"), 10, 64)
	req := &productpb.ListProductReq{
		CategoryId: categoryId,
		Page:       page,
		PageSize:   pageSize,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	productList, err := gclient.ProductClient.GetProductList(ctx, req)
	cancel()
	
	return productList, err
}

func (m *Product) Detail() (*productpb.ProductDetail, error) {
	productId, _ := strconv.ParseUint(m.DefaultQuery("product_id", "0"), 10, 64)
	if productId == 0 {
		return nil, fmt.Errorf("获取商品失败")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	productList, err := gclient.ProductClient.GetProductListByProductSpecIds(ctx, &productpb.ProductSpecIdsReq{
		ProductId: []uint64{productId},
	})
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取商品失败, err:%v", err)
	}
	
	if len(productList.Products) == 0 {
		return nil, fmt.Errorf("获取商品失败, err:%v", err)
	}
	
	return productList.Products[0], nil
}
