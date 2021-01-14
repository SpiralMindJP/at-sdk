package content

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strconv"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CreatePageHandler は、以下のリクエストを処理するハンドラーです。
// GET /settings/contents/new
func CreatePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeCreatePage(w, r, nil, "")
	}
}

// CreateHandler は、以下のリクエストを処理するハンドラーです。
// POST /settings/contents/new
func CreateHandler(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		conn := middleware.GRPCConnFromContext(ctx)

		service := pb.NewContentServiceClient(conn)

		if !webutil.ParseMultipartForm(w, r) {
			return
		}

		name := r.PostForm.Get("name")
		if name == "" {
			writeCreatePage(w, r, &contentData{Name: name}, "コンテンツ名が入力されていません。")
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			webutil.WriteError(w, r, "failed to open upload file", err, http.StatusInternalServerError)
			return
		}
		defer file.Close()

		filename := header.Filename

		uploadURL, err := service.Upload(ctx, &pb.ContentUploadRequest{
			TeamId:   user.TeamID(),
			Name:     name,
			Type:     pb.ContentType_CONTENT_TYPE_VIDEO,
			FileName: filename,
		})
		if err != nil {
			webutil.WriteError(w, r, "failed to upload a new content", err, http.StatusInternalServerError)
			return
		}

		hash := md5.New()
		tr := io.TeeReader(file, hash)

		mediaType := mime.TypeByExtension(filepath.Ext(filename))
		if mediaType == "" {
			mediaType = "application/octet-stream"
		}

		req, err := http.NewRequest("PUT", uploadURL.GetUrl(), tr)
		if err != nil {
			err = fmt.Errorf("failed to post file [%s]: %w", filename, err)
			webutil.WriteError(w, r, "failed to upload a new content", err, http.StatusInternalServerError)
			return
		}

		req.Header.Set("Content-Type", mediaType)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			err = fmt.Errorf("failed to post file [%s]: %w", filename, err)
			webutil.WriteError(w, r, "failed to upload a new content", err, http.StatusInternalServerError)
			return
		}

		if res.StatusCode != http.StatusOK {
			err = fmt.Errorf("failed to post file [%s]: status %s", filename, res.Status)
			webutil.WriteError(w, r, "failed to upload a new content", err, http.StatusInternalServerError)
			return
		}

		md5Hash := hash.Sum(nil)

		_, err = service.FinishUpload(ctx, &pb.FinishUploadRequest{
			TeamId:    user.TeamID(),
			ContentId: uploadURL.GetContentId(),
			Md5:       md5Hash,
		})
		if grpc.Code(err) == codes.NotFound || grpc.Code(err) == codes.FailedPrecondition {
			writeCreatePage(w, r, &contentData{Name: name}, "コンテンツのアップロードに失敗しました。再度アップロードを試してください。")
		}
		if err != nil {
			err = fmt.Errorf("failed to finish upload content: %w", err)
			webutil.WriteError(w, r, "failed to upload a new content", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/contents
		redirectToListPage(w, s)
	}
}

func objectName(spaceID uint64, key, ext string) string {
	// (space id)/(random key)(.ext)
	return path.Join(strconv.FormatUint(spaceID, 10), key+ext)
}

func writeCreatePage(w http.ResponseWriter, r *http.Request, content *contentData, errMessage string) {
	// input not valid

	if content == nil {
		content = new(contentData)
	}

	data := template.NewData(r)
	data.Title = "コンテンツ登録"
	data.PageCommands = []*template.PageCommand{
		{Path: "/settings/contents", Name: "コンテンツ設定に戻る", Icon: "chevron_left"},
	}
	data.Data = content
	data.ErrMessage = errMessage

	template.WriteTemplate(w, r, "content", "create.tmpl", data)
}
