// Package dashboard は、Dashboard 関連のリクエストを処理するハンドラーを提供します。
package dashboard

import (
	"net/http"
	"net/url"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/dashboard/jsondata"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/path"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type dasbhardData struct {
	EventStreamURL string
}

type Server interface {
	BuildWebSocketURL(path string) *url.URL
}

// Handler は、以下のリクエストを処理するハンドラーです。
// GET /dashboard
func Handler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := template.NewData(r)
		data.Title = "Dashboard"
		data.Data = &dasbhardData{
			EventStreamURL: s.BuildWebSocketURL(path.DashboardEventStream).String(),
		}

		template.WriteTemplate(w, r, "dashboard", "dashboard.tmpl", data)
	}
}

// DashboardAPIHandler は、以下のリクエストを処理するハンドラーです。
// GET /api/dashboard
func DashboardAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeDashboardJSON(w, r)
	}
}

// JoinAPIHandler は、以下のリクエストを処理するハンドラーです。
// POST /api/dashboard/rooms/{{roomID}}/join
func JoinAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		room := middleware.RoomFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		if user.DeviceID() == 0 {
			writeErrorJSON(w, r, jsondata.ErrorDeviceNotRegistered, "", nil)
			return
		}

		dashboardService := pb.NewDashboardServiceClient(conn)

		state, err := dashboardService.JoinRoom(ctx, &pb.JoinRoomRequest{
			TeamId:   user.TeamID(),
			RoomId:   room.GetRoomId(),
			DeviceId: user.DeviceID(),
		})
		if err != nil {
			writeErrorJSON(w, r, jsondata.ErrorJoinRoomFailed, "failed to join the room", err)
			return
		}

		writeRoomJSON(w, r, state)
	}
}

// LeaveAPIHandler は、以下のリクエストを処理するハンドラーです。
// POST /api/dashboard/rooms/{{roomID}}/leave
func LeaveAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		room := middleware.RoomFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		if user.DeviceID() == 0 {
			writeErrorJSON(w, r, jsondata.ErrorDeviceNotRegistered, "", nil)
			return
		}

		dashboardService := pb.NewDashboardServiceClient(conn)

		state, err := dashboardService.LeaveRoom(ctx, &pb.LeaveRoomRequest{
			TeamId:   user.TeamID(),
			RoomId:   room.GetRoomId(),
			DeviceId: user.DeviceID(),
		})
		if err != nil {
			writeErrorJSON(w, r, jsondata.ErrorLeaveRoomFailed, "failed to leave the room", err)
			return
		}

		writeRoomJSON(w, r, state)
	}
}

// PlayAPIHandler は、以下のリクエストを処理するハンドラーです。
// POST /api/dashboard/contents/{{contentID}}/play
func PlayAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		content := middleware.ContentFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		if user.DeviceID() == 0 {
			writeErrorJSON(w, r, jsondata.ErrorDeviceNotRegistered, "", nil)
			return
		}

		dashboardService := pb.NewDashboardServiceClient(conn)

		_, err := dashboardService.PlayContent(ctx, &pb.PlayContentRequest{
			TeamId:    user.TeamID(),
			DeviceId:  user.DeviceID(),
			ContentId: content.GetContentId(),
		})
		if grpc.Code(err) == codes.NotFound {
			writeErrorJSON(w, r, jsondata.ErrorContentNotFound, "the content is not found", err)
			return
		}
		if err != nil {
			writeErrorJSON(w, r, jsondata.ErrorUnknown, "failed to play the content", err)
			return
		}

		writeDashboardJSON(w, r)
	}
}

// StopAPIHandler は、以下のリクエストを処理するハンドラーです。
// POST /api/dashboard/stop
func StopAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		if user.DeviceID() == 0 {
			writeErrorJSON(w, r, jsondata.ErrorDeviceNotRegistered, "", nil)
			return
		}

		dashboardService := pb.NewDashboardServiceClient(conn)

		_, err := dashboardService.StopContent(ctx, &pb.StopContentRequest{
			TeamId:   user.TeamID(),
			DeviceId: user.DeviceID(),
		})
		if err != nil {
			writeErrorJSON(w, r, jsondata.ErrorUnknown, "failed to stop the content", err)
			return
		}

		writeDashboardJSON(w, r)
	}
}

func writeDashboardJSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.UserFromContext(ctx)
	conn := middleware.GRPCConnFromContext(ctx)

	dashboardService := pb.NewDashboardServiceClient(conn)

	dashboard, err := dashboardService.Get(ctx, &pb.DashboardRequest{
		TeamId: user.TeamID(),
	})
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			writeErrorJSON(w, r, jsondata.ErrorUnknown, "failed to get dashboard information", err)
			return
		}
	}

	data := jsondata.DashboardData{
		Rooms:    newRooms(dashboard, user.DeviceID()),
		Contents: newContents(dashboard),
	}

	webutil.WriteJSON(w, r, data)
}

func writeRoomJSON(w http.ResponseWriter, r *http.Request, state *pb.RoomState) {
	ctx := r.Context()
	user := auth.UserFromContext(ctx)
	conn := middleware.GRPCConnFromContext(ctx)

	roomService := pb.NewRoomServiceClient(conn)

	room, err := roomService.Get(ctx, &pb.RoomRequest{
		TeamId: user.TeamID(),
		RoomId: state.GetRoomId(),
	})
	if grpc.Code(err) != codes.NotFound {
		writeDashboardJSON(w, r)
		return
	}
	if err != nil {
		writeErrorJSON(w, r, jsondata.ErrorUnknown, "failed to get dashboard information", err)
	}

	data := jsondata.DashboardData{
		Room: newRoom(room, state),
	}
	data.Room.SetStatus(user.DeviceID())

	webutil.WriteJSON(w, r, data)
}

func newRooms(dashboard *pb.Dashboard, deviceID int64) (rooms []*jsondata.Room) {
	if dashboard == nil {
		return
	}

	states := map[int64]*pb.RoomState{}
	for _, state := range dashboard.GetRoomStates() {
		states[state.GetRoomId()] = state
	}

	rooms = make([]*jsondata.Room, len(dashboard.GetRooms()))
	for i, room := range dashboard.GetRooms() {
		rooms[i] = newRoom(room, states[room.GetRoomId()])
		rooms[i].SetStatus(deviceID)
	}

	return
}

func newRoom(room *pb.Room, state *pb.RoomState) *jsondata.Room {
	var opDeviceID int64
	if state != nil {
		for _, op := range state.GetOperators() {
			opDeviceID = op.GetDeviceId()
			break
		}
	}

	return &jsondata.Room{
		ID:               room.GetRoomId(),
		Name:             room.GetName(),
		OperatorDeviceID: opDeviceID,
	}
}

func newContents(dashboard *pb.Dashboard) (contents []*jsondata.Content) {
	if dashboard == nil {
		return
	}

	contents = make([]*jsondata.Content, len(dashboard.GetContents()))
	for i, content := range dashboard.GetContents() {
		contents[i] = &jsondata.Content{
			ID:   content.GetContentId(),
			Name: content.GetName(),
		}
	}

	return
}

func writeErrorJSON(w http.ResponseWriter, r *http.Request, code jsondata.ErrorCode, message string, err error) {
	if err != nil {
		logger := log.FromContext(r.Context())
		logger.Error(err)
	}

	if message == "" {
		message = code.String()
	}

	data := jsondata.DashboardData{
		Error: &jsondata.Error{
			Code:    code,
			Message: message,
		},
	}
	webutil.WriteJSON(w, r, data)
}
