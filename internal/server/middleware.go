package server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"nausea-web/internal/db"
	"nausea-web/internal/templates"
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

func notFoundMiddleware(t *template.Template, db *db.DB, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			meta, err := db.GetMeta(ctx)
			if err != nil {
				t.ExecuteTemplate(w, "/404", templates.TemplateData{})
				return
			}
			t.ExecuteTemplate(w, "/404", templates.TemplateData{
				"Title":    "Nausea",
				"HomePage": true,
				"Meta":     meta,
			})
			return
		}
		h.ServeHTTP(w, r)
	})
}

func routeLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := ""
		if r.URL.RawQuery != "" {
			params = "?" + r.URL.RawQuery
		}
		log.Printf("%v: %v%v\n", r.Method, r.URL.Path, params)
		h.ServeHTTP(w, r)
	})
}
