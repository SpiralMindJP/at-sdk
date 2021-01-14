package user

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	webauth "github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
)

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/user/password
func ChangePasswordPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeChangePasswordPage(w, r, "")
	}
}

// ListPageHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/user/password
func ChangePasswordHandler(client *auth.Client, s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := webauth.UserFromContext(ctx)

		if !webutil.ParseForm(w, r) {
			return
		}

		password := r.PostForm.Get("password")
		if password == "" {
			writeChangePasswordPage(w, r, "新しいパスワードが入力されていません。")
			return
		}
		confirmPassword := r.PostForm.Get("confirmpassword")
		if confirmPassword == "" {
			writeChangePasswordPage(w, r, "確認用の新しいパスワードが入力されていません。")
			return
		}

		if password != confirmPassword {
			writeChangePasswordPage(w, r, "確認用の新しいパスワードが一致しません。")
			return
		}

		params := new(auth.UserToUpdate).
			Password(password).
			Disabled(false)

		if _, err := client.UpdateUser(ctx, user.UID(), params); err != nil {
			webutil.WriteError(w, r, "failed to change the password", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/user
		redirectToUserPage(w, s)
	}
}

func writeChangePasswordPage(w http.ResponseWriter, r *http.Request, errMessage string) {
	// input not valid

	data := template.NewData(r)
	data.Title = "パスワード変更"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/user", Name: "ユーザー設定に戻る", Icon: "person"},
	}
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "user", "password.tmpl", data)
}
