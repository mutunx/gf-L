package library

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 数据返回通用结构体
type JsonRes struct {
	Data     interface{} `json:"data"`    //返回数据,业务接口定义具体数据结构
	Message  string      `json:"message"` // 提示信息
	Code     int         `json:"code"`    // 错误码 (0:成功 1:失败 >1:错误码)
	Redirect string      `json:"redirect"`
}

// 返回标准json数据
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	_ = r.Response.WriteJson(&JsonRes{
		Data:    responseData,
		Message: message,
		Code:    code,
	})
}

// 返回标准json数据并退出当前http执行函数
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

// 返回标准json数据并引导客户端跳转
func JsonRedirect(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	_ = r.Response.WriteJson(JsonRes{
		Data:     responseData,
		Message:  message,
		Code:     code,
		Redirect: redirect,
	})
}

// 返回标准json数据引导跳转,并退出当前HTTP执行的函数
func JsonRedirectExit(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	JsonRedirect(r, code, message, redirect, data...)
	r.Exit()
}
