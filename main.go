package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	port := ":8080"
	fmt.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func home(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello world")
}
