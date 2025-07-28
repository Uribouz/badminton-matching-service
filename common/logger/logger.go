package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	Log *Logger
)

type Logger struct {
	logger *zap.SugaredLogger
}

func InitLogger() *Logger {
	sync.OnceFunc(func() {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic("Failed to initialize logger: " + err.Error())
		}
		sugar := logger.Sugar()
		Log = &Logger{sugar}
	})()
	return Log
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
func (l *Logger) Sync() error {
	return l.logger.Sync()
}
