package auth

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
)

type User interface {
	UID() string
	Email() string
	Name() string
	TeamID() int64
	DeviceID() int64
	Admin() bool
}

type user struct {
	uid      string
	email    string
	name     string
	teamID   int64
	deviceID int64
	admin    bool
}

var _ User = (*user)(nil)

func (u *user) UID() string {
	return u.uid
}

func (u *user) Email() string {
	return u.email
}

func (u *user) Name() string {
	return u.name
}

func (u *user) TeamID() int64 {
	return u.teamID
}

func (u *user) DeviceID() int64 {
	return u.deviceID
}

func (u *user) Admin() bool {
	return u.admin
}

type Server interface {
	RedirectToLogin(w http.ResponseWriter, r *http.Request, addRedirect bool)
}

func UserProfileContext(client *auth.Client, s Server) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// クライアントから送信された ID token を取得
			cookie, err := r.Cookie("session")
			if err != nil {
				// セッション Cookie が利用できない
				// ログインにリダイレクト
				s.RedirectToLogin(w, r, true)
				return
			}

			ctx := r.Context()

			// セッション Cookie を検証
			// ユーザーの Firebase セッションが取り消された、ユーザーが削除・無効化されたなどを検出するため、
			// 追加の検証を行う。
			decoded, err := client.VerifySessionCookieAndCheckRevoked(ctx, cookie.Value)
			if err != nil {
				// セッション Cookie が不正
				// ログインにリダイレクト
				s.RedirectToLogin(w, r, true)
				return
			}

			record, err := client.GetUser(ctx, decoded.UID)
			if err != nil {
				// ユーザーが取得できない
				webutil.WriteError(w, r, "Failed to get user profile", err, http.StatusInternalServerError)
				return
			}

			teamIDFloat, ok := record.CustomClaims["team_id"].(float64)
			if !ok || teamIDFloat == 0 {
				// チームIDが取得できない
				webutil.WriteError(w, r, "Failed to get team ID that the user belongs to", err, http.StatusInternalServerError)
				return
			}
			teamID := int64(teamIDFloat)

			deviceIDFloat, ok := record.CustomClaims["device_id"].(float64)
			if !ok {
				deviceIDFloat = 0
			}
			deviceID := int64(deviceIDFloat)

			admin, ok := record.CustomClaims["admin"].(bool)
			if !ok {
				admin = false
			}

			// TODO : ユーザー情報をキャッシュする仕組みを入れる
			u := &user{
				uid:      record.UID,
				email:    record.Email,
				name:     record.DisplayName,
				teamID:   teamID,
				deviceID: deviceID,
				admin:    admin,
			}

			next.ServeHTTP(w, r.WithContext(NewContextWithUser(ctx, u)))
		})
	}
}
