package framework

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

// 代表请求包含的方法
type IRequest interface {
	// 请求地址 url 中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// 路由匹配中带的参数
	// 形如 /book/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	// form 表单中带的参数
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// =========== request =================

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false

}

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Query(key string) interface{} {
	panic("not implemented") // TODO: Implement
}

// 路由匹配中带的参数
// 形如 /book/:id
func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Param(key string) interface{} {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) FormBool(key string, def bool) (bool, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Form(key string) interface{} {
	panic("not implemented") // TODO: Implement
}

// xml body
func (ctx *Context) BindXml(obj interface{}) error {
	panic("not implemented") // TODO: Implement
}

// 其他格式
func (ctx *Context) GetRawData() ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// 基础信息
func (ctx *Context) Uri() string {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Method() string {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Host() string {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ClientIp() string {
	panic("not implemented") // TODO: Implement
}

// header
func (ctx *Context) Headers() map[string][]string {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Header(key string) (string, bool) {
	panic("not implemented") // TODO: Implement
}

// cookie
func (ctx *Context) Cookies() map[string]string {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) Cookie(key string) (string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return make(map[string][]string)
}
