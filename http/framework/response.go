package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"io"
)

// IResponse 代表返回方法
type IResponse interface {
	// Json 输出
	Json(obj interface{}) IResponse

	// Jsonp 输出
	Jsonp(obj interface{}) IResponse

	//xml 输出
	Xml(obj interface{}) IResponse

	// html 输出
	Html(template string, obj interface{}) IResponse

	// string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置 200 状态
	SetOkStatus() IResponse
}

// Jsonp 输出
func (ctx *Context) Jsonp(obj interface{}) IResponse {
	// 获取请求参数 callback
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成 XSS 攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}
	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}
	// 输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

// xml 输出
func (ctx *Context) Xml(obj interface{}) IResponse {
	panic("not implemented") // TODO: Implement
}

// html 输出
func (ctx *Context) Html(file string, obj interface{}) IResponse {
	// 读取模版文件，创建 template 实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	// 执行 Execute 方法将 obj 和模版进行结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

// 重定向
func (ctx *Context) Redirect(path string) IResponse {
	panic("not implemented") // TODO: Implement
}

// header
func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

// Cookie
func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	panic("not implemented") // TODO: Implement
}

// 设置状态码
func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

// 设置 200 状态
func (ctx *Context) SetOkStatus() IResponse {
	panic("not implemented") // TODO: Implement
}

// =========== response =================

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := io.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}

		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}
