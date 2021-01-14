package log

import "context"

type loggerCtxKey struct{}

// FromContext は context.Context にセットされた Logger を取得します。
func FromContext(ctx context.Context) Logger {
	logger, ok := ctx.Value(loggerCtxKey{}).(Logger)
	if !ok {
		return NewNopLogger()
	}

	return logger
}

// NewContext は logger を設定した parent のコピーを返します。
func NewContext(parent context.Context, logger Logger) context.Context {
	return context.WithValue(parent, loggerCtxKey{}, logger)
}
