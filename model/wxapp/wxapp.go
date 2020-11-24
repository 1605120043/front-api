package wxapp

type WxSessionInfo struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	WxErrInfo
}

type WxErrInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WxPhoneNumber struct {
	PhoneNumber     string       `json:"phoneNumber"`     //用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string       `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string       `json:"countryCode"`     //区号
	Watermark       *WxWatermark `json:"watermark"`       //数据水印( watermark )
}

//微信加密数据结构
type WxUserInfo struct {
	OpenId    string       `json:"openId"`
	NickName  string       `json:"nickName"`
	Gender    int8         `json:"gender"`
	City      string       `json:"city"`
	Province  string       `json:"province"`
	Country   string       `json:"country"`
	AvatarUrl string       `json:"avatarUrl"`
	UnionId   string       `json:"unionId"`
	Watermark *WxWatermark `json:"watermark"` //数据水印( watermark )
}

type WxWatermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type WXPayNotify struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int64  `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int64  `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	CouponFee     int64  `xml:"coupon_fee"`
	CouponCount   int64  `xml:"coupon_count"`
	CouponID0     string `xml:"coupon_id_0"`
	CouponFee0    int64  `xml:"coupon_fee_0"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}
