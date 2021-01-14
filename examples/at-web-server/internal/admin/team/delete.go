package team

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// DeletePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /admin/teams/:id/delete
func DeletePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		team := middleware.TeamFromContext(r.Context())

		writeDeletePage(w, r, newTeamData(team), "")
	}
}

// DeleteHandler は、以下のリクエストを処理するハンドラーです。
// POST /admin/teams/:id/delete
func DeleteHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		team := middleware.TeamFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewTeamServiceClient(conn)

		_, err := service.Delete(ctx, &pb.TeamRequest{
			TeamId: team.GetTeamId(),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to delete the team", err, http.StatusInternalServerError)
			return
		}

		// redirect to /admin/teams
		redirectToListPage(w, s)
	}
}

func writeDeletePage(w http.ResponseWriter, r *http.Request, team *teamData, errMessage string) {
	// input not valid

	if team == nil {
		team = new(teamData)
	}

	data := template.NewData(r)
	data.Title = "チーム削除"
	data.PageCommands = []*template.PageCommand{
		{Path: "/admin/teams", Name: "チーム設定に戻る", Icon: "chevron_left"},
	}
	data.Data = team
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "admin/team", "delete.tmpl", data)
}
