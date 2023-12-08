package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	tmpl := NewTemplate()
	f := NewFirebaseClient(context.Background(), "personal-406007")
	store := NewStore(*f)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", home(tmpl, store))
	http.HandleFunc("/about", about(tmpl, store))

	// TODO move to env
	port := ":8080"
	fmt.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
