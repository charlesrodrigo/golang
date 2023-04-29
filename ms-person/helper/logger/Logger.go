package logger

import (
	"context"
	"fmt"
	"os"
	"strings"

	"br.com.charlesrodrigo/ms-person/helper/constants"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func Init(appName string) *zap.Logger {
	initialCustomFields := make(map[string]interface{})
	initialCustomFields["service"] = appName

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
		InitialFields: initialCustomFields,
	}

	zapLog, _ := logConfig.Build()

	defer zapLog.Sync()

	log = zapLog

	return zapLog
}

func addTrace(c context.Context, message string, args ...interface{}) (fields []zap.Field) {
	span := trace.SpanFromContext(c)
	args = append(args, constants.REQUEST_ID, c.Value(constants.REQUEST_ID))
	args = append(args, "trace-ID", span.SpanContext().TraceID().String())
	attrs := make([]attribute.KeyValue, 0)
	fields = make([]zap.Field, 0)

	if len(args)%2 == 0 {
		for i, arg := range args {
			if i%2 == 0 {
				str := fmt.Sprintf("%v", arg)
				value := fmt.Sprintf("%v", args[i+1])
				key := attribute.Key(str)
				attrs = append(attrs, key.String(value))
				fields = append(fields, zap.String(str, value))
			}
		}
	}

	logMessageKey := attribute.Key("message")
	attrs = append(attrs, logMessageKey.String(message))

	span.AddEvent(zapcore.InfoLevel.String(), trace.WithAttributes(attrs...))

	return
}

func Info(message string) {
	log.Info(message)
}

func InfoWithContext(c context.Context, message string, args ...interface{}) {
	fields := addTrace(c, message, args...)
	log.Info(message, fields...)
}

func ErrorWithContext(c context.Context, message string, args ...interface{}) {
	fields := addTrace(c, message, args...)
	log.Error(message, fields...)
}

func PanicWithContext(c context.Context, message string, args ...interface{}) {
	fields := addTrace(c, message, args...)
	log.Panic(message, fields...)
}

func FatalWithContext(c context.Context, message string, args ...interface{}) {
	fields := addTrace(c, message, args...)
	log.Fatal(message, fields...)
}

func Error(message string) {
	log.Error(message)
}

func Panic(message string) {
	log.Panic(message)
}

func Fatal(message string) {
	log.Fatal(message)
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
