// Package log は Avatar Teleporter の各種サーバーのログ出力を行います。
package log

// Level は、ログレベルを表します。
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

// Logger は、ログ出力を行うインターフェースです。
type Logger interface {
	// ログレベルを設定します。
	SetLevel(level Level)

	// Debug ログを出力します。
	Debug(v ...interface{})
	// メッセージをフォーマットして Debug ログを出力します。
	Debugf(format string, v ...interface{})
	// msg で指定されたメッセージの Debug ログを出力します。
	// ksv は、追加で出力したい情報のキーと値の組み合わせです。
	Debugw(msg string, kvs ...interface{})

	// Info ログを出力します。
	Info(v ...interface{})
	// メッセージをフォーマットして Info ログを出力します。
	Infof(format string, v ...interface{})
	// msg で指定されたメッセージの Info ログを出力します。
	// ksv は、追加で出力したい情報のキーと値の組み合わせです。
	Infow(msg string, kvs ...interface{})

	// Warn ログを出力します。
	Warn(v ...interface{})
	// メッセージをフォーマットして Warn ログを出力します。
	Warnf(format string, v ...interface{})
	// msg で指定されたメッセージの Warn ログを出力します。
	// ksv は、追加で出力したい情報のキーと値の組み合わせです。
	Warnw(msg string, kvs ...interface{})

	// Error ログを出力します。
	Error(v ...interface{})
	// メッセージをフォーマットして Error ログを出力します。
	Errorf(format string, v ...interface{})
	// msg で指定されたメッセージの Error ログを出力します。
	// ksv は、追加で出力したい情報のキーと値の組み合わせです。
	Errorw(msg string, kvs ...interface{})

	// キーと値を指定した Logger を返します。
	With(kvs ...interface{}) Logger
}
