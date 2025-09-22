package lib

import (
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
	"strings"
)

func echo(args []string, session *model.TerminalSession) error {
	out := template.HTML(strings.Join(args, " "))
	session.StdOut = append(session.StdOut, out)
	return nil
}
