package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type ZapLogger struct {
	config LoggerConfig
	logger *zap.SugaredLogger
}

func NewZapLogger(config LoggerConfig) (*ZapLogger, error) {
	// lumberjack.Instance is already safe for concurrent use, so we don't need to
	// lock it.
	encoderCfg := zapcore.EncoderConfig{ //Cấu hình logging, sẽ không có stacktracekey
		MessageKey:   "message",
		TimeKey:      "time",
		LevelKey:     "level",
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder, //Lấy dòng code bắt đầu log
		EncodeLevel:  CustomLevelEncoder,        //Format cách hiển thị level log
		EncodeTime:   SyslogTimeEncoder,         //Format hiển thị thời điểm log
	}

	// Create a writer that logs to the console (os.Stdout)
	consoleDebugging := zapcore.Lock(os.Stdout)
	// Create a writer that logs to a file using lumberjack logger
	fileLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
	})
	// Combine both writers into a MultiWriteSyncer
	writeSyncer := zapcore.NewMultiWriteSyncer(consoleDebugging, zapcore.AddSync(fileLogger))
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		writeSyncer,
		zapcore.Level(config.Level),
	)
	logger := zap.New(core)
	return &ZapLogger{config: config, logger: logger.Sugar()}, nil
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func (z ZapLogger) Debug(a ...any) {
	z.logger.Debug(a...)
}

func (z ZapLogger) Info(a ...any) {
	z.logger.Info(a...)
}

func (z ZapLogger) Warning(a ...any) {
	z.logger.Warn(a...)
}
func (z ZapLogger) Error(a ...any) {
	z.logger.Error(a...)
}

func (z ZapLogger) Debugf(template string, a ...any) {
	z.logger.Debugf(template, a...)
}

func (z ZapLogger) Infof(template string, a ...any) {
	z.logger.Infof(template, a...)
}

func (z ZapLogger) Warningf(template string, a ...any) {
	z.logger.Warnf(template, a...)
}
func (z ZapLogger) Errorf(template string, a ...any) {
	z.logger.Errorf(template, a...)
}
