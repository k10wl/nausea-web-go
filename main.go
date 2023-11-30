package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := NewTemplate()
	http.HandleFunc("/", home(tmpl))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// TODO move to env
	port := ":8080"
	fmt.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func home(tmpl *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFound(tmpl)(w, r)
			return
		}
		tmpl.ExecuteTemplate(w, "home", TemplalteData{
			"Title": "Nausea",
		})
	}
}

func notFound(tmpl *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "did we lost this page? maybe? maybe it did not exist?")
	}
}
