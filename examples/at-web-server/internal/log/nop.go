package log

type nopLogger struct{}

// NewNopLogger は何も行わない log.Logger を返します。
func NewNopLogger() Logger {
	return nopLogger{}
}

func (l nopLogger) SetLevel(level Level)                   {}
func (l nopLogger) Debug(v ...interface{})                 {}
func (l nopLogger) Debugf(format string, v ...interface{}) {}
func (l nopLogger) Debugw(msg string, kvs ...interface{})  {}
func (l nopLogger) Info(v ...interface{})                  {}
func (l nopLogger) Infof(format string, v ...interface{})  {}
func (l nopLogger) Infow(msg string, kvs ...interface{})   {}
func (l nopLogger) Warn(v ...interface{})                  {}
func (l nopLogger) Warnf(format string, v ...interface{})  {}
func (l nopLogger) Warnw(msg string, kvs ...interface{})   {}
func (l nopLogger) Error(v ...interface{})                 {}
func (l nopLogger) Errorf(format string, v ...interface{}) {}
func (l nopLogger) Errorw(msg string, kvs ...interface{})  {}

func (l nopLogger) With(kvs ...interface{}) Logger {
	return l
}
