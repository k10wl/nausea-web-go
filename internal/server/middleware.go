package server

import (
	"fmt"
	"html/template"
	"nausea-web/internal/templates"
	"net/http"
	"path/filepath"
)

func gzipMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		w.Header().Set("Content-Encoding", "gzip")
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".ttf":
			w.Header().Set("Content-Type", "font/ttf")
		case ".webp":
			w.Header().Set("Content-Type", "image/webp")
		default:
			w.Header().Set("Content-Type", "text/plain")
		}
		h.ServeHTTP(w, r)
	})
}

func notFoundMiddleware(t *template.Template, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.ExecuteTemplate(w, "/404", templates.TemplateData{})
			return
		}
		h.ServeHTTP(w, r)
	})
}

func routeLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v: %v\n", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
