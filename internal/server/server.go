package server

import (
	"html/template"
	"net/http"
)

func NewServer(t *template.Template) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, t)
	return mux
}
