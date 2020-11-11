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
	"goshop/front-api/model/cart"
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

func (m *Cart) Delete(isAll int8, memberId uint64, cartIdList []uint64) error {
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

func (m *Cart) Index() (*cart.Carts, error) {
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
		amount    float64
		promotion float64
	)
	delCartIds := make([]uint64, 0, 32)
	cartProducts := make([]*cart.CartProducts, 0, len(cartList.Carts))
	for k := range cartList.Carts {
		if _, ok := ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId]; !ok {
			delCartIds = append(delCartIds, cartList.Carts[k].CartId)
			continue
		}
		specStruct := ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Spec
		
		// 商品总金额
		productAmount, _ := decimal.NewFromFloat(specStruct.Price).Mul(decimal.NewFromFloat(float64(cartList.Carts[k].Nums))).Float64()
		
		var checked bool
		if cartList.Carts[k].IsSelect == 1 {
			checked = true
			amount, _ = decimal.NewFromFloat(amount).Add(decimal.NewFromFloat(float64(productAmount))).Float64()
		}
		
		buf := &cart.CartProducts{
			CartId:        cartList.Carts[k].CartId,
			ProductId:     cartList.Carts[k].ProductId,
			ProductSpecId: cartList.Carts[k].ProductSpecId,
			Stock:         specStruct.Stock,
			ProductName:   ProductSpecKeyByProductSpecId[cartList.Carts[k].ProductSpecId].Detail.Name,
			SpecName:      specStruct.Sku,
			Image:         specStruct.Image,
			Price:         specStruct.Price,
			Num:           cartList.Carts[k].Nums,
			Checked:       checked,
			ProductAmount: productAmount,
		}
		cartProducts = append(cartProducts, buf)
		
	}
	
	if len(delCartIds) > 0 {
		// TODO 商品不存在 移除购物车
		gclient.CartClient.DelCart(ctx, &memberpb.DelCartReq{
			CartIds:  delCartIds,
			MemberId: memberId,
		})
	}
	cancel()
	
	return &cart.Carts{
		Amount:    amount,
		Promotion: promotion,
		Products:  cartProducts,
	}, nil
}

func (m *Cart) Count() (total uint64) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.CartClient.GetCartCountByMemberId(ctx, &memberpb.ListCartReq{
		MemberId: memberId,
	})
	cancel()
	if err != nil {
		return 0
	}
	total = resp.Count
	return
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

//结算计算
func (m *Cart) Buy(req []*cart.BuyReq) (*cart.BuyRes, error) {
	productIds := make([]uint64, 0, len(req))
	productSpecId := make([]uint64, 0, len(req))
	reqMap := make(map[uint64]map[uint64]uint64, 0)
	for k := range req {
		if req[k].Num == 0 {
			continue
		}
		productIds = append(productIds, req[k].ProductId)
		productSpecId = append(productSpecId, req[k].ProductSpecId)
		if _, ok := reqMap[req[k].ProductId]; ok {
			reqMap[req[k].ProductId][req[k].ProductSpecId] = req[k].Num
		} else {
			buf := make(map[uint64]uint64)
			buf[req[k].ProductSpecId] = req[k].Num
			reqMap[req[k].ProductId] = buf
		}
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.ProductClient.GetProductListByProductSpecIds(ctx, &productpb.ProductSpecIdsReq{
		ProductId:     productIds,
		ProductSpecId: productSpecId,
	})
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("获取商品失败, err:%v", err)
	}
	
	if len(res.Products) == 0 {
		return nil, fmt.Errorf("获取商品失败")
	}
	
	buyProducts := make([]*cart.BuyProducts, 0, len(res.Products))
	var (
		orderAmount     float64
		orderPromotion  float64
		couponPromotion float64
		orderWeight     float64
		costFreight     float64
		payAmount       float64
	)
	for k := range res.Products {
		p := res.Products[k]
		// 找不到规格跳过
		if len(p.Spec) == 0 {
			continue
		}
		
		for s := range p.Spec {
			num := reqMap[res.Products[k].ProductId][p.Spec[s].ProductSpecId]
			//商品总重
			productWeight, _ := decimal.NewFromFloat(p.Spec[s].Weight).Mul(decimal.NewFromFloat(float64(num))).Float64()
			//商品总金额
			productAmount, _ := decimal.NewFromFloat(p.Spec[s].Price).Mul(decimal.NewFromFloat(float64(num))).Float64()
			
			buf := &cart.BuyProducts{
				ProductId:        p.ProductId,
				ProductSpecId:    p.Spec[s].ProductSpecId,
				Stock:            p.Spec[s].Stock,
				ProductName:      p.Name,
				SpecName:         p.Spec[s].Sku,
				Image:            p.Spec[s].Image,
				Price:            p.Spec[s].Price,
				Num:              reqMap[res.Products[k].ProductId][p.Spec[s].ProductSpecId],
				Weight:           p.Spec[s].Weight,
				ProductWeight:    productWeight,
				ProductAmount:    productAmount,
				ProductPromotion: 0,
			}
			if p.Spec[s].Stock < reqMap[res.Products[k].ProductId][p.Spec[s].ProductSpecId] {
				// 库存不足
				buf.Error = fmt.Errorf("库存不足")
				continue
			}
			buyProducts = append(buyProducts, buf)
			
			// 订单总金额
			orderAmount, _ = decimal.NewFromFloat(orderAmount).Add(decimal.NewFromFloat(float64(productAmount))).Float64()
			// 订单总重
			orderWeight, _ = decimal.NewFromFloat(orderWeight).Add(decimal.NewFromFloat(float64(productWeight))).Float64()
		}
	}
	
	// 实付金额 = 商品金额加上运费金额减去优惠金额
	payAmount, _ = decimal.NewFromFloat(orderAmount).Add(decimal.NewFromFloat(float64(costFreight))).Float64()
	
	buyRes := &cart.BuyRes{
		OrderAmount:     orderAmount,
		OrderPromotion:  orderPromotion,
		CouponPromotion: couponPromotion,
		OrderWeight:     orderWeight,
		CostFreight:     costFreight,
		PromotionList:   []string{},
		CouponList:      []string{},
		Products:        buyProducts,
		PayAmount:       payAmount,
	}
	return buyRes, nil
}
