package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {

	return &Core{
		router: map[string]*Tree{
			"GET":    NewTree(),
			"POST":   NewTree(),
			"PUT":    NewTree(),
			"DELETE": NewTree(),
		},
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) addMethod(m, url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router[m].AddRouter(url, all); err != nil {
		log.Fatalf("add router error %s %v", m, err)
	}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	c.addMethod("GET", url, handlers...)
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	c.addMethod("POST", url, handlers...)
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	c.addMethod("PUT", url, handlers...)
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	c.addMethod("DELETE", url, handlers...)
}

func (c *Core) FindRouteByRequest(r *http.Request) []ControllerHandler {

	m := strings.ToUpper(r.Method)
	uri := strings.ToUpper(r.URL.Path)

	if mh, ok := c.router[m]; ok {
		return mh.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(req, resp)

	handlers := c.FindRouteByRequest(req)
	if handlers == nil {
		ctx.Json(404, "not found")
		return
	}

	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(500, "internal error")
		return
	}

}
