package middleware

import (
	"context"
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type roomCtxKey struct{}

func RoomContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := GRPCConnFromContext(ctx)
		user := auth.UserFromContext(ctx)

		client := pb.NewRoomServiceClient(conn)

		var room *pb.Room
		if roomID, ok := webutil.URLParamInt64(r, "roomID"); ok {
			var err error
			room, err = client.Get(ctx, &pb.RoomRequest{
				TeamId: user.TeamID(),
				RoomId: roomID,
			})
			if err != nil {
				if grpc.Code(err) != codes.NotFound {
					webutil.WriteError(w, r, "failed to get the room", err, http.StatusInternalServerError)
				}
				return
			}
		}
		if room == nil {
			http.NotFound(w, r)
			return
		}

		ctx = context.WithValue(ctx, roomCtxKey{}, room)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoomFromRequest(r *http.Request) *pb.Room {
	room := RoomFromContext(r.Context())
	if room == nil {
		panic("room is not found")
	}

	return room
}

func RoomFromContext(ctx context.Context) *pb.Room {
	room, ok := ctx.Value(roomCtxKey{}).(*pb.Room)
	if !ok {
		return nil
	}

	return room
}
