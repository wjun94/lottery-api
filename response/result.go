package response

import (
	"encoding/json"
	"net/http"
)

// ResultSuccess 请求成功
func ResultSuccess(w http.ResponseWriter, result interface{}) {
	response, _ := json.Marshal(result)
	w.Write(response)
}
