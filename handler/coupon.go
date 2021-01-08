package handler

import (
	"encoding/json"
	"io/ioutil"
	"mango-api/response"
	"net/http"

	"mango-api/models"
	"sync"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月18日
 * 修改时间：2019年10月18日
 * 描述信息：地址
 */

// AddCoupon 添加地址接口
func AddCoupon(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		CommodityID    string `validate:"required" label:"商品ID"` // 商品id
		Name           string `validate:"required" label:"优惠券名"` // 优惠券名
		Describe       string `validate:"required" label:"描述"`   // 描述
		Unit           string `validate:"required" label:"单位"`   // 单位
		Type           byte   `validate:"required" label:"类型"`   // 类型
		ValidityPeriod int    `validate:"required" label:"有效期"`  // 有效期
	}
	var (
		req models.Coupon
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		result := models.InsertCoupon(req)
		response.ResultSuccess(w, result)
	}()
	wg.Wait()
}

// GetAllCoupon 获取所有优惠券
func GetAllCoupon(w http.ResponseWriter, r *http.Request) {
	current, pageSize := GetPageParams(w, r)
	result := models.FindAllCoupon(current, pageSize)
	response.ResultSuccess(w, result)
}

// GetCouponByID 查找优惠券
func GetCouponByID(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["id"]
	result := models.FindCouponByID(param[0])
	response.ResultSuccess(w, result)
}

// DeleteCoupon 删除优惠券
func DeleteCoupon(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Coupon
	)
	json.Unmarshal(body, &req)
	result := models.DeleteCoupon(req.ID)
	response.ResultSuccess(w, result)
}

// UpdateCoupon 更新优惠券
func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Coupon
	)
	json.Unmarshal(body, &req)
	result := models.UpdateCoupon(req)
	response.ResultSuccess(w, result)
}
