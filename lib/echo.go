package lib

import (
	"html/template"
	"strings"
)

func echo(args []string, session *TerminalSession) error {
	out := template.HTML(strings.Join(args, " "))
	session.StdOut = append(session.StdOut, out)
	return nil
}
