package sdkhttp

import (
	"net/http"
	"time"
)

var (
	DefaultServerHost   string        = "0.0.0.0"
	DefaultServerPort   string        = "2000"
	DefaultReadTimeout  time.Duration = 5 * time.Second
	DefaultWriteTimeout time.Duration = 5 * time.Second
	DefaultIdleTimeout  time.Duration = 10 * time.Second
)

type serverOptions struct {
	host         string
	port         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	handler      http.Handler
}

func defaultServerOptions() *serverOptions {
	return &serverOptions{
		host:         DefaultServerHost,
		port:         DefaultServerPort,
		readTimeout:  DefaultReadTimeout,
		writeTimeout: DefaultWriteTimeout,
		idleTimeout:  DefaultIdleTimeout,
		handler:      http.DefaultServeMux,
	}
}

type ServerOption func(*serverOptions)

func WithServerHost(host string) ServerOption {
	return func(so *serverOptions) { so.host = host }
}

func WithServerPort(port string) ServerOption {
	return func(so *serverOptions) { so.port = port }
}

func WithServerReadTimeout(timeout time.Duration) ServerOption {
	return func(so *serverOptions) { so.readTimeout = timeout }
}

func WithServerWriteTimeout(timeout time.Duration) ServerOption {
	return func(so *serverOptions) { so.writeTimeout = timeout }
}

func WithServerIdleTimeout(timeout time.Duration) ServerOption {
	return func(so *serverOptions) { so.idleTimeout = timeout }
}

func WithServerHandler(handler http.Handler) ServerOption {
	return func(so *serverOptions) { so.handler = handler }
}
