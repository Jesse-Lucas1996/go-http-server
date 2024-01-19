package main

import (
	"html/template"
	"io"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(paths ...string) *Template {
	tmp := &Template{
		Templates: template.New(""),
	}

	for _, path := range paths {
		template.Must(tmp.Templates.ParseGlob(path))
	}

	return tmp
}
