package app

import (
	"context"
	"net/http"
	"project/pkg/logging"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) StartServer(host, port string, logger *logging.Logger) {

	s.httpServer = &http.Server{
		Addr:         host + ":" + port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Infof("server listening oh %s:%s", host, port)
	logger.Fatal(s.httpServer.ListenAndServe())

}

func (s *Server) StopServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
