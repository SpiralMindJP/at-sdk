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

type deviceListData struct {
	Operators []*deviceData
	Customers []*deviceData
}

type deviceData struct {
	ID     int64
	Name   string
	RoomID int64

	Avatar *avatarData
}

type avatarData struct {
	ID   int64
	Name string
}

func newDeviceData(device *pb.Device, avatar *pb.Avatar) *deviceData {
	var a *avatarData
	if avatar != nil {
		a = newAvatarData(avatar)
	}

	return &deviceData{
		ID:     device.GetDeviceId(),
		Name:   device.GetName(),
		RoomID: device.GetRoomId(),
		Avatar: a,
	}
}

func newAvatarData(avatar *pb.Avatar) *avatarData {
	return &avatarData{
		ID:   avatar.GetAvatarId(),
		Name: avatar.GetName(),
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

		teamID := user.TeamID()

		service := pb.NewDeviceServiceClient(conn)
		avatarService := pb.NewAvatarServiceClient(conn)

		result, err := service.List(ctx, &pb.DeviceListRequest{
			TeamId: teamID,
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get rooms", err, http.StatusInternalServerError)
				return
			}
		}

		avatars, err := avatarService.List(ctx, &pb.AvatarListRequest{
			TeamId: teamID,
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get avatars", err, http.StatusInternalServerError)
				return
			}
		}

		opAvatars, err := avatarService.ListOperatorAvatar(ctx, &pb.OperatorAvatarListRequest{
			TeamId: teamID,
		})
		if err != nil {
			if grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to get operator avatars", err, http.StatusInternalServerError)
				return
			}
		}

		var devAvatarMap map[int64]*pb.Avatar
		if len(opAvatars.GetOperatorAvatars()) > 0 {
			avatarMap := map[int64]*pb.Avatar{}
			for _, avatar := range avatars.GetAvatars() {
				avatarMap[avatar.GetAvatarId()] = avatar
			}

			devAvatarMap = map[int64]*pb.Avatar{}
			for _, opAvatar := range opAvatars.GetOperatorAvatars() {
				avatar, ok := avatarMap[opAvatar.GetAvatarId()]
				if !ok {
					continue
				}
				devAvatarMap[opAvatar.GetDeviceId()] = avatar
			}
		}

		var list deviceListData
		for _, device := range result.GetDevices() {
			switch device.GetType() {
			case pb.DeviceType_DEVICE_TYPE_OPERATOR:
				avatar := devAvatarMap[device.GetDeviceId()]
				list.Operators = append(list.Operators, newDeviceData(device, avatar))
			case pb.DeviceType_DEVICE_TYPE_CUSTOMER:
				list.Customers = append(list.Customers, newDeviceData(device, nil))
			}
		}

		data := template.NewData(r)
		data.Title = "デバイス設定"
		data.PageCommands = []*template.PageCommand{
			{Path: "/settings/devices/new", Name: "デバイス登録", Icon: "add_circle"},
		}
		data.Data = list

		template.WriteTemplate(w, r, "device", "list.tmpl", data)
	}
}

func redirectToListPage(w http.ResponseWriter, s Server) {
	location := s.BuildURL(path.Device)

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}
