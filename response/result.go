package response

import (
	"job-api/model"
	"job-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var util = new(utils.Utils)

// ResultSuccess 请求成功
func ResultSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, Success())
}

// ResultJSON 请求成功并返回值
func ResultJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, JSON(data))
}

// ResultErrSuccess 请求成功
func ResultErrSuccess(c *gin.Context, err *model.SQLError) {
	if err != nil {
		ResultSQLError(c, err.Number, err.Message)
		return
	}
	c.JSON(http.StatusOK, Success())
}

// ResultErrJSON 处理错误
func ResultErrJSON(c *gin.Context, data interface{}, err *model.SQLError) {
	if err != nil {
		ResultSQLError(c, err.Number, err.Message)
		return
	}
	c.JSON(http.StatusOK, JSON(data))
}

// ResultPageJSON 请求成功并返回值
func ResultPageJSON(c *gin.Context, data interface{}, total uint32, pageNum uint16) {
	c.JSON(http.StatusOK, PageJSON(data, total, pageNum))
}

// ResultSQLError 错误
func ResultSQLError(c *gin.Context, code uint16, message string) {
	err := message
	if ErrorCode[code] != "" {
		err = ErrorCode[code]
	}
	c.JSON(int(code), JSONError(err))
}
