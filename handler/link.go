package handler

import (
	"encoding/json"
	"io/ioutil"
	"mango-api/response"
	"net/http"

	"mango-api/models"
)

// CreateLink 添加h5链接
func CreateLink(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		Name      string `validate:"required" label:"名称"`
		Platform  byte   `validate:"required" label:"平台"`
		AccountID string `validate:"required" label:"账户ID"`
	}
	var (
		req models.Link
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	if req.Platform == 2 {
		// 微信小程序生成平台连接
		str := ByteQiniu(models.GetQRCode(req.ID), req.Name)
		req.QrCode = str
	}
	result := models.InsertLink(req)
	response.ResultSuccess(w, result)
}

// GetAllLink 获取所有链接
func GetAllLink(w http.ResponseWriter, r *http.Request) {
	current, pageSize := GetPageParams(w, r)
	result := models.FindAllLink(current, pageSize)
	response.ResultSuccess(w, result)
}

// SelectLink 获取对应链接
func SelectLink(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["id"]
	result := models.SelectLink(param[0])
	response.ResultSuccess(w, result)
}

// DeleteLink 删除链接
func DeleteLink(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Link
	)
	json.Unmarshal(body, &req)
	result := models.DeleteLink(req.ID)
	response.ResultSuccess(w, result)
}

// UpdateLink 更新链接
func UpdateLink(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		AccountID string `validate:"required" label:"账户ID"`
	}
	var (
		req models.Link
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	result := models.UpdateLink(req)
	response.ResultSuccess(w, result)
}
