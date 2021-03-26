package device

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

// EditPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/devices/:id
func EditPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device := middleware.DeviceFromRequest(r)

		writeEditPage(w, r, newDeviceData(device, nil), "")
	}
}

// EditHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/devices/:id
func EditHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		device := middleware.DeviceFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewDeviceServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeEditPage(w, r, &deviceData{ID: device.GetDeviceId(), Name: name}, "デバイス名が入力されていません。")
			return
		}

		_, err := service.Update(ctx, &pb.DeviceUpdateRequest{
			TeamId:   user.TeamID(),
			DeviceId: device.GetDeviceId(),
			Name:     name,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to update the device", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/devices
		redirectToListPage(w, s)
	}
}

func writeEditPage(w http.ResponseWriter, r *http.Request, device *deviceData, errMessage string) {
	// input not valid

	if device == nil {
		device = new(deviceData)
	}

	data := template.NewData(r)
	data.Title = "デバイス編集"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/devices", Name: "デバイス設定に戻る", Icon: "chevron_left"},
	}
	data.Data = device
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "device", "edit.tmpl", data)
}
