package controller

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
	"strings"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("test"))

// HomeHandler is HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(fmt.Sprintf("%s", body))

	fmt.Fprintln(w, "Welcome!")
}

// GetHandler1 get请求
func GetHandler1(w http.ResponseWriter, r *http.Request) {
	println(r.Header["User-Agent"][0])
	println(strings.Split(r.RemoteAddr, ":")[0])

	session, _ := store.Get(r, "123555")
	// Set some session values.
	session.Values["name"] = r.Header["User-Agent"][0]
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	// 获取设备信息，判断是否是从不同设备登入
	var res = map[string]string{"result": "succ", "name": r.Header["User-Agent"][0]}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetHandler2 测试
func GetHandler2(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["name"]

	session, _ := store.Get(r, "123555")
	fmt.Println(session.Values["name"] == r.Header["User-Agent"][0])

	var res = map[string]string{"result": "succ", "name": param[0]}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
