package server

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/aaorlov/stream/log"
)

// Server wraps http server
type Server struct {
	http *http.Server
	r    *mux.Router
}

// Srv interface
type Srv interface {
	Start()
	Shutdown(ctx context.Context) error
}

// New create ws server
func New(addr ...string) Srv {

	address := resolveAddress(addr)
	r := NewRouter()

	httpSrv := &http.Server{
		Addr:    address,
		Handler: r,
	}

	srv := &Server{
		httpSrv,
		r,
	}

	return srv
}

// Start server
func (srv *Server) Start() {
	go func() {
		if err := srv.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen: %s\n", err)
		}
	}()

	log.Infof("Listening and serving HTTP on %s\n", srv.http.Addr)
}

// Shutdown server
func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.http.Shutdown(ctx); err != nil {
		return err
	}
	log.Infof("Server exited properly")

	return nil
}
