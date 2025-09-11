package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
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

func readBootText() []string {
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
	return strings.Split(string(b), "\n")
}

func handleTerminal(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		log.Println("Connected!!")
		defer ws.Close()
		count := 0
		bootText := readBootText()
		for {
			err := websocket.Message.Send(ws, fmt.Sprintf("<div id='terminal' hx-swap-oob='beforeend'>%s\n</div>", bootText[count]))
			if err != nil {
				c.Logger().Error(err)
			}
			time.Sleep(30 * time.Millisecond)
			count++

			if count >= len(bootText) {
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = newTemplate()

	e.Static("/static", "static")
	e.Static("/css", "css")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/terminal-output", handleTerminal)

	e.Logger.Fatal(e.Start(":4000"))
}
