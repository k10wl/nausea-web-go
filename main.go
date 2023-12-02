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
