package content

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/contents/:id
func EditPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := middleware.ContentFromRequest(r)

		writeEditPage(w, r, newContentData(content), "")
	}
}

// EditHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/contents/:id
func EditHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		content := middleware.ContentFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewContentServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeEditPage(w, r, &contentData{ID: content.GetContentId(), Name: name}, "コンテンツ名が入力されていません。")
			return
		}

		_, err := service.Update(ctx, &pb.ContentUpdateRequest{
			TeamId:    user.TeamID(),
			ContentId: content.GetContentId(),
			Name:      name,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to update the content", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/contents
		redirectToListPage(w, s)
	}
}

func writeEditPage(w http.ResponseWriter, r *http.Request, content *contentData, errMessage string) {
	// input not valid

	if content == nil {
		content = new(contentData)
	}

	data := template.NewData(r)
	data.Title = "コンテンツ編集"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/contents", Name: "コンテンツ設定に戻る", Icon: "chevron_left"},
	}
	data.Data = content
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "content", "edit.tmpl", data)
}
