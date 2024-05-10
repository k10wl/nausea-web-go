package templates

import (
	"html/template"
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

type TemplateData map[string]any

func NewTemplate() *template.Template {
	tmpl := template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{"mdToLink": mdToLink})
	tmpl.ParseGlob("./views/layouts/*.html")
	tmpl.ParseGlob("./views/404.html")
	tmpl.ParseGlob("./views/home.html")
	tmpl.ParseGlob("./views/about.html")
	tmpl.ParseGlob("./views/contacts.html")
	return tmpl
}
