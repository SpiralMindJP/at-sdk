package middleware

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

type grpcConnCtxKey struct{}

// GRPCConnFromContext は context.Context にセットされた *grpc.ClientConn を取得します。
func GRPCConnFromContext(ctx context.Context) *grpc.ClientConn {
	conn, ok := ctx.Value(grpcConnCtxKey{}).(*grpc.ClientConn)
	if !ok {
		return nil
	}

	return conn
}

// NewContextWithGRPCConn は u を設定した parent のコピーを返します。
func NewContextWithGRPCConn(parent context.Context, conn *grpc.ClientConn) context.Context {
	return context.WithValue(parent, grpcConnCtxKey{}, conn)
}

// GRPCConnContext は、gRPC のクライアントコネクション conn を HTTP リクエストのコンテキストにセットするミドルウェアを返します。
func GRPCConnContext(conn *grpc.ClientConn) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := NewContextWithGRPCConn(r.Context(), conn)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
