package content

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/mime"
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
		err := HandleUploadContent(w, r, pb.ContentType_CONTENT_TYPE_IMAGE, pb.ContentType_CONTENT_TYPE_VIDEO, pb.ContentType_CONTENT_TYPE_AVATAR, pb.ContentType_CONTENT_TYPE_ANIMATION)
		if err != nil {
			var upErr *UploadError
			if errors.As(err, &upErr) {
				switch upErr.Code {
				case InvalidForm:
				case InvalidName:
					writeCreatePage(w, r, nil, "コンテンツ名が入力されていません。")
				case InvalidType:
					writeCreatePage(w, r, nil, "コンテンツタイプが選択されていません。")
				case InvalidFile:
					webutil.WriteError(w, r, upErr.Msg, err, http.StatusInternalServerError)
				case FailedUpload:
					webutil.WriteError(w, r, upErr.Msg, err, http.StatusInternalServerError)
				default:
					webutil.WriteError(w, r, upErr.Msg, err, http.StatusInternalServerError)
				}
				return
			}

			webutil.WriteError(w, r, "failed to upload a new content", err, http.StatusInternalServerError)
			return
		}

		// redirect to /settings/contents
		redirectToListPage(w, s)
	}
}

type ErrorCode int

const (
	Unknown ErrorCode = iota
	InvalidForm
	InvalidName
	InvalidFile
	InvalidType
	FailedUpload
)

type UploadError struct {
	Code ErrorCode
	Msg  string
	Err  error
}

func (err *UploadError) Error() string {
	if err == nil {
		return ""
	}

	var s strings.Builder
	if err.Msg == "" {
		s.WriteString("failed to upload")
	} else {
		s.WriteString(err.Msg)
	}

	switch err.Code {
	case InvalidForm:
		s.WriteString(": invalid form")
	case InvalidName:
		s.WriteString(": invalid name")
	case InvalidFile:
		s.WriteString(": invalid file")
	case InvalidType:
		s.WriteString(": invalid type")
	case FailedUpload:
		s.WriteString(": failed to upload")
	}

	if err.Err != nil {
		s.WriteString(": ")
		s.WriteString(err.Err.Error())
	}

	return s.String()
}

func (err *UploadError) Unwrap() error {
	if err == nil {
		return nil
	}

	return err.Err
}

func HandleUploadContent(w http.ResponseWriter, r *http.Request, acceptedTypes ...pb.ContentType) error {
	ctx := r.Context()
	user := auth.UserFromContext(ctx)
	conn := middleware.GRPCConnFromContext(ctx)

	service := pb.NewContentServiceClient(conn)

	if !webutil.ParseMultipartForm(w, r) {
		// Bad Request
		return &UploadError{Code: InvalidForm}
	}

	name := r.PostForm.Get("name")
	if name == "" {
		return &UploadError{Code: InvalidName}
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return &UploadError{Code: InvalidFile, Msg: "failed to open upload file", Err: err}
	}
	defer file.Close()

	ctype, err := webutil.PostFormInt(r, "type")
	if err != nil {
		return &UploadError{Code: InvalidType, Err: err}
	} else {
		validType := false
		ctype := pb.ContentType(ctype)
		for _, tp := range acceptedTypes {
			if ctype == tp {
				validType = true
				break
			}
		}
		if !validType {
			return &UploadError{Code: InvalidType}
		}
	}

	filename := header.Filename

	uploadURL, err := service.Upload(ctx, &pb.ContentUploadRequest{
		TeamId:   user.TeamID(),
		Name:     name,
		Type:     pb.ContentType(ctype),
		FileName: filename,
	})
	if err != nil {
		return &UploadError{Code: Unknown, Msg: "failed to upload a new content", Err: err}
	}

	hash := md5.New()
	tr := io.TeeReader(file, hash)

	req, err := http.NewRequest("PUT", uploadURL.GetUrl(), tr)
	if err != nil {
		err = fmt.Errorf("failed to post file [%s]: %w", filename, err)
		return &UploadError{Code: Unknown, Msg: "failed to upload a new content", Err: err}
	}

	req.Header.Set("Content-Type", mime.TypeByExtension(filepath.Ext(filename)))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to post file [%s]: %w", filename, err)
		return &UploadError{Code: Unknown, Msg: "failed to upload a new content", Err: err}
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to post file [%s]: status %s", filename, res.Status)
		return &UploadError{Code: Unknown, Msg: "failed to upload a new content", Err: err}
	}

	md5Hash := hash.Sum(nil)

	_, err = service.FinishUpload(ctx, &pb.FinishUploadRequest{
		TeamId:    user.TeamID(),
		ContentId: uploadURL.GetContentId(),
		Md5:       md5Hash,
	})
	if grpc.Code(err) == codes.NotFound || grpc.Code(err) == codes.FailedPrecondition {
		return &UploadError{Code: FailedUpload, Err: err}
	}
	if err != nil {
		return &UploadError{Code: Unknown, Msg: "failed to finish upload content", Err: err}
	}

	return nil
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
