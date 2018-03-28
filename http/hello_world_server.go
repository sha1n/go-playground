package hellohttp

import (
	"context"
	"io"
	"net/http"
)

var server http.Server

/*
StartHTTPServer this is an exported function comment format
*/
func StartHTTPServer(addr string) {
	handler := http.HandlerFunc(helloHandler)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	server.ListenAndServe()
}

/*
StopHTTPServer this is an exported function comment format
*/
func StopHTTPServer() {
	server.Shutdown(context.Background())
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
