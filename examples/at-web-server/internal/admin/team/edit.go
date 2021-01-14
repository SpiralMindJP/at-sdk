package team

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /admin/teams/:id
func EditPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		team := middleware.TeamFromContext(r.Context())

		writeEditPage(w, r, newTeamData(team), "")
	}
}

// EditHandler は、以下のリクエストを処理するハンドラーです。
// POST /admin/teams/:id
func EditHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		team := middleware.TeamFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewTeamServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeEditPage(w, r, &teamData{ID: team.GetTeamId(), Name: name}, "チーム名が入力されていません。")
			return
		}

		_, err := service.Update(ctx, &pb.TeamUpdateRequest{
			TeamId: user.TeamID(),
			Name:   name,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to update the team", err, http.StatusInternalServerError)
			return
		}

		// redirect to /admin/teams
		redirectToListPage(w, s)
	}
}

func writeEditPage(w http.ResponseWriter, r *http.Request, team *teamData, errMessage string) {
	// input not valid

	if team == nil {
		team = new(teamData)
	}

	data := template.NewData(r)
	data.Title = "チーム編集"
	data.PageCommands = []*template.PageCommand{
		{Path: "/admin/teams", Name: "チーム設定に戻る", Icon: "chevron_left"},
	}
	data.Data = team
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "admin/team", "edit.tmpl", data)
}
