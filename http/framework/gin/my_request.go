package gin

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

// 代表请求包含的方法
type IRequest interface {

	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

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
	Headers() map[string]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return map[string][]string(ctx.queryCache)
}

// 请求地址url中带的参数
// 形如: foo.com?a=1&b=bar&c[]=bar
func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], ok
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals, ok
		}
	}
	return def, false
}

// 路由匹配中带的参数
// 形如 /book/:id
func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultParam(key string) interface{} {
	panic("not implemented") // TODO: Implement
}

// form表单中带的参数
func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultFormFile(key string) (*multipart.FileHeader, error) {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) DefaultForm(key string) interface{} {
	panic("not implemented") // TODO: Implement
}

// json body
func (ctx *Context) BindJson(obj interface{}) error {
	panic("not implemented") // TODO: Implement
}

// xml body
func (ctx *Context) BindXml(obj interface{}) error {
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
func (ctx *Context) Headers() map[string]string {
	panic("not implemented") // TODO: Implement
}

// cookie
func (ctx *Context) Cookies() map[string]string {
	panic("not implemented") // TODO: Implement
}
