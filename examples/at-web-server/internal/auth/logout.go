package auth

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
)

func LogoutHandler(client *auth.Client, s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			// セッション Cookie が利用できない
			// ログインにリダイレクト
			s.RedirectToLogin(w, r, false)
			return
		}

		decoded, err := client.VerifySessionCookie(r.Context(), cookie.Value)
		if err != nil {
			// セッション Cookie が不正
			// ログインにリダイレクト
			s.RedirectToLogin(w, r, false)
			return
		}

		if err := client.RevokeRefreshTokens(r.Context(), decoded.UID); err != nil {
			webutil.WriteError(w, r, "failed to revoke refresh token", err, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: 0,
		})
		s.RedirectToLogin(w, r, false)
	}
}
