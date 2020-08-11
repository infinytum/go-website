package ui

import "github.com/infinytum/mustache"

var fp = &mustache.FileProvider{
	Paths:      []string{"static/templates/", "./"},
	Extensions: []string{".mustache"},
}

func RenderFileInLayout(filename string, layoutFile string, context ...interface{}) (string, error) {
	layoutTmpl, err := mustache.ParseFilePartials(layoutFile, fp)
	if err != nil {
		return "", err
	}

	tmpl, err := mustache.ParseFilePartials(filename, fp)
	if err != nil {
		return "", err
	}

	return tmpl.RenderInLayout(layoutTmpl, context...)
}
