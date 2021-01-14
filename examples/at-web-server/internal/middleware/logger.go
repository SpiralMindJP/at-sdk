package middleware

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
)

// LoggerContext は、logger を HTTP リクエストのコンテキストにセットするミドルウェアを返します。
func LoggerContext(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := log.NewContext(r.Context(), logger)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
