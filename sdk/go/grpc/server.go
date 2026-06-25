package sdkgrpc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type ServiceRegistration func(grpc.ServiceRegistrar)

type Server struct {
	address    string
	grpcServer *grpc.Server
}

func NewServer(opts ...ServerOption) *Server {
	options := defaultServerOptions()
	for _, opt := range opts {
		opt(options)
	}

	address := net.JoinHostPort(options.host, options.port)

	grpcServer := grpc.NewServer(options.grpcServerOptions...)
	for _, serviceReg := range options.serviceRegistrations {
		serviceReg(grpcServer)
	}

	return &Server{
		address:    address,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start(ctx context.Context) error {
	lisCfg := &net.ListenConfig{}
	lis, err := lisCfg.Listen(ctx, "tcp", s.address)
	if err != nil {
		return fmt.Errorf("cannot listen expected address %s: %w", s.address, err)
	}

	errCh := make(chan error)
	go func() {
		select {
		case errCh <- s.grpcServer.Serve(lis):
		case <-ctx.Done():
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if errors.Is(err, grpc.ErrServerStopped) {
			return nil
		}

		return fmt.Errorf("cannot serve expected address %s: %w", s.address, err)
	case <-ctx.Done():
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		return nil

	case <-ctx.Done():
		s.grpcServer.Stop()
		return fmt.Errorf("failed to stop server: %w", ctx.Err())
	}
}
