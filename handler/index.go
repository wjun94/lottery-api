package handler

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月12日
 * 修改时间：2019年10月12日
 * 描述信息：案例
 */
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HomeHandler is HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(fmt.Sprintf("%s", body))
	fmt.Fprintln(w, "Welcome!")
}

// GetHandler get请求
func GetHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["name"]
	var res = map[string]string{"result": "succ", "name": param[0]}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
