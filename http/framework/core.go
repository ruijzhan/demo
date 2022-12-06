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

func (c *Core) FindRouteNodeByRequest(req *http.Request) *node {
	uri := req.URL.Path
	method := strings.ToUpper(req.Method)

	if h, ok := c.router[method]; ok {
		return h.root.matchNode(uri)
	}

	return nil
}

func (c *Core) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := NewContext(req, resp)

	node := c.FindRouteNodeByRequest(req)
	if node == nil {
		ctx.SetStatus(404).Json("not found")
		return
	}

	ctx.SetHandlers(node.handlers)

	parems := node.parseParamsFromEndNode(req.URL.Path)
	ctx.SetParams(parems)

	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("internal error")
		return
	}

}
