package templates

import "html/template"

type TemplateData map[string]any

func NewTemplate() *template.Template {
	tmpl := template.New("")
	tmpl.ParseGlob("./views/layouts/*.html")
	tmpl.ParseGlob("./views/404.html")
	tmpl.ParseGlob("./views/home.html")
	tmpl.ParseGlob("./views/about.html")
	return tmpl
}
