package winter

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(viper *viper.Viper) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   viper.GetString("logging.file"),
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}
	write := zapcore.AddSync(&hook)

	var level zapcore.Level

	switch viper.GetString("logging.level") {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := ecszap.NewDefaultEncoderConfig()

	var writes = []zapcore.WriteSyncer{write}

	if level == zap.DebugLevel {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}

	core := ecszap.NewCore(
		encoderConfig,
		zapcore.NewMultiWriteSyncer(writes...),
		level,
	)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.Fields(
			zap.String("application", viper.GetString("spring.application.name")),
			zap.String("serverIp", LocalIP()),
			zap.Int("port", viper.GetInt("server.port")),
			zap.String("profile", viper.GetString("spring.profiles.active")),
			zap.String("logTime", time.Now().Format(viper.GetString("logging.date-time-format"))),
		),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}
