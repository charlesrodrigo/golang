package logger

import (
	"os"
	"strings"

	"br.com.charlesrodrigo/ms-person/helper/constants"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugar *zap.SugaredLogger
)

func Init() *zap.Logger {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    os.Getenv(constants.LOG_TYPE),
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zap, _ := logConfig.Build()

	defer zap.Sync()

	sugar = zap.Sugar()

	return zap
}

func Info(message string, args ...interface{}) {
	sugar.Infof(message, args)
}

func Error(message string, args ...interface{}) {
	sugar.Errorf(message, args)
}

func Panic(message string, args ...interface{}) {
	sugar.Panicf(message, args)
}

func Fatal(message string, args ...interface{}) {
	sugar.Fatalf(message, args)
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(constants.LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(constants.LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
