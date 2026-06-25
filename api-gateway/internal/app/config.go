package app

import (
	sdkhttp "github.com/uniqelus/opentrade/sdk/go/http"
	sdklog "github.com/uniqelus/opentrade/sdk/go/log/slog"
)

type Config struct {
	Log        sdklog.LogConfig     `yaml:"log"`
	HTTPServer sdkhttp.ServerConfig `yaml:"http_server"`
}
