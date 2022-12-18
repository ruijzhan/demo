package gin

import (
	"context"

	"github.com/ruijzhan/demo/http/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

func (engin *Engine) Bind(provider framework.ServiceProvider) error {
	return engin.container.Bind(provider)
}

func (engin *Engine) IsBind(key string) bool {
	return engin.container.IsBind(key)
}

func (ctx *Context) Make(key string) (any, error) {
	return ctx.container.Make(key)
}

func (ctx *Context) MustMake(key string) any {
	return ctx.container.MustMake(key)
}

func (ctx *Context) MakeNew(key string, params []any) (any, error) {
	return ctx.container.MakeNew(key, params)
}
