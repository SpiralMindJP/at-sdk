package auth

import (
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
)

func LoginHandler(client *auth.Client, secure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		idToken, ok := getIDTokenFromBody(w, r)
		if !ok {
			return
		}

		decoded, err := client.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			webutil.WriteError(w, r, "Invalid ID token", err, http.StatusUnauthorized)
			return
		}
		// Return error if the sign-in is older than 5 minutes.
		if time.Now().Unix()-int64(decoded.Claims["auth_time"].(float64)) > 5*60 {
			webutil.WriteError(w, r, "Recent sign-in required", err, http.StatusUnauthorized)
			return
		}

		expiresIn := time.Hour * 24 * 5
		cookie, err := client.SessionCookie(r.Context(), idToken, expiresIn)
		if err != nil {
			webutil.WriteError(w, r, "Failed to create a session cookie", err, http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    cookie,
			MaxAge:   int(expiresIn.Seconds()),
			HttpOnly: true,
			Secure:   secure,
		})

		data := struct {
			Status string `json:"status"`
		}{
			Status: "success",
		}
		webutil.WriteJSON(w, r, data)
	}
}

func getIDTokenFromBody(w http.ResponseWriter, r *http.Request) (string, bool) {
	var parsedBody struct {
		IDToken string `json:"id_token"`
	}

	if !webutil.DecodeBody(w, r, &parsedBody) {
		return "", false
	}

	return parsedBody.IDToken, true
}
