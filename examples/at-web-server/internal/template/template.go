package template

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"path"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
)

type PageCommand struct {
	Path string
	Name string
	Icon string
}

type Data struct {
	Title        string
	LoginUser    auth.User
	PageCommands []*PageCommand
	Data         interface{}
	ErrMessage   string
}

func NewData(r *http.Request) *Data {
	data := new(Data)

	user := auth.UserFromContext(r.Context())
	if user != nil {
		data.LoginUser = user
	}

	return data
}

func LoadTemplate(tmplFS fs.FS) {
	loadTemplate(tmplFS, "login")
	loadTemplate(tmplFS, "dashboard")
	loadTemplate(tmplFS, "room")
	loadTemplate(tmplFS, "device")
	loadTemplate(tmplFS, "content")
	loadTemplate(tmplFS, "user")
	loadTemplate(tmplFS, "admin/team")
}

var templates = map[string]*template.Template{}

func loadTemplate(tmplFS fs.FS, dir string) {
	t := template.New(dir)
	t.Funcs(template.FuncMap{
		"safeurl":  func(u string) template.URL { return template.URL(u) },
		"safehtml": func(h string) template.HTML { return template.HTML(h) },
	})

	// 共通テンプレートを先に読み込む
	template.Must(t.ParseFS(tmplFS, "templates/common/*.tmpl"))

	template.Must(t.ParseFS(tmplFS, path.Join("templates", dir, "*.tmpl")))

	templates[dir] = t
}

func WriteTemplate(w http.ResponseWriter, r *http.Request, dir, name string, data interface{}) {
	w.Header().Add("Content-Type", "text/html")

	t, ok := templates[dir]
	if !ok {
		err := fmt.Errorf("template is not found: %s", dir)
		webutil.WriteError(w, r, http.StatusText(http.StatusInternalServerError), err, http.StatusInternalServerError)
	}

	var buf bytes.Buffer
	err := t.ExecuteTemplate(&buf, name, data)
	if err != nil {
		err = fmt.Errorf("execute template error: %w", err)
		webutil.WriteError(w, r, http.StatusText(http.StatusInternalServerError), err, http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, &buf)
	if err != nil {
		logger := log.FromContext(r.Context())
		logger.Errorw("failed to write HTTP response", "path", r.URL.Path, "err", err)
	}
}
