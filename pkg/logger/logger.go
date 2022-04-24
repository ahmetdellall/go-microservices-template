package logger

type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode  string `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func NewLoggerConfig(loglevel string, devMode string, encoder string) *Config {
	return &Config{LogLevel: loglevel, DevMode: devMode, Encoder: encoder}
}

type Logger interface {
	Errorf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Infof(template string, args ...interface{})
}
