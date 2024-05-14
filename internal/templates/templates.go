package templates

import (
	"html/template"
	"nausea-web/internal/models"
	"regexp"
)

func mdToLink(md string) template.HTML {
	regexPattern := `\[([^\]]+)\]\(([^)]+)\)`
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		panic(err)
	}
	return template.HTML(
		re.ReplaceAllString(md, `<a class="contacts__link" href="$2">$1</a>`),
	)
}

type TemplateData struct {
	Title    string
	HomePage bool
	Meta     *models.Meta
	Props    map[string]interface{}
}

func NewTemplate() *template.Template {
	tmpl := template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{"mdToLink": mdToLink})
	tmpl.ParseGlob("./views/layouts/*.html")
	tmpl.ParseGlob("./views/404.html")
	tmpl.ParseGlob("./views/home.html")
	tmpl.ParseGlob("./views/about.html")
	tmpl.ParseGlob("./views/contacts.html")
	tmpl.ParseGlob("./views/gallery/index.html")
	tmpl.ParseGlob("./views/gallery/{folder_id}.html")
	return tmpl
}
