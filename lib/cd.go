package lib

import (
	"fmt"
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
)

func cd(args []string, session *model.TerminalSession) error {
	if len(args) == 0 {
		session.Cwd = session.HomeDir
	} else {
		path := args[0]
		absPath := pathToAbsolutePath(path, session)
		_, ok := session.Root[absPath]

		if ok {
			session.Cwd = absPath
		} else {
			errorMsg := fmt.Sprintf("cd: cannot access %s: No such file or directory", path)
			session.StdOut = append(session.StdOut, template.HTML(errorMsg))
		}
	}
	return nil
}
