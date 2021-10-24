package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ip       string
	port     string
	listener net.Listener
}

func New(port string) (*Server, error) {
	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", addr, err)
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}

func (s *Server) ServeHTTP(ctx context.Context, srv *http.Server) error {

	// Spawn a goroutine that listens for context closure. When the context is
	// closed, the server is stopped.
	errChan := make(chan error, 1)
	go func() {
		<-ctx.Done()

		shutDownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		errChan <- srv.Shutdown(shutDownCtx)
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.Serve(s.listener); err != nil && errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	if err := <-errChan; err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}

func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: handler,
	})
}
