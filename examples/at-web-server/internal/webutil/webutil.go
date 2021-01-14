package webutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strconv"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/go-chi/chi"
)

func URLParamInt64(r *http.Request, key string) (value int64, ok bool) {
	var param string
	if param = chi.URLParam(r, key); param == "" {
		return
	}

	n, err := strconv.ParseUint(param, 10, 64)
	if err == nil {
		ok = true
	}
	value = int64(n)

	return
}

func ParseForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		const msg = "failed to parse HTTP request form"
		err = fmt.Errorf("%s: %w", msg, err)
		WriteWarn(w, r, msg, err, http.StatusBadRequest)
		return false
	}

	return true
}

func ParseMultipartForm(w http.ResponseWriter, r *http.Request) bool {
	// 32MB data are stored in memory
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		const msg = "failed to parse HTTP request multipart form"
		err = fmt.Errorf("%s: %w", msg, err)
		WriteWarn(w, r, msg, err, http.StatusBadRequest)
		return false
	}

	return true
}

var (
	ErrFormValueEmpty   = errors.New("form value is empty")
	ErrInvalidFormValue = errors.New("invalid form value")
)

func PostFormInt(r *http.Request, key string) (int, error) {
	s := r.PostForm.Get(key)
	if s == "" {
		return 0, ErrFormValueEmpty
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, ErrInvalidFormValue
	}

	return v, nil
}

func PostFormInt64(r *http.Request, key string) (int64, error) {
	s := r.PostForm.Get(key)
	if s == "" {
		return 0, ErrFormValueEmpty
	}

	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, ErrInvalidFormValue
	}

	return int64(v), nil
}

func ValidContentType(t string, types ...string) bool {
	mt, _, err := mime.ParseMediaType(t)
	if err != nil {
		return false
	}

	for i := range types {
		if mt == types[i] {
			return true
		}
	}

	return false
}

func DecodeBody(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	ctype := r.Header.Get("Content-Type")
	if !ValidContentType(ctype, "text/json", "application/json") {
		WriteWarn(w, r, fmt.Sprintf("invalid Content-Type: %s", ctype), nil, http.StatusUnsupportedMediaType)
		return false
	}

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		WriteWarn(w, r, "failed to decode request JSON", nil, http.StatusBadRequest)
		return false
	}

	return true
}

func WriteJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		const msg = "failed to encode response JSON"
		err = fmt.Errorf("%s: %w", msg, err)
		WriteError(w, r, msg, err, http.StatusInternalServerError)
		return false
	}

	w.Header().Add("Content-Type", "text/json")
	if _, err := io.Copy(w, &buf); err != nil {
		logger := log.FromContext(r.Context())
		logger.Errorw("failed to write HTTP response", "err", err)
		return false
	}

	return true
}

func WriteError(w http.ResponseWriter, r *http.Request, msg string, err error, code int) {
	if msg == "" {
		if err != nil {
			msg = err.Error()
		} else {
			msg = http.StatusText(code)
		}
	}

	if err != nil {
		logger := log.FromContext(r.Context())
		logger.Errorw(msg, "status_code", code, "path", r.URL.Path, "err", err)
	}
	http.Error(w, msg, code)
}

func WriteWarn(w http.ResponseWriter, r *http.Request, msg string, err error, code int) {
	if msg == "" {
		if err != nil {
			msg = err.Error()
		} else {
			msg = http.StatusText(code)
		}
	}

	if err != nil {
		logger := log.FromContext(r.Context())
		logger.Warnw(msg, "status_code", code, "path", r.URL.Path, "err", err)
	}
	http.Error(w, msg, code)
}
