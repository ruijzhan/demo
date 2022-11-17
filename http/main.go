package main

import (
	"context"
	"fmt"
	"html"
	"net"
	"net/http"

	"github.com/ruijzhan/demo/http/framework"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	server := &http.Server{
		Handler: framework.NewCore(),
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
