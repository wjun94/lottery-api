package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mango-api/response"
	"net/http"
	"strings"

	"mango-api/models"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月25日
 * 修改时间：2019年10月25日
 * 描述信息：首页商品
 */

type ImgType struct {
	BaseURL string `json:"baseUrl"`
	URL     string `json:"url"`
	UID     string `json:"uid"`
	Type    string `json:"type"`
}

type rImg struct {
	AdImg      []ImgType `json:"adImg"`
	Img        []ImgType `json:"img"`
	Imgs       []ImgType `json:"imgs"`
	DetailImgs []ImgType `json:"detailImgs"`
}

type Ad struct {
	IsAd bool `json:"isAd"`
}

// UploadImg 图片上传到七牛云
func UploadImg(data []ImgType) string {
	imgData := make([]string, 0, 50)
	for _, v := range data {
		if strings.Contains(v.BaseURL, "base64") == true {
			result := Base64Qiniu(v.BaseURL, fmt.Sprintf("%s.%s", v.UID, strings.Split(v.Type, "/")[1]))
			imgData = append(imgData, result)
		} else {
			imgData = append(imgData, v.URL)
		}
	}
	return strings.Join(imgData, ",")
}

// AddCommodity 添加地址接口
func AddCommodity(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	type verify struct {
		AdImg      []ImgType `validate:"required" label:"商品广告图"`
		Img        []ImgType `validate:"required" label:"商品主图"`
		Imgs       []ImgType `validate:"required" label:"商品轮播图"`
		DetailImgs []ImgType `validate:"required" label:"商品详情图"`
		Name       string    `validate:"required" label:"商品名"`
		Describe   string    `validate:"required" label:"描述"`
		Price      float64   `validate:"required,ltfield=OldPrice" label:"价格"`
		OldPrice   float64   `validate:"required" label:"原价"`
		Unit       string    `validate:"required" label:"单位"`
	}
	var (
		req      models.Commodity
		data     rImg
		isAdData Ad
		vf       verify
	)
	json.Unmarshal(body, &vf)
	if !response.ParseError(w, vf) {
		return
	}
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &data)
	json.Unmarshal(body, &isAdData)
	req.AdImg = UploadImg(data.AdImg)
	req.Img = UploadImg(data.Img)
	req.Imgs = UploadImg(data.Imgs)
	req.DetailImgs = UploadImg(data.DetailImgs)
	result := models.InsertCommodity(req, isAdData.IsAd)
	response.ResultSuccess(w, result)
}

// GetAllCommodity 获取所有商品
func GetAllCommodity(w http.ResponseWriter, r *http.Request) {
	result := models.FindAllCommodity()
	response.ResultSuccess(w, result)
}

// GetRelatedCommodity 商品和优惠券关联查询
func GetRelatedCommodity(w http.ResponseWriter, r *http.Request) {
	current, pageSize := GetPageParams(w, r)
	result := models.FindRelatedCommodity(current, pageSize)
	response.ResultSuccess(w, result)
}

// GetCommodityUserCouponByID 查找用户商品对应的优惠券
func GetCommodityUserCouponByID(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["id"]
	info := r.Context().Value("info")
	md, _ := info.(map[string]string)
	userID := md["UserID"]
	result := models.FindCommodityUserCouponByID(param[0], userID)
	response.ResultSuccess(w, result)
}

// UpdateCommodity 更新商品
func UpdateCommodity(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req      models.Commodity
		data     rImg
		isAdData Ad
	)
	json.Unmarshal(body, &req)
	json.Unmarshal(body, &data)
	json.Unmarshal(body, &isAdData)
	req.AdImg = UploadImg(data.AdImg)
	req.Img = UploadImg(data.Img)
	req.Imgs = UploadImg(data.Imgs)
	req.DetailImgs = UploadImg(data.DetailImgs)
	result := models.UpdateCommodity(req, isAdData.IsAd)
	response.ResultSuccess(w, result)
}

// removeImg 批量删除七牛云图片
func removeImg(data []string) {
	for _, v := range data {
		DeleteQiniu(strings.Split(v, "?")[0])
	}
}

// DeleteCommodity 删除商品
func DeleteCommodity(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var (
		req models.Commodity
	)
	json.Unmarshal(body, &req)
	// find := models.FindCommodity(req.ID)
	// removeImg(find.AdImg)
	// removeImg(find.Img)
	// removeImg(find.Imgs)
	// removeImg(find.DetailImgs)
	result := models.DeleteCommodity(req.ID)
	response.ResultSuccess(w, result)
}

// GetCommodityAdByID 获取广告位图片
func GetCommodityAdByID(w http.ResponseWriter, r *http.Request) {
	result := models.FindCommodityAd()
	response.ResultSuccess(w, result)
}

// GetCommodityList 小程序首页使用(获取首页列表数据)
func GetCommodityList(w http.ResponseWriter, r *http.Request) {
	result := models.FindCommodityList()
	response.ResultSuccess(w, result)
}
