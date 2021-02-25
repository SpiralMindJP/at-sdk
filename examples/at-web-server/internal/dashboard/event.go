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

			switch event.GetType() {
			case pb.DashboardEventType_JOIN_ROOM, pb.DashboardEventType_LEAVE_ROOM:
				state := event.GetRoomState()
				if state == nil {
					continue
				}

				room, err := roomService.Get(ctx, &pb.RoomRequest{
					TeamId: user.TeamID(),
					RoomId: state.GetRoomId(),
				})
				if grpc.Code(err) == codes.NotFound {
					continue
				}
				if err != nil {
					logger.Errorw("failed to get the room", "team_id", user.TeamID(), "room_id", state.GetRoomId())
					continue
				}

				e := &jsondata.Event{
					Type: jsondata.EventType(event.GetType()),
					Room: newRoom(room, state),
				}
				e.Room.SetStatus(user.DeviceID())
				conn.WriteJSON(e)
			default:
				continue
			}

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
