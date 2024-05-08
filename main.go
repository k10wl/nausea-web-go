package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"nausea-web/internal/compress"
	"nausea-web/internal/dist"
	"nausea-web/internal/minify"
)

func init() {
	err := dist.NewDistCreator("./assets", minify.NewTdewolffMinifier(), compress.NewGzip()).
		RebuildDist()
	if err != nil {
		panic(err)
	}
}

func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		w.Header().Set("Content-Encoding", "gzip")
		switch ext {
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".ttf":
			w.Header().Set("Content-Type", "font/ttf")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	tmpl := NewTemplate()

	http.Handle(
		"/dist/",
		gzipMiddleware(http.StripPrefix("/dist/", http.FileServer(http.Dir("dist")))),
	)
	http.HandleFunc("/", home(tmpl))
	http.HandleFunc("/about", about(tmpl))

	// TODO move to env
	port := ":8080"
	fmt.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
