package dashboard

import (
	"context"
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/dashboard/jsondata"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	readBufferSize  = 1 * 1024
	writeBufferSize = 1 * 1024
)

var upgrader = &websocket.Upgrader{ReadBufferSize: readBufferSize, WriteBufferSize: writeBufferSize}

// EventStreamHandler は、以下のリクエストを処理する WebSocket ハンドラーです。
// GET /api/dashboard/event
func EventStreamHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			webutil.WriteError(w, r, "failed to upgrade websocket", err, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		grpcConn := middleware.GRPCConnFromContext(ctx)

		dashboardService := pb.NewDashboardServiceClient(grpcConn)
		roomService := pb.NewRoomServiceClient(grpcConn)

		client, err := dashboardService.EventStream(ctx, &pb.DashboardRequest{
			TeamId: user.TeamID(),
		})
		if err != nil {
			sendErrorJSON(ctx, conn, jsondata.ErrorUnknown, "failed to get dashbard event stream", err)
			return
		}
		defer client.CloseSend()

		logger := log.FromContext(ctx).With("method", "dashboard.EventStreamHandler")
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			event, err := client.Recv()
			if err != nil {
				sendErrorJSON(ctx, conn, jsondata.ErrorUnknown, "failed to receive dashboard event", err)
				return
			}

			if event.GetType() != pb.DashboardEventType_DASHBOARD_EVENT_ROOM_EVENT {
				continue
			}

			roomEvent := event.GetRoomEvent()
			roomState := roomEvent.GetRoomState()
			if roomState == nil {
				continue
			}

			var typ jsondata.EventType
			switch roomEvent.GetType() {
			case pb.RoomEventType_ROOM_EVENT_ROOM_CREATED:
				typ = jsondata.RoomCreated
			case pb.RoomEventType_ROOM_EVENT_ROOM_DELETED:
				typ = jsondata.RoomDeleted
			case pb.RoomEventType_ROOM_EVENT_DEVICE_JOINED:
				typ = jsondata.DeviceJoined
			case pb.RoomEventType_ROOM_EVENT_DEVICE_LEAVED:
				typ = jsondata.DeviceLeaved
			case pb.RoomEventType_ROOM_EVENT_DEVICE_DELETED:
				typ = jsondata.DeviceDeleted
			case pb.RoomEventType_ROOM_EVENT_DEVICE_OFFLINE:
				typ = jsondata.DeviceOffline
			case pb.RoomEventType_ROOM_EVENT_DEVICE_ONLINE:
				typ = jsondata.DeviceOnline
			default:
				continue
			}

			room, err := roomService.Get(ctx, &pb.RoomRequest{
				TeamId: user.TeamID(),
				RoomId: roomState.GetRoomId(),
			})
			if grpc.Code(err) == codes.NotFound {
				continue
			}
			if err != nil {
				logger.Errorw("failed to get the room", "team_id", user.TeamID(), "room_id", roomState.GetRoomId())
				continue
			}

			e := &jsondata.Event{
				Type: typ,
				Room: newRoom(room, roomState),
			}
			e.Room.SetStatus(user.DeviceID())
			conn.WriteJSON(e)

		}
	}
}

func sendErrorJSON(ctx context.Context, conn *websocket.Conn, code jsondata.ErrorCode, message string, err error) {
	if err != nil {
		logger := log.FromContext(ctx)
		logger.Error(err)
	}

	if message == "" {
		message = code.String()
	}

	conn.WriteJSON(&jsondata.Error{
		Code:    code,
		Message: message,
	})
}
