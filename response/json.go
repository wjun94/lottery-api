package response

import (
	"encoding/json"
	"mango-api/utils"
	"net/http"
)

// Response : JSON Response Object
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total,omitempty"`
}

type OrderResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total,omitempty"`
	Count   float64     `json:"count"`
	Income  float64     `json:"income"`
}

// JSON data
func JSON(d interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    d,
	}
}

// PageJSON 返回页面
func PageJSON(d interface{}, total int) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    d,
		Total:   total,
	}
}

// OrderJSON 订单返回数据
func OrderJSON(d interface{}, total int, count float64, income float64) OrderResponse {
	return OrderResponse{
		Code:    200,
		Message: "success",
		Data:    d,
		Total:   total,
		Count:   count,
		Income:  income,
	}
}

// JSONSuccess ...
func JSONSuccess() Response {
	return Response{
		Code:    200,
		Message: "success",
	}
}

// JSONErrorCode ...
func JSONErrorCode(code int) Response {
	return Response{
		Code:    code,
		Message: ErrorCode[code],
	}
}

// JSONError 数据库报错
func JSONError(message string) Response {
	return Response{
		Code:    500,
		Message: message,
	}
}

// ParseError 解析报错
func ParseError(w http.ResponseWriter, data interface{}) bool {
	if err := utils.Trans(data); err != nil {
		a := JSONError(*err)
		b, _ := json.Marshal(a)
		w.Write(b)
		return false
	}
	return true
}

// JSONErrorMessage 自定义输出
func JSONErrorMessage(code int, message string) Response {
	return Response{
		Code:    500,
		Message: message,
	}
}
