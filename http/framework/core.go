package framework

import (
	"net/http"
)

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//TODO
	req.Context()

	// fs := http.FileServer(http.Dir("/home/bob/static"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))
}
