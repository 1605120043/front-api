package filter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	
	"github.com/shinmigo/pb/basepb"
	"goshop/front-api/model/cart"
	"goshop/front-api/model/order"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Order struct {
	validation validation.Validation
	*gin.Context
}

func NewOrder(c *gin.Context) *Order {
	return &Order{Context: c, validation: validation.Validation{}}
}

func (m *Order) Index() (list *order.ListOrderRes, err error) {
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	orderStatus := m.DefaultQuery("order_status", "0")
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("拉取格式不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("拉取格式不正确")
	valid.Match(orderStatus, regexp.MustCompile(`^[0-9]+$`)).Message("订单状态错误")
	if valid.HasError() {
		return nil, valid.GetError()
	}
	
	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	orderStatusNum, _ := strconv.ParseUint(orderStatus, 10, 64)
	listCategoryReq := &orderpb.ListOrderReq{
		Page:        pageNum,
		PageSize:    pageSizeNum,
		MemberId:    memberId,
		OrderStatus: orderpb.OrderStatus(orderStatusNum),
	}
	list, err = service.NewOrder(m.Context).Index(listCategoryReq)
	return
}

func (m *Order) Info() (*order.DetailOrderRes, error) {
	orderId := m.Query("order_id")
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	valid := validation.Validation{}
	valid.Required(orderId).Message("请选择订单！")
	valid.Match(orderId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("订单信息错误！")
	if valid.HasError() {
		return nil, valid.GetError()
	}
	
	orderIdNum, _ := strconv.ParseUint(orderId, 10, 64)
	return service.NewOrder(m.Context).Info(memberId, orderIdNum)
}

func (m *Order) GetUserOrderStatusCount() (list []*order.UserOrderStatusCountRes, err error) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	list, err = service.NewOrder(m.Context).GetUserOrderStatusCount(memberId)
	return
}

func (m *Order) CreateOrder() (*basepb.AnyRes, error) {
	products := m.PostForm("products")
	addressId := m.DefaultPostForm("address_id", "0")
	node := m.DefaultPostForm("node", "")
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	valid := validation.Validation{}
	valid.Required(products).Message("请选择商品！")
	valid.Required(addressId).Message("请选择收货地址！")
	valid.Numeric(addressId).Message("请选择收货地址！")
	
	if valid.HasError() {
		return nil, valid.GetError()
	}
	
	addressIdNum, _ := strconv.ParseUint(addressId, 10, 64)
	
	req := make([]*cart.BuyReq, 0, 32)
	if err := json.Unmarshal([]byte(products), &req); err != nil {
		return nil, fmt.Errorf("参数错误, err: %v", err)
	}
	
	cartIds := make([]uint64, 0, len(req))
	orderProducts := make([]*orderpb.Order_Products, 0, len(req))
	for k := range req {
		if req[k].CartId > 0 {
			cartIds = append(cartIds, req[k].CartId)
		}
		buf := &orderpb.Order_Products{
			ProductId:     req[k].ProductId,
			ProductSpecId: req[k].ProductSpecId,
			Qty:           req[k].Num,
		}
		orderProducts = append(orderProducts, buf)
	}
	
	orderRes, err := service.NewOrder(m.Context).CreateOrder(memberId, addressIdNum, node, orderProducts)
	if err != nil {
		return nil, err
	}
	
	// 清空购物车
	if len(cartIds) > 0 {
		service.NewCart(m.Context).Delete(0, memberId, cartIds)
	}
	return orderRes, nil
}

func (m *Order) CancelOrder() (err error) {
	orderId := m.PostForm("order_id")
	storeId := m.PostForm("store_id")
	storeIdLen := len(storeId)
	
	valid := validation.Validation{}
	valid.Required(orderId).Message("请选择订单！")
	valid.Match(orderId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("订单信息错误！")
	if storeIdLen > 0 {
		valid.Match(storeId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("门店信息错误！")
	}
	if valid.HasError() {
		return valid.GetError()
	}
	
	var storeIdNum uint64
	orderIdNum, _ := strconv.ParseUint(orderId, 10, 64)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	if storeIdLen > 0 {
		storeIdNum, _ = strconv.ParseUint(storeId, 10, 64)
	}
	
	err = service.NewOrder(m.Context).CancelOrder(orderIdNum, memberId, storeIdNum)
	return
}

func (m *Order) DeleteOrder() (err error) {
	orderId := m.PostForm("order_id")
	storeId := m.PostForm("store_id")
	storeIdLen := len(storeId)
	
	valid := validation.Validation{}
	valid.Required(orderId).Message("请选择订单！")
	valid.Match(orderId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("订单信息错误！")
	if storeIdLen > 0 {
		valid.Match(storeId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("门店信息错误！")
	}
	if valid.HasError() {
		return valid.GetError()
	}
	
	var storeIdNum uint64
	orderIdNum, _ := strconv.ParseUint(orderId, 10, 64)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	if storeIdLen > 0 {
		storeIdNum, _ = strconv.ParseUint(storeId, 10, 64)
	}
	
	err = service.NewOrder(m.Context).DeleteOrder(orderIdNum, memberId, storeIdNum)
	return
}
