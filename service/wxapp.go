package service

import (
	"encoding/json"
	"fmt"
	
	"goshop/front-api/model/wxapp"
	"goshop/front-api/pkg/utils"
)

func GetAccessTokenServer(code string) (wxSessionInfo wxapp.WxSessionInfo, err error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		utils.C.WxAppID,
		utils.C.WxAppSecret,
		code,
	)
	
	err = utils.HttpGet(url).ToJson(&wxSessionInfo)
	return
}

//获取微信用户信息
func GetWxUserInfo(sessionKey, encryptedData, iv string) (userInfo wxapp.WxUserInfo, err error) {
	plaintext, err := utils.AesCBCDecrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &userInfo)
	return
}

//获取微信用户绑定的手机号
func GetPhoneNumber(sessionKey, encryptedData, iv string) (wpn wxapp.WxPhoneNumber, err error) {
	plaintext, err := utils.AesCBCDecrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &wpn)
	return
}
