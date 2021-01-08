package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"mango-api/models"
	"mango-api/response"
)

/**
 * 描述信息：账户
 */

// GetAuth 获取用户等级
func GetAuth(w http.ResponseWriter, r *http.Request) {
	info := r.Context().Value("info")
	md, _ := info.(map[string]interface{})
	tp := md["Type"]
	result := response.JSON(map[string]interface{}{"type": tp})
	response, _ := json.Marshal(result)
	w.Write(response)
}

// AddAccount 添加账户接口
func AddAccount(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		Name     string `validate:"required" label:"账号"`
		Password string `validate:"required,min=6,max=11" label:"密码"`
		Phone    string `validate:"required,numeric,min=6,max=11" label:"手机号"`
		Type     byte   `validate:"required,lte=3" label:"类型"`
	}
	var (
		req models.Account
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	result := models.InsertAccount(req)
	response, _ := json.Marshal(result)
	w.Write(response)
}

// GetAllAccount 获取所有账户信息
func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	current, pageSize := GetPageParams(w, r)
	info := r.Context().Value("info")
	md, _ := info.(map[string]interface{})
	id := md["UserID"].(string)
	result := models.FindAllAccount(current, pageSize, id)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	info := r.Context().Value("info")
	md, _ := info.(map[string]interface{})
	id := md["UserID"].(string)
	result := models.GetAccountByID(id)
	response, _ := json.Marshal(result)
	w.Write(response)
}

// UpdateAccount 更新优惠券
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		ID       string `validate:"required,len=36" label:"id"`
		Password string `validate:"min=6,max=11" label:"密码"`
		Phone    string `validate:"numeric,min=6,max=11" label:"手机号"`
		Type     byte   `validate:"lte=3" label:"类型"`
	}
	var (
		req models.Account
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	result := models.UpdateAccount(req)
	response.ResultSuccess(w, result)
}

// DeleteAccount 删除优惠券
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Account
	)
	json.Unmarshal(body, &req)
	result := models.DeleteAccount(req.ID)
	response.ResultSuccess(w, result)
}
