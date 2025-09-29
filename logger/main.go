package logger

import (
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type ConcurrentLogger struct {
	logger zerolog.Logger
	mu     sync.Mutex
}

func New(appName, appLogLevel string) *ConcurrentLogger {
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

	return &ConcurrentLogger{
		logger: logger,
	}
}

func (l *ConcurrentLogger) Log(level zerolog.Level, msg string, fields ...map[string]any) {
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

func (l *ConcurrentLogger) Debug(msg string, fields ...map[string]any) {
	l.Log(zerolog.DebugLevel, msg, fields...)
}

func (l *ConcurrentLogger) Info(msg string, fields ...map[string]any) {
	l.Log(zerolog.InfoLevel, msg, fields...)
}

func (l *ConcurrentLogger) Warn(msg string, fields ...map[string]any) {
	l.Log(zerolog.WarnLevel, msg, fields...)
}

func (l *ConcurrentLogger) Error(msg string, fields ...map[string]any) {
	l.Log(zerolog.ErrorLevel, msg, fields...)
}
