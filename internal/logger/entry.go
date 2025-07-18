package logger

import "go.uber.org/zap"

type Entry struct {
	sug *zap.SugaredLogger
}

func (e *Entry) WithFields(f Fields) *Entry {
	return &Entry{
		sug: e.sug.With(f.toKeysAndValues()...),
	}
}

func (e *Entry) Info(msg string, args ...any) {
	e.sug.Infof(msg, args...)
}

func (e *Entry) Debug(msg string, args ...any) {
	e.sug.Debugf(msg, args...)
}

func (e *Entry) Warn(msg string, args ...any) {
	e.sug.Warnf(msg, args...)
}

func (e *Entry) Error(msg string, args ...any) {
	e.sug.Errorf(msg, args...)
}

func (e *Entry) Fatal(msg string, args ...any) {
	e.sug.Fatalf(msg, args...)
}

func (e *Entry) InfoF(msg string, f Fields) {
	e.sug.Infow(msg, f.toKeysAndValues()...)
}

func (e *Entry) DebugF(msg string, f Fields) {
	e.sug.Debugw(msg, f.toKeysAndValues()...)
}

func (e *Entry) WarnF(msg string, f Fields) {
	e.sug.Warnw(msg, f.toKeysAndValues()...)
}

func (e *Entry) ErrorF(msg string, f Fields) {
	e.sug.Errorw(msg, f.toKeysAndValues()...)
}

func (e *Entry) FatalF(msg string, f Fields) {
	e.sug.Fatalw(msg, f.toKeysAndValues()...)
}
