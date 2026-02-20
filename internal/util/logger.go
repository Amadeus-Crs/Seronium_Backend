package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

func encodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(Green + "INFO" + Reset)
	case zapcore.WarnLevel:
		enc.AppendString(Yellow + "WARN" + Reset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(Red + "ERROR" + Reset)
	default:
		enc.AppendString(level.String())
	}
}

func InitLogger() {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = encodeLevel
	logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func L() *zap.Logger {
	return logger
}
