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
	isSelect := m.PostForm("is_select")
	isSelectVal := int32(1)
	if isSelect == "false" {
		isSelectVal = 0
	}
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	isPlus := m.DefaultPostForm("is_plus", "true") // 默认是叠加
	isPlusVal := true
	if isPlus == "false" {
		isPlusVal = false
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductClient.GetProductListByProductSpecIds(ctx, &productpb.ProductSpecIdsReq{
		ProductId:     []uint64{productId},
		ProductSpecId: []uint64{productSpecId},
	})
	
	if err != nil {
		return fmt.Errorf("添加购物车失败, err:%v", err)
	}
	productSpec := resp.Products[0].Spec
	
	if len(productSpec) == 0 {
		return fmt.Errorf("商品不存在")
	}
	
	if resp.Products[0].Status != productpb.ProductStatus_Disabled {
		return fmt.Errorf("商品:%s, 已下架", resp.Products[0].Name)
	}
	
	if int64(productSpec[0].Stock) < nums {
		return fmt.Errorf("商品:%s (%s), 库存不足", resp.Products[0].Name, productSpec[0].Sku)
	}
	
	// 加入购物车
	addCart, err := gclient.CartClient.AddCart(ctx, &memberpb.AddCartReq{
		MemberId:      memberId,
		ProductId:     productId,
		ProductSpecId: productSpecId,
		IsSelect:      memberpb.CartIsSelect(isSelectVal),
		Nums:          nums,
		IsPlus:        isPlusVal, //true是叠加
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
	isAll, _ := strconv.ParseInt(m.DefaultPostForm("is_all", "0"), 10, 8)
	
	cartIdList := make([]uint64, 0, 32)
	if isAll == 0 { //根据id删除
		cartIds := m.PostForm("cart_ids")
		
		if err := json.Unmarshal([]byte(cartIds), &cartIdList); err != nil {
			return fmt.Errorf("请选择购物车商品, err: %v", err)
		}
	}
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.CartClient.DelCart(ctx, &memberpb.DelCartReq{
		IsAll:    int32(isAll),
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
	
	cartLen := len(cartList.Carts)
	productIds := make([]uint64, 0, cartLen)
	productSpecIds := make([]uint64, 0, cartLen)
	for k := range cartList.Carts {
		productIds = append(productIds, cartList.Carts[k].ProductId)
		productSpecIds = append(productSpecIds, cartList.Carts[k].ProductSpecId)
	}
	productList, err := gclient.ProductClient.GetProductListByProductSpecIds(ctx, &productpb.ProductSpecIdsReq{
		ProductId:     productIds,
		ProductSpecId: productSpecIds,
	})
	
	if err != nil {
		return nil, fmt.Errorf("获取购物车列表失败, err:%v", err)
	}
	
	type ProductSpec struct {
		Detail *productpb.ProductDetail
		Spec   *productpb.ProductSpec
	}
	ProductSpecKeyByProductSpecId := make(map[uint64]ProductSpec, 32)
	for _, p := range productList.Products {
		// 无规格跳过
		if len(p.Spec) == 0 {
			continue
		}
		for _, s := range p.Spec {
			ProductSpecKeyByProductSpecId[s.ProductSpecId] = ProductSpec{
				Detail: p,
				Spec:   s,
			}
		}
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
		var checked bool
		if cartList.Carts[k].IsSelect == 1 {
			checked = true
		}
		buf := map[string]interface{}{
			"cart_id":         cartList.Carts[k].CartId,
			"member_id":       cartList.Carts[k].MemberId,
			"product_id":      cartList.Carts[k].ProductId,
			"product_spec_id": cartList.Carts[k].ProductSpecId,
			"nums":            cartList.Carts[k].Nums,
			"checked":         checked,
			//"product":         ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId],
			"image":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Detail.Images[0],
			"attr_val": ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec.Sku,
			"title":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Detail.Name,
			"stock":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec.Stock,
			"price":    ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec.Price,
		}
		
		//总金额
		amount, _ = decimal.NewFromFloat(amount).Add(decimal.NewFromFloat(ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec.OldPrice)).Float64()
		
		// 商品总金额
		productAmount, _ = decimal.NewFromFloat(productAmount).Add(decimal.NewFromFloat(ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec.Price)).Float64()
		
		//订单促销金额
		//商品促销金额
		//优惠券优惠金额
		//促销列表
		//运费
		
		//商品总重
		weight, _ = decimal.NewFromFloat(productAmount).Add(decimal.NewFromFloat(ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec.Weight)).Float64()
		
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

func (m *Cart) Checked() error {
	cartChecked := m.PostForm("cart_checked")
	buf := make([]*memberpb.SelectCart, 0, 32)
	if err := json.Unmarshal([]byte(cartChecked), &buf); err != nil {
		return fmt.Errorf("请选择购物车商品, err: %v", err)
	}
	
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.CartClient.SelectCart(ctx, &memberpb.SelectCartReq{
		SelectCart: buf,
		MemberId:   memberId,
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
