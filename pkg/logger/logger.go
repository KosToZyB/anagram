package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})

	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}

func GetLogger(logLevel string) (log Logger, err error) {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	level := getLevel(logLevel)
	atom.SetLevel(level)

	sugar := logger.Sugar() //if you want more faster logger, you don't use sugar

	return sugar, nil
}

func getLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
