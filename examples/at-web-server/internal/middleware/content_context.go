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

type contentContextKey struct{}

func ContentContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := GRPCConnFromContext(ctx)
		user := auth.UserFromContext(ctx)

		client := pb.NewContentServiceClient(conn)

		var content *pb.Content
		if contentID, ok := webutil.URLParamInt64(r, "contentID"); ok {
			var err error
			content, err = client.Get(ctx, &pb.ContentRequest{
				TeamId:    user.TeamID(),
				ContentId: contentID,
			})
			if err != nil {
				if grpc.Code(err) != codes.NotFound {
					webutil.WriteError(w, r, "failed to get the content", err, http.StatusInternalServerError)
				}
				return
			}
		}
		if content == nil {
			http.NotFound(w, r)
			return
		}

		ctx = context.WithValue(ctx, contentContextKey{}, content)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ContentFromRequest(r *http.Request) *pb.Content {
	content := ContentFromContext(r.Context())
	if content == nil {
		panic("content is not found")
	}

	return content
}

func ContentFromContext(ctx context.Context) *pb.Content {
	content, ok := ctx.Value(contentContextKey{}).(*pb.Content)
	if !ok {
		return nil
	}

	return content
}
