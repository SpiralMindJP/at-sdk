package user

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	webauth "github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
)

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/user/edit
func EditPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := webauth.UserFromContext(r.Context())

		writeEditPage(w, r, newUserData(user, nil), "")
	}
}

// EditHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/user/edit
func EditHandler(client *auth.Client, s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := webauth.UserFromContext(ctx)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeEditPage(w, r, &userData{Name: name}, "ユーザー名が入力されていません。")
			return
		}

		params := new(auth.UserToUpdate).
			DisplayName(name).
			Disabled(false)
		_, err := client.UpdateUser(ctx, user.UID(), params)
		if err != nil {
			webutil.WriteError(w, r, "failed to update the user", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/user
		redirectToUserPage(w, s)
	}
}

func writeEditPage(w http.ResponseWriter, r *http.Request, user *userData, errMessage string) {
	// input not valid

	data := template.NewData(r)
	data.Title = "ユーザー情報編集"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/user", Name: "ユーザー設定に戻る", Icon: "person"},
	}
	data.Data = user
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "user", "edit.tmpl", data)
}
