package sdkhttp

import "time"

type ServerConfig struct {
	Host         string        `yaml:"host" env:"HTTP_SERVER_HOST" env-default:"0.0.0.0"`
	Port         string        `yaml:"port" env:"HTTP_SERVER_PORT" env-default:"2000"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"HTTP_SERVER_READ_TIMEOUT" env-default:"5s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"HTTP_SERVER_WRITE_TIMEOUT" env-default:"5s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"10s"`
}

func (sc ServerConfig) Options() []ServerOption {
	return []ServerOption{
		WithServerHost(sc.Host),
		WithServerPort(sc.Port),
		WithServerReadTimeout(sc.ReadTimeout),
		WithServerWriteTimeout(sc.WriteTimeout),
		WithServerIdleTimeout(sc.IdleTimeout),
	}
}
