package httpserver

import "net/http"

type HTTPServer struct {
	handler http.Handler
	port    string
}

func NewHTTPServer(handler http.Handler, port string) *HTTPServer {
	return &HTTPServer{
		handler: handler,
		port:    port,
	}
}

func (h *HTTPServer) StartServer() error {
	return http.ListenAndServe(":"+h.port, h.handler)
}
