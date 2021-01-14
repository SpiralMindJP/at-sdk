// Package room は、Room 関連のリクエストを処理するハンドラーを提供します。
package room

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

type roomData struct {
	ID   int64
	Name string

	Device *deviceData
}

type deviceData struct {
	ID     int64
	Name   string
	RoomID int64
}

func newRoomData(room *pb.Room) *roomData {
	var device *deviceData
	if room.GetDeviceId() > 0 {
		device = &deviceData{
			ID:     room.GetDeviceId(),
			Name:   room.GetDeviceName(),
			RoomID: room.GetRoomId(),
		}
	}

	return &roomData{
		ID:     room.GetRoomId(),
		Name:   room.GetName(),
		Device: device,
	}
}

type Server interface {
	BuildURL(path string) *url.URL
}

// ListPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/rooms
func ListPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewRoomServiceClient(conn)

		result, err := service.List(ctx, &pb.RoomListRequest{
			TeamId: user.TeamID(),
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get rooms", err, http.StatusInternalServerError)
				return
			}
		}
		rooms := make([]*roomData, len(result.GetRooms()))
		for i, room := range result.GetRooms() {
			rooms[i] = newRoomData(room)
		}

		data := template.NewData(r)
		data.Title = "ルーム設定"
		data.PageCommands = []*template.PageCommand{
			{Path: "/settings/rooms/new", Name: "ルーム登録", Icon: "add_circle"},
		}
		data.Data = rooms

		template.WriteTemplate(w, r, "room", "list.tmpl", data)
	}
}

func redirectToListPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.Room)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
