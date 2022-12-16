package gin

import (
	"fmt"
	"net/http"

	"github.com/ruijzhan/demo/http/framework/gin/internal/json"
)

// IResponse代表返回方法
type IResponse interface {
	// Json输出
	IJson(obj interface{}) IResponse

	// Jsonp输出
	IJsonp(obj interface{}) IResponse

	//xml输出
	IXml(obj interface{}) IResponse

	// html输出
	IHtml(template string, obj interface{}) IResponse

	// string
	IText(format string, values ...interface{}) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// header
	ISetHeader(key string, val string) IResponse

	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置200状态
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IJsonp(obj interface{}) IResponse {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) IXml(obj interface{}) IResponse {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) IHtml(template string, obj interface{}) IResponse {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Content-Type", "application/text")
	ctx.Writer.Write([]byte(out))
	return ctx
}

func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetHeader(key string, val string) IResponse {
	ctx.Writer.Header().Add(key, val)
	return ctx
}

func (ctx *Context) ISetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	panic("not implemented") // TODO: Implement
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}
