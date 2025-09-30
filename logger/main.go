package logger

import (
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type concurrentLogger struct {
	logger zerolog.Logger
	mu     sync.Mutex
}

func New(appName, appLogLevel string) *concurrentLogger {
	rotator := &lumberjack.Logger{
		Filename:   "./log/" + appName + ".log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	}

	var logLevel zerolog.Level
	switch appLogLevel {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	default:
		logLevel = zerolog.WarnLevel
	}

	logger := zerolog.New(rotator).Level(logLevel).With().Timestamp().Logger()

	return &concurrentLogger{
		logger: logger,
	}
}

func (l *concurrentLogger) log(level zerolog.Level, msg string, fields ...map[string]any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	event := l.logger.WithLevel(level)

	for _, fieldMap := range fields {
		for k, v := range fieldMap {
			event = event.Interface(k, v)
		}
	}

	event.Msg(msg)
}

func (l *concurrentLogger) Debug(msg string, fields ...map[string]any) {
	l.log(zerolog.DebugLevel, msg, fields...)
}

func (l *concurrentLogger) Info(msg string, fields ...map[string]any) {
	l.log(zerolog.InfoLevel, msg, fields...)
}

func (l *concurrentLogger) Warn(msg string, fields ...map[string]any) {
	l.log(zerolog.WarnLevel, msg, fields...)
}

func (l *concurrentLogger) Error(msg string, fields ...map[string]any) {
	l.log(zerolog.ErrorLevel, msg, fields...)
}
