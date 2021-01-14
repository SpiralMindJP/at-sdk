// Package content は、Content 関連のリクエストを処理するハンドラーを提供します。
package content

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

type contentData struct {
	ID   int64
	Name string
}

func newContentData(content *pb.Content) *contentData {
	return &contentData{
		ID:   content.GetContentId(),
		Name: content.GetName(),
	}
}

type Server interface {
	BuildURL(path string) *url.URL
}

// ListPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/contents
func ListPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewContentServiceClient(conn)

		result, err := service.List(ctx, &pb.ContentListRequest{
			TeamId: user.TeamID(),
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get contents", err, http.StatusInternalServerError)
				return
			}
		}
		contents := make([]*contentData, len(result.GetContents()))
		for i, room := range result.GetContents() {
			contents[i] = newContentData(room)
		}

		data := template.NewData(r)
		data.Title = "コンテンツ設定"
		data.PageCommands = []*template.PageCommand{
			{Path: "/settings/contents/new", Name: "コンテンツ登録", Icon: "add_circle"},
		}
		data.Data = contents

		template.WriteTemplate(w, r, "content", "list.tmpl", data)
	}
}

func redirectToListPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.Content)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
