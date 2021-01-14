package room

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// DeletePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/rooms/:id/delete
func DeletePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room := middleware.RoomFromRequest(r)

		writeDeletePage(w, r, newRoomData(room), "")
	}
}

// DeleteHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/rooms/:id/delete
func DeleteHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		room := middleware.RoomFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewRoomServiceClient(conn)

		_, err := service.Delete(ctx, &pb.RoomRequest{
			TeamId: user.TeamID(),
			RoomId: room.GetRoomId(),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to delete the room", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/rooms
		redirectToListPage(w, s)
	}
}

func writeDeletePage(w http.ResponseWriter, r *http.Request, room *roomData, errMessage string) {
	// input not valid

	if room == nil {
		room = new(roomData)
	}

	data := template.NewData(r)
	data.Title = "ルーム削除"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/rooms", Name: "ルーム設定に戻る", Icon: "chevron_left"},
	}
	data.Data = room
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "room", "delete.tmpl", data)
}
