package room

import (
	"net/http"
	"strings"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// CreatePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/rooms/new
func CreatePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCreatePage(w, r, nil, "")
	}
}

// CreateHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/rooms/new
func CreateHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewRoomServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		room, errMessages := verifyPost(r)
		if len(errMessages) > 0 {
			writeCreatePage(w, r, room, strings.Join(errMessages, "<br />"))
			return
		}

		_, err := service.Create(ctx, &pb.RoomCreateRequest{
			TeamId: user.TeamID(),
			Name:   room.Name,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to create new room", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/rooms
		redirectToListPage(w, s)
	}
}

func verifyPost(r *http.Request) (*roomData, []string) {
	var errMessages []string

	name := r.PostForm.Get("name")
	if name == "" {
		errMessages = append(errMessages, "ルーム名が入力されていません。")
	}

	return &roomData{
		Name: name,
	}, errMessages
}

func writeCreatePage(w http.ResponseWriter, r *http.Request, room *roomData, errMessage string) {
	// input not valid

	if room == nil {
		room = new(roomData)
	}

	data := template.NewData(r)
	data.Title = "ルーム登録"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/rooms", Name: "ルーム設定に戻る", Icon: "chevron_left"},
	}
	data.Data = room
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "room", "create.tmpl", data)
}
