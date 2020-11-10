package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/model/product"
	"goshop/front-api/pkg/grpc/gclient"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
}

func (m *Product) Index() ([]*product.ProductList, error) {
	categoryId, _ := strconv.ParseUint(m.DefaultQuery("category_id", "0"), 10, 64)
	tagId, _ := strconv.ParseUint(m.DefaultQuery("tag_id", "0"), 10, 64)
	page, _ := strconv.ParseUint(m.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseUint(m.DefaultQuery("page_size", "20"), 10, 64)
	name := m.DefaultPostForm("name", "")
	req := &productpb.ListProductReq{
		CategoryId: categoryId,
		TagId:      tagId,
		Name:       name,
		Page:       page,
		PageSize:   pageSize,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	listRes, err := gclient.ProductClient.GetProductList(ctx, req)
	cancel()
	
	list := make([]*product.ProductList, 0, pageSize)
	if listRes == nil {
		return list, fmt.Errorf("获取商品失败")
	}
	
	for k := range listRes.Products {
		var image string
		if len(listRes.Products[k].Images) > 0 {
			image = listRes.Products[k].Images[0]
		}
		
		list = append(list, &product.ProductList{
			ProductId:        listRes.Products[k].ProductId,
			Name:             listRes.Products[k].Name,
			Image:            image,
			Price:            listRes.Products[k].Price,
			ShortDescription: listRes.Products[k].ShortDescription,
		})
	}
	return list, err
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

func (m *Product) Tag() ([]*product.TagList, error) {
	req := &productpb.ListTagReq{
		Page:     1,
		PageSize: 100,
		Display:  productpb.TagDisplay_Show,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.TagClient.GetTagList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取标签失败, err:%v", err)
	}
	
	tagList := make([]*product.TagList, 0, res.Total)
	for k := range res.Tags {
		tagList = append(tagList, &product.TagList{
			Id:   res.Tags[k].TagId,
			Name: res.Tags[k].Name,
		})
	}
	
	return tagList, nil
}
