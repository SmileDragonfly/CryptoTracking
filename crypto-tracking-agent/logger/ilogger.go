package logger

import "errors"

type ILogger interface {
	Debug(a ...any)
	Info(a ...any)
	Warning(a ...any)
	Error(a ...any)
	Debugf(template string, a ...any)
	Infof(template string, a ...any)
	Warningf(template string, a ...any)
	Errorf(template string, a ...any)
}

var Instance ILogger

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
		Instance = logger
		return nil
	default:
		return errors.New("Invalid logger type")
	}
}

func Debug(a ...any) {
	Instance.Debug(a...)
}

func Info(a ...any) {
	Instance.Info(a...)
}

func Warning(a ...any) {
	Instance.Warning(a...)
}

func Error(a ...any) {
	Instance.Error(a...)
}

func Debugf(template string, a ...any) {
	Instance.Debugf(template, a...)
}

func Infof(template string, a ...any) {
	Instance.Infof(template, a...)
}

func Warningf(template string, a ...any) {
	Instance.Warningf(template, a...)
}

func Errorf(template string, a ...any) {
	Instance.Errorf(template, a...)
}
