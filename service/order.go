package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	
	"github.com/shinmigo/pb/basepb"
	"goshop/front-api/model/order"
	"goshop/front-api/pkg/grpc/gclient"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Order struct {
	*gin.Context
}

func NewOrder(o *gin.Context) *Order {
	return &Order{Context: o}
}

func getOrderStatusName(status orderpb.OrderStatus) string {
	orderStatusMap := make(map[orderpb.OrderStatus]string, 6)
	
	orderStatusMap[orderpb.OrderStatus_PendingPayment] = "待付款"
	orderStatusMap[orderpb.OrderStatus_PendingReview] = "待发货"
	orderStatusMap[orderpb.OrderStatus_PendingShipment] = "待发货"
	orderStatusMap[orderpb.OrderStatus_PendingReceiving] = "待收货"
	orderStatusMap[orderpb.OrderStatus_PendingComment] = "已完成"
	orderStatusMap[orderpb.OrderStatus_Completed] = "待评价"
	
	return orderStatusMap[status]
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
			OrderId:         listRes.Orders[i].OrderId,
			GrandTotal:      listRes.Orders[i].GrandTotal,
			TotalQtyOrdered: listRes.Orders[i].TotalQtyOrdered,
			OrderStatus:     listRes.Orders[i].OrderStatus,
			OrderStatusName: getOrderStatusName(listRes.Orders[i].OrderStatus),
			OrderItems:      listRes.Orders[i].OrderItems,
			CreatedAt:       listRes.Orders[i].CreatedAt,
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
		return nil, fmt.Errorf("获取订单信息失败")
	}
	if listRes.Total == 0 {
		return nil, fmt.Errorf("获取订单信息失败")
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
		OrderStatus:   listRes.Orders[0].OrderStatus,
		OrderStatusName: getOrderStatusName(listRes.Orders[0].OrderStatus),
		RefundStatus:  listRes.Orders[0].RefundStatus,
		ReturnStatus:  listRes.Orders[0].ReturnStatus,
		UserNote:      listRes.Orders[0].UserNote,
		OrderItems:    listRes.Orders[0].OrderItems,
		OrderAddress:  listRes.Orders[0].OrderAddress,
		OrderShipment: listRes.Orders[0].OrderShipment,
		CreatedAt:     listRes.Orders[0].CreatedAt,
	}
	
	return
}

func (o *Order) GetUserOrderStatusCount(userId uint64) (orderStatusCountList []*order.UserOrderStatusCountRes, err error) {
	req := &orderpb.GetOrderStatusReq{
		MemberId: userId,
	}
	
	orderStatusCountList = make([]*order.UserOrderStatusCountRes, 0, 16)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	info, err := gclient.OrderClient.GetOrderStatus(ctx, req)
	cancel()
	if err != nil {
		return orderStatusCountList, fmt.Errorf("获取订单信息失败， err：%v", err)
	}
	if info == nil || info.OrderStatistics == nil {
		return
	}
	for i := range info.OrderStatistics {
		orderStatusCountList = append(orderStatusCountList, &order.UserOrderStatusCountRes{
			OrderStatus: info.OrderStatistics[i].OrderStatus,
			Count:       info.OrderStatistics[i].Count,
		})
	}
	
	return
}

func (o *Order) CreateOrder(memberId, addressId uint64, node string, products []*orderpb.Order_Products) (*basepb.AnyRes, error) {
	req := orderpb.Order{
		MemberId:  memberId,
		AddressId: addressId,
		UserNode:  node,
		Products:  products,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.OrderClient.AddOrder(ctx, &req)
	cancel()
	
	if err != nil {
		return nil, err
	}
	if res == nil || res.Id == 0 {
		return nil, errors.New("订单创建失败！")
	}
	
	return res, nil
}

func (o *Order) CancelOrder(orderId, memberId, storeId uint64) error {
	req := orderpb.CancelOrderReq{
		MemberId: memberId,
		OrderId:  []uint64{orderId},
		StoreId:  storeId,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.OrderClient.CancelOrder(ctx, &req)
	cancel()
	if err != nil {
		return err
	}
	if res == nil || res.Id == 0 {
		return errors.New("取消订单失败！")
	}
	
	return nil
}

func (o *Order) DeleteOrder(orderId, memberId, storeId uint64) error {
	req := orderpb.DeleteOrderReq{
		MemberId: memberId,
		OrderId:  []uint64{orderId},
		StoreId:  storeId,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.OrderClient.DeleteOrder(ctx, &req)
	cancel()
	if err != nil {
		return err
	}
	if res == nil || res.Id == 0 {
		return errors.New("删除订单失败！")
	}
	
	return nil
}
