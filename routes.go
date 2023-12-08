package main

import (
	"html/template"
	"net/http"
)

func home(tmpl *template.Template, store *Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			tmpl.ExecuteTemplate(w, "404", TemplalteData{
				"Title": "Nausea",
			})
			return
		}
		tmpl.ExecuteTemplate(w, "home", TemplalteData{
			"Title":    "Nausea",
			"HomePage": true,
		})
	}
}

func about(tmpl *template.Template, store *Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		tmpl.ExecuteTemplate(w, "about", TemplalteData{
			"Title": "Nausea",
			"Data":  store.GetAbout().Bio,
		})
	}
}
