package app

import (
	sdkgrpc "github.com/uniqelus/opentrade/sdk/go/grpc"
	sdklog "github.com/uniqelus/opentrade/sdk/go/log/slog"
)

type Config struct {
	Log        sdklog.LogConfig     `yaml:"log"`
	GRPCServer sdkgrpc.ServerConfig `yaml:"grpc_server"`
}
