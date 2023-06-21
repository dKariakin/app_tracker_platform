package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log *zap.SugaredLogger
var defaultLvl = zap.NewAtomicLevelAt(zapcore.DebugLevel)

func init() {
	SetLogger(New())
}

func SetLogger(logger *zap.SugaredLogger) {
	log = logger
}

// SetLogLvl sets the specified log level
// Log levels could be defined as debug, info, warn, error, panic, fatal
func SetLogLvl(lvl string) error {
	err := defaultLvl.UnmarshalText([]byte(lvl))

	if err != nil {
		Error("failed to set log level as "+lvl, "reason", err.Error())
	} else {
		Debug("log level was successfully set as " + lvl)
	}

	return err
}

func New() *zap.SugaredLogger {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(zapcore.NewCore(
		zapcore.NewJSONEncoder(conf),
		zapcore.Lock(os.Stderr),
		defaultLvl,
	))
	l := zap.New(core)

	return l.Sugar()
}

func Info(msg string, args ...interface{}) {
	log.Infow(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Debugw(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Warnw(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Errorw(msg, args...)
}

func Panic(msg string, args ...interface{}) {
	log.Panicw(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	log.Fatalw(msg, args...)
}
