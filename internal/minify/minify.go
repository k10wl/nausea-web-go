package minify

import (
	"bytes"
	"io"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

type Minifier interface {
	Minify(w io.Writer, r io.Reader, name string) error
	CanMinify(name string) bool
}

type TdewolffMinifier struct {
	m *minify.M
}

var ExtToMime = map[string]string{
	".js":  "application/javascript; charset=utf-8",
	".css": "text/css; charset=utf-8",
}

func NewTdewolffMinifier() *TdewolffMinifier {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)
	return &TdewolffMinifier{m: m}
}

func (m *TdewolffMinifier) Minify(w io.Writer, r io.Reader, name string) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	t := detectMime(data, name)
	return m.m.Minify(t, w, bytes.NewReader(data))
}

func (m *TdewolffMinifier) CanMinify(name string) bool {
	ext := filepath.Ext(name)
	_, ok := ExtToMime[ext]
	return ok
}

func detectMime(data []byte, name string) string {
	mimeType := http.DetectContentType(data)
	if mimeType == "text/plain; charset=utf-8" {
		ext := filepath.Ext(name)
		if val, ok := ExtToMime[ext]; ok {
			return val
		}
	}
	return mimeType
}
