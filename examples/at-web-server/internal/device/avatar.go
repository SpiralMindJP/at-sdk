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

type setAvatarPageData struct {
	DeviceID int64
	AvatarID int64
	Avatars  []*avatarData
}

// SetAvatarPageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/devices/:id/avatar
func SetAvatarPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device := middleware.DeviceFromRequest(r)
		if device.GetType() != pb.DeviceType_DEVICE_TYPE_OPERATOR {
			http.NotFound(w, r)
			return
		}

		writeSetAvatarPage(w, r, "")
	}
}

// SetAvatarHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/devices/:id/avatar
func SetAvatarHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		device := middleware.DeviceFromRequest(r)
		conn := middleware.GRPCConnFromContext(ctx)

		if device.GetType() != pb.DeviceType_DEVICE_TYPE_OPERATOR {
			http.NotFound(w, r)
			return
		}

		teamID := user.TeamID()
		deviceID := device.GetDeviceId()

		service := pb.NewAvatarServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		avatarID, err := webutil.PostFormInt64(r, "avatar")
		if err != nil {
			webutil.WriteError(w, r, "invalid avatar id is received", err, http.StatusBadRequest)
			return
		}

		if avatarID == 0 {
			_, err := service.DeleteOperatorAvatar(ctx, &pb.OperatorAvatarRequest{
				TeamId:   teamID,
				DeviceId: deviceID,
			})
			if err != nil && grpc.Code(err) != codes.NotFound {
				webutil.WriteError(w, r, "failed to delete the avatar from the operator device", err, http.StatusInternalServerError)
				return
			}
		} else {
			_, err := service.SetOperatorAvatar(ctx, &pb.OperatorAvatarSetRequest{
				TeamId:   teamID,
				DeviceId: deviceID,
				AvatarId: avatarID,
			})
			if grpc.Code(err) == codes.NotFound {
				writeSetAvatarPage(w, r, "デバイスまたは選択されたアバターが見つかりません。再度選択してください。")
				return
			}
			if err != nil {
				webutil.WriteError(w, r, "failed to set the avatar to the operator device", err, http.StatusInternalServerError)
				return
			}
		}

		// redirect to /settings/devices
		redirectToListPage(w, s)
	}
}

func writeSetAvatarPage(w http.ResponseWriter, r *http.Request, errMessage string) {
	ctx := r.Context()
	user := auth.UserFromContext(ctx)
	device := middleware.DeviceFromRequest(r)
	conn := middleware.GRPCConnFromContext(ctx)

	teamID := user.TeamID()
	deviceID := device.GetDeviceId()

	service := pb.NewAvatarServiceClient(conn)

	opAvatar, err := service.GetOperatorAvatar(ctx, &pb.OperatorAvatarRequest{
		TeamId:   teamID,
		DeviceId: deviceID,
	})
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			webutil.WriteError(w, r, "failed to get operator avatar", err, http.StatusInternalServerError)
			return
		}
	}
	var avatarID int64
	if opAvatar != nil {
		avatarID = opAvatar.GetAvatarId()
	}

	result, err := service.List(ctx, &pb.AvatarListRequest{
		TeamId: user.TeamID(),
	})
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			webutil.WriteError(w, r, "failed to get avatars", err, http.StatusInternalServerError)
			return
		}
	}

	avatars := make([]*avatarData, len(result.GetAvatars()))
	for i, avatar := range result.GetAvatars() {
		avatars[i] = &avatarData{
			ID:   avatar.GetAvatarId(),
			Name: avatar.GetName(),
		}
	}

	data := template.NewData(r)
	data.Title = "オペレーターアバター設定"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/devices", Name: "デバイス設定に戻る", Icon: "list_alt"},
	}
	data.Data = setAvatarPageData{
		DeviceID: deviceID,
		AvatarID: avatarID,
		Avatars:  avatars,
	}
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "device", "avatar.tmpl", data)
}
