package main

import (
	"context"
	"fmt"
	"html"
	"net"
	"net/http"

	"github.com/ruijzhan/demo/http/framework"
	"github.com/ruijzhan/demo/http/framework/middleware"
	// _ "net/http/pprof"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	core := framework.NewCore()

	core.Use(middleware.Recovery())

	// go http.ListenAndServe("localhost:8082", nil)
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
		BaseContext: func(l net.Listener) context.Context {
			return context.Background()
		},
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			return context.Background()
		},
	}

	server.ListenAndServe()
}
