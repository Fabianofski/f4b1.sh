package lib

import "html/template"

func clear(session *TerminalSession) error {
	session.StdOut = []template.HTML{}
	return nil
}
