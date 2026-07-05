package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer() (*Renderer, error) {
	tmpl, err := template.ParseGlob(filepath.Join("web", "templates", "**", "*.html"))
	if err != nil {
		return nil, err
	}

	return &Renderer{
		templates: tmpl,
	}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data any) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return r.templates.ExecuteTemplate(w, name, data)
}
