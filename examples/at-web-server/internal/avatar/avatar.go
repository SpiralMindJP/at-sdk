// Package avatar は、Avatar 関連のリクエストを処理するハンドラーを提供します。
package avatar

import (
	"net/http"
	"net/url"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/path"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type avatarData struct {
	ID   int64
	Name string

	Avatar    *contentData
	Animation *contentData
}

type contentData struct {
	ID   int64
	Name string
}

func newAvatarData(avatar *pb.Avatar, avatarContent *pb.Content, animationContent *pb.Content) *avatarData {
	var avc *contentData
	if avatarContent != nil {
		avc = &contentData{
			ID:   avatarContent.GetContentId(),
			Name: avatarContent.GetName(),
		}
	}

	var anc *contentData
	if animationContent != nil {
		anc = &contentData{
			ID:   animationContent.GetContentId(),
			Name: animationContent.GetName(),
		}
	}

	return &avatarData{
		ID:        avatar.GetAvatarId(),
		Name:      avatar.GetName(),
		Avatar:    avc,
		Animation: anc,
	}
}

type Server interface {
	BuildURL(path string) *url.URL
}

// ListPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/avatars
func ListPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		teamID := user.TeamID()

		service := pb.NewAvatarServiceClient(conn)
		contentService := pb.NewContentServiceClient(conn)

		result, err := service.List(ctx, &pb.AvatarListRequest{
			TeamId: teamID,
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get avatars", err, http.StatusInternalServerError)
				return
			}
		}
		contents, err := contentService.List(ctx, &pb.ContentListRequest{
			TeamId: teamID,
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get contents", err, http.StatusInternalServerError)
				return
			}
		}

		var contentMap map[int64]*pb.Content
		if len(contents.GetContents()) > 0 {
			contentMap = map[int64]*pb.Content{}

			for _, content := range contents.GetContents() {
				contentMap[content.GetContentId()] = content
			}
		}

		avatars := make([]*avatarData, len(result.GetAvatars()))
		for i, avatar := range result.GetAvatars() {
			avatarContent := contentMap[avatar.GetAvatarContentId()]
			animationContent := contentMap[avatar.GetAnimationContentId()]
			avatars[i] = newAvatarData(avatar, avatarContent, animationContent)
		}

		data := template.NewData(r)
		data.Title = "アバター設定"
		data.PageCommands = []*template.PageCommand{
			{Path: "/settings/avatars/new", Name: "アバター登録", Icon: "add_circle"},
		}
		data.Data = avatars

		template.WriteTemplate(w, r, "avatar", "list.tmpl", data)
	}
}

func redirectToListPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.Avatar)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
