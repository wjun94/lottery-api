package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"reflect"
	"regexp"
	"strconv"
)

type Utils struct{}

// GetPlatform 获取设备信息
func (this *Utils) GetPlatform(c *gin.Context) string {
	mobileRe, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP|Windows|Macintosh|Linux|Ubuntu)")
	return mobileRe.FindString(c.GetHeader("User-Agent"))
}

// GetBrowser 获取浏览器信息
func (this *Utils) GetBrowser(c *gin.Context) string {
	mobileRe, _ := regexp.Compile("(?i:msie|firefox|chrome|safari)")
	return mobileRe.FindString(c.GetHeader("User-Agent"))
}

// 获取中文(删除非中文字符)
func (this *Utils) GetCN(str string) string {
	chiReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	return chiReg.ReplaceAllString(str, "")
}

// Empty 判断是否为空
func (this *Utils) Empty(str string) bool {
	return len(str) == 0
}

// HasOwnProperty 判断是否包含属性
func (this *Utils) HasOwnProperty(data interface{}, key string) bool {
	return reflect.ValueOf(data).MethodByName("").IsValid()
}

// strToUin16 string 转 uint16
func (this *Utils) StrToUin16(str string) uint16 {
	num, _ := strconv.Atoi(str)
	return uint16(num)
}

// UUID 设置32位id
func (this *Utils) UUID() string {
	return uuid.New().String()
}
