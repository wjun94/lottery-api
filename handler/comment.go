package handler

import (
	"encoding/json"
	"io/ioutil"
	"mango-api/response"
	"net/http"

	"mango-api/models"
)

// CreateComment 创建评论
func CreateComment(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		OrderID     string `validate:"required" label:"订单号"`
		CommodityID string `validate:"required" label:"商品ID"`
		Product     int    `validate:"required" label:"产品评级"`
		Logistics   int    `validate:"required" label:"物流评级"`
		Detail      string `validate:"required" label:"描述"`
	}
	var (
		req models.Comment
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	req.UserID = md["UserID"]
	result := models.InsertComment(req)
	response.ResultSuccess(w, result)
}

// GetComment 查看评论
func GetComment(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["orderId"]
	result := models.FindComment(param[0])
	response.ResultSuccess(w, result)
}

// GetCommodityComments 获取商品对应评论
func GetCommodityComments(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["commodityId"]
	current, pageSize := GetPageParams(w, r)
	result := models.FindCommodityComments(param[0], current, pageSize)
	response.ResultSuccess(w, result)
}
