package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(config LoggerConfig) (*ZapLogger, error) {
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
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
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
	})

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		w,
		zapcore.Level(config.Level),
	)
	logger := zap.New(core)
	return &ZapLogger{logger: logger.Sugar()}, nil
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func (z ZapLogger) Debug(a ...any) {
	z.logger.Debug(a)
}

func (z ZapLogger) Info(a ...any) {
	z.logger.Info(a)
}

func (z ZapLogger) Warning(a ...any) {
	z.logger.Warn(a)
}
func (z ZapLogger) Error(a ...any) {
	z.logger.Error(a)
}
