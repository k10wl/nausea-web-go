package server

import (
	"html/template"
	"nausea-web/internal/db"
	"net/http"
)

func NewServer(
	t *template.Template,
	db *db.DB,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, t, db)
	return mux
}
