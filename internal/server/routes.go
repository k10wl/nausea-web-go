package server

import (
	"html/template"
	"net/http"
)

func addRoutes(mux *http.ServeMux, t *template.Template) {
	mux.Handle("/", notFoundMiddleware(t, routeLoggerMiddleware(handleHome(t))))
	mux.Handle("/dist/", handleDist())
}
