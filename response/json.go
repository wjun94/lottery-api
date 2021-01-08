package response

// Response : JSON Response Object
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   uint32      `json:"total,omitempty"`
	PageNum uint16      `json:"pageNum,omitempty"`
}

type ResponseSuccess struct {
	Message string `json:"message"`
}

// JSON data
func Success() ResponseSuccess {
	return ResponseSuccess{
		Message: "success",
	}
}

// JSON data
func JSON(d interface{}) Response {
	return Response{
		Message: "success",
		Data:    d,
	}
}

func PageJSON(d interface{}, total uint32, PageNum uint16) Response {
	return Response{
		Message: "success",
		Data:    d,
		Total:   total,
		PageNum: PageNum,
	}
}

// JSONError 自定义输出
func JSONError(message string) Response {
	return Response{
		Message: message,
	}
}
