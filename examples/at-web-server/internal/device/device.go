// Package device は、Device 関連のリクエストを処理するハンドラーを提供します。
package device

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

type deviceData struct {
	ID     int64
	Name   string
	RoomID int64
}

func newDeviceData(device *pb.Device) *deviceData {
	return &deviceData{
		ID:     device.GetDeviceId(),
		Name:   device.GetName(),
		RoomID: device.GetRoomId(),
	}
}

type Server interface {
	BuildURL(path string) *url.URL
}

// ListPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/devices
func ListPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewDeviceServiceClient(conn)

		result, err := service.List(ctx, &pb.DeviceListRequest{
			TeamId: user.TeamID(),
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get rooms", err, http.StatusInternalServerError)
				return
			}
		}
		devices := make([]*deviceData, len(result.GetDevices()))
		for i, device := range result.GetDevices() {
			devices[i] = newDeviceData(device)
		}

		data := template.NewData(r)
		data.Title = "デバイス設定"
		data.PageCommands = []*template.PageCommand{
			{Path: "/settings/devices/new", Name: "デバイス登録", Icon: "add_circle"},
		}
		data.Data = devices

		template.WriteTemplate(w, r, "device", "list.tmpl", data)
	}
}

func redirectToListPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.Device)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
