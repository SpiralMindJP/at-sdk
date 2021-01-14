package content

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// DeletePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/contents/:id/delete
func DeletePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := middleware.ContentFromRequest(r)

		writeDeletePage(w, r, newContentData(content), "")
	}
}

// DeleteHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/contents/:id/delete
func DeleteHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		content := middleware.ContentFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewContentServiceClient(conn)

		_, err := service.Delete(ctx, &pb.ContentRequest{
			TeamId:    user.TeamID(),
			ContentId: content.GetContentId(),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to delete the content", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/contents
		redirectToListPage(w, s)
	}
}

func writeDeletePage(w http.ResponseWriter, r *http.Request, content *contentData, errMessage string) {
	// input not valid

	if content == nil {
		content = new(contentData)
	}

	data := template.NewData(r)
	data.Title = "コンテンツ削除"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/contents", Name: "コンテンツ設定に戻る", Icon: "chevron_left"},
	}
	data.Data = content
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "content", "delete.tmpl", data)
}
