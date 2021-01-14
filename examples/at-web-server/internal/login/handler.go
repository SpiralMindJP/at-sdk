// Package login は、ログイン関連のリクエストを処理するハンドラーを提供します。
package login

import (
	"net/http"
	"net/url"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/config"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/path"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
)

type loginPageData struct {
	LoginURL    string
	RedirectURL string

	APIKey            string
	AuthDomain        string
	ProjectID         string
	StorageBucket     string
	MessagingSenderID string
	AppID             string
}

// Handler は、以下のリクエストを処理するハンドラーです。
// GET /login
func Handler(firebaseConfig *config.Firebase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := template.NewData(r)
		data.Title = "ログイン"
		data.Data = loginPageData{
			LoginURL:          loginURL(),
			RedirectURL:       redirectURL(r),
			APIKey:            firebaseConfig.APIKey,
			AuthDomain:        firebaseConfig.AuthDomain,
			ProjectID:         firebaseConfig.ProjectID,
			StorageBucket:     firebaseConfig.StorageBucket,
			MessagingSenderID: firebaseConfig.MessagingSenderID,
			AppID:             firebaseConfig.AppID,
		}

		template.WriteTemplate(w, r, "login", "login.tmpl", data)
	}
}

func loginURL() string {
	loginURL := new(url.URL)
	loginURL.Path = path.Login

	return loginURL.String()
}

func redirectURL(r *http.Request) string {
	redirect := r.Form.Get("redirect")
	redirect, err := url.QueryUnescape(redirect)
	if err != nil || redirect == "" {
		// Dashboard のパスを返す
		return dashboardURL()
	}

	redirectURL, err := url.Parse(redirect)
	if err != nil {
		return dashboardURL()
	}

	return redirectURL.String()
}

func dashboardURL() string {
	return (&url.URL{Path: path.Dashboard}).String()
}
