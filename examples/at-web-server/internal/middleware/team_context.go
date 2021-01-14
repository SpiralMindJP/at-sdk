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

type teamCtxKey struct{}

func TeamContext() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			conn := GRPCConnFromContext(ctx)

			user := auth.UserFromContext(ctx)
			if !user.Admin() {
				http.NotFound(w, r)
				return
			}

			client := pb.NewTeamServiceClient(conn)

			var team *pb.Team

			if teamID, ok := webutil.URLParamInt64(r, "teamID"); ok {
				var err error
				team, err = client.Get(ctx, &pb.TeamRequest{
					TeamId: teamID,
				})
				if err != nil {
					if grpc.Code(err) != codes.NotFound {
						webutil.WriteError(w, r, "failed to get the team", err, http.StatusInternalServerError)
					}
					return
				}
			}
			if team == nil {
				http.NotFound(w, r)
				return
			}

			ctx = context.WithValue(ctx, teamCtxKey{}, team)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func TeamFromContext(ctx context.Context) *pb.Team {
	team, ok := ctx.Value(teamCtxKey{}).(*pb.Team)
	if !ok {
		return nil
	}

	return team
}
