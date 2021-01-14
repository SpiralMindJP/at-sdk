package middleware

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
)

func AdminContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.UserFromContext(r.Context())

		if !user.Admin() {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
