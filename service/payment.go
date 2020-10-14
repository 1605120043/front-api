package service

import (
	"fmt"
	"time"
	
	"github.com/davecgh/go-spew/spew"
	aliPayment "github.com/shinmigo/gopay/alipay/payment"
	wxPayment "github.com/shinmigo/gopay/wxpay/payment"
	"github.com/shopspring/decimal"
	"goshop/front-api/pkg/utils"
)

func WechatPay(paymentId, tradeType string, money float64) (map[string]string, error) {
	totalFee, _ := decimal.NewFromFloat(money).Mul(decimal.NewFromFloat(float64(100))).Float64()
	wxPaymentClient := wxPayment.Payment{Client: utils.WxPayClient}
	
	payRes, err := wxPaymentClient.Pay(&wxPayment.Trade{
		Body:       "支付订单",
		Detail:     "",
		OutTradeNo: paymentId,
		TotalFee:   uint64(totalFee),
		TradeType:  tradeType,
		NotifyUrl:  "www.xxx.com/notify",
	})
	if err != nil {
		return nil, err
	}
	
	time_stamp := time.Now().String()
	
	prePayParams := map[string]string{
		"appid":     utils.WxPayConf.AppId,
		"partnerid": utils.WxPayConf.MchId,
		"prepayid":  payRes.PrepayId,
		"noncestr":  payRes.NonceStr,
		"package":   "Sign=WXPay",
		"timestamp": time_stamp,
	}
	
	return prePayParams, nil
}

func AliPay(paymentId, tradeType string, money float64) (map[string]string, error) {
	totalAmount := fmt.Sprintf("%f", money)
	paymentTrade := aliPayment.Payment{Client: utils.AliPayClient}
	switch tradeType {
	case "Wap":
		//手机网站支付
		payRes, err := paymentTrade.Wap(&aliPayment.Wap{
			Trade: aliPayment.Trade{
				Subject:     "支付订单",
				OutTradeNo:  paymentId,
				TotalAmount: totalAmount,
			},
		})
		if err != nil {
			return nil, err
		}
		spew.Dump(payRes)
		break
	case "App":
		//APP支付
		payRes, err := paymentTrade.App(&aliPayment.App{
			Trade: aliPayment.Trade{
				Subject:     "支付订单",
				OutTradeNo:  paymentId,
				TotalAmount: totalAmount,
			},
		})
		if err != nil {
			return nil, err
		}
		spew.Dump(payRes)
		break
	case "PC":
		//PC网站支付
		payRes, err := paymentTrade.Page(&aliPayment.Page{
			Trade: aliPayment.Trade{
				Subject:     "支付订单",
				OutTradeNo:  paymentId,
				TotalAmount: totalAmount,
			},
		})
		if err != nil {
			return nil, err
		}
		spew.Dump(payRes)
		break
	default:
		break
		
	}
	return nil, nil
}
