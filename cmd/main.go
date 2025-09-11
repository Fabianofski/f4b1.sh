package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

type Terminal struct {
	StdOut string
	StdErr string
	StdIn  string
}

func readBootText() string {
	file, err := os.Open("static/boot.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	return string(b)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.Static("/static", "static")

	terminal := Terminal{StdOut: "", StdErr: "", StdIn: ""}
	e.GET("/", func(c echo.Context) error {
		terminal.StdOut = readBootText()
		return c.Render(200, "index", terminal)
	})

	e.POST("/input", func(c echo.Context) error {
		terminal.StdIn = c.FormValue("stdin")
		return c.Render(200, "terminal", terminal)
	})

	e.POST("/submit", func(c echo.Context) error {
		log.Println("Submit")
		terminal.StdOut += "\n" + terminal.StdIn
		terminal.StdIn = ""
		return c.Render(200, "terminal", terminal)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
