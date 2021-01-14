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

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/rooms/:id
func EditPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room := middleware.RoomFromRequest(r)

		writeEditPage(w, r, newRoomData(room), "")
	}
}

// EditHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/rooms/:id
func EditHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		room := middleware.RoomFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewRoomServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		postRoom, errMessages := verifyPost(r)
		if len(errMessages) > 0 {
			writeEditPage(w, r, postRoom, strings.Join(errMessages, "<br />"))
			return
		}

		_, err := service.Update(ctx, &pb.RoomUpdateRequest{
			TeamId: user.TeamID(),
			RoomId: room.GetRoomId(),
			Name:   postRoom.Name,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to update the room", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/rooms
		redirectToListPage(w, s)
	}
}

func writeEditPage(w http.ResponseWriter, r *http.Request, room *roomData, errMessage string) {
	// input not valid

	if room == nil {
		room = new(roomData)
	}

	data := template.NewData(r)
	data.Title = "ルーム編集"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/rooms", Name: "ルーム設定に戻る", Icon: "chevron_left"},
	}
	data.Data = room
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "room", "edit.tmpl", data)
}
