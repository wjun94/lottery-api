package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"mango-api/models"
	"mango-api/response"
	_ "mango-api/utils"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月18日
 * 修改时间：2019年10月18日
 * 描述信息：地址
 */

// CreateOrder 添加订单接口
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		Method     string            `validate:"required" label:"支付方式"`
		Price      float64           `label:"商品金额"` // 计算后商品总价
		TotalPrice float64           `label:"总价"`
		Progress   byte              `label:"邮费"`
		UserCoupon models.UserCoupon `label:"用户优惠券"`
		Count      int               `label:"商品总个数"`
		Val        float64           `label:"商品总个数"`
	}
	var (
		req    models.Order
		result response.Response
		vf     verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	id := md["UserID"]
	req.UserID = id
	req.Status = 1

	var val float64 // 关联商品价格(包邮套件)
	for k, v := range req.Commodity {
		// 关联商品价格计算
		if k >= 1 {
			val += v.Price * float64(v.Count)
		}
	}
	calc := Calc(vf.UserCoupon, req.Commodity[0].Options, req.Commodity[0].Count, req.Commodity[0].Price, val)
	totalF1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", calc), 64)
	totalF2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", vf.TotalPrice), 64)
	if totalF1 == totalF2 {
		result = models.InsertOrder(req)
	} else {
		result = response.JSONError("金额不匹配")
	}
	response.ResultSuccess(w, result)
}

// GetOrder 获取订单
func GetOrder(w http.ResponseWriter, r *http.Request) {
	current, pageSize := GetPageParams(w, r)
	vals := r.URL.Query()
	orderID, _ := vals["orderId"]
	status, _ := vals["status"]
	start, _ := vals["start"]
	end, _ := vals["end"]
	vLinks, _ := vals["links"]
	str := start[0]
	ed := end[0]
	var statusArr []string
	switch status[0] {
	case "0":
		statusArr = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	case "1":
		statusArr = []string{"3", "4", "5", "6", "7", "8"}
	case "2":
		statusArr = []string{"1", "2", "9"}
	}
	var links []string
	if vLinks[0] != "all" {
		// 获取所有订单
		links = strings.Split(vLinks[0], ",")
	}
	result := models.GetAllOrder(current, pageSize, orderID[0], statusArr, str, ed, links)
	response.ResultSuccess(w, result)
}

// GetWeappOrderList 获取微信订单
func GetWeappOrderList(w http.ResponseWriter, r *http.Request) {
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	id := md["UserID"]
	current, pageSize := GetPageParams(w, r)
	result := models.WeappOrderList(current, pageSize, id)
	response.ResultSuccess(w, result)
}

// UpdateOrder 更新订单
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type t struct {
		Openid        string `json:"openid"`
		TransactionID string `json:"transactionId"`
	}
	var (
		req    models.Order
		target t
	)
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &target)
	models.UpdateOrder(w, req, target.Openid, target.TransactionID)
	// response.ResultSuccess(w, result)
}

// UpdateOrderAddress 更新订单地址
func UpdateOrderAddress(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.OrderAddress
	)
	json.Unmarshal(body, &req)
	result := models.UpdateOrderAddress(req)
	response.ResultSuccess(w, result)
}

// GetOrderDetails 获取订单详情
func GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["id"]
	result := models.OrderDetails(param[0])
	response.ResultSuccess(w, result)
}

// Calc 计算价格方法
// order:订单信息  userCoupons:优惠券  options:附加选项 count:商品个数  price:商品价格  val: 关联商品价格(包邮套件)
func Calc(userCoupons models.UserCoupon, options byte, count int, price float64, val float64) float64 {
	var progress float64 = 8 // 运费
	var other float64
	if options == 1 {
		// 封塑
		other = float64(count) * 0.5
	}
	var rstValue float64
	switch userCoupons.Type {
	case 1:
		// 无门槛券
		var tol int = 0
		if count-int(userCoupons.Value) > 0 {
			tol = count - int(userCoupons.Value) // 减去免费个数
		}
		cPrice := float64(tol)*price + other
		rstValue = cPrice + val
	case 2:
		// 满减
		commodity := float64(count)*price + other
		var reduceCommodity = commodity
		if commodity >= userCoupons.Full {
			reduceCommodity = commodity - userCoupons.Reduce
		}
		jump := reduceCommodity
		if reduceCommodity < 0 {
			jump = 0
		}
		rstValue = jump + val + progress
	default:
		rstValue = float64(count)*price + val + other
	}
	if rstValue < 20 {
		rstValue += progress
	}
	if rstValue < 0 {
		rstValue = 0
	}
	return rstValue
}
