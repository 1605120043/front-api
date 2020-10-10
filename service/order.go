package service

import (
	"context"
	"fmt"
	"goshop/front-api/model/order"
	"goshop/front-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Order struct {
	*gin.Context
}

func NewOrder(o *gin.Context) *Order {
	return &Order{Context: o}
}

func (o *Order) Index(param *orderpb.ListOrderReq) (list *order.ListOrderRes, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	listRes, err := gclient.OrderClient.GetOrderList(ctx, param)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取订单列表失败， err：%v", err)
	}
	list = &order.ListOrderRes{}
	if listRes == nil {
		return
	}
	if listRes.Total == 0 {
		return
	}

	list.Total = listRes.Total
	listDetail := make([]*order.ListDetailOrderRes, 0, param.PageSize)
	for i := range listRes.Orders {
		listDetail = append(listDetail, &order.ListDetailOrderRes{
			OrderId:     listRes.Orders[i].OrderId,
			GrandTotal:  listRes.Orders[i].GrandTotal,
			OrderStatus: listRes.Orders[i].OrderStatus,
			OrderItems:  listRes.Orders[i].OrderItems,
			CreatedAt:   listRes.Orders[i].CreatedAt,
		})
	}
	list.Orders = listDetail

	return
}

func (o *Order) Info(userId, orderId uint64) (info *order.DetailOrderRes, err error) {
	req := &orderpb.ListOrderReq{
		Page:     1,
		PageSize: 10,
		MemberId: userId,
		OrderId:  orderId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	listRes, err := gclient.OrderClient.GetOrderList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取订单信息失败， err：%v", err)
	}

	if listRes == nil {
		return nil, nil
	}
	if listRes.Total == 0 {
		return nil, nil
	}
	info = &order.DetailOrderRes{
		OrderId:        listRes.Orders[0].OrderId,
		Subtotal:       listRes.Orders[0].Subtotal,
		GrandTotal:     listRes.Orders[0].GrandTotal,
		TotalPaid:      listRes.Orders[0].TotalPaid,
		ShippingAmount: listRes.Orders[0].ShippingAmount,
		DiscountAmount: listRes.Orders[0].DiscountAmount,
		PaymentType:    listRes.Orders[0].PaymentType,
		PaymentStatus:  listRes.Orders[0].PaymentStatus,
		PaymentTime:    listRes.Orders[0].PaymentTime,
		ShippingStatus: listRes.Orders[0].ShippingStatus,
		ShippingTime:   listRes.Orders[0].PaymentTime,
		Confirm:        listRes.Orders[0].Confirm,
		//ConfirmTime:    listRes.Orders[0].ConfigTime,
		OrderStatus:    listRes.Orders[0].OrderStatus,
		RefundStatus:   listRes.Orders[0].RefundStatus,
		ReturnStatus:   listRes.Orders[0].ReturnStatus,
		UserNote:       listRes.Orders[0].UserNote,
		OrderItems:     listRes.Orders[0].OrderItems,
		OrderAddress:   listRes.Orders[0].OrderAddress,
		OrderPayment:   listRes.Orders[0].OrderPayment,
		OrderShipment:  listRes.Orders[0].OrderShipment,
		CreatedAt:      listRes.Orders[0].CreatedAt,
	}

	return
}
