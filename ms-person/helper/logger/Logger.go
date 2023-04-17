package logger

import (
	"os"
	"strings"

	"br.com.charlesrodrigo/ms-person/helper/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func Init() *zap.Logger {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    constants.LOG_TYPE,
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	config, _ := logConfig.Build()
	log = config.Sugar()

	return config
}

func Info(message ...interface{}) {
	log.Info(message)
	log.Sync()
}

func Error(message ...interface{}) {
	log.Error(message)
	log.Sync()
}

func Panic(message ...interface{}) {
	log.Panic(message)
	log.Sync()
}

func Fatal(message ...interface{}) {
	log.Fatal(message)
	log.Sync()
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
