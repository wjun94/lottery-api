package handler

/**
 * 文档作者: wjun94
 * 创建时间：2019年11月05日
 * 修改时间：2019年11月05日
 * 描述信息：用户优惠券
 */

import (
	"encoding/json"
	"io/ioutil"
	"mango-api/models"
	"mango-api/response"
	_ "mango-api/utils"
	"net/http"
)

// CreateUserCoupon 创建用户优惠券
func CreateUserCoupon(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		CommodityID    string `validate:"required" label:"商品ID"`
		Name           string `validate:"required" label:"优惠券名"`
		Describe       string `validate:"required" label:"描述"`
		Type           byte   `validate:"required" label:"类型"`
		ValidityPeriod int    `validate:"required" label:"有效期"`
	}
	var (
		req models.UserCoupon
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
	result := models.InsertUserCoupon(req)
	response.ResultSuccess(w, result)
}

// GetAllUserCoupon 获取用户对应的所有优惠券
func GetAllUserCoupon(w http.ResponseWriter, r *http.Request) {
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	id := md["UserID"]
	result := models.FindAllUserCoupon(id)
	response.ResultSuccess(w, result)
}

// DeleteUserCoupon 删除用户优惠券
func DeleteUserCoupon(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.UserCoupon
	)
	json.Unmarshal(body, &req)
	result := models.DeleteUserCouponByID(req.ID)
	response.ResultSuccess(w, result)
}
