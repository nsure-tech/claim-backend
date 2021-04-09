package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"nsure/vote/common"
	"nsure/vote/config"
	"os"
	"sync"
)

var logger *zap.Logger
var configOnce sync.Once

func GetLog() *zap.Logger {
	configOnce.Do(func() {
		logger = NewLog()
	})
	return logger
}
func NewLog() *zap.Logger {
	conf := config.GetConfig()
	hook := lumberjack.Logger{
		Filename:   conf.Log.Filename,
		MaxSize:    conf.Log.MaxSize,
		MaxBackups: conf.Log.MaxBackups,
		MaxAge:     conf.Log.MaxAge,
		Compress:   conf.Log.Compress,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	if conf.Log.Level == common.LevelDebug {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else if conf.Log.Level == common.LevelInfo {
		atomicLevel.SetLevel(zap.InfoLevel)
	} else {
		atomicLevel.SetLevel(zap.WarnLevel)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)

	caller := zap.AddCaller()

	development := zap.Development()
	logger := zap.New(core, caller, development)
	logger.Info("Success Log Init")
	return logger
}
