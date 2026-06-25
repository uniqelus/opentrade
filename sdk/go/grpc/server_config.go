package sdkgrpc

type ServerConfig struct {
	Host string `yaml:"host" env:"GRPC_SERVER_HOST" env-default:"0.0.0.0"`
	Port string `yaml:"port" env:"GRPC_SERVER_PORT" env-default:"3000"`
}

func (sc ServerConfig) Options() []ServerOption {
	return []ServerOption{
		WithServerHost(sc.Host),
		WithServerPort(sc.Port),
	}
}
