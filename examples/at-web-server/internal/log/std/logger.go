// Package std は、標準出力にログを出力する log.Logger です。
package std

import (
	"fmt"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
)

type stdLogger struct {
	level log.Level
	w     *writer
	kvs   []interface{}
}

func NewStdLogger() log.Logger {
	return &stdLogger{
		w: newWriter(),
	}
}

func (l *stdLogger) SetLevel(level log.Level) {
	l.level = level
}

func (l *stdLogger) Debug(v ...interface{}) {
	l.log(log.Debug, v...)
}

func (l *stdLogger) Debugf(format string, v ...interface{}) {
	l.logf(log.Debug, format, v...)
}

func (l *stdLogger) Debugw(msg string, kvs ...interface{}) {
	l.logw(log.Debug, msg, kvs...)
}

func (l *stdLogger) Info(v ...interface{}) {
	l.log(log.Info, v...)
}

func (l *stdLogger) Infof(format string, v ...interface{}) {
	l.logf(log.Info, format, v...)
}

func (l *stdLogger) Infow(msg string, kvs ...interface{}) {
	l.logw(log.Info, msg, kvs...)
}

func (l *stdLogger) Warn(v ...interface{}) {
	l.log(log.Warn, v...)
}

func (l *stdLogger) Warnf(format string, v ...interface{}) {
	l.logf(log.Warn, format, v...)
}

func (l *stdLogger) Warnw(msg string, kvs ...interface{}) {
	l.logw(log.Warn, msg, kvs...)
}

func (l *stdLogger) Error(v ...interface{}) {
	l.log(log.Error, v...)
}

func (l *stdLogger) Errorf(format string, v ...interface{}) {
	l.logf(log.Error, format, v...)
}

func (l *stdLogger) Errorw(msg string, kvs ...interface{}) {
	l.logw(log.Error, msg, kvs...)
}

func (l *stdLogger) log(level log.Level, v ...interface{}) {
	if level < l.level {
		return
	}

	msg := fmt.Sprintln(v...)
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}

	l.logw(level, msg)
}

func (l *stdLogger) logf(level log.Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logw(level, msg)
}

func (l *stdLogger) logw(level log.Level, msg string, kvs ...interface{}) {
	if level < l.level {
		return
	}

	var kvs_ []interface{}
	if len(l.kvs) > 0 && len(kvs) > 0 {
		kvs_ = make([]interface{}, 0, len(l.kvs)+len(kvs))
		kvs_ = append(kvs_, l.kvs...)
		kvs_ = append(kvs_, kvs...)
	} else if len(l.kvs) > 0 {
		kvs_ = l.kvs
	} else if len(kvs) > 0 {
		kvs_ = kvs
	}

	l.w.logw(level, msg, kvs_...)
}

func (l *stdLogger) With(kvs ...interface{}) log.Logger {
	var kvs_ []interface{}
	c := len(l.kvs) + len(kvs)
	if c > 0 {
		kvs_ = make([]interface{}, 0, c)
		kvs_ = append(kvs_, l.kvs...)
		kvs_ = append(kvs_, kvs...)
	}

	return &stdLogger{
		level: l.level,
		w:     l.w,
		kvs:   kvs_,
	}
}
