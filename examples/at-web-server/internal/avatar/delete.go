package avatar

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// DeletePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/avatars/:id/delete
func DeletePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		avatar := middleware.AvatarFromRequest(r)

		writeDeletePage(w, r, newAvatarData(avatar, nil), "")
	}
}

// DeleteHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/avatars/:id/delete
func DeleteHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		avatar := middleware.AvatarFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewAvatarServiceClient(conn)

		_, err := service.Delete(ctx, &pb.AvatarRequest{
			TeamId:   user.TeamID(),
			AvatarId: avatar.GetAvatarId(),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to delete the avatar", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/avatars
		redirectToListPage(w, s)
	}
}

func writeDeletePage(w http.ResponseWriter, r *http.Request, avatar *avatarData, errMessage string) {
	// input not valid

	if avatar == nil {
		avatar = new(avatarData)
	}

	data := template.NewData(r)
	data.Title = "アバター削除"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/avatars", Name: "アバター設定に戻る", Icon: "chevron_left"},
	}
	data.Data = avatar
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "avatar", "delete.tmpl", data)
}
