package handler
/**
 * 文档作者: wjun94
 * 创建时间：2019年11月02日
 * 修改时间：2019年11月02日
 * 描述信息：工具类
 */
import (
	"strconv"
	"net/http"
	// "encoding/json"
	// "github.com/dgrijalva/jwt-go"
)

// GetAllCoupon 获取所有地址
func GetPageParams(w http.ResponseWriter, r *http.Request)(int, int) {
	vals := r.URL.Query()
	var current int = 1	// 当前页
	var pageSize = 10	// 一页显示几个 
	if _, ok := vals["current"]; ok {
		num, _ := strconv.Atoi(vals["current"][0])
		current = num
	}
	if _, ok := vals["pageSize"]; ok {
		num, _ := strconv.Atoi(vals["pageSize"][0])
		pageSize = num
	}
	return current, pageSize
}