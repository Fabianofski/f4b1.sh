package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Fabianofski/f4b1.sh/lib"
	"github.com/Fabianofski/f4b1.sh/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func readBootText() []template.HTML {
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
	htmlLines := []template.HTML{}
	for _, v := range strings.Split(string(b), "\n") {
		htmlLines = append(htmlLines, template.HTML(v))
	}
	return htmlLines
}

func SendBootText(ws *websocket.Conn, templates *model.Templates, session *model.TerminalSession) error {
	count := 0
	bootText := readBootText()
	for {
		session.StdOut = bootText[:count]
		err := SendTerminalSession(ws, templates, session)
		if err != nil {
			return err
		}
		time.Sleep(30 * time.Millisecond)
		count++
		if count >= len(bootText) {
			break
		}
	}
	return nil
}

func SendTerminalSession(ws *websocket.Conn, templates *model.Templates, session *model.TerminalSession) error {
	html, err := templates.RenderToString("terminal-line", session)
	if err != nil {
		return err
	}
	err = websocket.Message.Send(ws, html)
	if err != nil {
		return err
	}
	return nil
}

func handleTerminal(c echo.Context, templates *model.Templates) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		session := &model.TerminalSession{
			Cwd:      "/home/guest/",
			CwdShort: "~",
			HomeDir:  "/home/guest/",
			Root:     DefaultFileTree,
		}

		err := SendBootText(ws, templates, session)
		if err != nil {
			c.Logger().Error(err)
		}

		session.InputAllowed = true
		err = SendTerminalSession(ws, templates, session)
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			var msg string
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				if err.Error() == "EOF" {
					c.Logger().Info("WebSocket closed by server")
					break
				}
				c.Logger().Error(err)
				continue
			}

			var m model.Message
			if err := json.Unmarshal([]byte(msg), &m); err != nil {
				c.Logger().Error(err)
				continue
			}
			log.Printf("%s\n", m.Input)
			lib.ParseCommand(m.Input, session)
			SendTerminalSession(ws, templates, session)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	templates := model.NewTemplate()
	e.Renderer = templates

	e.Static("/static", "static")
	e.Static("/css", "css")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/terminal-output", func(c echo.Context) error {
		return handleTerminal(c, templates)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
