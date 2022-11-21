package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {

	return &Core{
		router: map[string]map[string]ControllerHandler{
			"GET":    {},
			"POST":   {},
			"PUT":    {},
			"DELETE": {},
		},
	}
}

func (c *Core) addMethod(m, url string, h ControllerHandler) {
	up := strings.ToUpper(url)
	c.router[m][up] = h
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
	url := strings.ToUpper(r.URL.Path)

	if mm, ok := c.router[m]; ok {
		if h, ok := mm[url]; ok {
			return h
		}
	}
	return nil
}

func (c *Core) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(req, resp)

	handler := c.FindRouteByRequest(req)
	if handler == nil {
		ctx.Json(404, "not found")
	}

	if err := handler(ctx); err != nil {
		ctx.Json(500, "internal error")
		return
	}

}
