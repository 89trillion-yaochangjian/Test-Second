package StructInfo

var (
	OK  = response(200, "OK")    // 通用成功
	Err = response(500, "ERROR") // 通用错误

	EmptyErr   = response(1001, "字符串不能为空")
	LegErr     = response(1002, "字符串不合法")
	DivisorErr = response(1003, "除数不能为零")
	StrconvErr = response(1004, "数据类型转换异常")
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息

func (res *Response) WithMsg(message string, err error) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
