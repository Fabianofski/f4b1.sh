package lib

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/Fabianofski/f4b1.sh/model"
)

func ParseCommand(input string, session *model.TerminalSession) error {
	cmd := strings.Split(strings.Trim(input, " "), " ")
	args := cmd[1:]
	session.StdOut = append(session.StdOut, template.HTML(fmt.Sprintf("$ %s guest@f4b1.dev > %s", session.CwdShort, input)))
	switch cmd[0] {
	case "echo":
		echo(args, session)
	case "clear":
		clear(session)
	case "ls":
		ls(args, session)
	case "cat":
		cat(args, session)
	case "cd":
		cd(args, session)
	case "pwd":
		pwd(session)
	case "":
		return nil
	default:
		out := template.HTML("[f4b1.sh] command not found " + cmd[0])
		session.StdOut = append(session.StdOut, out)
	}
	return nil
}
