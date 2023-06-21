package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log *zap.SugaredLogger
var defaultLvl zap.AtomicLevel

func init() {
	setEnvLogLvl()
	SetLogger(New())
}

// SetLogger replaces a default logger created on the start of the application with a provided logger
func SetLogger(logger *zap.SugaredLogger) {
	log = logger
}

// getEnvLogLvl searches a log level in an environment variable LOG_LVL and sets it as default.
// debug is set as default level in case if LOG_LVL is not set
func setEnvLogLvl() {
	lvl := "debug"
	_ = SetLogLvl(lvl)

	if logEnv, ok := os.LookupEnv("LOG_LVL"); ok {
		_ = SetLogLvl(logEnv)
	}
}

// SetLogLvl sets the specified log level
// Log levels could be defined as debug, info, warn, error, panic, fatal
func SetLogLvl(lvl string) error {
	err := defaultLvl.UnmarshalText([]byte(lvl))

	if err != nil {
		Error("failed to set log level as "+lvl, "reason", err.Error())
		Warn("log level is set as ", defaultLvl.String())
	} else {
		Debug("log level was successfully set as " + lvl)
	}

	return err
}

// New creates a new instance of zap.SugaredLogger with json encoder, default level specified in LOG_LVL env variable
// and a time format equals to 2006-01-02T15:04:05.000Z0700
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
