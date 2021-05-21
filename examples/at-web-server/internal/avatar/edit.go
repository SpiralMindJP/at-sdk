package avatar

import (
	"net/http"
	"strings"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/avatars/:id
func EditPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		avatar := middleware.AvatarFromRequest(r)

		writeEditPage(w, r, newEditAvatarData(avatar), "")
	}
}

// EditHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/avatars/:id
func EditHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		avatar := middleware.AvatarFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewAvatarServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		postAvatar, errMessages := verifyPost(r)
		if len(errMessages) > 0 {
			writeEditPage(w, r, postAvatar, strings.Join(errMessages, "<br />"))
			return
		}

		_, err := service.Update(ctx, &pb.AvatarUpdateRequest{
			TeamId:             user.TeamID(),
			AvatarId:           avatar.GetAvatarId(),
			Name:               postAvatar.Name,
			AvatarContentId:    postAvatar.AvatarContentID,
			AnimationContentId: postAvatar.AnimationContentID,
			Animations:         newPBAnimations(postAvatar.Animations),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to update the avatar", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/avatars
		redirectToListPage(w, s)
	}
}

func writeEditPage(w http.ResponseWriter, r *http.Request, avatar *editAvatarData, errMessage string) {
	ctx := r.Context()

	if avatar == nil {
		avatar = new(editAvatarData)
	}

	avatarContents, animationContents, err := getContents(ctx)
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			webutil.WriteError(w, r, "failed to get contents", err, http.StatusInternalServerError)
			return
		}
	}
	avatar.AvatarContents = avatarContents
	avatar.AnimationContents = animationContents

	data := template.NewData(r)
	data.Title = "アバター編集"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/avatars", Name: "アバター設定に戻る", Icon: "chevron_left"},
	}
	data.Data = avatar
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "avatar", "edit.tmpl", data)
}
