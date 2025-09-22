package lib

import (
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
)

func clear(session *model.TerminalSession) error {
	session.StdOut = []template.HTML{}
	return nil
}
