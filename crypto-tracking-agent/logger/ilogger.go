package logger

import "errors"

type ILogger interface {
	Debug(a ...any)
	Info(a ...any)
	Warning(a ...any)
	Error(a ...any)
}

var Logger ILogger

type LoggerConfig struct {
	Type       string
	FileName   string
	MaxSize    int // MB
	MaxBackups int
	MaxAge     int
	Level      int // Debug/Info/Warning/Error = -1/0/1/2
}

const (
	ZAP_LOGGER string = "zap"
)

func NewLogger(config LoggerConfig) error {
	switch config.Type {
	case ZAP_LOGGER:
		logger, err := NewZapLogger(config)
		if err != nil {
			return err
		}
		Logger = logger
		return nil
	default:
		return errors.New("Invalid logger type")
	}
}
