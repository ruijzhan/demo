package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type ControllerHandler func(*Context) error

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	hasTimeout bool
	writerMux  *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

// =========== base =================

// =========== context =================

func (ctx *Context) BaseContext() context.Context {
	return ctx.ctx
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Deadline() (time.Time, bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}

var _ context.Context = &Context{}

// =========== request =================

// =========== response =================
