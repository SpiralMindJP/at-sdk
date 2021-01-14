package user

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	webauth "github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type deviceData struct {
	ID   int64
	Name string
}

func newDeviceData(device *pb.Device) *deviceData {
	return &deviceData{
		ID:   device.GetDeviceId(),
		Name: device.GetName(),
	}
}

// SetDevicePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/user/device
func SetDevicePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeSetDevicePage(w, r, nil, "")
	}
}

// SetDeviceHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/user/device
func SetDeviceHandler(client *auth.Client, s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := webauth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewDeviceServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeSetDevicePage(w, r, &deviceData{Name: name}, "デバイス名が入力されていません。")
			return
		}

		otp, err := webutil.PostFormInt(r, "otp")
		if err != nil {
			message := "ワンタイムパスワードが正しくありません。"
			if err == webutil.ErrFormValueEmpty {
				message = "ワンタイムパスワードが入力されていません。"
			}

			writeSetDevicePage(w, r, &deviceData{Name: name}, message)
			return
		}

		device, err := service.Register(ctx, &pb.DeviceRegistrationRequest{
			TeamId: user.TeamID(),
			Name:   name,
			Otp:    int32(otp),
		})
		if grpc.Code(err) == codes.NotFound || grpc.Code(err) == codes.AlreadyExists {
			writeSetDevicePage(w, r, &deviceData{Name: name}, "ワンタイムパスワードが間違っているか期限切れです。")
			return
		}
		if err != nil {
			webutil.WriteError(w, r, "failed to set a new device to the user", err, http.StatusInternalServerError)
			return
		}

		claims := map[string]interface{}{
			"team_id":   user.TeamID(),
			"admin":     user.Admin(),
			"device_id": device.GetDeviceId(),
		}
		if err := client.SetCustomUserClaims(ctx, user.UID(), claims); err != nil {
			webutil.WriteError(w, r, "failed to set a new device to the user", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/devices
		redirectToUserPage(w, s)
	}
}

func writeSetDevicePage(w http.ResponseWriter, r *http.Request, device *deviceData, errMessage string) {
	// input not valid

	if device == nil {
		device = new(deviceData)
	}

	data := template.NewData(r)
	data.Title = "使用デバイス登録"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/user", Name: "ユーザー設定に戻る", Icon: "person"},
	}
	data.Data = device
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "user", "device.tmpl", data)
}
