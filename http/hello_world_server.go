package hello_http

import (
	"context"
	"io"
	"net/http"
)

var server http.Server

func StartHttpServer(addr string) {
	handler := http.HandlerFunc(helloHandler)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	server.ListenAndServe()
}

func StopHttpServer() {
	server.Shutdown(context.Background())
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
