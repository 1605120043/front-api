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
