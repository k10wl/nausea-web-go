package main

import (
	"fmt"
	"log"
	"net/http"

	"nausea-web/internal/compress"
	"nausea-web/internal/dist"
	"nausea-web/internal/minify"
	"nausea-web/internal/server"
	"nausea-web/internal/templates"
)

func init() {
	err := dist.NewDistCreator(
		"./assets",
		minify.NewTdewolffMinifier(),
		compress.NewGzip(),
	).
		RebuildDist()
	if err != nil {
		panic(err)
	}
}

func main() {
	port := ":8080"
	t := templates.NewTemplate()
	server := server.NewServer(t)
	fmt.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, server))
}
