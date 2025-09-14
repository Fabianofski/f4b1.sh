package lib

import (
	"html/template"
	"strings"
)

type TerminalSession struct {
	StdOut       []template.HTML
	InputAllowed bool
	Cwd          string
	FileTree     []string
}

func ParseCommand(input string, session *TerminalSession) error {
	cmd := strings.Split(input, " ")
	args := cmd[1:]
	session.StdOut = append(session.StdOut, template.HTML("$ guest@f4b1.dev > "+input))
	switch cmd[0] {
	case "echo":
		echo(args, session)
	case "clear":
		clear(session)
	case "ls":
		ls(args, session)
	default:
		out := template.HTML("[f4b1.sh] command not found " + cmd[0])
		session.StdOut = append(session.StdOut, out)
	}
	return nil
}
