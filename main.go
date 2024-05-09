package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"nausea-web/internal/compress"
	"nausea-web/internal/db"
	"nausea-web/internal/dist_generator"
	"nausea-web/internal/firestore"
	"nausea-web/internal/minify"
	"nausea-web/internal/server"
	"nausea-web/internal/templates"
)

func init() {
	err := dist_generator.NewDistCreator(
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
	projectID := os.Getenv("PROJECT_ID")
	port := ":8080"
	t := templates.NewTemplate()
	f := firestore.NewFirestoreClient(projectID)
	db := db.NewDB(f)
	server := server.NewServer(t, db)
	fmt.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, server))
}
