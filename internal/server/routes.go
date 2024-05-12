package server

import (
	"html/template"
	"net/http"

	"nausea-web/internal/db"
)

func addRoutes(
	mux *http.ServeMux,
	t *template.Template,
	db *db.DB,
) {
	mux.Handle("/", notFoundMiddleware(t, db, handleHome(t, db)))
	mux.Handle("/about", handleAbout(t, db))
	mux.Handle("/contacts", handleContacts(t, db))
	mux.Handle("/dist/", handleDist())
}
