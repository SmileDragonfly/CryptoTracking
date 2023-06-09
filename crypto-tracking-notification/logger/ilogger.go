package logger

import (
	"encoding/json"
	"errors"
	"os"
)

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

func NewLogger(sPath string) error {
	// Init logger
	byteCfg, err := os.ReadFile(sPath)
	if err != nil {
		panic(err)
	}
	var logCfg LoggerConfig
	err = json.Unmarshal(byteCfg, &logCfg)
	if err != nil {
		panic(err)
	}
	switch logCfg.Type {
	case ZAP_LOGGER:
		logger, err := NewZapLogger(logCfg)
		if err != nil {
			return err
		}
		Instance = logger
		break
	default:
		return errors.New("Invalid logger type")
	}
	Instance.Info("==================================================")
	Instance.Info("Start logger succesfully")
	return nil
}
