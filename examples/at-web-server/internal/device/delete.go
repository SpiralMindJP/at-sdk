package device

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// DeletePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/devices/:id/delete
func DeletePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device := middleware.DeviceFromRequest(r)

		writeDeletePage(w, r, newDeviceData(device), "")
	}
}

// DeleteHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/devices/:id/delete
func DeleteHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		device := middleware.DeviceFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewDeviceServiceClient(conn)

		_, err := service.Delete(ctx, &pb.DeviceRequest{
			TeamId:   user.TeamID(),
			DeviceId: device.GetDeviceId(),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to delete the device", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/devices
		redirectToListPage(w, s)
	}
}

func writeDeletePage(w http.ResponseWriter, r *http.Request, device *deviceData, errMessage string) {
	// input not valid

	if device == nil {
		device = new(deviceData)
	}

	data := template.NewData(r)
	data.Title = "デバイス削除"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/devices", Name: "デバイス設定に戻る", Icon: "chevron_left"},
	}
	data.Data = device
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "device", "delete.tmpl", data)
}
