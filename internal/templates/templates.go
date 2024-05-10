package templates

import (
	"html/template"
	"nausea-web/internal/minify"
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
	tmpl = template.Must(
		minify.CompileTemplates(
			"./views/404.html",
			"./views/about.html",
			"./views/contacts.html",
			"./views/home.html",
			"./views/layouts/_head.html",
			"./views/layouts/_tail.html",
			"./views/layouts/_navigation.html",
		),
	)
	tmpl = tmpl.Funcs(template.FuncMap{"mdToLink": mdToLink})
	return tmpl
}
