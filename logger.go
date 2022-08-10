package main

import (
	"github.com/tsotosa/atmm/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	globalLogger *zap.Logger
)

// Sync call it in defer
func Sync() error {
	return globalLogger.Sync()
}

func InitLogger(logFile string) (logger *zap.Logger, err error) {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // short path encoder
		EncodeName:     zapcore.FullNameEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	if config.Conf.Environment == "dev" {
		encoderConfig = zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // short path encoder
			EncodeName:     zapcore.FullNameEncoder,
		}
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    config.Conf.LogRotateMaxLogFileSize, //MB
		MaxBackups: config.Conf.LogRotateMaxNumOfBackups,
		MaxAge:     config.Conf.LogRotateMaxAgeOfBackups, //days
		Compress:   config.Conf.LogRotateCompressBackups,
	})
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), ws),
		getLevel(),
	)

	var _globalLogger *zap.Logger
	_globalLogger = zap.New(core, zap.AddCaller())
	globalLogger = _globalLogger
	return globalLogger, nil
}

func getLevel() zap.AtomicLevel {
	if config.Conf.LogLevel != "" {
		switch config.Conf.LogLevel {
		case "debug":
			return zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info":
			return zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warn":
			return zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error":
			return zap.NewAtomicLevelAt(zap.ErrorLevel)
		default:
			return zap.NewAtomicLevelAt(zap.DebugLevel)
		}
	}
	return zap.NewAtomicLevelAt(zap.DebugLevel)
}
