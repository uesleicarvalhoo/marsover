package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *zap.SugaredLogger

func Configure(params Params) {
	zl := buildZap(params)

	zap.ReplaceGlobals(zl)
	defaultLogger = zl.Sugar()
}

func init() {
	Configure(Params{
		Level:          "INFO",
		ServiceName:    "mars-rover",
		ServiceVersion: "0.0.0",
		Env:            "dev",
	})
}

func buildZap(params Params) *zap.Logger {
	var lvl zapcore.Level
	if err := lvl.Set(params.Level); err != nil {
		panic(err)
	}

	enc := zap.NewProductionEncoderConfig()
	enc.MessageKey = "message"
	enc.TimeKey = "time"
	enc.LevelKey = "level"
	enc.EncodeTime = zapcore.ISO8601TimeEncoder
	enc.EncodeCaller = zapcore.ShortCallerEncoder

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(lvl),
		Development:       params.Env != "prd",
		DisableCaller:     false,
		DisableStacktrace: true,
		Encoding:          "json",
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig:     enc,
		InitialFields: map[string]any{
			"logVersion": "2.0.0",
			"app": map[string]string{
				"name":    params.ServiceName,
				"version": params.ServiceVersion,
				"env":     params.Env,
			},
		},
	}

	return zap.Must(cfg.Build()).WithOptions(zap.AddCallerSkip(1))
}

func Info(msg string, args ...any) {
	defaultLogger.Infof(msg, args...)
}

func Debug(msg string, args ...any) {
	defaultLogger.Debugf(msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Warnf(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Errorf(msg, args...)
}

func Fatal(msg string, args ...any) {
	defaultLogger.Fatalf(msg, args...)
}

func InfoF(msg string, f Fields) {
	defaultLogger.Infow(msg, f.toKeysAndValues()...)
}

func DebugF(msg string, f Fields) {
	defaultLogger.Debugw(msg, f.toKeysAndValues()...)
}

func WarnF(msg string, f Fields) {
	defaultLogger.Warnw(msg, f.toKeysAndValues()...)
}

func ErrorF(msg string, f Fields) {
	defaultLogger.Errorw(msg, f.toKeysAndValues()...)
}

func FatalF(msg string, f Fields) {
	defaultLogger.Fatalw(msg, f.toKeysAndValues()...)
}
