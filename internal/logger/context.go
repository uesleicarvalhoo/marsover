package logger

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey string

var loggerCtxKey ctxKey = "logger_ctx_key"

func WithContext(ctx context.Context) *Entry {
	if ctx == nil {
		return &Entry{sug: defaultLogger}
	}
	if v := ctx.Value(loggerCtxKey); v != nil {
		if sug, ok := v.(*zap.SugaredLogger); ok {
			return &Entry{sug: sug}
		}
	}
	return &Entry{sug: defaultLogger}
}

func ContextWithFields(ctx context.Context, f Fields) context.Context {
	sug := defaultLogger
	if v := ctx.Value(loggerCtxKey); v != nil {
		if s2, ok := v.(*zap.SugaredLogger); ok {
			sug = s2
		}
	}

	child := sug.With(f.toKeysAndValues()...)

	return context.WithValue(ctx, loggerCtxKey, child)
}
