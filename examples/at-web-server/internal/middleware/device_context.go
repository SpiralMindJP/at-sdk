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

type deviceContextKey struct{}

func DeviceContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := GRPCConnFromContext(ctx)
		user := auth.UserFromContext(ctx)

		client := pb.NewDeviceServiceClient(conn)

		var device *pb.Device
		if deviceID, ok := webutil.URLParamInt64(r, "deviceID"); ok {
			var err error
			device, err = client.Get(ctx, &pb.DeviceRequest{
				TeamId:   user.TeamID(),
				DeviceId: deviceID,
			})
			if err != nil {
				if grpc.Code(err) != codes.NotFound {
					webutil.WriteError(w, r, "failed to get the device", err, http.StatusInternalServerError)
				}
				return
			}
		}
		if device == nil {
			http.NotFound(w, r)
			return
		}

		ctx = context.WithValue(ctx, deviceContextKey{}, device)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func DeviceFromRequest(r *http.Request) *pb.Device {
	device := DeviceFromContext(r.Context())
	if device == nil {
		panic("device is not found")
	}

	return device
}

func DeviceFromContext(ctx context.Context) *pb.Device {
	device, ok := ctx.Value(deviceContextKey{}).(*pb.Device)
	if !ok {
		return nil
	}

	return device
}
