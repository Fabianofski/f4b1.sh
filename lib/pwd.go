package lib

import (
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
)

func pwd(session *model.TerminalSession) error {
	session.StdOut = append(session.StdOut, template.HTML(session.Cwd))
	return nil
}
