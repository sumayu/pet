package logger

import (
	"strings"
	"sync"

	"github.com/sumayu/pet/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log     *zap.Logger
	once    sync.Once
	logLock sync.RWMutex
)

func Init() {
	once.Do(func() {
		initializeLogger()
	})
	
}


func initializeLogger() {
	cfg := config.LoadFromEnv()

	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(parseLevel(cfg.Level)),
		Encoding:    cfg.Encoding,
		OutputPaths: cfg.OutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
EncodeCaller:   zapcore.ShortCallerEncoder,		},
	}

	if !cfg.EnableCaller {
		zapConfig.EncoderConfig.CallerKey = ""
	}

	builtLogger, err := zapConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(parseLevel(cfg.StacktraceLevel)),
	)
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
builtLogger.Debug("Logger initialized", 
    zap.Strings("output_paths", cfg.OutputPaths),)
	logLock.Lock()
	log = builtLogger
	logLock.Unlock()
}
func GetLogger() *zap.Logger {
	logLock.RLock()
	defer logLock.RUnlock()
	return log
}
func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel 
	}
}
func Debug(msg string, fields ...zap.Field) {
	if log.Core().Enabled(zapcore.DebugLevel) {
		log.Debug(msg, fields...)
	}
}
func Info(msg string, fields ...zap.Field) {
	if log.Core().Enabled(zapcore.InfoLevel) {
		log.Info(msg, fields...)
	}
}
func Warn(msg string, fields ...zap.Field) {
	if log.Core().Enabled(zapcore.WarnLevel) {
		log.Warn(msg, fields...)
	}
}
func Error(msg string, fields ...zap.Field) {
	if log.Core().Enabled(zapcore.ErrorLevel) {
		log.Error(msg, fields...)
	}
}
func Fatal(msg string, fields ...zap.Field) {
	if log.Core().Enabled(zapcore.FatalLevel) {
		log.Fatal(msg, fields...)
	}
}
func Sync() error {
	
	logLock.RLock()
	defer logLock.RUnlock()
	if log != nil {
		return log.Sync()
	}
	
	return nil
	
}