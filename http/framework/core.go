package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
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

func (c *Core) addMethod(m, url string, h ControllerHandler) {
	if err := c.router[m].AddRouter(url, h); err != nil {
		log.Fatalf("add router error %s %v", m, err)
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.addMethod("GET", url, handler)
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.addMethod("POST", url, handler)
}

func (c *Core) Put(url string, handler ControllerHandler) {
	c.addMethod("PUT", url, handler)
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	c.addMethod("DELETE", url, handler)
}

func (c *Core) FindRouteByRequest(r *http.Request) ControllerHandler {

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

	handler := c.FindRouteByRequest(req)
	if handler == nil {
		ctx.Json(404, "not found")
		return
	}

	if err := handler(ctx); err != nil {
		ctx.Json(500, "internal error")
		return
	}

}
