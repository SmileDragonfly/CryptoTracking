package logger

import (
	"encoding/json"
	"errors"
	"os"
	"testing"
)

func TestZapLogger_Debug(t *testing.T) {
	type fields struct {
		logger *ZapLogger
	}
	type args struct {
		a []any
	}
	// Get logger config
	byteCfg, err := os.ReadFile("../config/logcfg.json")
	if err != nil {
		t.Errorf("Read config file failed: %s", err)
		return
	}
	// Get zap logger
	var logCfg LoggerConfig
	err = json.Unmarshal(byteCfg, &logCfg)
	if err != nil {
		t.Errorf("Unmarshal config file failed: %s", err)
		return
	}
	zapLogger, err := NewZapLogger(logCfg)
	if err != nil {
		t.Errorf("Create zap logger failed: %s", err)
		return
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"TestZapLogger_Debug",
			fields{
				logger: zapLogger,
			},
			args{[]any{errors.New("This is test debug log 1"), errors.New("This is test debug log 2")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := ZapLogger{tt.fields.logger.}
			z.Debug(tt.args.a...)
		})
	}
}

func TestZapLogger_Error(t *testing.T) {
	type fields struct {
		logger *ZapLogger
	}
	type args struct {
		a []any
	}
	// Get logger config
	byteCfg, err := os.ReadFile("../config/logcfg.json")
	if err != nil {
		t.Errorf("Read config file failed: %s", err)
		return
	}
	// Get zap logger
	var logCfg LoggerConfig
	err = json.Unmarshal(byteCfg, &logCfg)
	if err != nil {
		t.Errorf("Unmarshal config file failed: %s", err)
		return
	}
	zapLogger, err := NewZapLogger(logCfg)
	if err != nil {
		t.Errorf("Create zap logger failed: %s", err)
		return
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"TestZapLogger_Debug",
			fields{
				logger: zapLogger,
			},
			args{[]any{errors.New("This is test error log 1"), errors.New("This is test error log 2")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := ZapLogger{tt.fields.logger.}
			z.Error(tt.args.a...)
		})
	}
}
