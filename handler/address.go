package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"mango-api/models"
	"mango-api/response"
	"mango-api/utils"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月18日
 * 修改时间：2019年10月18日
 * 描述信息：地址
 */

// AddAddress 添加地址接口
func AddAddress(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req       models.Address
		verify    models.VerifyAddress
		isDefault struct {
			IsDefault bool `json:"isDefault"`
		}
	)
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &isDefault)
	json.Unmarshal(body, &verify)
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	id := md["UserID"]
	req.UserID = id
	verify.UserID = id
	if err := utils.Trans(verify); err != nil {
		a := response.JSONError(*err)
		b, _ := json.Marshal(a)
		w.Write(b)
		return
	}
	result := models.InsertAddress(req, isDefault.IsDefault)
	response.ResultSuccess(w, result)
}

// GetAllAddress 获取所有地址
func GetAllAddress(w http.ResponseWriter, r *http.Request) {
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	userID := md["UserID"]
	result := models.FindAllAddress(userID)
	response.ResultSuccess(w, result)
}

// UpdateAddress 更新地址
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req       models.Address
		isDefault struct {
			IsDefault bool `json:"isDefault"`
		}
	)
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &isDefault)
	result := models.UpdateAddress(req, isDefault.IsDefault)
	response.ResultSuccess(w, result)
}

// DeleteAddress 删除地址
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Address
	)
	json.Unmarshal(body, &req)
	result := models.DeleteAddress(req.ID)
	response.ResultSuccess(w, result)
}
