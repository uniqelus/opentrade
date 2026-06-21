package sdkhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	address    string
	httpServer *http.Server
}

func NewServer(opts ...ServerOption) *Server {
	options := defaultServerOptions()
	for _, opt := range opts {
		opt(options)
	}

	address := net.JoinHostPort(options.host, options.port)
	return &Server{
		address: address,
		httpServer: &http.Server{
			Handler:      options.handler,
			ReadTimeout:  options.readTimeout,
			WriteTimeout: options.writeTimeout,
			IdleTimeout:  options.idleTimeout,
		},
	}
}

func (s *Server) Run(ctx context.Context) error {
	lisCfg := &net.ListenConfig{}
	lis, err := lisCfg.Listen(ctx, "tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen expected address %s: %w", s.address, err)
	}

	errCh := make(chan error)
	go func() {
		select {
		case errCh <- s.httpServer.Serve(lis):
		case <-ctx.Done():
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("failed to serve expected address %s: %w", s.address, err)
	case <-ctx.Done():
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
