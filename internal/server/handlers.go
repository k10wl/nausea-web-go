package server

import (
	"context"
	"html/template"
	"nausea-web/internal/db"
	"nausea-web/internal/templates"
	"net/http"
	"time"
)

func handleHome(t *template.Template, db *db.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			meta, err := db.GetMeta(ctx)
			if err != nil {
				t.ExecuteTemplate(w, "/404", templates.TemplateData{})
				return
			}
			t.ExecuteTemplate(w, "/", templates.TemplateData{
				"Title":    "Nausea",
				"HomePage": true,
				"Meta":     meta,
			})
		},
	)
}

func handleDist() http.Handler {
	return gzipMiddleware(
		http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))),
	)
}
