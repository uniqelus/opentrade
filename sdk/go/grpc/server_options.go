package sdkgrpc

import "google.golang.org/grpc"

const (
	DefaultServerHost string = "0.0.0.0"
	DefaultServerPort string = "3000"
)

type serverOptions struct {
	host                 string
	port                 string
	grpcServerOptions    []grpc.ServerOption
	serviceRegistrations []ServiceRegistration
}

func defaultServerOptions() *serverOptions {
	return &serverOptions{
		host:                 DefaultServerHost,
		port:                 DefaultServerPort,
		grpcServerOptions:    []grpc.ServerOption{},
		serviceRegistrations: []ServiceRegistration{},
	}
}

type ServerOption func(*serverOptions)

func WithServerHost(host string) ServerOption {
	return func(so *serverOptions) { so.host = host }
}

func WithServerPort(port string) ServerOption {
	return func(so *serverOptions) { so.port = port }
}

func WithGRPCServerOptions(opts ...grpc.ServerOption) ServerOption {
	return func(so *serverOptions) { so.grpcServerOptions = opts }
}

func WithServiceRegistrations(regs ...ServiceRegistration) ServerOption {
	return func(so *serverOptions) { so.serviceRegistrations = regs }
}
