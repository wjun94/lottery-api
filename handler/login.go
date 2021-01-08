package handler

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月17日
 * 修改时间：2019年10月18日
 * 描述信息：登入
 */
import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/medivhzhan/weapp"

	"mango-api/config"
	"mango-api/models"
	"mango-api/response"
)

// 客户端请求的js_code
type wxLogin struct {
	Code string `json:"code"`
}

// LoginWxHandler 微信登入
func LoginWxHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req wxLogin
	)
	json.Unmarshal(body, &req)
	appID := config.Config.App.AppID
	appSecret := config.Config.App.AppSecret
	loginInfo, err := weapp.Login(appID, appSecret, req.Code)
	if err != nil {
		// 处理一般错误信息
		fmt.Println(err)
		return
	}
	result := models.SelectUser(loginInfo.OpenID)
	response.ResultSuccess(w, result)
}

// CreateWxUser 创建微信用户信息
func CreateWxUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req  wxLogin
		user models.User
	)
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &user)
	appID := config.Config.App.AppID
	appSecret := config.Config.App.AppSecret
	loginInfo, err := weapp.Login(appID, appSecret, req.Code)
	if err != nil {
		// 处理一般错误信息
		fmt.Println(err)
		return
	}
	user.WxID = loginInfo.OpenID
	result := models.InsertUser(user)
	response.ResultSuccess(w, result)
}

// LoginAdminHandler 后台登入
func LoginAdminHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		Name     string `validate:"required" label:"账号"`
		Password string `validate:"required,min=6,max=11" label:"密码"`
	}
	var (
		user models.Account
		info models.LoginInfo
		vf   verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}

	json.Unmarshal(body, &user)
	info.IP = r.RemoteAddr
	result := models.FindAdminUser(user, info)
	response.ResultSuccess(w, result)
}
