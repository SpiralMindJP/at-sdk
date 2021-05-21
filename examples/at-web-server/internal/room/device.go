package room

import (
	"net/http"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type setDevicePageData struct {
	ID      int64
	Devices []*deviceData
}

// SetDevicePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/rooms/:id/devices
func SetDevicePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeSetDevicePage(w, r, "")
	}
}

// SetDeviceHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/room/device
func SetDeviceHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		room := middleware.RoomFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewRoomServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		deviceID, err := webutil.PostFormInt64(r, "device")
		if err != nil {
			webutil.WriteError(w, r, "invalid device id is received", err, http.StatusBadRequest)
			return
		}

		if deviceID > 0 {
			_, err = service.SetDevice(ctx, &pb.RoomDeviceRequest{
				TeamId:   user.TeamID(),
				RoomId:   room.GetRoomId(),
				DeviceId: deviceID,
				Force:    true,
			})
		} else {
			_, err = service.DeleteDevice(ctx, &pb.RoomRequest{
				TeamId: user.TeamID(),
				RoomId: room.GetRoomId(),
			})
		}
		if grpc.Code(err) == codes.NotFound {
			writeSetDevicePage(w, r, "ルームまたは選択されたデバイスが見つかりません。再度選択してください。")
			return
		}
		if err != nil {
			webutil.WriteError(w, r, "failed to set the device to the room", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/rooms
		redirectToListPage(w, s)
	}
}

func writeSetDevicePage(w http.ResponseWriter, r *http.Request, errMessage string) {
	ctx := r.Context()
	user := auth.UserFromContext(ctx)
	room := middleware.RoomFromRequest(r)
	conn := middleware.GRPCConnFromContext(ctx)

	service := pb.NewDeviceServiceClient(conn)

	result, err := service.ListByType(ctx, &pb.DeviceListByTypeRequest{
		TeamId: user.TeamID(),
		Type:   pb.DeviceType_DEVICE_TYPE_CUSTOMER,
	})
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			webutil.WriteError(w, r, "failed to get devices", err, http.StatusInternalServerError)
			return
		}
	}

	devices := make([]*deviceData, len(result.GetDevices()))
	for i, device := range result.GetDevices() {
		devices[i] = &deviceData{
			ID:     device.GetDeviceId(),
			Name:   device.GetName(),
			RoomID: device.GetRoomId(),
		}
	}

	data := template.NewData(r)
	data.Title = "ルームデバイス設定"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/rooms", Name: "ルーム設定に戻る", Icon: "list_alt"},
	}
	data.Data = setDevicePageData{
		ID:      room.GetRoomId(),
		Devices: devices,
	}
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "room", "device.tmpl", data)
}
