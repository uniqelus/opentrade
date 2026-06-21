package sdklog

type LogConfig struct {
	Mode  string `yaml:"mode" env:"LOG_MODE" env-default:"development"`
	Level string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
}

func (lc LogConfig) Options() []LogOption {
	return []LogOption{
		WithLogMode(lc.Mode),
		WithLogLevel(lc.Level),
	}
}
