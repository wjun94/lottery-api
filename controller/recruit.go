package controller

import (
	"encoding/json"
	"job-api/model"
	"job-api/response"
	"job-api/service"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var recruitService = new(service.RecruitService)

// CreateRecruit 创建招聘信息
func CreateRecruit(c *gin.Context) {
	type verify struct {
		Name   string   `json:"name" validate:"required" label:"岗位名称"`
		Type   string   `json:"type" validate:"required" binding:"required" label:""`
		Job    string   `json:"job" validate:"required"  binding:"required" label:"岗位类型"`
		Status byte     `json:"status"  label:"岗位状态"`                 // (1:未发布，2:已发布，3:暂停，4:到期)
		Nat    *bool    `json:"nat" validate:"required" label:"工作性质"` // (全职/兼职)
		Benf   []string `json:"benf" validate:"required" label:"福利待遇"`
		Educ   byte     `json:"educ" validate:"required" label:"学历要求"`
		Exp    byte     `json:"exp" validate:"required" label:"工作经验"`
		Sex    byte     `json:"sex" validate:"required" label:"性别"`
		Rcrt   byte     `json:"rcrt" validate:"required" label:"招聘人数"`
		Prov   string   `json:"prov" validate:"required" label:"省"`
		City   string   `json:"city" validate:"required" label:"市"`
		Area   string   `json:"area" label:"区"`
		Desc   string   `json:"desc" validate:"required" label:"描述"`
		Cont   string   `json:"cont" validate:"required" label:"联系人"`
		Phone  string   `json:"phone" validate:"required" label:"联系电话"`
		Email  string   `json:"email" validate:"required" label:"联系邮箱"`
		Losal  int      `json:"losal" validate:"required" label:"最低薪资"`
		Hisal  int      `json:"hisal" validate:"required" label:"最高薪资"`
	}
	var (
		recruit model.Recruit
		vf      verify
	)
	c.ShouldBindBodyWith(&vf, binding.JSON)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	c.ShouldBindBodyWith(&recruit, binding.JSON)
	recruit.Benf = strings.Join(vf.Benf, ",")
	recruit.ID = Utils.UUID()
	recruit.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	recruit.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	recruit.UserID = info.UserID
	recruit.Status = 2
	_, err := recruitService.Insert(recruit)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultSuccess(c)
}

// ListRecruit 获取列表
func ListRecruit(c *gin.Context) {
	current, pageSize := GetPageParams(c)
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	res, count, err := recruitService.SelectList(current, pageSize, info.UserID)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	var list []model.ListRcrt
	rp, _ := json.Marshal(res)
	json.Unmarshal(rp, &list)
	response.ResultPageJSON(c, list, count, current)
}

// ListRcrtByComp 求职者网站列表
func ListRcrtByComp(c *gin.Context) {
	current, pageSize := GetPageParams(c)
	res, count, err := recruitService.SelectListRcrtByComp(current, pageSize)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultPageJSON(c, res, count, current)
}

// RcrtDetail 求职者网站列表
func RcrtDetail(c *gin.Context) {
	id := c.Query("id")
	res, err := recruitService.SelectDetail(id)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultJSON(c, res)
}

// SelectRecruit 获取列表
func SelectRecruit(c *gin.Context) {
	id := c.Query("id")
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	if id == "" {
		response.ResultSQLError(c, 501, "请传入id")
		return
	}
	result, err := recruitService.Select(id, info.UserID)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	var rt model.SelectRcrt
	rp, _ := json.Marshal(result)
	json.Unmarshal(rp, &rt)
	response.ResultJSON(c, rt)
}

// UpdateRecruit 更新招聘数据
func UpdateRecruit(c *gin.Context) {
	var rcrt model.UpdateRcrt
	if err := c.ShouldBind(&rcrt); err != nil {
		response.ResultSQLError(c, http.StatusBadRequest, err.Error())
		return
	}
	rcrt.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	err := recruitService.Update(rcrt)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
	}
	response.ResultSuccess(c)
}

// UpdateRecruits 批量更新(是否急聘/清空访问量)
func UpdateRecruits(c *gin.Context) {
	var rcrts model.UpdateRcrts
	if err := c.ShouldBind(&rcrts); err != nil {
		response.ResultSQLError(c, http.StatusBadRequest, err.Error())
		return
	}
	target := make(map[string]interface{})
	target["ids"] = rcrts.Ids
	if rcrts.Ugt != nil {
		target["ugt"] = rcrts.Ugt
	} else if rcrts.Views != nil {
		target["views"] = rcrts.Views
	} else {
		response.ResultSQLError(c, http.StatusBadRequest, "缺少参数")
		return
	}
	err := recruitService.Updates(target)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
	}
	response.ResultSuccess(c)
}

// DeleteRecruit 删除职位
func DeleteRecruit(c *gin.Context) {
	var arr struct {
		Ids []string `json:"ids" validate:"required"`
	}
	c.BindJSON(&arr)
	if err := Utils.Trans(arr); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	if len(arr.Ids) == 0 {
		response.ResultSQLError(c, 400, "数组要大于0")
		return
	}
	err := recruitService.Delete(arr.Ids)
	if err != nil {
		response.ResultSQLError(c, 500, err.Message)
		return
	}
	response.ResultSuccess(c)
}
