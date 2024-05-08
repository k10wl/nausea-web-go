package server

import (
	"html/template"
	"nausea-web/internal/templates"
	"net/http"
)

func handleHome(t *template.Template) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t.ExecuteTemplate(w, "/", templates.TemplateData{
				"Title":    "Nausea",
				"HomePage": true,
			})
		},
	)
}

func handleDist() http.Handler {
	return gzipMiddleware(
		http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))),
	)
}
