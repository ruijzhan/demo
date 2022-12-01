package framework

import (
	"bytes"
	"encoding/json"
	"errors"

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
	panic("not implemented") // TODO: Implement
}

// xml 输出
func (ctx *Context) Xml(obj interface{}) IResponse {
	panic("not implemented") // TODO: Implement
}

// html 输出
func (ctx *Context) Html(template string, obj interface{}) IResponse {
	panic("not implemented") // TODO: Implement
}

// 重定向
func (ctx *Context) Redirect(path string) IResponse {
	panic("not implemented") // TODO: Implement
}

// header
func (ctx *Context) SetHeader(key string, val string) IResponse {
	panic("not implemented") // TODO: Implement
}

// Cookie
func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	panic("not implemented") // TODO: Implement
}

// 设置状态码
func (ctx *Context) SetStatus(code int) IResponse {
	panic("not implemented") // TODO: Implement
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
		//TODO: close body here?
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

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	_, err = ctx.responseWriter.Write(byt)
	return err
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
