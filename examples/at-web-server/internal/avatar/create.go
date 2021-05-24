package avatar

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const maxAnimations = 10

type editAvatarData struct {
	ID                 int64
	Name               string
	AvatarContentID    int64
	AnimationContentID int64

	AvatarContents    []*contentData
	AnimationContents []*contentData

	Animations []*animationData
}

type animationData struct {
	Name string
}

func newEditAvatarData(avatar *pb.Avatar) *editAvatarData {
	animations := make([]*animationData, maxAnimations)
	for i := 0; i < maxAnimations; i++ {
		var animation *animationData

		if i < len(avatar.GetAnimations()) {
			animation = newAnimationData(avatar.GetAnimations()[i])
		}

		if animation == nil {
			animation = &animationData{}
		}

		animations[i] = animation
	}

	return &editAvatarData{
		ID:                 avatar.GetAvatarId(),
		Name:               avatar.GetName(),
		AvatarContentID:    avatar.GetAvatarContentId(),
		AnimationContentID: avatar.GetAnimationContentId(),
		Animations:         animations,
	}
}

func newAnimationData(animation *pb.Animation) *animationData {
	return &animationData{
		Name: animation.GetName(),
	}
}

func newPBAnimation(animation *animationData) *pb.Animation {
	return &pb.Animation{
		Name: animation.Name,
	}
}

func newPBAnimations(animations []*animationData) []*pb.Animation {
	var list []*pb.Animation
	for _, animation := range animations {
		list = append(list, newPBAnimation(animation))
	}
	return list
}

// CreatePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/avatars/new
func CreatePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCreatePage(w, r, newEditAvatarData(nil), "")
	}
}

// CreateHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/avatars/new
func CreateHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewAvatarServiceClient(conn)

		if !webutil.ParseForm(w, r) {
			return
		}

		avatar, errMessages := verifyPost(r)
		if len(errMessages) > 0 {
			writeCreatePage(w, r, avatar, strings.Join(errMessages, "<br />"))
			return
		}

		_, err := service.Create(ctx, &pb.AvatarCreateRequest{
			TeamId:             user.TeamID(),
			Name:               avatar.Name,
			AvatarContentId:    avatar.AvatarContentID,
			AnimationContentId: avatar.AnimationContentID,
			Animations:         newPBAnimations(avatar.Animations),
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to create new avatar", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/avatars
		redirectToListPage(w, s)
	}
}

func verifyPost(r *http.Request) (*editAvatarData, []string) {
	var errMessages []string

	name := r.PostForm.Get("name")
	if name == "" {
		errMessages = append(errMessages, "アバター名が入力されていません。")
	}

	avatarContentID, err := webutil.PostFormInt64(r, "avatar_content")
	if err != nil || avatarContentID <= 0 {
		errMessages = append(errMessages, "アバターコンテンツが選択されていません。")
	}

	animationContentID, err := webutil.PostFormInt64(r, "animation_content")
	if err != nil || animationContentID < 0 {
		errMessages = append(errMessages, "アニメーションコンテンツが選択されていません。")
	}

	animations := getAnimations(r)

	return &editAvatarData{
		Name:               name,
		AvatarContentID:    avatarContentID,
		AnimationContentID: animationContentID,
		Animations:         animations,
	}, errMessages
}

func getAnimations(r *http.Request) []*animationData {
	animations := make([]*animationData, maxAnimations)
	for i := 0; i < maxAnimations; i++ {
		name := r.PostForm.Get(fmt.Sprintf("animation%d", i))
		animations[i] = &animationData{
			Name: name,
		}
	}

	return animations
}

func writeCreatePage(w http.ResponseWriter, r *http.Request, avatar *editAvatarData, errMessage string) {
	ctx := r.Context()

	avatarContents, animationContents, err := getContents(ctx)
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			webutil.WriteError(w, r, "failed to get contents", err, http.StatusInternalServerError)
			return
		}
	}
	avatar.AvatarContents = avatarContents
	avatar.AnimationContents = animationContents

	data := template.NewData(r)
	data.Title = "アバター登録"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/avatars", Name: "アバター設定に戻る", Icon: "chevron_left"},
	}
	data.Data = avatar
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "avatar", "create.tmpl", data)
}

func getContents(ctx context.Context) ([]*contentData, []*contentData, error) {
	user := auth.UserFromContext(ctx)
	conn := middleware.GRPCConnFromContext(ctx)

	service := pb.NewContentServiceClient(conn)

	result, err := service.List(ctx, &pb.ContentListRequest{
		TeamId: user.TeamID(),
	})
	if err != nil {
		if grpc.Code(err) != codes.NotFound {
			return nil, nil, err
		}
	}

	var avatarContents []*contentData
	var animationContents []*contentData
	for _, content := range result.GetContents() {
		switch content.GetType() {
		case pb.ContentType_CONTENT_TYPE_AVATAR:
			avatarContents = append(avatarContents, &contentData{
				ID:   content.GetContentId(),
				Name: content.GetName(),
			})
		case pb.ContentType_CONTENT_TYPE_ANIMATION:
			animationContents = append(animationContents, &contentData{
				ID:   content.GetContentId(),
				Name: content.GetName(),
			})
		}
	}

	return avatarContents, animationContents, nil
}
