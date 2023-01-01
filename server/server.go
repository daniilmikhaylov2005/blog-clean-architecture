package server

import (
	"net/http"
	"time"
  "log"
	"golang.org/x/net/context"
)

type Server struct {
  httpServer *http.Server
}

func (s *Server) Run (handler http.Handler, config map[string]string) error {
  s.httpServer = &http.Server{
    Addr: ":" + config["http_port"],
    Handler: handler,
    MaxHeaderBytes: 1 << 20,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }
  log.Printf("Starting http server on port %s\n", config["http_port"])
  return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
  log.Println("server is shutting down...")
  return s.httpServer.Shutdown(ctx)
}
