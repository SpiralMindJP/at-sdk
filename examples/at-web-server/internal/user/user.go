// Package user は、ユーザー設定関連のリクエストを処理するハンドラーを提供します。
package user

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

type userData struct {
	Email      string
	Name       string
	DeviceName string
}

func newUserData(user auth.User, device *pb.Device) *userData {
	var deviceName string
	if device != nil {
		deviceName = device.GetName()
	}

	return &userData{
		Email:      user.Email(),
		Name:       user.Name(),
		DeviceName: deviceName,
	}
}

type Server interface {
	BuildURL(path string) *url.URL
}

// Handler は、以下のリクエストを処理するハンドラーです。
// GET /settings/user
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewDeviceServiceClient(conn)

		var device *pb.Device
		if user.DeviceID() > 0 {
			var err error
			device, err = service.Get(ctx, &pb.DeviceRequest{
				TeamId:   user.TeamID(),
				DeviceId: user.DeviceID(),
			})
			if err != nil {
				if grpc.Code(err) != codes.NotFound {
					webutil.WriteError(w, r, "failed to get the device", err, http.StatusInternalServerError)
					return
				}
				device = nil
			}
		}

		data := template.NewData(r)
		data.Title = "ユーザー設定"
		data.PageCommands = []*template.PageCommand{
			{Path: "/settings/user/edit", Name: "ユーザー情報編集", Icon: "edit"},
			{Path: "/settings/user/password", Name: "パスワード変更", Icon: "lock"},
			{Path: "/settings/user/device", Name: "デバイス登録", Icon: "smartphone"},
		}
		data.Data = newUserData(user, device)

		template.WriteTemplate(w, r, "user", "user.tmpl", data)
	}
}

func redirectToUserPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.User)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
