package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	alipayKernel "github.com/shinmigo/gopay/alipay/kernel"
	wxpayKernel "github.com/shinmigo/gopay/wxpay/kernel"
	"github.com/shinmigo/pb/shoppb"
	"goshop/front-api/pkg/grpc/gclient"
	"goshop/front-api/pkg/utils"
)

func GetPaymentConfig() {
	for {
		//每5分钟动态加载一次配置
		GetPaymentListGrpc()
		if len(utils.Payments.Payment) > 0 {
			time.Sleep(5 * time.Minute)
		} else {
			spew.Dump("第一次加载支付方式=========")
			time.Sleep(3 * time.Second)
		}
	}
}

func GetPaymentListGrpc() {
	if gclient.PaymentClient == nil {
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := gclient.PaymentClient.GetPaymentList(ctx, &shoppb.ListPaymentReq{
		Status: shoppb.PaymentStatus_Open,
	})
	if err != nil {
		//log.Fatal("获取支付方式失败， err：%v", err)
		fmt.Printf("获取支付方式失败，支付服务没有开启，请检查，err：%v", err)
		return
	}
	cancel()
	
	configList := make([]*utils.PaymentConfig, 0, len(res.Payments))
	for k := range res.Payments {
		switch res.Payments[k].Code {
		case "Wechat":
			wxpay(res.Payments[k].Params)
			break
		
		case "Alipay":
			alipay(res.Payments[k].Params)
			break
		
		default:
			break
		}
		configList = append(configList, &utils.PaymentConfig{
			Name:   res.Payments[k].Name,
			Code:   res.Payments[k].Code,
			Status: int32(res.Payments[k].Status),
		})
	}
	
	utils.Payments.Payment = configList
	return
}

func alipay(params string) {
	buf := map[string]string{}
	if err := json.Unmarshal([]byte(params), &buf); err != nil {
		return
	}
	
	aliPayConf := &utils.AliPayConfig{
		AppId:         buf["appid"],
		NotifyUrl:     buf["notify_url"],
		EncryptKey:    buf["encrypt_key"],
		IsProd:        true,
		LocalTimeZone: "PRC",
		SignType:      "RSA",
	}
	
	aliPayClient, err := alipayKernel.NewAliPayClient(&alipayKernel.Config{
		AppId:                  aliPayConf.AppId,
		AliPayPublicKeyPath:    "./pkg/cert/ali/rsa_ali_public_key.pem",
		MerchantPrivateKeyPath: "./pkg/cert/ali/rsa_private_key.pem",
		NotifyUrl:              aliPayConf.NotifyUrl,
		EncryptKey:             aliPayConf.EncryptKey,
		IsProd:                 aliPayConf.IsProd,
		LocalTimeZone:          aliPayConf.LocalTimeZone,
		SignType:               aliPayConf.SignType,
	})
	if err != nil || aliPayClient == nil {
		panic(fmt.Sprintf("初始化支付宝客户端出错, err: %v", err))
	}
	
	utils.AliPayConf = aliPayConf
	utils.AliPayClient = aliPayClient
	return
}

func wxpay(params string) {
	buf := map[string]string{}
	if err := json.Unmarshal([]byte(params), &buf); err != nil {
		return
	}
	
	wxPayConf := &utils.WxPayConfig{
		AppId:     buf["appid"],
		MchId:     buf["mch_id"],
		NotifyUrl: buf["notify_url"],
		Md5key:    buf["md5_key"],
		IsProd:    true,
	}
	wxPayClient := wxpayKernel.NewWxClient(wxPayConf.AppId, wxPayConf.MchId, wxPayConf.Md5key, wxPayConf.IsProd)
	if wxPayClient == nil {
		panic(fmt.Sprintf("初始化微信客户端出错"))
	}
	
	utils.WxPayConf = wxPayConf
	utils.WxPayClient = wxPayClient
	return
}
