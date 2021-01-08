package handler

import (
	"encoding/json"
	"net/http"

	"io/ioutil"
	"mango-api/models"
	"mango-api/response"
)

type img struct {
	Img []ImgType `json:"img"`
}

// CreatePromote 创建弹出框
func CreatePromote(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		Name       string `validate:"required" label:"名称"`
		createTime *bool  `validate:"required" label:"创建时间"`
		Img        img    `validate:"required" label:"图片"`
		LinkID     string `label:"推广渠道链接"`
		CouponID   string `validate:"required" label:"关联的优惠券ID"`
	}
	var (
		req models.Promote
		img img
		vf  verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &img)
	req.Img = UploadImg(img.Img)
	result := models.InsertPromote(req)
	response.ResultSuccess(w, result)
}

// GetPromote 获取所有弹窗
func GetPromote(w http.ResponseWriter, r *http.Request) {
	current, pageSize := GetPageParams(w, r)
	result := models.FindPromote(current, pageSize)
	response.ResultSuccess(w, result)
}

// GetPromoteByLinkID 查找对应用户(LinkID)的上架弹窗
func GetPromoteByLinkID(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["linkId"]
	result := models.FindPromoteByLinkID(param[0])
	response.ResultSuccess(w, result)
}

// UpdatePromote 更新弹窗
func UpdatePromote(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Promote
	)
	json.Unmarshal(body, &req)
	result := models.UpdatePromote(req)
	response.ResultSuccess(w, result)
}

// DeletePromote 删除弹窗
func DeletePromote(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Promote
	)
	json.Unmarshal(body, &req)
	result := models.DeletePromote(req.ID)
	response.ResultSuccess(w, result)
}
