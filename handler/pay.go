package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mango-api/config"
	"mango-api/response"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
)

// Config 配置文件
var Config = config.Config.App

// GetOpenID 获取小程序openID
func GetOpenID(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var req struct {
		Code string `json:"code"`
	}
	json.Unmarshal(body, &req)
	u, err := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	paras.Set("appid", Config.AppID)
	paras.Set("secret", Config.AppSecret)
	paras.Set("js_code", req.Code)
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	jMap := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&jMap)
	// err = json.NewDecoder(resp.Body).Decode(&jMap)
	// if err != nil {
	// 	fmt.Fprintln(w, "request openid response json parse err :")
	// }
	// println(jMap)
	// if jMap["openid"] == nil {
	// 	fmt.Fprintln(w, w.Write(errors.New("request openid response json parse err :"+err.Error())))
	// }
	response, _ := json.Marshal(jMap)
	w.Write(response)
}

// WxPay 微信支付
func WxPay(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var req struct {
		Code     string  `json:"code"`     // 登入二维码
		NonceStr string  `json:"nonceStr"` // 订单号
		Body     string  `json:"body"`     // 支付订单的商品名
		TotalFee float64 `json:"totalFee"` // 价格
	}
	json.Unmarshal(body, &req)
	sessionRsp, _ := wechat.Code2Session(Config.AppID, Config.AppSecret, req.Code)
	client := wechat.NewClient(Config.AppID, Config.MchID, Config.Key, true)
	client.SetCountry(wechat.China)
	// number := gopay.GetRandomString(16)
	bm := make(gopay.BodyMap)
	outTradeNo := gopay.GetRandomString(32)
	bm.Set("nonce_str", req.NonceStr) // 随机字符串 订单号
	bm.Set("body", req.Body)
	bm.Set("out_trade_no", outTradeNo) // 商户订单号
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://www.gopay.ink")
	// bm.Set("notify_url", "https://cnicu.cn")
	bm.Set("device_info", "WEB")
	totalF, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", req.TotalFee*100), 64)
	bm.Set("total_fee", totalF)
	bm.Set("trade_type", wechat.TradeType_Mini)
	bm.Set("openid", sessionRsp.Openid)

	sign := wechat.GetParamSign(Config.AppID, Config.MchID, Config.Key, bm)
	bm.Set("sign", sign)
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println("wxRsp:", *wxRsp)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	//获取小程序支付需要的paySign
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := wechat.GetMiniPaySign(Config.AppID, req.NonceStr, pac, wechat.SignType_MD5, timeStamp, Config.Key)
	// fmt.Println("paySign:", paySign)
	result := make(map[string]interface{})
	result["paySign"] = paySign
	result["package"] = pac
	result["timeStamp"] = timeStamp
	result["nonceStr"] = req.NonceStr
	result["outTradeNo"] = outTradeNo
	responseJSON := response.JSON(result)
	response.ResultSuccess(w, responseJSON)
}
