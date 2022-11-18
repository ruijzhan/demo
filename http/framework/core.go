package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(req, resp)

	router := c.router["foo"]
	if router == nil {
		return
	}

	log.Println("core.router")
	router(ctx)
}
