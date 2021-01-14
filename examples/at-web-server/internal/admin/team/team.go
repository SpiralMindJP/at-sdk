// Package team は、Team 関連のリクエストを処理するハンドラーを提供します。
package team

import (
	"net/http"
	"net/url"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/path"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type teamData struct {
	ID   int64
	Name string
}

func newTeamData(team *pb.Team) *teamData {
	return &teamData{
		ID:   team.GetTeamId(),
		Name: team.GetName(),
	}
}

type Server interface {
	BuildURL(path string) *url.URL
}

// ListPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /admin/teams
func ListPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewTeamServiceClient(conn)

		result, err := service.List(ctx, &pb.TeamListRequest{})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get teams", err, http.StatusInternalServerError)
				return
			}
		}
		teams := make([]*teamData, len(result.GetTeams()))
		for i, team := range result.GetTeams() {
			teams[i] = newTeamData(team)
		}

		data := template.NewData(r)
		data.Title = "チーム設定"
		data.PageCommands = []*template.PageCommand{}
		data.PageCommands = []*template.PageCommand{
			{Path: "/admin/teams/new", Name: "チーム登録", Icon: "add_circle"},
		}
		data.Data = teams

		template.WriteTemplate(w, r, "admin/team", "list.tmpl", data)
	}
}

func redirectToListPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.Team)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
