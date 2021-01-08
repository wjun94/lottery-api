package models

import (
	"fmt"
	"io/ioutil"
	"mango-api/config"

	"github.com/medivhzhan/weapp/v2"
)

// GetQRCode 获取小程序二维码
func GetQRCode(param string) []byte {
	appID := config.Config.App.AppID
	appSecret := config.Config.App.AppSecret
	token, err := weapp.GetAccessToken(appID, appSecret)
	if err != nil {
		// 处理一般错误信息
		return nil
	}
	// if err := token.GetResponseError(); err != nil {
	// 	// 处理微信返回错误信息
	// 	fmt.Println(err)
	// 	return nil
	// }
	// fmt.Printf("返回结果: %#v", token)
	getter := weapp.UnlimitedQRCode{
		// Scene:     "id=" + param,
		Scene:     param,
		Page:      "pages/index/index",
		Width:     430,
		AutoColor: true,
		LineColor: weapp.Color{"0", "0", "0"},
		IsHyaline: true,
	}
	resp, res, err := getter.Get(token.AccessToken)
	if err != nil {
		// 处理一般错误信息
		fmt.Println(err)
		return nil
	}
	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return content
}
