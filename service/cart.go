package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
	"github.com/shinmigo/pb/productpb"
	"github.com/shopspring/decimal"
	"goshop/front-api/pkg/grpc/gclient"
)

type Cart struct {
	*gin.Context
}

func NewCart(c *gin.Context) *Cart {
	return &Cart{Context: c}
}

func (m *Cart) Add() error {
	productId, _ := strconv.ParseUint(m.PostForm("product_id"), 10, 64)
	productSpecId, _ := strconv.ParseUint(m.PostForm("product_spec_id"), 10, 64)
	nums, _ := strconv.ParseInt(m.PostForm("nums"), 10, 32)
	isSelect, _ := strconv.ParseInt(m.PostForm("is_select"), 10, 32)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	isPlus := m.GetBool("is_plus")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductClient.GetProductListByProductSpecIds(ctx, &productpb.ProductSpecIdsReq{ProductSpecId: []uint64{productSpecId}})
	
	if err != nil {
		return fmt.Errorf("添加购物车失败, err:%v", err)
	}
	productSpec := resp.ProductSpecs[0]
	
	if productSpec.ProductId != productId {
		return fmt.Errorf("商品不存在")
	}
	
	if productSpec.Product.Status != productpb.ProductStatus_Disabled {
		return fmt.Errorf("商品:%s, 已下架", productSpec.Product.Name)
	}
	
	if int64(productSpec.Stock) < nums {
		return fmt.Errorf("商品:%s (%s), 库存不足", productSpec.Product.Name, productSpec.Sku)
	}
	
	// 加入购物车
	addCart, err := gclient.CartClient.AddCart(ctx, &memberpb.AddCartReq{
		MemberId:      memberId,
		ProductId:     productId,
		ProductSpecId: productSpecId,
		IsSelect:      memberpb.CartIsSelect(isSelect),
		Nums:          nums,
		IsPlus:        isPlus, //true是叠加
	})
	cancel()
	
	if err != nil {
		return fmt.Errorf("添加购物车失败, err:%v", err)
	}
	
	if addCart.State == 0 {
		return fmt.Errorf("添加失败")
	}
	return nil
}

func (m *Cart) Delete() error {
	cartIds := m.PostForm("cart_ids")
	cartIdList := make([]uint64, 0, 32)
	if err := json.Unmarshal([]byte(cartIds), &cartIdList); err != nil {
		return fmt.Errorf("请选择购物车商品, err: %v", err)
	}
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.CartClient.DelCart(ctx, &memberpb.DelCartReq{
		CartIds:  cartIdList,
		MemberId: memberId,
	})
	cancel()
	
	if err != nil {
		return fmt.Errorf("移除购物车失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("移除失败")
	}
	return nil
}

func (m *Cart) Index() (map[string]interface{}, error) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	cartList, err := gclient.CartClient.GetCartListByMemberId(ctx, &memberpb.ListCartReq{MemberId: memberId})
	
	if err != nil {
		return nil, fmt.Errorf("获取购物车列表失败, err:%v", err)
	}
	
	productSpecIds := make([]uint64, 0, len(cartList.Carts))
	for k := range cartList.Carts {
		productSpecIds = append(productSpecIds, cartList.Carts[k].ProductSpecId)
	}
	productSpecList, err := gclient.ProductClient.GetProductListByProductSpecIds(ctx, &productpb.ProductSpecIdsReq{ProductSpecId: productSpecIds})
	
	if err != nil {
		return nil, fmt.Errorf("获取购物车列表失败, err:%v", err)
	}
	
	ProductSpecKeyByProductSpecId := make(map[uint64]*productpb.ListProductSpecRes_ProductSpec, len(productSpecList.ProductSpecs))
	for _, spec := range productSpecList.ProductSpecs {
		ProductSpecKeyByProductSpecId[spec.ProductSpecId] = spec
	}
	
	var (
		amount           float64
		productAmount    float64
		orderPromotion   float64
		productPromotion float64
		couponPromotion  float64
		promotionList    []string
		costFreight      float64
		weight           float64
	)
	delCartIds := make([]uint64, 0, 32)
	list := make([]map[string]interface{}, 0, len(cartList.Carts))
	for k := range cartList.Carts {
		if _, ok := ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId]; !ok {
			delCartIds = append(delCartIds, cartList.Carts[k].CartId)
			continue
		}
		buf := map[string]interface{}{
			"cart_id":         cartList.Carts[k].CartId,
			"member_id":       cartList.Carts[k].MemberId,
			"product_id":      cartList.Carts[k].ProductId,
			"product_spec_id": cartList.Carts[k].ProductSpecId,
			"nums":            cartList.Carts[k].Nums,
			"checked":         cartList.Carts[k].IsSelect,
			//"product":         ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId],
			"image":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Image,
			"attr_val": ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Sku,
			"title":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Product.Name,
			"stock":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Stock,
			"price":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Price,
		}
		
		// 总金额
		amount, _ = decimal.NewFromFloat(amount).Add(decimal.NewFromFloat(ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].OldPrice)).Float64()
		
		// 商品总金额
		productAmount, _ = decimal.NewFromFloat(productAmount).Add(decimal.NewFromFloat(ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Price)).Float64()
		
		//订单促销金额
		//商品促销金额
		//优惠券优惠金额
		//促销列表
		//运费
		
		//商品总重
		weight, _ = decimal.NewFromFloat(productAmount).Add(decimal.NewFromFloat(ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Weight)).Float64()
		
		list = append(list, buf)
	}
	
	cartMap := map[string]interface{}{
		"amount":            amount,
		"product_amount":    productAmount,
		"order_promotion":   orderPromotion,
		"product_promotion": productPromotion,
		"coupon_promotion":  couponPromotion,
		"promotion_list":    promotionList,
		"cost_freight":      costFreight,
		"weight":            weight, //商品总重
		"carts":             list,
	}
	
	if len(delCartIds) > 0 {
		// TODO 商品不存在 移除购物车
		gclient.CartClient.DelCart(ctx, &memberpb.DelCartReq{
			CartIds:  delCartIds,
			MemberId: memberId,
		})
	}
	cancel()
	
	return cartMap, nil
}

func (m *Cart) Selected() error {
	cartIds := m.PostForm("cart_ids")
	cartIdList := make([]uint64, 0, 32)
	if err := json.Unmarshal([]byte(cartIds), &cartIdList); err != nil {
		return fmt.Errorf("请选择购物车商品, err: %v", err)
	}
	isSelect, _ := strconv.ParseInt(m.PostForm("is_select"), 10, 32)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.CartClient.SelectCart(ctx, &memberpb.SelectCartReq{
		CartIds:  cartIdList,
		MemberId: memberId,
		IsSelect: memberpb.CartIsSelect(isSelect),
	})
	cancel()
	
	if err != nil {
		return fmt.Errorf("操作购物车失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("操作失败")
	}
	return nil
}
