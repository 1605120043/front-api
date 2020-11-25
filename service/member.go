package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/shopspring/decimal"
	"goshop/front-api/model/wxapp"
	"goshop/front-api/pkg/grpc/gclient"
	"goshop/front-api/pkg/utils"
	
	"github.com/shinmigo/pb/orderpb"
	"github.com/shinmigo/pb/shoppb"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
)

type Member struct {
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c}
}

func (m *Member) Info() (*memberpb.LoginRes, error) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	row, err := gclient.MemberClient.GetMemberForLogin(ctx, &memberpb.MemberIdReq{MemberId: memberId})
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取失败， err：%v", err)
	}
	return row, nil
}

func (m *Member) Update(req *memberpb.Member) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	
	res, err := gclient.MemberClient.EditMember(ctx, req)
	cancel()
	if err != nil {
		return false, fmt.Errorf("修改信息失败， err：%v", err)
	}
	if res.State == 0 {
		return false, fmt.Errorf("修改信息失败， err：%v", err)
	}
	
	return true, nil
}

func (m *Member) Pay(memberId, orderId uint64, paymentCode, tradeType string) (map[string]string, error) {
	paymentCodeVal := memberpb.PaymentCode_value[paymentCode]
	if paymentCodeVal == 0 {
		return nil, fmt.Errorf("选择正确的支付方式")
	}
	
	var isOpen bool
	for _, k := range utils.Payments.Payment {
		if k.Code == paymentCode && k.Status == int32(shoppb.PaymentStatus_Open) {
			isOpen = true
		}
	}
	
	if !isOpen {
		return nil, fmt.Errorf("未开启该支付方式")
	}
	
	var openid string
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if tradeType == "JSAPI" {
		// 获取用openid
		memberThird, err := gclient.MemberClient.GetMemberOpenid(ctx, &memberpb.MemberIdReq{MemberId: memberId})
		if err != nil {
			return nil, fmt.Errorf("获取openid失败， err：%v", err)
		}
		openid = memberThird.OpenId
	}
	
	orderRes, err := gclient.OrderClient.GetOrderList(ctx, &orderpb.ListOrderReq{
		Page:     1,
		PageSize: 1,
		MemberId: memberId,
		OrderId:  orderId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取订单信息失败， err：%v", err)
	}
	if orderRes.Total == 0 {
		return nil, fmt.Errorf("获取订单信息失败， err：%v", err)
	}
	
	money := orderRes.Orders[0].GrandTotal
	idStr := strconv.FormatUint(orderId, 10)
	params := make([]*memberpb.PaymentParams, 0, 1)
	params = append(params, &memberpb.PaymentParams{
		SourceId: idStr,
		Money:    money,
	})
	req := memberpb.ToAdd{
		MemberId:    memberId,
		Type:        memberpb.PaymentType_Order, // 当前只有订单支付
		PaymentCode: memberpb.PaymentCode(paymentCodeVal),
		Ip:          utils.GetClientIp(),
		Params:      params,
	}
	
	res, err := gclient.MemberPaymentClient.AddPay(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("发起支付失败， err：%v", err)
	}
	if res.State == 0 {
		return nil, fmt.Errorf("发起发起失败， err：%v", err)
	}
	
	if money == 0 {
		// 直接支付成功
		r, err := gclient.MemberPaymentClient.EditPay(ctx, &memberpb.ToEdit{
			PaymentId:   res.PaymentId,
			Status:      memberpb.PaymentStatus_PaySuccess,
			PaymentCode: memberpb.PaymentCode(paymentCodeVal),
			Money:       0,
			PayedMsg:    "金额为0，自动支付成功",
		})
		
		if err != nil {
			return nil, fmt.Errorf("支付失败， err：%v", err)
		}
		
		if r.State == 0 {
			return nil, fmt.Errorf("支付失败， err：%v", err)
		}
		
		return nil, nil
	}
	cancel()
	
	// 打开支付
	if memberpb.PaymentCode(paymentCodeVal) == memberpb.PaymentCode_Wechat {
		return WechatPay(res.PaymentId, tradeType, money, openid)
	}
	
	if memberpb.PaymentCode(paymentCodeVal) == memberpb.PaymentCode_Alipay {
		return AliPay(res.PaymentId, tradeType, money)
	}
	
	return nil, nil
}

func (m *Member) WxNotify(wxn wxapp.WXPayNotify) error {
	if wxn.ReturnCode != "SUCCESS" { // return_code 表示通信状态，不代表支付状态
		return fmt.Errorf("通信失败，请稍后再通知我")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	buf := decimal.NewFromFloat(float64(wxn.TotalFee)).Div(decimal.NewFromFloat(100))
	money, _ := buf.Float64()
	var (
		status    int32
		statusMsg string
	)
	if wxn.ResultCode == "SUCCESS" { // 支付成功
		status = int32(memberpb.PaymentStatus_PaySuccess)
		statusMsg = "支付成功"
	} else {
		// 支付失败
		status = int32(memberpb.PaymentStatus_StatusOther)
		statusMsg = "支付失败"
	}
	
	res, err := gclient.MemberPaymentClient.EditPay(ctx, &memberpb.ToEdit{
		PaymentId:   wxn.OutTradeNo,
		Status:      memberpb.PaymentStatus(status),
		PaymentCode: 1,
		Money:       money,
		PayedMsg:    statusMsg,
		TradeNo:     wxn.OutTradeNo,
	})
	
	if err != nil {
		return fmt.Errorf("支付失败， err：%v", err)
	}
	
	if res.State == 0 {
		return fmt.Errorf("支付失败， err：%v", err)
	}
	cancel()
	return nil
}
