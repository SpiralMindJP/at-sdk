package device

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

// CreatePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/devices/new
func CreatePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCreatePage(w, r, nil, "")
	}
}

// CreateHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/devices/new
func CreateHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewDeviceServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeCreatePage(w, r, &deviceData{Name: name}, "デバイス名が入力されていません。")
			return
		}

		otp, err := webutil.PostFormInt(r, "otp")
		if err != nil {
			message := "ワンタイムパスワードが正しくありません。"
			if err == webutil.ErrFormValueEmpty {
				message = "ワンタイムパスワードが入力されていません。"
			}

			writeCreatePage(w, r, &deviceData{Name: name}, message)
			return
		}

		_, err = service.Register(ctx, &pb.DeviceRegistrationRequest{
			TeamId: user.TeamID(),
			Name:   name,
			Otp:    int32(otp),
		})
		if grpc.Code(err) == codes.NotFound || grpc.Code(err) == codes.AlreadyExists {
			writeCreatePage(w, r, &deviceData{Name: name}, "ワンタイムパスワードが間違っているか期限切れです。")
			return
		}
		if err != nil {
			webutil.WriteError(w, r, "failed to create new device", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/devices
		redirectToListPage(w, s)
	}
}

func writeCreatePage(w http.ResponseWriter, r *http.Request, device *deviceData, errMessage string) {
	// input not valid

	if device == nil {
		device = new(deviceData)
	}

	data := template.NewData(r)
	data.Title = "デバイス登録"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/devices", Name: "デバイス設定に戻る", Icon: "chevron_left"},
	}
	data.Data = device
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "device", "create.tmpl", data)
}
