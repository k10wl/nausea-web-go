package server

import (
	"context"
	"html/template"
	"nausea-web/internal/db"
	"nausea-web/internal/models"
	"nausea-web/internal/templates"
	"net/http"
	"time"
)

func handleHome(t *template.Template, db *db.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			meta, err := db.GetMeta(ctx)
			if err != nil {
				t.ExecuteTemplate(w, "/404", templates.TemplateData{})
				return
			}
			t.ExecuteTemplate(w, "/", templates.TemplateData{
				"Title":    "Nausea",
				"HomePage": true,
				"Meta":     meta,
			})
		},
	)
}

func handleContacts(t *template.Template, db *db.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			echan := make(chan error, 1)
			var meta *models.Meta
			mchan := make(chan *models.Meta, 1)
			var contacts *models.Contacts
			cchan := make(chan *models.Contacts, 1)
			go func() {
				c, err := db.GetContacts(ctx)
				if err != nil {
					echan <- err
					return
				}
				cchan <- c
			}()
			go func() {
				m, err := db.GetMeta(ctx)
				if err != nil {
					echan <- err
					return
				}
				mchan <- m
			}()
			var err error
			for i := 0; i < 2; i++ {
				select {
				case e := <-echan:
					err = e
					break
				case m := <-mchan:
					meta = m
				case c := <-cchan:
					contacts = c
				case <-ctx.Done():
					http.Error(w, "Request timed out", http.StatusRequestTimeout)
					return
				}
			}
			if err != nil {
				t.ExecuteTemplate(w, "/404", templates.TemplateData{})
				return
			}
			t.ExecuteTemplate(w, "/contacts", templates.TemplateData{
				"Title":    "Contacts",
				"HomePage": false,
				"Contacts": contacts,
				"Meta":     meta,
			},
			)
		})
}

func handleDist() http.Handler {
	return gzipMiddleware(
		http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))),
	)
}
