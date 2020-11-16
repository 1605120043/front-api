package utils

import (
	aliPayClient "github.com/shinmigo/gopay/alipay/kernel"
	wxClient "github.com/shinmigo/gopay/wxpay/kernel"
)

var (
	C            *Config
	Payments     PaymentConfigList
	AliPayConf   *AliPayConfig
	AliPayClient *aliPayClient.AliPayClient
	WxPayConf    *WxPayConfig
	WxPayClient  *wxClient.WxClient
)

type Config struct {
	*Base
	*Redis
	*Mysql
	*Etcd
	*Grpc
	*Sms
	*WxApp
}

type Base struct {
	Name    string
	Version string
	Webhost string
}

type Redis struct {
	Host     string
	Password string
	Database int
}

type Mysql struct {
	Host     string
	User     string
	Password string
	Database string
}

type Etcd struct {
	Host []string
}

type Grpc struct {
	Name map[string]string
	Host string
}

type PaymentConfigList struct {
	Payment []*PaymentConfig
}

type PaymentConfig struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int32  `json:"status"`
}

type AliPayConfig struct {
	AppId         string
	NotifyUrl     string
	EncryptKey    string
	IsProd        bool
	LocalTimeZone string
	SignType      string
}

type WxPayConfig struct {
	AppId     string
	MchId     string
	NotifyUrl string
	Md5key    string
	IsProd    bool
}

type Sms struct {
	Url    string
	Params string
}

type WxApp struct {
	WxAppID     string
	WxAppSecret string
}
