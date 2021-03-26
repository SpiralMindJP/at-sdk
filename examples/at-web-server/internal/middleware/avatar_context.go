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

type avatarCtxKey struct{}

func AvatarContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := GRPCConnFromContext(ctx)
		user := auth.UserFromContext(ctx)

		client := pb.NewAvatarServiceClient(conn)

		var avatar *pb.Avatar
		if avatarID, ok := webutil.URLParamInt64(r, "avatarID"); ok {
			var err error
			avatar, err = client.Get(ctx, &pb.AvatarRequest{
				TeamId:   user.TeamID(),
				AvatarId: avatarID,
			})
			if err != nil {
				if grpc.Code(err) != codes.NotFound {
					webutil.WriteError(w, r, "failed to get the avatar", err, http.StatusInternalServerError)
				}
				return
			}
		}
		if avatar == nil {
			http.NotFound(w, r)
			return
		}

		ctx = context.WithValue(ctx, avatarCtxKey{}, avatar)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AvatarFromRequest(r *http.Request) *pb.Avatar {
	avatar := AvatarFromContext(r.Context())
	if avatar == nil {
		panic("avatar is not found")
	}

	return avatar
}

func AvatarFromContext(ctx context.Context) *pb.Avatar {
	avatar, ok := ctx.Value(avatarCtxKey{}).(*pb.Avatar)
	if !ok {
		return nil
	}

	return avatar
}
