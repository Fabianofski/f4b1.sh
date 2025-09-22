package model

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"strings"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Templates) RenderToString(name string, data any) (string, error) {
	var buf strings.Builder
	if err := t.templates.ExecuteTemplate(&buf, name, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
