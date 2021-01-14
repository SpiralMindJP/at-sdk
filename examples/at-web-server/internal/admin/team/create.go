package team

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// CreatePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /admin/teams/new
func CreatePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCreatePage(w, r, nil, "")
	}
}

// CreateHandler は、以下のリクエストを処理するハンドラーです。
// POST /admin/teams/new
func CreateHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewTeamServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeEditPage(w, r, &teamData{Name: name}, "チーム名が入力されていません。")
			return
		}

		_, err := service.Create(ctx, &pb.TeamCreateRequest{
			Name: name,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to create a new team", err, http.StatusInternalServerError)
			return
		}

		// redirect to /admin/teams
		redirectToListPage(w, s)
	}
}

func writeCreatePage(w http.ResponseWriter, r *http.Request, team *teamData, errMessage string) {
	// input not valid

	if team == nil {
		team = new(teamData)
	}

	data := template.NewData(r)
	data.Title = "チーム登録"
	data.PageCommands = []*template.PageCommand{
		{Path: "/admin/teams", Name: "チーム設定に戻る", Icon: "chevron_left"},
	}
	data.Data = team
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "admin/team", "create.tmpl", data)
}
